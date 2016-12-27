var loginname="";                 //用户名
var socket_json = {};             //数据内容
$(function () {
    var inputme         = $('input[class=input-box]');//输入框
    var socketcontent   = $('.chitchat ul');//内容控件

    getuserlist("chatdata",codeid);
    //建立websocket连接
    socket.on("connect", function(){ 
        if(loginname.length>0)
        {
            var sendmsg = {};     //发送的内容
            var loginuser = $('input[name=webUsername]').val();
            if(loginuser.length>0){
                sendmsg["islogin"] = "true";
                sendmsg["uname"] =  loginuser;
            }else{
                sendmsg["islogin"] = "false";
                sendmsg["uname"]   =  getnamecookie(0);
                sendmsg["userIcon"]=  getnamecookie(1);
                socket_json["Authortype"]="游客";
            }
            loginname= sendmsg["uname"];
            sendmsg["Codeid"] =  codeid;//公司房间标识符
            var sendstr = JSON.stringify(eval(sendmsg));
            sendstr = base64encode(utf16to8(sendstr));
            socket.emit('all connection', sendstr);
        }
    });

    socket.on("all connection", function(json){
        var data = JSON.stringify(json);
        data  = jQuery.parseJSON(data);
        if(data.msg != null){
            switch(data.utype){
                case "emit":
                    if (data.kickout!=null){
                        location.href = "/chat/kickout";//location.href实现客户端页面的跳转
                    }
                    break;
                case "broad":
                    var uname    = utf8to16(base64decode(data.msg.Uname));
                    var mydate = new Date();
                    var datatime = mydate.getHours()+":"+mydate.getMinutes();

                    var connectmsg="<li><span class='time'>";
                    connectmsg+=datatime;
                    connectmsg+="</span><span>欢迎</span>";
                    connectmsg+="<span>"+"“ ";
                    connectmsg+=uname;
                    connectmsg+=" ”"+"</span><span>进入房间</span></li>";
                    socketcontent.append(connectmsg);

                    if($('.chitchat ul')[0].scrollHeight>$('.chitchat ul').height()){
                        $('.chitchat ul').scrollTop($('.chitchat ul')[0].scrollHeight-$('.chitchat ul').height());
                    }
                    break;
                default:
                    break;
            }
        }
    });

    socket.on("all message", function(msg){
        var data = JSON.stringify(msg);
        data  = jQuery.parseJSON(data);
        data  = data.msg;
        data.Author     = utf8to16(base64decode(data.Author));
        data.Authortype = utf8to16(base64decode(data.Authortype));
        data.Content    = utf8to16(base64decode(data.Content));
        data.Chat       = utf8to16(base64decode(data.Chat));
        var connectmsg ="";
        switch(data.Chat){
            case  "allchat":
                connectmsg=chatulcontent(data,0);
                break;
            case  "sayhim":
            case "privatechat":
                data.Username     = utf8to16(base64decode(data.Username));
                data.Usertype = utf8to16(base64decode(data.Usertype));
                connectmsg=chatulcontent(data,1);
                break;
            default:
                break;
        }
        socketcontent.append(connectmsg);

        if($('.chitchat ul')[0].scrollHeight>$('.chitchat ul').height()){
            $('.chitchat ul').scrollTop($('.chitchat ul')[0].scrollHeight-$('.chitchat ul').height());
        }
    });


     socket.on("all broadmsg", function(msg){
        var data = JSON.stringify(msg);
        data  = jQuery.parseJSON(data);
        data  = data.msg;
        data.Author     = utf8to16(base64decode(data.Author));
        data.Content    = utf8to16(base64decode(data.Content));
        
        var connectmsg=chatulcontent(data,0);
        socketcontent.append(connectmsg);
        if($('.chitchat ul')[0].scrollHeight>$('.chitchat ul').height()){
            $('.chitchat ul').scrollTop($('.chitchat ul')[0].scrollHeight-$('.chitchat ul').height());
        }

    });

    $('img[class=send]').click(function(){
        SendPhoneMsg();
    });

    //通过“回车”提交聊天信息
    inputme.keydown(function(e) {
        if (e.keyCode == 13) {
            SendPhoneMsg();
        }
    });

    function SendPhoneMsg(){
        var msg = inputme.val();
        inputme.val('');
        if (!msg) return;
        socket_json["Chat"]        = "allchat";
        socket_json["Author"]      = loginname;
        socket_json["Content"]     = msg;
        socket_json["Codeid"]           =  codeid;//公司房间标识符
        var sendstr = JSON.stringify(eval(socket_json));
        sendstr = base64encode(utf16to8(sendstr));
        socket.emit('all message', sendstr);
    }
});


function chatulcontent(data,sel){
    var connectmsg="";
    switch(sel){
        case 0:
        {
            connectmsg="<li><span class='time'>";
            connectmsg+=data.Time;
            connectmsg+="</span>";
            connectmsg+="<span class='name' style='cursor: pointer;' title="+data.Author+">";
            connectmsg+=data.Author+"：";
            connectmsg+="</span><span>";
            connectmsg+=data.Content;
            connectmsg+="</span></li>";
            break;
        }
        case 1:
        {
            connectmsg="<li><span class='time'>";
            connectmsg+=data.Time;
            connectmsg+="</span>";
            connectmsg+="<span class='name'>";
            connectmsg+=data.Author;
            connectmsg+="</span>";
            connectmsg+="<span class=Dui'>对</span>";
            connectmsg+="<span class='name'>";
            connectmsg+=data.Username+"：";
            connectmsg+="</span>";
            connectmsg+="<span>";
            connectmsg+=data.Content;
            connectmsg+="</span></li>";
            break;
        }
        default:
            break;
    }
    return connectmsg;
}

//发送
function getuserlist(mode,codeid){
    $.ajax({
        type: "post",
        url: "/chat/user/list",
        dataType: "json",
        data: {
            method:mode,
            ucodeid:codeid
        },
        success: function (result) {
           if(mode=="chatdata")
            {
                if(result.uname != null){ 
                    loginname= utf8to16(base64decode(result.uname));

                    var sendmsg = {};     //发送的内容
                    var loginuser = $('input[name=webUsername]').val();
                    if(loginuser.length>0){
                        sendmsg["islogin"] = "true";
                        sendmsg["uname"] =  loginuser;
                    }else{
                        sendmsg["islogin"] = "false";
                        sendmsg["uname"]   =  getnamecookie(0);
                        sendmsg["userIcon"]=  getnamecookie(1);
                        //sendmsg["gagtime"]    =  result.gagtime;  //禁言时间间隔
                        socket_json["gagtime"]  =  result.gagtime;  //禁言时间间隔
                        socket_json["Authortype"]="游客";
                    }
                    loginname= sendmsg["uname"];
                    sendmsg["Codeid"] =  codeid;//公司房间标识符
                    var sendstr = JSON.stringify(eval(sendmsg));
                    sendstr = base64encode(utf16to8(sendstr));
                    socket.emit('all connection', sendstr);
                }
            } 
        },
        error: function (msg) { 
            console.log(msg);
        }
    });
}