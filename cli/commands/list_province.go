package commands

import (
	"api-bed-covid/service/rest"
	"api-bed-covid/utils"
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
	provinces := rest.MapProvinceID

	if args[0] == "object" {
		fmt.Println(utils.JSONIndentFormatter(provinces.GetListForOptions()).String())
		return nil
	}

	if args[0] == "name" {
		fmt.Println(utils.JSONIndentFormatter(provinces.GetKeys()).String())
		return nil
	}

	return nil
}
