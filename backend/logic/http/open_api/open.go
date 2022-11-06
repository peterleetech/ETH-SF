package open_api

import (
	"veric-backend-mvp/logic/db"
	"veric-backend-mvp/logic/http/http_util"
)

func TransactionList(r *http_util.HTTPContext) (resp interface{}, err error) {
	return db.GetEventByPage(r.QueryWithDefaultInt("page", 1), r.QueryWithDefaultInt("page_size", 15))
}
