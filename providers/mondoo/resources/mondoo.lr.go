// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by resources. DO NOT EDIT.

package resources

import (
	"errors"
	"time"

	"go.mondoo.com/cnquery/v11/llx"
	"go.mondoo.com/cnquery/v11/providers-sdk/v1/plugin"
	"go.mondoo.com/cnquery/v11/types"
)

var resourceFactories map[string]plugin.ResourceFactory

func init() {
	resourceFactories = map[string]plugin.ResourceFactory {
		"mondoo.client": {
			// to override args, implement: initMondooClient(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createMondooClient,
		},
		"mondoo.organization": {
			Init: initMondooOrganization,
			Create: createMondooOrganization,
		},
		"mondoo.space": {
			Init: initMondooSpace,
			Create: createMondooSpace,
		},
		"mondoo.asset": {
			// to override args, implement: initMondooAsset(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createMondooAsset,
		},
		"mondoo.resource": {
			// to override args, implement: initMondooResource(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createMondooResource,
		},
	}
}

// NewResource is used by the runtime of this plugin to create new resources.
// Its arguments may be provided by users. This function is generally not
// used by initializing resources from recordings or from lists.
func NewResource(runtime *plugin.Runtime, name string, args map[string]*llx.RawData) (plugin.Resource, error) {
	f, ok := resourceFactories[name]
	if !ok {
		return nil, errors.New("cannot find resource " + name + " in this provider")
	}

	if f.Init != nil {
		cargs, res, err := f.Init(runtime, args)
		if err != nil {
			return res, err
		}

		if res != nil {
			id := name+"\x00"+res.MqlID()
			if x, ok := runtime.Resources.Get(id); ok {
				return x, nil
			}
			runtime.Resources.Set(id, res)
			return res, nil
		}

		args = cargs
	}

	res, err := f.Create(runtime, args)
	if err != nil {
		return nil, err
	}

	id := name+"\x00"+res.MqlID()
	if x, ok := runtime.Resources.Get(id); ok {
		return x, nil
	}

	runtime.Resources.Set(id, res)
	return res, nil
}

// CreateResource is used by the runtime of this plugin to create resources.
// Its arguments must be complete and pre-processed. This method is used
// for initializing resources from recordings or from lists.
func CreateResource(runtime *plugin.Runtime, name string, args map[string]*llx.RawData) (plugin.Resource, error) {
	f, ok := resourceFactories[name]
	if !ok {
		return nil, errors.New("cannot find resource " + name + " in this provider")
	}

	res, err := f.Create(runtime, args)
	if err != nil {
		return nil, err
	}

	id := name+"\x00"+res.MqlID()
	if x, ok := runtime.Resources.Get(id); ok {
		return x, nil
	}

	runtime.Resources.Set(id, res)
	return res, nil
}

var getDataFields = map[string]func(r plugin.Resource) *plugin.DataRes{
	"mondoo.client.mrn": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooClient).GetMrn()).ToDataRes(types.String)
	},
	"mondoo.organization.name": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooOrganization).GetName()).ToDataRes(types.String)
	},
	"mondoo.organization.mrn": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooOrganization).GetMrn()).ToDataRes(types.String)
	},
	"mondoo.organization.spaces": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooOrganization).GetSpaces()).ToDataRes(types.Array(types.Resource("mondoo.space")))
	},
	"mondoo.space.name": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooSpace).GetName()).ToDataRes(types.String)
	},
	"mondoo.space.mrn": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooSpace).GetMrn()).ToDataRes(types.String)
	},
	"mondoo.space.assets": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooSpace).GetAssets()).ToDataRes(types.Array(types.Resource("mondoo.asset")))
	},
	"mondoo.asset.name": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetName()).ToDataRes(types.String)
	},
	"mondoo.asset.mrn": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetMrn()).ToDataRes(types.String)
	},
	"mondoo.asset.platform": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetPlatform()).ToDataRes(types.String)
	},
	"mondoo.asset.annotations": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetAnnotations()).ToDataRes(types.Map(types.String, types.String))
	},
	"mondoo.asset.labels": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetLabels()).ToDataRes(types.Map(types.String, types.String))
	},
	"mondoo.asset.updatedAt": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetUpdatedAt()).ToDataRes(types.Time)
	},
	"mondoo.asset.scoreValue": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetScoreValue()).ToDataRes(types.Int)
	},
	"mondoo.asset.scoreGrade": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetScoreGrade()).ToDataRes(types.String)
	},
	"mondoo.asset.resources": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooAsset).GetResources()).ToDataRes(types.Array(types.Resource("mondoo.resource")))
	},
	"mondoo.resource.name": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooResource).GetName()).ToDataRes(types.String)
	},
	"mondoo.resource.id": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondooResource).GetId()).ToDataRes(types.String)
	},
}

