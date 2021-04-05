package lib

import "fmt"

type Controller struct {
  Moveable []Moveable
  Move     bool
  Copy     bool
}

func PrepareController(Moveable []Moveable, Move bool, Copy bool) *Controller {
  return &Controller{
    Move:     Move,
    Copy:     Copy,
    Moveable: Moveable,
  }
}

func (c *Controller) Start() {
  switch true {
  case c.Move:
    c.moveFile()
  case c.Copy:
    c.copyFile()
  }
}

func (c *Controller) moveFile() {
  for _, d := range c.Moveable {
    fmt.Println("Moving", "["+d.FileName+"]")
    if err := MoveFile(d.Source, d.Dest); err != nil {
      fmt.Println(err.Error())
      return
    }
  }
}

func (c *Controller) copyFile() {
  for _, d := range c.Moveable {
    fmt.Println("Cloning", "["+d.FileName+"]")
    if err := CopyFile(d.Size, d.Source, d.Dest); err != nil {
      fmt.Println(err.Error())
      return
    }
  }
}
