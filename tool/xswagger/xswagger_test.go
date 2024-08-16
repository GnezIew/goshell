package xswagger

import "testing"

func TestXSwagger_CompareSwaggerJson(t *testing.T) {
	xs := NewXswagger()
	xs.CompareSwaggerJson("./question.json", "./question2.json")
}
