package entity

type RawProblem struct {
	Title        string
	Description  string
	TimeLimit    int64
	MemoryLimit  int64
	Input        string
	Output       string
	SampleInput  string
	SampleOutput string
	Hint         string
	Source       string
}
