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
 * 友链管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        ,form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Name', width: 250, title: '友链名称', align: 'center'}
            , {field: 'Image', width: 100, title: '友链图片', align: 'center', templet: function (d) {
                    var imageStr = "";
                    if (d.Image) {
                        imageStr = '<a href="' + d.Image + '" target="_blank"><img src="' + d.Image + '" height="26" /></a>';
                    }
                    return imageStr;
                }
            }
            , {field: 'Type', width: 100, title: '类型', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Type == 1) {
                        // 友情链接
                        cls = "layui-btn-normal";
                    } else if (d.Type == 2) {
                        // 合作伙伴
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.TypeName + '</span>';
                }
            }
            , {field: 'Url', width: 200, title: '友链地址', align: 'center', templet(d) {
                    return "<a href='" + d.Url + "' target='_blank'>" + d.Url + "</a>";
                }
            }
            , {field: 'Platform', width: 100, title: '平台', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Platform == 1) {
                        // PC站
                        cls = "layui-btn-normal";
                    } else if (d.Platform == 2) {
                        // WAP站
                        cls = "layui-btn-danger";
                    } else if (d.Platform == 3) {
                        // 微信小程序
                        cls = "layui-btn-warm";
                    } else if (d.Platform == 4) {
                        // APP应用
                        cls = "layui-btn-primary";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.PlatformName + '</span>';
                }
            }
            , {field: 'Form', width: 100, title: '友链形式', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Form == 1) {
                        // 文字链接
                        cls = "layui-btn-normal";
                    } else if (d.Form == 2) {
                        // 图片链接
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.FormName + '</span>';
                }
            }
            , {field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.Id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.Status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'Sort', width: 100, title: '显示顺序', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("友链");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    } else {
        //【监听友链类型】
        var link_form = $('#form').val();
        if (link_form == 1) {
            //文字
            $(".image").addClass("layui-hide");
        } else if (link_form == 2) {
            //图片
            $(".image").removeClass("layui-hide");
        }
        form.on('select(form)', function (data) {
            if (data.value == 1) {
                $(".image").addClass("layui-hide");
            } else if (data.value == 2) {
                $(".image").removeClass("layui-hide");
            }
        });
    }
});
