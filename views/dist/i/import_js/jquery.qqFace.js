// QQ表情插件
(function($) {
	$.fn.qqFace = function(options) {
		var defaults = {
			id: 'facebox',
			path: 'i/face/',
			assign: 'content',
			tip: 'em_',
			faceArr: ["[微笑]", "[撇嘴]", "[色]", "[发呆]", "[得意]", "[流泪]", "[害羞]", "[闭嘴]", "[睡]", "[大哭]", "[尴尬]", "[发怒]", "[调皮]", "[呲牙]", "[惊讶]", "[难过]", "[酷]", "[冷汗]", "[抓狂]", "[吐]", "[偷笑]", "[愉快]", "[白眼]", "[傲慢]", "[饥饿]", "[困]", "[惊恐]", "[流汗]", "[憨笑]", "[悠闲]", "[奋斗]", "[咒骂]", "[疑问]", "[嘘]", "[晕]", "[疯了]", "[衰]", "[骷髅]", "[敲打]", "[再见]", "[擦汗]", "[抠鼻]", "[鼓掌]", "[糗大了]", "[坏笑]", "[左哼哼]", "[右哼哼]", "[哈欠]", "[鄙视]", "[委屈]", "[快哭了]", "[阴险]", "[亲亲]", "[吓]", "[可怜]", "[菜刀]", "[西瓜]", "[啤酒]", "[篮球]", "[乒乓]", "[咖啡]", "[饭]", "[猪头]", "[玫瑰]", "[凋谢]", "[嘴唇]", "[爱心]", "[心碎]", "[蛋糕]", "[闪电]", "[炸弹]", "[刀]", "[足球]", "[瓢虫]", "[便便]", "[月亮]", "[太阳]", "[礼物]", "[拥抱]", "[强]", "[弱]", "[握手]", "[胜利]", "[抱拳]", "[勾引]", "[拳头]", "[差劲]", "[爱你]", "[NO]", "[OK]", "[爱情]", "[飞吻]", "[跳跳]", "[发抖]", "[怄火]", "[转圈]", "[磕头]", "[回头]", "[跳绳]", "[投降]"]
		};
		var option = $.extend(defaults, options);
		var assign = $('#' + option.assign);
		var id = option.id;
		var path = option.path;
		var tip = option.tip;

		$(this).click(function(e) {
			//var objId = $(this).parent().siblings('div').attr("id");
			var objId = 'msgcontent';
			var strFace, labFace;
			if ($('#' + id).length <= 0) {
				strFace = '<div id="' + id + '" style="position:absolute;display:none;z-index:1000;" class="qqFace">' +
					'<table border="0" cellspacing="0" cellpadding="0"><tr>';
				for (var i = 1; i <= option.faceArr.length; i++) {
					labFace = option.path + i + ".gif";
					labflag = option.faceArr[i - 1]
					strFace += '<td><img src="' + labFace + '" onclick="$(\'#' + option.assign + '\').insertAtCaret(\'' + labFace + '\',\'' + objId + '\',\'' + labflag + '\');" /></td>';
					if (i % 15 == 0) strFace += '</tr><tr>';
				}
				strFace += '</tr></table></div>';
			}
			//$(this).parent().append(strFace);
			$('#faceurl').append(strFace);
			var offset = $(this).position();
			var top = offset.top + $(this).outerHeight();
			$('#' + id).css('top', top);
			$('#' + id).css('left', offset.left);
			$('#' + id).show();
			e.stopPropagation();
		});

		$(document).click(function() {
			$('#' + id).hide();
			$('#' + id).remove();
		});
	};


})(jQuery);

