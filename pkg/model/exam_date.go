package model

import (
	"time"
)

type ExamDate struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	ExamName     string     `json:"examName"`
	ExamDate     time.Time  `json:"examDate"`
	Duration     time.Duration
	DurationTime string
}

type ExamDateDTO struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	ExamName     string `json:"examName"`
	ExamDate     string `json:"examDate"`
	Duration     time.Duration
	DurationTime string
}

// ToExamDate converts ExamDateDTO to ExamDate
func ToExamDate(examDateDTO *ExamDateDTO) *ExamDate {
	// examDate, _ := time.Parse("2006-01-02", examDateDTO.ExamDate)
	examDate, err := time.Parse("2006-01-02", examDateDTO.ExamDate)
	if err != nil {
		return nil
	}
	return &ExamDate{
		ExamName:     examDateDTO.ExamName,
		ExamDate:     examDate,
		Duration:     examDateDTO.Duration,
		DurationTime: examDateDTO.DurationTime,
	}
}

// ToExamDateDTO converts ExamDate to ExamDateDTO
func ToExamDateDTO(examDate *ExamDate) *ExamDateDTO {
	return &ExamDateDTO{
		ID:           examDate.ID,
		ExamName:     examDate.ExamName,
		ExamDate:     examDate.ExamDate.Format("2006-01-02"),
		Duration:     examDate.Duration,
		DurationTime: examDate.DurationTime,
	}
}
