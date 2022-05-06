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

// 字典项列表查询条件
type DictDataPageReq struct {
	DictId int    `form:"dictId"` // 字典ID
	Name   string `form:"name"`   // 字典项名称
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加字典项
type DictDataAddReq struct {
	Name   string `form:"Name,unique" validate:"required"` // 字典项名称
	Code   string `form:"Code" validate:"required"`        // 字典项值
	DictId int    `form:"DictId" validate:"int"`           // 字典类型ID
	Note   string `form:"Note"`                            // 备注
	Sort   int    `form:"Sort" validate:"int"`             // 显示顺序
}

// 添加字典项表单验证
func (a DictDataAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "字典项名称不能为空.",
		"Code.required": "字典项编码不能为空.",
		"DictId.int":    "请选择数据字典.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新字典项
type DictDataUpdateReq struct {
	Id     int    `form:"Id" validate:"int"`
	Name   string `form:"Name,unique" validate:"required"` // 字典项名称
	Code   string `form:"Code" validate:"required"`        // 字典项值
	DictId int    `form:"DictId" validate:"int"`           // 字典类型ID
	Note   string `form:"Note"`                            // 备注
	Sort   int    `form:"Sort" validate:"int"`             // 显示顺序
}

// 更新字典项表单验证
func (a DictDataUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "字典项ID不能为空.",
		"Name.required": "字典项名称不能为空.",
		"Code.required": "字典项编码不能为空.",
		"DictId.int":    "请选择数据字典.",
		"Sort.int":      "排序不能为空.",
	}
}
