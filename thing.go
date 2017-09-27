package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	uuid "github.com/satori/go.uuid"
	render "github.com/unrolled/render"
)

var ren = render.New()

// Thing holds all logs
type Thing struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Action string    `json:"action"`
	Time   time.Time `json:"time"`
	Logged time.Time `json:"logged"`
}

// FieldMap function for json body binding validation
func (th *Thing) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&th.Action: binding.Field{
			Form:     "action",
			Required: true,
		},
		&th.Time: binding.Field{
			Form:     "time",
			Required: true,
		},
	}
}

func logThing(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID, errs := uuid.FromString(vars["userid"])
	if errs != nil {
		http.Error(w, "Invalid UUID for User ID", http.StatusUnprocessableEntity)
		return
	}

	thing := new(Thing)
	thing.ID = uuid.NewV4()
	thing.UserID = userID
	thing.Logged = time.Now()

	if errs := binding.Bind(req, thing); errs != nil {
		http.Error(w, errs.Error(), http.StatusUnprocessableEntity)
		return
	}

	ren.JSON(w, http.StatusOK, thing)
}
