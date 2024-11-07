package instance

import (
	"fmt"
	"reflect"
)

// ApplicationContext定义
type ApplicationContext struct {
	instanceTypeMap map[reflect.Type]interface{}
	instanceNameMap map[string]interface{}
}

// NewApplicationContext创建一个新的ApplicationContext
func NewApplicationContext() *ApplicationContext {
	return &ApplicationContext{
		instanceTypeMap: make(map[reflect.Type]interface{}),
		instanceNameMap: make(map[string]interface{}),
	}
}

// RegisterInstance注册一个实例
func (ctx *ApplicationContext) RegisterInstance(instance interface{}) {
	ctx.RegisterInstanceWithName(instance, GetStructName(instance))
}

func (ctx *ApplicationContext) RegisterInstanceWithName(instance interface{}, instanceName string) {
	instanceType := getInstanceType(instance)
	if instanceName == "" {
		instanceName = GetStructName(instance)
	}

	// 不允许重复注册
	if _, exists := ctx.instanceTypeMap[instanceType]; exists {
		panic("duplicate register instance error: " + instanceName)
	}
	if _, exists := ctx.instanceNameMap[instanceName]; exists {
		panic("duplicate register instance error: " + instanceName)
	}

	ctx.instanceTypeMap[instanceType] = instance
	ctx.instanceNameMap[instanceName] = instance
}

func GetStructName(i interface{}) string {
	// 通过反射获取类型信息
	t := reflect.TypeOf(i)
	// 如果是指针，获取指针指向的元素类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

// GetInstance根据类型获取实例
func (ctx *ApplicationContext) GetInstanceByType(instanceType interface{}) (interface{}, error) {
	refType := getInstanceType(instanceType)
	instance, exists := ctx.instanceTypeMap[refType]
	if !exists {
		return nil, fmt.Errorf("no instance found for type: %v", refType)
	}
	return instance, nil
}

// GetInstance根据InstanceName获取实例
func (ctx *ApplicationContext) GetInstanceByName(instanceName string) (interface{}, error) {
	instance, exists := ctx.instanceNameMap[instanceName]
	if !exists {
		return nil, fmt.Errorf("no instance found for name: %v", instanceName)
	}
	return instance, nil
}

// MustGetInstance获取实例，如果不存在则panic
func (ctx *ApplicationContext) MustGetInstance(instanceTypeOrName interface{}) interface{} {
	if instanceName, ok := instanceTypeOrName.(string); ok {
		instance, err := ctx.GetInstanceByName(instanceName)
		if err != nil {
			panic(err)
		}
		return instance
	} else {
		instance, err := ctx.GetInstanceByType(instanceTypeOrName)
		if err != nil {
			panic(err)
		}
		return instance
	}
}

func getInstanceType(instance interface{}) reflect.Type {
	return reflect.TypeOf(instance).Elem()
}
