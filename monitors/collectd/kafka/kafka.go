package kafka

import (
	"github.com/signalfx/neo-agent/monitors"
	"github.com/signalfx/neo-agent/monitors/collectd/genericjmx"
)

const monitorType = "collectd/kafka"

var serviceName = "kafka"

// Monitor is the main type that represents the monitor
type Monitor struct {
	*genericjmx.MonitorCore
}

func init() {
	monitors.Register(monitorType, func() interface{} {
		return &Monitor{
			genericjmx.Instance(),
		}
	}, &genericjmx.Config{})
}

// Configure configures and runs the plugin in collectd
func (km *Monitor) Configure(conf *genericjmx.Config) bool {
	conf.Common.ServiceName = &serviceName

	conf.Common.MBeanDefinitions = conf.Common.MBeanDefinitions.MergeWith(defaultMBeans)
	km.AddConfiguration(conf)
	return true
}
