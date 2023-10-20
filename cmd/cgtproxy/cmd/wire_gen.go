// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/black-desk/cgtproxy/pkg/cgtproxy/config"
	"github.com/black-desk/cgtproxy/pkg/interfaces"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func injectedCGTProxy(configConfig *config.Config, sugaredLogger *zap.SugaredLogger) (interfaces.CGTProxy, error) {
	cgroupRoot := provideCgroupRoot(configConfig)
	cGroupMonitor, err := provideCgrougMontior(cgroupRoot, sugaredLogger)
	if err != nil {
		return nil, err
	}
	bypass := provideBypass(configConfig)
	nftManager, err := provideNFTManager(cgroupRoot, bypass, sugaredLogger)
	if err != nil {
		return nil, err
	}
	cmdChans := provideChans()
	v := provideInputChan(cmdChans)
	routeManager, err := provideRuleManager(nftManager, configConfig, v, sugaredLogger)
	if err != nil {
		return nil, err
	}
	cgtProxy, err := provideCGTProxy(cGroupMonitor, routeManager, sugaredLogger, configConfig)
	if err != nil {
		return nil, err
	}
	return cgtProxy, nil
}