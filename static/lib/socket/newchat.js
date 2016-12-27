var loginname   = "";             //用户名
var socket_json = {};             //数据内容
var chatname    = "所有人";       //私聊的用户名
var indexlayer  = 0;              //layer的索引值
$(document).ready(function () {
    var msgcount        = 0;                        //私聊的消息个数
    var inputme         = $('div[class=import]');//输入框
    var socketcontent   = $('#chat_size ul');     //内容控件

    // 绑定表情
    //建立websocket连接
    socket.on("connect", function(){ 
        if(loginname.length>0)
        {
            var sendmsg = {};     //发送的内容
            var loginuser = $('input[name=webUsername]').val();
            if(loginuser.length>0){
                sendmsg["islogin"] = "true";
            }else{
                sendmsg["islogin"] = "false";
                sendmsg["userIcon"]=  getnamecookie(1);
            }
            sendmsg["uname"]       =  loginname;
            sendmsg["Codeid"]      =  codeid;//公司房间标识符
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
                    {
                        if (data.kickout!=null){
                            location.href = "/chat/kickout";//location.href实现客户端页面的跳转
                        }else{
                            var k=data.msg.length;
                            loginname                    = utf8to16(base64decode(data.msg[k-1].Uname));
                            socket_json["Authortype"]    = utf8to16(base64decode(data.msg[k-1].Titlerole));
                            if(data.msg[k-1].UserIcon != null){
                                socket_json["UserIcon"]  = utf8to16(base64decode(data.msg[k-1].UserIcon));
                            }

                            //此处设置下登入属性
                            if(socket_json["Authortype"]=="游客"){
                                $('span[name=ykname]').text(loginname);
                            }
                            $('span[name=ykname]').css("color","#FFFF00");

                            $('.newsone .outer_news_ul ul li').remove(); //清除讲师助理列表数据
                            $('.news .outer_news_ul ul li').remove();    //清除游客列表数据
                            for (var i = 0; i <k-1; i++) {
                               loaduserlist(data.msg[i]);
                            };

                            if(indexlayer>0){
                              layer.close(indexlayer);
                           }

                           FunLeftSet();//设置左边列表信息
                        }
                        break;
                    }
                case "broad":
                    loaduserlist(data.msg);
                    break;
                default:
                    break;
            }
        }
    });

    $('div[class=send]').click(function(){
        setsocketjson();
    });
  
    //通过“回车”提交聊天信息
    inputme.keydown(function(e) {
        if (e.keyCode == 13) {
            setsocketjson();

            setTimeout(function(){
                inputme.html('');
            },10);
        }
    });

    function setsocketjson(){
        var msg = inputme.html();
         inputme.html('');
        if (!msg) return;
        if (loginname != socket_json["Author"]) {
            socket_json["Author"]           = loginname;
            socket_json["Codeid"]           =  codeid;//公司房间标识符
            if(!socket_json["Chat"]){
                socket_json["Chat"]         = "allchat";
            }
            if(!socket_json["Authortype"]){
                socket_json["Authortype"]   = "游客";
            }
        }
        socket_json["Content"]        = msg;
        var sendstr = JSON.stringify(eval(socket_json));
        sendstr = base64encode(utf16to8(sendstr));
        socket.emit('all message', sendstr);
    }

    socket.on('all gagtime', function(msg){
        //发言间隔时间提示
        var showmsg="请间隔"+msg+"秒在发言!";
        layer.msg(showmsg, {icon: 1});
    });

    socket.on('all message', function(msg){
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

        //用户聊天点击事件
        contentevent(data);
        //控制滚动条
        FunControlscrollbar(".chat");
    });

    socket.on("all broadmsg", function(msg){
    });

    socket.on('all disconnection',function(msg){
        var data = JSON.stringify(msg);
        data  = jQuery.parseJSON(data);
        if(msg.leave==null){
            if(data.user != null){
                data.user.Uname          = utf8to16(base64decode(data.user.Uname));
                data.user.Titlerole      = utf8to16(base64decode(data.user.Titlerole));
                $(".news .outer_news_ul ul li").each(function(index){
                     var spantext = $(this).attr("title");
                     if(spantext == data.user.Uname){
                        $(".news .outer_news_ul ul li").eq(index).remove();
                     }
                });

                $(".newsone .outer_news_ul ul li").each(function(index){
                     var spantext = $(this).attr("title");
                     if(spantext == data.user.Uname){
                        $(".newsone .outer_news_ul ul li").eq(index).remove();
                     }
                });
            }
        }else{
            getuserlist("leave","","","",codeid);
        }
    });


    socket.on('all kickout',function(msg){
        $.ajax({
            type: "post",
            url: "/chat/user/list",
            dataType: "json",
            data: {
                method:"leave"
            },
            success: function (result) {
                location.href = "/chat/kickout";    //location.href实现客户端页面的跳转
            },
            error: function (msg) { 
                console.log(msg);
            }
        });
    });


    //禁言五分钟
    $(".popup ul li div[class=disablemsg]").click(function(){
        Checkuserrights("gagcontrol","5",loginname,chatname,codeid);
    });

    //恢复发言
    $(".popup ul li div[class=enablemsg]").click(function(){
        Checkuserrights("gagcontrol","0",loginname,chatname,codeid);
    });

    //黑名单操作
    $(".popup ul li div[class=addblacklist]").click(function(){
        Checkuserrights("blacklist","-1",loginname,chatname,codeid);
    });

    //踢出一小时
    $(".popup ul li div[class=out1hour]").click(function(){
        Checkuserrights("kickout","0",loginname,chatname,codeid);
    });

    //对他说
    $(".popup ul li div[class=saidto]").click(function(){
        socket_json["Chat"]     = "sayhim";
        socket_json["Username"] = chatname;                       //私聊的用户名
        //按钮控制
        $('span[class=Object]').text(chatname);//私聊的用户名
        $('span[class=Object]').append("<span class='Object-x'> × </span>");
        $('span[class=Object-x]').click(function(){
            //单击
            socket_json["Chat"]     = "allchat";            //私聊
            socket_json["Username"] = "";                   //私聊的用户名
            chatname                = "所有人";             //私聊的用户名
            $('span[class=Object]').text(chatname);//私聊的用户名
        });
        inputme.focus();//把焦点设置在input上
    }); 

    //私聊
    $(".popup ul li div[class=saidtosecret]").click(function(){
    });

    $(document).bind('click',function(){ 
        $('.popup').css('display','none'); 
        $('.expression').css('display','none');
    });

    //$(".huanfu span").click(function() {
     //var index    = $('.huanfu span').index($(this));
     //$('.huanfu span').removeClass("hook");
     //$('.huanfu span').eq(index).addClass("hook");
     //});

    $(".public").click(function() {
        $("#chat_one").css("display","");
        $("#chat_tow").css("display","none");
        $('.public').css("background-color","#05B105");
        $('.private').css("background-color","rgba(0, 0, 0, 0.25)");
    });
    
    $(".private").click(function() {
        $("#chat_one").css("display","none");
        $("#chat_tow").css("display","");
        $('.public').css("background-color","rgba(0, 0, 0, 0.25)");
        $('.private').css("background-color","#05B105");
    });
    
    
    $(".huanfu span").click(function() {
        var index  = $('.huanfu span').index($(this));
        switch(index){
            case 0:
                break;
            case 1:
                $('.huanfu span').removeClass("hook");
                $('.huanfu span').eq(index).addClass("hook");
                $('body').css("background","url(/static/images/i/img/body_bg.jpg)");
                $('.theme').css("background-color","#663300");
                $('.partition').css("background","url(/static/images/newimg/btn_ge.png) no-repeat center center");
                $('.news_head').css("background-color","#895705");
                $('.Broadcast').css("border-bottom","1px solid #cc9933");
                $('.bottom').css("color","#ffc047");
                $('.bottom').css("background","#8e5e00");
                $("img[name=logo]").attr("src","/static/images/newimg/logo1.jpg");
                break;
            case 2:
                $('.huanfu span').removeClass("hook");
                $('.huanfu span').eq(index).addClass("hook");
                $('body').css("background","url(/static/images/i/img/body_bg_blue.jpg)");
                $('.theme').css("background-color","#0e477c");
                $('.partition').css("background","url(/static/images/newimg/btn_ge_blue.jpg) no-repeat center center");
                $('.news_head').css("background-color","#895705");
                $('.Broadcast').css("border-bottom","1px solid #7dccf3");
                $('.bottom').css("color","white");
                $('.bottom').css("background","#1c5e99");
                $("img[name=logo]").attr("src","/static/images/newimg/logo.jpg");
                break;
            case 3:
                $('.huanfu span').removeClass("hook");
                $('.huanfu span').eq(index).addClass("hook");
                $('body').css("background","url(/static/images/i/img/body_bg_purple.jpg)");
                $('.theme').css("background-color","#572375");
                $('.partition').css("background","url(/static/images/newimg/btn_ge_gray.jpg) no-repeat center center");
                $('.news_head').css("background-color","rgb(101, 15, 142)");
                $('.Broadcast').css("border-bottom","1px solid #7f4897");
                $('.bottom').css("color","#ffc047");
                $('.bottom').css("background","#642c7e");
                $("img[name=logo]").attr("src","/static/images/newimg/logo2.jpg");
                break;
            case 4:
                $('.huanfu span').removeClass("hook");
                $('.huanfu span').eq(index).addClass("hook");
                $('body').css("background","url(/static/images/i/img/body_bg_gray.jpg)");
                $('.theme').css("background-color","#323232");
                $('.partition').css("background","url(/static/images/newimg/btn_ge_gray1.jpg) no-repeat center center");
                $('.news_head').css("background-color","rgb(27, 24, 24)");
                $('.Broadcast').css("border-bottom","1px solid #434343");
                $('.bottom').css("color","white");
                $('.bottom').css("background","#393939");
                $("img[name=logo]").attr("src","/static/images/newimg/logo3.jpg");
                break;
        }
    });

    $(document).bind('click',function(){ 
        $('.popup').css('display','none'); 
        $('.expression').css('display','none');
    });
    
    //聊天样式事件
    $('.expression').css("display","none");
    $('.face').click(function(e){
        if (e.stopPropagation){
            e.stopPropagation(); 
        }else{
            e.cancelBubble = true; 
        }
        $('.expression').css("display","");
        Faceurlload();//加载图片
     });

    $('div[class=import]').bind('DOMNodeInserted', function(e) { 
    }); 


    //轮播
    $(".carousel").css('display','none');
    //设置宽度
    var sWidth = $(".carousel").width(); //获取焦点图的宽度（显示面积）
    var len = $("#focus ul li").length; //获取焦点图个数
    var index = 0;
    var picTimer;
    
    //以下代码添加数字按钮和按钮后的半透明长条
    var btn = "<div class='btnBg'></div><div class='btn'>";
    for(var i=0; i < len; i++) {
        btn += "<span>" + (i+1) + "</span>";
    }
    btn += "</div>"
    $("#focus").append(btn);
    $("#focus .btnBg").css("opacity",0.5);
    
    //为数字按钮添加鼠标滑入事件，以显示相应的内容
    $("#focus .btn span").mouseenter(function() {
        index = $("#focus .btn span").index(this);
        showPics(index);
    }).eq(0).trigger("mouseenter");
    
    //本例为左右滚动，即所有li元素都是在同一排向左浮动，所以这里需要计算出外围ul元素的宽度
    $("#focus ul").css("width",sWidth * (len + 1));
    
    //鼠标滑入某li中的某div里，调整其同辈div元素的透明度，由于li的背景为黑色，所以会有变暗的效果
    $("#focus ul li div").hover(function() {
        $(this).siblings().css("opacity",0.7);
    },function() {
        $("#focus ul li div").css("opacity",1);
    });
    
    //鼠标滑上焦点图时停止自动播放，滑出时开始自动播放
    $("#focus").hover(function() {
        clearInterval(picTimer);
    },function() {
        picTimer = setInterval(function() {
            sWidth = $(".carousel").width(); //获取焦点图的宽度（显示面积）
            if(index == len) { //如果索引值等于li元素个数，说明最后一张图播放完毕，接下来要显示第一张图，即调用showFirPic()，然后将索引值清零
                showFirPic();
                index = 0;
            } else { //如果索引值不等于li元素个数，按普通状态切换，调用showPics()
                showPics(index);
            }
            index++;
        },3000); //此3000代表自动播放的间隔，单位：毫秒
    }).trigger("mouseleave");
    
    //显示图片函数，根据接收的index值显示相应的内容
    function showPics(index) { //普通切换
        var nowLeft = -index*sWidth; //根据index值计算ul元素的left值
        $("#focus ul").stop(true,false).animate({"left":nowLeft},sWidth/2+10); //通过animate()调整ul元素滚动到计算出的position
        $("#focus .btn span").removeClass("on").eq(index).addClass("on"); //为当前的按钮切换到选中的效果
    }
    
    function showFirPic() { //最后一张图自动切换到第一张图时专用
        $("#focus ul").append($("#focus ul li:first").clone());
        var nowLeft = -len*sWidth; //通过li元素个数计算ul元素的left值，也就是最后一个li元素的右边
        $("#focus ul").stop(true,false).animate({"left":nowLeft},sWidth/2+10,function() {
            //通过callback，在动画结束后把ul元素重新定位到起点，然后删除最后一个复制过去的元素
            $("#focus ul").css("left","0");
            $("#focus ul li:last").remove();
        }); 
        $("#focus .btn span").removeClass("on").eq(0).addClass("on"); //为第一个按钮添加选中的效果
    }
});

