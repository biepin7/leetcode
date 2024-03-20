package _020

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	// 这玩意的目的是配对，如果当前字符是一个右括号，则 pairs[s[i]] 将返回其对应的左括号；否则，如果当前字符不是右括号，则返回零值。
	// 为啥返回 0 值：当我们试图从一个映射中获取一个不存在的键时，映射会返回该值类型的零值。
	// 为什么非要 pairs 而不是 if ，因为后续做增加的时候，好判断
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { //如果当前是右括号，尝试匹配左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1] // 去掉头
			}
		} else { // 是左括号，入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
