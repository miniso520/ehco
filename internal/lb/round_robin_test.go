package lb

import (
	"testing"
)

func Test_roundrobin_Next(t *testing.T) {
	remotes := []string{
		"127.0.0.1",
		"127.0.0.2",
		"127.0.0.3",
	}
	nodeList := make([]*Node, len(remotes))
	for i := range remotes {
		nodeList[i] = &Node{Address: remotes[i]}
	}
	rb := NewRoundRobin(nodeList)

	for i := 0; i < len(remotes); i++ {
		if node := rb.Next(); node.Address != remotes[i] {
			t.Fatalf("need %s got %s", remotes[i], node.Address)
		}
	}
}