function Faceurlload(){
    var tableface=$('.expression table tbody');
    $('.expression table tbody tr').remove();//先清除显示的列表数据
    for (var i = 0; i < 5; i++) {
        var tabcontent="<tr>";
        for (var j = 1; j <10; j++) {
            var tabfacesrc="/static/images/face/"+(i*9+j).toString()+".gif";
            tabcontent+="<td>"+"<img src="+tabfacesrc+" style='cursor:pointer;' name='showface'>"+"</td>";
        };
        tabcontent+="</tr>";
        tableface.append(tabcontent);

        for (var j = 1; j <10; j++) {
            $('img[name=showface]').eq(i*9+j-1).click(function(e){
                var sendface=$(this).attr("src");
                sendface="<img src="+sendface+">";
                //var msg = $('div[class=import]').html()+sendface;
                //$('div[class=import]').html(msg);
                //$('div[class=import]').select();
                $('div[class=import]').focus();
                insertHtmlAtCaret(sendface);
            });  
        };
    };
}

//左边列表高度设置
 function FunLeftSet(){
    var mydiv=$("#newslist");
    var divheight =$(".main").height()-$(".left_logo").height()-$(".left_menu ul").height()-$("#newslistone").height()-30;
    mydiv.css("height",divheight);
 }

//右边聊天高度设置
  function FunRightSet(){
   var bottom_size=$("#chat_size");
    var chatheight =$(".main").height()-257;
    bottom_size.css("height",chatheight);
 }

