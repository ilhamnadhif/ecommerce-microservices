package dto

import (
	pb "api-gateway/proto"
	"fmt"
	"time"
)

type QueryData struct {
	ID   int64
	Role pb.Role
}

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}
