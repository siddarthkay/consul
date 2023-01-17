package service

import (
	"context"
	"fmt"
	"io"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"

	libcluster "github.com/hashicorp/consul/test/integration/consul-container/libs/cluster"
	"github.com/hashicorp/consul/test/integration/consul-container/libs/utils"
)

func CreateAndRegisterStaticServerAndSidecar(node libcluster.Agent) (Service, Service, error) {
	// Do some trickery to ensure that partial completion is correctly torn
	// down, but successful execution is not.
	var deferClean utils.ResettableDefer
	defer deferClean.Execute()

	// Create a service and proxy instance
	serverService, err := NewExampleService(context.Background(), "static-server", 8080, 8079, node)
	if err != nil {
		return nil, nil, err
	}
	deferClean.Add(func() {
		_ = serverService.Terminate()
	})

	serverConnectProxy, err := NewConnectService(context.Background(), "static-server-sidecar", "static-server", 8080, node) // bindPort not used
	if err != nil {
		return nil, nil, err
	}
	deferClean.Add(func() {
		_ = serverConnectProxy.Terminate()
	})

	// Register the static-server service and sidecar
	req := &api.AgentServiceRegistration{
		Name: "static-server",
		Port: 8080,
		Connect: &api.AgentServiceConnect{
			SidecarService: &api.AgentServiceRegistration{
				Proxy: &api.AgentServiceConnectProxyConfig{},
			},
		},
		Check: &api.AgentServiceCheck{
			Name:     "Static Server Listening",
			TCP:      fmt.Sprintf("127.0.0.1:%d", 8080),
			Interval: "10s",
			Status:   api.HealthPassing,
		},
	}

	err = node.GetClient().Agent().ServiceRegister(req)
	if err != nil {
		return serverService, serverConnectProxy, err
	}

	// disable cleanup functions now that we have an object with a Terminate() function
	deferClean.Reset()

	return serverService, serverConnectProxy, nil
}

func CreateAndRegisterStaticClientSidecar(
	node libcluster.Agent,
	peerName string,
	localMeshGateway bool,
) (*ConnectContainer, error) {
	// Do some trickery to ensure that partial completion is correctly torn
	// down, but successful execution is not.
	var deferClean utils.ResettableDefer
	defer deferClean.Execute()

	// Create a service and proxy instance
	clientConnectProxy, err := NewConnectService(context.Background(), "static-client-sidecar", "static-client", 5000, node)
	if err != nil {
		return nil, err
	}
	deferClean.Add(func() {
		_ = clientConnectProxy.Terminate()
	})

	mgwMode := api.MeshGatewayModeRemote
	if localMeshGateway {
		mgwMode = api.MeshGatewayModeLocal
	}

	// Register the static-client service and sidecar
	req := &api.AgentServiceRegistration{
		Name: "static-client",
		Port: 8080,
		Connect: &api.AgentServiceConnect{
			SidecarService: &api.AgentServiceRegistration{
				Proxy: &api.AgentServiceConnectProxyConfig{
					Upstreams: []api.Upstream{{
						DestinationName:  "static-server",
						DestinationPeer:  peerName,
						LocalBindAddress: "0.0.0.0",
						LocalBindPort:    5000,
						MeshGateway: api.MeshGatewayConfig{
							Mode: mgwMode,
						},
					}},
				},
			},
		},
	}

	err = node.GetClient().Agent().ServiceRegister(req)
	if err != nil {
		return clientConnectProxy, err
	}

	// disable cleanup functions now that we have an object with a Terminate() function
	deferClean.Reset()

	return clientConnectProxy, nil
}

func GetEnvoyConfigDump(port int) (string, error) {
	client := cleanhttp.DefaultClient()
	url := fmt.Sprintf("http://localhost:%d/config_dump?include_eds", port)

	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}