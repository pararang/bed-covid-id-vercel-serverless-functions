package commands

import (
	"api-bed-covid/service/scraper"
	"api-bed-covid/utils"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// GetAvailablityInProvince ...
func GetAvailablityInProvince() *cobra.Command {
	return &cobra.Command{
		Use:   "get-availablity-in-province",
		Short: "Scrap to get data availibility in selected province",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAvailablityInProvince(args[0])
		},
	}
}

func getAvailablityInProvince(id string) error {
	defer utils.TimeTrack(time.Now())

	scraperServices := scraper.New()

	provinceID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, err = scraperServices.GetProvinceAvailability(provinceID)
	if err != nil {
		return err
	}

	// fmt.Println(utils.JSONIndentFormatter(data))

	return nil
}
