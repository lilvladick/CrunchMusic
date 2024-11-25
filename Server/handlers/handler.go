package handlers

// func Handler(query string, w Http.Response, r *Http.Request, dest interface{}, args ...interface{}) {
// 	jsonData, err := postgres.GetResultsJson(query, dest, args...)
// 	if err != nil {
// 		log.Printf("error %v", err)
// 		w.WriteHeader(Http.StatusInternalServerError)
// 		fmt.Println(Http.GetStatusText(Http.StatusInternalServerError))
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonData)))
// 	w.Write(jsonData)
// }
