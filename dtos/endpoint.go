package dtos

type Endpoint struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url`
	Domain  string
	Method  string            `yaml:"method,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Body    string            `yaml:"body,omitempty"`
}
