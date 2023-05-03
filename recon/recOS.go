package recon

import (
  "runtime"
)

func GetOS() string {
  return runtime.GOOS
}
