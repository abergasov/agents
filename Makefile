adopt_copilot: ## Copy agent files to ~/.copilot/agents with model overrides
	go run main.go -src ./agents -dst ~/.claude/agents -system copilot

adopt_opencode: ## Copy agent files to ~/.opencode/agents with model overrides
	go run main.go -src ./agents -dst ~/.opencode/agents -system opencode
