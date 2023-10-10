package config

import "github.com/seal-io/walrus/utils/vars"

// TlsCertified is a flag to indicate whether the server is TLS certified.
var TlsCertified = vars.SetOnce[bool]{}
