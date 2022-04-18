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

package services

import (
	"easygoadmin/app/constant"
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

var Ad = new(adService)

type adService struct{}

func (s *adService) GetList(req dto.AdPageReq) ([]vo.AdInfoVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Ad)).Filter("mark", 1)
	// 广告标题查询
	if req.Title != "" {
		query = query.Filter("title__contains", req.Title)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Ad, 0)
	query.All(&lists)

	// 数据处理
	result := make([]vo.AdInfoVo, 0)
	for _, v := range lists {
		item := vo.AdInfoVo{}
		item.Ad = v
		// 广告封面
		if v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 广告类型
		if v.Type > 0 {
			item.TypeName = constant.AD_TYPE_LIST[v.Type]
		}
		// 所属广告位
		if v.AdSortId > 0 {
			adSortInfo := &models.AdSort{Id: v.AdSortId}
			err := adSortInfo.Get()
			if err == nil {
				item.AdSortDesc = adSortInfo.Description + " >> " + gconv.String(adSortInfo.LocId)
			}
		}
		result = append(result, item)
	}
	// 返回结果
	return result, count, nil
}

func (s *adService) Add(req dto.AdAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 富文本处理
	content := utils.SaveImageContent(req.Content, req.Title, "ad")

	// 实例化对象
	var entity models.Ad
	entity.Title = req.Title
	entity.AdSortId = req.AdSortId
	entity.Type = req.Type
	entity.Description = req.Description
	entity.Content = content
	entity.Url = req.Url
	entity.Width = req.Width
	entity.Height = req.Height
	// 开始日期转换
	startTime, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
	entity.StartTime = startTime
	// 结束日期转换
	endTime, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)
	entity.EndTime = endTime
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1

	// 广告封面
	if entity.Type == 1 {
		// 图片
		cover, err := utils.SaveImage(req.Cover, "ad")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		entity.Cover = ""
	}

	// 插入数据
	return entity.Insert()
}

func (s *adService) Update(req dto.AdUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Ad{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}

	// 富文本处理
	content := utils.SaveImageContent(req.Content, req.Title, "ad")

	// 设置对象
	entity.Title = req.Title
	entity.AdSortId = req.AdSortId
	entity.Type = req.Type
	entity.Description = req.Description
	entity.Content = content
	entity.Url = req.Url
	entity.Width = req.Width
	entity.Height = req.Height
	// 开始日期转换
	startTime, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
	entity.StartTime = startTime
	// 结束日期转换
	endTime, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)
	entity.EndTime = endTime
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 广告封面
	if entity.Type == 1 {
		// 图片
		cover, err := utils.SaveImage(req.Cover, "ad")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		entity.Cover = ""
	}

	// 更新信息
	return entity.Update()
}

func (s *adService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Ad{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Ad{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *adService) Status(req dto.AdStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Ad{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}