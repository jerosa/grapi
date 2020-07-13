package cmd

import (
	"fmt"

	readlistgrpc "github.com/jerosa/grapi/proto"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Store a new read list",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		r := &readlistgrpc.ReadList{
			Name:   name,
			Status: readlistgrpc.ReadList_ACTIVE,
		}

		res, err := cli.Create(ctx, &readlistgrpc.CreateReadListReq{ReadList: r})
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of read list")
	_ = createCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createCmd)
}
