var socket = io.connect(serverurl, {
    'path': '/wechatSocket'
}, {
    'force new connection': true
});
var socketid = ""; //socketid

console.info(serverurl);

//系统设置参数
//Imagesize     int64  //图片大小 单位Kb
//Imagetype     string //图片类型
//Guestchat     int64  //0 禁止聊天 1 允许聊天
//ChatInterval  int64  //0 无间隔  其它数字为间隔时间（秒）
//HistoryMsg    int64  //是否显示历史消息 0显示  1 不显示
//HistoryCount  int64  //显示历史记录条数
//WelcomeMsg    string //欢迎语
var loginname = ""; //用户名
var chatname = ""; //私聊的用户名
var socket_json = {}; //数据内容

var base64EncodeChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
var base64DecodeChars = new Array(　　-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 　　-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 　　-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63, 　　52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1, 　　-1, 　0, 　1, 　2, 　3, 4, 　5, 　6, 　7, 　8, 　9, 10, 11, 12, 13, 14, 　　15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, 　　-1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 　　41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1);

function base64encode(data) {
    var str = unicodejson(0, data);
    var out, i, len;　　
    var c1, c2, c3;
    out = "";
    if (str == null) {
        return out;
    }　　
    len = str.length;　　
    i = 0;　　
    while (i < len) {
        c1 = str.charCodeAt(i++) & 0xff;
        if (i == len) {　　
            out += base64EncodeChars.charAt(c1 >> 2);　　
            out += base64EncodeChars.charAt((c1 & 0x3) << 4);　　
            out += "==";　　
            break;
        }
        c2 = str.charCodeAt(i++);
        if (i == len) {　　
            out += base64EncodeChars.charAt(c1 >> 2);　　
            out += base64EncodeChars.charAt(((c1 & 0x3) << 4) | ((c2 & 0xF0) >> 4));　　
            out += base64EncodeChars.charAt((c2 & 0xF) << 2);　　
            out += "=";　　
            break;
        }
        c3 = str.charCodeAt(i++);
        out += base64EncodeChars.charAt(c1 >> 2);
        out += base64EncodeChars.charAt(((c1 & 0x3) << 4) | ((c2 & 0xF0) >> 4));
        out += base64EncodeChars.charAt(((c2 & 0xF) << 2) | ((c3 & 0xC0) >> 6));
        out += base64EncodeChars.charAt(c3 & 0x3F);　　
    }　　
    return out;
}

function base64decode(str) {
    var c1, c2, c3, c4;　　
    var i, len, out;
    out = "";
    if (str == null) {
        return out;
    }　　
    len = str.length;　　
    i = 0;　　
    while (i < len) {
        /* c1 */
        do {　　
            c1 = base64DecodeChars[str.charCodeAt(i++) & 0xff];
        } while (i < len && c1 == -1);
        if (c1 == -1)　　 break;
        /* c2 */
        do {　　
            c2 = base64DecodeChars[str.charCodeAt(i++) & 0xff];
        } while (i < len && c2 == -1);
        if (c2 == -1)　　 break;
        out += String.fromCharCode((c1 << 2) | ((c2 & 0x30) >> 4));
        /* c3 */
        do {　　
            c3 = str.charCodeAt(i++) & 0xff;　　
            if (c3 == 61)　 return out;　　
            c3 = base64DecodeChars[c3];
        } while (i < len && c3 == -1);
        if (c3 == -1)　　 break;
        out += String.fromCharCode(((c2 & 0XF) << 4) | ((c3 & 0x3C) >> 2));
        /* c4 */
        do {　　
            c4 = str.charCodeAt(i++) & 0xff;　　
            if (c4 == 61)　 return out;　　
            c4 = base64DecodeChars[c4];
        } while (i < len && c4 == -1);
        if (c4 == -1)　　 break;
        out += String.fromCharCode(((c3 & 0x03) << 6) | c4);　　
    }
    return out;
}

function utf16to8(str) {　　
    var out, i, len, c;　　
    out = "";　　
    len = str.length;　　
    for (i = 0; i < len; i++) {
        c = str.charCodeAt(i);
        if ((c >= 0x0001) && (c <= 0x007F)) {　　
            out += str.charAt(i);
        } else if (c > 0x07FF) {　　
            out += String.fromCharCode(0xE0 | ((c >> 12) & 0x0F));　　
            out += String.fromCharCode(0x80 | ((c >> 　6) & 0x3F));　　
            out += String.fromCharCode(0x80 | ((c >> 　0) & 0x3F));
        } else {　　
            out += String.fromCharCode(0xC0 | ((c >> 　6) & 0x1F));　　
            out += String.fromCharCode(0x80 | ((c >> 　0) & 0x3F));
        }　　
    }　　
    return out;
}


