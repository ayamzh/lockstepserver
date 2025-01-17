package api

import (
	_ "embed"
	"fmt"
	"github.com/byebyebruce/lockstepserver/pkg/util"
	"html/template"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"

	"github.com/byebyebruce/lockstepserver/logic"
)

//go:embed index.html
var index string

// WebAPI http api
type WebAPI struct {
	m *logic.RoomManager
}

// NewWebAPI 构造
func NewWebAPI(port int, m *logic.RoomManager) *WebAPI {
	r := &WebAPI{
		m: m,
	}

	http.HandleFunc("/", r.index)
	http.HandleFunc("/create", r.createRoom)
	http.HandleFunc("/quit", r.quitRoom)

	go func() {
		fmt.Printf("web api listen on 127.0.0.1:%d \n", port)
		e := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		util.PanicIfErr(e)
	}()

	return r
}

func (h *WebAPI) index(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	if 0 == len(query) {
		t, err := template.New("test").Parse(index)
		if err != nil {
			w.Write([]byte("error"))
		} else {
			t.Execute(w, nil)
		}
		return
	}
}

func (h *WebAPI) createRoom(w http.ResponseWriter, r *http.Request) {

	ret := "error"

	defer func() {
		w.Write([]byte(ret))
	}()

	query := r.URL.Query()

	roomStr := query.Get("room")
	roomID, _ := strconv.ParseUint(roomStr, 10, 64)

	ps := make([]uint64, 0, 10)

	members := query.Get("member")
	if len(members) > 0 {

		a := strings.Split(members, ",")

		for _, v := range a {
			id, _ := strconv.ParseUint(v, 10, 64)
			ps = append(ps, id)
		}

	}

	room, err := h.m.CreateRoom(roomID, 0, ps, 0, "test")
	if nil != err {
		ret = err.Error()
	} else {
		ret = fmt.Sprintf("room.ID=[%d] room.Secret=[%s] room.Time=[%d], room.Member=[%v]", room.ID(), room.SecretKey(), room.TimeStamp(), members)
	}

}

func (h *WebAPI) quitRoom(w http.ResponseWriter, r *http.Request) {

	ret := "error"

	defer func() {
		w.Write([]byte(ret))
	}()

	query := r.URL.Query()

	roomStr := query.Get("room")
	roomID, _ := strconv.ParseUint(roomStr, 10, 64)

	playerStr := query.Get("player")
	playerID, _ := strconv.ParseUint(playerStr, 10, 64)

	room := h.m.GetRoom(roomID)
	if room == nil {
		ret = "room not exist"
		return
	}

	if !room.HasPlayer(playerID) {
		ret = "player not exist"
		return
	}

	ret = "player is leave room"
	return
}
