package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/stianeikeland/go-rpio"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
    text_tmpl "text/template"
)

var (
    // Use mcu pin 10, corresponds to physical pin 19 on the pi
    pin             = rpio.Pin(14)
    pin1            = rpio.Pin(15)
    currentPosition = -1
)

const (
    POSITION1    = 1
    POSITION2    = 2
    POSITION3    = 3
    POSITION4    = 4
    POSITION5    = 5
    RESET        = 6
    HARDRESET    = 7
    POSITIONTIME = 3
    BaseUrl = "http://192.168.86.211:8070"
)

func main() {
    if currentPosition == -1 {
        //read info from file and setup "currentPosition" var
        //in case the file does not exist or it is empty the var should be set to 0
        currentPosition = 0 //this should change later when reading from the file. this is a temp line
    }
    defineEndpoints()
}

func rootMethod(w http.ResponseWriter, r *http.Request) {
    // Open and map memory to access gpio, check for errors
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Unmap gpio memory when done
    defer rpio.Close()

    // Set pin to output mode
    pin.Output()

    w.WriteHeader(http.StatusOK)
    // Toggle pin 20 times
    //for x := 0; x < 5; x++ {
    pin.Toggle()
    //	time.Sleep(time.Second*3)
    //}

    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func on(w http.ResponseWriter, r *http.Request) {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
    pin.Output()
    pin1.Output()

    pin1.Low()
    pin.High()

    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func off(w http.ResponseWriter, r *http.Request) {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
    pin.Output()
    pin1.Output()

    pin.Low()
    pin1.High()

    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func stop(w http.ResponseWriter, r *http.Request) {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
    pin.Output()
    pin1.Output()

    pin.Low()
    pin1.Low()

    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func position(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    position := vars["number"]
    var err error
    var tempPosition int64
    tempPosition, err = strconv.ParseInt(position, 10, 64)
    if err != nil {
        fmt.Println(err.Error())
    }
    var positionInt int
    positionInt = int(tempPosition)
    positionMovement := determineTableMovement(positionInt)

    currentPosition = currentPosition + positionMovement
    if positionMovement > 0 {
        on(w, r)
        time.Sleep(time.Duration(positionMovement) * POSITIONTIME * time.Second)
    } else {
        off(w, r)
        positionMovement = -1 * positionMovement
        time.Sleep(time.Duration(positionMovement) * POSITIONTIME * time.Second)
    }

    stop(w, r)
    jsonReturn := map[string]interface{}{"version": "1.0"}
    json.NewEncoder(w).Encode(jsonReturn)
}

func defineEndpoints() {
    router := mux.NewRouter()
    //router.HandleFunc("/", rootMethod).Methods("GET")

    router.HandleFunc("/up", on).Methods("GET")
    router.HandleFunc("/down", off).Methods("GET")
    router.HandleFunc("/stop", stop).Methods("GET")
    router.HandleFunc("/position/{number}", position).Methods("GET")
    router.HandleFunc("/index", renderAdminPage).Methods("GET")
    router.HandleFunc("/cdn/{filename}", cdnProvider).Methods("GET")
    router.HandleFunc("/cdn/{filename}/{data_type}", cdnProvider).Methods("GET")

    //router.HandleFunc("/", rootMethod).Methods("POST")
    log.Fatal(http.ListenAndServe(":8070", router))

}

//render admin page from the index.html generated from React Polaris
func renderAdminPage(w http.ResponseWriter, r *http.Request) {
    //index := template.Must(template.ParseFiles("static/index.html"))
    data := map[string]interface{}{"Asi": "daleeeeee"}
    renderText(w, "index.html", data, "")
}

//this function determine how many positions the table has to be move and the way. Positive goes up and negative goes down
func determineTableMovement(goTo int) int {
    if goTo == POSITION1 || goTo == POSITION2 || goTo == POSITION3 || goTo == POSITION4 || goTo == POSITION5 || goTo == RESET || goTo == HARDRESET{
        if goTo == RESET {
            return -1 * currentPosition
        }
	if goTo == HARDRESET {
	    return -15
	}
        return goTo - currentPosition
    }
    return 0
}

//provide files being asked throughout the api
func cdnProvider(writer http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    filename := vars["filename"]
    data_type := vars["data_type"]
    fmt.Println(data_type)
    folder := ""
    if data_type == "" {
        fmt.Println("empty")
    }
    if data_type != "" && data_type == "css" {
        writer.Header().Set("Content-Type", "text/css")
        folder = "css/"
    } else if data_type == "js" {
        folder = "js/"
    }

    if filename == "test.js" {
        folder = "js/"
        parseVars := map[string]interface{}{"baseUrl": BaseUrl}
        renderText(writer, filename, parseVars, folder)
    } else if filename == "index.js" {
        folder = "js/"
        parseVars := map[string]interface{}{"Asi": BaseUrl}
        renderText(writer, filename, parseVars, folder)
    } else {
        renderText(writer, filename, nil, folder)
    }
}

//render file content based on a template
func renderText(w http.ResponseWriter, tmpl string, data interface{}, folder string) {
    tmpl = fmt.Sprintf("static/"+folder+"%s", tmpl)
    t, err := text_tmpl.ParseFiles(tmpl)
    //text_tmpl
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, data)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}
