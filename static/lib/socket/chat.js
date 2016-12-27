var loginname = ""; //用户名
var socket_json = {}; //数据内容
var chatname = "所有人"; //私聊的用户名
var indexlayer = 0; //layer的索引值
$(document).ready(function() {
    var msgcount = 0; //私聊的消息个数
    var inputme = $('div[class=inputsend]'); //输入框
    var socketcontent = $('.list1 #right-chat #Dialog'); //内容控件

    //隐藏上传文件图标
    $("input[name=img_upload]").css("display", "none");

    //建立websocket连接
    socket.on("connect", function() {
        if (loginname.length > 0) {
            var sendmsg = {}; //发送的内容
            if (socket_json["AuthorRole"] == "guest") {
                sendmsg["islogin"] = "false";
                sendmsg["userIcon"] = getnamecookie(1);
            } else {
                sendmsg["islogin"] = "true";
            }
            sendmsg["uname"] = loginname;
            sendmsg["Codeid"] = codeid; //公司房间标识符
            var sendstr = JSON.stringify(eval(sendmsg));
            sendstr = base64encode(utf16to8(sendstr));
            socket.emit('all connection', sendstr);
        }
    });

    socket.on("all connection", function(json) {
        var data = JSON.stringify(json);
        data = jQuery.parseJSON(data);
        if (data.msg != null) {
            switch (data.utype) {
                case "emit":
                    {
                        if (data.kickout != null) {
                            location.href = "/chat/kickout"; //location.href实现客户端页面的跳转
                        } else {
                            var k = data.msg.length;
                            loginname = utf8to16(base64decode(data.msg[k - 1].Uname));
                            socket_json["AuthorRole"] = utf8to16(base64decode(data.msg[k - 1].RoleName));
                            socket_json["Authortype"] = utf8to16(base64decode(data.msg[k - 1].Titlerole));
                            if (data.msg[k - 1].UserIcon != null) {
                                socket_json["UserIcon"] = utf8to16(base64decode(data.msg[k - 1].UserIcon));
                            }
                            $('#Border-list .Userbox li').remove(); //先清除显示的列表数据

                            $(".MENU .selectlist").removeClass("selectlist").addClass("noselectlist");
                            $('.MENU span').eq(0).removeClass("noselectlist").addClass("selectlist");


                            for (var i = 0; i < k - 1; i++) {
                                var authorname = utf8to16(base64decode(data.msg[i].RoleName));
                                /*
                                if (authorname != "guest" && authorname != "member") {
                                    loaduserlist(data.msg[i]);
                                }
                                */
                                if (authorname == "member") {
                                    loaduserlist(data.msg[i]);
                                }
                            };

                            //控制用户列表滚动条
                            FunControlscrollbar(".Userbox");
                            if (indexlayer > 0) {
                                layer.close(indexlayer);
                            }
                        }
                        break;
                    }
                case "broad":
                    loaduserlist(data.msg);
                    //控制用户列表滚动条
                    FunControlscrollbar(".Userbox");
                    break;
                default:
                    break;
            }
        }
    });

    $('img[class=send]').click(function() {
        setsocketjson();
    });


    //通过“回车”提交聊天信息
    inputme.keydown(function(e) {
        if (e.keyCode == 13) {
            setsocketjson();

            setTimeout(function() {
                inputme.html('');
            }, 10);
        }
    });

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

    socket.on('all message', function(msg) {
        var data = JSON.stringify(msg);
        data = jQuery.parseJSON(data);
        data = data.msg;
        data.Author = utf8to16(base64decode(data.Author));
        data.AuthorRole = utf8to16(base64decode(data.AuthorRole));
        data.Authortype = utf8to16(base64decode(data.Authortype));
        data.Content = utf8to16(base64decode(data.Content));
        data.Chat = utf8to16(base64decode(data.Chat));
        var connectmsg = "";
        var slectlist = ""; //哪个控件
        switch (data.Chat) {
            case "allchat":
                connectmsg = chatulcontent(data, 0);
                slectlist = ".list1"; //哪个控件
                break;
            case "sayhim":
            case "privatechat":
                data.Username = utf8to16(base64decode(data.Username));
                data.UserRole = utf8to16(base64decode(data.UserRole));
                data.Usertype = utf8to16(base64decode(data.Usertype));
                connectmsg = chatulcontent(data, 1);
                slectlist = ".list1"; //哪个控件
                break;
            default:
                break;
        }
        if (data.Chat == "privatechat") {
            var index = $(".bigone .titleone .selecttab").index('.bigone .titleone a');
            switch (index) {
                case 0:
                    {
                        if (0 == msgcount) {
                            var quantity = $('.bigone .titleone .noselecttab')
                            quantity.append("<span class='quantity'>" + (msgcount + 1).toString() + "</span>");
                        } else {
                            var quantity = $('.bigone .titleone .noselecttab .quantity')
                            if (quantity.length > 0) {
                                quantity.text((msgcount + 1).toString());
                            }
                        }
                        if (data.Author == loginname) {
                            $(".bigone .titleone .selecttab").removeClass("selecttab").addClass("noselecttab");
                            $('.bigone .titleone a').eq(1).removeClass("noselecttab").addClass("selecttab");
                        }
                        msgcount++; //私聊的消息个数
                        break;
                    }
                default:
                    break;
            }
            if ((data.AuthorRole == data.UserRole) && (data.UserRole == "guest" || data.UserRole == "member")) {
                connectmsg = "";
            }
            slectlist = ".list2"; //哪个控件
        }
        var socketcontent = $(slectlist + " #right-chat #Dialog"); //内容控件
        socketcontent.append(connectmsg);

        //用户聊天点击事件
        contentevent(slectlist, data);
    });


    socket.on("all broadmsg", function(msg) {
        var data = JSON.stringify(msg);
        data = jQuery.parseJSON(data);
        data = data.msg;
        data.Author = utf8to16(base64decode(data.Author));
        data.Content = utf8to16(base64decode(data.Content));
        $('marquee').text(data.Content);

        var authorcss = utf8to16(base64decode(data.Authorcss));
        /*
        var authorcss =new Array();
        if(data.Authorcss != null){
            var n=data.Authorcss.length;
            for (var i=0; i<n;i++) {
                 authorcss[i]= utf8to16(base64decode(data.Authorcss[i])); 
            };
        }
        var csslength=authorcss.length;
        */
        var connectmsg = "";
        connectmsg = "<li class='grade'><span class='grade-time'>";
        connectmsg += data.Time;
        connectmsg += "</span>";
        /*
        for (var i = 0; i<csslength; i++) {
            var urlcss="/upload/usertitle/"+authorcss[i];
            connectmsg+="<img src="+urlcss+" class='grade-picture'>";
        };
        */
        var urlcss = "/upload/usertitle/" + authorcss;
        connectmsg += "<img src=" + urlcss + " class='grade-picture'>";
        if (data.Author == loginname) {
            connectmsg += "<span  class='name-color green-color' name='chatauthor' style='cursor: pointer;' title=" + data.Author + ">";
        } else {
            connectmsg += "<span  class='name-color blue-color' name='chatauthor' style='cursor: pointer;' title=" + data.Author + ">";
        }
        connectmsg += data.Author;
        connectmsg += "</span><div><span class='grade-speak'>";
        connectmsg += data.Content;
        connectmsg += "</span></div></li>";
        $('.list1 #right-chat #Dialog').append(connectmsg);

        //控制滚动条
        FunControlscrollbar(".list1 #Dialog");
    });


    //禁言五分钟
    $(".popup ul li div[class=disablemsg]").click(function() {
        Checkuserrights("gagcontrol", "5", loginname, chatname, codeid);
    });

    //恢复发言
    $(".popup ul li div[class=enablemsg]").click(function() {
        Checkuserrights("gagcontrol", "0", loginname, chatname, codeid);
    });

    //黑名单操作
    $(".popup ul li div[class=addblacklist]").click(function() {
        Checkuserrights("blacklist", "-1", loginname, chatname, codeid);
    });

    //踢出一小时
    $(".popup ul li div[class=out1hour]").click(function() {
        Checkuserrights("kickout", "0", loginname, chatname, codeid);
    });

    //对他说
    $(".popup ul li div[class=saidto]").click(function() {
        socket_json["Chat"] = "sayhim";
        socket_json["Username"] = chatname; //私聊的用户名
        //按钮控制
        $('span[class=Object]').text(chatname); //私聊的用户名
        $('span[class=Object]').append("<span class='Object-x'> × </span>");
        $('span[class=Object-x]').click(function() {
            //单击
            socket_json["Chat"] = "allchat"; //私聊
            socket_json["Username"] = ""; //私聊的用户名
            chatname = "所有人"; //私聊的用户名
            $('span[class=Object]').text(chatname); //私聊的用户名
        });
        inputme[0].focus(); //把焦点设置在input上
    });

    //私聊
    $(".popup ul li div[class=saidtosecret]").click(function() {
        var index = $(".bigone .titleone .selecttab").index('.bigone .titleone a');
        switch (index) {
            case 0:
                index = 1;
                taboperation(index);
                break;
            default:
                break;
        }

        var quantity = $('.bigone .titleone .quantity');
        if (quantity.length > 0) {
            quantity.remove();
            msgcount = 0; //私聊的消息个数
        }
        socket_json["Chat"] = "privatechat";
        socket_json["Username"] = chatname; //私聊的用户名

        //按钮控制
        $('span[class=Object]').text(chatname); //私聊的用户名
        $('span[class=Object]').append("<span class='Object-x'> × </span>");
        $('span[class=Object-x]').click(function() {
            //单击
            socket_json["Chat"] = "allchat"; //私聊
            socket_json["Username"] = ""; //私聊的用户名
            chatname = "所有人"; //私聊的用户名
            $('span[class=Object]').text(chatname); //私聊的用户名
            switch (index) {
                case 1:
                    taboperation(0);
                    break;
                default:
                    break;
            }
        });
        inputme[0].focus(); //把焦点设置在input上
    });

    $(document).bind('click', function() {
        $('.popup').css('display', 'none');
        $('.right-image .expression').css('display', 'none');
    });

    socket.on('all disconnection', function(msg) {
        var data = JSON.stringify(msg);
        data = jQuery.parseJSON(data);
        if (msg.leave == null) {
            if (data.user != null) {
                data.user.Uname = utf8to16(base64decode(data.user.Uname));
                data.user.Titlerole = utf8to16(base64decode(data.user.Titlerole));
                //var menutext = $(".MENU .selectlist").text();
                //if(menutext ==  data.user.Titlerole){
                $("#Border-list .Userbox li").each(function(index) {
                    var spantext = $(this).attr("title");
                    if (spantext == data.user.Uname) {
                        $("#Border-list .Userbox li").eq(index).remove();
                    }
                });
            }
        } else {
            getuserlist("leave", "", "", "", codeid);
        }
    });


    socket.on('all kickout', function(msg) {
        $.ajax({
            type: "post",
            url: "/chat/user/list",
            dataType: "json",
            data: {
                method: "leave"
            },
            success: function(result) {
                location.href = "/chat/kickout"; //location.href实现客户端页面的跳转
            },
            error: function(msg) {
                console.log(msg);
            }
        });
    });

    /*
    socket.on('join room', function(json){
    });
    */

    $('.MENU span').click(function() {
        var index = $('.MENU span').index($(this));
        $(".MENU .selectlist").removeClass("selectlist").addClass("noselectlist");
        $('.MENU span').eq(index).removeClass("noselectlist").addClass("selectlist");
        //alert($(".MENU .selectlist").text());
        //alert($(".MENU .selectlist").index('.MENU span'));
        getuserlist("read", index.toString(), loginname, "", codeid);

        //初始化搜索框
        $("#Border-seek").val("搜索成员");
        $("#Border-seek").css('color', '#BCB4B4');
        $("#eliminate").css('display', 'none');
    });

    $('.bigone .titleone a').click(function() {
        var index = $('.bigone .titleone a').index($(this));
        if (1 == index) {
            var quantity = $('.bigone .titleone .quantity');
            if (quantity.length > 0) {
                quantity.remove();
                msgcount = 0; //私聊的消息个数
            }
        }
        taboperation(index);
        //alert($(".bigone .titleone .selecttab").text());
        //alert($(".bigone .titleone .selecttab").index('.bigone .titleone a'));
    });

    //初始化
    $('.list1').show();
    $(".list2").hide();
    $("#Border-seek").val("搜索成员");
    $("#Border-seek").css('color', '#BCB4B4');
    $("img[name=tumbleBody]").attr("src", "/static/images/i/img/05.png");

    $("img[name=clearBody]").click(function() {
        var index = $(".bigone .titleone .selecttab").index('.bigone .titleone a');
        switch (index) {
            case 0:
                $(".list1 #right-chat #Dialog").html("");
                break;
            case 1:
                $(".list2 #right-chat #Dialog").html("");
                break;
            default:
                break;
        }
    });

    $("img[name=tumbleBody]").click(function() {
        if ($(this).attr("src") == "/static/images/i/img/05_dark.png") {
            $(this).attr("src", "/static/images/i/img/05.png");
        } else {
            $(this).attr("src", "/static/images/i/img/05_dark.png");
        }
    });

    /*
    $('#img_upload-button').hover(function(){
       $('#img_upload-button').css("background-image","url(/static/images/i/img/tupian_01.png)");
    },function(){
        $('#img_upload-button').css("background-image","url(/static/images/i/img/tupian_02.png)");
    });
    //$('#img_upload-button').css("background-image","url(/static/images/i/img/tupian_01.png)");
    //$('#img_upload-button').css("background-image","url(/static/images/i/img/tupian_02.png)");
    */

    $('img[name=tuichu]').hover(function() {
        $(this).attr("src", "/static/images/i/img/1_18.png");
    }, function() {
        $(this).attr("src", "/static/images/i/img/1_17.png");
    });

    $('img[name=shezhi]').hover(function() {
        $(this).attr("src", "/static/images/i/img/1_14.png");
    }, function() {
        $(this).attr("src", "/static/images/i/img/1_13.png");
    });

    $('img[name=face]').hover(function() {
        $(this).attr("src", "/static/images/i/img/biaoqing_02.png");
    }, function() {
        $(this).attr("src", "/static/images/i/img/biaoqing_01.png");
    });


    $('img[name=clearBody]').hover(function() {
        $(this).attr("src", "/static/images/i/img/04_bright.png");
    }, function() {
        $(this).attr("src", "/static/images/i/img/04.png");
    });

    $("#Border-seek").focusin(function() {
        if ($(this).val() == "搜索成员") {
            $("#Border-seek").css('color', '#BCB4B4');
            $(this).val("");
            $("#eliminate").css('display', '');
            $("#eliminate").css('cursor', 'pointer');
        }
    });

    $("#Border-seek").focusout(function() {
        if ($(this).val() == "") {
            $(this).val("搜索成员");
            $(this).css('color', '#BCB4B4');
            $("#eliminate").css('display', 'none');
        }
    });

    $("#eliminate").click(function(event) {
        $("#Border-seek").val("搜索成员");
        $("#Border-seek").css('color', '#BCB4B4');
        $(this).css('display', 'none');

        $("#Border-list .Userbox li").each(function(index) {
            $("#Border-list .Userbox li").eq(index).slideDown();
        });
    });

    $("#Border-seek").change(function() {
        var filter = $(this).val();
        $("#Border-list .Userbox li").each(function(index) {
            if (namesearch(filter, $(this).attr("title"))) {
                $("#Border-list .Userbox li").eq(index).slideDown();
            } else {
                $("#Border-list .Userbox li").eq(index).slideUp();
            }
        });
        return false;
    }).keyup(function() {
        $(this).change();
    });


    $('#right-bottom img[name=face]').click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }
        Faceurlload(); //加载图片
    });

    $('.expression a').click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }
        Faceurlload(); //加载图片
    });

    $(".Expert a").click(function() {
        var index = $('.Expert a').index($(this));
        switch (index) {
            case 0:
                $(".Expert_issue").css("display", "none");
                $(".information").css("display", "");
                break;

            case 1:
                $(".Expert_issue").css("display", "");
                $(".information").css("display", "none");
                break;
            case 2:
                break;
        }
    });

    function setsocketjson() {
        /*
        if (chat_sysconfig != null) {
            if (socket_json["AuthorRole"] == "guest" && chat_sysconfig.Guestchat == 0) { //0 禁止聊天 1 允许聊天
                var showmsg = "游客禁止聊天!";
                layer.msg(showmsg, {
                    icon: 2
                });
                return;
            }
        }
        */
        var msg = inputme.html();
        inputme.html('');
        if (!msg) return;
        if (loginname != socket_json["Author"]) {
            socket_json["Author"] = loginname;
            socket_json["Codeid"] = codeid; //公司房间标识符
            if (!socket_json["Chat"]) {
                socket_json["Chat"] = "allchat";
            }
        }
        socket_json["Content"] = msg;
        var sendstr = JSON.stringify(eval(socket_json));
        sendstr = base64encode(utf16to8(sendstr));
        socket.emit('all message', sendstr);
    }
});

