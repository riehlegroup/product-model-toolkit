package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

		if typeValue == "" {
			fmt.Println("The `type` flag is compulsory")
			return
		}
		if path == "" {
			fmt.Println("The `path` flag is compulsory")
			return

		}
		err := createBomWithType(typeValue, path)
		if err != nil {
			log.Fatalf("an error occurred: %v\n", err)
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

type BomClient struct {
	pb.BomServiceClient
	*grpc.ClientConn
}

func createGrpcClient() (*BomClient, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:56985", grpc.WithInsecure()) // TODO(change the hard coded address)
	if err != nil {
		log.Printf("Did not connect: %v\n", err)
		return nil, err
	}

	client := pb.NewBomServiceClient(conn)
	bomClient := &BomClient{
		client,
		conn,
	}
	return bomClient, nil
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
	bomClient, err := createGrpcClient()
	if err != nil {
		return err
	}
	defer bomClient.ClientConn.Close()

	normalTypeValue := normaliseTypeValue(typeValue)
	inputValue := &pb.InputValue{
		FileName: path,
		Type:     normalTypeValue,
	}

	r, err := bomClient.CreateBom(context.Background(), inputValue)
	if err != nil {
		return err
	}

	// check if the bom is not created
	if !r.Created {
		return errors.New("an error occurred during creating the BoM, the input path is invalid")
	}
	// if bom is created: store the product into the db

	// then create the spdx/human readable/custom report file
	// return the generated file location

	// TODO(change)
	log.Printf("Bom created: %t", r.Created)
	return nil
}
