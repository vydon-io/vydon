package transformers

type TemplateData struct {
	Name        string
	Description string
	Example     string
}

type VydonTransformer interface {
	ParseOptions(opts map[string]any) (any, error)
	GetJsTemplateData() (*TemplateData, error)
	Transform(value any, opts any) (any, error)
}

type VydonGenerator interface {
	ParseOptions(opts map[string]any) (any, error)
	GetJsTemplateData() (*TemplateData, error)
	Generate(opts any) (any, error)
}
