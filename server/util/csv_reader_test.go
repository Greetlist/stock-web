package util

import (
    "os"
    "testing"
    "fmt"
)

func TestCsvReader(t *testing.T) {
    csvReader := CsvReader{}
    fileName := "/home/greetlist/macd/data_storage/603788.XSHG/stock_daily_data/kline_daily.csv"
    if err := csvReader.ReadCsv(fileName); err != nil {
        os.Exit(1)
    }

    recordNum := len(csvReader.Data)
    fmt.Printf("Total Data is : %v.\n", recordNum)
    for i := recordNum - 10; i < recordNum; i++ {
        fmt.Printf("%f.\n", csvReader.Data[i].Open)
    }
}
