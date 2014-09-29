package gooozie

import (
	"encoding/json"
)

const (
	jobsResource = ""
	jobsResource = "/jobs?jobtype="
)

func bundles() {
	resource := "/jobs?jobtype=bundle"
}

func workflows() {
	resource := "/jobs?jobtype=wf"
}

func coordinators() {
	resource := "/jobs?jobtype=coord"
}

func jobString(jobId string, queryType string) (res string) {
	res = "/job/" + jobId + "?show=" + queryType
	return res
}

func log(jobId string) {
	//resource := "/job/" + jobId + "?show=log"
	resource := jobString(jobId, "log")
}

func definition(jobId string) {
	//resource := "/job/" + jobId + "?show=definition"
	resource := jobString(jobId, "definition")
}

func info(jobId string) {
	//resource := "/job/" + jobId + "?show=info"
	resource := jobString(jobId, "info")
}
