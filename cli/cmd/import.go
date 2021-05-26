package cmd

import (
	//"fmt"

	"context"
	"errors"
	"log"

	"fmt"
	"github.com/spf13/cobra"
	pb "pmt/model"
)

var (
	importType string
	importPath string
)

func importCommandHandler(t, p string) error {
	client, err := createGrpcClient()
	if err != nil {
		return err
	}
	defer client.ClientConn.Close()
	normalTypeValue := normaliseTypeValue(t)
	inputValue := &pb.InputValue{
		FilePath: importPath,
		Type:     normalTypeValue,
	}

	r, err := client.CreateBom(context.Background(), inputValue)
	if err != nil {
		return err
	}


	// check if the bom is not created
	if !r.Result.Created {
		return errors.New("an error occurred during creating the BoM, the input path is invalid")
	}
	// if bom is created: store the product into the db

	// then create the spdx/human readable/custom report file
	// return the generated file location

	// TODO(change)
	log.Printf("Bom created: %t", r.Result.Created)
	return nil

}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import product model information",
	Long: `This command will import information from a SPDX, a composer, and a hasher file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(importType)
		fmt.Println(importPath)
		err := importCommandHandler(importType, importPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func makeFlagsRequired(list ...string) {
	for _, v := range list {
		importCmd.MarkPersistentFlagRequired(v)
	}
}

func init() {
	rootCmd.AddCommand(importCmd)

	// define the flags
	importCmd.PersistentFlags().StringVarP(&importType, "type", "t", "", "Specify the type of import")
	importCmd.PersistentFlags().StringVarP(&importPath, "path", "p", "", "Path to the import file")

	// make the flags required
	makeFlagsRequired("type", "path")

	startOfUsage := "\n================= USAGE ====================\n"
	endOfUsage := "==========================================\n"
	usageInstructions := []string{
		"`type` values are spdx, composer or hasher\n",
		"`path` is the path to the required file\n",
	}

	usageTemplate := fmt.Sprintf("%v%v%v%v",
		startOfUsage,
		usageInstructions[0],
		usageInstructions[1],
		endOfUsage)
	importCmd.SetUsageTemplate(usageTemplate)
}
