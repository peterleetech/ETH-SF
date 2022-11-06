package open_api

import (
	"errors"
	"veric-backend-mvp/logic/db"
	"veric-backend-mvp/logic/http/http_util"
)

func GetUserVCList(r *http_util.HTTPContext) (resp interface{}, err error) {
	user, err := FindUser(r)
	if err != nil {
		return nil, err
	}

	return db.GetVCByUserIdAndPage(user.ID, r.QueryWithDefaultInt("page", 1), r.QueryWithDefaultInt("page_size", 15))
}

func DownloadUserVC(r *http_util.HTTPContext) (resp interface{}, err error) {
	user, err := FindUser(r)
	if err != nil {
		return nil, err
	}

	vcId := r.QueryWithDefaultInt("vc_id", -1)
	vc, err := db.GetVCById(uint(vcId))
	if err != nil {
		return nil, err
	}

	if vc.UserId != user.ID {
		return nil, errors.New("invalid user")
	}

	return vc.VC, nil
}
