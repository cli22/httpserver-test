package error

import (
	"google.golang.org/grpc/codes"
)

const (
	// ErrOk            成功
	ErrOk codes.Code = 0
	// ErrUnknown       未知错误
	ErrUnknown codes.Code = 101
	// ErrArgsInvalid   参数异常
	ErrArgsInvalid codes.Code = 102
	// ErrArgsEmpty     参数为空
	ErrArgsEmpty codes.Code = 103
	// ErrSystem        系统错误
	ErrSystem codes.Code = 104
	// ErrDB            数据库错误
	ErrDB codes.Code = 105
	// ErrNoServe       未提供服务
	ErrNoServe codes.Code = 106
	// ErrUidNotExist 没有这个用户，user_id不存在
	ErrUidNotExist codes.Code = 80001
	// ErrOtherUidNotExist 没有这个用户，other_user_id不存在
	ErrOtherUidNotExist codes.Code = 80002
	// ErrStateInvalid  状态只能是liked或disliked
	ErrStateInvalid codes.Code = 80003
	// ErrNameEmpty    用户名不能为空
	ErrNameEmpty codes.Code = 80004
	// ErrGetRelationship
	ErrGetRelationship codes.Code = 80005
	// ErrCreateRelationship
	ErrPutRelationship codes.Code = 80006
	// ErrListUser
	ErrGetUser codes.Code = 80007
	// ErrCreateUser
	ErrCreateUser codes.Code = 80008
)

// Msg 异常消息映射表
var Msg = map[codes.Code]string{
	ErrOk:               "成功",
	ErrUnknown:          "未知错误",
	ErrArgsInvalid:      "参数异常",
	ErrArgsEmpty:        "参数为空",
	ErrSystem:           "系统错误",
	ErrDB:               "数据库错误",
	ErrNoServe:          "未提供服务",
	ErrUidNotExist:      "没有这个用户，user_id不存在",
	ErrOtherUidNotExist: "没有这个用户，other_user_id不存在",
	ErrStateInvalid:     "状态只能是liked或disliked",
	ErrNameEmpty:        "用户名不能为空",
	ErrGetRelationship:  "获取用户关系错误",
	ErrPutRelationship:  "更新用户关系错误",
	ErrGetUser:          "获取所有用户错误",
	ErrCreateUser:       "创建用户错误",
}
