package models

// Stats struct
type Stats struct {
	TotalAssigns        int64   `json:"total_assigns" db:"total"`
	TotalSuccessAssigns int64   `json:"total_success_assigns"`
	TotalFailAssigns    int64   `json:"total_fail_assigns"`
	AVGSuccessAssigns   float32 `json:"avg_success_assigns"`
	AVGFailAssigns      float32 `json:"avg_fail_assigns"`
}
