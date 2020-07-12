package collector

import (
	"context"
	"fildr-cli/internal/config"
	"fildr-cli/internal/module"
	"fmt"
	"time"
)

var _ module.Module = (*Collector)(nil)

type Collector struct {
	config *config.TomlConfig
}

func New(ctx context.Context, config *config.TomlConfig) (*Collector, error) {
	return &Collector{config: config}, nil
}

func (c *Collector) Name() string {
	return "collector"
}

func (c *Collector) Start() error {
	evaluation := c.config.Gateway.Evaluation
	if evaluation == 0 {
		evaluation = 5
	}

	instance := c.config.Gateway.Instance

	c.execute("node", "node", instance, time.Duration(evaluation))

	return nil
}

func (c *Collector) Stop() {

}

func (c *Collector) execute(namespace string, job string, instanceName string, evaluation time.Duration) {
	//go func() {
	instance := GetInstance(namespace)
	//instance.SetJob(job)
	//instance.SetInstance(instanceName)
	for range time.Tick(time.Second * evaluation) {
		fmt.Println("print metrics ...")
		fmt.Println(instance.GetMetrics())
	}
	//}()
}
