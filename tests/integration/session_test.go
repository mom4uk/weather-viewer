package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"weather-viewer/internal/domain"
	"weather-viewer/tests/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Session

func TestSession_error_sessionExpired(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	ctx := context.Background()

	err = db.Redis.Del(ctx, sessionID).Err()
	require.NoError(t, err)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/getLocations",
		nil,
		sessionID,
	)

	testutils.AssertStatus(t, rr, http.StatusUnauthorized)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Unauthorized",
	}

	assert.Equal(t, expected, got)
}
