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

package services

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"strings"
	"time"
)

var ItemCate = new(itemCateService)

type itemCateService struct{}

func (s *itemCateService) GetList(req dto.ItemCateQueryReq) []vo.ItemCateInfoVo {
	// 创建查询实例
	query := orm.NewOrm().QueryTable(new(models.ItemCate)).Filter("mark", 1)
	// 部门名称
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []models.ItemCate
	query.All(&list)

	// 数据处理
	var result []vo.ItemCateInfoVo
	for _, v := range list {
		item := vo.ItemCateInfoVo{}
		item.ItemCate = v
		// 站点封面
		if v.IsCover == 1 && v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 获取栏目
		if v.ItemId > 0 {
			var itemInfo models.Item
			orm.NewOrm().QueryTable(new(models.Item)).Filter("id", item.Id).One(&itemInfo)
			item.ItemName = itemInfo.Name
		}
		// 加入数组
		result = append(result, item)
	}
	return result
}

func (s *itemCateService) Add(req dto.ItemCateAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.ItemCate
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if entity.IsCover == 1 {
		// 有封面
		if req.Cover != "" {
			cover, err := utils.SaveImage(req.Cover, "item_cate")
			if err != nil {
				return 0, err
			}
			entity.Cover = cover
		}
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *itemCateService) Update(req dto.ItemCateUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.ItemCate{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if entity.IsCover == 1 {
		// 有封面
		if req.Cover != "" {
			cover, err := utils.SaveImage(req.Cover, "item_cate")
			if err != nil {
				return 0, err
			}
			entity.Cover = cover
		}
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 更新记录
	return entity.Update()
}

func (s *itemCateService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.ItemCate{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.ItemCate{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *itemCateService) GetCateName(cateId int, delimiter string) string {
	// 声明数组
	list := make([]string, 0)
	for {
		if cateId <= 0 {
			// 退出
			break
		}
		// 业务处理
		info := &models.ItemCate{Id: cateId}
		err := info.Get()
		if err != nil {
			break
		}
		// 上级栏目ID
		cateId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		// 拼接字符串
		return strings.Join(list, delimiter)
	}
	return ""
}

// 获取子级菜单
func (s *itemCateService) GetCateTreeList(itemId int, pid int) ([]*vo.CateTreeNode, error) {
	var cateNote vo.CateTreeNode
	// 创建查询实例
	query := orm.NewOrm().QueryTable(new(models.ItemCate)).Filter("mark", 1)
	// 站点ID
	if itemId > 0 {
		query = query.Filter("item_id", itemId)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询所有
	data := make([]models.ItemCate, 0)
	// 指定返回字段id,name,pid
	_, err := query.All(&data, "id", "name", "pid")
	if err != nil {
		return nil, errors.New("系统错误")
	}
	makeCateTree(data, &cateNote)
	return cateNote.Children, nil
}

//递归生成分类列表
func makeCateTree(cate []models.ItemCate, tn *vo.CateTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &vo.CateTreeNode{}
			child.ItemCate = c
			tn.Children = append(tn.Children, child)
			makeCateTree(cate, child)
		}
	}
}

// 数据源转换
func (s *itemCateService) MakeList(data []*vo.CateTreeNode) []map[string]string {
	cateList := make([]map[string]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			item := map[string]string{}
			item["id"] = gconv.String(val.Id)
			item["name"] = val.Name
			cateList = append(cateList, item)

			// 二级栏目
			for _, v := range val.Children {
				item2 := map[string]string{}
				item2["id"] = gconv.String(v.Id)
				item2["name"] = "|--" + v.Name
				cateList = append(cateList, item2)

				// 三级栏目
				for _, vt := range v.Children {
					item3 := map[string]string{}
					item3["id"] = gconv.String(vt.Id)
					item3["name"] = "|--|--" + vt.Name
					cateList = append(cateList, item3)
				}
			}
		}
	}
	return cateList
}