function utf8to16(str) {　　
    var out, i, len, c;　　
    var char2, char3;　　
    out = "";　　
    len = str.length;　　
    i = 0;　　
    while (i < len) {
        c = str.charCodeAt(i++);
        switch (c >> 4) {　
            case 0:
            case 1:
            case 2:
            case 3:
            case 4:
            case 5:
            case 6:
            case 7:
                // 0xxxxxxx
                out += str.charAt(i - 1);　　
                break;　
            case 12:
            case 13:
                // 110x xxxx　 10xx xxxx
                char2 = str.charCodeAt(i++);　　
                out += String.fromCharCode(((c & 0x1F) << 6) | (char2 & 0x3F));　　
                break;　
            case 14:
                　　 // 1110 xxxx　10xx xxxx　10xx xxxx
                　　char2 = str.charCodeAt(i++);　　
                char3 = str.charCodeAt(i++);　　
                out += String.fromCharCode(((c & 0x0F) << 12) | 　　　　((char2 & 0x3F) << 6) | 　　　　((char3 & 0x3F) << 0));　　
                break;
        }　　
    }　　
    return out;
}

//对汉子加码解码
function _Unicodetxt(code, str) {
    if (1 == code) {
        return unescape(str.replace(/\\/g, "%"));
    } else {
        return escape(str).replace(/%/g, "\\");
    }
}

//数据解析
function unicodejson(type, data) {
    var sendstr;
    switch (type) {
        case 0:
            sendstr = JSON.stringify(eval(data));
            sendstr = utf16to8(sendstr);
            break;
        case 1:
            sendstr = JSON.stringify(data);
            sendstr = jQuery.parseJSON(sendstr);
            break;
        case 2:
            sendstr = utf8to16(base64decode(data));
            break;
        default:
            break;
    }
    return sendstr;
}

//获取cookie的值
function getnamecookie(sel) {
    var mycookie = "";
    switch (sel) {
        case 0:
            mycookie = $.cookie('weserver_name');
            if (!mycookie) {
                mycookie = _getRandomString(6);
                $.cookie('weserver_name', mycookie, {
                    expires: 7,
                    path: '/',
                    secure: false
                });
            }
            break;
        case 1:
            mycookie = $.cookie('weserver_role');
            if (!mycookie) {
                var irandnum = randomNumminmax(1, 16);
                mycookie = irandnum.toString() + ".jpg";
                $.cookie('weserver_role', mycookie, {
                    expires: 7,
                    path: '/',
                    secure: false
                });
            }
            break;
        default:
            break;
    }
    return mycookie;
}


//number.toFixed(2)//保留小数点
//生成指定位数的随机整数
function randomNum(n) {
    var t = '';
    for (var i = 0; i < n; i++) {
        t += Math.floor(Math.random() * 10);
    }
    return t;
}

//生成指定范围内的随机整数
function randomNumminmax(minNum, maxNum) {
    return parseInt(Math.random() * (maxNum - minNum + 1) + minNum);
}

//生成随机字母和数字
function _getRandomString(len) {
    len = len || 32;
    var $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678'; // 默认去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1  
    var maxPos = $chars.length;
    var pwd = '';
    for (i = 0; i < len; i++) {
        pwd += $chars.charAt(Math.floor(Math.random() * maxPos));
    }
    return pwd;
}

function namesearch(objname, uname) {
    var test = (uname || "").toUpperCase().indexOf(objname.toUpperCase()) >= 0;
    return test;
};

//控制input焦点位置
function setSelection(editor, pos) {
    if (editor.setSelectionRange) {
        editor.focus();
        editor.setSelectionRange(pos, pos);
    } else if (editor.createTextRange) {
        var textRange = editor.createTextRange();
        textRange.collapse(true);
        textRange.moveEnd("character", pos);
        textRange.moveStart("character", pos);
        textRange.select();
    }
}

