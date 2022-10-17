package main

import (
	"fmt"
	"os"
	"errors"
	"os/exec"
)

func toCodes(command string) (string, string, error) {
  switch command {
  case "dp":
    return "0x60", "0x0f", nil
  case "mdp":
    return "0x60", "0x11", nil
  case "hdmi":
    return "0x60", "0x11", nil
  case "typec", "type-c":
    return "0x60", "0x1b", nil
  case "standard":
    return "0xdc", "0x00", nil
  case "movie":
    return "0xdc", "0x03", nil
  case "game":
    return "0xdc", "0x05", nil
  case "comfort":
    return "0x0f", "0x0c", nil
  }

  return "", "", errors.New("Invalid command")
}

func toModel(target string) (string, error) {
  switch target {
  case "main":
    return "DELL U2720QM", nil
  case "sub":
    return "DELL U2718Q", nil
  }

  return "", errors.New("Invalid target")
}

func main() {
  if len(os.Args) < 3 {
    fmt.Println(`Not enough arguments

display-control <TARGET> <COMMAND>

  TARGET    main sub
  COMMAND   dp hdmi mdp htmi typec standard movie game comfort`)
    os.Exit(1)
  }

  target := os.Args[1]
  command := os.Args[2]

  first, second, err := toCodes(command)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  model, err := toModel(target)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  out, err := exec.Command("sudo", "ddcutil", "--model", model, "setvcp", first, second).Output()
  if (err != nil) {
    fmt.Printf("Failed: %s\n", err)
    return
  }

  fmt.Println(string(out))
}
