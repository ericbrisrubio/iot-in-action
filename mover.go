package main

import (
    "fmt"
    "github.com/stianeikeland/go-rpio"
    "os"
    "sync"
    "time"
)

var (
    pin  = rpio.Pin(14)
    pin1 = rpio.Pin(15)
)

const (
    POSITION1 = iota + 1
    POSITION2
    POSITION3
    POSITION4
    POSITION5
    RESET
    HARDRESET
)

type Mover interface {
    GetCurrentPosition() int
    GetSpeedCm() int
    InitialSetup() Mover
    On()
    Off()
    Position(int)
    Stop()
}

type Table struct {
    mu                   *sync.RWMutex
    currentPosition      int
    speed                int
    timeBetweenPositions int // time in seconds interval between consecutive positions
}

func (t *Table) On() {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
    pin.Output()
    pin1.Output()

    pin1.Low()
    pin.High()
}

func (t *Table) Off() {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
    pin.Output()
    pin1.Output()

    pin.Low()
    pin1.High()
}

func (t *Table) Stop() {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
    pin.Output()
    pin1.Output()

    pin.Low()
    pin1.Low()
}

func (t *Table) Position(positionMovement int) {
    positionMovement = t.DetermineTableMovement(positionMovement)
    if positionMovement == -15 {
        t.currentPosition = 0;
    }
    if positionMovement > 0 {
        t.mu.Lock()
        t.On()
        time.Sleep(time.Duration(positionMovement) * time.Duration(t.timeBetweenPositions) * time.Second)
        t.currentPosition = positionMovement
        t.mu.Unlock()
    } else {
        t.mu.Lock()
        t.Off()
        positionMovement = -1 * positionMovement
        time.Sleep(time.Duration(positionMovement) * time.Duration(t.timeBetweenPositions) * time.Second)
        t.currentPosition = positionMovement
        t.mu.Unlock()
    }
    t.Stop()
}

func (t *Table) GetCurrentPosition() int {
    t.mu.RLock()
    currentPositionR := t.currentPosition
    t.mu.RUnlock()
    return currentPositionR
}

func (t *Table) GetSpeedCm() int {
    return t.speed
}

func (t *Table) DetermineTableMovement(goTo int) int {
    if goTo == POSITION1 || goTo == POSITION2 || goTo == POSITION3 || goTo == POSITION4 || goTo == POSITION5 || goTo == RESET || goTo == HARDRESET {
        if goTo == RESET {
            return -1 * t.currentPosition
        }
        if goTo == HARDRESET {
            return -15
        }
        return goTo - t.currentPosition
    }
    return 0
}

func (t *Table) InitialSetup() Mover{
    t.currentPosition = 0
    t.speed = 12
    t.timeBetweenPositions = 3
    return t
}
