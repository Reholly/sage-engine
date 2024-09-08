package model

const (
	OnCommitFailure   Behaviour = "onCommitFailure"
	OnCommitSuccess   Behaviour = "onCommitSuccess"
	OnRollbackFailure Behaviour = "onRollbackFailure"
	OnRollbackSuccess Behaviour = "onRollbackSuccess"
	OnTimeout         Behaviour = "onTimeout"
	OnPrepareFailure  Behaviour = "onPrepareFailure"
	OnPrepareSuccess  Behaviour = "onPrepareSuccess"
)

const (
	Log   Action = "log"
	Alert Action = "alert"
	Retry Action = "retry"
)

// Behaviour describes Action when accident occurs. As example,
// you can choose behaviour when Rollback at some step failed. You can add
// Log, Alert, Retry if you configured alert-bot.
// or log server.
type Behaviour string

// Action describes possible actions to be taken in response to an accident.
type Action string
