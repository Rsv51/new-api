// api/index.go
package api

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 示例：返回查询结果
	resp := map[string]interface{}{
		"code": 200,
		"msg":  "Hello from NewAPI-Go on Vercel!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
