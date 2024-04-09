package errGroup

import "fmt"

var NotEnoughTickets = fmt.Errorf("not enough tickets")
var EventNotFound = fmt.Errorf("event not found")
