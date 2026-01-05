package agent

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

// CodexAgent runs code reviews using the Codex CLI
type CodexAgent struct {
	Command string // The codex command to run (default: "codex")
}

// NewCodexAgent creates a new Codex agent
func NewCodexAgent(command string) *CodexAgent {
	if command == "" {
		command = "codex"
	}
	return &CodexAgent{Command: command}
}

func (a *CodexAgent) Name() string {
	return "codex"
}

func (a *CodexAgent) Review(ctx context.Context, repoPath, commitSHA, prompt string) (string, error) {
	// Use codex review subcommand for non-interactive review
	args := []string{
		"review",
		"--commit", commitSHA,
		"-c", `model_reasoning_effort="high"`,
	}

	cmd := exec.CommandContext(ctx, a.Command, args...)
	cmd.Dir = repoPath

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("codex failed: %w\nstderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

func init() {
	Register(NewCodexAgent(""))
}
