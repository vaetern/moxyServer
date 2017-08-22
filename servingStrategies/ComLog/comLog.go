package ComLog

import "strings"

const failureString = "<Failure>"
const errorsString = "<Errors>"

type CommunicationLog struct{
	Target       string
	ResponseKey  string
	ResponseBody string
}

func (cl CommunicationLog) IsResponseFaulty () bool{
	return strings.Contains(cl.ResponseBody, failureString) || strings.Contains(cl.ResponseBody, errorsString)
}