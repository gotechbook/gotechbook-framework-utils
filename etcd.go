package utils

import (
	v3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/tests/v3/framework/integration"
	"testing"
)

func GetTestEtcd(t *testing.T) (*integration.Cluster, *v3.Client) {
	t.Helper()
	integration.BeforeTest(t)
	c := integration.NewCluster(t, &integration.ClusterConfig{Size: 1})
	cli := c.RandClient()
	return c, cli
}
