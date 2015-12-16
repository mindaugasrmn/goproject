package clients

import (
	"net/http"
	"encoding/json"
	"github.com/mindaugasrmn/goproject/core/mongo"
    "github.com/mindaugasrmn/goproject/services"
	"github.com/mindaugasrmn/goproject/utils"
	"fmt"
	"log"
)

func GetMeController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    db := mongo.GetDb(r)
    user := utils.GetUserByIdMe(db,services.GetUserFromToken(r))
    db.FindRef(user.Position).One(&user.Pos)
    log.Print(user.Pos)
    data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", data)
	
}