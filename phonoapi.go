package main
import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"strings"
	"strconv"
	"time"
	"math/rand"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	router.HandleFunc("/api/v1/phonoapi/{word}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]
		fmt.Println(word)
		phonoapi(w, r, word)
	})

	
	fmt.Println("Greetings Human\nI come in peace")
}
