// Code generated by "gogen option"; DO NOT EDIT.
// Exec: "gogen option -n ClientOption -f Client -o option.client.go"
// Version: 0.0.2

package gnet

import (
	process "github.com/aggronmagi/walle/net/process"
	zaplog "github.com/aggronmagi/walle/zaplog"
	goframe "github.com/smallnest/goframe"
)

var _ = walleClient()

// ClientOption
type ClientOptions struct {
	Network string
	// Addr Server Addr
	Addr string
	// Process Options
	ProcessOptions []process.ProcessOption
	// process router
	Router Router
	// log interface
	Logger (*zaplog.Logger)
	// AutoReconnect auto reconnect server. zero means not reconnect!
	AutoReconnectTime int
	// StopImmediately when session finish,business finish immediately.
	StopImmediately bool
	EncodeConfig    (*goframe.EncoderConfig)
	DecodeConfig    (*goframe.DecoderConfig)
}

func WithClientOptionsNetwork(v string) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.Network
		cc.Network = v
		return WithClientOptionsNetwork(previous)
	}
}

// Addr Server Addr
func WithClientOptionsAddr(v string) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.Addr
		cc.Addr = v
		return WithClientOptionsAddr(previous)
	}
}

// Process Options
func WithClientOptionsProcessOptions(v ...process.ProcessOption) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.ProcessOptions
		cc.ProcessOptions = v
		return WithClientOptionsProcessOptions(previous...)
	}
}

// process router
func WithClientOptionsRouter(v Router) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.Router
		cc.Router = v
		return WithClientOptionsRouter(previous)
	}
}

// log interface
func WithClientOptionsLogger(v *zaplog.Logger) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.Logger
		cc.Logger = v
		return WithClientOptionsLogger(previous)
	}
}

// AutoReconnect auto reconnect server. zero means not reconnect!
func WithClientOptionsAutoReconnectTime(v int) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.AutoReconnectTime
		cc.AutoReconnectTime = v
		return WithClientOptionsAutoReconnectTime(previous)
	}
}

// StopImmediately when session finish,business finish immediately.
func WithClientOptionsStopImmediately(v bool) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.StopImmediately
		cc.StopImmediately = v
		return WithClientOptionsStopImmediately(previous)
	}
}
func WithClientOptionsEncodeConfig(v *goframe.EncoderConfig) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.EncodeConfig
		cc.EncodeConfig = v
		return WithClientOptionsEncodeConfig(previous)
	}
}
func WithClientOptionsDecodeConfig(v *goframe.DecoderConfig) ClientOption {
	return func(cc *ClientOptions) ClientOption {
		previous := cc.DecodeConfig
		cc.DecodeConfig = v
		return WithClientOptionsDecodeConfig(previous)
	}
}

// SetOption modify options
func (cc *ClientOptions) SetOption(opt ClientOption) {
	_ = opt(cc)
}

// ApplyOption modify options
func (cc *ClientOptions) ApplyOption(opts ...ClientOption) {
	for _, opt := range opts {
		_ = opt(cc)
	}
}

// GetSetOption modify and get last option
func (cc *ClientOptions) GetSetOption(opt ClientOption) ClientOption {
	return opt(cc)
}

// ClientOption option define
type ClientOption func(cc *ClientOptions) ClientOption

// NewClientOptions create options instance.
func NewClientOptions(opts ...ClientOption) *ClientOptions {
	cc := newDefaultClientOptions()
	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogClientOptions != nil {
		watchDogClientOptions(cc)
	}
	return cc
}

// InstallClientOptionsWatchDog install watch dog
func InstallClientOptionsWatchDog(dog func(cc *ClientOptions)) {
	watchDogClientOptions = dog
}

var watchDogClientOptions func(cc *ClientOptions)

// newDefaultClientOptions new option with default value
func newDefaultClientOptions() *ClientOptions {
	cc := &ClientOptions{
		Network:           "tcp",
		Addr:              "localhost:8080",
		ProcessOptions:    nil,
		Router:            nil,
		Logger:            zaplog.Default,
		AutoReconnectTime: 5,
		StopImmediately:   false,
		EncodeConfig:      DefaultClientEncodeConfig,
		DecodeConfig:      DefaultClientDecodeConfig,
	}
	return cc
}