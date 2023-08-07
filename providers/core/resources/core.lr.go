// Code generated by resources. DO NOT EDIT.
package resources

import (
	"errors"
	"time"

	"go.mondoo.com/cnquery/llx"
	"go.mondoo.com/cnquery/providers-sdk/v1/plugin"
	"go.mondoo.com/cnquery/types"
)

var resourceFactories map[string]plugin.ResourceFactory

func init() {
	resourceFactories = map[string]plugin.ResourceFactory {
		"mondoo": {
			// to override args, implement: initMondoo(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createMondoo,
		},
		"asset": {
			// to override args, implement: initAsset(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createAsset,
		},
		"time": {
			// to override args, implement: initTime(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createTime,
		},
		"regex": {
			// to override args, implement: initRegex(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createRegex,
		},
		"parse": {
			// to override args, implement: initParse(runtime *plugin.Runtime, args map[string]*llx.RawData) (map[string]*llx.RawData, plugin.Resource, error)
			Create: createParse,
		},
		"uuid": {
			Init: initUuid,
			Create: createUuid,
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
			if x, ok := runtime.Resources[id]; ok {
				return x, nil
			}
			runtime.Resources[id] = res
			return res, nil
		}

		args = cargs
	}

	res, err := f.Create(runtime, args)
	if err != nil {
		return nil, err
	}

	id := name+"\x00"+res.MqlID()
	if x, ok := runtime.Resources[id]; ok {
		return x, nil
	}

	runtime.Resources[id] = res
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
	if x, ok := runtime.Resources[id]; ok {
		return x, nil
	}

	runtime.Resources[id] = res
	return res, nil
}

