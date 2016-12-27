$(document).ready(function() {
    //后台广播事件
    $("button[name=btnbroad]").click(function() {
        sendroommsg();
    });

    $("textarea[name=Remark]").keydown(function(e) {
        if (e.keyCode == 13) {
            sendroommsg();
        }
    });

    socketdata();
});

function socketdata() {
    socket.on('all showmsg', function(msg) {
        //消息提示
        layer.msg(msg, {
            icon: 2
        });
    });

    socket.on('all success', function(msg) {
        //消息提示
        layer.msg(msg, {
            icon: 1
        });
    });
}

//后台广播事件
function sendroommsg() {
    var selindex = $("select[name=role]").get(0).selectedIndex;
    if (selindex > 0) {
        var msg = $('#remarkcontent').html();
        if (msg.length <= 0) {
            return;
        } else {
            $('#remarkcontent').html('');
            $("button[name=btnbroad]").focus(); //设置焦点位置
            var seltext = $("select[name=role]").find("option:selected").text();

            var sendmsg = {}; //发送的内容
            sendmsg["Author"] = $("input[name=webUsername]").val();
            sendmsg["Content"] = msg;
            switch (selindex) {
                case 1:
                    sendmsg["Codeid"] = companycode + "_all"; //公司房间标识符
                    break;
                default:
                    sendmsg["Codeid"] = companycode + "_" + seltext; //公司房间标识符
                    break;
            }
            sendmsg["WebIp"] = $("input[name=webIp]").val(); //IP
            sendmsg["WebPro"] = $("input[name=webPro]").val(); //省市
            var sendstr = base64encode(sendmsg);
            socket.emit('all roombroadmsg', sendstr);
            oTable.fnReloadAjax(oTable.fnSettings());
        }
    } else {
        $('#remarkcontent').html('');
        layer.msg("请选择房间号。。。。。。。", {
            icon: 2
        });
    }
}
