package apps

import (
	"dhis2cli/client"
	"fmt"
	"github.com/spf13/cobra"
)

var appKey string
var DeleteAppCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an app",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := client.Dhis2Client.DeleteResource(fmt.Sprintf("apps/%v", appKey))
		if err != nil {
			// Log failed to delete
			fmt.Printf("Failed to delete app: %v\n", err)
			return
		}
		fmt.Println("App deleted successfully:")
	},
}

func init() {
	DeleteAppCmd.Flags().StringVarP(&appKey, "app-key", "", "", "App Key for App to delete")
	_ = DeleteAppCmd.MarkFlagRequired("app-key")
}
