// Code generated by interfacer; DO NOT EDIT

package interfaces

import (
	"github.com/google/nftables"
)

// NetlinkConnector is an interface generated for "github.com/black-desk/cgtproxy/pkg/nftman/connector.Connector".
type NetlinkConnector interface {
	Connect() (*nftables.Conn, error)
	Release() error
}