package domain

type DoctorRepositoryStub struct {
	doctors []Doctor
}

func (s DoctorRepositoryStub) FindAll() ([]Doctor, error) {
	return s.doctors, nil
}

func NewDoctorRespositoryStub() DoctorRepositoryStub {
	doctors := []Doctor{
		{
			DoctorId: "1",
			Fullname: "john",
			Availability: []Timings{
				{
					From: "10",
					To:   "12",
				},
				{
					From: "4",
					To:   "6",
				},
			},
		},

		{
			DoctorId: "2",
			Fullname: "clark",
			Availability: []Timings{
				{
					From: "08",
					To:   "10",
				},
				{
					From: "18",
					To:   "20",
				},
			},
		},

		{
			DoctorId: "3",
			Fullname: "bruce",
			Availability: []Timings{
				{
					From: "12",
					To:   "14",
				},
				{
					From: "20",
					To:   "22",
				},
			},
		},
	}
	return DoctorRepositoryStub{doctors}
}