function Faceurlload() {
    $('.right-image .expression').show();
    var tableface = $('.expression table tbody');
    $('.expression table tbody tr').remove(); //先清除显示的列表数据
    for (var i = 0; i < 5; i++) {
        var tabcontent = "<tr>";
        for (var j = 1; j < 10; j++) {
            var tabfacesrc = "/static/images/face/" + (i * 9 + j).toString() + ".gif";
            tabcontent += "<td>" + "<img src=" + tabfacesrc + " alt='' class='clap1' name='showface'>" + "</td>";
        };
        tabcontent += "</tr>";
        tableface.append(tabcontent);

        for (var j = 1; j < 10; j++) {
            $('img[name=showface]').eq(i * 9 + j - 1).click(function(e) {
                var sendface = $(this).attr("src");
                sendface = "<img src=" + sendface + ">";
                //var msg = $('div[class=inputsend]').html()+sendface;
                //$('div[class=inputsend]').html(msg);
                $('div[class=inputsend]').focus();
                insertHtmlAtCaret(sendface);
            });
        };
    };
}

function chatulcontent(data, sel) {
    var connectmsg = "";
    switch (sel) {
        case 0:
            {
                var authorcss = utf8to16(base64decode(data.Authorcss));
                /*
                var authorcss =new Array();
                if(data.Authorcss != null){
                    var n=data.Authorcss.length;
                    for (var i=0; i<n;i++) {
                         authorcss[i]= utf8to16(base64decode(data.Authorcss[i])); 
                    };
                }
                var csslength=authorcss.length;
                */
                connectmsg = "<li class='grade'><span class='grade-time'>";
                connectmsg += data.Time;
                connectmsg += "</span>";
                if (data.AuthorRole == "guest") {
                    connectmsg += "<img src='/static/images/i/img/guest.png' class='grade-picture'>";
                } else {
                    /*
                    for (var i = 0; i<csslength; i++) {
                        var urlcss="/upload/usertitle/"+authorcss[i];
                        connectmsg+="<img src="+urlcss+" class='grade-picture'>";
                    };
                    */
                    var urlcss = "/upload/usertitle/" + authorcss;
                    connectmsg += "<img src=" + urlcss + " class='grade-picture'>";
                }
                if (data.Author == loginname) {
                    connectmsg += "<span  class='name-color green-color' name='chatauthor' style='cursor: pointer;' title=" + data.Author + ">";
                } else {
                    connectmsg += "<span  class='name-color blue-color' name='chatauthor' style='cursor: pointer;' title=" + data.Author + ">";
                }
                connectmsg += data.Author;
                connectmsg += "</span><div><span class='grade-speak'>";
                connectmsg += data.Content;
                connectmsg += "</span></div></li>";
                break;
            }
        case 1:
            {
                var authorcss = utf8to16(base64decode(data.Authorcss));
                var usercss = utf8to16(base64decode(data.Usercss));
                /*
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
                */
                connectmsg = "<li class='grade'><span class='grade-time'>";
                connectmsg += data.Time;
                connectmsg += "</span>";
                if (data.AuthorRole == "guest") {
                    connectmsg += "<img src='/static/images/i/img/guest.png' class='grade-picture'>";
                } else {
                    var urlcss = "/upload/usertitle/" + authorcss;
                    connectmsg += "<img src=" + urlcss + " class='grade-picture'>";
                    /*
                    for (var i = 0; i<csslength; i++) {
                        var urlcss="/upload/usertitle/"+authorcss[i];
                        connectmsg+="<img src="+urlcss+" class='grade-picture'>";
                    };
                    */
                }
                connectmsg += "<span class='name1-color green-color' name='chatauthor' style='cursor: pointer;' title=" + data.Author + ">";
                connectmsg += data.Author;
                connectmsg += "</span><span class='Face'>对</span>";
                //csslength=usercss.length;
                if (data.UserRole == "guest") {
                    connectmsg += "<img src='/static/images/i/img/guest.png' class='grade-picture'>";
                } else {
                    var urlcss = "/upload/usertitle/" + usercss;
                    connectmsg += "<img src=" + urlcss + " class='grade-picture'>";
                    /*
                    for (var i = 0; i<csslength; i++) {
                        var urlcss="/upload/usertitle/"+usercss[i];
                        connectmsg+="<img src="+urlcss+" class='grade-picture'>";
                    };
                    */
                }

                connectmsg += "<span class='name-color blue-color' name='chatuser' style='cursor: pointer;' title=" + data.Username + ">";
                connectmsg += data.Username;
                connectmsg += "</span><div><span class='grade-speak'>";
                connectmsg += data.Content;
                connectmsg += "</span></div></li>";
                break;
            }
        default:
            break;
    }
    return connectmsg;
}

