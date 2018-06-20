# Test

```go
package view

import (
	"fmt"
	"html/template"
	"testing"
)

// indexModel index.html 模型
type indexModel struct {
	Title string
	Body  string
}

var (
	views = make(map[string]*View)
)

// registerViewUseNew 注册模板
func registerViewUseNew(name, filename string) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}
	view := New()
	view.SetTemplate(tmpl)
	// or
	// view.SetTemplateParseFiles(filename)
	views[name] = view
}

// registerViewUseNewViewParseFiles 注册模板
func registerViewUseNewViewParseFiles(name, filename string) {
	view, err := NewViewParseFiles(filename)
	if err != nil {
		panic(err)
	}
	views[name] = view
}

// TestInit 测试初始化
func testInit() {
	//registerViewUseNew("index", "template/index.html")
	// or
	registerViewUseNewViewParseFiles("index", "template/index.html")
}
func TestGenerateHTML(t *testing.T) {
	testInit()
	index := indexModel{
		Title: "this is a title.",
		Body:  "this is a body.",
	}
	view := views["index"]
	view.SetModel(&index)
	fmt.Println(view.HTML())
}
```

# index.html

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.Title}}</title>
</head>
<body>
    {{.Body}}
</body>
</html>
```