//控制轮播
function FunCenterSet(){
    var playheight=$("#playheight");
    var play =document.documentElement.clientHeight-$(".carousel").height()-95;
    playheight.css("height",play+"px");
 }

function contentevent(data){
}

window.onload=function  () {
    //轮播设置
    windowbanner();
    //聊天历史记录用户信息发送
    var loginuser = $('input[name=webUsername]').val();
    if(loginuser.length>0){
        loginname                 =loginuser;
        socket_json["Authortype"] = "";
    }else{
        loginname                 =getnamecookie(0);
        socket_json["Authortype"] = "游客";
    }
    getuserlist("chatdata","",loginname,"",codeid);
}

function chatulcontent(data,sel){
    var connectmsg="";
    switch(sel){
        case 0:
        {
            var authorcss =new Array();
            if(data.Authorcss != null){
                var n=data.Authorcss.length;
                for (var i=0; i<n;i++) {
                     authorcss[i]= utf8to16(base64decode(data.Authorcss[i])); 
                };
            }
            var csslength=authorcss.length;
            connectmsg="<li>";
            if(data.Authortype=="游客"){
                connectmsg+="<img class='title' src='/static/images/i/img/guest.png' class='grade-picture'>";
            }else{
                for (var i = 0; i<csslength; i++) {
                    var urlcss="/upload/usertitle/"+authorcss[i];
                    connectmsg+="<img src="+urlcss+" class='title'>";
                };
            }

             if(data.Author==loginname){
                connectmsg+="<div><span class='timesend'><span name='chatauthor' style='cursor: pointer;' title="+data.Author+"></span>";
             }else{
                 connectmsg+="<div><span class='timesend'><span name='chatauthor' style='cursor: pointer;' title="+data.Author+"></span>";
             }
            connectmsg+=data.Author+"  ";
            connectmsg+=data.Newtime;
            connectmsg+="</span><div class='chat_word'>";
            connectmsg+=data.Content;
            connectmsg+="</div></div></li>";
            break;
        }
        case 1:
        {
            var authorcss =new Array();
            var usercss   =new Array();
            if(data.Authorcss != null){
                var n=data.Authorcss.length;
                for (var i=0; i<n;i++) {
                     authorcss[i]= utf8to16(base64decode(data.Authorcss[i])); 
                };
            }
            if(data.Usercss != null){
                var n=data.Usercss.length;
                for (var i=0; i<n;i++) {
                     usercss[i]= utf8to16(base64decode(data.Usercss[i])); 
                };
            }
            var csslength=authorcss.length;
            connectmsg="<li>";
            if(data.Authortype=="游客"){
                connectmsg+="<img src='/static/images/i/img/guest.png' class='title'>";
            }else{
                for (var i = 0; i<csslength; i++) {
                    var urlcss="/upload/usertitle/"+authorcss[i];
                    connectmsg+="<img src="+urlcss+" class='title'>";
                };
            }
            connectmsg+="<div><span class='timesend'><span class='leftFD' name='chatauthor' style='cursor: pointer;' title="+data.Author+">";
            connectmsg+=data.Author;
            connectmsg+="</span> <span class='toward'>对</span>";
            csslength=usercss.length;
            if(data.Usertype=="游客"){
                connectmsg+="<img src='/static/images/i/img/guest.png' class='title'>";
            }else{
                for (var i = 0; i<csslength; i++) {
                    var urlcss="/upload/usertitle/"+usercss[i];
                    connectmsg+="<img src="+urlcss+" class='title'>";
                };
            }
            connectmsg+="<span class='leftFD' name='chatuser' style='cursor: pointer;' title="+data.Username+"></span>";
            connectmsg+=data.Username;
            connectmsg+="</span></span><div class='chat_word'>";
            connectmsg+=data.Content;
            connectmsg+="</div></div></li>";
            break;
        }
        default:
            break;
    }
    return connectmsg;
}

