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

/**
 * 用户管理
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
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Realname', width: 120, title: '用户姓名', align: 'center'}
            , {field: 'Gender', width: 60, title: '性别', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Gender == 1) {
                        // 男
                        cls = "layui-btn-normal";
                    } else if (d.Gender == 2) {
                        // 女
                        cls = "layui-btn-danger";
                    } else if (d.Gender == 3) {
                        // 保密
                        cls = "layui-btn-warm";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.GenderName + '</span>';
                }
            }
            , {field: 'Nickname', width: 100, title: '用户昵称', align: 'center'}
            , {field: 'Avatar', width: 80, title: '头像', align: 'center', templet: function (d) {
                    if (d.Avatar != "") {
                        return '<a href="' + d.Avatar + '" target="_blank"><img src="' + d.Avatar + '" height="26" /></a>';
                    }
                }
            }
            , {field: 'Username', width: 100, title: '登录名', align: 'center'}
            , {field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.Status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'DeptName', width: 150, title: '所属部门', align: 'center'}
            , {field: 'LevelName', width: 120, title: '职级名称', align: 'center'}
            , {field: 'PositionName', width: 120, title: '岗位名称', align: 'center'}
            , {field: 'Mobile', width: 130, title: '手机号码', align: 'center'}
            , {field: 'Email', width: 200, title: '邮箱地址', align: 'center'}
            , {field: 'Birthday', width: 120, title: '出生日期', align: 'center', templet:"<div>{{layui.util.toDateString(d.Birthday, 'yyyy-MM-dd')}}</div>"}
            , {field: 'Sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 250, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            if (layEvent === 'resetPwd') {
                layer.confirm('您确定要初始化当前用户的密码吗？', {
                    icon: 3,
                    skin: 'layer-ext-moon',
                    btn: ['确认', '取消'] //按钮
                }, function (index) {
                    //初始化密码
                    var url = cUrl + "/resetPwd";
                    // 切记采用FormData表单提交
                    var formData = new FormData();
                    formData.append("id", data.Id);
                    func.ajaxPost(url, formData, function (data, success) {
                        console.log("重置密码：" + (success ? "成功" : "失败"));
                        // 关闭弹窗
                        layer.close(index);
                    })
                });
            }
        });

        //【设置弹框】
        func.setWin("用户");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
