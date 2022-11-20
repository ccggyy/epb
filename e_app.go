package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type EApp struct {
	Layout       int    `json:"layout"`
	Identifier   string `json:"identifier"`
	Hidden       bool   `json:"hidden"`
	Build        int    `json:"build"`
	ProgressFlag bool   `json:"progressFlag"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ReleaseNote  string `json:"releaseNote"`
	Version      string `json:"version"`
	KeyWord      string `json:"keyWord"`
}

func (a *EApp) increaseVersion(idx int) error {
	av := EVersion{0, 0, 0}
	err := av.InitWithString(a.Version)
	if err != nil {
		return err
	}

	err = av.Increase(idx)
	if err != nil {
		return err
	}

	a.Version = av.ToString()
	return nil
}

func (a *EApp) InitWithJsonFile(fp string) error {
	bytes, err := os.ReadFile(fp)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &a)
}

func (a *EApp) IncrBuilds() {
	a.Build += 1
}

func (a *EApp) IncreaseLeadingVersion() error {
	return a.increaseVersion(0)
}

func (a *EApp) IncreaseCenterVersion() error {
	return a.increaseVersion(1)
}

func (a *EApp) IncreaseTrailingVersion() error {
	return a.increaseVersion(2)
}

func (a *EApp) WriteToFile(path string) error {
	bytes, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0622)
}

func (a *EApp) OutputZipFileName() string {
	return fmt.Sprintf("%s-%s-%s-%d.zip", a.Name, a.Identifier, a.Version, a.Build)
}
