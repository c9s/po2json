package po

import (
	"encoding/json"
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

func (self Dictionary) String() string {
	jsonBytes, err := json.MarshalIndent(self, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

func ParseFile(filename string) (*Dictionary, error) {

	// process(filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")

	lastMsgId := []string{}
	lastMsgStr := []string{}

	dictionary := Dictionary{}

	state := STATE_MSGID

	commentRegExp := regexp.MustCompile("^\\s*#")
	emptyLineRegExp := regexp.MustCompile("^\\s*$")
	msgIdRegExp := regexp.MustCompile("^msgid\\s+\"(.*)\"")
	msgStrRegExp := regexp.MustCompile("^msgstr\\s+\"(.*)\"")
	stringRegExp := regexp.MustCompile("\"(.*)\"")

	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' ||
			commentRegExp.MatchString(line) ||
			emptyLineRegExp.MatchString(line) {
			continue
		}

		if strings.HasPrefix(line, "msgid") || msgIdRegExp.MatchString(line) {
			if len(lastMsgId) > 0 && len(lastMsgStr) > 0 {
				// push to the dictionary
				dictionary.AddMessage(strings.Join(lastMsgId, ""), strings.Join(lastMsgStr, ""))
				lastMsgId = []string{}
				lastMsgStr = []string{}
			}

			state = STATE_MSGID
			msgId := msgIdRegExp.FindStringSubmatch(line)[1]
			lastMsgId = append(lastMsgId, msgId)

		} else if strings.HasPrefix(line, "msgstr") || msgStrRegExp.MatchString(line) {
			state = STATE_MSGSTR
			msgStr := msgStrRegExp.FindStringSubmatch(line)[1]
			lastMsgStr = append(lastMsgStr, msgStr)
		} else if stringRegExp.MatchString(line) {
			var str = stringRegExp.FindStringSubmatch(line)[1]
			if state == STATE_MSGID {
				lastMsgId = append(lastMsgId, str)
			} else if state == STATE_MSGSTR {
				lastMsgStr = append(lastMsgStr, str)
			}
		}
	}

	return &dictionary, nil
}
