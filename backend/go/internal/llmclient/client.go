package llmclient

var llmLimit = make(chan struct{}, 3)