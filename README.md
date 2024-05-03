# Simple Slack Message

A minimal, static executable to send a message to a Slack channel.

The idea was to have something that can be embedded into CI/CD pipelines to send notifications.

## Usage

```
export SLACK_API_TOKEN=xoxb-abcdef123456789
# Also supports SLACK_CLI_TOKEN
echo "Changelog\n- Updated foo\n- Refactored bar" | simple-slack-message --channel "deploy" --pretext "Build Deployed"
```

## Downloading

Pre-built binary executables can be downloaded from the [Releases](https://github.com/TJC/simple-slack-message/releases) page.

## Building from source

- Install Golang
- Run `go build`


# License

Copyright 2024 Toby Corkindale

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://github.com/TJC/simple-slack-message/blob/main/LICENSE

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

