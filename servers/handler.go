package servers

import (
	"fmt"
	"net/http"

	"github.com/HyeockJinKim/jira-transform/destination"
	"github.com/labstack/echo/v4"
)

type HandlerArgs struct {
	Destinations map[string]destination.Destination
}

type Handler struct {
	destinations map[string]destination.Destination
}

func NewHandler(args HandlerArgs) *Handler {
	return &Handler{
		destinations: args.Destinations,
	}
}

func (h *Handler) Mirror(c echo.Context) error {
	req := new(JIRAMirrorRequest)
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}
	dest, ok := h.destinations[req.Destination]
	if !ok {
		return fmt.Errorf("unsupported destination: %s", req.Destination)
	}
	if err := dest.Mirror(c.Request().Context(), req.ToIssue()); err != nil {
		return fmt.Errorf("failed to mirror: %w", err)
	}
	return c.NoContent(http.StatusOK)
}
