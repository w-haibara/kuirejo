package main

import ()

type FailState struct {
	CommonState
	Cause string `json:"Cause"`
	Error string `json:"Error"`
}
