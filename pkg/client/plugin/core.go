// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
)

type coreConfig struct {
	RestApi            bool
	RemoteRepoUser     string
	RemoteRepoPass     string
	SaveResultsLocally bool
	PathDirResults     string
	PathDirLogs        string
}

var coreConfigValues coreConfig
var configLoaded = false

const envRemoteRepoUser = "REMOTEREPO_USER"
const envRemoteRepoPass = "REMOTEREPO_PASS"

func StartCore(tool string, regFile string, inDir string) {
	coreCfg, err := loadCoreEngineConfig()
	if err != nil {
		log.Printf("[Core] Unable to load core engine configuration: %v", err.Error())
		return
	}

	var pluginRegistry Register = LoadPluginRegistry(regFile)

	var pluginsExec []Plugin

	if tool == "all" {
		log.Printf("[Core] All plugins selected")
		pluginsExec = pluginRegistry.Available()
	}

	if tool != "all" {
		scn, found := pluginRegistry.FromStr(tool)
		if !found {
			scn = pluginRegistry.Default()
			log.Printf("[Core] Unable to find plugin with name '%s'; fallback to default plugin with name '%s'", tool, scn.Name)
		}
		if found {
			log.Printf("[Core] Selected plugin: %s %s", scn.Name, scn.Version)
		}
		pluginsExec = []Plugin{scn}
	}

	err = initializeFilestore(len(pluginsExec))
	if err != nil {
		log.Printf("[Core] Unable to initialize filestore: %v", err.Error())
		return
	}

	log.Printf("[Core] Input directory: %v", inDir)

	var wg sync.WaitGroup
	for i := 0; i < len(pluginsExec); i++ {
		wg.Add(1)
		cfg := &Config{Plugin: pluginsExec[i], InDir: inDir, ResultDir: coreCfg.PathDirResults, Id: i}
		var agent agent = createAgent(cfg)
		go func() {
			err := agent.execPlugin(&wg)
			if err != nil {
				log.Printf("[Core] Error during plugin execution: %v", err.Error())
				return
			}
		}()
	}
	wg.Wait()
}

func loadCoreEngineConfig() (*coreConfig, error) {
	if configLoaded == true {
		return &coreConfigValues, nil
	}

	handle, err := os.Open("pkg/client/plugin/config/core_engine_config.json")
	if err != nil {
		return &coreConfig{}, err
	}
	defer handle.Close()

	err = json.NewDecoder(handle).Decode(&coreConfigValues)
	if err != nil {
		return &coreConfig{}, err
	}

	if coreConfigValues.SaveResultsLocally == true && coreConfigValues.PathDirResults == "" {
		coreConfigValues.PathDirResults, err = ioutil.TempDir("", "pmtclient-*")
		if err != nil {
			log.Print("[Core] Unable to create a temporary directory for result files\nUnable to proceed")
			os.Exit(-1)
		}
	}

	if coreConfigValues.PathDirLogs == "" {
		coreConfigValues.PathDirLogs, err = ioutil.TempDir("", "pmtclient-logs-*")
		if err != nil {
			log.Print("[Core] Unable to create a temporary directory for log files\nUnable to proceed")
			os.Exit(-1)
		}
	}

	configLoaded = true

	return &coreConfigValues, nil
}

func LoadPluginRegistry(file string) *Registry {
	pluginRegistry, err := NewRegistry(file)
	if err != nil {
		log.Printf("[Core] Unable to create new plugin registry from file '%s'. Error: %s\nUnable to proceed", file, err.Error())
		os.Exit(-1)
	}
	if pluginRegistry.IsEmpty() {
		log.Print("[Core] Unable to proceed with empty plugin registry")
		os.Exit(-1)
	}
	return pluginRegistry
}

// getRemoteRepoAuth returns authentication string necessary to pull container from container registry
func getRemoteRepoAuth() (string, error) {
	user := os.Getenv(envRemoteRepoUser)
	if coreConfigValues.RemoteRepoUser != "" {
		user = coreConfigValues.RemoteRepoUser
	}
	pass := os.Getenv(envRemoteRepoPass)
	if coreConfigValues.RemoteRepoPass != "" {
		pass = coreConfigValues.RemoteRepoPass
	}

	if user == "" || pass == "" {
		log.Println("[Core] No authentication credentials for remote repository provided, please check configuration or environment variables")
		return "", errors.New("no authentication credentials provided")
	}

	authConfig := types.AuthConfig{
		Username: user,
		Password: pass,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	return authStr, nil
}

func createLogFile(pluginName string) (string, error) {
	file, err := os.Create(filepath.Join(coreConfigValues.PathDirLogs, fmt.Sprintf("pmt_container_output_%v_%v.log", pluginName, time.Now().Format("2006-01-02_15-04-05"))))
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func writeToLogFile(logFile string, cmd string, stream string, text string) error {
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = src.WriteString(fmt.Sprintf("%s of command %s\n%s\n", stream, cmd, text))
	if err != nil {
		return err
	}

	return nil
}
