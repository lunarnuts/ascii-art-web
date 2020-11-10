package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	ascii "../ascii"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var data input
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ErrorHandler(400, w, r)
		return
	}
	out, e := ascii.AsciiArt([]string{data.Text, data.Banner})
	if e != 0 {
		ErrorHandler(500, w, r)
		return
	}
	fmt.Println(out)
	ExportFile(out)
	w.WriteHeader(200)
	var result = make(map[string]string)
	result["result"] = out
	json.NewEncoder(w).Encode(result)
}
