package header

// HTTPHeader Http 请求 头
type HTTPHeader map[string]string

// Set 设置值
func (h HTTPHeader) Set(key, value string) {
	h[key] = value
}

// Get 获取值
func (h HTTPHeader) Get(key string) string {
	return h[key]
}
