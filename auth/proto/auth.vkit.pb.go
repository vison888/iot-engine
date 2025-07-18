// Code generated by protoc-gen-vkit. DO NOT EDIT.
// versions:
// - protoc-gen-vkit v1.0.0
// - protoc             v3.21.1
// source: auth.proto

package proto

import (
	context "context"

	grpcx "github.com/vison888/go-vkit/grpcx"
	grpc "google.golang.org/grpc"
)

var _ = new(context.Context)
var _ = new(grpc.CallOption)
var _ = new(grpcx.Client)

type AuthServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *AuthServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) LoginByPhone(ctx context.Context, in *LoginByPhoneReq, opts ...grpc.CallOption) (*LoginByPhoneResp, error) {
	out := new(LoginByPhoneResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.LoginByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) LoginByEmail(ctx context.Context, in *LoginByEmailReq, opts ...grpc.CallOption) (*LoginByEmailResp, error) {
	out := new(LoginByEmailResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.LoginByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	out := new(LogoutResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) VerificationCode(ctx context.Context, in *VerificationCodeReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, c.name, "AuthService.VerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) ForgetPassword(ctx context.Context, in *ForgetPasswordReq, opts ...grpc.CallOption) (*ForgetPasswordResp, error) {
	out := new(ForgetPasswordResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.ForgetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordResp, error) {
	out := new(VerifyPasswordResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.VerifyPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) GenToken(ctx context.Context, in *GenTokenReq, opts ...grpc.CallOption) (*GenTokenResp, error) {
	out := new(GenTokenResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.GenToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) Inspect(ctx context.Context, in *InspectReq, opts ...grpc.CallOption) (*InspectResp, error) {
	out := new(InspectResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.Inspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	out := new(RefreshTokenResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) DelToken(ctx context.Context, in *DelTokenReq, opts ...grpc.CallOption) (*DelTokenResp, error) {
	out := new(DelTokenResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.DelToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AuthServiceClient) APIPermissions(ctx context.Context, in *APIPermissionsReq, opts ...grpc.CallOption) (*APIPermissionsResp, error) {
	out := new(APIPermissionsResp)
	err := c.cc.Invoke(ctx, c.name, "AuthService.APIPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewAuthServiceClient(name string, cc grpcx.Client) *AuthServiceClient {
	return &AuthServiceClient{name, cc}
}

type AppServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *AppServiceClient) Page(ctx context.Context, in *AppPageReq, opts ...grpc.CallOption) (*AppPageResp, error) {
	out := new(AppPageResp)
	err := c.cc.Invoke(ctx, c.name, "AppService.Page", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AppServiceClient) Add(ctx context.Context, in *AppAddReq, opts ...grpc.CallOption) (*AppAddResp, error) {
	out := new(AppAddResp)
	err := c.cc.Invoke(ctx, c.name, "AppService.Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AppServiceClient) Del(ctx context.Context, in *AppDelReq, opts ...grpc.CallOption) (*AppDelResp, error) {
	out := new(AppDelResp)
	err := c.cc.Invoke(ctx, c.name, "AppService.Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AppServiceClient) Update(ctx context.Context, in *AppUpdateReq, opts ...grpc.CallOption) (*AppUpdateResp, error) {
	out := new(AppUpdateResp)
	err := c.cc.Invoke(ctx, c.name, "AppService.Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *AppServiceClient) Get(ctx context.Context, in *AppGetReq, opts ...grpc.CallOption) (*AppGetResp, error) {
	out := new(AppGetResp)
	err := c.cc.Invoke(ctx, c.name, "AppService.Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewAppServiceClient(name string, cc grpcx.Client) *AppServiceClient {
	return &AppServiceClient{name, cc}
}

type RoleServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *RoleServiceClient) List(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	out := new(RoleListResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RoleServiceClient) ListByCode(ctx context.Context, in *RoleListByCodeReq, opts ...grpc.CallOption) (*RoleListByCodeResp, error) {
	out := new(RoleListByCodeResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.ListByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RoleServiceClient) Add(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	out := new(RoleAddResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RoleServiceClient) Del(ctx context.Context, in *RoleDelReq, opts ...grpc.CallOption) (*RoleDelResp, error) {
	out := new(RoleDelResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RoleServiceClient) Update(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	out := new(RoleUpdateResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RoleServiceClient) Get(ctx context.Context, in *RoleGetReq, opts ...grpc.CallOption) (*RoleGetResp, error) {
	out := new(RoleGetResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RoleServiceClient) Page(ctx context.Context, in *RolePageReq, opts ...grpc.CallOption) (*RolePageResp, error) {
	out := new(RolePageResp)
	err := c.cc.Invoke(ctx, c.name, "RoleService.Page", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewRoleServiceClient(name string, cc grpcx.Client) *RoleServiceClient {
	return &RoleServiceClient{name, cc}
}

type UserServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *UserServiceClient) Add(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	out := new(UserAddResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) Page(ctx context.Context, in *UserPageReq, opts ...grpc.CallOption) (*UserPageResp, error) {
	out := new(UserPageResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.Page", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) PageBySql(ctx context.Context, in *UserPageBySqlReq, opts ...grpc.CallOption) (*UserPageBySqlResp, error) {
	out := new(UserPageBySqlResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.PageBySql", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) Update(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	out := new(UserUpdateResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) Del(ctx context.Context, in *UserDelReq, opts ...grpc.CallOption) (*UserDelResp, error) {
	out := new(UserDelResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) Get(ctx context.Context, in *UserGetReq, opts ...grpc.CallOption) (*UserGetResp, error) {
	out := new(UserGetResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := c.cc.Invoke(ctx, c.name, "UserService.UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) LinkUserRoles(ctx context.Context, in *LinkUserRolesReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, c.name, "UserService.LinkUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *UserServiceClient) UnlinkUserRoles(ctx context.Context, in *UnlinkUserRolesReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, c.name, "UserService.UnlinkUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewUserServiceClient(name string, cc grpcx.Client) *UserServiceClient {
	return &UserServiceClient{name, cc}
}

type ResourceServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *ResourceServiceClient) Add(ctx context.Context, in *AddResourceReq, opts ...grpc.CallOption) (*AddResourceResp, error) {
	out := new(AddResourceResp)
	err := c.cc.Invoke(ctx, c.name, "ResourceService.Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ResourceServiceClient) Update(ctx context.Context, in *UpdateResourceReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, c.name, "ResourceService.Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ResourceServiceClient) Del(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, c.name, "ResourceService.Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ResourceServiceClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResourceGetResp, error) {
	out := new(ResourceGetResp)
	err := c.cc.Invoke(ctx, c.name, "ResourceService.Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ResourceServiceClient) Page(ctx context.Context, in *ResourcePageReq, opts ...grpc.CallOption) (*ResourcePageResp, error) {
	out := new(ResourcePageResp)
	err := c.cc.Invoke(ctx, c.name, "ResourceService.Page", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ResourceServiceClient) RuleType(ctx context.Context, in *ResourceRuleTypeReq, opts ...grpc.CallOption) (*ResourceRuleTypeResp, error) {
	out := new(ResourceRuleTypeResp)
	err := c.cc.Invoke(ctx, c.name, "ResourceService.RuleType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewResourceServiceClient(name string, cc grpcx.Client) *ResourceServiceClient {
	return &ResourceServiceClient{name, cc}
}

type ApiWhiteListServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *ApiWhiteListServiceClient) Refresh(ctx context.Context, in *ApiWhiteListRefreshReq, opts ...grpc.CallOption) (*ApiWhiteListRefreshResp, error) {
	out := new(ApiWhiteListRefreshResp)
	err := c.cc.Invoke(ctx, c.name, "ApiWhiteListService.Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ApiWhiteListServiceClient) Add(ctx context.Context, in *ApiWhiteListAddReq, opts ...grpc.CallOption) (*ApiWhiteListAddResp, error) {
	out := new(ApiWhiteListAddResp)
	err := c.cc.Invoke(ctx, c.name, "ApiWhiteListService.Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ApiWhiteListServiceClient) Update(ctx context.Context, in *ApiWhiteListUpdateReq, opts ...grpc.CallOption) (*ApiWhiteListUpdateResp, error) {
	out := new(ApiWhiteListUpdateResp)
	err := c.cc.Invoke(ctx, c.name, "ApiWhiteListService.Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ApiWhiteListServiceClient) Del(ctx context.Context, in *ApiWhiteListDelReq, opts ...grpc.CallOption) (*ApiWhiteListDelResp, error) {
	out := new(ApiWhiteListDelResp)
	err := c.cc.Invoke(ctx, c.name, "ApiWhiteListService.Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ApiWhiteListServiceClient) Page(ctx context.Context, in *ApiWhiteListPageReq, opts ...grpc.CallOption) (*ApiWhiteListPageResp, error) {
	out := new(ApiWhiteListPageResp)
	err := c.cc.Invoke(ctx, c.name, "ApiWhiteListService.Page", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewApiWhiteListServiceClient(name string, cc grpcx.Client) *ApiWhiteListServiceClient {
	return &ApiWhiteListServiceClient{name, cc}
}

type PermissionServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *PermissionServiceClient) Page(ctx context.Context, in *PermissionPageReq, opts ...grpc.CallOption) (*PermissionPageResp, error) {
	out := new(PermissionPageResp)
	err := c.cc.Invoke(ctx, c.name, "PermissionService.Page", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *PermissionServiceClient) Add(ctx context.Context, in *PermissionAddReq, opts ...grpc.CallOption) (*PermissionAddResp, error) {
	out := new(PermissionAddResp)
	err := c.cc.Invoke(ctx, c.name, "PermissionService.Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *PermissionServiceClient) Update(ctx context.Context, in *PermissionUpdateReq, opts ...grpc.CallOption) (*PermissionUpdateResp, error) {
	out := new(PermissionUpdateResp)
	err := c.cc.Invoke(ctx, c.name, "PermissionService.Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *PermissionServiceClient) Del(ctx context.Context, in *PermissionDelReq, opts ...grpc.CallOption) (*PermissionDelResp, error) {
	out := new(PermissionDelResp)
	err := c.cc.Invoke(ctx, c.name, "PermissionService.Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *PermissionServiceClient) Get(ctx context.Context, in *PermissionGetReq, opts ...grpc.CallOption) (*PermissionGetResp, error) {
	out := new(PermissionGetResp)
	err := c.cc.Invoke(ctx, c.name, "PermissionService.Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *PermissionServiceClient) GetMenu(ctx context.Context, in *PermissionGetMenuReq, opts ...grpc.CallOption) (*PermissionGetMenuResp, error) {
	out := new(PermissionGetMenuResp)
	err := c.cc.Invoke(ctx, c.name, "PermissionService.GetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewPermissionServiceClient(name string, cc grpcx.Client) *PermissionServiceClient {
	return &PermissionServiceClient{name, cc}
}
