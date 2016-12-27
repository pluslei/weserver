var UserData; //数据验证
var InterValObj; //timer变量，控制时间
var count = 120; //间隔函数，1秒执行
var curCount = 0; //当前剩余秒数

layer.use('extend/layer.ext.js')
$(function() {
	//输入用户名事件
	$('.next a').click(function() {
		var username = $('input[name=account]').val();
		var errorcontrol = $('.controlmain .account-content span');
		if (username.length > 0) {
			$.ajax({
				url: '/index/checkedname',
				type: 'POST',
				data: {
					uname: username
				},
				success: function(data) {
					if (data.msg != null) {
						UserData = data.msg;
						if (UserData.Exist) {
							if (UserData.ShowMsg.length > 0) {
								errorcontrol.text(UserData.ShowMsg);
							} else {
								$('.contentfillin').css("display", "none")
								$('.contentreset').css("display", "none")
								$('.contentcomplete').css("display", "none")

								$('.contentverifi').css("display", "")
								$('.fillinactive').css("background-image", "url('/static/img/step/1-2.png')")
								$('.verifiactive').css("background-image", "url('/static/img/step/2-1.png')")
								$('.resetactive').css("background-image", "url('/static/img/step/2-2.png')")
								$('.completeactive').css("background-image", "url('/static/img/step/3-2.png')")

								$('.message-account .message-tel .tel-num').text(UserData.Phone);
							}
						} else {
							errorcontrol.text("账号不存在");
						}
					}
				}
			});
		} else {
			errorcontrol.text("账号不能为空");
		}
	});

	//验证事件
	$('.resetpwd a').click(function() {
		var codecontrol = $('input[name=prove-code]').val();
		var errorcontrol = $('.prove-code-content .prove-code-info');
		var provecode = $(".message-tel .send-prove-code").html();
		if (codecontrol.length > 0) {
			var code = parseInt(codecontrol);
			if (code > 0) {
				if (codecontrol.length == 6) {
					if (UserData != null) {
						$.ajax({
							url: '/index/verification/code',
							type: 'POST',
							data: {
								uname: UserData.Uname,
								uphone: UserData.AecPhone,
								ucode: code
							},
							success: function(data) {
								if (data.msg != null) {
									UserData = data.msg;
									if (UserData.ShowMsg.length > 0) {
										errorcontrol.text(UserData.ShowMsg);
									} else {
										$('.contentfillin').css("display", "none")
										$('.contentverifi').css("display", "none")
										$('.contentcomplete').css("display", "none")

										$('.contentreset').css("display", "")
										$('.fillinactive').css("background-image", "url('/static/img/step/1-2.png')")
										$('.verifiactive').css("background-image", "url('/static/img/step/2-2.png')")
										$('.resetactive').css("background-image", "url('/static/img/step/2-1.png')")
										$('.completeactive').css("background-image", "url('/static/img/step/3-2.png')")
									}
								}
							}
						});
					} else {
						errorcontrol.text("请发送验证码");
					}
				} else {
					errorcontrol.text("验证码为6位数字");
				}
			} else {
				errorcontrol.text("验证码为6位数字");
			}
		} else {
			errorcontrol.text("验证码不能为空");
		}
	});

	//修改密码事件
	$('.confirm-reset a').click(function() {
		if (UserData != null) {
			var newpwd = $("input[name=newpwd]").val();
			var confirmpwd = $("input[name=confirmpwd]").val();
			var newerror = $(".new-pwd .input-error-text");
			var confirmerror = $(".confirm-pwd .input-error-text");
			if (newpwd.length > 0) {
				if (newpwd.length >= 6) {
					if (confirmpwd.length > 0) {
						if (newpwd == confirmpwd) {
							$.ajax({
								url: '/index/updata/psw',
								type: 'POST',
								data: {
									uname: UserData.Uname,
									uphone: UserData.AecPhone,
									ucode: UserData.PhoneCode,
									newpsw: newpwd
								},
								success: function(data) {
									if (data.msg != null) {
										if (data.msg.ShowMsg.length > 0) {
											newerror.text("");
											confirmerror.text(data.msg.ShowMsg);
										} else {
											if (data.msg.Exist) {
												$('.contentfillin').css("display", "none")
												$('.contentverifi').css("display", "none")
												$('.contentreset').css("display", "none")
												$('.contentcomplete').css("display", "")

												$('.fillinactive').css("background-image", "url('/static/img/step/1-2.png')")
												$('.verifiactive').css("background-image", "url('/static/img/step/2-2.png')")
												$('.resetactive').css("background-image", "url('/static/img/step/2-2.png')")
												$('.completeactive').css("background-image", "url('/static/img/step/3-1.png')")
											}
										}
									}
								}
							});
						} else {
							confirmerror.text("新密码和确认密码不一致");
						}
					} else {
						confirmerror.text("确认密码不能为空");
					}
				} else {
					newerror.text("新密码最小长度为6位");
				}

			} else {
				newerror.text("新密码不能为空");
			}
		}
	});

	//发送验证码
	$(".message-tel .send-prove-code").click(function() {
		if (curCount == 0) {
			curCount = count;
			var errorcontrol = $('.prove-code-content .prove-code-info');
			if (UserData != null) {
				$.ajax({
					url: '/index/send/code',
					type: 'POST',
					data: {
						uname: UserData.Uname,
						uphone: UserData.AecPhone
					},
					success: function(data) {
						if (data.msg != null) {
							UserData = data.msg;
							if (UserData.ShowMsg.length > 0) {
								errorcontrol.text(UserData.ShowMsg);
							} else {
								$(".message-tel .send-prove-code").attr("disabled", "true");
								$(".message-tel .send-prove-code").html(curCount + "秒之后");
								$(".message-tel .send-prove-code").css('background-color', '#ADABAB');

								InterValObj = window.setInterval(SetRemainTime, 1000); //启动计时器，1秒执行一次
								errorcontrol.text("验证码15分钟有效");
							}
						}
					}
				});
			} else {
				errorcontrol.text("验证码发送失败");
			}
		}
	});

	$('.top_right ul li[name=return]').click(function() {
		window.location = "/";
	});

	$('.top_right ul li[name=login]').click(function() {
		UserLogin(1);
	});

	$('span[name=register]').click(function() {
		UserLogin(2);
	});


	$('.top_right ul li[name=return]').hover(function() {
		$('.top_right ul li[name=return] img').attr("src", "/static/img/step/b_2.png");
		$('.top_right ul li[name=return] span').css("color", "#17a3dd")
	}, function() {
		$('.top_right ul li[name=return] img').attr("src", "/static/img/step/b_1.png");
		$('.top_right ul li[name=return] span').css("color", "#616169");
	});

	$('.top_right ul li[name=login]').hover(function() {
		$('.top_right ul li[name=login] img').attr("src", "/static/img/step/c_2.png");
		$('.top_right ul li[name=login] span').css("color", "#17a3dd")
	}, function() {
		$('.top_right ul li[name=login] img').attr("src", "/static/img/step/c_1.png");
		$('.top_right ul li[name=login] span').css("color", "#616169");
	});

	$('span[name=register]').hover(function() {
		$(this).css("color", "#17a3dd")
	}, function() {
		$(this).css("color", "#616169");
	});

	$('.top_right ul li[name=return]').css("cursor", "pointer");
	$('.top_right ul li[name=login]').css("cursor", "pointer");
	$('span[name=register]').css("cursor", "pointer");



	$('.contentfillin').css("display", "")
	$('.contentverifi').css("display", "none")
	$('.contentreset').css("display", "none")
	$('.contentcomplete').css("display", "none")
});


//timer处理函数
function SetRemainTime() {
	if (curCount == 0) {
		window.clearInterval(InterValObj); //停止计时器
		$(".message-tel .send-prove-code").removeAttr("disabled"); //启用按钮
		$(".message-tel .send-prove-code").html("重新发送");
		$(".message-tel .send-prove-code").css('background-color', '#2795dc');
	} else {
		curCount--;
		$(".message-tel .send-prove-code").html(curCount + "秒之后");
	}
}


function UserLogin(type) {
	layer.open({
		type: 2,
		title: '用户登录',
		shadeClose: true,
		shade: [0.01],
		area: ['500px', '450px'],
		content: '/reglogin?type=' + type
	});
}
