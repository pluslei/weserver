
  var voiceindex = 0; // 声音样式
  var voiceobj = Object//
  var voice = {
      localId: '',
      serverId: ''
  };
  var images = {
    localId: [],
    serverId: [],
    viewImages: [],
    localImg:[],
  };

  wx.ready(function () {
      // 监听录音自动停止
      wx.onVoiceRecordEnd({
        complete: function (res) {
          voice.localId = res.localId
          autoRecordStop(res.localId)
        }
      });
      // 监听播放结束
      wx.onVoicePlayEnd({
        complete: function (res) {
          autoPlayStop(res.localId)
        }
      });

      if(!localStorage.rainAllowRecord || localStorage.rainAllowRecord !== 'true'){
          wx.startRecord({
              success: function(){
                  localStorage.rainAllowRecord = 'true';
                  wx.stopRecord();
              },
              cancel: function () {
                  alert('用户拒绝授权录音');
              }
          });
      }

      $(".inputText").each(function(index, el) {
        imgid = $(this).children('span').attr("id");
        if (imgid != undefined) {
          imgid = $(this).children('span').attr("id");
          wx.downloadImage({
              serverId: imgid, // 需要下载的图片的服务器端ID，由uploadImage接口获得
              isShowProgressTips: 0, // 默认为1，显示进度提示
              success: function (res) {
                images.localImg.push(res.localId);
                $("#"+imgid).children("img").attr("src",res.localId);
                console.info("get the images serverid",res,res.localId);
              }
          });
        };
      });

      
  });

  wx.error(function(res) {
    alert(res);
  });



  $(function() {
    var startY, endY;
    // document.getElementById("startRecord").addEventListener("touchstart", touchStart, false);
    // document.getElementById("startRecord").addEventListener("touchmove", touchMove, false);
    // document.getElementById("startRecord").addEventListener("touchend", touchEnd, false);
    
    function touchStart(event) {
      startY = event.touches[0].clientY;
      endY = event.touches[0].clientY;

      // 开始录音
      event.preventDefault();
      START = new Date().getTime();
      recordTimer = setTimeout(function(){
          wx.startRecord({
              success:function (){
                localStorage.rainAllowRecord = 'true';
              },
              cancel: function () {
                  alert('用户拒绝授权录音');
                  $(this).html("按住 说话")
              }
          });
      },200);
      $(this).html("松开 结束") 
    }

    function touchMove(event) {
      endY = event.touches[0].clientY;
      console.info("endY的值",endY);
    }

    function touchEnd(event) {
      $(this).html("按住 说话")
      event.preventDefault();
      END = new Date().getTime();
      if((END - START) < 200){  //小于200ms，不录音
          START = 0;
          END = 0;
          clearTimeout(recordTimer);
          return false;
      }

      if (startY - endY <= 100) {
        wx.stopRecord({
          success: function (res) {
            voice.localId = res.localId;
            uploadVoice(res.localId);
          },
          fail: function (res) {
            alert(JSON.stringify(res));
          }
        });
      }else{
        // TODO 撤销语音
      }
      console.info("Y轴移动大小：" + (startY - endY));
    }

    // 上传图片
    $("#uploadImage").click(function(){
      count:1,
      wx.chooseImage({
        success: function (res) {
          var localIds = res.localIds;
          console.info("1:",localIds)
          syncUpload(localIds);
        }
      });
    })

  })
  
  // 单张图片上传
  var syncUpload = function(localIds){
    var localId = localIds.pop();
    console.info("2:",localId)
    wx.uploadImage({
      localId: localId,
      isShowProgressTips: 1,
      success: function (res) {
        var serverId = res.serverId; // 返回图片的服务器端ID
        console.info("3:",serverId)
        //TODO 其他对serverId做处理的代码
        if(serverId.length > 0){
          const sendmsg = {
            Chat: 'allchat',
            Codeid: userinfor.Codeid,
            Author: userinfor.Nickname,
            UserIcon: userinfor.UserIcon,
            RoleName: userinfor.RoleName,
            Titlerole: userinfor.Titlerole,
            Authorcss: userinfor.Authorcss,
            Insider: userinfor.Insider,
            Content: serverId,
            IsLogin: userinfor.IsLogin,
            Sendtype: 'IMG',
          };
          const sendstr = base64encode(sendmsg);
          mysocket.emit('all message', sendstr)
        }
        syncDownloadImg(serverId)
      }
    });
  };

  // 下载图片文件
  function syncDownloadImg(serverId){
    console.info("syncDownloadImg:",serverId)
    wx.downloadImage({
        serverId: serverId, // 需要下载的图片的服务器端ID，由uploadImage接口获得
        isShowProgressTips: 0, // 默认为1，显示进度提示
        success: function (res) {
          images.localImg.push(res.localId);
          $("#"+serverId).children("img").attr("src",res.localId);
          console.info("get the images serverid",res,res.localId);
        }
    });
  }

  // 预览图片文件
  function preViewImage(obj){
    currentimg = $(obj).children("img").attr("src");
    if (currentimg.length > 0) {
      images.viewImages.push(currentimg);
      wx.previewImage({
        current: currentimg, // 当前显示图片的http链接
        urls: images.viewImages // 需要预览的图片http链接列表
      });
    };
  }


  //上传录音
  function uploadVoice(localid){
    wx.uploadVoice({
        localId: localid, 
        isShowProgressTips: 1, 
        success: function (res) {
          console.info("voice.localId res",res,res.serverId);
          // TODO进行数据库入库操作
          const sendmsg = {
            Chat: 'allchat',
            Codeid: userinfor.Codeid,
            Author: userinfor.Nickname,
            UserIcon: userinfor.UserIcon,
            RoleName: userinfor.RoleName,
            Titlerole: userinfor.Titlerole,
            Authorcss: userinfor.Authorcss,
            Insider: userinfor.Insider,
            Content: res.serverId,
            IsLogin: userinfor.IsLogin,
            Sendtype: 'VOICE',
          };
          const sendstr = base64encode(sendmsg);
          mysocket.emit('all message', sendstr);
        }
    });
  }

  // 播放录音
  function playVoice(obj, sel){
    status = $(obj).attr("data-status");
    voicesrc = $(obj).attr("data-voicersrc")
    
    stopLastVoice();
    voiceindex = sel;
    voiceobj = obj;
    console.info(status,voicesrc);
    if (status == 'start') {
      wx.downloadVoice({
          serverId: voicesrc, 
          isShowProgressTips: 1, 
          success: function (res) {
              voice.serverId = res.localId;
              console.info("voicesrc:",voice.serverId);
              if (res.localId == '') {
                alert('语音播放失败');
                return;
              }

              overPlayStyle(obj,sel)
              wx.playVoice({
                localId: res.localId
              });
          }
      });
    }else{
      // 结束播放
      stopPlayVoice(voice.serverId)
      // 改为开始播放样式
      startPlayStyle(obj,sel)
    }
  }

  // 监听自动停止
  function autoRecordStop(localid){
    uploadVoice(localid)
    $("#startRecord").html("按住 说话")
  }

  // 监听播放停止
  function autoPlayStop(localid){
    stopLastVoice()    
  }

  // 结束播放接口
  function stopPlayVoice(voiceid){
    wx.stopVoice({
      localId: voiceid
    });
  }



  // 预览图片
  function previewImage(imgurl){
    wx.previewImage({
        current: imgurl, // 当前显示图片的http链接
        urls: [] // 需要预览的图片http链接列表
    });
  }


  // 开始播放时候样式
  function startPlayStyle(obj,sel){
    switch(sel) {
      case 100:
        {
          $(obj).attr("src","../i/videoOther.png")
          break;
        }
      case 101:
        {
          $(obj).attr("src","../i/videoGuanli.png")
          break;
        }
      case 102:
        {
          $(obj).attr("src","../i/videomyself.png")
          break;
        }
      default:
        break;
    }
    $(obj).attr("data-status","start");
  }

  // 结束播放时候样式
  function overPlayStyle(obj,sel){
    switch(sel) {
      case 100:
        {
          $(obj).attr("src","../i/video.gif")
          break;
        }
      case 101:
        {
          $(obj).attr("src","../i/video.gif")
          break;
        }
      case 102:
        {
          $(obj).attr("src","../i/myselfa.gif")
          break;
        }
      default:
        break;
    }
    $(obj).attr("data-status","end");
  }

  // 结束上一段语言
  function stopLastVoice(){
    stopPlayVoice(voice.serverId)
    switch(voiceindex) {
      case 100:
        {
          $(voiceobj).attr("src","../i/videoOther.png")
          break;
        }
      case 101:
        {
          $(voiceobj).attr("src","../i/videoGuanli.png")
          break;
        }
      case 102:
        {
          $(voiceobj).attr("src","../i/videomyself.png")
          break;
        }
      default:
        break;
    }
    $(voiceobj).attr("data-status","start");
  }

