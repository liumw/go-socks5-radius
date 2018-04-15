go-socks5-radius
================

Provides the `socks5` package that implements a [SOCKS5 server](http://en.wikipedia.org/wiki/SOCKS).
SOCKS (Secure Sockets) is used to route traffic between a client and server through
an intermediate proxy layer. This can be used to bypass firewalls or NATs.

Attention
=========

This is a fork of [go-socks5](https://github.com/armon/go-socks5) that uses RADIUS server instead of in-app algorithm.
Use it at your own risk. I haven't tested it in various environments but it works for me. Low code quality included.

Example
=======

Below is a simple example of usage

```go
	rad := socks5.RadiusCfg{
		Server: "radius.example.com:1812",
		Secret: "radiusSecret",
	}
	cator := socks5.RadiusAuthenticator{Radius: rad}
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
```

