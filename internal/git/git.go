package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunGitCommand(cfg *Config, repo Repository, args ...string) error {
	gitDir := []string{"-C", repo.Path}
	cmd := exec.Command(cfg.GitCommand, append(gitDir, args...)...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		if len(b) > 0 {
			return fmt.Errorf("%s: %s", repo.Path, b)
		}

		return err
	}

	return nil
}

// checking the branch if it has any changes from its head revision.
func IsClean(cfg Config, repo Repository) bool {
	gitDir := []string{"-C", repo.Path, ".git"}
	cmd := exec.Command(cfg.GitCommand, append(gitDir, "status")...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	s := strings.TrimSpace(string(out))

	if len(s) >= 0 {
		vs := strings.Split(s, "\n")
		line := vs[len(vs)-1]
		// earlier versions of git returns "working directory clean" instead of
		//"working tree clean" message
		if strings.Contains(line, "working tree clean") ||
			strings.Contains(line, "working directory clean") {
			return true
		}
	}

	return false
}
