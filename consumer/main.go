package main

import (
  "fmt"
  "os"
  "io/ioutil"

  "github.com/jzelinskie/geddit"
)

func main() {
  passwordBytes, err := ioutil.ReadFile(".credentials")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  session, err := geddit.NewLoginSession("wonderb0lt", string(passwordBytes), "test")

  if err != nil {
    fmt.Println(err)
    os.Exit(2)
  }

  submissions, err := session.Frontpage()

  if (err != nil) {
    fmt.Println(err)
    os.Exit(3)
  }

  fmt.Println(submissions)
}
