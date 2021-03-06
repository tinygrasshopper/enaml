package generators

const (
	structTemplate = `package {{.PackageName}} 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type {{.JobName}} struct {
{{ range $key, $value := .Elements }}
	/*{{ $value.ElementName }} - {{ $value.ElementComments }}*/
	{{ $value.ElementName }} {{ $value.ElementType }} {{$value.ElementYamlName}}
{{ end }}
}`
)