//检查用户权限
function Checkuserrights(mode,opeart,uname,objname,codeid){
   $.ajax({
        type: "post",
        url: "/chat/user/list",
        dataType: "json",
        data: {
            method:"checkrole",
            myname:uname,
            username:objname,
            ucodeid:codeid
        },
        success: function (result) {
            if(result.msg!= null){
                if(result.msg=="yes"){
                    switch(mode){
                        case "gagcontrol":
                        {
                            switch(opeart){
                                case "0":
                                {
                                    var showmsg=objname+"用户将恢复发言！";
                                    layer.alert(showmsg, {
                                        title:'恢复发言提示',
                                        icon: 7,
                                        skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                                    }, function(){
                                        getuserlist("gagcontrol","0",uname,objname,codeid);
                                    });
                                    break;
                                }
                                case "5":
                                {
                                    var showmsg=objname+"用户将禁言五分钟！";
                                    layer.alert(showmsg, {
                                        title:'禁言五分钟提示',
                                        icon: 7,
                                        skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                                    }, function(){
                                        getuserlist("gagcontrol","5",uname,objname,codeid);
                                    });
                                    break;
                                }
                                default:
                                    break;
                            }
                            break;
                        }
                        case "blacklist":
                        {
                            var showmsg=objname+"用户将加入黑名单！";
                            layer.alert(showmsg, {
                                title:'黑名单提示',
                                icon: 7,
                                skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                            }, function(){
                                getuserlist("blacklist","-1",uname,objname,codeid);
                            });
                            break;
                        }
                        case "kickout":
                        {
                            var showmsg=objname+"用户将被踢出一小时！";
                            layer.alert(showmsg, {
                                title:'踢出一小时提示',
                                icon: 7,
                                skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                            }, function(){
                                getuserlist("kickout","0",uname,objname,codeid);
                            });
                            break;
                        }
                        default:
                            break;
                    }
                }else{
                    var showmsg=uname+"用户无此权限!";
                    layer.msg(showmsg, {icon: 1});
                }
            }
        },
        error: function (msg) {
            console.log(msg);
        }
    });
}

