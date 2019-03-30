package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "strconv"
    text_tmpl "text/template"
)

var (
    currentPosition = -1
)

const (
    BaseUrl      = "http://192.168.86.211:8070"
)

func main() {
    defineEndpoints()
}

func (mm *moverManager)on(w http.ResponseWriter, r *http.Request) {
    mm.MoverObj.On()
    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func (mm *moverManager)off(w http.ResponseWriter, r *http.Request) {
    mm.MoverObj.Off()
    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func (mm *moverManager)stop(w http.ResponseWriter, r *http.Request) {
    mm.MoverObj.Stop()
    jsonReturn := map[string]interface{}{"version": "1.0", "shouldEndSession": true}
    json.NewEncoder(w).Encode(jsonReturn)
}

func (mm *moverManager)position(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    position := vars["number"]
    var err error
    var tempPosition int64
    tempPosition, err = strconv.ParseInt(position, 10, 64)
    if err != nil {
        fmt.Println(err.Error())
        w.WriteHeader(http.StatusForbidden)
    }
    var positionInt int
    positionInt = int(tempPosition)
    //positionMovement := mm.DetermineTableMovement(positionInt)

    //currentPosition = currentPosition + positionMovement
    mm.MoverObj.Position(positionInt)
    jsonReturn := map[string]interface{}{"version": "1.0"}
    json.NewEncoder(w).Encode(jsonReturn)
}

/*
Controller of Mover methods
 */
type moverManager struct {
    MoverObj Mover
}

func defineEndpoints() {
    router := mux.NewRouter()
    //router.HandleFunc("/", rootMethod).Methods("GET")

    mm := moverManager{}
    mm.MoverObj = new(Table).InitialSetup()

    router.HandleFunc("/up", mm.on).Methods("GET")
    router.HandleFunc("/down", mm.off).Methods("GET")
    router.HandleFunc("/stop", mm.stop).Methods("GET")
    router.HandleFunc("/position/{number}", mm.position).Methods("GET")
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
    template_admin, err := renderText("index.html", data, "")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        err = template_admin.Execute(w, data)
        if err != nil {
            log.Print("template executing error: ", err)
        }
    }
}

//this function determine how many positions the table has to be move and the way. Positive goes up and negative goes down


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

    var template *text_tmpl.Template
    var err error
    var parseVars map[string]interface{} = nil
    if filename == "test.js" {
        folder = "js/"
        parseVars = map[string]interface{}{"baseUrl": BaseUrl}
        template, err =renderText(filename, parseVars, folder)
    } else if filename == "index.js" {
        folder = "js/"
        parseVars = map[string]interface{}{"Asi": BaseUrl}
        template, err = renderText(filename, parseVars, folder)

    } else {
        template, err = renderText(filename, nil, folder)
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        err = template.Execute(writer, parseVars)
        if err != nil {
            log.Print("template executing error: ", err)
        }
    }
}

//render file content based on a template
func renderText(tmpl string, data interface{}, folder string) (*text_tmpl.Template, error){
    tmpl = fmt.Sprintf("static/"+folder+"%s", tmpl)
    t, err := text_tmpl.ParseFiles(tmpl)
    //text_tmpl
    if err != nil {
        //log.Print("template parsing error: ", err)
        return nil, err
    }
    return t, nil
    /*err = t.Execute(w, data)
    if err != nil {
        log.Print("template executing error: ", err)
    }*/
}
