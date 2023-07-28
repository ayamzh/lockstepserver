// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package xconfig

import (
	"sync/atomic"
	"unsafe"

	"github.com/byebyebruce/lockstepserver/frame/network"
)

// Config should use NewConfig to initialize it
type Config struct {
	OptionUsage string          `xconf:"option_usage"`
	HttpPort    int             `xconf:"http_port" usage:"HTTP端口"`
	DebugLog    bool            `xconf:"debug_log" usage:"打开DEBUGLOG"`
	Server      *network.Config `xconf:"server"`
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

// WithOptionUsage option func for OptionUsage
func WithOptionUsage(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.OptionUsage
		cc.OptionUsage = v
		return WithOptionUsage(previous)
	}
}

// HTTP端口
// WithHttpPort option func for HttpPort
func WithHttpPort(v int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.HttpPort
		cc.HttpPort = v
		return WithHttpPort(previous)
	}
}

// 打开DEBUGLOG
// WithDebugLog option func for DebugLog
func WithDebugLog(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.DebugLog
		cc.DebugLog = v
		return WithDebugLog(previous)
	}
}

// WithServer option func for Server
func WithServer(v *network.Config) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Server
		cc.Server = v
		return WithServer(previous)
	}
}

// NewConfig new Config
func NewConfig(opts ...ConfigOption) *Config {
	cc := newDefaultConfig()
	for _, opt := range opts {
		opt(cc)
	}
	if watchDogConfig != nil {
		watchDogConfig(cc)
	}
	return cc
}

// InstallConfigWatchDog the installed func will called when NewConfig  called
func InstallConfigWatchDog(dog func(cc *Config)) { watchDogConfig = dog }

// watchDogConfig global watch dog
var watchDogConfig func(cc *Config)

// newDefaultConfig new default Config
func newDefaultConfig() *Config {
	cc := &Config{}

	for _, opt := range [...]ConfigOption{
		WithOptionUsage(optionUsage),
		WithHttpPort(80),
		WithDebugLog(true),
		WithServer(network.NewServerConf()),
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
func (cc *Config) GetOptionUsage() string     { return cc.OptionUsage }
func (cc *Config) GetHttpPort() int           { return cc.HttpPort }
func (cc *Config) GetDebugLog() bool          { return cc.DebugLog }
func (cc *Config) GetServer() *network.Config { return cc.Server }

// ConfigVisitor visitor interface for Config
type ConfigVisitor interface {
	GetOptionUsage() string
	GetHttpPort() int
	GetDebugLog() bool
	GetServer() *network.Config
}

// ConfigInterface visitor + ApplyOption interface for Config
type ConfigInterface interface {
	ConfigVisitor
	ApplyOption(...ConfigOption) []ConfigOption
}
