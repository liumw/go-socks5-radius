package socks5

type RadiusServer interface {
	getServer() string
	getSecret() string
}

type RadiusCfg struct {
	Server string
	Secret string
}

func (s RadiusCfg) getServer() string {
	return s.Server
}

func (s RadiusCfg) getSecret() string {
	return s.Secret
}