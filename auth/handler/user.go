// Code generated by protoc-gen-vkit.
// versions:
// - protoc-gen-vkit v1.0.0

package handler

import (
	context "context"

	"github.com/vison888/iot-engine/auth/errno"
	pb "github.com/vison888/iot-engine/auth/proto"
	"github.com/vison888/iot-engine/auth/utils"
	"gorm.io/gorm"

	"github.com/vison888/iot-engine/auth/model"

	"github.com/vison888/iot-engine/auth/handler/auth"
	"github.com/vison888/iot-engine/auth/handler/convert"
	"github.com/vison888/iot-engine/auth/handler/user"

	"github.com/vison888/go-vkit/errorsx"
)

type UserService struct {
}

func (the *UserService) Add(ctx context.Context, req *pb.UserAddReq, resp *pb.UserAddResp) error {
	if len(req.Item.RoleCode) == 0 {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "必须选择角色"
		return nil
	}
	if len(req.Item.Account) == 0 && len(req.Item.Email) == 0 && len(req.Item.Phone) == 0 {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "账号、邮箱或手机号至少有一个不为空"
		return nil
	}
	// 账号 邮箱 手机唯一
	if req.Item.Account != "" {
		_, err := model.UserGetByAccount(nil, req.Item.Account)
		if err != nil && err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		if err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = "账号已经被注册"
			return nil
		}
	}

	if req.Item.Phone != "" {
		_, err := model.UserGetByPhone(nil, req.Item.Phone)
		if err != nil && err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		if err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = "电话号码已经被注册"
			return nil
		}
	}

	if req.Item.Email != "" {
		_, err := model.UserGetByEmail(nil, req.Item.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		if err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = "邮箱已经被注册"
			return nil
		}
	}

	u := &model.UserModel{
		Account:    req.Item.Account,
		Password:   user.EncodePassword(req.Item.Password),
		NickName:   req.Item.NickName,
		Email:      req.Item.Email,
		Phone:      req.Item.Phone,
		CreateUser: utils.GetUserIdFromContext(ctx),
	}
	roles, err := model.RoleListByCodes(nil, req.Item.RoleCode)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	err = model.UserAddWithTransaction(nil, u, roles)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	resp.Code = errorsx.OK.Code
	resp.Id = u.Id
	return nil
}

