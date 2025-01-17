// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package network

import (
	"sync/atomic"
	"time"
	"unsafe"
)

// Config should use NewServerConf to initialize it
type Config struct {
	UdpPort                int           `xconf:"udp_port" usage:"udp端口"`
	PacketSendChanLimit    uint32        `xconf:"packet_send_chan_limit" usage:"the limit of packet send channel"`
	PacketReceiveChanLimit uint32        `xconf:"packet_receive_chan_limit" usage:"the limit of packet receive channel"`
	ConnReadTimeout        time.Duration `xconf:"conn_read_timeout" usage:"read timeout"`
	ConnWriteTimeout       time.Duration `xconf:"conn_write_timeout" usage:"write timeout"`
}

// ApplyOption apply mutiple new option and return the old ones
// sample:
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
func (cc *Config) ApplyOption(opts ...ConfigOption) []ConfigOption {
	var previous []ConfigOption
	for _, opt := range opts {
		previous = append(previous, opt(cc))
	}
	return previous
}

// ConfigOption option func
type ConfigOption func(cc *Config) ConfigOption

// udp端口
// WithUdpPort option func for UdpPort
func WithUdpPort(v int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.UdpPort
		cc.UdpPort = v
		return WithUdpPort(previous)
	}
}

// the limit of packet send channel
// WithPacketSendChanLimit option func for PacketSendChanLimit
func WithPacketSendChanLimit(v uint32) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.PacketSendChanLimit
		cc.PacketSendChanLimit = v
		return WithPacketSendChanLimit(previous)
	}
}

// the limit of packet receive channel
// WithPacketReceiveChanLimit option func for PacketReceiveChanLimit
func WithPacketReceiveChanLimit(v uint32) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.PacketReceiveChanLimit
		cc.PacketReceiveChanLimit = v
		return WithPacketReceiveChanLimit(previous)
	}
}

// read timeout
// WithConnReadTimeout option func for ConnReadTimeout
func WithConnReadTimeout(v time.Duration) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.ConnReadTimeout
		cc.ConnReadTimeout = v
		return WithConnReadTimeout(previous)
	}
}

// write timeout
// WithConnWriteTimeout option func for ConnWriteTimeout
func WithConnWriteTimeout(v time.Duration) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.ConnWriteTimeout
		cc.ConnWriteTimeout = v
		return WithConnWriteTimeout(previous)
	}
}

// NewServerConf new Config
func NewServerConf(opts ...ConfigOption) *Config {
	cc := newDefaultConfig()
	for _, opt := range opts {
		opt(cc)
	}
	if watchDogConfig != nil {
		watchDogConfig(cc)
	}
	return cc
}

// InstallConfigWatchDog the installed func will called when NewServerConf  called
func InstallConfigWatchDog(dog func(cc *Config)) { watchDogConfig = dog }

// watchDogConfig global watch dog
var watchDogConfig func(cc *Config)

// newDefaultConfig new default Config
func newDefaultConfig() *Config {
	cc := &Config{}

	for _, opt := range [...]ConfigOption{
		WithUdpPort(9000),
		WithPacketSendChanLimit(1024),
		WithPacketReceiveChanLimit(1024),
		WithConnReadTimeout(time.Second * 5),
		WithConnWriteTimeout(time.Second * 5),
	} {
		opt(cc)
	}

	return cc
}

// AtomicSetFunc used for XConf
func (cc *Config) AtomicSetFunc() func(interface{}) { return AtomicConfigSet }

// atomicConfig global *Config holder
var atomicConfig unsafe.Pointer

// onAtomicConfigSet global call back when  AtomicConfigSet called by XConf.
// use ConfigInterface.ApplyOption to modify the updated cc
// if passed in cc not valid, then return false, cc will not set to atomicConfig
var onAtomicConfigSet func(cc ConfigInterface) bool

// InstallCallbackOnAtomicConfigSet install callback
func InstallCallbackOnAtomicConfigSet(callback func(cc ConfigInterface) bool) {
	onAtomicConfigSet = callback
}

// AtomicConfigSet atomic setter for *Config
func AtomicConfigSet(update interface{}) {
	cc := update.(*Config)
	if onAtomicConfigSet != nil && !onAtomicConfigSet(cc) {
		return
	}
	atomic.StorePointer(&atomicConfig, (unsafe.Pointer)(cc))
}

// AtomicConfig return atomic *ConfigVisitor
func AtomicConfig() ConfigVisitor {
	current := (*Config)(atomic.LoadPointer(&atomicConfig))
	if current == nil {
		defaultOne := newDefaultConfig()
		if watchDogConfig != nil {
			watchDogConfig(defaultOne)
		}
		atomic.CompareAndSwapPointer(&atomicConfig, nil, (unsafe.Pointer)(defaultOne))
		return (*Config)(atomic.LoadPointer(&atomicConfig))
	}
	return current
}

// all getter func
func (cc *Config) GetUdpPort() int                    { return cc.UdpPort }
func (cc *Config) GetPacketSendChanLimit() uint32     { return cc.PacketSendChanLimit }
func (cc *Config) GetPacketReceiveChanLimit() uint32  { return cc.PacketReceiveChanLimit }
func (cc *Config) GetConnReadTimeout() time.Duration  { return cc.ConnReadTimeout }
func (cc *Config) GetConnWriteTimeout() time.Duration { return cc.ConnWriteTimeout }

// ConfigVisitor visitor interface for Config
type ConfigVisitor interface {
	GetUdpPort() int
	GetPacketSendChanLimit() uint32
	GetPacketReceiveChanLimit() uint32
	GetConnReadTimeout() time.Duration
	GetConnWriteTimeout() time.Duration
}

// ConfigInterface visitor + ApplyOption interface for Config
type ConfigInterface interface {
	ConfigVisitor
	ApplyOption(...ConfigOption) []ConfigOption
}
