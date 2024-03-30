package wordchange

import (
	"strings"
	"unicode"
)

// ConvertStringFormat 转换字符串到指定的格式
func ConvertStringFormat(src, format string) string {
	if src == "" {
		return ""
	}
	src = ToPascalCaseWithSpace(src)

	// 识别原始字符串格式（这里简化处理，实际识别可能需要更复杂的逻辑）

	// 转换为指定的格式
	switch format {
	case "up":
		return strings.ToUpper(src)
	case "low":
		return strings.ToLower(src)
	case "camel":
		return toCamelCase(src)
	case "pascal":
		return toPascalCase(src)
	case "snake", "sn", "snakeCamel":
		return strings.ToLower(toSnakeCase(src, true))
	case "snakeLower":
		return strings.ToLower(toSnakeCase(src, false))
	case "snakeUpper":
		return strings.ToUpper(toSnakeCase(src, false))
	case "snakePascal":
		return toSnakeCase(src, false)
	case "dash":
		return ToDashCase(src)
	default:
		return src
	}
}

// 辅助函数：转换为小驼峰格式
func toCamelCase(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

// 辅助函数：转换为大驼峰格式
func toPascalCase(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// ToPascalCase 将包含空格的字符串转换为大驼峰形式,另外一种形式，可以处理空格
func ToPascalCaseWithSpace(s string) string {
	var result strings.Builder
	// 标记是否遇到第一个字母，用于确保首字母大写
	firstLetter := true
	// 标记是否遇到空格，用于将下一个字母转换为大写
	nextToUpper := false

	for _, r := range s {
		if unicode.IsSpace(r) {
			// 如果遇到空格，则设置下一个字母为大写
			nextToUpper = true
			continue
		}

		if firstLetter || nextToUpper {
			// 将首字母或空格后的首字母转换为大写
			result.WriteRune(unicode.ToUpper(r))
			firstLetter = false
			nextToUpper = false
		} else {
			// 其他字母保持原样
			result.WriteRune(r)
		}
	}

	return result.String()
}

// // 辅助函数：转换为蛇形格式，camel参数控制是否驼峰蛇形
// func toSnakeCase(s string, camel bool) string {
// 	var result []rune
// 	for i, r := range s {
// 		if unicode.IsUpper(r) {
// 			if i > 0 {
// 				result = append(result, '_')
// 			}
// 			if camel {
// 				result = append(result, unicode.ToLower(r))
// 			} else {
// 				result = append(result, r)
// 			}
// 		} else {
// 			result = append(result, r)
// 		}
// 	}
// 	return string(result)
// }

func toSnakeCase(s string, camel bool) string {
	var result []rune
	needsUnderscore := false

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) && (unicode.IsLower(rune(s[i-1])) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
			// 如果当前字符是大写，并且前一个字符是小写，或者下一个字符是小写，则需要插入下划线
			needsUnderscore = true
		}

		if needsUnderscore && len(result) > 0 {
			result = append(result, '_')
			needsUnderscore = false
		}

		if camel {
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func ToDashCase(s string) string {
	var result strings.Builder
	var prevChar rune
	needsDash := false

	for i, r := range s {
		// 当前字符转换为小写
		currentChar := unicode.ToLower(r)

		if i > 0 && unicode.IsUpper(r) && (unicode.IsLower(prevChar) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
			// 如果当前字符是大写，并且前一个字符是小写，或者下一个字符是小写，则需要插入短横线
			needsDash = true
		}

		if needsDash && result.Len() > 0 {
			result.WriteRune('-')
			needsDash = false
		}

		result.WriteRune(currentChar)
		prevChar = r
	}

	return result.String()
}

func WordChangeDemo(original string) {

	formats := []string{"up", "low", "camel", "pascal", "snake", "snakeCamel", "snakeLower", "snakeUpper", "snakePascal", "dash"}

	for _, format := range formats {
		converted := ConvertStringFormat(original, format)
		println(format+":", converted)
	}
}
