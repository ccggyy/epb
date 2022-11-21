package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type EVersion [3]int

func (av *EVersion) InitWithString(v string) error {
	strVersions := strings.Split(strings.TrimSpace(v), ".")
	if len(strVersions) != 3 {
		return errors.New(fmt.Sprintf("invalid version '%s'", v))
	}

	for i, sv := range strVersions {
		iv, err := strconv.Atoi(sv)
		if err != nil {
			return errors.New(fmt.Sprintf("invalid '%s' in version '%s'", sv, v))
		}
		av[i] = iv
	}

	return nil
}

func (av *EVersion) ToString() string {
	return strings.Trim(strings.Replace(fmt.Sprint(*av), " ", ".", -1), "[]")
}

func (av *EVersion) Increase(idx int) error {
	if idx < 0 || idx >= len(av) {
		return errors.New("invalid idx")
	}
	av[idx] += 1
	for i := range av {
		if i > idx {
			av[i] = 0
		}
	}
	return nil
}
