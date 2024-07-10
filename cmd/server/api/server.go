package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/klaital/todolist/internal/storage/queries"
	"log/slog"
	"net/http"
	"time"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	q *queries.Queries
}

func NewServer(q *queries.Queries) Server {
	return Server{
		q: q,
	}
}

func pointerTo[T any](s T) *T {
	return &s
}

// GET /lists
func (s *Server) GetLists(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) GetChores(w http.ResponseWriter, r *http.Request) {
	var userId int64 = 1 // hardcoded until login workflow is implemented
	chores, err := s.q.ListChores(r.Context(), int32(userId))
	if err != nil {
		// TODO: discern between "not found" and actual db errors
		slog.Error("Failed to query chores", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	slog.Debug("Got chores from db", "chores", len(chores))

	resp := ListChoresResponse{
		Chores: make([]Chore, len(chores)),
	}
	for i, c := range chores {
		resp.Chores[i] = Chore{
			Id:            c.ID,
			LastCompleted: c.LastCompleted.UnixMilli(),
			Name:          c.Name,
			NextDeadline:  c.NextDeadline.UnixMilli(),
		}
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal listchoresresponse", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func (s *Server) MarkChoreCompleted(w http.ResponseWriter, r *http.Request, choreId int64) {
	err := s.q.UpdateChoreCompleted(r.Context(), queries.UpdateChoreCompletedParams{
		LastCompleted: time.Now(),
		NextDeadline:  time.Now().AddDate(0, 0, 7),
		ID:            choreId,
	})
	if err != nil {
		slog.Error("Failed to update chore", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the updated Chore back to the caller so they know the new deadline
	c, err := s.q.GetChore(r.Context(), choreId)
	if err != nil {
		slog.Error("Failed to refetch chore", "choreId", choreId, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Success!
	respBytes, err := json.Marshal(c)
	if err != nil {
		slog.Error("Failed to marshal response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(respBytes)
}
