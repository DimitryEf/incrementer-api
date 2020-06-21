package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Api struct {
	Inc *Incrementer
}

type Rpc struct {
	Func string   `json:"func"`
	Args []string `json:"args"`
}

func NewApi(inc *Incrementer) *Api {
	return &Api{
		Inc: inc,
	}
}

func (api *Api) Do(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rpc Rpc
	err = json.Unmarshal(b, &rpc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch rpc.Func {
	case "GetNumber":
		num, err := api.Inc.GetNumber()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		_, err = fmt.Fprintf(w, "%d", num)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		return
	case "IncrementNumber":
		err := api.Inc.IncrementNumber()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		return
	case "SetParams":
		maximumValue, err := strconv.Atoi(rpc.Args[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		stepValue, err := strconv.Atoi(rpc.Args[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		err = api.Inc.SetParams(int64(maximumValue), int64(stepValue))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		return

	}

}
