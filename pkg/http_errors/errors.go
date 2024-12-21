package http_errors

import (
	"encoding/json"
	"net/http"
)

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func BadRequest(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "https://example.com/probs/bad-request",
		Title:    "Bad Request",
		Status:   http.StatusBadRequest,
		Detail:   detail,
		Instance: instance,
	}
}

func Unauthorized(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "https://example.com/probs/unauthorized",
		Title:    "Unauthorized",
		Status:   http.StatusUnauthorized,
		Detail:   detail,
		Instance: instance,
	}
}

func Forbidden(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "https://example.com/probs/forbidden",
		Title:    "Forbidden",
		Status:   http.StatusForbidden,
		Detail:   detail,
		Instance: instance,
	}
}

func NotFound(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "https://example.com/probs/not-found",
		Title:    "Not Found",
		Status:   http.StatusNotFound,
		Detail:   detail,
		Instance: instance,
	}
}

func InternalServerError(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "https://example.com/probs/internal-server-error",
		Title:    "Internal Server Error",
		Status:   http.StatusInternalServerError,
		Detail:   detail,
		Instance: instance,
	}
}

func ServiceUnavailable(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "https://example.com/probs/service-unavailable",
		Title:    "Service Unavailable",
		Status:   http.StatusServiceUnavailable,
		Detail:   detail,
		Instance: instance,
	}
}

func WriteProblemDetails(w http.ResponseWriter, problem ProblemDetails) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(problem.Status)
	err := json.NewEncoder(w).Encode(problem)
	if err != nil {
		return
	}
}
