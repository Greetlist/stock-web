package util
import (
    "strings"
)

func selectExchange(stockCode string) string {
    if strings.Contains(stockCode, "SH") {
        return "SSE"
    } else if strings.Contains(stockCode, "SZ") {
        return "SZSE"
    }
    return ""
}
