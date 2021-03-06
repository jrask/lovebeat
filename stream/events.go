package stream

import (
	"encoding/json"
	"github.com/boivie/lovebeat/service"
)

type serviceChangedArgs struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type serviceChangedMsg struct {
	Message string             `json:"m"`
	Args    serviceChangedArgs `json:"args"`
}

func serviceStateChanged(ev service.ServiceStateChangedEvent) {
	msg := serviceChangedMsg{
		Message: "service_changed",
		Args: serviceChangedArgs{
			Name:  ev.Service.Name,
			State: ev.Current,
		},
	}

	var encoded, _ = json.Marshal(msg)
	h.broadcast <- encoded
}

type viewChangedArgs struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type viewChangedMsg struct {
	Message string          `json:"m"`
	Args    viewChangedArgs `json:"args"`
}

func viewStateChanged(ev service.ViewStateChangedEvent) {
	msg := viewChangedMsg{
		Message: "view_changed",
		Args: viewChangedArgs{
			Name:  ev.View.Name,
			State: ev.Current,
		},
	}

	var encoded, _ = json.Marshal(msg)
	h.broadcast <- encoded
}
