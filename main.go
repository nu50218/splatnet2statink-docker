package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// {"api_key": "", "cookie": "", "user_lang": "", "session_token": ""}

type ConfigData struct {
	APIKey       string `json:"api_key"`
	Cookie       string `json:"cookie"`
	UserLang     string `json:"user_lang"`
	SessionToken string `json:"session_token"`
}

// generate config.txt from environment variable
func prepare() error {
	if os.Getenv("USE_ENVIRONMENT_VARIABLE") != "1" {
		return nil
	}

	config := &ConfigData{
		APIKey:       os.Getenv("API_KEY"),
		Cookie:       os.Getenv("COOKIE"),
		UserLang:     os.Getenv("USER_LANG"),
		SessionToken: os.Getenv("SESSION_TOKEN"),
	}

	f, err := os.Create("config.txt")
	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.NewEncoder(f).Encode(config); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func run() int {
	cmd := exec.Command("python", append([]string{"splatnet2statink.py"}, os.Args[1:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func main() {
	if err := prepare(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(run())
}
