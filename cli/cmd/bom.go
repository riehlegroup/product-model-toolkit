package cmd

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"log"
	pb "pmt/model"
)

// bomCmd represents the bom command
var bomCmd = &cobra.Command{
	Use:   "bom",
	Short: "Generate bill of materials of your software",
	Long:  `This command generates BoM artifacts as SPDX, human readable or custom reports.`,
	Run: func(cmd *cobra.Command, args []string) {
		typeValue, _ := cmd.Flags().GetString("type")
		path, _ := cmd.Flags().GetString("path")

		err := createBomWithType(typeValue, path)
		if err != nil {
			log.Fatalf("error, %v\n", err)
			return
		}
	},
}



func init() {
	rootCmd.AddCommand(bomCmd)
	bomCmd.SetUsageTemplate("The `type` and `path` flags are compulsory.\n" +
		"`type` can be:\n" +
		"0 which means SPDX\n" +
		"1 which means Human Readable\n" +
		"2 which means Custom Reports\n")
	bomCmd.PersistentFlags().String("type", "", "BoM type")
	bomCmd.PersistentFlags().String("path", "", "path to the input directory")
}


func normaliseTypeValue(typeValue string) pb.InputType {
	switch typeValue {
	case "0":
		return pb.InputType_SPDX
	case "1":
		return pb.InputType_HUMAN
	case "2":
		return pb.InputType_CUSTOM
	default:
		return pb.InputType_SPDX
	}
}

func createBomWithType(path, typeValue string) error {
	client, err := createGrpcClient()
	if err != nil {
		return err
	}
	defer client.ClientConn.Close()

	normalTypeValue := normaliseTypeValue(typeValue)
	inputValue := &pb.InputValue{
		FilePath: path,
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
