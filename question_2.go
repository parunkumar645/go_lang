package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
    Ev    string `json:"ev"`
    Et    string `json:"et"`
    ID    string `json:"id"`
    UID   string `json:"uid"`
    MID   string `json:"mid"`
    T     string `json:"t"`
    P     string `json:"p"`
    L     string `json:"l"`
    SC    string `json:"sc"`
    ATRK1 string `json:"atrk1"`
    ATRV1 string `json:"atrv1"`
    ATRT1 string `json:"atrt1"`
    ATRK2 string `json:"atrk2"`
    ATRV2 string `json:"atrv2"`
    ATRT2 string `json:"atrt2"`
    UATRK1 string `json:"uatrk1"`
    UATRV1 string `json:"uatrv1"`
    UATRT1 string `json:"uatrt1"`
    UATRK2 string `json:"uatrk2"`
    UATRV2 string `json:"uatrv2"`
    UATRT2 string `json:"uatrt2"`
    UATRK3 string `json:"uatrk3"`
    UATRV3 string `json:"uatrv3"`
    UATRT3 string `json:"uatrt3"`
}

type NewRequestData struct {
    Event           string `json:"event"`
    EventType       string `json:"event_type"`
    AppID           string `json:"app_id"`
    UserID          string `json:"user_id"`
    MessageID       string `json:"message_id"`
    PageTitle       string `json:"page_title"`
    PageURL         string `json:"page_url"`
    BrowserLanguage string `json:"browser_language"`
    ScreenSize      string `json:"screen_size"`
    Attributes      map[string]struct {
        Value string `json:"value"`
        Type  string `json:"type"`
    } `json:"attributes"`
    Traits map[string]struct {
        Value string `json:"value"`
        Type  string `json:"type"`
    } `json:"traits"`
}

func worker(input chan RequestData) {
    for requestData := range input {
        // Transform data to the new format
        newRequestData := NewRequestData{
            Event:           requestData.Ev,
            EventType:       requestData.Et,
            AppID:           requestData.ID,
            UserID:          requestData.UID,
            MessageID:       requestData.MID,
            PageTitle:       requestData.T,
            PageURL:         requestData.P,
            BrowserLanguage: requestData.L,
            ScreenSize:      requestData.SC,
            Attributes: map[string]struct {
                Value string `json:"value"`
                Type  string `json:"type"`
            }{
                requestData.ATRK1: {
                    Value: requestData.ATRV1,
                    Type:  requestData.ATRT1,
                },
                requestData.ATRK2: {
                    Value: requestData.ATRV2,
                    Type:  requestData.ATRT2,
                },
            },
            Traits: map[string]struct {
                Value string `json:"value"`
                Type  string `json:"type"`
            }{
                requestData.UATRK1: {
                    Value: requestData.UATRV1,
                    Type:  requestData.UATRT1,
                },
                requestData.UATRK2: {
                    Value: requestData.UATRV2,
                    Type:  requestData.UATRT2,
                },
                requestData.UATRK3: {
                    Value: requestData.UATRV3,
                    Type:  requestData.UATRT3,
                },
            },
        }

        // Do something with the transformed data (e.g., send it to another service)
        fmt.Printf("Received data in new format: %+v\n", newRequestData)
    }
}

func main() {
    inputChannel := make(chan RequestData, 10) // Buffer the channel to handle multiple requests

    // Start worker goroutine
    go worker(inputChannel)

    // HTTP handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        var requestData RequestData
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&requestData)
        if err != nil {
            http.Error(w, "Error decoding JSON", http.StatusBadRequest)
            return
        }

        // Send the request data to the worker through the channel
        inputChannel <- requestData

        // Respond with a success message
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Data received successfully"))
    })

    port := 8080
    fmt.Printf("Server is listening on :%d...\n", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

