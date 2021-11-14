package util
import (
    "strings"
    "os"
    "time"
    "path"
    "github.com/go-gota/gota/dataframe"
    "greetlist/stock-web/server/conf"
)

func selectExchange(stockCode string) string {
    if strings.Contains(stockCode, "SH") {
        return "SSE"
    } else if strings.Contains(stockCode, "SZ") {
        return "SZSE"
    }
    return ""
}


func LastTradingDayDirStr() string {
    tradingDayPath := path.Join(conf.StockRawBaseDir, "trading_date.csv")
    f, _ := os.OpenFile(tradingDayPath, os.O_RDWR, os.ModePerm)
    df := dataframe.ReadCSV(f)
    todayDateStr := strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", "")
    tradeDateList := df.Col("cal_date").Records()
    for index, value := range(tradeDateList) {
        if value >= todayDateStr {
            targetDate := tradeDateList[index-1]
            return targetDate[0:4] + "/" + targetDate[4:6] + "/" + targetDate[6:]
        }
    }
    return "1970/01/01"
}
