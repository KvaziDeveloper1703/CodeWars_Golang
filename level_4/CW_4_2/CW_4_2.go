/*
Write a function that formats a given non-negative number of seconds into a human-readable duration.
	- If the input is 0, return "now";
	- Otherwise, express the time using years, days, hours, minutes, seconds.

Rules:
	- Use only units with non-zero values;
	- Use each unit at most once and always in order from largest to smallest;
	- Use plural forms when needed;
	- Separate parts with ", " and use " and " before the last part;
	- Always convert time using the largest possible.

Напишите функцию, которая форматирует неотрицательное число секунд в человекочитаемый вид.
	- Если число равно 0, вернуть "now";
	- Иначе представить время через годы, дни, часы, минуты, секунды.

Правила:
	- Использовать только ненулевые единицы;
	- Единицы идут строго от больших к меньшим, без повторов;
	- Множественное число - только если значение больше 1;
	- Части разделяются ", ", а перед последней - " and ";
	- Использовать максимально крупные единицы.
*/

package kata

import "fmt"

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	year := int64(365 * 24 * 3600)
	day := int64(24 * 3600)
	hour := int64(3600)
	minute := int64(60)

	years := seconds / year
	seconds %= year

	days := seconds / day
	seconds %= day

	hours := seconds / hour
	seconds %= hour

	minutes := seconds / minute
	seconds %= minute

	parts := []string{}

	if years > 0 {
		parts = append(parts, fmt.Sprintf("%d year%s", years, plural(years)))
	}
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d day%s", days, plural(days)))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hour%s", hours, plural(hours)))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d minute%s", minutes, plural(minutes)))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%d second%s", seconds, plural(seconds)))
	}

	if len(parts) == 1 {
		return parts[0]
	}
	if len(parts) == 2 {
		return parts[0] + " and " + parts[1]
	}

	result := ""
	for i := 0; i < len(parts)-1; i++ {
		result += parts[i] + ", "
	}
	result = result[:len(result)-2]
	result += " and " + parts[len(parts)-1]

	return result
}

func plural(n int64) string {
	if n == 1 {
		return ""
	}
	return "s"
}