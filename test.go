package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/google/uuid"
)

func main() {
  ex, _:= os.Executable()
  exPath := filepath.Dir(ex)
  s1 := fmt.Sprintf("The following is running: %s", exPath)
  instance := uuid.New().String()
  s2 := fmt.Sprintf("And this is a uuid: %s", instance)
  token := os.Getenv("ENV_VAR")
  s3 := fmt.Sprintf("And this is an env var: %s", token)
  s4 := fmt.Sprintf("And this is a split string: %s",  strings.SplitN("1=1+0", "=", 2))
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println(s3)
  fmt.Println(s4)
}