func GetData(resource plugin.Resource, field string, args map[string]*llx.RawData) *plugin.DataRes {
	f, ok := getDataFields[resource.MqlName()+"."+field]
	if !ok {
		return &plugin.DataRes{Error: "cannot find '" + field + "' in resource '" + resource.MqlName() + "'"}
	}

	return f(resource)
}

var setDataFields = map[string]func(r plugin.Resource, v *llx.RawData) bool {
	"mondoo.client.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlMondooClient).__id, ok = v.Value.(string)
			return
		},
	"mondoo.client.mrn": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooClient).Mrn, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.organization.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlMondooOrganization).__id, ok = v.Value.(string)
			return
		},
	"mondoo.organization.name": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooOrganization).Name, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.organization.mrn": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooOrganization).Mrn, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.organization.spaces": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooOrganization).Spaces, ok = plugin.RawToTValue[[]interface{}](v.Value, v.Error)
		return
	},
	"mondoo.space.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlMondooSpace).__id, ok = v.Value.(string)
			return
		},
	"mondoo.space.name": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooSpace).Name, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.space.mrn": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooSpace).Mrn, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.space.assets": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooSpace).Assets, ok = plugin.RawToTValue[[]interface{}](v.Value, v.Error)
		return
	},
	"mondoo.asset.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlMondooAsset).__id, ok = v.Value.(string)
			return
		},
	"mondoo.asset.name": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).Name, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.asset.mrn": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).Mrn, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.asset.platform": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).Platform, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.asset.annotations": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).Annotations, ok = plugin.RawToTValue[map[string]interface{}](v.Value, v.Error)
		return
	},
	"mondoo.asset.labels": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).Labels, ok = plugin.RawToTValue[map[string]interface{}](v.Value, v.Error)
		return
	},
	"mondoo.asset.updatedAt": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).UpdatedAt, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"mondoo.asset.scoreValue": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).ScoreValue, ok = plugin.RawToTValue[int64](v.Value, v.Error)
		return
	},
	"mondoo.asset.scoreGrade": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).ScoreGrade, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.asset.resources": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooAsset).Resources, ok = plugin.RawToTValue[[]interface{}](v.Value, v.Error)
		return
	},
	"mondoo.resource.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlMondooResource).__id, ok = v.Value.(string)
			return
		},
	"mondoo.resource.name": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooResource).Name, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.resource.id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondooResource).Id, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
}

func SetData(resource plugin.Resource, field string, val *llx.RawData) error {
	f, ok := setDataFields[resource.MqlName() + "." + field]
	if !ok {
		return errors.New("[mondoo] cannot set '"+field+"' in resource '"+resource.MqlName()+"', field not found")
	}

	if ok := f(resource, val); !ok {
		return errors.New("[mondoo] cannot set '"+field+"' in resource '"+resource.MqlName()+"', type does not match")
	}
	return nil
}

func SetAllData(resource plugin.Resource, args map[string]*llx.RawData) error {
	var err error
	for k, v := range args {
		if err = SetData(resource, k, v); err != nil {
			return err
		}
	}
	return nil
}

// mqlMondooClient for the mondoo.client resource
type mqlMondooClient struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlMondooClientInternal it will be used here
	Mrn plugin.TValue[string]
}