function contentevent(slectlist, data) {
    $(slectlist + " span[name=chatauthor]:last").click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }

        //获取鼠标点击的高度    
        var mouseClick; //鼠标点击的Y坐标
        var dialog = $("#Dialog").height(); //聊天窗口高度
        var dialogTop = $("#Dialog").offset().top; //聊天窗口离顶部距离
        e = e || window.event;
        var point = {
            x: 0,
            y: 0
        };
        if (e.pageX || e.pageY) {
            mouseClick = e.pageY;
        } else { //兼容ie  
            mouseClick = e.clientY + document.documentElement.scrollTop;
        }

        //单击
        chatname = $(this).text();
        if (chatname == loginname) {} else {
            if ((socket_json["AuthorRole"] == data.UserRole) && (data.UserRole == "guest" || data.UserRole == "member")) {} else {
                var popupHeight = $(".popup").height(); //权限菜单的高度    192
                var windowHeight = $(window).height(); //当前浏览器的高度     978
                if (mouseClick + popupHeight > dialog + dialogTop) {
                    $(".popup").show().css({
                        "top": $(this).offset().top - popupHeight,
                        "left": $(this).offset().left
                    });
                } else {
                    $(".popup").show().css({
                        "top": $(this).offset().top + 20,
                        "left": $(this).offset().left
                    });
                }
            }
        }
    });


    $(slectlist + " span[name=chatuser]:last").click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }
        //单击
        chatname = $(this).text();
        if (chatname == loginname) {} else {
            if ((socket_json["AuthorRole"] == data.Usertype) && (data.UserRole == "guest" || data.UserRole == "member")) {} else {
                $(".popup").show().css({
                    "top": $(this).offset().top + 20,
                    "left": $(this).offset().left - 20
                });
            }
        }

        //控制滚动条
        FunControlscrollbar(slectlist + " #Dialog");
    });


    if ($("img[name=tumbleBody]").attr("src") == "/static/images/i/img/05.png") {
        //控制滚动条
        FunControlscrollbar(slectlist + " #Dialog");
    }
}


