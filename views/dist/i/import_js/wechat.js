  
  var voiceindex = 0; // 声音样式
  var voiceobj = Object//
  var voice = {
      localId: '',
      serverId: ''
  };
<<<<<<< HEAD
  var images = {
    localId: [],
    serverId: []
  };
=======

  $(function(){
    // 录制声音上传
    $("#startRecord").on('touchstart', function(event) {
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
    });

    $("#startRecord").on('touchend', function(event) {
      $(this).html("按住 说话")
      event.preventDefault();
      END = new Date().getTime();
      if((END - START) < 200){
          END = 0;
          START = 0;
          //小于200ms，不录音
          clearTimeout(recordTimer);
      }else{
          wx.stopRecord({
            success: function (res) {
              voice.localId = res.localId;
              uploadVoice(res.localId);
            },
            fail: function (res) {
              alert(JSON.stringify(res));
            }
          });
      }
    });
  })
>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef

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
  });

  wx.error(function(res) {
    alert(res);
  });

<<<<<<< HEAD
  $(function() {
    var startY, endY;
    document.getElementById("startRecord").addEventListener("touchstart", touchStart, false);
    document.getElementById("startRecord").addEventListener("touchmove", touchMove, false);
    document.getElementById("startRecord").addEventListener("touchend", touchEnd, false);

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

    // 预览上传图片
    $("#uploadImage").click(function(){
      images.localId = [];
      wx.chooseImage({
        success: function (res) {
          images.localId = res.localIds;
          if (images.localId.length == 0) {
              alert('请先使用 chooseImage 接口选择图片');
              return;
          }
          if(images.localId.length > 1) {
              alert('目前仅支持单张图片上传,请重新上传');
              images.localId = [];
              return;
          }
          var i = 0, length = images.localId.length;
          images.serverId = [];
          function upload() {
            wx.uploadImage({
              localId: images.localId[i],
              success: function (res) {
                i++;                             
                if (i < length) {
                    upload();
                }
              },
              fail: function (res) {
                alert(JSON.stringify(res));
              }
            });
          }
          upload();
        }
      });
    })
  })
  
=======
>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef

  //上传录音
  function uploadVoice(localid){
    wx.uploadVoice({
        localId: localid, 
        isShowProgressTips: 1, 
        success: function (res) {
          console.info("voice.localId res",res,res.serverId);
          // TODO进行数据库入库操作
<<<<<<< HEAD
=======
          // var voicedata = '<img src="../i/videomyself.png" alt="videomyself" class="videomyself am-margin-left-sm playVoice" data-status="start" data-voicersrc="'+res.serverId+'" onclick="playVoice(this)" />'
>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef
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
<<<<<<< HEAD
      // 结束播放
      stopPlayVoice(voice.serverId)
      // 改为开始播放样式
=======
      wx.stopVoice({
        localId: voice.serverId
      });
>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef
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
<<<<<<< HEAD
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


=======
    // // alert(localid);
    // $('.playVoice').attr('src', '../i/videomyself.gif');
    // $(".playVoice").attr('data-status','start');
  }

>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef
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
<<<<<<< HEAD
    stopPlayVoice(voice.serverId)
=======
>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef
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
<<<<<<< HEAD
  }

=======
  }
>>>>>>> 11996a6b23372c992730476cbcf434112e9f88ef
