$(function(){
	$("#refreshimgcode").bind("click",function(e){
		url = 'imagecode?'+Math.random();
		$(this).attr('src',"/phone/" + url);
		return false;
	});
});

$(document).ready(function() {
	$('#phone_login').click(function() {
		if ($('#username').val() == "") {
			$('#username').focus();
			$('#userCue').show().html("<font color='red'><b>×请输入用户名</b></font>");
			return false;
		}
		if ($('#password').val() == "") {
			$('#password').focus();
			$('#userCue').show().html("<font color='red'><b>×请输入密码</b></font>");
			return false;
		}
        $.post('/login', {Username: $("#username").val(),Password:$("#password").val(),Logintime:$("#logintime").val(),ValidataCode:$("#c").val(),Xsrf:$('meta[name=_xsrf]').attr('content')}, function(jsondata) {
            if (jsondata.status) {	
            		alert("登陆成功.")
            		window.location = "/phone?roomid="+roomid;
            }else{
            	url = 'imagecode?'+Math.random();
				$("#refreshimgcode").attr('src',url);
				return false;
                $("#showLoginErr").show().html("");
                $("#userCue").hide();
            };
        });
		   
    });
});