package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	VersionPartLeading  = "leading"
	VersionPartCenter   = "center"
	VersionPartTrailing = "trailing"
)

type ECmd struct {
	BundlePath     string
	DistPath       string
	OutputPath     string
	ConfigFileName string
	VersionPart    string
	PrintVersion   bool
}

func (c *ECmd) InitWithFlag() error {
	flag.BoolVar(&c.PrintVersion, "V", false, "Print EPB version.")
	flag.StringVar(&c.BundlePath, "p", "", "Path to needs handle package")
	flag.StringVar(&c.DistPath, "d", "", "Web dist path")
	flag.StringVar(&c.OutputPath, "o", "", "Output path")
	flag.StringVar(&c.ConfigFileName, "f", "package.json", "Config filename")
	flag.StringVar(&c.VersionPart, "v", "trailing", "Needs increase part of version. Options: leading,center,trailing")
	flag.Parse()

	if c.PrintVersion {
		fmt.Printf("EPB version %s\n", version)
	}

	if c.BundlePath == "" ||
		c.DistPath == "" ||
		c.OutputPath == "" {
		return errors.New("BundlePath(-p) and DistPath(-d) and OutputPath(-o) is required")
	}
	if c.VersionPart != VersionPartTrailing &&
		c.VersionPart != VersionPartCenter &&
		c.VersionPart != VersionPartLeading {
		return errors.New("invalid version part. options: leading,center,trailing")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if !strings.HasPrefix(c.BundlePath, "/") {
		c.BundlePath = path.Join(pwd, c.BundlePath)
	}
	if !strings.HasPrefix(c.DistPath, "/") {
		c.DistPath = path.Join(pwd, c.DistPath)
	}
	if !strings.HasPrefix(c.OutputPath, "/") {
		c.OutputPath = path.Join(pwd, c.OutputPath)
	}

	return nil
}

func (c *ECmd) FullConfigFilePath() string {
	return path.Join(c.BundlePath, c.ConfigFileName)
}
