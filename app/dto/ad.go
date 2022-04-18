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

// 列表查询
type AdPageReq struct {
	Title string `form:"title"` // 广告标题
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加广告
type AdAddReq struct {
	Title       string `form:"title" validate:"required"`       // 广告标题
	AdSortId    int    `form:"adSortId" validate:"int"`         // 广告位ID
	Cover       string `form:"cover"`                           // 广告图片
	Type        int    `form:"type" validate:"int"`             // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" validate:"required"` // 广告描述
	Content     string `form:"content"`                         // 广告内容
	Url         string `form:"url" validate:"required"`         // 广告链接
	Width       int    `form:"width" validate:"int"`            // 广告宽度
	Height      int    `form:"height" validate:"int"`           // 广告高度
	StartTime   string `form:"startTime" validate:"required"`   // 开始时间
	EndTime     string `form:"endTime" validate:"required"`     // 结束时间
	Status      int    `form:"status" validate:"int"`           // 状态：1在用 2停用
	Sort        int    `form:"sort" validate:"int"`             // 排序
	Note        string `form:"note"`                            // 备注
}

// 更新广告
type AdUpdateReq struct {
	Id          int    `form:"id" validate:"int"`
	Title       string `form:"title" validate:"required"`       // 广告标题
	AdSortId    int    `form:"adSortId" validate:"int"`         // 广告位ID
	Cover       string `form:"cover"`                           // 广告图片
	Type        int    `form:"type" validate:"int"`             // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" validate:"required"` // 广告描述
	Content     string `form:"content"`                         // 广告内容
	Url         string `form:"url" validate:"required"`         // 广告链接
	Width       int    `form:"width" validate:"int"`            // 广告宽度
	Height      int    `form:"height" validate:"int"`           // 广告高度
	StartTime   string `form:"startTime" validate:"required"`   // 开始时间
	EndTime     string `form:"endTime" validate:"required"`     // 结束时间
	Status      int    `form:"status" validate:"int"`           // 状态：1在用 2停用
	Sort        int    `form:"sort" validate:"int"`             // 排序
	Note        string `form:"note"`                            // 备注
}

// 设置状态
type AdStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}