function uploadimgicon() {
	$('#docformfile').uploadify({
		'swf': '/static/img/uploadify.swf', // uploadify.swf 文件的相对JS文件的路径
		'uploader': '/updateusericon', //后台处理程序的相对路径
		'floder': '/upload/usericon',
		'queueID': 'fileQueue',
		'queueSizeLimit': 999999999,
		'uploadLimit': 999999999,
		//'progressData': 'speed',
		'auto': true,
		'multi': false,
		'buttonClass': 'some-class', //自定义的样式类名
		'buttonTitle': '提示文字',
		'buttonImage': '/static/images/i/img/upload_file.png',
		'requeueErrors': true,
		'debug': false,
		'height': 28, // The height of the flash button
		'width': 155, // The width of the flash button
		'itemTemplate': false,
		'successTimeout': 30, //设置文件上传后等待服务器响应的秒数，超出这个时间，将会被认为上传成功，默认为30秒
		'fileTypeExts': '*.jpg;*.jpge;*.gif;*.png',
		'fileSizeLimit': '1MB',
		'fileDataName': 'Filedata',
		'formData': {
			'content': ''
		},
		'onUploadStart': function(file) {
			var contentval = $('span[name=user-login-name]').text();
			$("#docformfile").uploadify("settings", "formData", {
				'content': contentval
			});
		},
		'onUploadSuccess': function(file, data, response) {
			if (data != null) {
				data = jQuery.parseJSON(data);
				if (data.status) {
					var str = data.url
					var strurl = str.substring(str.lastIndexOf('/') + 1, str.length);
					$("img[name=user-login-ico]").attr("src", str);
					var contentval = $('span[name=user-login-name]').text();
					$(".searchList ul li").each(function(index) {
						var spantext = $(".searchList ul li .hyname").eq(index).text();
						if (spantext == contentval) {
							$(".searchList ul li .hytb").eq(index).attr("src", str);
						}
					});

					//修改redius的Icon
					$.ajax({
						type: "post",
						url: "/chat/modify/icon",
						dataType: "json",
						data: {
							uname: contentval,
							uicon: strurl,
							ucodeid: codeid
						},
						success: function(result) {},
						error: function(msg) {
							console.log(msg);
						}
					});
				}
			}
		},
		//加上此句会重写onSelectError方法【需要重写的事件】
		'overrideEvents': ['onSelectError', 'onDialogClose'],
		//返回一个错误，选择文件的时候触发
		'onSelectError': function(file, data, response) {},
	})
}
