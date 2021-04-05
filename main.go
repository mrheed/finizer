package main

import (
  "errors"
  "flag"
  "fmt"
  "os"
  "time"

  lib "github.com/syahidnurrohim/file_classifier/src"
)

var (
  move                bool
  copy                bool
  help                bool
  maxStairs           int
  startDir            string
  customFormatPath    string
  ignoreDefaultFormat bool
)

func log(err error) {
  fmt.Println(err.Error())
}

func init() {
  currentDirectory, err := os.Getwd()
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(0)
  }
  flag.BoolVar(&ignoreDefaultFormat, "ignoreDefaultFormat", false, lib.DefaultMimeHelper)
  flag.StringVar(&customFormatPath, "pathToFormat", "", lib.CustomFormatPathHelper)
  flag.IntVar(&maxStairs, "maxStairs", -1, lib.MaxStairsHelper)
  flag.StringVar(&startDir, "begin", currentDirectory, lib.BeginHelper)
  flag.BoolVar(&move, "move", false, lib.MoveHelper)
  flag.BoolVar(&copy, "copy", false, lib.CopyHelper)
  flag.BoolVar(&help, "help", false, lib.AllHelper)
  flag.Parse()
}

func main() {
  switch true {
  case help:
    lib.CallHelpers()
    return
  case move && copy:
    log(errors.New("parameters error: you cannot use --move argument alongside with --copy and otherwise"))
    return
  case !move && !copy:
    log(errors.New("parameter error: you should put at least one argument from one of --move and --copy"))
    return
  case ignoreDefaultFormat && customFormatPath == "":
    log(errors.New("parameters error: specify --pathToFormat directory in order to locate the custom format filtering rules"))
    return
  default:
    fmt.Println("Start to investigating files, this perform may take a moment depends on the number of files...")
    for range time.Tick(time.Second * 1) {
      break
    }
  }
  dir := lib.NewDir(startDir, customFormatPath, ignoreDefaultFormat, maxStairs)
  moveable, err := dir.ReadDir(0)
  if err != nil {
    log(err)
    return
  }
  lib.Sort(&moveable, 0, len(moveable)-1)
  lib.PrepareController(moveable, move, copy).Start()
  fmt.Println("Highest stairs", dir.HighestStairs)
  fmt.Println("Total folder explored", dir.FolderExplored)
  fmt.Println("Total file explored", dir.FileExplored)
  fmt.Println("Total items", dir.TotalItems)
  fmt.Println("Total skipped files", dir.FileSkipped)
}
