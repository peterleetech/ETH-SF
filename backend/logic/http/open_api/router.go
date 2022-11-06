package open_api

import (
	"github.com/gorilla/mux"
	"reflect"
	"veric-backend-mvp/logic/http/http_util"
)

func RegisterRouter(r *mux.Router) {
	r.Handle("/airswift_callback", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(AirSwiftCallbackRequest{}), AirSwiftCallback),
	})
	r.Handle("/user/login", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(LoginUseEthSignatureRequest{}), LoginUseEthSignature),
	})
	r.Handle("/user/vc", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetUserVCList),
	})
	r.Handle("/user/vc/download", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(DownloadUserVC),
	})
	r.Handle("/user/vp/withdraw", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(VPWithdrawRequest{}), VPWithdraw),
	})
	r.Handle("/open/transaction_list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(TransactionList),
	})
}
