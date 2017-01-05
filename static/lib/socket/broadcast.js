$(document).ready(function() {
    //后台广播事件
    $("#remarkcontent").keydown(function(e) {
        if (e.keyCode == 13) {
            sendbroad();
        }
    });
    socketdata();
});

function socketdata() {
    socket.on('all success', function(msg) {
        //消息提示
        layer.msg(msg, {
            icon: 1
        });
        parent.location.reload();
    });
    socket.on('all showmsg', function(msg) {
        //消息提示
        layer.msg(msg, {
            icon: 2
        });
    });
}

//后台广播事件
function sendbroad() {
    var msg = $('#remarkcontent').val();
    if (msg.length <= 0) {
        layer.msg("请输入广播内容", {
            icon: 2
        });
        return;
    } else {
        var sendmsg = {}; //发送的内容
        sendmsg["Author"] = username;
        sendmsg["Content"] = msg;
        sendmsg["Codeid"] = codeid; //公司房间标识符
        sendmsg["WebIp"] = ipaddress; //IP
        var sendstr = base64encode(sendmsg);
        socket.emit('all roombroadmsg', sendstr);
    }
}
