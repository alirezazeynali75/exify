package payment

import "errors"


var ErrNoProviderFound = errors.New("there is no provider to handle the request")
var ErrNoDepositAgentIsAvailable = errors.New("there is no deposit agent available")