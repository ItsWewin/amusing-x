package password

import (
	"amusingx.fit/amusingx/protos/amusingxuserserv/userservice"
	"amusingx.fit/amusingx/services/amusinguserserv/rpcserver/userservice/handler/passwordapp"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandlerResetPassword(ctx context.Context,
	req *userservice.ResetPasswordRequest) (*userservice.ResetPasswordResponse, *xerror.Error) {

	err := getAndValidRequest(req)
	if err != nil {
		return nil, err
	}

	err = resetPassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return &userservice.ResetPasswordResponse{Result: "密码重置成功"}, nil
}

func resetPassword(ctx context.Context, req *userservice.ResetPasswordRequest) *xerror.Error {
	domain := passwordapp.NewDomain()

	err := domain.SetResetPasswordInfo(req)
	if err != nil {
		return xerror.NewError(err, err.Code, "SetResetPasswordInfo failed. ")
	}

	err = domain.SetUserModelInfo(ctx)
	if err != nil {
		return xerror.NewError(err, err.Code, "SetUserModelInfo failed. ")
	}

	err = domain.ResetPassword(ctx, domain.ResetPasswordInfo.Password)
	if err != nil {
		return xerror.NewError(err, err.Code, "ResetPassword failed. ")
	}

	return nil
}

func getAndValidRequest(request *userservice.ResetPasswordRequest) *xerror.Error {
	xErr := request.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}
