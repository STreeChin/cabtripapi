package cabsvc

import (
	"github.com/STreeChin/cabtripapi/domain/model"
	"net/http"
)

func DeleteCache(w http.ResponseWriter, r *http.Request) {
	cache := model.GetCacheInstance()
	cache.ResetCache()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
