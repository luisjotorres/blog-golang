package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type PayloadForTokenVisiting struct {
	IpAddress          string    `json:"ip_address"`
	TimeFromStartVisit time.Time `json:"time_from_start_visit"`
	Type               string    `json:"type"`
	jwt.RegisteredClaims
}
