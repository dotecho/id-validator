// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package main

import (
	idvalidator "github.com/dotecho/id-validator"
	"gopkg.in/ffmt.v1"
)

func main() {
	// 验证身份证号合法性
	ffmt.P(idvalidator.IsValid("500154199301135886", true))  // 严格模式验证大陆居民身份证18位
	ffmt.P(idvalidator.IsValid("500154199301135886", false)) // 非严格模式验证大陆居民身份证18位
	ffmt.P(idvalidator.IsValid("11010119900307803X", false)) // 大陆居民身份证末位是X18位
	ffmt.P(idvalidator.IsValid("610104620927690", false))    // 大陆居民身份证15位
	ffmt.P(idvalidator.IsValid("810000199408230021", false)) // 港澳居民居住证18位
	ffmt.P(idvalidator.IsValid("830000199201300022", false)) // 台湾居民居住证18位

	// 获取身份证号信息
	ffmt.P(idvalidator.GetInfo("500154199301135886", true))  // 严格模式获取身份证号信息
	ffmt.P(idvalidator.GetInfo("500154199301135886", false)) // 非严格模式获取身份证号信息

	// 生成可通过校验的假身份证号
	ffmt.P(idvalidator.FakeId())                                // 随机生成
	ffmt.P(idvalidator.FakeRequireId(true, "江苏省", "200001", 1)) // 生成出生于2000年1月江苏省的男性居民身份证

	// 15位号码升级为18位
	ffmt.P(idvalidator.UpgradeId("610104620927690"))
}
