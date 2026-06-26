adopt_copilot: ## Copy agent files to ~/.claude with model overrides
	go run main.go -dst ~/.claude -system copilot

adopt_opencode: ## Copy agent files to ~/.config/opencode with model overrides
	go run main.go -dst ~/.config/opencode -system opencode

adopt_claude: ## Copy agent files to ~/.claude with model overrides
	go run main.go -dst ~/.claude -system claude

all: adopt_copilot adopt_opencode adopt_claude
