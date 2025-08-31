func IsValidBase64(content string) bool {

	s := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // 删除空白字符
		}
		return r
	}, string(content))

	// 1. 检查长度：Base64 编码后的字符串长度必须是 4 的倍数
	if len(s)%4 != 0 {
		return false
	}

	// 2. 检查字符：只允许 Base64 字符集和末尾的填充符 '='
	// 并且 '=' 只能出现在末尾
	paddingStart := -1
	for i, r := range s {
		if r == '=' {
			if paddingStart == -1 {
				paddingStart = i // 记录第一个 '=' 的位置
			}
		} else {
			if paddingStart != -1 {
				// 在 '=' 之后又出现了非 '=' 字符，无效
				return false
			}
			// 检查字符是否在 Base64 字符集中
			if !((r >= 'A' && r <= 'Z') ||
				(r >= 'a' && r <= 'z') ||
				(r >= '0' && r <= '9') ||
				r == '+' || r == '/') {
				return false
			}
		}
	}

	// 3. 检查填充符数量：末尾的 '=' 数量只能是 0, 1, 或 2 个
	if paddingStart != -1 {
		paddingCount := len(s) - paddingStart
		if paddingCount > 2 {
			return false
		}
	}

	// 4. 尝试解码，看是否成功
	_, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}

	// 所有检查通过
	return true
}
