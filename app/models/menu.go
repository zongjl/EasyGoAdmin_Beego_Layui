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

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Menu struct {
	Id         int       `orm:"column(id);auto" description:"唯一性标识"`
	Name       string    `orm:"column(name);size(30)" description:"菜单名称"`
	Icon       string    `orm:"column(icon);size(50);null" description:"图标"`
	Url        string    `orm:"column(url);size(150);null" description:"URL地址"`
	Param      string    `orm:"column(param);size(150);null" description:"参数"`
	Pid        int       `orm:"column(pid)" description:"上级ID"`
	Type       int       `orm:"column(type);null" description:"类型：1模块 2导航 3菜单 4节点"`
	Permission string    `orm:"column(permission);size(150);null" description:"权限标识"`
	Status     int       `orm:"column(status)" description:"是否显示：1显示 2不显示"`
	Target     int       `orm:"column(target);null" description:"打开方式：1内部打开 2外部打开"`
	Note       string    `orm:"column(note);size(255);null" description:"菜单备注"`
	Sort       int       `orm:"column(sort);null" description:"显示顺序"`
	CreateUser int       `orm:"column(create_user);null" description:"添加人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	UpdateUser int       `orm:"column(update_user);null" description:"更新人"`
	Mark       int       `orm:"column(mark)" description:"有效标识"`
}

func (t *Menu) TableName() string {
	return "sys_menu"
}

func init() {
	orm.RegisterModel(new(Menu))
}

// 根据条件查询单条数据
func (t *Menu) Get() error {
	err := orm.NewOrm().QueryTable(new(Menu)).Filter("id", t.Id).One(t)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return errors.New("查询到了多条记录")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return errors.New("未查询到记录")
	}
	return nil
}

// 插入数据
func (t *Menu) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Menu) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Menu) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}