package sql

import (
	"database/sql"
	"time"
)

func ToPtrTime(value sql.NullTime) *time.Time {
	var time *time.Time
	if value.Valid {
		time = &value.Time
	}

	return time
}

func ToPtrString(nullableString sql.NullString) *string {
	var value *string
	if nullableString.Valid {
		value = &nullableString.String
	}

	return value
}

func ToNullableString(value *string) sql.NullString {
	if value != nil {
		return sql.NullString{Valid: true, String: *value}
	}

	return sql.NullString{}
}

func ToNullableTime(value *time.Time) sql.NullTime {
	if value != nil {
		return sql.NullTime{Valid: true, Time: *value}
	}

	return sql.NullTime{}
}
