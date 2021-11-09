// Code generated by "gogen option"; DO NOT EDIT.
// Exec: "gogen option -n ProcessOption -o option.process.go"
// Version: 0.0.2

package process

import (
	packet "github.com/aggronmagi/walle/net/packet"
	zaplog "github.com/aggronmagi/walle/zaplog"
)

var _ = walleProcessOption()

// ProcessOption process option
type ProcessOptions struct {
	// log interface
	Logger (*zaplog.Logger)
	// packet pool
	PacketPool packet.PacketPool
	// packet encoder
	PacketEncode PacketEncoder
	// packet codec
	PacketCodec PacketCodec
	// message codec
	MsgCodec MessageCodec
	// dispatch packet data filter
	DispatchDataFilter PacketDispatcherFilter
	// load limit. return true to ignore packet.
	LoadLimitFilter func(ctx Context, count int64, req *packet.Packet) bool
}

// log interface
func WithLogger(v *zaplog.Logger) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.Logger
		cc.Logger = v
		return WithLogger(previous)
	}
}

// packet pool
func WithPacketPool(v packet.PacketPool) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.PacketPool
		cc.PacketPool = v
		return WithPacketPool(previous)
	}
}

// packet encoder
func WithPacketEncode(v PacketEncoder) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.PacketEncode
		cc.PacketEncode = v
		return WithPacketEncode(previous)
	}
}

// packet codec
func WithPacketCodec(v PacketCodec) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.PacketCodec
		cc.PacketCodec = v
		return WithPacketCodec(previous)
	}
}

// message codec
func WithMsgCodec(v MessageCodec) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.MsgCodec
		cc.MsgCodec = v
		return WithMsgCodec(previous)
	}
}

// dispatch packet data filter
func WithDispatchDataFilter(v PacketDispatcherFilter) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.DispatchDataFilter
		cc.DispatchDataFilter = v
		return WithDispatchDataFilter(previous)
	}
}

// load limit. return true to ignore packet.
func WithLoadLimitFilter(v func(ctx Context, count int64, req *packet.Packet) bool) ProcessOption {
	return func(cc *ProcessOptions) ProcessOption {
		previous := cc.LoadLimitFilter
		cc.LoadLimitFilter = v
		return WithLoadLimitFilter(previous)
	}
}

// SetOption modify options
func (cc *ProcessOptions) SetOption(opt ProcessOption) {
	_ = opt(cc)
}

// ApplyOption modify options
func (cc *ProcessOptions) ApplyOption(opts ...ProcessOption) {
	for _, opt := range opts {
		_ = opt(cc)
	}
}

// GetSetOption modify and get last option
func (cc *ProcessOptions) GetSetOption(opt ProcessOption) ProcessOption {
	return opt(cc)
}

// ProcessOption option define
type ProcessOption func(cc *ProcessOptions) ProcessOption

// NewProcessOptions create options instance.
func NewProcessOptions(opts ...ProcessOption) *ProcessOptions {
	cc := newDefaultProcessOptions()
	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogProcessOptions != nil {
		watchDogProcessOptions(cc)
	}
	return cc
}

// InstallProcessOptionsWatchDog install watch dog
func InstallProcessOptionsWatchDog(dog func(cc *ProcessOptions)) {
	watchDogProcessOptions = dog
}

var watchDogProcessOptions func(cc *ProcessOptions)

// newDefaultProcessOptions new option with default value
func newDefaultProcessOptions() *ProcessOptions {
	cc := &ProcessOptions{
		Logger:             zaplog.Default,
		PacketPool:         packet.DefaultPacketPool,
		PacketEncode:       &EmtpyPacketCoder{},
		PacketCodec:        PacketCodecProtobuf,
		MsgCodec:           MessageCodecProtobuf,
		DispatchDataFilter: DefaultPacketFilter,
		LoadLimitFilter: func(ctx Context, count int64, req *packet.Packet) bool {
			return false
		},
	}
	return cc
}