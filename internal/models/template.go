package models

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Return struct {
	Type string `json:"type"`
}

type Signature struct {
	FunctionName string  `json:"function_name"`
	Parameters   []Param `json:"parameters"`
	Returns      Return  `json:"returns"`
}

type TemplateRequest struct {
	QuestionID  string    `json:"question_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Signature   Signature `json:"signature"`
	Language    string    `json:"language"`
}
