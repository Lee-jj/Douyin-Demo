package common

import "errors"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var (
	ErrorUserNameEmpty    = errors.New("用户名为空")
	ErrorUserNameInvalid  = errors.New("用户名长度超过16位")
	ErrorPasswordEmpty    = errors.New("密码为空")
	ErrorPasswordInvalid  = errors.New("密码长度不足6位或超过16位")
	ErrorGetConfigFaild   = errors.New("加载配置文件失败")
	ErrorDBMigrateFaild   = errors.New("数据库迁移失败")
	ErrorUserNotFound     = errors.New("未查询到该用户")
	ErrorSQLFaild         = errors.New("SQL执行错误")
	ErrorUserExist        = errors.New("该用户已经存在")
	ErrorCreateUserFaild  = errors.New("创建用户失败")
	ErrorPasswordWrong    = errors.New("密码错误")
	ErrorHasNoToken       = errors.New("用户未登录")
	ErrorHasNoTitle       = errors.New("标题为空")
	ErrorTokenFaild       = errors.New("Token解析失败")
	ErrorCreateVideoFaild = errors.New("上传视频数据库失败")
	ErrorTokenExpired     = errors.New("Token过期")
	ErrorWrongArgument    = errors.New("错误参数")
	ErrorAlreadyLiked     = errors.New("点赞失败，已经点赞过该视频了")
	ErrorNotLiked         = errors.New("取消失败，没有点赞过该视频")
	ErrorEncrypteFaild    = errors.New("密码加密失败")
	ErrorAlreadyFollowed  = errors.New("关注失败，已经关注过该用户了")
	ErrorNotFollowed      = errors.New("取消失败，没有关注过该用户")
)
