adopt_copilot: ## Copy agent files to ~/.copilot/agents with model overrides
	go run main.go -dst ~/.claude -system copilot

adopt_opencode: ## Copy agent files to ~/.opencode/agents with model overrides
	go run main.go -dst ~/.opencode -system opencode
