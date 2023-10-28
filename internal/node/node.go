package node

import (
	"net"

	"github.com/google/uuid"

	"github.com/megalania/erebus/internal/network"
)

type Node struct {
	ID uuid.UUID
	PrivateIPv4Address net.IP
	PublicIPv4Address  net.IP
}

func NewNode() (*Node, error) {
	privateIPv4Address, err := network.GetPrivateIPv4Address()
	if err != nil {
		return nil, err
	}
	publicIPv4Address, err := network.GetPublicIPv4Address()
	if err != nil {
		return nil, err
	}
	return &Node{
		ID: uuid.New(),
		PrivateIPv4Address: privateIPv4Address,
		PublicIPv4Address: publicIPv4Address,
	}, nil
}
