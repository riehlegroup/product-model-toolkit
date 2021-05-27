package cmd

import (
	//"fmt"

	"context"
	"errors"
	"google.golang.org/grpc"
	"fmt"
	"github.com/spf13/cobra"
	pb "pmt/model"
	"strings"
)

var (
	importType string
	importPath string
)


func importCommandHandler(t, p string) error {
	type Client struct {
		pb.ImportServiceClient
		*grpc.ClientConn
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:56985", grpc.WithInsecure()) // TODO(change the hard coded address)
	if err != nil {
		return err
	}

	// create a new grpc client
	importServiceClient := pb.NewImportServiceClient(conn)
	client := &Client{
		importServiceClient,
		conn,
	}

	// handle the error from grpc client phase
	if err != nil {
		return err
	}

	// close the client connection in defer
	defer client.ClientConn.Close()

	// create the import value
	importInput := &pb.ImportInput{
		Type: strings.ToLower(t),
		FilePath: strings.ToLower(p),
	}

	// call the create import function of the client
	r, err := client.CreateImport(context.Background(), importInput)

	// handle possible errors
	if err != nil {
		return err
	}

	// check if the import is done
	if r.Result.Created {
		fmt.Println("File successfully imported")
		return nil
	}

	// return the result
	return errors.New("couldn't import the file")
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

	startOfUsage := "Usage:\n"
	usageInstructions := []string{
		"  type: values are spdx, composer or hasher\n",
		"  path: path to the file\n",
	}

	usageTemplate := fmt.Sprintf("%v%v%v",
		startOfUsage,
		usageInstructions[0],
		usageInstructions[1])
	importCmd.SetUsageTemplate(usageTemplate)
}
