package main

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	pda "pda-enqueue-registration-documents/generated"
)

func main() {
	r := gin.Default()
	r.POST("/api/documents", enqueueDocuments)
	r.GET("/api/fintest", finTest)
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "9119"
	}

	r.Run(":" + port)
}

func finTest(c *gin.Context) {
	fin := &pda.Financial{}
	fin.Amount = 1.1
	fin.Operation = 0
	fin.Currency = 122
	fin.NodeId = 2
	fin.UserId = 23
	ts := timestamp.Timestamp{}
	ts.Seconds = 123
	fin.Date = &ts

	finJson, err := protojson.Marshal(fin)
	fmt.Fprintf(os.Stdout, "%s", finJson)
	if err != nil {
		log.Printf("Fin marshal error %v", err)
	}

	c.JSON(200, finJson)
}

func enqueueDocuments(c *gin.Context) {
	documents := &pda.Documents{}
	bodyRead, err := ioutil.ReadAll(c.Request.Body)
	err = protojson.Unmarshal(bodyRead, documents)

	if err != nil {
		log.Printf("Protojson unmarshal error: %v", err)
	}

	for i, s := range documents.GetDocuments() {
		log.Printf("Received document: #%v %v", i, s)
	}

	log.Printf("Protojson unmarshalled json: %v", documents.String())

	c.JSON(http.StatusOK, documents.String())
}
