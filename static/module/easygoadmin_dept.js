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

/**
 * 部门管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'Name', width: 250, title: '部门名称', align: 'left'}
            , {field: 'Code', width: 100, title: '部门编码', align: 'center'}
            , {field: 'Fullname', width: 200, title: '部门全称', align: 'center'}
            , {field: 'Type', width: 100, title: '类型', align: 'center', templet(d) {
                if (d.Type == 1) {
                    // 公司
                    return '<span class="layui-btn layui-btn-normal layui-btn-xs">公司</span>';
                } else if (d.Type == 2) {
                    // 子公司
                    return '<span class="layui-btn layui-btn-warm layui-btn-xs">子公司</span>';
                } else if (d.Type == 3) {
                    // 部门
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">部门</span>';
                } else if (d.Type == 4) {
                    // 小组
                    return '<span class="layui-btn layui-btn-primary layui-btn-xs">小组</span>';
                }
            }}
            , {field: 'Note', width: 200, title: '备注', align: 'center'}
            , {field: 'Sort', width: 80, title: '排序', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {width: 220, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList", true);

        //【设置弹框】
        func.setWin("部门", 500, 420);

    }
});