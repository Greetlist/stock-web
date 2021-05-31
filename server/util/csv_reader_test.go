package util

import (
    "os"
    "testing"
    "fmt"
)

func TestCsvReader(t *testing.T) {
    csvReader := CsvReader{}
    if err := csvReader.ReadCsv("/home/greetlist/macd/data_storage/603788.XSHG/stock_daily_data/kline_daily.csv"); err != nil {
        os.Exit(1)
    }
    fmt.Printf("Total Data is : %v.\n", len(csvReader.Data[10]))
    fmt.Printf("Total Data is : %v.\n", csvReader.Data[10])
    fmt.Printf("Total Data Head is : %v.\n", csvReader.Head)
}
