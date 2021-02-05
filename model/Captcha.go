package model

type Captcha struct {
	Value         string `json:"value"`
	ValueSolution string `json:"valueSolution" binding:"required"`
}
