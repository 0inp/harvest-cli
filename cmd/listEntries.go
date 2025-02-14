/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
  c "./config"
	"github.com/spf13/viper"
)

const serverPort = 3333

// listEntriesCmd represents the listEntries command
var listEntriesCmd = &cobra.Command{
	Use:   "listEntries",
	Short: "list all Harvest entries",
	Long:  `list all Harvest entries`,
	Run: func(cmd *cobra.Command, args []string) {
    // Set the file name of the configurations file
    viper.SetConfigName("config")

    // Set the path to look for the configurations file
    viper.AddConfigPath(".")

    // Enable VIPER to read Environment Variables
    viper.AutomaticEnv()

    viper.SetConfigType("yml")
    var configuration c.Configurations

    if err := viper.ReadInConfig(); err != nil {
      fmt.Printf("Error reading config file, %s", err)
    }
    
    // Set undefined variables
    viper.SetDefault("database.dbname", "test_db")

    err := viper.Unmarshal(&configuration)
    if err != nil {
      fmt.Printf("Unable to decode into struct, %v", err)
    }

		client := http.Client{Timeout: 5 * time.Second}
		requestURL := "https://api.harvestapp.com/api/v2/time_entries?from=2025-02-04T17:32:23"
		res, err := http.Get(requestURL)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		req.Header.Set("Authorization", os.ExpandEnv(configuration.HARVEST_AUTH_TOKEN))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Harvest API Example")
		req.Header.Set("Harvest-Account-ID", configuration.HARVEST_ACCOUNT_ID)

		res, err = client.Do(req)
		if err != nil {
			panic(err)
		}

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(listEntriesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listEntriesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listEntriesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
