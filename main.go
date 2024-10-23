package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type LanguageData struct {
	Translations map[string]string
}

var translations = map[string]LanguageData{
	"en": {
		Translations: map[string]string{
			"greeting": "Hello can you help me in write code used Go and add featured multilanguage",
		},
	},
	"id": {
		Translations: map[string]string{
			"greeting": "Halo, bisakah Anda membantu saya dalam menulis kode menggunakan Go dan menambahkan fitur multibahasa",
		},
	},
	"korea": {
		Translations: map[string]string{
			"greeting": "안녕하세요, Go를 사용하여 코드를 작성하고 특집 다국어를 추가하는 데 도움을 주실 수 있나요?",
		},
	},
}

func getGreeting(w http.ResponseWriter, r *http.Request) {
	lang := mux.Vars(r)["lang"]

	if data, ok := translations[lang]; ok {
		response := map[string]string{"message": data.Translations["greeting"]}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Language not supported", http.StatusNotFound)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greeting/{lang}", getGreeting).Methods("GET")

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
