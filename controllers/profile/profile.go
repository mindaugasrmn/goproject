package mecontroller

import (
	"net/http"
	"encoding/json"
	"github.com/mindaugasrmn/goproject/apicore/db/mongo"
	"github.com/mindaugasrmn/goproject/apicore/utils"
	"fmt"
	"github.com/gorilla/mux"
)


func GetMyProfileController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	username := vars["username"]
    db := mongo.GetDb(r)
    user := utils.GetProfileByUsername(db,username)
    data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", data)
	
}
