package errorAssets

import "github.com/ken-house/go-contrib/prototype/errorAssets"

var (
	// ERR_PARAM 公共错误
	ERR_PARAM         = errorAssets.NewError(10000, "参数错误")
	ERR_SYSTEM        = errorAssets.NewError(10001, "系统错误")
	ERR_CERT          = errorAssets.NewError(10002, "证书错误")
	ERR_CALL_FUNC     = errorAssets.NewError(10003, "调用方法出错")
	ERR_DIAL          = errorAssets.NewError(10004, "连接错误")
	ERR_SMS_SEND_FAIL = errorAssets.NewError(10006, "短信发送失败")

	// ERR_LOGIN 登录注册认证相关错误
	ERR_LOGIN         = errorAssets.NewError(20000, "用户名或密码不正确")
	ERR_LOGIN_REMOTE  = errorAssets.NewError(20001, "账号已在其他设备登录")
	ERR_LOGIN_FAILURE = errorAssets.NewError(20002, "当前登录已失效，请尝试请求refresh_token获取新令牌")
	ERR_REFRESH_TOKEN = errorAssets.NewError(20003, "刷新令牌失败，请重新登录")

	// ERR_EXPOSE 导入导出相关错误
	ERR_EXPORT     = errorAssets.NewError(30000, "导出失败")
	ERR_IMPORT     = errorAssets.NewError(30001, "导入失败")
	ERR_FILE_PARSE = errorAssets.NewError(30002, "文件解析失败")
)
