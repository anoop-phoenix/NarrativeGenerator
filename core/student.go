package core

import "time"

type Assignment int

type Subject struct {
	Name        string
	Assignments []Assignment
}

type Student struct {
	Name     string
	Subjects []Subject
	Ielts    bool
}

func evaluate_Subjects(subjects []Subject) Narrative {
	narrative := Narrative{
		Desc:   "All subjects is $gte 8.5",
		Status: true,
	}
	for _, subject := range subjects {
		assignments_narrative := evaluate_Assignments(subject.Assignments)
		narrative.Childs = append(narrative.Childs)
		if assignments_narrative.Status == false {
			narrative.Status = false
			return narrative
		}
	}
	return narrative
}

func evaluate_Assignments(assignment []Assignment) Narrative {
	narrative := Narrative{
		Desc:   "All assignments is over 5",
		Status: true,
	}

	for _, assignment := range assignment {
		if assignment < 5 {
			narrative.Status = false
			return narrative
		}
	}
	return narrative
}

func (s Student) Evaluate() EvaluateResult {
	evaluateResult := EvaluateResult{
		EvaluatedDate: time.Now(),
	}
	narrative := Narrative{}
	// Evaluate Subjects
	subjects_narrative := evaluate_Subjects(s.Subjects)
	narrative.Childs = append(narrative.Childs, subjects_narrative)
	// Evaluate Ielts
	evaluateResult.Narrative = narrative
	return evaluateResult
}
