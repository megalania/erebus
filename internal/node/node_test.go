package node_test

import (
	"testing"

	"github.com/megalania/erebus/internal/node"
)

func TestNewNode(t *testing.T) {
	node, err := node.NewNode()
	if err != nil {
		t.Fatal(
			err.Error(),
		)
	}
	t.Logf(
		"ID: %s, PrivateIPv4Address: %s PublicIPv4Address: %s",
		node.ID,
		node.PrivateIPv4Address,
		node.PublicIPv4Address,
	)
}
