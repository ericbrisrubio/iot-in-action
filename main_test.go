package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "net/http/httptest"
    "testing"
)

func SetupMoverManager() *moverManager{
    tableObj := SetupTableMover()
    moverManager := new(moverManager)
    moverManager.MoverObj = tableObj
    return moverManager
}


func SetupTableMover() *TableTest{
    tableObj := new(TableTest)
    tableObj.InitialSetup()
    return tableObj
}

func TestDetermineTableMovement(t *testing.T){
    table := SetupTableMover()
    expectedMovement := table.DetermineTableMovement(5)
    if expectedMovement != 5 {
        t.Log("Failed to move to the expected position 5")
        t.Fail()
    }
    table.currentPosition = 5
    expectedMovement = table.DetermineTableMovement(3)
    if expectedMovement != -2 {
        t.Log("Failed to move to the expected position -2")
        t.Fail()
    }
}

func TestPosition(t *testing.T){

    mm := SetupMoverManager()
    r := mux.NewRouter()
    r.HandleFunc("/position/{number}", mm.position).Name("func1")

    req := httptest.NewRequest("GET", "http://192.168.86.211:8070/position/3", nil)

    recorder := httptest.NewRecorder()

    r.ServeHTTP(recorder,req)
    if mm.MoverObj.GetCurrentPosition() != 3 {
        t.Error("The position is not the desired")
    }
}

/**
Test the error when the position is different from an int value
 */
func TestPositionError(t *testing.T){

    mm := SetupMoverManager()
    r := mux.NewRouter()
    r.HandleFunc("/position/{number}", mm.position).Name("func1")

    req := httptest.NewRequest("GET", "http://192.168.86.211:8070/position/test", nil)

    recorder := httptest.NewRecorder()

    r.ServeHTTP(recorder,req)
    if recorder.Code != http.StatusForbidden {
        t.Error("The position is not the desired")
    }
}

func TestRenderText(t *testing.T) {

}
