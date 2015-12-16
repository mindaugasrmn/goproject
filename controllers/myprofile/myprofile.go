package mecontroller

import (
	"net/http"
	"encoding/json"
	"github.com/mindaugasrmn/goproject/apicore/db/mongo"
    "github.com/mindaugasrmn/goproject/apicore/services/auth"
	"github.com/mindaugasrmn/goproject/apicore/utils"
	"fmt"
)

type omit bool

func GetMyProfileController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    db := mongo.GetDb(r)
    user := utils.GetUserByIdMe(db,services.GetUserFromToken(r))
    db.FindRef(user.Position).One(&user.Pos)
    data, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", data)
	
}
