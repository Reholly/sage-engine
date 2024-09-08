package model

const (
	RollbackFailure Signal = iota
	RollbackSuccess
	CommitFailure
	CommitSuccess
	TimeoutFailure
	PrepareFailure
	PrepareSuccess
)

// Signal is a type of system information what different services can receive and send.
// As example, if transaction failed at some step application a client must do rollback and if it is successful
// send RollbackSuccess signal. Otherwise, it must send RollbackFailure.
// If commit at your step successful you must send CommitSuccess.
// Sage-engine allows you to use 7 types of signal. CommitSuccess, CommitFailure for commit signal information,
// RollbackFailure, RollbackSuccess for rollback signal information, TimeoutFailure for timeout at operation signal,
// PrepareFailure, PrepareSuccess for prepare quiz.
type Signal uint8
