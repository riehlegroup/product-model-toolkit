package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	bomPb "pmt/bom/proto"
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
			fmt.Println("The `type` subcommand is compulsory")
			return
		}
		if path == "" {
			fmt.Println("The `path` subcommand is compulsory")
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
	bomCmd.PersistentFlags().String("type", "", "BoM type")
	bomCmd.PersistentFlags().String("path", "", "path to the input directory")
}

type BomClient struct {
	bomPb.BomServiceClient
	*grpc.ClientConn
}

func createGrpcClient() (*BomClient, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:56985", grpc.WithInsecure()) // TODO(change the hard coded address)
	if err != nil {
		log.Printf("Did not connect: %v\n", err)
		return nil, err
	}

	client := bomPb.NewBomServiceClient(conn)
	bomClient := &BomClient{
		client,
		conn,
	}
	return bomClient, nil
}

func normaliseTypeValue(typeValue string) bomPb.InputType {
	switch typeValue {
	case "0":
		return bomPb.InputType_SPDX
	case "1":
		return bomPb.InputType_HUMAN
	case "2":
		return bomPb.InputType_CUSTOM
	default:
		return bomPb.InputType_SPDX
	}
}

func createBomWithType(path, typeValue string) error {
	bomClient, err := createGrpcClient()
	if err != nil {
		return err
	}
	defer bomClient.ClientConn.Close()

	normalTypeValue := normaliseTypeValue(typeValue)
	inputValue := &bomPb.InputValue{
		FileName: path,
		Type:     normalTypeValue,
	}

	r, err := bomClient.CreateBom(context.Background(), inputValue)
	if err != nil {
		return err
	}

	// TODO(change)
	log.Printf("Created: %t", r.Created)
	return nil
}