function taboperation(Indexes) {
    $(".bigone .titleone .selecttab").removeClass("selecttab").addClass("noselecttab");
    $('.bigone .titleone a').eq(Indexes).removeClass("noselecttab").addClass("selecttab");
    switch (Indexes) {
        case 0:
            {
                $('.list1').show();
                $(".list2").hide();

                //初始化
                $(".list1 #Dialog").fadeIn();



                if (!FunControlscrollbar(".list2 #Dialog")) {
                    $(".list2 #Dialog").fadeOut();
                }
                break;
            }
        case 1:
            {
                $('.list2').show();
                $(".list1").hide();
                //初始化
                $(".list2 #Dialog").fadeIn();

                if (!FunControlscrollbar(".list1 #Dialog")) {
                    $(".list1 #Dialog").fadeOut();
                }
                break;
            }
        default:
            break;
    }
}


window.onload = function() {
    var linksone = getClass("titleone")[0].getElementsByTagName("a");
    var listsone = getClass("listone");
    tablist(linksone, listsone);

    //初始化
    if ($("#bottom-center ul li").length > 0) {
        var interv = 2000; //切换间隔时间
        var interv2 = 10; //切换速速
        var opac1 = 80; //文字背景的透明度
        var source = "bottom-center" //焦点轮换图片容器的id名称
            //获取对象
        function getTag(tag, obj) {
            if (obj == null) {
                return document.getElementsByTagName(tag)
            } else {
                return obj.getElementsByTagName(tag)
            }
        }

        function getid(id) {
            return document.getElementById(id)
        };
        var opac = 0,
            j = 0,
            t = 63,
            num, scton = 0,
            timer, timer2, timer3;
        var id = getid(source);
        id.removeChild(getTag("div", id)[0]);
        var li = getTag("li", id);
        var div = document.createElement("div");
        var title = document.createElement("div");
        var span = document.createElement("span");
        var button = document.createElement("div");
        button.className = "button";
        for (var i = 0; i < li.length; i++) {
            var a = document.createElement("a");
            a.innerHTML = i + 1;
            a.onclick = function() {
                clearTimeout(timer);
                clearTimeout(timer2);
                clearTimeout(timer3);
                j = parseInt(this.innerHTML) - 1;
                scton = 0;
                t = 63;
                opac = 0;
                fadeon();
            };
            a.className = "b1";
            a.onmouseover = function() {
                this.className = "b2"
            };
            a.onmouseout = function() {
                this.className = "b1";
                sc(j)
            };
            button.appendChild(a);
        }

        //控制图层透明度
        function alpha(obj, n) {
            if (document.all) {
                obj.style.filter = "alpha(opacity=" + n + ")";
            } else {
                obj.style.opacity = (n / 100);
            }
        }
        //控制焦点按钮
        function sc(n) {
            for (var i = 0; i < li.length; i++) {
                button.childNodes[i].className = "b1"
            };
            button.childNodes[n].className = "b2";
        }
        title.className = "num_list";
        alpha(title, opac1);
        id.className = "d1";
        div.className = "d2";
        id.appendChild(div);
        id.appendChild(title);
        id.appendChild(button);
        //渐显
        var fadeon = function() {
            opac += 5;
            div.innerHTML = li[j].innerHTML;
            span.innerHTML = getTag("a", li[j])[0].alt;
            alpha(div, opac);
            if (scton == 0) {
                sc(j);
                num = -2;
                scrolltxt();
                scton = 1
            };
            if (opac < 100) {
                timer = setTimeout(fadeon, interv2)
            } else {
                timer2 = setTimeout(fadeout, interv);
            };
        }

        //渐隐
        var fadeout = function() {
            opac -= 5;
            div.innerHTML = li[j].innerHTML;
            alpha(div, opac);
            if (scton == 0) {
                num = 2;
                scrolltxt();
                scton = 1
            };
            if (opac > 0) {
                timer = setTimeout(fadeout, interv2)
            } else {
                if (j < li.length - 1) {
                    j++
                } else {
                    j = 0
                };
                fadeon()
            };
        }

        //滚动文字
        var scrolltxt = function() {
            t += num;
            span.style.marginTop = t + "px";
            if (num < 0 && t > 3) {
                timer3 = setTimeout(scrolltxt, interv2)
            } else if (num > 0 && t < 62) {
                timer3 = setTimeout(scrolltxt, interv2)
            } else {
                scton = 0
            }
        };
        fadeon();
    }

    //页面设置信息
    windowset();
    //聊天历史记录用户信息发送
    var loginuser = $('input[name=webUsername]').val();
    if (loginuser.length > 0) {
        loginname = loginuser;
        socket_json["AuthorRole"] = "";
    } else {
        loginname = getnamecookie(0);
        socket_json["AuthorRole"] = "guest";
    }
    getuserlist("chatdata", "", loginname, "", codeid);
}

