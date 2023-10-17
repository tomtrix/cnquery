// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package confluence

import (
	"errors"
	"os"

	"github.com/ctreminiom/go-atlassian/confluence"
	"go.mondoo.com/cnquery/v9/providers-sdk/v1/inventory"
	"go.mondoo.com/cnquery/v9/providers/atlassian/connection/shared"
)

const (
	Confluence shared.ConnectionType = "confluece"
)

type ConfluenceConnection struct {
	id     uint32
	Conf   *inventory.Config
	asset  *inventory.Asset
	client *confluence.Client
	name   string
}

func NewConnection(id uint32, asset *inventory.Asset, conf *inventory.Config) (*ConfluenceConnection, error) {
	host := conf.Options["host"]
	if host == "" {
		host = os.Getenv("ATLASSIAN_HOST")
	}
	if host == "" {
		return nil, errors.New("you need to provide atlassian hostname via ATLASSIAN_HOST env or via --host flag")
	}

	user := conf.Options["user"]
	if user == "" {
		user = os.Getenv("ATLASSIAN_USER")
	}
	if user == "" {
		return nil, errors.New("you need to provide atlassian username via ATLASSIAN_USER env or via --user flag")
	}

	token := conf.Options["user-token"]
	if token == "" {
		token = os.Getenv("ATLASSIAN_USER_TOKEN")
	}
	if token == "" {
		return nil, errors.New("you need to provide atlassian user token via ATLASSIAN_USER_TOKEN env or via --user-token flag")
	}

	client, err := confluence.New(nil, host)
	if err != nil {
		return nil, err
	}

	client.Auth.SetBasicAuth(user, token)
	client.Auth.SetUserAgent("curl/7.54.0")

	conn := &ConfluenceConnection{
		Conf:   conf,
		id:     id,
		asset:  asset,
		client: client,
		name:   host,
	}

	return conn, nil
}

func (c *ConfluenceConnection) Name() string {
	return c.name
}

func (c *ConfluenceConnection) ID() uint32 {
	return c.id
}

func (c *ConfluenceConnection) Asset() *inventory.Asset {
	return c.asset
}

func (c *ConfluenceConnection) Client() *confluence.Client {
	return c.client
}

func (c *ConfluenceConnection) Type() shared.ConnectionType {
	return Confluence
}

func (c *ConfluenceConnection) ConnectionType() string {
	return "confluence"
}

func (c *ConfluenceConnection) Config() *inventory.Config {
	return c.Conf
}