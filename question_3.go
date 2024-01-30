package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InputData struct {
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
	ATRK3 string `json:"atrk3"`
	ATRV3 string `json:"atrv3"`
	ATRT3 string `json:"atrt3"`
	ATRK4 string `json:"atrk4"`
	ATRV4 string `json:"atrv4"`
	ATRT4 string `json:"atrt4"`
	UATRK1 string `json:"uatrk1"`
	UATRV1 string `json:"uatrv1"`
	UATRT1 string `json:"uatrt1"`
	UATRK2 string `json:"uatrk2"`
	UATRV2 string `json:"uatrv2"`
	UATRT2 string `json:"uatrt2"`
	UATRK3 string `json:"uatrk3"`
	UATRV3 string `json:"uatrv3"`
	UATRT3 string `json:"uatrt3"`
	UATRK4 string `json:"uatrk4"`
	UATRV4 string `json:"uatrv4"`
	UATRT4 string `json:"uatrt4"`
	UATRK5 string `json:"uatrk5"`
	UATRV5 string `json:"uatrv5"`
	UATRT5 string `json:"uatrt5"`
	UATRK6 string `json:"uatrk6"`
	UATRV6 string `json:"uatrv6"`
	UATRT6 string `json:"uatrt6"`
}

type OutputData struct {
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

func convertToOutputFormat(inputData InputData) OutputData {
	return OutputData{
		Event:           inputData.Ev,
		EventType:       inputData.Et,
		AppID:           inputData.ID,
		UserID:          inputData.UID,
		MessageID:       inputData.MID,
		PageTitle:       inputData.T,
		PageURL:         inputData.P,
		BrowserLanguage: inputData.L,
		ScreenSize:      inputData.SC,
		Attributes: map[string]struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		}{
			"button_text": {
				Value: inputData.ATRV1,
				Type:  inputData.ATRT1,
			},
			"color_variation": {
				Value: inputData.ATRV2,
				Type:  inputData.ATRT2,
			},
			"page_path": {
				Value: inputData.ATRV3,
				Type:  inputData.ATRT3,
			},
			"source": {
				Value: inputData.ATRV4,
				Type:  inputData.ATRT4,
			},
		},
		Traits: map[string]struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		}{
			"user_score": {
				Value: inputData.UATRV1,
				Type:  inputData.UATRT1,
			},
			"gender": {
				Value: inputData.UATRV2,
				Type:  inputData.UATRT2,
			},
			"tracking_code": {
				Value: inputData.UATRV3,
				Type:  inputData.UATRT3,
			},
			"phone": {
				Value: inputData.UATRV4,
				Type:  inputData.UATRT4,
			},
			"coupon_clicked": {
				Value: inputData.UATRV5,
				Type:  inputData.UATRT5,
			},
			"opt_out": {
				Value: inputData.UATRV6,
				Type:  inputData.UATRT6,
			},
		},
	}
}

func worker(input chan InputData) {
	for inputData := range input {
		// Convert the message to the desired output format
		outputData := convertToOutputFormat(inputData)

		// Print the converted output
		outputJSON, err := json.MarshalIndent(outputData, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling output data:", err)
			continue
		}

		fmt.Println("Converted Output:")
		fmt.Println(string(outputJSON))
	}
}

func main() {
	inputChannel := make(chan InputData, 10) // Buffer the channel to handle multiple requests

	// Start worker goroutine
	go worker(inputChannel)

	// HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var inputData InputData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&inputData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Send the input data to the worker through the channel
		inputChannel <- inputData

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