func (the *UserService) Page(ctx context.Context, req *pb.UserPageReq, resp *pb.UserPageResp) error {
	if len(req.RoleCodes) == 0 && len(req.AppCode) > 0 {
		//获取应用对应的所有角色
		req.RoleCodes = auth.RoleCodesByAppCode(req.AppCode)
	}
	list, total, err := model.UserPage(nil, req.PageIndex, req.PageSize, req.OrderKey, req.Desc,
		req.Account, req.NickName, req.Phone, req.Email, req.CreateUser, req.RoleCodes)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	resp.List = []*pb.User{}
	for _, user := range list {
		pbUser, err := convert.UserToUserPb(user, req.AppCode)
		if err != nil {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		resp.List = append(resp.List, pbUser)
	}
	resp.Code = errorsx.OK.Code
	resp.Total = total
	return nil
}

func (the *UserService) Update(ctx context.Context, req *pb.UserUpdateReq, resp *pb.UserUpdateResp) error {
	u, err := model.UserGetFromCache(nil, req.Item.Id)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "用户信息获取失败" + err.Error()
		return nil
	}
	if len(req.Item.Password) != 0 {
		u.Password = user.EncodePassword(req.Item.Password)
	}
	u.Account = req.Item.Account
	u.NickName = req.Item.NickName
	u.Email = req.Item.Email
	u.Phone = req.Item.Phone
	u.UpdateUser = utils.GetUserIdFromContext(ctx)
	if len(req.Item.Account) == 0 && len(req.Item.Email) == 0 && len(req.Item.Phone) == 0 {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "账号、邮箱或手机号至少有一个不为空"
		return nil
	}

	// 账号 邮箱 手机唯一
	if req.Item.Account != "" {
		m, err := model.UserGetByAccount(nil, req.Item.Account)
		if err != nil && err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		if m.Id != "" && m.Id != req.Item.Id {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = "账号已经被注册"
			return nil
		}
	}

	if req.Item.Phone != "" {
		m, err := model.UserGetByPhone(nil, req.Item.Phone)
		if err != nil && err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		if m.Id != "" && m.Id != req.Item.Id {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = "电话号码已经被注册"
			return nil
		}
	}

	if req.Item.Email != "" {
		m, err := model.UserGetByEmail(nil, req.Item.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		if m.Id != "" && m.Id != req.Item.Id {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = "邮箱已经被注册"
			return nil
		}
	}

	err = model.UserUpdateAndRole(nil, u, req.Item.RoleCode)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	resp.Code = errorsx.OK.Code
	return nil
}

func (the *UserService) Del(ctx context.Context, req *pb.UserDelReq, resp *pb.UserDelResp) error {
	if req.AppCode == "" {
		codes, err := model.RoleCodeListByUserId(nil, req.Id)
		if err != nil {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		for _, code := range codes {
			if code == "LEADER" || code == "MEMBER" {
				resp.Code = errorsx.FAIL.Code
				resp.Msg = "因队长和队员角色关联其他业务数据，请勿删除！"
				return nil
			}
		}
		err = model.UserDel(nil, req.Id)
		if err != nil {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
	} else {
		roles := auth.RoleCodesByAppCode(req.AppCode)
		err := model.UserRoleBatchDelByRoleCode(nil, req.Id, roles)
		if err != nil {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
	}
	resp.Code = errorsx.OK.Code
	return nil
}

func (the *UserService) Get(ctx context.Context, req *pb.UserGetReq, resp *pb.UserGetResp) error {
	u, err := model.UserGetFromCache(nil, req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.Code = errno.USER_NOT_FOUND_ERR_NO.Code
			resp.Msg = errno.USER_NOT_FOUND_ERR_NO.Msg
			return nil
		}
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}

	resp.Item, err = convert.UserToUserPb(u, req.AppCode)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	resp.Code = errorsx.OK.Code
	return nil
}

func (the *UserService) PageBySql(ctx context.Context, req *pb.UserPageBySqlReq, resp *pb.UserPageBySqlResp) error {
	list, total, err := model.UserPageBySql(nil, req.Filter, req.PageIndex, req.PageSize, req.OrderKey, req.Desc)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	resp.List = []*pb.User{}
	for _, user := range list {
		pbUser, err := convert.UserToUserPb(user, req.AppCode)
		if err != nil {
			resp.Code = errorsx.FAIL.Code
			resp.Msg = err.Error()
			return nil
		}
		resp.List = append(resp.List, pbUser)
	}
	resp.Code = errorsx.OK.Code
	resp.Total = total
	return nil
}

func (the *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq, resp *pb.UpdatePasswordResp) error {
	u, err := model.UserGet(nil, req.UserId)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	if user.EncodePassword(req.CurPassword) != u.Password {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "密码错误"
		return nil
	}
	u.Password = user.EncodePassword(req.NewPassword)
	_, err = model.UserUpdate(nil, u)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	resp.Code = errorsx.OK.Code
	return nil
}

func (the *UserService) LinkUserRoles(ctx context.Context, req *pb.LinkUserRolesReq, resp *pb.Response) error {
	if len(req.RoleCodes) == 0 {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "角色列表为空"
		return nil
	}
	err := model.UserRoleBatchAddByRoleCode(nil, req.UserId, req.RoleCodes)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	return nil
}

func (the *UserService) UnlinkUserRoles(ctx context.Context, req *pb.UnlinkUserRolesReq, resp *pb.Response) error {
	if len(req.RoleCodes) == 0 {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = "角色列表为空"
		return nil
	}
	err := model.UserRoleBatchDelByRoleCode(nil, req.UserId, req.RoleCodes)
	if err != nil {
		resp.Code = errorsx.FAIL.Code
		resp.Msg = err.Error()
		return nil
	}
	return nil
}
