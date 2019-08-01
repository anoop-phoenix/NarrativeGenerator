package main

import (
	"encoding/json"
	"fmt"

	"github.com/zquangu112z/NarrativeGenerator/core"
)

// print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
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
	PrettyPrint(student1_narrative)

	student2_narrative := student2.Evaluate()
	PrettyPrint(student2_narrative)

	student3_narrative := student3.Evaluate()
	PrettyPrint(student3_narrative)
}
