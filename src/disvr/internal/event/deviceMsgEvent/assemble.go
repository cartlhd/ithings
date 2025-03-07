package deviceMsgEvent

import (
	"github.com/i-Things/things/shared/domain/application"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgGateway"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgThing"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

func ToDmDevicesCore(devices []*msgGateway.Device) (ret []*dm.DeviceCore) {
	for _, v := range devices {
		ret = append(ret, &dm.DeviceCore{
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
		})
	}
	return
}

func ToParamValues(tp map[string]msgThing.Param) map[string]application.ParamValue {
	ret := make(map[string]application.ParamValue, len(tp))
	for k, v := range tp {
		ret[k] = ToParamValue(v)
	}
	return ret
}

func ToParamValue(p msgThing.Param) application.ParamValue {
	var ret application.ParamValue
	ret.Type = p.Value.Type
	switch p.Value.Type {
	case schema.DataTypeStruct:
		v, ok := p.Value.Value.(map[string]msgThing.Param)
		if ok == false {
			panic("struct Param is not find")
		}
		val := make(map[string]application.ParamValue, len(v)+1)
		for _, tp := range v {
			val[tp.Identifier] = ToParamValue(tp)
		}
		ret.Value = val
		return ret
	case schema.DataTypeArray:
		array, ok := p.Value.Value.([]any)
		if ok == false {
			panic("array Param is not find")
		}
		val := make([]any, 0, len(array)+1)
		for _, value := range array {
			switch value.(type) {
			case map[string]msgThing.Param:
				valMap := make(map[string]application.ParamValue, len(array)+1)
				for _, tp := range value.(map[string]msgThing.Param) {
					valMap[tp.Identifier] = ToParamValue(tp)
				}
				val = append(val, valMap)
			default:
				val = append(val, value)
			}
		}
		ret.Value = val
		return ret
	default:
		ret.Value = p.Value.Value
		return ret
	}
}
