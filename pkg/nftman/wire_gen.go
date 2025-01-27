// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package nftman

import (
	"github.com/black-desk/cgtproxy/internal/tests/logger"
	"github.com/black-desk/cgtproxy/pkg/cgtproxy/config"
	"github.com/black-desk/cgtproxy/pkg/interfaces"
	"github.com/black-desk/cgtproxy/pkg/nftman/connector"
	"github.com/black-desk/cgtproxy/pkg/nftman/lastingconnector"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func injectedNFTManagerWithLastingConnector(cGroupRoot config.CGroupRoot) (*NFTManager, error) {
	sugaredLogger, err := logger.ProvideLogger()
	if err != nil {
		return nil, err
	}
	netlinkConnector, err := provideLastingConnector()
	if err != nil {
		return nil, err
	}
	nftManager, err := provideNFTManager(cGroupRoot, sugaredLogger, netlinkConnector)
	if err != nil {
		return nil, err
	}
	return nftManager, nil
}

func injectedNFTManagerWithConnector(cGroupRoot config.CGroupRoot) (*NFTManager, error) {
	sugaredLogger, err := logger.ProvideLogger()
	if err != nil {
		return nil, err
	}
	netlinkConnector, err := provideLastingConnector()
	if err != nil {
		return nil, err
	}
	nftManager, err := provideNFTManager(cGroupRoot, sugaredLogger, netlinkConnector)
	if err != nil {
		return nil, err
	}
	return nftManager, nil
}

// wire.go:

func provideLastingConnector() (ret interfaces.NetlinkConnector, err error) {
	return lastingconnector.New()
}

func provideConnector() (ret interfaces.NetlinkConnector, err error) {
	return connector.New()
}

func provideNFTManager(
	root config.CGroupRoot, logger2 *zap.SugaredLogger, connector2 interfaces.NetlinkConnector,
) (
	ret *NFTManager, err error,
) {

	return New(WithCgroupRoot(root), WithLogger(logger2), WithConnFactory(connector2))

}

var testWithLastingConnectorSet = wire.NewSet(
	provideLastingConnector,
	provideNFTManager, logger.ProvideLogger,
)

var testWithConnectorSet = wire.NewSet(
	provideConnector,
	provideNFTManager, logger.ProvideLogger,
)