// createMondooClient creates a new instance of this resource
func createMondooClient(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlMondooClient{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	// to override __id implement: id() (string, error)

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("mondoo.client", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlMondooClient) MqlName() string {
	return "mondoo.client"
}

func (c *mqlMondooClient) MqlID() string {
	return c.__id
}

func (c *mqlMondooClient) GetMrn() *plugin.TValue[string] {
	return &c.Mrn
}

// mqlMondooOrganization for the mondoo.organization resource
type mqlMondooOrganization struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlMondooOrganizationInternal it will be used here
	Name plugin.TValue[string]
	Mrn plugin.TValue[string]
	Spaces plugin.TValue[[]interface{}]
}

// createMondooOrganization creates a new instance of this resource
func createMondooOrganization(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlMondooOrganization{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	if res.__id == "" {
	res.__id, err = res.id()
		if err != nil {
			return nil, err
		}
	}

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("mondoo.organization", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlMondooOrganization) MqlName() string {
	return "mondoo.organization"
}

func (c *mqlMondooOrganization) MqlID() string {
	return c.__id
}

func (c *mqlMondooOrganization) GetName() *plugin.TValue[string] {
	return &c.Name
}

func (c *mqlMondooOrganization) GetMrn() *plugin.TValue[string] {
	return &c.Mrn
}

func (c *mqlMondooOrganization) GetSpaces() *plugin.TValue[[]interface{}] {
	return plugin.GetOrCompute[[]interface{}](&c.Spaces, func() ([]interface{}, error) {
		if c.MqlRuntime.HasRecording {
			d, err := c.MqlRuntime.FieldResourceFromRecording("mondoo.organization", c.__id, "spaces")
			if err != nil {
				return nil, err
			}
			if d != nil {
				return d.Value.([]interface{}), nil
			}
		}

		return c.spaces()
	})
}

// mqlMondooSpace for the mondoo.space resource
type mqlMondooSpace struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlMondooSpaceInternal it will be used here
	Name plugin.TValue[string]
	Mrn plugin.TValue[string]
	Assets plugin.TValue[[]interface{}]
}

// createMondooSpace creates a new instance of this resource
func createMondooSpace(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlMondooSpace{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	if res.__id == "" {
	res.__id, err = res.id()
		if err != nil {
			return nil, err
		}
	}

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("mondoo.space", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlMondooSpace) MqlName() string {
	return "mondoo.space"
}

func (c *mqlMondooSpace) MqlID() string {
	return c.__id
}

func (c *mqlMondooSpace) GetName() *plugin.TValue[string] {
	return &c.Name
}

func (c *mqlMondooSpace) GetMrn() *plugin.TValue[string] {
	return &c.Mrn
}

func (c *mqlMondooSpace) GetAssets() *plugin.TValue[[]interface{}] {
	return plugin.GetOrCompute[[]interface{}](&c.Assets, func() ([]interface{}, error) {
		if c.MqlRuntime.HasRecording {
			d, err := c.MqlRuntime.FieldResourceFromRecording("mondoo.space", c.__id, "assets")
			if err != nil {
				return nil, err
			}
			if d != nil {
				return d.Value.([]interface{}), nil
			}
		}

		return c.assets()
	})
}

// mqlMondooAsset for the mondoo.asset resource
type mqlMondooAsset struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlMondooAssetInternal it will be used here
	Name plugin.TValue[string]
	Mrn plugin.TValue[string]
	Platform plugin.TValue[string]
	Annotations plugin.TValue[map[string]interface{}]
	Labels plugin.TValue[map[string]interface{}]
	UpdatedAt plugin.TValue[*time.Time]
	ScoreValue plugin.TValue[int64]
	ScoreGrade plugin.TValue[string]
	Resources plugin.TValue[[]interface{}]
}

// createMondooAsset creates a new instance of this resource
func createMondooAsset(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlMondooAsset{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	if res.__id == "" {
	res.__id, err = res.id()
		if err != nil {
			return nil, err
		}
	}

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("mondoo.asset", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlMondooAsset) MqlName() string {
	return "mondoo.asset"
}

func (c *mqlMondooAsset) MqlID() string {
	return c.__id
}

func (c *mqlMondooAsset) GetName() *plugin.TValue[string] {
	return &c.Name
}

func (c *mqlMondooAsset) GetMrn() *plugin.TValue[string] {
	return &c.Mrn
}

func (c *mqlMondooAsset) GetPlatform() *plugin.TValue[string] {
	return &c.Platform
}

func (c *mqlMondooAsset) GetAnnotations() *plugin.TValue[map[string]interface{}] {
	return &c.Annotations
}

func (c *mqlMondooAsset) GetLabels() *plugin.TValue[map[string]interface{}] {
	return &c.Labels
}

func (c *mqlMondooAsset) GetUpdatedAt() *plugin.TValue[*time.Time] {
	return &c.UpdatedAt
}

func (c *mqlMondooAsset) GetScoreValue() *plugin.TValue[int64] {
	return &c.ScoreValue
}

func (c *mqlMondooAsset) GetScoreGrade() *plugin.TValue[string] {
	return &c.ScoreGrade
}

func (c *mqlMondooAsset) GetResources() *plugin.TValue[[]interface{}] {
	return plugin.GetOrCompute[[]interface{}](&c.Resources, func() ([]interface{}, error) {
		if c.MqlRuntime.HasRecording {
			d, err := c.MqlRuntime.FieldResourceFromRecording("mondoo.asset", c.__id, "resources")
			if err != nil {
				return nil, err
			}
			if d != nil {
				return d.Value.([]interface{}), nil
			}
		}

		return c.resources()
	})
}

// mqlMondooResource for the mondoo.resource resource
type mqlMondooResource struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlMondooResourceInternal it will be used here
	Name plugin.TValue[string]
	Id plugin.TValue[string]
}

// createMondooResource creates a new instance of this resource
func createMondooResource(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlMondooResource{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	if res.__id == "" {
	res.__id, err = res.id()
		if err != nil {
			return nil, err
		}
	}

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("mondoo.resource", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlMondooResource) MqlName() string {
	return "mondoo.resource"
}

func (c *mqlMondooResource) MqlID() string {
	return c.__id
}

func (c *mqlMondooResource) GetName() *plugin.TValue[string] {
	return &c.Name
}

func (c *mqlMondooResource) GetId() *plugin.TValue[string] {
	return &c.Id
}
