package cabapi

import (
	cabsvc "github.com/STreeChin/cabtripapi/domain/service"
	"net/http"
)


func DeleteCacheCtrl(w http.ResponseWriter, r *http.Request){
	cabsvc.DeleteCache(w,r)
}