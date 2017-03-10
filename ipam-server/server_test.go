package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"strings"
)

var messageIndex = "/messages"
var messageGet = "/messages/"
var testMessage = "My test message"

var testMessageAdd = Message{1, testMessage}

func TestMessageAddNew(t *testing.T) {

	result, _ := messageAddNew(testMessage)

	expected := testMessageAdd

	if result.Id != expected.Id {
		t.Errorf("Expected the ID '%d' got '%d'\n", expected.Id, result.Id)
	}

	if result.Message != expected.Message {
		t.Errorf("Expected the message '%s' got '%s'\n", expected.Message, result.Id)
	}

}

func TestMessageFindById(t *testing.T) {

	result := messageFindById(1)

	expected := testMessage

	if result != expected {
		t.Errorf("Expected the message '%s' got '%s'\n", expected, result)
	}

	result = messageFindById(2)

	expected = ""

	if result != expected {
		t.Errorf("Expected the message '%s' got '%s'\n", expected, result)
	}
}

func TestMessagesIndex(t *testing.T) {

	server := httptest.NewServer(apiHandler())
	defer server.Close()

	messagesIndexUrl := server.URL + messageIndex
	resp, err := http.Get(messagesIndexUrl)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Got non-200 response: %d\n", resp.StatusCode)
	}

	expected := fmt.Sprintf("[1] %s\n", testMessage)

	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s' got '%s'\n", expected, string(actual))
	}

}

func TestMessagesAdd(t *testing.T) {

	server := httptest.NewServer(apiHandler())
	defer server.Close()

	messagesIndexUrl := server.URL + messageIndex
	resp, err := http.Post(messagesIndexUrl, "text/html", strings.NewReader(testMessage))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("Got non-201 response: %d\n", resp.StatusCode)
	}

	expected := fmt.Sprint("\n{\"id\":2}\n")

	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s' got '%s'\n", expected, string(actual))
	}

}

func TestMessagesGetById(t *testing.T) {

	server := httptest.NewServer(apiHandler())
	defer server.Close()

	messagesIndexUrl := server.URL + messageGet + "2"
	resp, err := http.Get(messagesIndexUrl)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Got non-200 response: %d\n", resp.StatusCode)
	}

	expected := fmt.Sprintf("\n%s\n", testMessage)

	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s' got '%s'\n", expected, string(actual))
	}


}
