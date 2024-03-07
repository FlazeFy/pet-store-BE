package models

type (
	GetDoctorSchedule struct {
		DoctorSlug   string `json:"doctor_slug"`
		DoctorName   string `json:"doctor_name"`
		ScheduleDay  string `json:"schedule_day"`
		ScheduleHour string `json:"schedule_hour"`
	}
	GetDoctors struct {
		DoctorSlug     string `json:"doctor_slug"`
		DoctorName     string `json:"doctors_name"`
		DoctorDesc     string `json:"doctors_desc"`
		IsReady        int    `json:"is_ready"`
		DoctorImageUrl string `json:"doctors_img_url"`

		// Props
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
