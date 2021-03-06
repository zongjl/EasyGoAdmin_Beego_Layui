// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package dto

import "github.com/gookit/validate"

// 分页查询条件
type MemberPageReq struct {
	Username string `form:"username"` // 用户名
	Gender   int    `form:"gender"`   // 性别（1男 2女 3未知）
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加会员
type MemberAddReq struct {
	Username     string `form:"username,unique" validate:"required"` // 用户名
	Password     string `form:"password"`                            // 登录密码
	MemberLevel  int    `form:"memberLevel" validate:"int"`          // 会员等级
	Realname     string `form:"realname" validate:"required"`        // 真实姓名
	Nickname     string `form:"nickname" validate:"required"`        // 用户昵称
	Gender       int    `form:"gender" validate:"int"`               // 性别（1男 2女 3未知）
	Avatar       string `form:"avatar" validate:"required"`          // 用户头像
	Birthday     string `form:"birthday" validate:"required"`        // 出生日期
	ProvinceCode string `form:"provinceCode" validate:"required"`    // 省份编号
	CityCode     string `form:"cityCode" validate:"required"`        // 市区编号
	DistrictCode string `form:"districtCode" validate:"required"`    // 区县编号
	Address      string `form:"address" validate:"required"`         // 详细地址
	Intro        string `form:"intro"`                               // 个人简介
	Signature    string `form:"signature"`                           // 个性签名
	Device       int    `form:"device" validate:"int"`               // 设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加
	Source       int    `form:"source" validate:"int"`               // 来源：1、APP注册；2、后台添加；
	Status       int    `form:"status" validate:"int"`               // 是否启用 1、启用  2、停用
}

// 添加会员表单验证
func (v MemberAddReq) Messages() map[string]string {
	return validate.MS{
		"Username.required":     "用户名不能为空.",
		"MemberLevel.int":       "请选择会员等级.",
		"Realname.required":     "真实姓名不能为空.",
		"Nickname.required":     "昵称不能为空.",
		"Gender.int":            "请选择性别.",
		"Avatar.required":       "请上传头像.",
		"Birthday.required":     "请选择出生日期.",
		"ProvinceCode.required": "请选择省份.",
		"CityCode.required":     "请选择城市.",
		"DistrictCode.required": "请选择县区.",
		"Device.int":            "请选择设备类型.",
		"Source.int":            "请选择注册来源.",
		"Status.int":            "请选择会员状态.",
	}
}

// 更新会员
type MemberUpdateReq struct {
	Id           int    `form:"id" validate:"int"`
	Username     string `form:"username,unique" validate:"required"` // 用户名
	Password     string `form:"password"`                            // 登录密码
	MemberLevel  int    `form:"memberLevel" validate:"int"`          // 会员等级
	Realname     string `form:"realname" validate:"required"`        // 真实姓名
	Nickname     string `form:"nickname" validate:"required"`        // 用户昵称
	Gender       int    `form:"gender" validate:"int"`               // 性别（1男 2女 3未知）
	Avatar       string `form:"avatar" validate:"required"`          // 用户头像
	Birthday     string `form:"birthday" validate:"required"`        // 出生日期
	ProvinceCode string `form:"provinceCode" validate:"required"`    // 省份编号
	CityCode     string `form:"cityCode" validate:"required"`        // 市区编号
	DistrictCode string `form:"districtCode" validate:"required"`    // 区县编号
	Address      string `form:"address" validate:"required"`         // 详细地址
	Intro        string `form:"intro"`                               // 个人简介
	Signature    string `form:"signature"`                           // 个性签名
	Device       int    `form:"device" validate:"int"`               // 设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加
	Source       int    `form:"source" validate:"int"`               // 来源：1、APP注册；2、后台添加；
	Status       int    `form:"status" validate:"int"`               // 是否启用 1、启用  2、停用
}

// 更新会员表单验证
func (v MemberUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":                "会员ID不能为空.",
		"Username.required":     "用户名不能为空.",
		"MemberLevel.int":       "请选择会员等级.",
		"Realname.required":     "真实姓名不能为空.",
		"Nickname.required":     "昵称不能为空.",
		"Gender.int":            "请选择性别.",
		"Avatar.required":       "请上传头像.",
		"Birthday.required":     "请选择出生日期.",
		"ProvinceCode.required": "请选择省份.",
		"CityCode.required":     "请选择城市.",
		"DistrictCode.required": "请选择县区.",
		"Device.int":            "请选择设备类型.",
		"Source.int":            "请选择注册来源.",
		"Status.int":            "请选择会员状态.",
	}
}

// 设置状态
type MemberStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

func (v MemberStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "会员ID不能为空.",
		"Status.int": "请选择会员状态.",
	}
}
