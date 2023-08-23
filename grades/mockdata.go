package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Alice",
			LastName:  "Smith",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 80,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Bob",
			LastName:  "Smith",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 82,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 70,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 86,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Frank",
			LastName:  "Smith",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 90,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 95,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 95,
				},
			},
		},
	}
}
