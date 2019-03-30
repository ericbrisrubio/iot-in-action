package main

import "fmt"

type TableTest struct {
    currentPosition      int
    speed                int
    timeBetweenPositions int
}

func (t *TableTest) GetCurrentPosition() int {
    return t.currentPosition
}

func (t *TableTest) GetSpeedCm() int {
    return t.speed
}

func (t *TableTest) On() {
    fmt.Println("Going up")
}

func (t *TableTest) Off() {
    fmt.Println("Going Down")
}

func (t *TableTest) Position(movement int) {
    position := t.DetermineTableMovement(movement)
    t.currentPosition = position
}

func (t *TableTest) Stop() {
    fmt.Println("Stopped")
}

func (t *TableTest) DetermineTableMovement(goTo int) int {
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

func (t *TableTest) InitialSetup() Mover{
    t.currentPosition = 0
    t.speed = 12
    t.timeBetweenPositions = 3
    return t
}
