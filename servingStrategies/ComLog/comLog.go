package ComLog

import "strings"

const faultyString = "<Failure>"

type CommunicationLog struct{
	Target       string
	ResponseKey  string
	ResponseBody string
}

func (cl CommunicationLog) IsResponseFaulty () bool{
	return strings.Contains(cl.ResponseBody,faultyString)
}