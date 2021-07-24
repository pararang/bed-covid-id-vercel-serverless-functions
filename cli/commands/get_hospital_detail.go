package commands

import (
	"api-bed-covid/service/scraper"
	"api-bed-covid/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// GetHospitalDetail ...
func GetHospitalDetail() *cobra.Command {
	return &cobra.Command{
		Use:   "get-hospital-detail",
		Short: "Scrap to get data hospital detail",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getHospitalDetail(args[0])
		},
	}
}

func getHospitalDetail(id string) error {
	scraperServices := scraper.New()

	data, err := scraperServices.GetHospitalDetail(id)
	if err != nil {
		return err
	}

	fmt.Println(utils.JSONIndentFormatter(data))

	return nil
}
