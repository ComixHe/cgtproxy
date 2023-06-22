// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package core

import (
	"github.com/black-desk/cgtproxy/pkg/cgtproxy/config"
)

// Injectors from wire.go:

func injectedComponents(configConfig *config.Config) (*components, error) {
	cgroupRoot := provideCgroupRoot(configConfig)
	watcher, err := provideWatcher(cgroupRoot)
	if err != nil {
		return nil, err
	}
	coreChans := provideChans()
	v := provideOutputChan(coreChans)
	monitor, err := provideMonitor(v, watcher, cgroupRoot)
	if err != nil {
		return nil, err
	}
	bypass := provideBypass(configConfig)
	table, err := provideTable(cgroupRoot, bypass)
	if err != nil {
		return nil, err
	}
	v2 := provideInputChan(coreChans)
	routeManager, err := provideRuleManager(table, configConfig, v2)
	if err != nil {
		return nil, err
	}
	coreComponents := provideComponents(watcher, monitor, routeManager)
	return coreComponents, nil
}
