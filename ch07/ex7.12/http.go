// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var itemTable = template.Must(template.New("itemTable").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <title>ex7.12</title>
    <style>
      table {
     border-collapse: collapse;
      }
      td, th {
     border: solid 1px;
     padding: 0.5em;
        text-align: right;
      }
    </style>
  </head>
  <body>
    <table>
      <tr>
        <th>Name</th>
        <th>Price</th>
      </tr>
      {{range $k, $v := .}}
      <tr>
        <td>{{$k}}</td>
        <td>{{$v}}</td>
      </tr>
      {{end}}
    <table>
  </body>
</html>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	itemTable.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}
	if _, exist := db[item]; exist {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Item %s already exsits.\n", item)
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, "Successfully created item. %s: %s", item, db[item])
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}
	if _, ok := db[item]; ok {
		db[item] = dollars(price)
		fmt.Fprintf(w, "Successfully updated item. %s: %s\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "Successfully deleted item %s\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
