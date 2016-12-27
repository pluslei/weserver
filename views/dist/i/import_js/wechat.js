  
  var voiceindex = 0; // 声音样式
  var voiceobj = Object//
  var voice = {
      localId: '',
      serverId: ''
  };

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


  //上传录音
  function uploadVoice(localid){
    wx.uploadVoice({
        localId: localid, 
        isShowProgressTips: 1, 
        success: function (res) {
          console.info("voice.localId res",res,res.serverId);
          // TODO进行数据库入库操作
          // var voicedata = '<img src="../i/videomyself.png" alt="videomyself" class="videomyself am-margin-left-sm playVoice" data-status="start" data-voicersrc="'+res.serverId+'" onclick="playVoice(this)" />'
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
      wx.stopVoice({
        localId: voice.serverId
      });
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
    // // alert(localid);
    // $('.playVoice').attr('src', '../i/videomyself.gif');
    // $(".playVoice").attr('data-status','start');
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