package models

type (
	GetDoctorSchedule struct {
		DoctorSlug   string `json:"doctor_slug"`
		DoctorName   string `json:"doctor_name"`
		ScheduleDay  string `json:"schedule_day"`
		ScheduleHour string `json:"schedule_hour"`
	}
)
