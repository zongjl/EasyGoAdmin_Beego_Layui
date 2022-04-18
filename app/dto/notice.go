// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package dto

import "github.com/gookit/validate"

// 分页查询
type NoticePageReq struct {
	Title  string `form:"title"`  // 通知标题
	Source int    `form:"source"` // 通知来源
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加通知公告
type NoticeAddReq struct {
	Title   string `form:"title" validate:"required"`   // 通知标题
	Content string `form:"content" validate:"required"` // 通知内容
	Source  int    `form:"source" validate:"int"`       // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"isTop" validate:"int"`        // 是否置顶：1是 2否
	Status  int    `form:"status" validate:"int"`       // 状态：1已发布 2待发布
	Sort    int    `form:"sort" validate:"int"`         // 排序号
}

// 添加通知表单验证
func (a NoticeAddReq) Messages() map[string]string {
	return validate.MS{
		"Title.required":   "通知标题不能为空.",
		"Content.required": "通知内容不能为空.",
		"Source.int":       "请选择通知来源.",
		"IsTop.int":        "请选择是否置顶.",
		"Status.int":       "请选择通知状态.",
		"Sort.int":         "排序号不能为空.",
	}
}

// 更新通知公告
type NoticeUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Title   string `form:"title" validate:"required"`   // 通知标题
	Content string `form:"content" validate:"required"` // 通知内容
	Source  int    `form:"source" validate:"int"`       // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"isTop" validate:"int"`        // 是否置顶：1是 2否
	Status  int    `form:"status" validate:"int"`       // 状态：1已发布 2待发布
	Sort    int    `form:"sort" validate:"int"`         // 排序号
}

// 添加通知表单验证
func (u NoticeUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":           "通知ID不能为空.",
		"Title.required":   "通知标题不能为空.",
		"Content.required": "通知内容不能为空.",
		"Source.int":       "请选择通知来源.",
		"IsTop.int":        "请选择是否置顶.",
		"Status.int":       "请选择通知状态.",
		"Sort.int":         "排序号不能为空.",
	}
}

// 设置状态
type NoticeStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (s NoticeStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "通知ID不能为空.",
		"Status.int": "请选择通知状态.",
	}
}