//控制div焦点位置
function insertHtmlAtCaret(html) {
    var sel, range;
    if (window.getSelection) {
        // IE9 and non-IE
        sel = window.getSelection();
        if (sel.getRangeAt && sel.rangeCount) {
            range = sel.getRangeAt(0);
            range.deleteContents();
            // Range.createContextualFragment() would be useful here but is
            // non-standard and not supported in all browsers (IE9, for one)
            var el = document.createElement("div");
            el.innerHTML = html;
            var frag = document.createDocumentFragment(),
                node, lastNode;
            while ((node = el.firstChild)) {
                lastNode = frag.appendChild(node);
            }
            range.insertNode(frag);
            // Preserve the selection
            if (lastNode) {
                range = range.cloneRange();
                range.setStartAfter(lastNode);
                range.collapse(true);
                sel.removeAllRanges();
                sel.addRange(range);
            }
        }
    } else if (document.selection && document.selection.type != "Control") {
        // IE < 9
        document.selection.createRange().pasteHTML(html);
    }
}

function getCursortPosition(ctrl) {
    var CaretPos = 0; // IE Support
    if (document.selection) {
        ctrl.focus();
        var Sel = document.selection.createRange();
        Sel.moveStart('character', -ctrl.value.length);
        CaretPos = Sel.text.length;
    }
    // Firefox support
    else if (ctrl.selectionStart || ctrl.selectionStart == '0')
        CaretPos = ctrl.selectionStart;
    return (CaretPos);
}

function setCaretPosition(ctrl, pos) {
    if (ctrl.setSelectionRange) {
        ctrl.focus();
        ctrl.setSelectionRange(pos, pos);
    } else if (ctrl.createTextRange) {
        var range = ctrl.createTextRange();
        range.collapse(true);
        range.moveEnd('character', pos);
        range.moveStart('character', pos);
        range.select();
    }
}

//控制input焦点位置
function setScrollstyle(Scrollname, color, type) {
    if (type) {
        $(Scrollname).getNiceScroll().remove();
    }
    $(Scrollname).niceScroll({
        touchbehavior: false,
        railoffset: true, // 可以使用top/left来修正位置
        cursorcolor: color,
        cursoropacitymax: 1.6,
        cursorwidth: "8px",
        //horizrailenabled:false,  
        cursorborder: "0",
        cursorborderradius: "5px",
        autohidemode: false
    });
}

//隐藏显示滚动条
function showScrollstyle(Scrollname, type) {
    switch (type) {
        case 0:
            $(Scrollname).getNiceScroll().hide();
            break;
        case 1:
            $(Scrollname).getNiceScroll().show();
            break;
        default:
            break;
    }
}

//控制滚动条
function funControlscrollbar(Scrollname, issettop) {
    if (issettop) {
        if ($(Scrollname)[0].scrollHeight > $(Scrollname).height()) {
            $(Scrollname).scrollTop($(Scrollname)[0].scrollHeight - $(Scrollname).height());
        }
    } else {
        $(Scrollname).scrollTop(0);
    }
    $(Scrollname).getNiceScroll().resize();
}

//判断图片是否存在
function checkimgexists(imgurl) {
    var ImgObj = new Image(); //判断图片是否存在  
    ImgObj.src = imgurl;
    //没有图片，则返回-1  
    if (ImgObj.fileSize > 0 || (ImgObj.width > 0 && ImgObj.height > 0)) {
        return true;
    } else {
        return false;
    }
}

