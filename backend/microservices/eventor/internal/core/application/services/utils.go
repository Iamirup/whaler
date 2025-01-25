package services

import (
	"log"
	"time"
)

// // formatNumber adds commas to a number for better readability
// func formatNumber(num float64) string {
// 	in := fmt.Sprintf("%.0f", num)
// 	out := []string{}
// 	for i, c := range in {
// 		if (len(in)-i)%3 == 0 && i != 0 {
// 			out = append(out, ",")
// 		}
// 		out = append(out, string(c))
// 	}
// 	return strings.Join(out, "")
// }

// isWithinLastDay checks if the given time string is within the last 24 hours
func isWithinLastAge(timeStr string, age int) bool {
	const layout = "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return time.Since(t) <= time.Duration(age)*time.Second
}
