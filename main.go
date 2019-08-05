package main

import (
	"github.com/zquangu112z/NarrativeGenerator/common"
	"github.com/zquangu112z/NarrativeGenerator/core"
)

func main() {
	// student1 := core.Student{
	// 	Name: "Nicholas",
	// 	Subjects: []core.Subject{
	// 		core.Subject{
	// 			Name:        "Science",
	// 			Assignments: []float64{8, 9, 10},
	// 		},
	// 		core.Subject{
	// 			Name:        "Physical Education",
	// 			Assignments: []float64{1, 2, 10},
	// 		},
	// 		core.Subject{
	// 			Name:        "Geography",
	// 			Assignments: []float64{10, 10, 10, 10},
	// 		},
	// 	},
	// 	Ielts: true,
	// 	Jlpt:  true,
	// }
	// common.PrettyPrint(student1.Evaluate())

	student2 := core.Student{
		Name: "Limmie",
		Subjects: []core.Subject{
			core.Subject{
				Name:        "Science",
				Assignments: []float64{8, 9, 10},
			},
			core.Subject{
				Name:        "Physical Education",
				Assignments: []float64{10, 10, 10},
			},
		},
		Ielts: true,
		Jlpt:  true,
	}
	common.PrettyPrint(student2.Evaluate())

	student3 := core.Student{
		Name: "Kilia",
		Subjects: []core.Subject{
			core.Subject{
				Name:        "Science",
				Assignments: []float64{8, 9, 10},
			},
			core.Subject{
				Name:        "Physical Education",
				Assignments: []float64{10, 10, 10},
			},
		},
		Ielts: true,
		Jlpt:  false,
	}
	common.PrettyPrint(student3.Evaluate())
	common.PrettyPrint(student3.Average())

}