var getDataFields = map[string]func(r plugin.Resource) *plugin.DataRes{
	"mondoo.version": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondoo).GetVersion()).ToDataRes(types.String)
	},
	"mondoo.build": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondoo).GetBuild()).ToDataRes(types.String)
	},
	"mondoo.arch": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondoo).GetArch()).ToDataRes(types.String)
	},
	"mondoo.jobEnvironment": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondoo).GetJobEnvironment()).ToDataRes(types.Dict)
	},
	"mondoo.capabilities": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlMondoo).GetCapabilities()).ToDataRes(types.Array(types.String))
	},
	"asset.name": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetName()).ToDataRes(types.String)
	},
	"asset.ids": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetIds()).ToDataRes(types.Array(types.String))
	},
	"asset.platform": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetPlatform()).ToDataRes(types.String)
	},
	"asset.kind": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetKind()).ToDataRes(types.String)
	},
	"asset.runtime": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetRuntime()).ToDataRes(types.String)
	},
	"asset.version": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetVersion()).ToDataRes(types.String)
	},
	"asset.arch": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetArch()).ToDataRes(types.String)
	},
	"asset.title": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetTitle()).ToDataRes(types.String)
	},
	"asset.family": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetFamily()).ToDataRes(types.Array(types.String))
	},
	"asset.fqdn": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetFqdn()).ToDataRes(types.String)
	},
	"asset.build": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetBuild()).ToDataRes(types.String)
	},
	"asset.labels": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlAsset).GetLabels()).ToDataRes(types.Map(types.String, types.String))
	},
	"time.now": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetNow()).ToDataRes(types.Time)
	},
	"time.second": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetSecond()).ToDataRes(types.Time)
	},
	"time.minute": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetMinute()).ToDataRes(types.Time)
	},
	"time.hour": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetHour()).ToDataRes(types.Time)
	},
	"time.day": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetDay()).ToDataRes(types.Time)
	},
	"time.today": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetToday()).ToDataRes(types.Time)
	},
	"time.tomorrow": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlTime).GetTomorrow()).ToDataRes(types.Time)
	},
	"regex.ipv4": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetIpv4()).ToDataRes(types.Regex)
	},
	"regex.ipv6": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetIpv6()).ToDataRes(types.Regex)
	},
	"regex.url": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetUrl()).ToDataRes(types.Regex)
	},
	"regex.email": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetEmail()).ToDataRes(types.Regex)
	},
	"regex.mac": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetMac()).ToDataRes(types.Regex)
	},
	"regex.uuid": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetUuid()).ToDataRes(types.Regex)
	},
	"regex.emoji": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetEmoji()).ToDataRes(types.Regex)
	},
	"regex.semver": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetSemver()).ToDataRes(types.Regex)
	},
	"regex.creditCard": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlRegex).GetCreditCard()).ToDataRes(types.Regex)
	},
	"uuid.value": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlUuid).GetValue()).ToDataRes(types.String)
	},
	"uuid.urn": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlUuid).GetUrn()).ToDataRes(types.String)
	},
	"uuid.version": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlUuid).GetVersion()).ToDataRes(types.Int)
	},
	"uuid.variant": func(r plugin.Resource) *plugin.DataRes {
		return (r.(*mqlUuid).GetVariant()).ToDataRes(types.String)
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
	"mondoo.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlMondoo).__id, ok = v.Value.(string)
			return
		},
	"mondoo.version": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondoo).Version, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.build": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondoo).Build, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.arch": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondoo).Arch, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"mondoo.jobEnvironment": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondoo).JobEnvironment, ok = plugin.RawToTValue[interface{}](v.Value, v.Error)
		return
	},
	"mondoo.capabilities": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlMondoo).Capabilities, ok = plugin.RawToTValue[[]interface{}](v.Value, v.Error)
		return
	},
	"asset.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlAsset).__id, ok = v.Value.(string)
			return
		},
	"asset.name": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Name, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.ids": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Ids, ok = plugin.RawToTValue[[]interface{}](v.Value, v.Error)
		return
	},
	"asset.platform": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Platform, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.kind": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Kind, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.runtime": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Runtime, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.version": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Version, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.arch": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Arch, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.title": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Title, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.family": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Family, ok = plugin.RawToTValue[[]interface{}](v.Value, v.Error)
		return
	},
	"asset.fqdn": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Fqdn, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.build": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Build, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"asset.labels": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlAsset).Labels, ok = plugin.RawToTValue[map[string]interface{}](v.Value, v.Error)
		return
	},
	"time.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlTime).__id, ok = v.Value.(string)
			return
		},
	"time.now": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Now, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"time.second": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Second, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"time.minute": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Minute, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"time.hour": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Hour, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"time.day": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Day, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"time.today": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Today, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"time.tomorrow": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlTime).Tomorrow, ok = plugin.RawToTValue[*time.Time](v.Value, v.Error)
		return
	},
	"regex.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlRegex).__id, ok = v.Value.(string)
			return
		},
	"regex.ipv4": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Ipv4, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.ipv6": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Ipv6, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.url": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Url, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.email": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Email, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.mac": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Mac, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.uuid": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Uuid, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.emoji": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Emoji, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.semver": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).Semver, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"regex.creditCard": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlRegex).CreditCard, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"parse.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlParse).__id, ok = v.Value.(string)
			return
		},
	"uuid.__id": func(r plugin.Resource, v *llx.RawData) (ok bool) {
			r.(*mqlUuid).__id, ok = v.Value.(string)
			return
		},
	"uuid.value": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlUuid).Value, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"uuid.urn": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlUuid).Urn, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
	"uuid.version": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlUuid).Version, ok = plugin.RawToTValue[int64](v.Value, v.Error)
		return
	},
	"uuid.variant": func(r plugin.Resource, v *llx.RawData) (ok bool) {
		r.(*mqlUuid).Variant, ok = plugin.RawToTValue[string](v.Value, v.Error)
		return
	},
}

func SetData(resource plugin.Resource, field string, val *llx.RawData) error {
	f, ok := setDataFields[resource.MqlName() + "." + field]
	if !ok {
		return errors.New("[core] cannot set '"+field+"' in resource '"+resource.MqlName()+"', field not found")
	}

	if ok := f(resource, val); !ok {
		return errors.New("[core] cannot set '"+field+"' in resource '"+resource.MqlName()+"', type does not match")
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

// mqlMondoo for the mondoo resource
type mqlMondoo struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlMondooInternal it will be used here
	Version plugin.TValue[string]
	Build plugin.TValue[string]
	Arch plugin.TValue[string]
	JobEnvironment plugin.TValue[interface{}]
	Capabilities plugin.TValue[[]interface{}]
}

// createMondoo creates a new instance of this resource
func createMondoo(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlMondoo{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	// to override __id implement: id() (string, error)

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("mondoo", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlMondoo) MqlName() string {
	return "mondoo"
}

func (c *mqlMondoo) MqlID() string {
	return c.__id
}

func (c *mqlMondoo) GetVersion() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Version, func() (string, error) {
		return c.version()
	})
}

func (c *mqlMondoo) GetBuild() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Build, func() (string, error) {
		return c.build()
	})
}

