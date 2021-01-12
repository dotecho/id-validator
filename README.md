# id-validator

> **中华人民共和国居民身份证**、**中华人民共和国港澳居民居住证**以及**中华人民共和国台湾居民居住证**号码验证工具。

![Tests](https://github.com/guanguans/id-validator/workflows/Tests/badge.svg)
![Check & fix styling](https://github.com/guanguans/id-validator/workflows/Check%20&%20fix%20styling/badge.svg)
[![codecov](https://codecov.io/gh/guanguans/id-validator/branch/main/graph/badge.svg?token=URGFAWS6S4)](https://codecov.io/gh/guanguans/id-validator)
[![Latest Stable Version](https://poser.pugx.org/guanguans/id-validator/v)](//packagist.org/packages/guanguans/id-validator)
[![Total Downloads](https://poser.pugx.org/guanguans/id-validator/downloads)](//packagist.org/packages/guanguans/id-validator)
[![License](https://poser.pugx.org/guanguans/id-validator/license)](//packagist.org/packages/guanguans/id-validator)

## Requirement

* Go >= 1.11

## Installation

``` shell script
$ go get -u github.com/guanguans/id-validator
```

## Usage

``` go
package main

import (
	"github.com/guanguans/id-validator"

	"gopkg.in/ffmt.v1"
)

func main() {

	// 验证身份证号合法性
	ffmt.P(idvalidator.IsValid("440308199901101512"))  // 大陆居民身份证18位
	ffmt.P(idvalidator.IsValid("610104620927690"))     // 大陆居民身份证15位
	ffmt.P(idvalidator.IsValid("810000199408230021"))  // 港澳居民居住证18位
	ffmt.P(idvalidator.IsValid("830000199201300022"))  // 台湾居民居住证18位

	// 获取身份证号信息
	ffmt.P(idvalidator.GetInfo("440308199901101512"))
	// []interface {}[
	// 	github.com/guanguans/id-validator.IdInfo{          // 身份证号信息
	// 		AddressCode: int(440308)                           // 地址码
	// 		Abandoned:   int(0)                                // 地址码是否废弃：1为废弃的，0为正在使用的
	// 		Address:     string("广东省深圳市盐田区")             // 地址
	// 		AddressTree: []string[                             // 省市区三级列表
	//			string("广东省")                                    // 省
	//			string("深圳市")                                    // 市
	//			string("盐田区")                                    // 区
	//		]
	// 		Birthday:      <1999-01-10 00:00:00 +0000 UTC>     // 出生日期
	// 		Constellation: string("摩羯座")                     // 星座
	// 		ChineseZodiac: string("卯兔")                       // 生肖
	// 		Sex:           int(1)                              // 性别：1为男性，0为女性
	// 		Length:        int(18)                             // 号码长度
	// 		CheckBit:      string("2")                         // 校验码
	// 	}
	// 	<nil>                                              // 错误信息
	// ]

	// 生成可通过校验的假身份证号
	ffmt.P(idvalidator.FakeId()) // 随机生成
	ffmt.P(idvalidator.FakeRequireId(true, "江苏省", "200001", 1)) // 生成出生于2000年1月江苏省的男性居民身份证

	// 15位号码升级为18位
	ffmt.P(idvalidator.UpgradeId("610104620927690"))
	// []interface {}[
	// 	string("610104196209276908") // 升级后号码
	// 	<nil>                        // 错误信息
	// ]
}
```

## Testing

``` bash
$ make ...
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

* [guanguans](https://github.com/guanguans)
* [All Contributors](../../contributors)

## 相关项目

* [PHP 版本](https://github.com/jxlwqq/id-validator)
* [Python 版本](https://github.com/jxlwqq/id-validator.py)
* [Ruby 版本](https://github.com/renyijiu/id_validator)
* [JavaScript 版本](https://github.com/mc-zone/IDValidator)

## 参考资料

* [中华人民共和国公民身份号码](https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码)
* [中华人民共和国民政部：行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)
* [中华人民共和国行政区划代码历史数据集](https://github.com/jxlwqq/address-code-of-china)
* [国务院办公厅关于印发《港澳台居民居住证申领发放办法》的通知](http://www.gov.cn/zhengce/content/2018-08/19/content_5314865.htm)
* [港澳台居民居住证](https://zh.wikipedia.org/wiki/港澳台居民居住证)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
