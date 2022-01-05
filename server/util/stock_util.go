package util
import (
    "strings"
    "os"
    "time"
    "path"
    "bufio"
    "fmt"
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


// return today_str and last_trading_date_str
func LastTradingDayDirStr() (string, string) {
    tradingDayPath := path.Join(conf.StockRawBaseDir, "trading_date.csv")
    f, _ := os.OpenFile(tradingDayPath, os.O_RDWR, os.ModePerm)
    df := dataframe.ReadCSV(f)
    todayStr := time.Now().Format("2006-01-02")
    todayDateStr := strings.ReplaceAll(todayStr, "-", "")
    tradeDateList := df.Col("cal_date").Records()
    for index, value := range(tradeDateList) {
        if value >= todayDateStr {
            targetDate := tradeDateList[index-1]
            return strings.ReplaceAll(todayStr, "-", "/"), targetDate[0:4] + "/" + targetDate[4:6] + "/" + targetDate[6:]
        }
    }
    return strings.ReplaceAll(todayStr, "-", "/"), "1970/01/01"
}

func ReadPredictionFileInOrder(filePath string) []string {
    var res []string
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Open File Error is : %s", err)
        return res
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasSuffix(line, "SH") || strings.HasSuffix(line, "SZ") {
            res = append(res, line)
        }
    }
    return res
}

