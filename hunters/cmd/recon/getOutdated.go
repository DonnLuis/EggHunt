package recon

import (
  "os/exec"
  "fmt"
)

func Updates() {
  cmd := exec.Command("brew","update")
  out, err := cmd.Output()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(out))
}
