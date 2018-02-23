package error

import "errors"

var (
	ErrUidNotExist        = errors.New("没有这个用户，user_id不存在")
	ErrAnotherUidNotExist = errors.New("没有这个用户，other_user_id不存在")
	ErrStateInvalid       = errors.New("参数错误，状态只能是liked或disliked")
	ErrNameEmpty          = errors.New("参数错误，用户名不能为空")
)
