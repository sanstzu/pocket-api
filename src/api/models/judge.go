package models

type JudgeRequest struct {
	Code     string   `json:"code"`
	Language string   `json:"language"`
	Stdin    []string `json:"stdin"`
}

type JudgeResponse struct {
	Verdict []string `json:"verdict"`
	Stdout  []string `json:"stdout"`
	Stderr  []string `json:"stderr"`
	Time    []int32  `json:"time"`
	Memory  []int32  `json:"memory"`
}
