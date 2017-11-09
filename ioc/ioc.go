package ioc

import (
	"reflect"
)

type typeInfo struct {
	depType []reflect.Type
	builder reflect.Value
}

var typeBuilder map[reflect.Type]typeInfo = map[reflect.Type]typeInfo{}

func dfs(t reflect.Type, visit map[reflect.Type]bool, cache map[reflect.Type]reflect.Value, myTypeBuilder map[reflect.Type]typeInfo) reflect.Value {
	result, isExist := cache[t]
	if isExist {
		return result
	}
	_, isVisit := visit[t]
	if isVisit {
		panic("loop dependence")
	}
	visit[t] = true

	info, isExist := myTypeBuilder[t]
	if isExist == false {
		panic("unknown type")
	}
	args := []reflect.Value{}
	for _, singleDepType := range info.depType {
		args = append(args, dfs(singleDepType, visit, cache, myTypeBuilder))
	}
	lastResult := info.builder.Call(args)
	cache[t] = lastResult[0]
	return lastResult[0]
}

func New(a interface{}, moc ...interface{}) interface{} {
	var myTypeBuilder map[reflect.Type]typeInfo
	if len(moc) == 0 {
		myTypeBuilder = typeBuilder
	} else {
		myTypeBuilder = map[reflect.Type]typeInfo{}
		for key, value := range typeBuilder {
			myTypeBuilder[key] = value
		}
		a, b := getRegisterInfo(moc[0])
		myTypeBuilder[a] = b
	}

	visit := map[reflect.Type]bool{}
	cache := map[reflect.Type]reflect.Value{}
	targetType := reflect.ValueOf(a).Type()
	if targetType.Elem().Kind() == reflect.Interface {
		targetType = targetType.Elem()
	}
	return dfs(targetType, visit, cache, myTypeBuilder).Interface()
}

func getRegisterInfo(createFun interface{}) (reflect.Type, typeInfo) {
	typeValue := reflect.ValueOf(createFun)
	typeType := typeValue.Type()
	if typeType.Kind() != reflect.Func {
		panic("invalid type")
	}
	numIn := []reflect.Type{}
	for i := 0; i != typeType.NumIn(); i++ {
		numIn = append(numIn, typeType.In(i))
	}
	if typeType.NumOut() != 1 {
		panic("invalid num out")
	}
	numOut := typeType.Out(0)
	return numOut, typeInfo{
		depType: numIn,
		builder: typeValue,
	}
}
func Register(createFun interface{}) {
	a, b := getRegisterInfo(createFun)
	typeBuilder[a] = b
}
