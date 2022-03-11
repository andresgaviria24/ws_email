package utils

import (
	"net/http"
	"regexp"
	"strconv"
	"time"
	"ws_notifications_email/domain/dto"
)

func StatusText(code int) string {
	return http.StatusText(code)
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func ResponseError(code int, err error) dto.Response {
	if err != nil {
		return dto.Response{
			Status:      code,
			Description: StatusText(code),
			Message:     err.Error(),
		}
	}
	return dto.Response{Status: http.StatusOK}
}

func ConvertInt64(number string) int64 {
	if number != "" {
		value, err := strconv.Atoi(number)
		if err != nil {
			return 0
		}
		return int64(value)
	}
	return 0
}

func TimeNow() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)
	return t.Format("2006-01-02 15:04:05")
}
