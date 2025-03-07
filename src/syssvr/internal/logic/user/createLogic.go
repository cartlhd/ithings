package userlogic

import (
	"context"
	"database/sql"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/internal/repo/mysql"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *CreateLogic) CheckPwd(in *sys.UserCreateReq) error {
	if l.svcCtx.Config.UserOpt.NeedPassWord &&
		utils.CheckPasswordLever(in.Password) < l.svcCtx.Config.UserOpt.PassLevel {
		return errors.PasswordLevel
	}
	return nil
}
func (l *CreateLogic) handlePassword(in *sys.UserCreateReq) (*sys.UserCreateResp, error) {
	//首先校验账号格式使用正则表达式，对用户账号做格式校验：只能是大小写字母，数字和下划线，减号
	ret := false
	if ret, _ = regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]{6,19}$", in.UserName); !ret {
		return nil, errors.UsernameFormatErr.AddDetail("账号必须以字母开头，且只能包含大小写字母和数字下划线和减号。 长度为6到20位之间")
	}
	//校验密码强度
	err := l.CheckPwd(in)
	if err != nil {
		return nil, err
	}

	//如果是账密，则in.Note为账号
	_, err = l.svcCtx.UserInfoModel.FindOneByUserName(l.ctx, sql.NullString{String: in.UserName, Valid: true})
	switch err {
	case nil: //已注册
		//提示重复注册
		return nil, errors.DuplicateRegister.AddDetail(in.UserName)
	case mysql.ErrNotFound: //未注册
		//1.生成uid
		uid_temp := l.svcCtx.UserID.GetSnowflakeId()

		//2.对密码进行md5加密
		password1 := utils.MakePwd(in.Password, uid_temp, false)
		ui := mysql.SysUserInfo{
			Uid:        uid_temp,
			UserName:   sql.NullString{String: in.UserName, Valid: true},
			Password:   password1,
			LastIP:     in.LastIP,
			RegIP:      in.RegIP,
			NickName:   in.NickName,
			City:       in.City,
			Country:    in.Country,
			Province:   in.Province,
			Language:   in.Language,
			HeadImgUrl: in.HeadImgUrl,
			Role:       in.Role,
			Sex:        in.Sex,
		}

		err := l.svcCtx.UserModel.Register(l.ctx, l.svcCtx.UserInfoModel, ui, mysql.Keys{Key: "userName", Value: in.UserName})
		if err != nil { //并发情况下有可能重复所以需要再次判断一次
			if err == mysql.ErrDuplicate {
				return nil, errors.DuplicateMobile.AddDetail(in.UserName)
			}
			l.Errorf("%s.Inserts err=%#v", utils.FuncName(), err)
			return nil, err
		}

		return &sys.UserCreateResp{Uid: ui.Uid}, nil
	default:
		break
	}

	return &sys.UserCreateResp{}, nil
}
func (l *CreateLogic) Create(in *sys.UserCreateReq) (*sys.UserCreateResp, error) {
	l.Infof("%s req=%+v", utils.FuncName(), in)
	switch in.ReqType {
	case "pwd":
		return l.handlePassword(in)
	default:
		l.Errorf("%s ReqType=%s  not support yet", utils.FuncName(), in.ReqType)
		return nil, errors.Parameter.AddDetail("reqType not support yet :" + in.ReqType)
	}
}
