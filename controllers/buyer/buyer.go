package buyer

import (
	"log"

	"net/http"
	"io/ioutil"
)

// AddBuyer adds new buyer.
func AddBuyer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}



	//w.Write([]byte(message))
}
