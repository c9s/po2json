package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	STATE_MSGID    = iota // waiting for msgid
	STATE_MSGSTR          // waiting for msgstr
	STATE_COMPLETE        // complete state
)

type Dictionary map[string]string

func (self Dictionary) AddMessage(msgId string, msgStr string) {
	self[msgId] = msgStr
}

func (self Dictionary) RemoveMessage(msgId string) {
	delete(self, msgId)
}

func ParsePOFile(filename string) error {

	// process(filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")

	lastMsgId := []string{}
	lastMsgStr := []string{}

	state := STATE_MSGID

	commentRegExp := regexp.MustCompile("^\\s*#")
	emptyLineRegExp := regexp.MustCompile("^\\s*$")
	msgIdRegExp := regexp.MustCompile("\\^msgid\\s+\"(.*)\"")
	msgStrRegExp := regexp.MustCompile("\\^msgstr\\s+\"(.*)\"")

	for _, line := range lines {
		if commentRegExp.MatchString(line) {
			continue
		}
		if emptyLineRegExp.MatchString(line) {
			continue
		}

		if msgIdRegExp.MatchString(line) {

		} else if msgStrRegExp.MatchString(line) {

		}

		fmt.Println(line)
	}

	_ = state
	_ = lines
	_ = lastMsgId
	_ = lastMsgStr
	return nil
}
