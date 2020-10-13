package main

import (
	"context"
	"fmt"
	"github.com/tkandal/oauth2autoconf"
	"net/url"
	"os"
	"time"
)

func main() {
	if err := realMain(); err != nil {
		os.Exit(1)
	}
}

func realMain() error {
	u := os.Getenv("CONF_URL")
	if len(u) == 0 {
		return fmt.Errorf("CONF_URL is empty")
	}
	if _, err := url.Parse(u); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cfg, err := oauth2autoconf.Get(ctx, u)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s", cfg)
	return nil
}
