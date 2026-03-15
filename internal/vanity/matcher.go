package vanity

import (
	"strings"
)

// Result 表示找到的匹配结果
type Result struct {
	PrivateHex    string
	AddressBase58 string
}

// Matcher 用于匹配地址的前缀和后缀
type Matcher struct {
	prefix string
	suffix string
}

// NewMatcher 创建一个新的匹配器
func NewMatcher(p, s string) *Matcher {
	// 清理输入并转换为大写（因为Base58Alphabet是大写的）
	p = strings.TrimSpace(p)
	s = strings.TrimSpace(s)
	return &Matcher{
		prefix: strings.ToUpper(p),
		suffix: strings.ToUpper(s),
	}
}

// Match 检查地址是否匹配前缀和后缀
func (m *Matcher) Match(addr string) bool {
	// 由于Base58Alphabet使用大写字母，生成的地址已经是大写的，不需要转换
	
	// 检查前缀
	if m.prefix != "" && !strings.HasPrefix(addr, m.prefix) {
		return false
	}
	
	// 检查后缀
	if m.suffix != "" && !strings.HasSuffix(addr, m.suffix) {
		return false
	}
	
	return true
}