func (c *mqlMondoo) GetArch() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Arch, func() (string, error) {
		return c.arch()
	})
}

func (c *mqlMondoo) GetJobEnvironment() *plugin.TValue[interface{}] {
	return plugin.GetOrCompute[interface{}](&c.JobEnvironment, func() (interface{}, error) {
		return c.jobEnvironment()
	})
}

func (c *mqlMondoo) GetCapabilities() *plugin.TValue[[]interface{}] {
	return plugin.GetOrCompute[[]interface{}](&c.Capabilities, func() ([]interface{}, error) {
		return c.capabilities()
	})
}

// mqlAsset for the asset resource
type mqlAsset struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlAssetInternal it will be used here
	Name plugin.TValue[string]
	Ids plugin.TValue[[]interface{}]
	Platform plugin.TValue[string]
	Kind plugin.TValue[string]
	Runtime plugin.TValue[string]
	Version plugin.TValue[string]
	Arch plugin.TValue[string]
	Title plugin.TValue[string]
	Family plugin.TValue[[]interface{}]
	Fqdn plugin.TValue[string]
	Build plugin.TValue[string]
	Labels plugin.TValue[map[string]interface{}]
}

// createAsset creates a new instance of this resource
func createAsset(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlAsset{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	// to override __id implement: id() (string, error)

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("asset", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlAsset) MqlName() string {
	return "asset"
}

func (c *mqlAsset) MqlID() string {
	return c.__id
}

func (c *mqlAsset) GetName() *plugin.TValue[string] {
	return &c.Name
}

func (c *mqlAsset) GetIds() *plugin.TValue[[]interface{}] {
	return &c.Ids
}

func (c *mqlAsset) GetPlatform() *plugin.TValue[string] {
	return &c.Platform
}

func (c *mqlAsset) GetKind() *plugin.TValue[string] {
	return &c.Kind
}

func (c *mqlAsset) GetRuntime() *plugin.TValue[string] {
	return &c.Runtime
}

func (c *mqlAsset) GetVersion() *plugin.TValue[string] {
	return &c.Version
}

func (c *mqlAsset) GetArch() *plugin.TValue[string] {
	return &c.Arch
}

func (c *mqlAsset) GetTitle() *plugin.TValue[string] {
	return &c.Title
}

func (c *mqlAsset) GetFamily() *plugin.TValue[[]interface{}] {
	return &c.Family
}

func (c *mqlAsset) GetFqdn() *plugin.TValue[string] {
	return &c.Fqdn
}

func (c *mqlAsset) GetBuild() *plugin.TValue[string] {
	return &c.Build
}

func (c *mqlAsset) GetLabels() *plugin.TValue[map[string]interface{}] {
	return &c.Labels
}

// mqlTime for the time resource
type mqlTime struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlTimeInternal it will be used here
	Now plugin.TValue[*time.Time]
	Second plugin.TValue[*time.Time]
	Minute plugin.TValue[*time.Time]
	Hour plugin.TValue[*time.Time]
	Day plugin.TValue[*time.Time]
	Today plugin.TValue[*time.Time]
	Tomorrow plugin.TValue[*time.Time]
}

// createTime creates a new instance of this resource
func createTime(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlTime{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	// to override __id implement: id() (string, error)

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("time", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlTime) MqlName() string {
	return "time"
}

func (c *mqlTime) MqlID() string {
	return c.__id
}

func (c *mqlTime) GetNow() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Now, func() (*time.Time, error) {
		return c.now()
	})
}

func (c *mqlTime) GetSecond() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Second, func() (*time.Time, error) {
		return c.second()
	})
}

func (c *mqlTime) GetMinute() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Minute, func() (*time.Time, error) {
		return c.minute()
	})
}