function tablist(links, lists) {
    for (var i = 0; i < links.length; i++) {
        links[i].index = i;
        links[i].onclick = function() {
            for (var j = 0; j < lists.length; j++) {
                lists[j].style.display = "none";
                links[j].style.background = "";
                links[j].style.color = "#000";
            }
            if (lists[this.index] != null) {
                lists[this.index].style.display = "block";
                //this.addClass(themename);
            }
        }
    }
}

function getClass(classname, obj) {
    var obj = obj || document
    var arr = [];
    if (obj.getElementsByClassName) {
        return obj.getElementsByClassName(classname);
    } else {
        var all = obj.getElementsByTagName("*");
        for (var i = 0; i < all.length; i++) {
            if (checkClass(all[i].className, classname)) {
                arr.push(all[i]);
            }
        }
        return arr;
    }
}

function checkClass(oldclass, newclass) {
    var arr = oldclass.split(" ");
    for (var i = 0; i < arr.length; i++) {
        if (arr[i] == newclass) {
            return true;
        }
    }
    return false;
}

//检查用户权限
function Checkuserrights(mode, opeart, uname, objname, codeid) {
    $.ajax({
        type: "post",
        url: "/chat/user/list",
        dataType: "json",
        data: {
            method: "checkrole",
            myname: uname,
            username: objname,
            ucodeid: codeid
        },
        success: function(result) {
            if (result.msg != null) {
                if (result.msg == "yes") {
                    switch (mode) {
                        case "gagcontrol":
                            {
                                switch (opeart) {
                                    case "0":
                                        {
                                            var showmsg = objname + "用户将恢复发言！";
                                            layer.alert(showmsg, {
                                                title: '恢复发言提示',
                                                icon: 7,
                                                skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                                            }, function() {
                                                getuserlist("gagcontrol", "0", uname, objname, codeid);
                                            });
                                            break;
                                        }
                                    case "5":
                                        {
                                            var showmsg = objname + "用户将禁言五分钟！";
                                            layer.alert(showmsg, {
                                                title: '禁言五分钟提示',
                                                icon: 7,
                                                skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                                            }, function() {
                                                getuserlist("gagcontrol", "5", uname, objname, codeid);
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
                                var showmsg = objname + "用户将加入黑名单！";
                                layer.alert(showmsg, {
                                    title: '黑名单提示',
                                    icon: 7,
                                    skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                                }, function() {
                                    getuserlist("blacklist", "-1", uname, objname, codeid);
                                });
                                break;
                            }
                        case "kickout":
                            {
                                var showmsg = objname + "用户将被踢出一小时！";
                                layer.alert(showmsg, {
                                    title: '踢出一小时提示',
                                    icon: 7,
                                    skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                                }, function() {
                                    getuserlist("kickout", "0", uname, objname, codeid);
                                });
                                break;
                            }
                        default:
                            break;
                    }
                } else {
                    var showmsg = uname + "用户无此权限!";
                    layer.msg(showmsg, {
                        icon: 1
                    });
                }
            }
        },
        error: function(msg) {
            console.log(msg);
        }
    });
}

//获取信息
function getuserlist(mode, menutext, uname, username, codeid) {
    $.ajax({
        type: "post",
        url: "/chat/user/list",
        dataType: "json",
        data: {
            method: mode,
            myname: uname,
            username: username,
            mydata: menutext,
            ucodeid: codeid
        },
        success: function(result) {
            switch (mode) {
                case "read":
                    {
                        indexlayer = layer.load(2, { //layer的索引值
                            shade: [0.1, '#fff'], //0.1透明度的白色背景
                            offset: ['300px', '110px']
                        });

                        var data = JSON.stringify(result);
                        data = jQuery.parseJSON(data);
                        $('#Border-list .Userbox li').remove(); //先清除显示的列表数据
                        if (data.msg != null) {
                            var listlen = data.msg.length;
                            for (var i = 0; i < listlen; i++) {
                                loaduserlist(data.msg[i]);
                            };
                            //控制用户列表滚动条
                            FunControlscrollbar(".Userbox");
                        }
                        if (indexlayer > 0) {
                            layer.close(indexlayer);
                        }
                        break;
                    }
                case "chatdata":
                    {
                        indexlayer = layer.load(2, { //layer的索引值
                            shade: [0.1, '#fff'], //0.1透明度的白色背景
                            offset: ['300px', '110px']
                        });

                        //数据解析
                        var data = JSON.stringify(result);
                        data = jQuery.parseJSON(data);
                        var sendmsg = {}; //发送的内容
                        if (data.sysconfig != null) {
                            //显示上传文件图标
                            chat_sysconfig = data.sysconfig;
                        }
                        if (socket_json["AuthorRole"] == "") {
                            //显示上传文件图标
                            $("input[name=img_upload]").css("display", "");
                            Img_Fileupload(); //上传文件

                            sendmsg["islogin"] = "true";
                        } else {
                            sendmsg["islogin"] = "false";
                            sendmsg["userIcon"] = getnamecookie(1);
                        }
                        sendmsg["uname"] = loginname;
                        sendmsg["Codeid"] = codeid; //公司房间标识符
                        var sendstr = JSON.stringify(eval(sendmsg));
                        sendstr = base64encode(utf16to8(sendstr));
                        socket.emit('all connection', sendstr);

                        //广播信息
                        if (data.broaddata != null) {
                            $('marquee').text(data.broaddata);
                        }
                        //聊天历史记录
                        if (data.historydata != null) {
                            var chatlength = data.historydata.length;
                            for (var i = 0; i < chatlength; i++) {
                                var markdata = data.historydata[i];
                                markdata.Author = utf8to16(base64decode(markdata.Author));
                                markdata.AuthorRole = utf8to16(base64decode(markdata.AuthorRole));
                                markdata.Authortype = utf8to16(base64decode(markdata.Authortype));
                                markdata.Content = utf8to16(base64decode(markdata.Content));
                                markdata.Chat = utf8to16(base64decode(markdata.Chat));
                                var connectmsg = "";
                                var slectlist = ""; //哪个控件
                                switch (markdata.Chat) {
                                    case "allchat":
                                        connectmsg = chatulcontent(markdata, 0);
                                        slectlist = ".list1"; //哪个控件
                                        break;
                                    case "sayhim":
                                    case "privatechat":
                                        markdata.Username = utf8to16(base64decode(markdata.Username));
                                        markdata.UserRole = utf8to16(base64decode(markdata.UserRole));
                                        markdata.Usertype = utf8to16(base64decode(markdata.Usertype));
                                        connectmsg = chatulcontent(markdata, 1);
                                        slectlist = ".list1"; //哪个控件
                                        break;
                                    default:
                                        break;
                                }
                                if (markdata.Chat == "privatechat") {
                                    slectlist = ".list2"; //哪个控件
                                }
                                var socketcontent = $(slectlist + " #right-chat #Dialog"); //内容控件
                                socketcontent.append(connectmsg);

                                //用户聊天点击事件
                                contentevent(slectlist, markdata);
                            };
                        }
                        break;
                    }
                case "gagcontrol":
                    {
                        switch (menutext) {
                            case "0":
                                {
                                    //提示操作
                                    var showmsg = username + "用户恢复发言,执行成功!";
                                    layer.msg(showmsg, {
                                        icon: 1
                                    });
                                    break;
                                }
                            case "5":
                                {
                                    //提示操作
                                    var showmsg = username + "用户禁言5分钟,执行成功!";
                                    layer.msg(showmsg, {
                                        icon: 1
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
                        //提示操作
                        var showmsg = username + "用户加入黑名单,执行成功!";
                        layer.msg(showmsg, {
                            icon: 1
                        });
                        break;
                    }
                case "kickout":
                    {
                        var sendmsg = {}; //发送的内容
                        sendmsg["Uname"] = uname;
                        sendmsg["Objname"] = username;
                        sendmsg["Codeid"] = codeid; //公司房间标识符
                        var sendstr = JSON.stringify(eval(sendmsg));
                        sendstr = base64encode(utf16to8(sendstr));
                        socket.emit('all kickout', sendstr);

                        //提示操作
                        var showmsg = username + "用户踢出一小时,执行成功!";
                        layer.msg(showmsg, {
                            icon: 1
                        });
                        break;
                    }
                case "leave":
                    location.href = "/"; //location.href实现客户端页面的跳转
                    break;
                default:
                    break;
            }
        },
        error: function(msg) {
            console.log(msg);
        }
    });
}

//加载左边列表信息
function loaduserlist(obj) {
    var menuindex = $(".MENU span").index($('.MENU .selectlist'));
    var menutext = $(".MENU .selectlist").text();
    var data = obj;
    data.Uname = utf8to16(base64decode(obj.Uname));
    data.RoleName = utf8to16(base64decode(obj.RoleName));
    data.Titlerole = utf8to16(base64decode(obj.Titlerole));
    if (obj.UserIcon != null) {
        data.UserIcon = utf8to16(base64decode(obj.UserIcon));
    }
    switch (menuindex) {
        case 0:
            //menutext="teacher";
            menutext = "member";
            break;
        case 1:
            //menutext="member";
            menutext = "guest";
            break;
        default:
            break;
    }
    if ((menutext == data.RoleName) || ((menutext == "teacher") && (data.RoleName != "member" && data.RoleName != "guest"))) {
        var existindex = 0;
        $("#Border-list .Userbox li").each(function(index) {
            var spantext = $(this).attr("title");
            if (spantext == data.Uname) {
                existindex++;
            }
        });
        if (existindex == 0) {
            var userul = $('#Border-list .Userbox');
            var username = data.Uname;
            var urlIcon = "";
            var contentmsg = "<li title=" + username + "><a href='javascript:void(0);'><div class='UserInformation'>";
            if (data.RoleName == "guest") {
                urlIcon = "/static/images/i/img/role/" + data.UserIcon;
            } else {
                urlIcon = "/upload/usericon/" + data.UserIcon;
            }
            contentmsg += "<img src=" + urlIcon + " id='Userimg' style='width: 100%; height:100%;'>";

            contentmsg += "</div>";
            contentmsg += "<div class='UserName'><span>";
            contentmsg += data.Titlerole;
            contentmsg += "</span><span class='Userbox-right'>";
            if (username.length > 8) {
                username = username.substring(0, 8) + "..";
            }
            contentmsg += username;
            contentmsg += "</span><div class='time'>";
            contentmsg += data.Logintime;
            contentmsg += "</div></div></a></li>";
            userul.append(contentmsg);

            $(".Userbox li:last").click(function(e) {
                if (e.stopPropagation) {
                    e.stopPropagation();
                } else {
                    e.cancelBubble = true;
                }
                //单击
                chatname = data.Uname;
                if (chatname == loginname) {} else {
                    if ((socket_json["AuthorRole"] == data.RoleName) && (data.RoleName == "guest" || data.RoleName == "member")) {} else {
                        $(".popup").show().css({
                            "top": $(this).offset().top + 42,
                            "left": $(this).offset().left + 85
                        });
                    }
                }
            });
        }
    }
}
