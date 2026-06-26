package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	dst           = flag.String("dst", "", "destination directory for agents")
	system        = flag.String("system", "", "target system")
	copilotMapper = map[string]string{
		"tech_lead":       "gpt-5.4",
		"code_reviewer":   "gpt-5.4",
		"code_researcher": "claude-sonnet-4.6",
		"code_writer":     "claude-sonnet-4.6",
		"test_engineer":   "gpt-5.4-mini",
	}
	opencodeMapper = map[string]string{
		"tech_lead":       "opencode/gpt-5",
		"code_reviewer":   "opencode/gpt-5",
		"code_researcher": "claude/sonnet-4.6",
		"code_writer":     "claude/sonnet-4.6",
		"test_engineer":   "claude/haiku-4.5",
	}
	claudeMapper = map[string]string{
		"tech_lead":       "opus",
		"code_reviewer":   "opus",
		"code_researcher": "sonnet",
		"code_writer":     "sonnet",
		"test_engineer":   "haiku",
	}
	// claudeAgentConfig maps the opencode permission intent onto Claude Code
	// frontmatter. `tools` is an allowlist that enforces the read-only vs.
	// edit-capable boundary; `skills` preloads each agent's skill at startup.
	// Per-skill hard deny is not expressible in Claude frontmatter (the Skill
	// tool is all-or-nothing), so the skill list is a preload binding, not a gate.
	claudeAgentConfig = map[string]struct {
		tools  string
		skills []string
	}{
		"tech_lead": {
			tools:  "Read, Grep, Glob, Bash, Skill, Agent, TodoWrite",
			skills: []string{"grill-with-docs", "repo-research", "implementation-plan"},
		},
		"code_researcher": {
			tools:  "Read, Grep, Glob, Bash, Skill, TodoWrite",
			skills: []string{"repo-research"},
		},
		"code_writer": {
			tools:  "Read, Write, Edit, Grep, Glob, Bash, Skill, TodoWrite",
			skills: []string{"scoped-implementation"},
		},
		"code_reviewer": {
			tools:  "Read, Grep, Glob, Bash, Skill, TodoWrite",
			skills: []string{"code-review"},
		},
		"test_engineer": {
			tools:  "Read, Write, Edit, Grep, Glob, Bash, Skill, TodoWrite",
			skills: []string{"golang-tests"},
		},
	}
	modelMapper = map[string]map[string]string{
		"copilot":  copilotMapper,
		"opencode": opencodeMapper,
		"claude":   claudeMapper,
	}
)

func main() {
	flag.Parse()
	if dst == nil || *dst == "" {
		log.Fatal("destination directory required")
	}
	if system == nil || *system == "" {
		log.Fatal("target system required")
	}

	if err := copyAgents("./agents", *dst, *system); err != nil {
		log.Fatal("failed to run agents", err)
	}
	if err := copySkills("./skills", *dst, *system); err != nil {
		log.Fatal("failed to run skills", err)
	}
}

func copySkills(srcDir, dstDir, system string) error {
	dstDir += "/skills"
	println(fmt.Sprintf("Copying skills %s to %s", srcDir, dstDir))
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	if err = os.RemoveAll(dstDir); err != nil {
		return fmt.Errorf("remove destination dir %q: %w", dstDir, err)
	}
	if err = os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("create destination dir %q: %w", dstDir, err)
	}

	var copied int
	for _, entry := range entries {
		if entry.IsDir() {
			skillName := entry.Name()
			srcPath := filepath.Join(srcDir, skillName)
			dstPath := filepath.Join(dstDir, skillName)
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}

			fmt.Printf("%s -> %s\n", srcPath, dstPath)
			copied++
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}
		srcPath := filepath.Join(srcDir, name)
		skillName := strings.TrimSuffix(name, ".md")
		content, errR := os.ReadFile(srcPath)
		if errR != nil {
			return fmt.Errorf("read %q: %w", srcPath, errR)
		}
		if err := os.MkdirAll(filepath.Join(dstDir, skillName), 0o755); err != nil {
			return fmt.Errorf("mkdir %q: %w", filepath.Join(dstDir, skillName), err)
		}
		dstPath := filepath.Join(dstDir, skillName, "SKILL.md")
		if err = os.WriteFile(dstPath, content, fs.FileMode(0o644)); err != nil {
			return fmt.Errorf("write %q: %w", dstPath, err)
		}

		fmt.Printf("%s -> %s\n", srcPath, dstPath)
		copied++
	}
	if copied == 0 {
		return fmt.Errorf("no skills found in %q", srcDir)
	}

	fmt.Printf("done: %d skill(s) adopted\n", copied)
	return nil
}

