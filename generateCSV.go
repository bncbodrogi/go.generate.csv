package main

import(
  "encoding/csv"
  "log"
  "math/rand"
  "time"
  "strconv"
  "fmt"
  "os"
  "github.com/mohae/struct2csv"
)

type MockData struct {
  Id int
  Description string
  Serial string
  Date string
  Parent_Id int
  Parent_description string
  Has_image bool
  Downloadable bool
  Main_url string
  Media_url string
}

func main(){
  now := time.Now()
  file, err := os.Create("result.csv")
  checkErr("Cannot create file", err)
  defer file.Close()

  writer := csv.NewWriter(file)
  defer writer.Flush()
  writer.Comma = ';'

  var data = make([]MockData, 20000)
  for i := 0; i < len(data); i++ {
    data[i] = generateMock()
  }

  enc := struct2csv.New()
  rows, err := enc.Marshal(data)
  checkErr("Can't marshal data ", err)

  for _, value := range rows {
    err := writer.Write(value)
    checkErr("Cannot write to file", err)
  }

  timeSpent := time.Since(now)
  spent := timeSpent.Seconds()
  log.Println("time spent: ")
  log.Println(spent)
}

func generateMock() MockData{

  letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

  t := time.Now()

  return MockData{
    Id: randomNumber(8),
    Description: fmt.Sprintf("Sample description%d", randomNumber(3)),
    Serial: fmt.Sprintf("%d-%c%c%c", randomNumber(3), letterRunes[randomNumber(1)], letterRunes[randomNumber(1)], letterRunes[randomNumber(1)]),
    Date: fmt.Sprintf("%d/%s/%d", t.Year(), t.Month(), t.Day()),
    Parent_Id: randomNumber(8),
    Parent_description: fmt.Sprintf("Sample description%d", randomNumber(3)),
    Has_image: randomNumber(1) % 2 == 0,
    Downloadable: randomNumber(1) % 2 == 0,
    Main_url: "http://main.url.com",
    Media_url: "http://media.url.com",
  }
}

func randomNumber(digits int) int {
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)

  result := ""
  for i := 0; i < digits; i++ {
    result = result + strconv.Itoa(random.Intn(10))
  }

  intRes, err := strconv.Atoi(result)
  checkErr("Can't convert string to int", err)

  return intRes
}

func checkErr(message string, err error){
  if err!= nil{
    log.Fatal(message, err)
  }
}
