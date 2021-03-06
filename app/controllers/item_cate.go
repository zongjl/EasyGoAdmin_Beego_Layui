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

package controllers

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/services"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gookit/validate"
)

var ItemCate = new(ItemCateController)

type ItemCateController struct {
	BaseController
}

func (ctl *ItemCateController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "item_cate/index.html"
}

func (ctl *ItemCateController) List() {
	// 查询对象
	var req dto.ItemCateQueryReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询列表接口
	lists := services.ItemCate.GetList(req)
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: lists,
	})
}

func (ctl *ItemCateController) Edit() {
	// 记录ID
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.ItemCate{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 封面
		if info.IsCover == 1 && info.Cover != "" {
			info.Cover = utils.GetImageUrl(info.Cover)
		}
		// 渲染模板
		ctl.Data["info"] = info
	} else {
		// 渲染模板
		ctl.Data["info"] = &models.ItemCate{}
	}
	// 站点列表
	result := make([]models.Item, 0)
	orm.NewOrm().QueryTable(new(models.Item)).Filter("mark", 1).All(&result)
	var itemList = map[int]string{}
	for _, v := range result {
		itemList[v.Id] = v.Name
	}
	ctl.Data["itemList"] = itemList
	ctl.Layout = "public/form.html"
	ctl.TplName = "item_cate/edit.html"
}

func (ctl *ItemCateController) Add() {
	// 添加对象
	var req dto.ItemCateAddReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 参数校验
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用添加方法
	rows, err := services.ItemCate.Add(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 添加成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (ctl *ItemCateController) Update() {
	// 更新对象
	var req dto.ItemCateUpdateReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 参数校验
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用更新方法
	rows, err := services.ItemCate.Update(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 更新成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (ctl *ItemCateController) Delete() {
	// 记录ID
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.ItemCate.Delete(ids)
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 删除成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (ctl *ItemCateController) GetCateTreeList() {
	itemId, _ := ctl.GetInt("itemId", 0)
	list, err := services.ItemCate.GetCateTreeList(itemId, 0)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 数据源转换
	result := services.ItemCate.MakeList(list)
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: result,
	})
}
