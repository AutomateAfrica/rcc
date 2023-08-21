package common

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	defaultRobocorpLocation = "%LOCALAPPDATA%\\robocorp"
	defaultHoloLocation     = "c:\\ProgramData\\robocorp\\ht"
)

var (
	variablePattern = regexp.MustCompile("%[a-zA-Z]+%")
)

func ExpandPath(entry string) string {
	intermediate := os.ExpandEnv(entry)
	intermediate = variablePattern.ReplaceAllStringFunc(intermediate, fromEnvironment)
	result, err := filepath.Abs(intermediate)
	if err != nil {
		return intermediate
	}
	return result
}

func fromEnvironment(form string) string {
	replacement, ok := os.LookupEnv(form[1 : len(form)-1])
	if ok {
		return replacement
	}
	replacement, ok = os.LookupEnv(form)
	if ok {
		return replacement
	}
	return form
}

func GenerateKillCommand(keys []int) string {
	command := []string{"taskkill /f"}
	for _, key := range keys {
		command = append(command, fmt.Sprintf("/pid %d", key))
	}
	return strings.Join(command, " ")
}