jQuery.extend({
	unselectContents: function() {
		if (window.getSelection)
			window.getSelection().removeAllRanges();
		else if (document.selection)
			document.selection.empty();
	},
	replaceReceiveHtml: function(msg) {
		msg = msg.replace(/\[微笑\]/g, "<img data-face='[微笑]' src='/static/images/face/1.gif' />");
		msg = msg.replace(/\[撇嘴\]/g, "<img data-face='[撇嘴]' src='/static/images/face/2.gif' />");
		msg = msg.replace(/\[色\]/g, "<img data-face='[色]' src='/static/images/face/3.gif' />");
		msg = msg.replace(/\[发呆\]/g, "<img data-face='[发呆]' src='/static/images/face/4.gif' />");
		msg = msg.replace(/\[得意\]/g, "<img data-face='[得意]' src='/static/images/face/5.gif' />");
		msg = msg.replace(/\[流泪\]/g, "<img data-face='[流泪]' src='/static/images/face/6.gif' />");
		msg = msg.replace(/\[害羞\]/g, "<img data-face='[害羞]' src='/static/images/face/7.gif' />");
		msg = msg.replace(/\[闭嘴\]/g, "<img data-face='[闭嘴]' src='/static/images/face/8.gif' />");
		msg = msg.replace(/\[睡\]/g, "<img data-face='[睡]' src='/static/images/face/9.gif' />");
		msg = msg.replace(/\[大哭\]/g, "<img data-face='[大哭]' src='/static/images/face/10.gif' />");
		msg = msg.replace(/\[尴尬\]/g, "<img data-face='[尴尬]' src='/static/images/face/11.gif' />");
		msg = msg.replace(/\[发怒\]/g, "<img data-face='[发怒]' src='/static/images/face/12.gif' />");
		msg = msg.replace(/\[调皮\]/g, "<img data-face='[调皮]' src='/static/images/face/13.gif' />");
		msg = msg.replace(/\[呲牙\]/g, "<img data-face='[呲牙]' src='/static/images/face/14.gif' />");
		msg = msg.replace(/\[惊讶\]/g, "<img data-face='[惊讶]' src='/static/images/face/15.gif' />");
		msg = msg.replace(/\[难过\]/g, "<img data-face='[难过]' src='/static/images/face/16.gif' />");
		msg = msg.replace(/\[酷\]/g, "<img data-face='[酷]' src='/static/images/face/17.gif' />");
		msg = msg.replace(/\[冷汗\]/g, "<img data-face='[冷汗]' src='/static/images/face/18.gif' />");
		msg = msg.replace(/\[抓狂\]/g, "<img data-face='[抓狂]' src='/static/images/face/19.gif' />");
		msg = msg.replace(/\[吐\]/g, "<img data-face='[吐]' src='/static/images/face/20.gif' />");
		msg = msg.replace(/\[偷笑\]/g, "<img data-face='[偷笑]' src='/static/images/face/21.gif' />");
		msg = msg.replace(/\[愉快\]/g, "<img data-face='[愉快]' src='/static/images/face/22.gif' />");
		msg = msg.replace(/\[白眼\]/g, "<img data-face='[白眼]' src='/static/images/face/23.gif' />");
		msg = msg.replace(/\[傲慢\]/g, "<img data-face='[傲慢]' src='/static/images/face/24.gif' />");
		msg = msg.replace(/\[饥饿\]/g, "<img data-face='[饥饿]' src='/static/images/face/25.gif' />");
		msg = msg.replace(/\[困\]/g, "<img data-face='[困]' src='/static/images/face/26.gif' />");
		msg = msg.replace(/\[惊恐\]/g, "<img data-face='[惊恐]' src='/static/images/face/27.gif' />");
		msg = msg.replace(/\[流汗\]/g, "<img data-face='[流汗]' src='/static/images/face/28.gif' />");
		msg = msg.replace(/\[憨笑\]/g, "<img data-face='[憨笑]' src='/static/images/face/29.gif' />");
		msg = msg.replace(/\[悠闲\]/g, "<img data-face='[悠闲]' src='/static/images/face/30.gif' />");
		msg = msg.replace(/\[奋斗\]/g, "<img data-face='[奋斗]' src='/static/images/face/31.gif' />");
		msg = msg.replace(/\[咒骂\]/g, "<img data-face='[咒骂]' src='/static/images/face/32.gif' />");
		msg = msg.replace(/\[疑问\]/g, "<img data-face='[疑问]' src='/static/images/face/33.gif' />");
		msg = msg.replace(/\[嘘\]/g, "<img data-face='[嘘]' src='/static/images/face/34.gif' />");
		msg = msg.replace(/\[晕\]/g, "<img data-face='[晕]' src='/static/images/face/35.gif' />");
		msg = msg.replace(/\[疯了\]/g, "<img data-face='[疯了]' src='/static/images/face/36.gif' />");
		msg = msg.replace(/\[衰\]/g, "<img data-face='[衰]' src='/static/images/face/37.gif' />");
		msg = msg.replace(/\[骷髅\]/g, "<img data-face='[骷髅]' src='/static/images/face/38.gif' />");
		msg = msg.replace(/\[敲打\]/g, "<img data-face='[敲打]' src='/static/images/face/39.gif' />");
		msg = msg.replace(/\[再见\]/g, "<img data-face='[再见]' src='/static/images/face/40.gif' />");
		msg = msg.replace(/\[擦汗\]/g, "<img data-face='[擦汗]' src='/static/images/face/41.gif' />");
		msg = msg.replace(/\[抠鼻\]/g, "<img data-face='[抠鼻]' src='/static/images/face/42.gif' />");
		msg = msg.replace(/\[鼓掌\]/g, "<img data-face='[鼓掌]' src='/static/images/face/43.gif' />");
		msg = msg.replace(/\[糗大了\]/g, "<img data-face='[糗大了]' src='/static/images/face/44.gif' />");
		msg = msg.replace(/\[坏笑\]/g, "<img data-face='[坏笑]' src='/static/images/face/45.gif' />");
		msg = msg.replace(/\[左哼哼\]/g, "<img data-face='[左哼哼]' src='/static/images/face/46.gif' />");
		msg = msg.replace(/\[右哼哼\]/g, "<img data-face='[右哼哼]' src='/static/images/face/47.gif' />");
		msg = msg.replace(/\[哈欠\]/g, "<img data-face='[哈欠]' src='/static/images/face/48.gif' />");
		msg = msg.replace(/\[鄙视\]/g, "<img data-face='[鄙视]' src='/static/images/face/49.gif' />");
		msg = msg.replace(/\[委屈\]/g, "<img data-face='[委屈]' src='/static/images/face/50.gif' />");
		msg = msg.replace(/\[快哭了\]/g, "<img data-face='[快哭了]' src='/static/images/face/51.gif' />");
		msg = msg.replace(/\[阴险\]/g, "<img data-face='[阴险]' src='/static/images/face/52.gif' />");
		msg = msg.replace(/\[亲亲\]/g, "<img data-face='[亲亲]' src='/static/images/face/53.gif' />");
		msg = msg.replace(/\[吓\]/g, "<img data-face='[吓]' src='/static/images/face/54.gif' />");
		msg = msg.replace(/\[可怜\]/g, "<img data-face='[可怜]' src='/static/images/face/55.gif' />");
		msg = msg.replace(/\[菜刀\]/g, "<img data-face='[菜刀]' src='/static/images/face/56.gif' />");
		msg = msg.replace(/\[西瓜\]/g, "<img data-face='[西瓜]' src='/static/images/face/57.gif' />");
		msg = msg.replace(/\[啤酒\]/g, "<img data-face='[啤酒]' src='/static/images/face/58.gif' />");
		msg = msg.replace(/\[篮球\]/g, "<img data-face='[篮球]' src='/static/images/face/59.gif' />");
		msg = msg.replace(/\[乒乓\]/g, "<img data-face='[乒乓]' src='/static/images/face/60.gif' />");
		msg = msg.replace(/\[咖啡\]/g, "<img data-face='[咖啡]' src='/static/images/face/61.gif' />");
		msg = msg.replace(/\[饭\]/g, "<img data-face='[饭]' src='/static/images/face/62.gif' />");
		msg = msg.replace(/\[猪头\]/g, "<img data-face='[猪头]' src='/static/images/face/63.gif' />");
		msg = msg.replace(/\[玫瑰\]/g, "<img data-face='[玫瑰]' src='/static/images/face/64.gif' />");
		msg = msg.replace(/\[凋谢\]/g, "<img data-face='[凋谢]' src='/static/images/face/65.gif' />");
		msg = msg.replace(/\[嘴唇\]/g, "<img data-face='[嘴唇]' src='/static/images/face/66.gif' />");
		msg = msg.replace(/\[爱心\]/g, "<img data-face='[爱心]' src='/static/images/face/67.gif' />");
		msg = msg.replace(/\[心碎\]/g, "<img data-face='[心碎]' src='/static/images/face/68.gif' />");
		msg = msg.replace(/\[蛋糕\]/g, "<img data-face='[蛋糕]' src='/static/images/face/69.gif' />");
		msg = msg.replace(/\[闪电\]/g, "<img data-face='[闪电]' src='/static/images/face/70.gif' />");
		msg = msg.replace(/\[炸弹\]/g, "<img data-face='[炸弹]' src='/static/images/face/71.gif' />");
		msg = msg.replace(/\[刀\]/g, "<img data-face='[刀]' src='/static/images/face/72.gif' />");
		msg = msg.replace(/\[足球\]/g, "<img data-face='[足球]' src='/static/images/face/73.gif' />");
		msg = msg.replace(/\[瓢虫\]/g, "<img data-face='[瓢虫]' src='/static/images/face/74.gif' />");
		msg = msg.replace(/\[便便\]/g, "<img data-face='[便便]' src='/static/images/face/75.gif' />");
		msg = msg.replace(/\[月亮\]/g, "<img data-face='[月亮]' src='/static/images/face/76.gif' />");
		msg = msg.replace(/\[太阳\]/g, "<img data-face='[太阳]' src='/static/images/face/77.gif' />");
		msg = msg.replace(/\[礼物\]/g, "<img data-face='[礼物]' src='/static/images/face/78.gif' />");
		msg = msg.replace(/\[拥抱\]/g, "<img data-face='[拥抱]' src='/static/images/face/79.gif' />");
		msg = msg.replace(/\[强\]/g, "<img data-face='[强]' src='/static/images/face/80.gif' />");
		msg = msg.replace(/\[弱\]/g, "<img data-face='[弱]' src='/static/images/face/81.gif' />");
		msg = msg.replace(/\[握手\]/g, "<img data-face='[握手]' src='/static/images/face/82.gif' />");
		msg = msg.replace(/\[胜利\]/g, "<img data-face='[胜利]' src='/static/images/face/83.gif' />");
		msg = msg.replace(/\[抱拳\]/g, "<img data-face='[抱拳]' src='/static/images/face/84.gif' />");
		msg = msg.replace(/\[勾引\]/g, "<img data-face='[勾引]' src='/static/images/face/85.gif' />");
		msg = msg.replace(/\[拳头\]/g, "<img data-face='[拳头]' src='/static/images/face/86.gif' />");
		msg = msg.replace(/\[差劲\]/g, "<img data-face='[差劲]' src='/static/images/face/87.gif' />");
		msg = msg.replace(/\[爱你\]/g, "<img data-face='[爱你]' src='/static/images/face/88.gif' />");
		msg = msg.replace(/\[NO\]/g, "<img data-face='[NO]' src='/static/images/face/89.gif' />");
		msg = msg.replace(/\[OK\]/g, "<img data-face='[OK]' src='/static/images/face/90.gif' />");
		msg = msg.replace(/\[爱情\]/g, "<img data-face='[爱情]' src='/static/images/face/91.gif' />");
		msg = msg.replace(/\[飞吻\]/g, "<img data-face='[飞吻]' src='/static/images/face/92.gif' />");
		msg = msg.replace(/\[跳跳\]/g, "<img data-face='[跳跳]' src='/static/images/face/93.gif' />");
		msg = msg.replace(/\[发抖\]/g, "<img data-face='[发抖]' src='/static/images/face/94.gif' />");
		msg = msg.replace(/\[怄火\]/g, "<img data-face='[怄火]' src='/static/images/face/95.gif' />");
		msg = msg.replace(/\[转圈\]/g, "<img data-face='[转圈]' src='/static/images/face/96.gif' />");
		msg = msg.replace(/\[磕头\]/g, "<img data-face='[磕头]' src='/static/images/face/97.gif' />");
		msg = msg.replace(/\[回头\]/g, "<img data-face='[回头]' src='/static/images/face/98.gif' />");
		msg = msg.replace(/\[跳绳\]/g, "<img data-face='[跳绳]' src='/static/images/face/99.gif' />");
		msg = msg.replace(/\[投降\]/g, "<img data-face='[投降]' src='/static/images/face/100.gif' />");
		return msg;
	}

});
jQuery.fn.extend({
	selectContents: function() {
		$(this).each(function(i) {
			var node = this;
			var selection, range, doc, win;
			if ((doc = node.ownerDocument) && (win = doc.defaultView) && typeof win.getSelection != 'undefined' && typeof doc.createRange != 'undefined' && (selection = window.getSelection()) && typeof selection.removeAllRanges != 'undefined') {
				range = doc.createRange();
				range.selectNode(node);
				if (i == 0) {
					selection.removeAllRanges();
				}
				selection.addRange(range);
			} else if (document.body && typeof document.body.createTextRange != 'undefined' && (range = document.body.createTextRange())) {
				range.moveToElementText(node);
				range.select();
			}
		});
	},

	setCaret: function() {
		if (!$.browser.msie) return;
		var initSetCaret = function() {
			var textObj = $(this).get(0);
			textObj.caretPos = document.selection.createRange().duplicate();
		};
		$(this).click(initSetCaret).select(initSetCaret).keyup(initSetCaret);
	},

	insertAtCaret: function(textFeildValue, textObj, lableflag) {
		$("#" + textObj).append("<img data-face='" + lableflag + "' src='" + textFeildValue + "'/>");
	}
});
