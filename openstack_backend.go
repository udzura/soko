package soko

import (
	"fmt"
	"io/ioutil"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

type OpenStackBackend struct {
	SectionConfig

	client *gophercloud.ServiceClient
}

const openstackTomlTemplate = `[default]
backend = "openstack"

[openstack]
username = "%s"
password = "%s"
tenant_name = "%s"
auth_url = "%s"
region = "%s"
`

func NewOpenStackBackend(config SectionConfig) (*OpenStackBackend, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: config["auth_url"],
		Username:         config["username"],
		Password:         config["password"],
		TenantName:       config["tenant_name"],
	}
	auth, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	regionOpts := gophercloud.EndpointOpts{Region: config["region"]}
	cli, err := openstack.NewComputeV2(auth, regionOpts)
	if err != nil {
		return nil, err
	}

	return &OpenStackBackend{
		SectionConfig: config,
		client:        cli,
	}, nil
}

func (b *OpenStackBackend) Save() error {
	config := b.SectionConfig
	data := fmt.Sprintf(
		openstackTomlTemplate,
		config["username"],
		config["password"],
		config["tenant_name"],
		config["auth_url"],
		config["region"],
	)
	return ioutil.WriteFile(defaultConfigPath, []byte(data), 0644)
}

func (b *OpenStackBackend) Get(serverID string, key string) (string, error) {
	metadata, err := b.prefetch(serverID)
	if err != nil {
		return "", err
	}

	if v, ok := metadata[key]; ok {
		if v == "" {
			sayEmpty(v)
		}

		return v, nil
	} else {
		sayEmpty(key)
		return "", nil
	}
}

func (b *OpenStackBackend) Put(serverID string, key string, value string) error {
	metadata, err := b.prefetch(serverID)
	if err != nil {
		return err
	}

	metadata[key] = value

	opts := servers.MetadataOpts(metadata)

	_, err = servers.UpdateMetadata(b.client, serverID, opts).Extract()
	return err
}

func (b *OpenStackBackend) Delete(serverID string, key string) error {
	metadata, err := b.prefetch(serverID)
	if err != nil {
		return err
	}
	delete(metadata, key)

	opts := servers.MetadataOpts(metadata)

	_, err = servers.ResetMetadata(b.client, serverID, opts).Extract()
	return err
}

func (b *OpenStackBackend) prefetch(serverID string) (map[string]string, error) {
	s, err := servers.Get(b.client, serverID).Extract()
	if err != nil {
		return nil, err
	}

	return toStringMap(s.Metadata), nil
}
