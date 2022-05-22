package main

import (
	"encoding/xml"
	"math/rand"
	"net/http"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	Say string `xml:",omitempty"`
}

func main() {
	http.HandleFunc("/twiml", twiml)
	http.ListenAndServe(":3000", nil)
}

func twiml(w http.ResponseWriter, r *http.Request) {
	twiml := TwiML{Say: getPositiveMsg()}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func getPositiveMsg() string {
	msgs := [5]string{"The last three or four reps is what makes the muscle grow. This area of pain divides a champion from someone who is not a champion.", "If you think lifting is dangerous, try being weak. Being weak is dangerous", "Whether you think you can, or you think you can’t, you’re right.", "A champion is someone who gets up when they can’t.", "If something stands between you and your success, move it. Never be denied."}
	idx := rand.Intn(5)
	msg := msgs[idx]
	return (msg)
}
