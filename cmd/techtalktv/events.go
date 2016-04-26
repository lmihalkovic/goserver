package main

import "sort"

// Model ----------------------------------------------------------------------

type Event struct {
	// Unique event identifier
	Id string     	`json:"id"`
	// Event given name
	Name string     `json:"name"`

	/* Internal */
	Index string	`json:"-"`
	Base string	`json:"-"`
}

type Events []Event
type EventsIndex map[string]Event

func SortEventsById(e Events) {
	sort.Sort(eventsSortedById(e))
}
func SortEventsByName(e Events) {
	sort.Sort(eventsSortedByName(e))
}

// helpers

type eventsSortedById []Event
func (s eventsSortedById) Len() int 		{ return len(s) }
func (s eventsSortedById) Less(i, j int) bool 	{ return s[i].Id < s[j].Id }
func (s eventsSortedById) Swap(i, j int) 	{ s[i], s[j] = s[j], s[i] }

type eventsSortedByName []Event
func (s eventsSortedByName) Len() int 		{ return len(s) }
func (s eventsSortedByName) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s eventsSortedByName) Swap(i, j int) 	{ s[i], s[j] = s[j], s[i] }

// ----------------------------------------------------------------------------
