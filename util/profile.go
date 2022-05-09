package util

import (
	"io/ioutil"
	"os"
	"strings"
)

// Append prefix and suffix to the profile.
func appendPrefixSuffix(profile string, prefix string, suffix string) {
	file, err := os.Open(profile)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	cb, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fileLine := strings.Split(string(cb), "\n")
	beginAt := -1
	endAt := -1

	for i, line := range fileLine {
		if strings.TrimSpace(line) == prefix && beginAt == -1 {
			beginAt = i
			continue
		}

		if strings.TrimSpace(line) == suffix && endAt == -1 {
			endAt = i
			continue
		}
	}

	if beginAt == -1 && endAt == -1 {
		fileLine = append(fileLine, prefix, suffix)
	}

	if beginAt == -1 && endAt != -1 {
		panic("no openning prefix")
	}

	if beginAt != -1 && endAt == -1 {
		panic("no closing suffix")
	}

	writeToFile(profile, fileLine)
}

func writeToFile(f string, content []string) {
	file, err := os.OpenFile(f, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	c := strings.Join(content, "\n")

	if _, err = file.WriteString(c); err != nil {
		panic(err)
	}
}

// SaveToProfile saves the session to the profile.
func SaveToProfile(f string, content []string, prefix, suffix string, force bool) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if force == true {
		appendPrefixSuffix(f, prefix, suffix)
	}

	cb, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fileLine := strings.Split(string(cb), "\n")

	// delete lines of the file inside prefix and suffix
	beginAt := -1
	endAt := -1
	for i, line := range fileLine {
		if beginAt == -1 && strings.TrimSpace(line) == prefix {
			beginAt = i
		}

		if endAt == -1 && strings.TrimSpace(line) == suffix {
			endAt = i
		}
	}

	// append content to the file inside prefix line and suffix line
	fileLine = append(fileLine[:beginAt+1], append(content, fileLine[endAt:]...)...)
	writeToFile(f, fileLine)
}
