// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type artifact struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Hash  hash   `json:"hashes"`
}

type hash struct {
	MD5    string `json:"md5,omitempty"`
	SHA1   string `json:"sha1,omitempty"`
	SHA256 string `json:"sha256,omitempty"`
}

func main() {
	inputDir, resultFile := getFlagArguments()
	fmt.Printf("Input dir: %s\nResult file: %s", inputDir, resultFile)

	artifacts, err := getAllArtifacts(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	calcHashForEachFile(artifacts)

	for _, item := range artifacts {
		fmt.Println(item.Path)
	}

	fmt.Printf("\nTotal artifacts %v\n", len(artifacts))

	writeAsJSONFile(artifacts, resultFile)
}

func getAllArtifacts(inputDir string) ([]artifact, error) {
	artifacts := make([]artifact, 0, 30)
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		a := artifact{
			Path:  path,
			Name:  info.Name(),
			IsDir: info.IsDir(),
		}

		artifacts = append(artifacts, a)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return artifacts, nil
}

func calcHashForEachFile(artifacts []artifact) {
	for index, item := range artifacts {
		if item.IsDir {
			continue
		}
		hash, err := calcHash(item.Path)
		if err != nil {
			continue
		}

		artifacts[index].Hash = hash
	}
}

func calcHash(filePath string) (hash, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return hash{}, err
	}

	return hash{
		MD5:    md5AsStr(data),
		SHA1:   sha1AsStr(data),
		SHA256: sha256AsStr(data),
	}, nil
}

func md5AsStr(data []byte) string {
	md5InBytes := md5.Sum(data)
	return hex.EncodeToString(md5InBytes[:])
}

func sha1AsStr(data []byte) string {
	sha1InBytes := sha1.Sum(data)
	return hex.EncodeToString(sha1InBytes[:])
}

func sha256AsStr(data []byte) string {
	sha256InBytes := sha256.Sum256(data)
	return hex.EncodeToString(sha256InBytes[:])
}

func (a *artifact) toString() string {
	fileInfo := fmt.Sprintf("Path: %v\nFilename: %v\nisDir: %v", a.Path, a.Name, a.IsDir)
	if a.IsDir {
		return fileInfo
	}

	hashInfo := fmt.Sprintf("\nmd5:    %s\nsha1:   %s\nsha256: %s", a.Hash.MD5, a.Hash.SHA1, a.Hash.SHA256)
	return fileInfo + hashInfo
}

func writeAsJSONFile(artifacts []artifact, filename string) {
	jsonString, _ := json.MarshalIndent(artifacts, "", "  ")
	ioutil.WriteFile(filename, jsonString, os.ModePerm)
}

func getFlagArguments() (string, string) {
	inputDir := flag.String("i", ".", "input folder to scan")
	resultFile := flag.String("o", "result.json", "result file")
	flag.Parse()
	return *inputDir, *resultFile
}
