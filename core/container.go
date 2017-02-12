package core

import (
	"reflect"
)

type Context struct {
	preHook  func()
	postHook func()
	beans    map[interface{}]reflect.Value
}

func (this *Context) Run(target interface{}, args []reflect.Value) []reflect.Value {
	result, isOk := this.beans[target]
	if isOk == false {
		panic("Not Exist Target")
	}
	if this.preHook != nil {
		this.preHook()
	}
	resultValue := result.Call(args)
	if this.postHook != nil {
		this.postHook()
	}
	return resultValue
}

func (this *Context) Mock(source interface{}, target interface{}) interface{} {
	oldValue := this.beans[source]
	this.beans[source] = reflect.ValueOf(target)
	return oldValue
}

func (this *Context) PreHook(target func()) {
	this.preHook = target
}

func (this *Context) PostHook(target func()) {
	this.postHook = target
}

var globalBeans map[interface{}]reflect.Value

func init() {
	globalBeans = map[interface{}]reflect.Value{}
}

func RegisterInterface(target interface{}) {
	targetValue := reflect.ValueOf(target).Elem()
	targetType := targetValue.Type()
	if targetType.Kind() != reflect.Func {
		panic("RegisterInterface is not function")
	}
	newTargetValue := reflect.MakeFunc(targetType, func(args []reflect.Value) (results []reflect.Value) {
		context := args[0].Interface().(Context)
		return context.Run(target, args)
	})
	targetValue.Set(newTargetValue)
}

func RegisterBean(source interface{}, target interface{}) {
	globalBeans[source] = reflect.ValueOf(target)
}

func NewContext() Context {
	context := Context{}
	context.beans = map[interface{}]reflect.Value{}
	for key, value := range globalBeans {
		context.beans[key] = value
	}
	return context
}
