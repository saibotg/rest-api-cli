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
	"strings"

	"github.com/saibot/rest-api-cli/nagios"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

type Check struct {
	url      string
	key      string
	regex    string
	username string
	password string
	authFile string
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
	checkCmd.Flags().StringVar(&check.authFile, "auth-file", "", "Path to file with auth settings")
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
		fmt.Printf("%s - error while creating request to url: %s\n", nagios.NagiosResultCriticalText, err)
		return nagios.NagiosResultCriticalCode
	}

	//create auth header, if username is set
	if c.authFile != "" {
		content, err := os.ReadFile(c.authFile)
		if err != nil {
			fmt.Printf("%s - error reading auth file %s\n", nagios.NagiosResultUnknownText, c.authFile)
			return nagios.NagiosResultUnknownCode
		}
		req.Header.Add("Authorization", strings.TrimSpace(string(content)))
	} else if c.username != "" {
		auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(c.username+":"+c.password))
		req.Header.Add("Authorization", auth)
	}

	//do the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s - error while calling the url: %s\n", nagios.NagiosResultCriticalText, err)
		return nagios.NagiosResultCriticalCode
	}
	defer resp.Body.Close()

	//check response code
	if resp.StatusCode >= 299 {
		fmt.Printf("%s - response code was %d\n", nagios.NagiosResultUnknownText, resp.StatusCode)
		return nagios.NagiosResultUnknownCode
	}

	//read the body as byte[]
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s - Error reading body: %s\n", nagios.NagiosResultCriticalText, err)
		return nagios.NagiosResultCriticalCode
	}

	//parse the body as gjson
	jsonResult := gjson.ParseBytes(body)

	//read the key
	keyResult := jsonResult.Get(c.key)

	//check if key exists
	if !keyResult.Exists() {
		fmt.Printf("%s - Key '%s' not found\n", nagios.NagiosResultUnknownText, c.key)
		return nagios.NagiosResultUnknownCode
	}

	//check if regex is valid
	re, err := regexp.Compile(c.regex)
	if err != nil {
		fmt.Printf("%s - Regex pattern is not valid\n", nagios.NagiosResultUnknownText)
		return nagios.NagiosResultUnknownCode
	}

	//check if value matches pattern
	if re.MatchString(keyResult.String()) {
		fmt.Printf("%s - Value '%s' matches the pattern\n", nagios.NagiosResultOkText, keyResult.String())
		return nagios.NagiosResultOkCode
	} else {
		fmt.Printf("%s - Value '%s' does not match the pattern.\n", nagios.NagiosResultCriticalText, keyResult.String())
		return nagios.NagiosResultCriticalCode
	}
}
