/*
Copyright © 2023 Tobias Grotheer <tobias@grotheer-web.de>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

type Check struct {
	url      string
	key      string
	regex    string
	username string
	password string
}

var check = Check{}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run a check agains the rest api",
	Long:  `Run a nagios compatible check against the api endpoint.`,
	Run:   runCheck,
}

func init() {
	checkCmd.Flags().StringVarP(&check.url, "url", "U", "", "The url to the endpoint, that should be checked")
	checkCmd.Flags().StringVarP(&check.key, "key", "K", "", "The key to the json entry, that should be checked")
	checkCmd.Flags().StringVarP(&check.regex, "regex", "R", "", "The regex to check the json value represented by the key")
	checkCmd.Flags().StringVar(&check.username, "username", "", "Username for basic auth")
	checkCmd.Flags().StringVar(&check.password, "password", "", "Password for basic auth")
	rootCmd.AddCommand(checkCmd)
}

func runCheck(cmd *cobra.Command, args []string) {
	result := check.Execute()
	os.Exit(result)
}

func (c Check) Execute() int {
	//create request
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		fmt.Printf("CRITICAL - error while creating request to url: %s\n", err)
		return 2
	}

	//create auth header, if username is set
	if c.username != "" {
		auth := base64.StdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
		req.Header.Add("Authorization", "Basic "+auth)
	}

	//do the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("CRITICAL - error while calling the url: %s\n", err)
		return 2
	}
	defer resp.Body.Close()

	//check response code
	if resp.StatusCode >= 299 {
		fmt.Printf("UNKNOWN - response code was %d\n", resp.StatusCode)
		return 3
	}

	//read the body as byte[]
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("CRITICAL - Error reading body: %s\n", err)
		return 2
	}

	//parse the body as gjson
	jsonResult := gjson.ParseBytes(body)

	//read the key
	keyResult := jsonResult.Get(c.key)

	//check if key exists
	if !keyResult.Exists() {
		fmt.Printf("UNKNOWN - Key '%s' not found\n", c.key)
		return 3
	}

	//check if regex is valid
	re, err := regexp.Compile(c.regex)
	if err != nil {
		fmt.Printf("UNKNOWN - Regex pattern is not valid\n")
		return 3
	}

	//check if value matches pattern
	if re.MatchString(keyResult.String()) {
		fmt.Printf("OK - Value '%s' matches the pattern\n", keyResult.String())
		return 1
	} else {
		fmt.Printf("CRITICAL - Value '%s' does not match the pattern.\n", keyResult.String())
		return 2
	}
}
