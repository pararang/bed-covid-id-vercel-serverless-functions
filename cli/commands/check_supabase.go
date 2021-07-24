package commands

import (
	"fmt"
	"os"

	supabase "github.com/nedpals/supabase-go"
	"github.com/spf13/cobra"
)

// CheckSupabase ...
func CheckSupabase() *cobra.Command {
	return &cobra.Command{
		Use:   "check-supabase",
		Short: "check supabase connect",
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkSupabase()
		},
	}
}

func checkSupabase() error {
	const urlTest = "http://test"

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabaseDB := os.Getenv("SUPABASE_DB")

	supabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey)

	defer supabaseClient.HTTPClient.CloseIdleConnections()

	var results map[string]interface{}
	err := supabaseClient.DB.From(supabaseDB).Select("*").Single().Filter("url", "eq", urlTest).Execute(&results)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Connected")

	return nil
}
