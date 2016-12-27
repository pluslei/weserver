$(document).ready(function() {
    $('div[class=chat-say]').bind('DOMNodeInserted', function(e) {
        // alert('element now contains: ' + $(e.target).html()); 
        $("div[class=chat-say] img").each(function(index) {
            var imagurl = $(this).attr("src");
            var imgindex = imagurl.indexOf("upload");
            if (imgindex > 0) {
                var upmaxwidth = parseInt($(".chatContent").width() * 0.8);
                var upimgwidth = $(this).width();
                var upimgheigth = $(this).height();
                var timeimg = setInterval(function() {
                    if (upimgwidth == 0) {
                        upimgwidth = $("div[class=chat-say] img").eq(index).width();
                        upimgheigth = $("div[class=chat-say] img").eq(index).height();
                    }
                    if (upimgwidth > 0 && upimgheigth > 0) {
                        clearInterval(timeimg);
                        if (upimgwidth > upmaxwidth) {
                            upimgheigth = parseInt(upimgheigth * upmaxwidth / upimgwidth);
                            upimgwidth = upmaxwidth;
                            $("div[class=chat-say] img").eq(index).attr("width", upimgwidth);
                            $("div[class=chat-say] img").eq(index).attr("height", upimgheigth);
                        }
                        $("div[class=chat-say] img").eq(index).attr("onclick", "_Funshowimg(" + '"' + imagurl + '"' + ")");
                        $("div[class=chat-say] img").eq(index).css("cursor", "pointer");
                    }
                }, 100); // 我这里设置的是100毫秒就扫描一次，可以自己调整
            }
        });
    });
});

function _Funshowimg(imgurl) {
    var upimgcontrol = $('img[name=upimg]');
    upimgcontrol.attr("src", imgurl);
    if (typeof(upimgcontrol.attr("width")) != "undefined") {
        upimgcontrol.removeAttr("width");
    }
    if (typeof(upimgcontrol.attr("height")) != "undefined") {
        upimgcontrol.removeAttr("height");
    }
    // 开定时器检查图片是否加载完
    var timeimg = setInterval(function() {
        var upimgwidth = upimgcontrol.width();
        var upimgheigth = upimgcontrol.height();
        if (upimgwidth > 0 && upimgheigth > 0) {
            clearInterval(timeimg);

            var upimgneww = upimgwidth;
            var upimgnewh = upimgheigth;
            if (upimgwidth > 1024) {
                upimgneww = 1000;
                upimgnewh = parseInt(upimgheigth * upimgneww / upimgwidth);
                if (upimgnewh > 600) {
                    upimgneww = parseInt(upimgneww * 600 / upimgnewh);
                    upimgnewh = 600;
                }
                upimgcontrol.attr("width", upimgneww);
                upimgcontrol.attr("height", upimgnewh);
            }
            upimgneww = (upimgneww + 100).toString() + 'px';
            upimgnewh = (upimgnewh + 100).toString() + 'px';

            layer.open({
                type: 1,
                title: false,
                closeBtn: 1,
                area: [upimgneww, upimgnewh],
                skin: 'layui-layer-rim', // 加上边框
                shadeClose: true,
                content: upimgcontrol
            });
        }
    }, 100); // 我这里设置的是100毫秒就扫描一次，可以自己调整
}
