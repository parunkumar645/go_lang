2.	package main
3.	
4.	import (
5.	    "encoding/json"
6.	    "fmt"
7.	    "net/http"
8.	)
9.	
10.	type RequestData struct {
11.	    Ev    string `json:"ev"`
12.	    Et    string `json:"et"`
13.	    ID    string `json:"id"`
14.	    UID   string `json:"uid"`
15.	    MID   string `json:"mid"`
16.	    T     string `json:"t"`
17.	    P     string `json:"p"`
18.	    L     string `json:"l"`
19.	    SC    string `json:"sc"`
20.	    ATRK1 string `json:"atrk1"`
21.	    ATRV1 string `json:"atrv1"`
22.	    ATRT1 string `json:"atrt1"`
23.	    ATRK2 string `json:"atrk2"`
24.	    ATRV2 string `json:"atrv2"`
25.	    ATRT2 string `json:"atrt2"`
26.	    UATRK1 string `json:"uatrk1"`
27.	    UATRV1 string `json:"uatrv1"`
28.	    UATRT1 string `json:"uatrt1"`
29.	    UATRK2 string `json:"uatrk2"`
30.	    UATRV2 string `json:"uatrv2"`
31.	    UATRT2 string `json:"uatrt2"`
32.	    UATRK3 string `json:"uatrk3"`
33.	    UATRV3 string `json:"uatrv3"`
34.	    UATRT3 string `json:"uatrt3"`
35.	}
36.	
37.	type NewRequestData struct {
38.	    Event           string `json:"event"`
39.	    EventType       string `json:"event_type"`
40.	    AppID           string `json:"app_id"`
41.	    UserID          string `json:"user_id"`
42.	    MessageID       string `json:"message_id"`
43.	    PageTitle       string `json:"page_title"`
44.	    PageURL         string `json:"page_url"`
45.	    BrowserLanguage string `json:"browser_language"`
46.	    ScreenSize      string `json:"screen_size"`
47.	    Attributes      struct {
48.	        FormVariant struct {
49.	            Value string `json:"value"`
50.	            Type  string `json:"type"`
51.	        } `json:"form_varient"`
52.	        Ref struct {
53.	            Value string `json:"value"`
54.	            Type  string `json:"type"`
55.	        } `json:"ref"`
56.	    } `json:"attributes"`
57.	    Traits struct {
58.	        Name struct {
59.	            Value string `json:"value"`
60.	            Type  string `json:"type"`
61.	        } `json:"name"`
62.	        Email struct {
63.	            Value string `json:"value"`
64.	            Type  string `json:"type"`
65.	        } `json:"email"`
66.	        Age struct {
67.	            Value string `json:"value"`
68.	            Type  string `json:"type"`
69.	        } `json:"age"`
70.	    } `json:"traits"`
71.	}
72.	
73.	func handleRequest(w http.ResponseWriter, r *http.Request) {
74.	    if r.Method != http.MethodPost {
75.	        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
76.	        return
77.	    }
78.	
79.	    var requestData RequestData
80.	    decoder := json.NewDecoder(r.Body)
81.	    err := decoder.Decode(&requestData)
82.	    if err != nil {
83.	        http.Error(w, "Error decoding JSON", http.StatusBadRequest)
84.	        return
85.	    }
86.	
87.	    // Transform data to the new format
88.	    newRequestData := NewRequestData{
89.	        Event:           requestData.Ev,
90.	        EventType:       requestData.Et,
91.	        AppID:           requestData.ID,
92.	        UserID:          requestData.UID,
93.	        MessageID:       requestData.MID,
94.	        PageTitle:       requestData.T,
95.	        PageURL:         requestData.P,
96.	        BrowserLanguage: requestData.L,
97.	        ScreenSize:      requestData.SC,
98.	        Attributes: struct {
99.	            FormVariant struct {
100.	                Value string `json:"value"`
101.	                Type  string `json:"type"`
102.	            } `json:"form_varient"`
103.	            Ref struct {
104.	                Value string `json:"value"`
105.	                Type  string `json:"type"`
106.	            } `json:"ref"`
107.	        }{
108.	            FormVariant: struct {
109.	                Value string `json:"value"`
110.	                Type  string `json:"type"`
111.	            }{
112.	                Value: requestData.ATRV1,
113.	                Type:  requestData.ATRT1,
114.	            },
115.	            Ref: struct {
116.	                Value string `json:"value"`
117.	                Type  string `json:"type"`
118.	            }{
119.	                Value: requestData.ATRV2,
120.	                Type:  requestData.ATRT2,
121.	            },
122.	        },
123.	        Traits: struct {
124.	            Name struct {
125.	                Value string `json:"value"`
126.	                Type  string `json:"type"`
127.	            } `json:"name"`
128.	            Email struct {
129.	                Value string `json:"value"`
130.	                Type  string `json:"type"`
131.	            } `json:"email"`
132.	            Age struct {
133.	                Value string `json:"value"`
134.	                Type  string `json:"type"`
135.	            } `json:"age"`
136.	        }{
137.	            Name: struct {
138.	                Value string `json:"value"`
139.	                Type  string `json:"type"`
140.	            }{
141.	                Value: requestData.UATRV1,
142.	                Type:  requestData.UATRT1,
143.	            },
144.	            Email: struct {
145.	                Value string `json:"value"`
146.	                Type  string `json:"type"`
147.	            }{
148.	                Value: requestData.UATRV2,
149.	                Type:  requestData.UATRT2,
150.	            },
151.	            Age: struct {
152.	                Value string `json:"value"`
153.	                Type  string `json:"type"`
154.	            }{
155.	                Value: requestData.UATRV3,
156.	                Type:  requestData.UATRT3,
157.	            },
158.	        },
159.	    }
160.	
161.	    // Do something with the transformed data
162.	    fmt.Printf("Received data in new format: %+v\n", newRequestData)
163.	
164.	    // Respond with a success message
165.	    w.WriteHeader(http.StatusOK)
166.	    w.Write([]byte("Data received successfully"))
167.	}
168.	
169.	func main() {
170.	    http.HandleFunc("/", handleRequest)
171.	    port := 8080
172.	    fmt.Printf("Server is listening on :%d...\n", port)
173.	    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
174.	    if err != nil {
175.	        fmt.Println("Error starting server:", err)
176.	    }
177.	}
178.	
