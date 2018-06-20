package view

import (
	"bytes"
	"html/template"
)

// View 视图
type View struct {
	buffer   *bytes.Buffer
	template *template.Template
}

func (v *View) Write(data []byte) (n int, err error) {
	return v.buffer.Write(data)
}

// HTML 获取Html
func (v *View) HTML() string {
	return v.buffer.String()
}

// New 创建一个新对象
func New() *View {
	return &View{
		buffer: new(bytes.Buffer),
	}
}

// NewViewParseFiles 创建视图，解析文件
func NewViewParseFiles(filename string) (view *View, err error) {
	template, err := template.ParseFiles(filename)
	if err != nil {
		return
	}
	return &View{
		buffer:   new(bytes.Buffer),
		template: template,
	}, nil
}

// SetTemplate 设置模板
func (v *View) SetTemplate(template *template.Template) {
	v.template = template
}

// SetTemplateParseFiles 解析文件
func (v *View) SetTemplateParseFiles(filename string) (err error) {
	v.template, err = template.ParseFiles(filename)
	return
}

// SetModel 设置数据模型
func (v *View) SetModel(model interface{}) {
	v.template.Execute(v, model)
}