//获取信息
function getuserlist(mode,menutext,uname,username,codeid){
    $.ajax({
        type: "post",
        url: "/chat/user/list",
        dataType: "json",
        data: {
            method:mode,
            myname:uname,
            username:username,
            mydata:menutext,
            ucodeid:codeid
        },
        success: function (result) {
            switch(mode){
                case "read":
                {
                    indexlayer = layer.load(2, {//layer的索引值
                      shade: [0.1,'#fff'], //0.1透明度的白色背景
                      offset: ['300px','110px']
                    });

                    var data = JSON.stringify(result);
                    data  = jQuery.parseJSON(data);
                    $('#Border-list .Userbox li').remove();//先清除显示的列表数据
                    if(data.msg != null){
                       var listlen=data.msg.length;
                       for (var i = 0; i <listlen; i++) {
                            loaduserlist(data.msg[i]);
                       };
                    }
                    if(indexlayer>0){
                      layer.close(indexlayer);
                   }
                    break;
                }
                case "chatdata":
                {
                    indexlayer = layer.load(2, {//layer的索引值
                      shade: [0.1,'#fff'], //0.1透明度的白色背景
                      offset: ['300px','110px']
                    });

  
                    var sendmsg = {};     //发送的内容
                    if(socket_json["Authortype"]==""){
                        sendmsg["islogin"] = "true";
                    }else{
                        sendmsg["islogin"]      = "false";
                        sendmsg["userIcon"]     =  getnamecookie(1);
                    }
                    sendmsg["uname"]=loginname;
                    sendmsg["Codeid"]      =  codeid;//公司房间标识符
                    var sendstr = JSON.stringify(eval(sendmsg));
                    sendstr = base64encode(utf16to8(sendstr));
                    socket.emit('all connection', sendstr);

                    //聊天历史记录
                    var data = JSON.stringify(result);
                    data  = jQuery.parseJSON(data);

                    //历史记录
                    if(data.historydata != null){
                        var chatlength = data.historydata.length;
                        for (var i = 0; i <chatlength; i++) {
                            var markdata = data.historydata[i];
                            markdata.Author     = utf8to16(base64decode(markdata.Author));
                            markdata.Authortype = utf8to16(base64decode(markdata.Authortype));
                            markdata.Content    = utf8to16(base64decode(markdata.Content));
                            markdata.Chat       = utf8to16(base64decode(markdata.Chat));
                            var connectmsg ="";
                            switch(markdata.Chat){
                                case  "allchat":
                                    connectmsg=chatulcontent(markdata,0);
                                    break;
                                case  "sayhim":
                                case "privatechat":
                                    markdata.Username     = utf8to16(base64decode(markdata.Username));
                                    markdata.Usertype     = utf8to16(base64decode(markdata.Usertype));
                                    connectmsg=chatulcontent(markdata,1);
                                    break;
                                default:
                                    break;
                            }
                            $('#chat_size ul').append(connectmsg);
                            //用户聊天点击事件
                            contentevent(markdata);
                        };
                        //控制滚动条
                        FunControlscrollbar(".chat");
                    }
                    break;
                }
                case "gagcontrol":
                {
                    switch(menutext){
                        case "0":
                        {
                            //提示操作
                            var showmsg=username+"用户恢复发言,执行成功!";
                            layer.msg(showmsg, {icon: 1});
                            break;
                        }
                        case "5":
                        {
                            //提示操作
                            var showmsg=username+"用户禁言5分钟,执行成功!";
                            layer.msg(showmsg, {icon: 1});
                            break;
                        }
                        default:
                            break;
                    }
                    break;
                }
                case "blacklist":
                {
                    //提示操作
                    var showmsg=username+"用户加入黑名单,执行成功!";
                    layer.msg(showmsg, {icon: 1});
                    break;
                }
                case "kickout":
                {
                    var sendmsg = {};     //发送的内容
                    sendmsg["Uname"]  =  uname;
                    sendmsg["Objname"]=  username;
                    sendmsg["Codeid"] =  codeid;//公司房间标识符
                    var sendstr = JSON.stringify(eval(sendmsg));
                    sendstr = base64encode(utf16to8(sendstr));
                    socket.emit('all kickout', sendstr);

                    //提示操作
                    var showmsg=username+"用户踢出一小时,执行成功!";
                    layer.msg(showmsg, {icon: 1});
                    break;
                }
                case "leave":
                    location.href = "/";//location.href实现客户端页面的跳转
                    break;
                default:
                    break;
            } 
        },
        error: function (msg) { 
            console.log(msg);
        }
    });
}

