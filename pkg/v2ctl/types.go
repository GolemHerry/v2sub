package v2ctl

type V2Ray interface {
	Start() error
	Stop() error
	Restart() error
}

type v2Ray struct {
	status string
	flags  map[string]string
	sonfig interface{}
}
