document.ready(function(){

	$('#userreg').click(function() {
		if (!CheckReg()){
			return false;
		}

		$.post('register', {Username:$("#UserName").val(),PassWord:$("#PassWord").val(),PassWord2:$("#PassWord2").val(),Regtime:$("#logintime").val(),Xsrf:$('meta[name=_xsrf]').attr('content')}, function(jsondata) {
            if (jsondata.status) {
            	parent.layer.msg("用户注册成功，请登录", {shade:[0],icon:1}, function(){
				   // window.location.href="registerindex?phone="+$("#Phone").val();
				   $("#qlogin").hide();
				   $("#registerindex").show();
				});
            }else{
            	parent.layer.msg(jsondata.info,{shade:[0],icon:1,time:2000,icon:2});
            	$("#userCue").show().html("");
            };
        });
	});
})

function CheckReg(){
		if ($('#UserName').val() == "") {
			$('#UserName').focus();
			$('#userCue').show().html("<font color='red'><b>×用户名不能为空</b></font>");
			return false;
		}

		if ($('#UserName').val().length < 6 || $('#UserName').val().length > 16) {
			$('#UserName').focus();
			$('#userCue').show().show().html("<font color='red'><b>×用户名位6-16字符</b></font>");
			return false;
		}

		var UserName = /^[a-zA-Z]\w{3,15}$/;
		if (!UserName.test($("#UserName").val())) {
			$('#UserName').focus();
			$('#userCue').show().html("<font color='red'><b>×只能以字母开头，包含字符,数字,下划线</b></font>");
			return false;
		};

		if ($('#PassWord').val().length < 6) {
			$('#PassWord').focus();
			$('#userCue').show().html("<font color='red'><b>×密码不能小于6位</b></font>");
			return false;
		}

		if ($('#Password2').val() != $('#PassWord').val()) {
			$('#Password2').focus();
			$('#userCue').show().html("<font color='red'><b>×两次密码不一致！</b></font>");
			return false;
		}

		return true
}