//上传文件
function img_Fileupload() {
    var FileSize = chat_sysconfig.Imagesize.toString() + "KB"; //图片大小
    var FileStyle = chat_sysconfig.Imagetype; //图片类型
    var RandNumber = 0;
    $('#chatsendImg').uploadify({
        'swf': '/static/img/uploadify.swf', // uploadify.swf 文件的相对JS文件的路径
        'uploader': '/chat/upload', //后台处理程序的相对路径
        'fileSizeLimit': FileSize, //'2MB'
        'floder': '/upload/img',
        'queueID': 'fileQueue',
        'queueSizeLimit': 999999999,
        'uploadLimit': 999999999,
        //'progressData': 'speed',
        'auto': true,
        'multi': false,
        'buttonClass': 'some-class', //自定义的样式类名
        //'queueID'  : 'some_file_queue',
        //'buttonCursor' : 'arrow',
        //'buttonText': '上传图片',
        'buttonTitle': '提示文字',
        'buttonImage': '/static/images/i/img/img.png',
        'requeueErrors': true,
        'debug': false,
        'height': 30, // The height of the flash button
        'width': 64, // The width of the flash button
        'itemTemplate': false,
        'successTimeout': 30, //设置文件上传后等待服务器响应的秒数，超出这个时间，将会被认为上传成功，默认为30秒
        'fileTypeExts': '*.bmp;*.jpg;.jpeg;*.png;*.gif', //FileStyle, //'*.bmp;*.jpg;.jpeg;*.png;*.gif'//设置允许上传的文件扩展名（也就是文件类型）。但手动键入文件名可以绕过这种级别的安全检查，所以你应该始终在服务端中检查文件类型。输入多个扩展名时用分号隔开('*.*;*.jpg;*.png;*.gif')
        'fileTypeDesc': '文件格式', //可选文件的描述。这个值出现在文件浏览窗口中的文件类型下拉选项中。（但我设置了好像没效果？原文：The description of the selectable files.  This string appears in the browse files dialog box in the file type drop down.）
        'fileObjName': 'Filedata', //设置后台的文件名
        'method': 'post',
        'formData': {
            'myname': '',
            'ucodeid': '',
            'randnum': ''
        },
        'onUploadStart': function(file) {
            RandNumber = randomNum(10);
            $('#chatsendImg').uploadify("settings", "formData", {
                'myname': uname,
                'ucodeid': codeid,
                'randnum': RandNumber.toString()
            });
        },
        // 更多的参数
        'onUploadProgress': function(file, bytesUploaded, bytesTotal, totalBytesUploaded, totalBytesTotal) {},
        'onSelect': function(file) {},
        'onCancel': function(file) {},
        'onUploadComplete': function(file) {
            var upimgcontrol = $('img[name=upimg]');
            var filename = file.name;
            if (filename.length > 3) {
                var indexnum = filename.indexOf('.');
                if (indexnum != -1) {
                    filename = RandNumber.toString() + filename.substring(indexnum, filename.length);
                }
                var imagurl = "/upload/img/" + filename;
                upimgcontrol.attr("src", imagurl);
                //开定时器检查图片是否加载完
                var timeimg = setInterval(function() {
                    var upimgwidth = upimgcontrol.width();
                    var upimgheigth = upimgcontrol.height();
                    if (upimgwidth > 0 && upimgheigth > 0) {
                        clearInterval(timeimg);
                        var inputcontrol = $('.chatInput .chat-say'); //输入框
                        var upmaxwidth = parseInt($(".chatContent").width() * 0.92);
                        var inputcontent = "<img src=" + imagurl + " >";
                        if (upimgwidth > upmaxwidth) {
                            upimgheigth = parseInt(upimgheigth * upmaxwidth / upimgwidth);
                            upimgwidth = upmaxwidth;
                            inputcontent = "<img src=" + imagurl + " width=" + upimgwidth + " height=" + upimgheigth + " ></img>";
                        }
                        //if (CheckImgExists(imagurl)) {
                        inputcontent = inputcontrol.html() + inputcontent;
                        inputcontrol.html(inputcontent);
                        //}
                    }
                }, 100); // 我这里设置的是100毫秒就扫描一次，可以自己调整
            }
        },
        // Triggered when a file is not added to the queue
        'onSelectError': function(file, errorCode, errorMsg) {
            // Load the swfupload settings
            var settings = this.settings;
            // Run the default event handler
            if ($.inArray('onSelectError', settings.overrideEvents) < 0) {
                switch (errorCode) {
                    case SWFUpload.QUEUE_ERROR.QUEUE_LIMIT_EXCEEDED:
                        if (settings.queueSizeLimit > errorMsg) {
                            this.queueData.errorMsg += '\nThe number of files selected exceeds the remaining upload limit (' + errorMsg + ').';
                        } else {
                            this.queueData.errorMsg += '\nThe number of files selected exceeds the queue size limit (' + settings.queueSizeLimit + ').';
                        }
                        break;
                    case SWFUpload.QUEUE_ERROR.FILE_EXCEEDS_SIZE_LIMIT:
                        //this.queueData.errorMsg += '\nThe file "' + file.name + '" exceeds the size limit (' + settings.fileSizeLimit + ').';
                        this.queueData.errorMsg = file.name + "文件大小超过(" + settings.fileSizeLimit + ").";
                        //alert("10000");
                        break;
                    case SWFUpload.QUEUE_ERROR.ZERO_BYTE_FILE:
                        this.queueData.errorMsg += '\nThe file "' + file.name + '" is empty.';
                        break;
                    case SWFUpload.QUEUE_ERROR.FILE_EXCEEDS_SIZE_LIMIT:
                        this.queueData.errorMsg += '\nThe file "' + file.name + '" is not an accepted file type (' + settings.fileTypeDesc + ').';
                        break;
                }
            }
            if (errorCode != SWFUpload.QUEUE_ERROR.QUEUE_LIMIT_EXCEEDED) {
                delete this.queueData.files[file.id];
            }

            // Call the user-defined event handler
            if (settings.onSelectError) settings.onSelectError.apply(this, arguments);
        },
        'onUploadSuccess': function(file, data, response) {},
        'onUploadError': function(file, errorCode, erorMsg, errorString) {}, //一个文件完成上传但返回错误时触发，有以下参数
        'onQueueComplete': function(queueData) { //队列中的文件都上传完后触发，返回queueDate参数，有以下属性：
        },
        'onUploadError': function(file, errorCode, erorMsg, errorString) {}
    });
}

