package jwt

import (
	"gin-api/pkg/jwt"
	"testing"
	"time"
)

const secret = "i1ydX9RtHyuJTrw7frcu"

func TestSign(t *testing.T) {
	token, err := jwt.New(secret).GenerateSign(1928374, 24*time.Hour)
	if err != nil {
		t.Error("generateSign error", err)
		return
	}
	t.Log(token)
}

func TestParse(t *testing.T)  {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjE5MjgzNzQsImV4cCI6MTYxNDc4MTIyNSwiaWF0IjoxNjE0Njk0ODI1LCJuYmYiOjE2MTQ2OTQ4MjV9.pdp5GqoomGIKgXM8OYkp9nIB58UenMd_jd2S_LyNV1k"
	claims, err := jwt.New(secret).ParseSign(token)
	if err != nil {
		t.Error("parseSign error", err)
		return
	}
	t.Log(claims.UserID)
}