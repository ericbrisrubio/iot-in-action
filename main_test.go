package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "net/http/httptest"
    "testing"
)


func TestDetermineTableMovement(t *testing.T){
    currentPosition = 0
    expectedMovement := determineTableMovement(5)
    if expectedMovement != 5 {
        t.Log("Failed to move to the expected position 5")
        t.Fail()
    }
    currentPosition = 5
    expectedMovement = determineTableMovement(3)
    if expectedMovement != -2 {
        t.Log("Failed to move to the expected position -2")
        t.Fail()
    }
}

func TestPosition(t *testing.T){

    r := mux.NewRouter()
    r.HandleFunc("/position/{number}", position).Name("func1")

    req, err := http.NewRequest("GET", "http://192.168.86.211:8070/position/3", nil)

    //contextG.Set(req, "number", 3)

    if err != nil {
        fmt.Errorf("Request mock failed to be created %s", err.Error())
    }
    recorder := httptest.NewRecorder()

    r.ServeHTTP(recorder,req)
    if currentPosition != 3 {
        t.Error("The position is not the desired")
    }
}

func TestRenderText(t *testing.T) {

}
