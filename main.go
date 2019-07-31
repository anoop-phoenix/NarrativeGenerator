package main

import (
	"fmt"

	"github.com/zquangu112z/NarrativeGenerator/core"
)

func main() {
	student1 := core.Student{
		Name: "Nicholas",
		Subjects: []core.Subject{
			core.Subject{
				Assignments: []core.Assignment{8, 9, 10},
			},
			core.Subject{
				Assignments: []core.Assignment{1, 2, 10},
			},
			core.Subject{
				Assignments: []core.Assignment{10, 10, 10, 10},
			},
		},
		Ielts: true,
	}

	student2 := core.Student{
		Name: "Limmie",
		Subjects: []core.Subject{
			core.Subject{
				Assignments: []core.Assignment{8, 9, 10},
			},
			core.Subject{
				Assignments: []core.Assignment{10, 10, 10},
			},
		},
		Ielts: false,
	}

	student3 := core.Student{
		Name: "Kilia",
		Subjects: []core.Subject{
			core.Subject{
				Assignments: []core.Assignment{8, 9, 10},
			},
			core.Subject{
				Assignments: []core.Assignment{10, 10, 10},
			},
		},
		Ielts: true,
	}

	student1_narrative := student1.Evaluate()
	fmt.Println(student1_narrative)

	student2_narrative := student2.Evaluate()
	fmt.Println(student2_narrative)

	student3_narrative := student3.Evaluate()
	fmt.Println(student3_narrative)
}
