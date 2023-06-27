/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Artifact operations",
	Long:  `A test flag`,
}

var artifactListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List artifacts",
	Long:  "List registered artifacts",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("cannot get 'url' key: %v", err)
			os.Exit(1)
		}

		apiKey, err := cmd.Flags().GetString("api-key")
		if err != nil {
			log.Fatalf("cannot get 'api-key' key: %v", err)
		}

		if url == "" {
			log.Fatalf("url is mandatory")
		}

		if apiKey == "" {
			log.Fatalf("api-key is mandatory")
		}

		req, err := http.NewRequest("GET", url+"/artifacts", nil)
		if err != nil {
			log.Fatalf("error getting artifacts list: %v", err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("failed to communicate with the server: %v", err)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("unpexpected response from the server: %d", resp.StatusCode)
		}

		defer resp.Body.Close()

		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("cannot read response body: %v", err)
		}

		var response map[string]interface{}

		err = json.Unmarshal(buf, &response)
		if err != nil {
			log.Fatalf("cannot unmarshal response: %v", err)
		}

		fmt.Println("listing artifacts")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(artifactListCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	testCmd.Flags().BoolP("wide", "w", false, "Produce wide output")
}
