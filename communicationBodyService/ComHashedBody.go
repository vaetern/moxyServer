package communicationBodyService

import (
	"strings"
	"crypto/md5"
	"encoding/hex"
)

const trimOffsetStart = "<soap:Body>"
const trimOffsetEnd = "</soap:Body>"

type ComHashedBody struct {
	input  string
	Output string
}

func NewComHashedBody(initial string) (hb *ComHashedBody) {
	hb = &ComHashedBody{input: initial, Output: stripAndGetHash(initial)}
	return hb
}

func stripAndGetHash(initial string) string {

	soapBodyString := getStringInBetween(initial, trimOffsetStart, trimOffsetEnd)

	soapBodyString = removeWhitespaces(soapBodyString)

	soapBodyString = replaceQuote(soapBodyString)

	soapBodyString = removeTimestamp(soapBodyString)

	soapBodyString = removeCorellationId(soapBodyString)

	hasher := md5.New()
	hasher.Write([]byte(soapBodyString))
	return hex.EncodeToString(hasher.Sum(nil))
}

func replaceQuote(s string) string {
	return strings.Replace(s, "\"", "|", -1)
}

func removeWhitespaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func removeTimestamp(str string) string {
	return removeBetweenPattern(str, "TimeStamp=|")
}

func removeCorellationId(str string) string {
	return removeBetweenPattern(str, "CorrelationID=|")
}

func removeBetweenPattern(str string, pattern string) string {
	start := strings.Index(str, pattern)
	if start == -1 {
		return str
	}

	startOffsetEnd := start + len(pattern)

	end := strings.Index(str[startOffsetEnd:], "|")

	return str[0:start] + str[startOffsetEnd+end:]
}

func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
}
