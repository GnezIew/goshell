package xswagger

type IxSwagger interface {
	CompareSwaggerJson(oriFilePath string, newFilePath string) string
}