func (c *mqlTime) GetHour() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Hour, func() (*time.Time, error) {
		return c.hour()
	})
}

func (c *mqlTime) GetDay() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Day, func() (*time.Time, error) {
		return c.day()
	})
}

func (c *mqlTime) GetToday() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Today, func() (*time.Time, error) {
		return c.today()
	})
}

func (c *mqlTime) GetTomorrow() *plugin.TValue[*time.Time] {
	return plugin.GetOrCompute[*time.Time](&c.Tomorrow, func() (*time.Time, error) {
		return c.tomorrow()
	})
}

// mqlRegex for the regex resource
type mqlRegex struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlRegexInternal it will be used here
	Ipv4 plugin.TValue[string]
	Ipv6 plugin.TValue[string]
	Url plugin.TValue[string]
	Email plugin.TValue[string]
	Mac plugin.TValue[string]
	Uuid plugin.TValue[string]
	Emoji plugin.TValue[string]
	Semver plugin.TValue[string]
	CreditCard plugin.TValue[string]
}

// createRegex creates a new instance of this resource
func createRegex(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlRegex{
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
		args, err = runtime.ResourceFromRecording("regex", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlRegex) MqlName() string {
	return "regex"
}

func (c *mqlRegex) MqlID() string {
	return c.__id
}

func (c *mqlRegex) GetIpv4() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Ipv4, func() (string, error) {
		return c.ipv4()
	})
}

func (c *mqlRegex) GetIpv6() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Ipv6, func() (string, error) {
		return c.ipv6()
	})
}

func (c *mqlRegex) GetUrl() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Url, func() (string, error) {
		return c.url()
	})
}

func (c *mqlRegex) GetEmail() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Email, func() (string, error) {
		return c.email()
	})
}

func (c *mqlRegex) GetMac() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Mac, func() (string, error) {
		return c.mac()
	})
}

func (c *mqlRegex) GetUuid() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Uuid, func() (string, error) {
		return c.uuid()
	})
}

func (c *mqlRegex) GetEmoji() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Emoji, func() (string, error) {
		return c.emoji()
	})
}

func (c *mqlRegex) GetSemver() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Semver, func() (string, error) {
		return c.semver()
	})
}

func (c *mqlRegex) GetCreditCard() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.CreditCard, func() (string, error) {
		return c.creditCard()
	})
}

// mqlParse for the parse resource
type mqlParse struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlParseInternal it will be used here
}

// createParse creates a new instance of this resource
func createParse(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlParse{
		MqlRuntime: runtime,
	}

	err := SetAllData(res, args)
	if err != nil {
		return res, err
	}

	// to override __id implement: id() (string, error)

	if runtime.HasRecording {
		args, err = runtime.ResourceFromRecording("parse", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlParse) MqlName() string {
	return "parse"
}

func (c *mqlParse) MqlID() string {
	return c.__id
}

// mqlUuid for the uuid resource
type mqlUuid struct {
	MqlRuntime *plugin.Runtime
	__id string
	// optional: if you define mqlUuidInternal it will be used here
	Value plugin.TValue[string]
	Urn plugin.TValue[string]
	Version plugin.TValue[int64]
	Variant plugin.TValue[string]
}

// createUuid creates a new instance of this resource
func createUuid(runtime *plugin.Runtime, args map[string]*llx.RawData) (plugin.Resource, error) {
	res := &mqlUuid{
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
		args, err = runtime.ResourceFromRecording("uuid", res.__id)
		if err != nil || args == nil {
			return res, err
		}
		return res, SetAllData(res, args)
	}

	return res, nil
}

func (c *mqlUuid) MqlName() string {
	return "uuid"
}

func (c *mqlUuid) MqlID() string {
	return c.__id
}

func (c *mqlUuid) GetValue() *plugin.TValue[string] {
	return &c.Value
}

func (c *mqlUuid) GetUrn() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Urn, func() (string, error) {
		return c.urn()
	})
}

func (c *mqlUuid) GetVersion() *plugin.TValue[int64] {
	return plugin.GetOrCompute[int64](&c.Version, func() (int64, error) {
		return c.version()
	})
}

func (c *mqlUuid) GetVariant() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Variant, func() (string, error) {
		return c.variant()
	})
}