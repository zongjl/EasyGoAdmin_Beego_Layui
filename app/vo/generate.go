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

package vo

// 数据库信息
type GenerateInfo struct {
	Name          string `json:"Name"`          // 表名
	Engine        string `json:"Engine"`        // 引擎
	Version       string `json:"Version"`       // 版本
	Collation     string `json:"Collation"`     // 编码
	Rows          string `json:"Rows"`          // 记录数
	DataLength    string `json:"DataLength"`    // 大小
	AutoIncrement string `json:"AutoIncrement"` // 自增索引
	Comment       string `json:"Comment"`       // 表备注
	CreateTime    string `json:"CreateTime"`    // 添加时间
	UpdateTime    string `json:"UpdateTime"`    // 更新时间
}
