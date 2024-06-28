package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func pointerTo[T any](s T) *T {
	return &s
}

// GET /lists
func (s Server) GetLists(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.NewV6()
	resp := GetListsResponse{
		Lists: &[]ListDescription{
			{
				Id:   &id,
				Name: pointerTo("grocery list"),
			},
		},
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error("Failed to serialize response: %s", err.Error())
	}
}
