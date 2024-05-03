package main

import (
	"flag"

	"fmt"

	"encoding/json"

	"errors"

	"io"

	"os"

	"github.com/slack-go/slack"
)

type responseInfo struct {
  Success bool `json:"success"`
  Error string `json:"error,omitempty"`
}

func errorAndExit(err error) {
    errorMsg := fmt.Sprintf("%s", err)
    r, _ := json.Marshal(responseInfo{
      Success: false,
      Error: errorMsg,
    })
    fmt.Println(string(r))
    os.Exit(1)
}

func successAndExit() {
    r, _ := json.Marshal(responseInfo{
      Success: true,
    })
    fmt.Println(string(r))
    os.Exit(0)
}

func showUsage() {
  println("echo \"hello\\n- world\\n- sky\" | simple-slack-message --channel general --pretext 'A message from your bot'")
  os.Exit(0)
}

// Gets an API token from the environment, trying to be compatible
// with some other CLI tools so this can be a drop-in replacement.
func getApiToken() string {
  api_token := os.Getenv("SLACK_API_TOKEN")
  cli_token := os.Getenv("SLACK_CLI_TOKEN")
  if (len(api_token) > 0) {
    return api_token
  }
  if (len(cli_token) > 0) {
    return cli_token
  }

  errorAndExit(errors.New("this requires either SLACK_API_TOKEN or SLACK_CLI_TOKEN to be defined in environment"))
  return ""
}

func main() {
  api_token := getApiToken()

  channel := flag.String("channel", "", "Channel to post to, without the # prefix")
  pretext := flag.String("pretext", "", "Pretext of message")
  helpFlag := flag.Bool("help", false, "Show usage information")
  versionFlag := flag.Bool("version", false, "Show version number")

  flag.Parse()

  if (*helpFlag) {
    showUsage()
  }

  if (*versionFlag) {
    fmt.Println("v0.1.0")
    os.Exit(0)
  }

  if (*channel == "") {
    errorAndExit(errors.New("missing channel"))
  }

  textBytes, err := io.ReadAll(os.Stdin)
  if (err != nil) {
    errorAndExit(err)
  }
  text := string(textBytes)


  api := slack.New(api_token)
  attachment := slack.Attachment{
    Pretext: *pretext,
    Text: text,
  }

  _, _, err = api.PostMessage(
    *channel,
    slack.MsgOptionText(text, false),
    slack.MsgOptionAttachments(attachment),
    slack.MsgOptionAsUser(true),
  )
  if err != nil {
    errorAndExit(err)
  }

  successAndExit()
}
