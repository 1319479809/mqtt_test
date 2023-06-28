package utils

//错误代码定义
const (
	EBUSY  = 40001
	EBUSYS = "系统繁忙请稍后再试"

	EEXPIRE  = 40002
	EEXPIRES = "您的账号已经过期该部分功能已经被限制"

	EAUTHLOGIN  = 401             //权限错误
	EAUTHLOGINS = "登录已经过期了需要重新登录" //权限错误

	SUCCESS      = 0
	NOFOUNTUSER  = 2
	EUNKNOWN     = 10001 //未知错误
	EDBERR       = 10002 //数据库错误
	EMTHNOTFOUND = 10003 //方法不存在
	EDTNOTFOUND  = 10004 //数据不存在
	ERENDER      = 10008 //渲染出错
	EAUTH        = 10009 //权限错误
	EFORMAT      = 10010 //参数格式错误
	EDATA        = 10011 //数据错误
	EMPLOGIN     = 10012 //登录错误
	EPASSWORD    = 10013 //密码错误

	ELOGINS = 20001
	ELOGINC = 20002
	ECAPT   = 20003
	EECAPT  = 20004
	EPASSWD = 20005
	ELOGIND = 20006

	EECAPTEXP = 20007

	ESYN  = 30001
	ESYNS = "命令发送失败"

	EUNKNOWNS     = "未知错误"
	EDBERRS       = "数据库错误"
	EMTHNOTFOUNDS = "方法不存在"
	EDTNOTFOUNDS  = "数据不存在"

	EAUTHS     = "权限错误"
	EFORMATS   = "数据格式错误"
	EDATAS     = "数据错误"
	SUCCESSS   = "成功"
	SUCCESSSYN = "成功,请等待同步"
	SUCCESSAS  = "添加成功"
	SUCCESSUS  = "修改成功"
	SUCCESSDS  = "删除成功"

	ELOGINSS     = "您已经登录了"
	NOFOUNTUSERS = "用户不存在"
	ELOGINCS     = "请合理使用资源"
	ELOGINDS     = "您的账号被禁用，请联系管理员"
	ECAPTS       = "验证码错误"
	EUSERPASSWDS = "请检查用户名和密码"
	EPASSWORDS   = "密码错误"

	SUCCESSLOGINS = "登陆成功"
	SUCCEPASSWD   = "密码修改成功"
	SUCCCAPT      = "验证码发送成功"

	CFILENOTFOUND = 90001
	CMD5          = 90002
	CTORAGE       = 9004

	// EFILENOTFOUND = "没有找到文件"
	// EMD5          = "MD5错误"
	// ESUCCESS      = "上传成功"
	// EFIAL         = "上传失败"
	// ETORAGE       = "存储空间不足"
)