function chatuiulcontent(chatdata, type) {
    var i, j = 0
    var length = chatdata.length;
    if (type == 1) {
        j = 0;
        $('.chatContent ul li').remove();
    } else {
        j = chatdata.length - 1;
    }
    for (var i = j; i < length; i++) {
        var data = chatdata[i];
        var connectmsg = "";
        if (data.Author.length <= 0) {
            return;
        }
        switch (data.Chat) {
            case "allchat":
                {
                    switch (data.AuditStatus) {
                        case 2:
                            if (data.IsEmitBroad) {
                                connectmsg = "<li class='distance chatPos'>";
                                connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                                connectmsg += data.Time;
                                connectmsg += "</span>";
                                connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                                connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                                connectmsg += data.Author;
                                connectmsg += "</a>";
                                connectmsg += "<div>";
                                connectmsg += "<span class='chatword am-margin-right'>";
                                connectmsg += data.Content;
                                connectmsg += '</span>';
                                connectmsg += '</div>';
                                connectmsg += '</li>';
                            } else {
                                if (socket_json['RoleName'] == "manager") {
                                    connectmsg = "<li class='distance chatPos'>";
                                    connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                                    connectmsg += data.Time;
                                    connectmsg += "</span>";
                                    connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                                    connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                                    connectmsg += data.Author;
                                    connectmsg += "</a>";
                                    connectmsg += "<button type='button' class='am-btn am-radius passed'>审 核 通 过</button>";
                                    connectmsg += "<div>";
                                    connectmsg += "<span class='chatword am-margin-right'>";
                                    connectmsg += data.Content;
                                    connectmsg += '</span>';
                                    connectmsg += '</div>';
                                    connectmsg += '</li>';
                                }
                            }
                            break;
                        case 100:
                            if (!data.IsEmitBroad && data.Objname != loginname) {
                                var checkcontent = true;
                                if (socket_json['RoleName'] == "manager") {
                                    var divcontent = $('<div></div>');
                                    var receivedata = data.Authorcss + data.Author;
                                    var receivecontent = data.Content;
                                    var passedlen = $(".chatPos .passed").length;
                                    var reg = new RegExp("&quot;", "g"); //创建正则RegExp对象
                                    receivecontent = receivecontent.replace(reg, '"');
                                    divcontent.html(receivecontent);
                                    divcontent.find('img').each(function(index) {
                                        receivedata += $(this).attr("src");
                                    });
                                    receivedata += divcontent.text();
                                    $(".chatPos .passed").each(function(index) {
                                        var resultdata = "";
                                        var datacontent = "";
                                        //resultdata += $(this).siblings('.getTime').text();
                                        resultdata += $(this).siblings('img').attr('src');
                                        resultdata += $(this).siblings('a').text();
                                        datacontent = $(this).siblings('div').children('.chatword').html();
                                        datacontent = datacontent.replace(reg, '"');
                                        divcontent.html(datacontent);
                                        divcontent.find('img').each(function(index) {
                                            resultdata += $(this).attr("src");
                                        });
                                        resultdata += divcontent.text();
                                        if (receivedata == resultdata) {
                                            $(".chatPos .passed").eq(index).after("<span class='am-margin-left-xs Checked am-padding-horizontal-xs am-radius'>已审核</span>");
                                            $(".chatPos .passed").eq(index).remove();
                                            checkcontent = false;
                                        }
                                    });
                                }
                                if (checkcontent && data.Author != loginname) {
                                    connectmsg = "<li class='distance chatPos'>";
                                    connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                                    connectmsg += data.Time;
                                    connectmsg += "</span>";
                                    connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                                    connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                                    connectmsg += data.Author;
                                    connectmsg += "</a>";
                                    connectmsg += "<div>";
                                    connectmsg += "<span class='chatword am-margin-right'>";
                                    connectmsg += data.Content;
                                    connectmsg += '</span>';
                                    connectmsg += '</div>';
                                    connectmsg += '</li>';
                                }
                            }
                            break;
                        default:
                            connectmsg = "<li class='distance chatPos'>";
                            connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                            connectmsg += data.Time;
                            connectmsg += "</span>";
                            connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                            connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                            connectmsg += data.Author;
                            connectmsg += "</a>";
                            connectmsg += "<div>";
                            switch (data.AuthorRole) {
                                case "manager":
                                    connectmsg += "<span class='chatword am-margin-right managerchat' style='background:#ff1b46;color:#fff;'>";
                                    break;
                                default:
                                    connectmsg += "<span class='chatword am-margin-right'>";
                                    break;
                            }
                            connectmsg += data.Content;
                            connectmsg += '</span>';
                            connectmsg += '</div>';
                            connectmsg += '</li>';
                            break;
                    }
                    break;
                }
            case "sayhim":
                {
                    switch (data.AuditStatus) {
                        case 2:
                            if (data.IsEmitBroad) {
                                connectmsg = "<li class='distance chattosay'>";
                                connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                                connectmsg += data.Time;
                                connectmsg += "</span>";
                                connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                                connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                                connectmsg += data.Author;
                                connectmsg += "</a>";
                                connectmsg += "<span class='am-padding-left-xs'>对</span>";
                                connectmsg += "<img src=" + data.Usercss + " alt='' width='48px' height='20px' />";
                                connectmsg += "<a class='am-badge am-badge-danger am-radius am-text-default am-margin-vertical-xs' name='chatuser'>";
                                connectmsg += data.Username;
                                connectmsg += "</a>";
                                connectmsg += "<div>";
                                connectmsg += "<span class='chatword am-margin-right'>";
                                connectmsg += data.Content;
                                connectmsg += '</span>';
                                connectmsg += '</div>';
                                connectmsg += '</li>';
                            } else {
                                if (socket_json['RoleName'] == "manager") {
                                    connectmsg = "<li class='distance chattosay'>";
                                    connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                                    connectmsg += data.Time;
                                    connectmsg += "</span>";
                                    connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                                    connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                                    connectmsg += data.Author;
                                    connectmsg += "</a>";
                                    connectmsg += "<span class='am-padding-left-xs'>对</span>";
                                    connectmsg += "<img src=" + data.Usercss + " alt='' width='48px' height='20px' />";
                                    connectmsg += "<a class='am-badge am-badge-danger am-radius am-text-default am-margin-vertical-xs' name='chatuser'>";
                                    connectmsg += data.Username;
                                    connectmsg += "</a>";
                                    connectmsg += "<button type='button' class='am-btn am-radius passed'>审 核 通 过</button>";
                                    connectmsg += "<div>";
                                    connectmsg += "<span class='chatword am-margin-right'>";
                                    connectmsg += data.Content;
                                    connectmsg += '</span>';
                                    connectmsg += '</div>';
                                    connectmsg += '</li>';
                                }
                            }
                            break;
                        case 100:
                            if (!data.IsEmitBroad && data.Objname != loginname) {
                                var checkcontent = true;
                                if (socket_json['RoleName'] == "manager") {
                                    var divcontent = $('<div></div>');
                                    var receivedata = data.Authorcss + data.Author;
                                    var receivecontent = data.Content;
                                    var passedlen = $(".chattosay .passed").length;
                                    var reg = new RegExp("&quot;", "g"); //创建正则RegExp对象
                                    receivecontent = receivecontent.replace(reg, '"');
                                    divcontent.html(receivecontent);
                                    divcontent.find('img').each(function(index) {
                                        receivedata += $(this).attr("src");
                                    });
                                    receivedata += divcontent.text();
                                    $(".chattosay .passed").each(function(index) {
                                        var resultdata = "";
                                        var datacontent = "";
                                        //resultdata += $(this).siblings('.getTime').text();
                                        resultdata += $(this).siblings('img').attr('src');
                                        resultdata += $(this).siblings('a').text();
                                        datacontent = $(this).siblings('div').children('.chatword').html();
                                        datacontent = datacontent.replace(reg, '"');
                                        divcontent.html(datacontent);
                                        divcontent.find('img').each(function(index) {
                                            resultdata += $(this).attr("src");
                                        });
                                        resultdata += divcontent.text();
                                        if (receivedata == resultdata) {
                                            $(".chattosay .passed").eq(index).after("<span class='am-margin-left-xs Checked am-padding-horizontal-xs am-radius'>已审核</span>");
                                            $(".chattosay .passed").eq(index).remove();
                                            checkcontent = false;
                                        }
                                    });
                                }
                                if (checkcontent && data.Author != loginname) {
                                    connectmsg = "<li class='distance chattosay'>";
                                    connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                                    connectmsg += data.Time;
                                    connectmsg += "</span>";
                                    connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                                    connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                                    connectmsg += data.Author;
                                    connectmsg += "</a>";
                                    connectmsg += "<span class='am-padding-left-xs'>对</span>";
                                    connectmsg += "<img src=" + data.Usercss + " alt='' width='48px' height='20px' />";
                                    connectmsg += "<a class='am-badge am-badge-danger am-radius am-text-default am-margin-vertical-xs' name='chatuser'>";
                                    connectmsg += data.Username;
                                    connectmsg += "</a>";
                                    connectmsg += "<div>";
                                    connectmsg += "<span class='chatword am-margin-right'>";
                                    connectmsg += data.Content;
                                    connectmsg += '</span>';
                                    connectmsg += '</div>';
                                    connectmsg += '</li>';
                                }
                            }
                            break;
                        default:
                            connectmsg = "<li class='distance chattosay'>";
                            connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                            connectmsg += data.Time;
                            connectmsg += "</span>";
                            connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                            connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                            connectmsg += data.Author;
                            connectmsg += "</a>";
                            connectmsg += "<span class='am-padding-left-xs'>对</span>";
                            connectmsg += "<img src=" + data.Usercss + " alt='' width='48px' height='20px' />";
                            connectmsg += "<a class='am-badge am-badge-danger am-radius am-text-default am-margin-vertical-xs' name='chatuser'>";
                            connectmsg += data.Username;
                            connectmsg += "</a>";
                            connectmsg += "<div>";
                            switch (data.AuthorRole) {
                                case "manager":
                                    connectmsg += "<span class='chatword am-margin-right managerchat' style='background:#ff1b46;color:#fff;'>";
                                    break;
                                default:
                                    connectmsg += "<span class='chatword am-margin-right'>";
                                    break;
                            }
                            connectmsg += data.Content;
                            connectmsg += '</span>';
                            connectmsg += '</div>';
                            connectmsg += '</li>';
                            break;
                    }
                    break;
                }
            case "privatechat":
                {
                    connectmsg = "<li class='distance'>";
                    connectmsg += "<span class='am-fl am-block getTime am-text-center am-margin-top-xs'>";
                    connectmsg += data.Time;
                    connectmsg += "</span>";
                    connectmsg += "<img src=" + data.Authorcss + " alt='' width='48px' height='20px' />";
                    connectmsg += "<a class='am-badge am-badge-success am-radius am-text-default am-margin-vertical-xs' name='chatauthor'>";
                    connectmsg += data.Author;
                    connectmsg += "</a>";
                    connectmsg += "<span class='am-padding-left-xs'>对</span>";
                    connectmsg += "<img src=" + data.Usercss + " alt='' width='48px' height='20px' />";
                    connectmsg += "<a class='am-badge am-badge-danger am-radius am-text-default am-margin-vertical-xs' name='chatuser'>";
                    connectmsg += data.Username;
                    connectmsg += "</a>";
                    connectmsg += "<div>";
                    switch (data.AuthorRole) {
                        case "manager":
                            connectmsg += "<span class='chatword am-margin-right managerchat' style='background:#ff1b46;color:#fff;'>";
                            break;
                        default:
                            connectmsg += "<span class='chatword am-margin-right'>";
                            break;
                    }
                    connectmsg += data.Content;
                    connectmsg += '</span>';
                    connectmsg += '</div>';
                    connectmsg += '</li>';
                    break;
                }
            default:
                break;
        }
        $('.chatContent > ul').append(connectmsg);
        contentevent(data);
    }
}

