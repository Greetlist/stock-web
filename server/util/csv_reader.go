package util

import (
    "os"
    "bufio"
    "glog"
    "io"
    "strings"
    "errors"
    "strconv"

    "greetlist/stock-web/server/model"
)

//type DataRow []string

type CsvReader struct {
    Head []string
    Data []model.SingleRecord
}

func (reader *CsvReader) ReadCsv(fileName string) error {
    file, err := os.Open(fileName)
    if err != nil {
        glog.Errorf("Open %s Error : %s.", fileName, err)
        return err
    }
    defer file.Close()
    fileReader := bufio.NewReader(file)
    var readHead bool = false
    for {
        var lineByte []byte
        var isPrefix bool
        var err error
        lineByte, isPrefix, err = fileReader.ReadLine()
        lineString := string(lineByte)
        if err == nil {
            if readHead == false {
                reader.Head = strings.Split(lineString, ",")
                if !checkHead(reader.Head) {
                    return errors.New("Head Has Empty Columns")
                }
                readHead = true
            } else {
                rawData := strings.Split(lineString, ",")
                singleRecord := convertRawData(rawData)
                reader.Data = append(reader.Data, singleRecord)
            }
        }
        if err != nil || isPrefix {
            if err == io.EOF {
                break
            }
            glog.Errorf("Read Csv File Line Error : %s.", err)
            return err
        }
    }
    return nil
}

func checkHead(head []string) bool {
    for _, v := range(head) {
        if v == "" {
            return false
        }
    }
    return true
}

func convertRawData(rawData []string) model.SingleRecord {
    record := model.SingleRecord{}
    record.Date = rawData[0]
    record.Open, _ = strconv.ParseFloat(rawData[1], 64)
    record.Close, _ = strconv.ParseFloat(rawData[2], 64)
    record.High, _ = strconv.ParseFloat(rawData[3], 64)
    record.Low, _ = strconv.ParseFloat(rawData[4], 64)
    record.Volume, _ = strconv.ParseFloat(rawData[5], 64)
    record.Money, _ = strconv.ParseFloat(rawData[6], 64)
    return record
}


