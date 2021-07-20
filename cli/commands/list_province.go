package commands

import (
	"api-bed-covid/cli/utils"
	"api-bed-covid/service"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// ListProvinces ...
func ListProvinces() *cobra.Command {
	return &cobra.Command{
		Use:   "list-province",
		Short: "Get list provinces",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getListProvince(args)
		},
	}
}

func getListProvince(args []string) error {
	log.Println(args)
	provinces := service.MapProvinceID

	if args[0] == "object" {
		fmt.Println(utils.JSONIndent(provinces.GetListForOptions()).String())
		return nil
	}

	if args[0] == "name" {
		fmt.Println(utils.JSONIndent(provinces.GetKeys()).String())
		return nil
	}

	return nil
}