function contentevent(data) {
    $(".chatContent a[name=chatauthor]:last").click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }
        //单击
        chatname = $(this).text();
        if (chatname != loginname) {
            for (var i = 0; i < 4; i++) {
                $(".chatList ul li").eq(i).hide();
            }
            $.ajax({
                url: '/chat/permisse',
                type: 'post',
                data: {
                    rolename: socket_json['RoleName']
                },
                success: function(data) {
                    if (data.chatname != null) {
                        var length = data.chatname.length;
                        for (var i = 0; i < length; i++) {
                            if (data.chatname[i] == "out1hour") { //踢出1小时
                                $(".chatList ul li").eq(0).show();
                            }
                            if (data.chatname[i] == "disablemsg") { //禁言5分钟
                                $(".chatList ul li").eq(1).show();
                            }
                            if (data.chatname[i] == "enablemsg") {
                                $(".chatList ul li").eq(2).show();
                            }
                            if (data.chatname[i] == "addblacklist") {
                                $(".chatList ul li").eq(3).show();
                            }
                        }
                    }
                }
            });
            if (socket_json['InSider'] == 0) {
                if (data.AuthorInSider == 0) {
                    $(".chatList ul li").eq(5).hide();
                }
            }
            $(".chatList").show().css({
                "top": $(this).position().top + 102,
                "left": $(this).position().left
            });
        }
    });

    $(".chatContent a[name=chatuser]:last").click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }
        //单击
        chatname = $(this).text();
        if (chatname != loginname) {
            for (var i = 0; i < 4; i++) {
                $(".chatList ul li").eq(i).hide();
            }
            $.ajax({
                url: '/chat/permisse',
                type: 'post',
                data: {
                    rolename: socket_json['RoleName']
                },
                success: function(data) {
                    if (data.chatname != null) {
                        var length = data.chatname.length;
                        for (var i = 0; i < length; i++) {
                            if (data.chatname[i] == "out1hour") { //踢出1小时
                                $(".chatList ul li").eq(0).show();
                            }
                            if (data.chatname[i] == "disablemsg") { //禁言5分钟
                                $(".chatList ul li").eq(1).show();
                            }
                            if (data.chatname[i] == "enablemsg") {
                                $(".chatList ul li").eq(2).show();
                            }
                            if (data.chatname[i] == "addblacklist") {
                                $(".chatList ul li").eq(3).show();
                            }
                        }
                    }
                }
            });
            if (socket_json['InSider'] == 0) {
                if (data.AuthorInSider == 0) {
                    $(".chatList ul li").eq(5).hide();
                }
            }
            $(".chatList").show().css({
                "top": $(this).position().top + 102,
                "left": $(this).position().left
            });
        }
    });

    $(".chatPos .passed:last").click(function(e) {
        var sendmsg = {}; //发送的内容
        var uname = $(this).parent().find('a[name=chatauthor]').text();
        $(this).attr('disabled', 'true');
        $(this).after("<span class='am-margin-left-xs Checked am-padding-horizontal-xs am-radius'>已审核</span>");
        $(this).remove();
        sendmsg["Objname"] = loginname;
        sendmsg["Author"] = uname;
        sendmsg["Codeid"] = codeid; //公司房间标识符
        sendmsg["Chat"] = "allchat";
        sendmsg["AuthorRole"] = data.AuthorRole;
        sendmsg["Content"] = data.Content;
        sendmsg["AuditStatus"] = 100;
        var sendstr = base64encode(sendmsg);
        socket.emit('all message', sendstr);
    });

    $(".chattosay .passed:last").click(function(e) {
        var sendmsg = {}; //发送的内容
        var objname = $(this).parent().find('a[name=chatauthor]').text();
        var username = $(this).parent().find('a[name=chatuser]').text();
        $(this).attr('disabled', 'true');
        $(this).after("<span class='am-margin-left-xs Checked am-padding-horizontal-xs am-radius'>已审核</span>");
        $(this).remove();
        sendmsg["Objname"] = loginname;
        sendmsg["Author"] = objname;
        sendmsg["Username"] = username;
        sendmsg["Codeid"] = codeid; //公司房间标识符
        sendmsg["Chat"] = "sayhim";
        sendmsg["AuthorRole"] = data.AuthorRole;
        sendmsg["Content"] = data.Content;
        sendmsg["AuditStatus"] = 100;
        var sendstr = base64encode(sendmsg);
        socket.emit('all message', sendstr);
    });
}

function strgetWidth(fontSize) {
    var span = document.getElementById("__getwidth");
    if (span == null) {
        span = document.createElement("span");
        span.id = "__getwidth";
        document.body.appendChild(span);
        span.style.visibility = "hidden";
        span.style.whiteSpace = "nowrap";
    }
    span.innerText = this;
    span.style.fontSize = fontSize + "px";
    return span.offsetWidth;
}

$(document).ready(function() {
    $('.chatInput .chatTool-expression').click(function(e) {
        if (e.stopPropagation) {
            e.stopPropagation();
        } else {
            e.cancelBubble = true;
        }
        $('.chatInput .look').show();
        $('.look .looktab a').eq(0).removeClass('no-active').addClass('active');
        $('.look .lookContent').css('display', '');
        var looktablen = $('.look .looktab a').length;
        for (var i = 1; i < looktablen; i++) {
            $('.look .looktab a').eq(i).removeClass('active').addClass('no-active');
        }
    });

    $(document).bind('click', function() {
        $('.chatList').css('display', 'none');
        $('.chatInput .look').css('display', 'none');
        $('.speak-content .trueUser-tosay').css('display', 'none');
    });
});