func copyDir(srcDir, dstDir string) error {
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("create destination dir %q: %w", dstDir, err)
	}
	return filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("relative path for %q: %w", path, err)
		}
		dstPath := filepath.Join(dstDir, relPath)
		if d.IsDir() {
			if relPath == "." {
				return nil
			}
			if err := os.MkdirAll(dstPath, 0o755); err != nil {
				return fmt.Errorf("mkdir %q: %w", dstPath, err)
			}
			return nil
		}
		return copyFile(path, dstPath)
	})
}

func copyFile(srcPath, dstPath string) error {
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("read %q: %w", srcPath, err)
	}
	if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
		return fmt.Errorf("mkdir %q: %w", filepath.Dir(dstPath), err)
	}
	if err := os.WriteFile(dstPath, content, fs.FileMode(0o644)); err != nil {
		return fmt.Errorf("write %q: %w", dstPath, err)
	}
	return nil
}

func copyAgents(srcDir, dstDir, system string) error {
	dstDir += "/agents"
	println(fmt.Sprintf("Copying agents %s to %s", srcDir, dstDir))
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("read source dir %q: %w", srcDir, err)
	}

	if err = os.RemoveAll(dstDir); err != nil {
		return fmt.Errorf("remove destination dir %q: %w", dstDir, err)
	}
	if err = os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("create destination dir %q: %w", dstDir, err)
	}

	var copied int
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		srcPath := filepath.Join(srcDir, name)
		agentName := strings.TrimSuffix(name, ".md")
		model := modelForAgent(agentName, system)

		content, errR := os.ReadFile(srcPath)
		if errR != nil {
			return fmt.Errorf("read %q: %w", srcPath, errR)
		}
		updated := prepareContent(content, agentName, model, system)
		dstName := normalizeAgentFileName(name, system)
		dstPath := filepath.Join(dstDir, dstName)

		if err = os.WriteFile(dstPath, updated, fs.FileMode(0o644)); err != nil {
			return fmt.Errorf("write %q: %w", dstPath, err)
		}

		fmt.Printf("%s -> %s (%s)\n", srcPath, dstPath, model)
		copied++
	}

	if copied == 0 {
		return fmt.Errorf("no .md agent files found in %q", srcDir)
	}

	fmt.Printf("done: %d agent(s) adopted\n", copied)
	return nil
}

func prepareContent(content []byte, agentName, model, system string) []byte {
	res := bytes.ReplaceAll(content, []byte("model_placeholder"), []byte(model))
	res = bytes.ReplaceAll(res, []byte("memory: user"), []byte(""))
	if system == "opencode" {
		res = bytes.ReplaceAll(res, []byte("permissions:"), []byte("permission:"))
	}
	if system == "claude" {
		res = claudeFrontmatter(res, agentName)
	}
	return res
}

// claudeFrontmatter rewrites the opencode frontmatter into valid Claude Code
// agent frontmatter. It drops keys Claude does not understand (agent, context,
// mode) and the opencode permissions/tools blocks, then injects a `tools`
// allowlist and `skills` preload from claudeAgentConfig — so the source `tools`
// block (opencode delegation gate) is rebuilt here, not passed through. The
// tools list enforces the read-only vs. edit-capable boundary; the skills list
// preloads each agent's skill (per-skill hard deny is not expressible in Claude
// frontmatter).
func claudeFrontmatter(content []byte, agentName string) []byte {
	text := string(content)
	if !strings.HasPrefix(text, "---\n") {
		return content
	}
	end := strings.Index(text[4:], "\n---")
	if end == -1 {
		return content
	}
	fmEnd := 4 + end
	front := text[4:fmEnd]
	rest := text[fmEnd:]

	var out []string
	skipBlock := false
	for _, line := range strings.Split(front, "\n") {
		trimmed := strings.TrimSpace(line)
		indented := len(line) > 0 && (line[0] == ' ' || line[0] == '\t')
		if skipBlock {
			if indented || trimmed == "" {
				continue
			}
			skipBlock = false
		}
		if trimmed == "" {
			continue
		}
		if !indented {
			key := trimmed
			if i := strings.Index(key, ":"); i != -1 {
				key = key[:i]
			}
			switch key {
			case "agent", "context", "mode":
				continue
			case "permissions", "tools":
				skipBlock = true
				continue
			}
		}
		out = append(out, line)
	}

	if cfg, ok := claudeAgentConfig[agentName]; ok {
		out = append(out, "tools: "+cfg.tools)
		if len(cfg.skills) > 0 {
			out = append(out, "skills:")
			for _, s := range cfg.skills {
				out = append(out, "  - "+s)
			}
		}
	}

	return []byte("---\n" + strings.Join(out, "\n") + rest)
}

func modelForAgent(agentName, system string) string {
	modelMapperForSystem, ok := modelMapper[system]
	if !ok {
		log.Fatalf("model mapper for system %q not found", system)
	}
	if model, ok := modelMapperForSystem[agentName]; ok {
		return model
	}
	log.Fatalf("model for agent %q not found", agentName)
	return ""
}

func normalizeAgentFileName(name, system string) string {
	return name
}
