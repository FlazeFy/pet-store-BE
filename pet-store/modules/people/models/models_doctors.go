package models

type (
	GetDoctorSchedule struct {
		DoctorSlug   string `json:"doctors_slug"`
		DoctorName   string `json:"doctors_name"`
		ScheduleDay  string `json:"schedule_day"`
		ScheduleHour string `json:"schedule_hour"`
	}
)