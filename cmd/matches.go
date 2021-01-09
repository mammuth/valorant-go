/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/mammuth/valorant-go/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// matchesCmd represents the matches command
var matchesCmd = &cobra.Command{
	Use:   "matches",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		region := viper.Get("region")
		username := viper.Get("username")
		password := viper.Get("password")
		client, _ := api.NewClient(region.(string), username.(string), password.(string))

		matches := client.GetMatchHistory()
		fmt.Printf("\nCurrent Elo: %d\n", Cyan(matches[0].TotalElo()))
		fmt.Println("Date\t\tMap\tElo Change")
		for _, match := range matches {
			prettyPrintMatch(match)
		}
	},
}

func prettyPrintMatch(m api.Match) {
	sign := ""
	Color := Red
	if m.EloChange() > 0 {
		Color = Green
		sign = "+"

	}
	elo := sign + strconv.FormatInt(int64(m.EloChange()), 10)
	fmt.Println(m.VerboseTime()+"\t"+m.VerboseMapName()+"\t", Color(elo))
}

func init() {
	rootCmd.AddCommand(matchesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matchesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matchesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
