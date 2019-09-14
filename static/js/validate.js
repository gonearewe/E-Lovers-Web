$(function () {
    $('#form-register').bootstrapValidator({
      message: 'This value is not valid',
      feedbackIcons: {
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
      },
      live: 'enabled',
      submitButtons: 'button[type="submit"]',
      fields: {
        username: {
          message: '用户名验证失败',
          validators: {
            notEmpty: {
              message: '用户名不能为空'
            },
            stringLength: {
              min: 4,
              max: 18,
              message: '用户名长度必须在4到18位之间'
            },
            regexp: {
              regexp: /^[a-zA-Z0-9_]+$/,
              message: '用户名只能包含大写、小写、数字和下划线'
            }
          }
        },
        password: {
          validators: {
            notEmpty: {
              message: '密码不能为空'
            },
            stringLength: {
                min: 8,
                max: 24,
                message: '密码长度必须在8到24位之间'
              },
            regexp: {
                regexp: /^[a-zA-Z0-9_]+$/,
                message: '密码只能包含大写、小写、数字和下划线'
            }
        }
      }
    },
    submitHandler: function (form) {
      var urlStr = "/register";
      // alert("urlStr:"+urlStr)
      $(form).ajaxSubmit({
          url:urlStr,
          async:false,
          type:"post",
          dataType:"json",
          cache:false,
          data:{"username":$('#username').val(),"password":$('#password').val(),"repassword":$('#repassword').val()},
          success:function (data,status) {
              alert("data:"+data.message)
              if (data.code == 0){
                  setTimeout(function () {
                      window.location.href="/login"
                  },1000)
              }
          },
          err:function (data,status) {
              alert("err:"+data.message+":"+status)
          }
      })
  }
  });
  
  $('#form-login').bootstrapValidator({
    message: 'This value is not valid',
    submitButtons: 'button[type="submit"]',
    feedbackIcons: {
      valid: 'glyphicon glyphicon-ok',
      invalid: 'glyphicon glyphicon-remove',
      validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    fields: {
      username: {
        message: '用户名验证失败',
        validators: {
          notEmpty: {
            message: '用户名不能为空'
          },
          stringLength: {
            min: 4,
            max: 18,
            message: '用户名长度必须在4到18位之间'
          },
          regexp: {
            regexp: /^[a-zA-Z0-9_]+$/,
            message: '用户名只能包含大写、小写、数字和下划线'
          }
        }
      },
      password: {
        validators: {
          notEmpty: {
            message: '密码不能为空'
          },
          stringLength: {
              min: 8,
              max: 24,
              message: '密码长度必须在8到24位之间'
            },
          regexp: {
              regexp: /^[a-zA-Z0-9_]+$/,
              message: '密码只能包含大写、小写、数字和下划线'
          }
      }
    }
  },
  submitHandler: function (form) {
    var urlStr = "/login";
    alert("urlStr:"+urlStr)
    $(form).ajaxSubmit({
        url:urlStr,
        async:false,
        type:"post",
        dataType:"json",
        cache:false,
        data:{"username":$('#login-username').val(),"password":$('#login-password').val()},
        success:function (data,status) {
            alert("data:"+data.message)
            if (data.code == 0){
                setTimeout(function () {
                    window.location.href="/"
                },1000)
            }
        },
        err:function (data,status) {
            alert("err:"+data.message+":"+status)
        }
    })
}
});



//添加文章
$('#form-article').bootstrapValidator({
  message: 'This value is not valid',
  submitButtons: 'button[type="submit"]',
  feedbackIcons: {
    valid: 'glyphicon glyphicon-ok',
    invalid: 'glyphicon glyphicon-remove',
    validating: 'glyphicon glyphicon-refresh'
  },
  live: 'enabled',
  fields: {
    title: {
      message: '文章上传失败',
      validators: {
        notEmpty: {
          message: '标题不能为空'
        },
        stringLength: {
          min: 4,
          max: 48,
          message: '标题长度必须在4到48位之间'
        },
      }
    },
    tags: {
      message: '文章上传失败',
      validators: {
        notEmpty: {
          message: '标签不能为空'
        },
      }
    },
    article: {
      message: '文章上传失败',
      validators: {
        notEmpty: {
          message: '需要选择文件'
        },
      }
    },
},
submitHandler: function (form) {
  var urlStr = "/article/add";
  var fd = new FormData();
  //fd.append('sealPicPathFile', $("#sealPicPathFile").val()); 不可以这样
fd.append('article', $('#article')[0].files[0]);
fd.append('title', $("#title").val());
fd.append('subtitle', $("#subtitle").val());
fd.append('tags', $("#tags").val());

$.ajax({
  url:urlStr,
  data: fd ,
  type:'post',
  dataType: 'json',
  processData:false,  //tell jQuery not to process the data
      contentType: false,  //tell jQuery not to set contentType
  success:function(request){
    console.log(JSON.stringify(request));
    if(request.code == 0){
      
      alert("保存成功");

      // $("#btnBack").click();
    }else{
       alert("保存失败，" + request.message);
    }
  }
});
}
});

// $('button').click(loginFun)
// function loginFun()
// {
//   $('form').ajaxSubmit({
//     url:'/login',
//     async:false,
//     type:"post",
//     dataType:"json",
//     cache:false,
//     data:{"username":$('#login-username').val(),"password":$('#login-password').val()},
//     success:function (data,status) {
//         alert("data:"+data.message)
//         if (data.code == 0){
//             setTimeout(function () {
//                 window.location.href="/"
//             },1000)
//         }
//     },
//     err:function (data,status) {
//         alert("err:"+data.message+":"+status)
//     }
// })
// }
});

