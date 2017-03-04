$(function(){
	$('#switch_qlogin').click(function(){
		$('#switch_login').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_qlogin').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_bottom').animate({left:'0px',width:'70px'});
		$('#qlogin').css('display','none');
		$('#web_qr_login').css('display','block');
	});
	$('#switch_login').click(function(){
		$('#switch_login').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_qlogin').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_bottom').animate({left:'154px',width:'70px'});
		$('#qlogin').css('display','block');
		$('#web_qr_login').css('display','none');
	});

	$("#refreshimgcode").bind("click",function(e){
		url = 'imagecode?'+Math.random();
		$(this).attr('src',url);
		return false;
	});
});

$(document).ready(function() {
	$('#reg').click(function() {
		if (!CheckRegPhone()){
			return false;
		}

		if ($("#PhoneCode").val().length != 6) {
			$("#PhoneCode").focus();
			$("#userCue").show().html("<font color='red'><b>验证码为六位数</b></font>")
			return false;
		};
		// if $('#agreement').prop('checked') == false{
		// 	alert(111)
		// }
		if ($('#agreement').prop('checked')){
		}else{
			$("#userCue").show().html("<font color='red'><b>请同意直播服务条款</b></font>")
			return false;
		}


		$.post('register', {Phone:$("#Phone").val(),PhoneCode:$("#PhoneCode").val(),Regtime:$("#logintime").val(),Xsrf:$('meta[name=_xsrf]').attr('content')}, function(jsondata) {
            if (jsondata.status) {
            	parent.layer.msg("手机注册成功,请设置帐户！", {shade:[0],icon:1}, function(){
				   // window.location.href="registerindex?phone="+$("#Phone").val();
				  $("#phonediv").hide();
				  $("#userCue").show().html("");
				  $("#userdiv").show(); 

				  
				});
            }else{
            	parent.layer.msg(jsondata.info,{shade:[0],icon:1,time:2000,icon:2});
            	$("#userCue").show().html("");
            };
        });
	});

	$('#reg1').click(function() {
		if (!CheckReg()){
			return false;
		}

		$.post('registeruser', {UserName:$("#UserName").val(),PassWord:$("#PassWord").val(),PassWord2:$("#Password2").val(),Phone:$("#Phone").val(),Regtime:$("#logintime").val(),Xsrf:$('meta[name=_xsrf]').attr('content')}, function(jsondata) {
            if (jsondata.status) {
            	parent.layer.msg("注册成功,请登陆！", {shade:[0],icon:1}, function(){
				  window.location.href="reglogin?type=1"
				});
            }else{
            	parent.layer.msg(jsondata.info,{shade:[0],icon:1,time:2000,icon:2});
            	$("#userCue").show().html("");
            };
        });
	});



	

	$('#userlogin').click(function() {
		if ($("#u").val().length <= 0 || $("#p").val().length <= 0 || $("#c").val().length <= 0) {
			$("#showLoginErr").show().html("请将表单填写完整");
			return false;
		};
        $.post('login', {Username: $("#u").val(),Password:$("#p").val(),Logintime:$("#Regtime").val(),ValidataCode:$("#c").val(),Xsrf:$('meta[name=_xsrf]').attr('content')}, function(jsondata) {
            if (jsondata.status) {
            	parent.layer.msg("用户登录成功", {shade:[0],icon:1,time:1000}, function(){
            		// userLoigSuccess(jsondata.info);
				    //parent.location.reload();
				    parent.location="/";
				})
            }else{
            	parent.layer.msg(jsondata.info,{shade:[0],icon:1,time:2000,icon:2});
            	url = 'imagecode?'+Math.random();
				$("#refreshimgcode").attr('src',url);
				return false;
                $("#showLoginErr").show().html("");
            };
        });
    });    
});

var InterValObj; //timer变量，控制时间
var count = 30; //间隔函数，1秒执行
var curCount = 0;//当前剩余秒数

function sendMessage() {
	if (curCount <= 0) {
		if (!CheckRegPhone()){
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
		    url: 'sendvalidata', //目标地址
		    data: {Phone: $("#Phone").val(),Logintime:$("#Regtime").val(),Xsrf: $('meta[name=_xsrf]').attr('content')},
		    success: function (jsoninfo){
		    	if (jsoninfo.status) {
		    		curCount = count;
					//设置button效果，开始计时
				    $("#btnSendCode").attr("disabled", "true");
				    $("#btnSendCode").html(curCount + "秒之后");
				    $(".nowsend").css('background', '#ADABAB');
				    $("#validataShow").removeClass('hide');
				    InterValObj = window.setInterval(SetRemainTime, 1000); //启动计时器，1秒执行一次
		    	}else{
		    		parent.layer.msg(jsoninfo.info,{shade:[0],icon:1,time:2000,icon:2});
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

function CheckRegPhone(){
		// if ($('#UserName').val() == "") {
		// 	$('#UserName').focus();
		// 	$('#userCue').show().html("<font color='red'><b>×用户名不能为空</b></font>");
		// 	return false;
		// }

		// if ($('#UserName').val().length < 6 || $('#UserName').val().length > 16) {
		// 	$('#UserName').focus();
		// 	$('#userCue').show().show().html("<font color='red'><b>×用户名位6-16字符</b></font>");
		// 	return false;
		// }

		// var UserName = /^[a-zA-Z]\w{3,15}$/;
		// if (!UserName.test($("#UserName").val())) {
		// 	$('#UserName').focus();
		// 	$('#userCue').show().html("<font color='red'><b>×只能以字母开头，包含字符,数字,下划线</b></font>");
		// 	return false;
		// };

		// if ($('#PassWord').val().length < 6) {
		// 	$('#PassWord').focus();
		// 	$('#userCue').show().html("<font color='red'><b>×密码不能小于6位</b></font>");
		// 	return false;
		// }

		// if ($('#Password2').val() != $('#PassWord').val()) {
		// 	$('#Password2').focus();
		// 	$('#userCue').show().html("<font color='red'><b>×两次密码不一致！</b></font>");
		// 	return false;
		// }

		var qphone = /^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\d{8}$/
		if (!qphone.test($("#Phone").val()) || $("#Phone").val().length != 11) {
			$('#Phone').focus();
			$('#userCue').show().html("<font color='red'><b>×请输入正确的手机号码</b></font>");
			return false;
		};


		
		return true
}


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