/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package SpringBoot

import (
	"os"

	"github.com/go-spring/go-spring/boot-starter"
	"github.com/go-spring/go-spring/spring-core"
)

// ctx 全局的 SpringContext 变量
var ctx = SpringCore.NewDefaultSpringContext()

// RunApplication 快速启动 SpringBoot 应用
func RunApplication(configLocation ...string) {

	app := newApplication(&defaultApplicationContext{
		SpringContext: ctx,
	}, configLocation)

	// 设置运行环境
	if profile, ok := os.LookupEnv(SpringProfile); ok {
		ctx.SetProfile(profile)
	}

	// 设置是否允许注入私有字段
	if access, ok := os.LookupEnv(SpringAccess); ok {
		ctx.SetAllAccess(access == "all")
	}

	BootStarter.Run(app)
}

// Exit 退出 SpringBoot 应用
func Exit() {
	BootStarter.Exit()
}

//////////////// SpringContext ////////////////////////

// GetProfile 返回运行环境
func GetProfile() string {
	return ctx.GetProfile()
}

// SetProfile 设置运行环境
func SetProfile(profile string) {
	ctx.SetProfile(profile)
}

// AllAccess 返回是否允许访问私有字段
func AllAccess() bool {
	return ctx.AllAccess()
}

// SetAllAccess 设置是否允许访问私有字段
func SetAllAccess(allAccess bool) {
	ctx.SetAllAccess(allAccess)
}

// RegisterBean 注册单例 Bean，不指定名称，重复注册会 panic。
func RegisterBean(bean interface{}) *SpringCore.BeanDefinition {
	return ctx.RegisterBean(bean)
}

// RegisterNameBean 注册单例 Bean，需指定名称，重复注册会 panic。
func RegisterNameBean(name string, bean interface{}) *SpringCore.BeanDefinition {
	return ctx.RegisterNameBean(name, bean)
}

// RegisterBeanFn 注册单例构造函数 Bean，不指定名称，重复注册会 panic。
func RegisterBeanFn(fn interface{}, tags ...string) *SpringCore.BeanDefinition {
	return ctx.RegisterBeanFn(fn, tags...)
}

// RegisterNameBeanFn 注册单例构造函数 Bean，需指定名称，重复注册会 panic。
func RegisterNameBeanFn(name string, fn interface{}, tags ...string) *SpringCore.BeanDefinition {
	return ctx.RegisterNameBeanFn(name, fn, tags...)
}

// RegisterMethodBean 注册成员方法单例 Bean，不指定名称，重复注册会 panic。
// selector 可以是 *BeanDefinition，可以是 BeanId，还可以是 (Type)(nil) 变量。
func RegisterMethodBean(selector interface{}, method string, tags ...string) *SpringCore.BeanDefinition {
	return ctx.RegisterMethodBean(selector, method, tags...)
}

// RegisterNameMethodBean 注册成员方法单例 Bean，需指定名称，重复注册会 panic。
// selector 可以是 *BeanDefinition，可以是 BeanId，还可以是 (Type)(nil) 变量。
func RegisterNameMethodBean(name string, selector interface{}, method string, tags ...string) *SpringCore.BeanDefinition {
	return ctx.RegisterNameMethodBean(name, selector, method, tags...)
}

// WireBean 绑定外部的 Bean 源
func WireBean(bean interface{}) {
	ctx.WireBean(bean)
}

// GetBean 根据类型获取单例 Bean，若多于 1 个则 panic；找到返回 true 否则返回 false。
func GetBean(i interface{}) bool {
	return ctx.GetBean(i)
}

// GetBeanByName 根据名称和类型获取单例 Bean，若多于 1 个则 panic；找到返回 true 否则返回 false。
func GetBeanByName(beanId string, i interface{}) bool {
	return ctx.GetBeanByName(beanId, i)
}

// CollectBeans 收集数组或指针定义的所有符合条件的 Bean 对象，收集到返回 true，否则返回 false。
func CollectBeans(i interface{}) bool {
	return ctx.CollectBeans(i)
}

// FindBeanByName 根据名称和类型获取单例 Bean，若多于 1 个则 panic；找到返回 true 否则返回 false。
func FindBeanByName(beanId string) (*SpringCore.BeanDefinition, bool) {
	return ctx.FindBeanByName(beanId)
}

// FindBean 获取单例 Bean，若多于 1 个则 panic；找到返回 true 否则返回 false。
// selector 可以是 BeanId，还可以是 (Type)(nil) 变量，Type 为接口类型时带指针。
func FindBean(selector interface{}) (*SpringCore.BeanDefinition, bool) {
	return ctx.FindBean(selector)
}

// GetAllBeanDefinitions 获取所有 Bean 的定义，一般仅供调试使用。
func GetAllBeanDefinitions() []*SpringCore.BeanDefinition {
	return ctx.GetAllBeanDefinitions()
}

// GetProperty 返回属性值，属性名称统一转成小写。
func GetProperty(name string) interface{} {
	return ctx.GetProperty(name)
}

// GetBoolProperty 返回布尔型属性值，属性名称统一转成小写。
func GetBoolProperty(name string) bool {
	return ctx.GetBoolProperty(name)
}

// GetIntProperty 返回有符号整型属性值，属性名称统一转成小写。
func GetIntProperty(name string) int64 {
	return ctx.GetIntProperty(name)
}

// GetUintProperty 返回无符号整型属性值，属性名称统一转成小写。
func GetUintProperty(name string) uint64 {
	return ctx.GetUintProperty(name)
}

// GetFloatProperty 返回浮点型属性值，属性名称统一转成小写。
func GetFloatProperty(name string) float64 {
	return ctx.GetFloatProperty(name)
}

// GetStringProperty 返回字符串型属性值，属性名称统一转成小写。
func GetStringProperty(name string) string {
	return ctx.GetStringProperty(name)
}

// GetDefaultProperty 返回属性值，如果没有找到则使用指定的默认值，属性名称统一转成小写。
func GetDefaultProperty(name string, defaultValue interface{}) (interface{}, bool) {
	return ctx.GetDefaultProperty(name, defaultValue)
}

// SetProperty 设置属性值，属性名称统一转成小写。
func SetProperty(name string, value interface{}) {
	ctx.SetProperty(name, value)
}

// GetPrefixProperties 返回指定前缀的属性值集合，属性名称统一转成小写。
func GetPrefixProperties(prefix string) map[string]interface{} {
	return ctx.GetPrefixProperties(prefix)
}

// GetAllProperties 返回所有的属性值，属性名称统一转成小写。
func GetAllProperties() map[string]interface{} {
	return ctx.GetAllProperties()
}

// BindProperty 根据类型获取属性值，属性名称统一转成小写。
func BindProperty(name string, i interface{}) {
	ctx.BindProperty(name, i)
}

// BindPropertyIf 根据类型获取属性值，属性名称统一转成小写。
func BindPropertyIf(name string, i interface{}, allAccess bool) {
	ctx.BindPropertyIf(name, i, allAccess)
}
