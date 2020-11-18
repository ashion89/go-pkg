/**
 * @Author pibing
 * @create 2020/11/14 1:26 PM
 */

package svc

import (
	"encoding/json"
	"go-pkg/model"
	"go-pkg/params"
	"go-pkg/pkg/db"
	"go-pkg/pkg/redis"
	"go-pkg/util"
	"gorm.io/gorm"
	"time"
)

// 用户登陆
func SignIn(param *params.SigninReq) (rsp *params.SigninRsp, err error) {
	var user model.User
	err = db.GetDB().Model(&model.User{}).Where(nil).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	if user.Password != param.Password {
		return nil, err
	}

	// 生成token
	rsp = &params.SigninRsp{}
	rsp.UID = user.ID
	rsp.Token = util.GenToken()

	ut := util.UserToken{
		Name:       user.Name,
		UID:        user.ID,
		Token:      rsp.Token,
		ExpireTime: time.Now().Unix() + util.TokenExpireTime,
	}

	tokenData, err := json.Marshal(ut)
	if err != nil {
		return nil, err
	}

	// 把以前的token干掉
	oldToken, _ := redis.Get(util.FormatUserTokenKey(rsp.UID))
	if oldToken != nil {
		redis.Delete(util.FormatTokenUserKey(string(oldToken)))
	}

	// 记录uid和token的互相关联关系
	redis.Set(util.FormatUserTokenKey(rsp.UID), rsp.Token)
	redis.Set(util.FormatTokenUserKey(rsp.Token), tokenData)

	// todo 把用户信息写到缓存，可提高访问效率

	return rsp, nil
}


// 用户登出
func SignOut(token string) error {
	tokenData, err := redis.Get(util.FormatTokenUserKey(token))
	if err != nil {
		return err
	}
	redis.Delete(util.FormatTokenUserKey(token))

	ut := util.UserToken{}
	err = json.Unmarshal(tokenData, &ut)
	if err != nil {
		return err
	}
	redis.Delete(util.FormatUserTokenKey(ut.UID))

	return nil
}
