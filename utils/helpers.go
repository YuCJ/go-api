package utils

import "database/sql"

// ToNullString invalidates a sql.NullString if empty, validates if not empty
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// GetNullString returns a invalid NullString
func GetNullString() sql.NullString {
	return sql.NullString{String: "", Valid: false}
}