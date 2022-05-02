package contact

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
	"golang.org/x/xerrors"
)

type Contact interface {
	utils.DingIdReduceAble
	GetUserInfo(userid string) (res *GetUserInfoResponseResult, err error)
}

func NewContact(oauth2 oauth2.Oauth2) (r Contact) {
	return &inner{
		oauth2: oauth2,
	}
}

type inner struct {
	oauth2 oauth2.Oauth2
	utils.DingIdReduceStruct
}

func (d *inner) GetUserInfo(userid string) (res *GetUserInfoResponseResult, err error) {
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	body := make(map[string]interface{})
	ctx := context.Background()
	userid = d.Reduce(ctx, utils.AttrUserid, userid)
	body["userid"] = userid
	body["language"] = "zh_CN"
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRquest(http.MethodPost, "/topapi/v2/user/get", query, body)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var result GetUserInfoResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = xerrors.Errorf("%s", result.Errmsg)
		return
	}
	res = result.Result
	return
}
