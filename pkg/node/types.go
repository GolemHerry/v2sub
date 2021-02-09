package node

import "time"

var nodes = make([]Node, 0)

type Node struct {
	Name       string `json:"ps"`
	Host       string `json:"host"`
	Path       string `json:"path"`
	TLS        string `json:"tls"`
	VerifyCert bool   `json:"verify_cert"`
	Add        string `json:"add"`
	Aid        int    `json:"aid"`
	Net        string `json:"net"`
	HeaderType string `json:"headerType"`
	Vip        string `json:"v"`
	Port       int    `json:"port"`
	Remark     string `json:"remark"`
	Id         string `json:"id"`
	Class      int    `json:"class"`
	IP         string
	Time       time.Duration
}
