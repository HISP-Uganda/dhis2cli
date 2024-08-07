package datastore

import "github.com/spf13/cobra"

var DataStoreCmd = &cobra.Command{
	Use:   "datastore",
	Short: "Manage data store",
}

func init() {
	DataStoreCmd.AddCommand(ListCmd, UpdateCmd, DeleteCmd, CreateCmd)
}
