package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	event struct {
		ID          int    `json:"ID"`
		Date        string `json:"Date"`
		Title       string `json:"Title"`
		Description string `json:"Description"`
		Type        string `json:"Type"`
		Priority    string `json:"Priority"`
		UserName    string `json:"Username"`
	}
	allEvents []event
)

//Post a Json struct of type event
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data")
	}

	json.Unmarshal(reqBody, &newEvent)
	AddEvent(password, newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

//GET all events
func getAllEvents(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(events)
}

//GET a single event.
func getSingleEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if strconv.Itoa(singleEvent.ID) == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

//Update event list.
func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please enter updated data")
	}
	json.Unmarshal(reqBody, &updatedEvent)
	for i, singleEvent := range events {
		if strconv.Itoa(singleEvent.ID) == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			singleEvent.Type = updatedEvent.Type
			singleEvent.Priority = updatedEvent.Priority
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

//Remove event
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	var eventID = mux.Vars(r)["id"]
	check := true
	for i, singleEvent := range events {
		if strconv.Itoa(singleEvent.ID) == eventID {
			RemoveEvent(password, singleEvent.ID)
			events = append(events[:i], events[i+1:]...)
			check = false
			fmt.Fprintf(w, "The event with id %v has been successfully removed", eventID)
		}

	}

	if check == true {
		temp, err := strconv.Atoi(eventID)
		if err == nil {
			ev := FetchEventByID(password, temp)
			//error because it doesn't return anything, fix by returning a event object
			if ev.Title != "" {
				RemoveEvent(password, temp)
				fmt.Fprintf(w, "The event with id %v has been successfully removed", eventID)
				check = false
			}
		}
	}
	if check == true {
		fmt.Print("sorry this event doesn't exist")
	}
}
