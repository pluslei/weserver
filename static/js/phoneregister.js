$(function(){
	
});

$(document).ready(function() {

	$('#register').click(function() {
		if (!CheckReg()){
			return false;
		}

		if ($("#PhoneCode").val().length != 6) {
			$("#PhoneCode").focus();
			$("#userCue").show().html("<font color='red'><b>验证码为六位数</b></font>")
			return false;
		};

		$.post('/register', {UserName: $("#UserName").val(),PassWord:$("#PassWord").val(),PassWord2:$("#Password2").val(),PhoneCode:$("#PhoneCode").val(),Phone:$("#Phone").val(),Regtime:$("#logintime").val(),Xsrf:$('meta[name=_xsrf]').attr('content')}, function(jsondata) {
            if (jsondata.status) {
            		alert("注册成功,请登陆!")
				    window.location.href="/phone/login";
            }else{
            	$("#userCue").show().html(jsondata.info);
            };
        });
	});


	
});

var InterValObj; //timer变量，控制时间
var count = 30; //间隔函数，1秒执行
var curCount = 0;//当前剩余秒数
function sendMessage() {
	if (curCount <= 0) {
		if (!CheckReg()){
			return false;
		}
		$.ajax({
			url: '/path/to/file',
			type: 'default GET (Other values: POST)',
			dataType: 'default: Intelligent Guess (Other values: xml, json, script, or html)',
			data: {param1: 'value1'},
		})
		.done(function() {
			console.log("success");
		})
		.fail(function() {
			console.log("error");
		})
		.always(function() {
			console.log("complete");
		});
		
	　　//向后台发送处理数据
    	$.ajax({
		    type: "POST", //用POST方式传输
		    url: '/sendvalidata', //目标地址
		    data: {Phone: $("#Phone").val(),Logintime:$("#Regtime").val(),Xsrf: $('meta[name=_xsrf]').attr('content')},
		    success: function (jsoninfo){
		    	if (jsoninfo.status) {
		    		curCount = 30;
					//设置button效果，开始计时
				    $("#btnSendCode").attr("disabled", "true");
				    $("#btnSendCode").html(curCount + "秒之后");
				    $(".nowsend").css('background', '#ADABAB');
				    $("#validataShow").removeClass('hide');
				    InterValObj = window.setInterval(SetRemainTime, 1000); //启动计时器，1秒执行一次
		    	}else{
		    		console.info(jsoninfo)
		    		// parent.layer.msg(jsoninfo.info,{shade:[0],icon:1,time:2000,icon:2});
		    	};
		    }
	    });
	};
}

//timer处理函数
function SetRemainTime() {
    if (curCount == 0) {          
        window.clearInterval(InterValObj);//停止计时器
        $("#btnSendCode").removeAttr("disabled");//启用按钮
        $("#btnSendCode").html("重新发送");
        $(".nowsend").css('background', '#2795dc');
    }else {
        curCount--;
        $("#btnSendCode").html(curCount + "秒之后");
    }
}

function CheckReg(){
		if ($('#UserName').val() == "") {
			$('#UserName').focus();
			$('#userCue').show().html("<font color='red'><b>×请输入用户名</b></font>");
			return false;
		}

		if ($('#UserName').val().length < 6 || $('#UserName').val().length > 16) {
			$('#UserName').focus();
			$('#userCue').show().html("<font color='red'><b>×用户名长度为6-16个字符</b></font>");
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
			$('#userCue').show().html("<font color='red'><b>×两次密码不一致</b></font>");
			return false;
		}

		var qphone = /^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\d{8}$/
		if (!qphone.test($("#Phone").val()) || $("#Phone").val().length != 11) {
			$('#Phone').focus();
			$('#userCue').show().html("<font color='red'><b>×请输入正确的手机号码</b></font>");
			return false;
		};
		
		return true
}