//加载左边列表信息
function loaduserlist(obj){
    var menujs = $('.newsone .outer_news_ul ul');
    var menuyk = $('.news .outer_news_ul ul');
    var data = obj;
    data.Uname          = utf8to16(base64decode(obj.Uname));
    data.Titlerole      = utf8to16(base64decode(obj.Titlerole));
    if(obj.UserIcon != null){
        data.UserIcon   = utf8to16(base64decode(obj.UserIcon));
    }

    var  existindex = 0;
    var username    = data.Uname;
    var contentmsg  ="<li title="+username+" style='display: table; opacity: 1;'>";
    contentmsg+="<span class='user_name'>";
    contentmsg+=username;
    contentmsg+="</span><span>";

    var authorcss =new Array();
    if(data.Authorcss != null){
        var n=data.Authorcss.length;
        for (var i=0; i<n;i++) {
             authorcss[i]= utf8to16(base64decode(data.Authorcss[i])); 
             break;
        };
    }
    var csslength=authorcss.length;
    if(data.Titlerole=="游客"){
        contentmsg+="<img src='/static/images/i/img/guest.png' width='48'>";
    }else{
        for (var i = 0; i<csslength; i++) {
            var urlcss="/upload/usertitle/"+authorcss[i];
            contentmsg+="<img src="+urlcss+" width='48'>";
            break;
        };
    }
    contentmsg+="</span></li>";
    if(data.Titlerole=="会员"||data.Titlerole=="游客"){
        $(".news .outer_news_ul ul li").each(function(index){
             var spantext = $(this).attr("title");
             if(spantext == data.Uname){
                existindex++;
             }
        });
        if(existindex==0){
            menuyk.append(contentmsg);
        }
    }else{
        $(".newsone .outer_news_ul ul li").each(function(index){
             var spantext = $(this).attr("title");
             if(spantext == data.Uname){
                existindex++;
             }
        });
        if(existindex==0){
            menujs.append(contentmsg);
        }
    }
}


//轮播函数
function windowbanner(){
    $(".carousel").css('display','');
    $(".carousel #focus").css("width",$(".carousel").width());
    $(".carousel #focus ul li").css("width",$(".carousel").width());
    $(".carousel #focus").css("height",$(".carousel").height());
    $(".carousel #focus ul").css("height",$(".carousel").height());
    $(".carousel #focus ul li").css("height",$(".carousel").height()+20);
    $(".carousel #focus .btnBg").css("bottom",0);
    $(".carousel #focus .btn span").css("margin-top",0);
}
