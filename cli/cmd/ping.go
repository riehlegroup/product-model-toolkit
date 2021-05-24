package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the service",
	Long: `This command checks the health of the system`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := createGrpcClient()
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		defer client.ClientConn.Close()
		fmt.Println("Ok")
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
