package errno

import baseNo "github.com/vison888/go-vkit/errorsx"

var (
	USER_NOT_FOUND_ERR_NO = baseNo.NewErrno(5, 1, "用户不存在")
)
