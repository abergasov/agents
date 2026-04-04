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
	src    = flag.String("src", "./agents", "source directory with agent markdown files")
	dst    = flag.String("dst", "", "destination directory for agents")
	system = flag.String("system", "", "target system")
)

var copilotMapper = map[string]string{
	"tech_lead":       "gpt-5.4",
	"code_reviewer":   "gpt-5.4",
	"code_researcher": "claude-sonnet-4.6",
	"code_writer":     "claude-sonnet-4.6",
	"test_engineer":   "gpt-5.4-mini",
}

func main() {
	flag.Parse()
	if dst == nil || *dst == "" {
		log.Fatal("destination directory required")
	}
	if system == nil || *system == "" {
		log.Fatal("target system required")
	}

	if err := run(*src, *dst, *system); err != nil {
		log.Fatal("failed to run", err)
	}
}

func run(srcDir, dstDir, system string) error {
	println(fmt.Sprintf("Copying %s to %s", srcDir, dstDir))
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
		model := modelForAgent(agentName)

		content, errR := os.ReadFile(srcPath)
		if errR != nil {
			return fmt.Errorf("read %q: %w", srcPath, errR)
		}
		updated := prepareContent(content, model, system)
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

func prepareContent(content []byte, model, system string) []byte {
	if system == "copilot" {
		res := bytes.ReplaceAll(content, []byte("model_placeholder"), []byte(model))
		res = bytes.ReplaceAll(res, []byte("memory: user"), []byte(""))
		return res
	}
	return content
}

func modelForAgent(agentName string) string {
	if model, ok := copilotMapper[agentName]; ok {
		return model
	}
	log.Fatalf("model for agent %q not found", agentName)
	return ""
}

func normalizeAgentFileName(name, system string) string {
	return name // return as for claude, copilot can extract them
	if system == "copilot" {
		return strings.ReplaceAll(name, ".md", ".agent.md")
	}
	log.Fatal("unknown system name", system)
	return ""
}
