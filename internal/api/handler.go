package api

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    "imgservice/internal/imgsrv"
)

const port string = ":8080"

const (
    apiType    = "api"
    apiVersion = "v1"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Method Not Allowed"))
        return
    }

    var img imgsrv.Image
    if err := json.NewDecoder(r.Body).Decode(&img); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    srv := imgsrv.NewImageProcessingService()
    if err := srv.Validate(ctx, &img); err != nil {
        http.Error(w, "Validation failed: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Validation succeeded"))
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Method Not Allowed"))
        return
    }

    var img imgsrv.Image
    if err := json.NewDecoder(r.Body).Decode(&img); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    srv := imgsrv.NewImageProcessingService()
    resp, err := srv.Process(ctx, &img)
    if err != nil {
        http.Error(w, "Processing failed: "+err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResp, err := json.Marshal(resp)
    if err != nil {
        http.Error(w, "Error creating response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
    
    // Simulate health fetch
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
}

func StartServer() {
	path := fmt.Sprintf("/%s/%s", apiType, apiVersion)
    http.HandleFunc(path + "/validate", ValidateHandler)
    http.HandleFunc(path + "/process", ProcessHandler)
    http.HandleFunc(path + "/health", HealthHandler)

    log.Printf("Listening on %s", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
