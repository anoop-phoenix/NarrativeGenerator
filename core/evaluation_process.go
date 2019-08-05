package core

import "fmt"

type EvaluationProcesses []EvaluationProcess

func (evaluationProcesses EvaluationProcesses) GenerateStory() {
	level := 0
	for _, evaluationProcess := range evaluationProcesses {
		printEvaluationProcess(evaluationProcess, level)
	}
}

func tabLiteral(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s = s + "\t"
	}
	return s
}
func printEvaluationProcess(evaluationProcess EvaluationProcess, level int) {
	fmt.Printf("%s%s\n", tabLiteral(level), evaluationProcess.Chapter)
	fmt.Printf("%s<%s: %v>\n", tabLiteral(level), evaluationProcess.ComposedValue.ValueType, evaluationProcess.ComposedValue.Value)
	for _, childEvaluationProcess := range evaluationProcess.Childs {
		printEvaluationProcess(childEvaluationProcess, level+1)
	}
}
