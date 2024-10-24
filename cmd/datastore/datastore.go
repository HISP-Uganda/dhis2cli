package datastore

import "github.com/spf13/cobra"

var DataStoreCmd = &cobra.Command{
	Use:   "datastore",
	Short: "Manage data store",
}

func init() {
	DataStoreCmd.AddCommand(ListCmd, UpdateCmd, DeleteCmd, CreateCmd)
	DataStoreCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "Namespace for the data store")
}
