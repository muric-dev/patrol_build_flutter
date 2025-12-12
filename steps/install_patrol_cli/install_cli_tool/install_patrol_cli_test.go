package install_cli_tool

import (
	"errors"
	"os"
	"testing"

	"patrol_install/commands"
)

func resetPatrolCLIVersionEnv() {
	os.Unsetenv("CUSTOM_PATROL_CLI_VERSION")
}

func TestInstallPatrolCLI_LatestVersion(t *testing.T) {
	resetPatrolCLIVersionEnv()
	called := false
	executor := func(cmd commands.Command) (string, error) {
		called = true
		if cmd.Name != commands.PatrolInstall.Name {
			t.Errorf("expected command name %q, got %q", commands.PatrolInstall.Name, cmd.Name)
		}
		return "installed latest", nil
	}
	output, err := InstallPatrolCLI(executor)
	resetPatrolCLIVersionEnv()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output != "installed latest" {
		t.Errorf("expected output 'installed latest', got %q", output)
	}
	if !called {
		t.Error("executor was not called")
	}
}

func TestInstallPatrolCLI_CustomVersion(t *testing.T) {
	os.Setenv("CUSTOM_PATROL_CLI_VERSION", "1.2.3")
	called := false
	executor := func(cmd commands.Command) (string, error) {
		called = true
		if len(cmd.Args) == 0 || cmd.Args[len(cmd.Args)-1] != "1.2.3" {
			t.Errorf("expected custom version in args, got %v", cmd.Args)
		}
		return "installed custom", nil
	}
	output, err := InstallPatrolCLI(executor)
	resetPatrolCLIVersionEnv()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output != "installed custom" {
		t.Errorf("expected output 'installed custom', got %q", output)
	}
	if !called {
		t.Error("executor was not called")
	}
}

func TestInstallPatrolCLI_Error(t *testing.T) {
	os.Setenv("CUSTOM_PATROL_CLI_VERSION", "")
	executor := func(cmd commands.Command) (string, error) {
		return "", errors.New("install failed")
	}
	output, err := InstallPatrolCLI(executor)
	resetPatrolCLIVersionEnv()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if output != "" {
		t.Errorf("expected empty output, got %q", output)
	}
}
