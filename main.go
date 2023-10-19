// Package main is the entry point of the application
package main

import (
	"github.com/mirshahriar/marketplace/cmd"
	_ "github.com/mirshahriar/marketplace/docs"
	_ "github.com/spf13/viper/remote"
)

func main() {
	cmd.Execute()
}
