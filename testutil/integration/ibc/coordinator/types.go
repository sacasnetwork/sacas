// Copyright SacDev0S . (SacaS)
package coordinator

// Endpoint defines the identifiers for a chain's client, connection, and channel.
type Endpoint struct {
	ChainID      string
	ClientID     string
	ConnectionID string
	ChannelID    string
	PortID       string
}

// IBCConnection defines the connection between two chains.
type IBCConnection struct {
	EndpointA Endpoint
	EndpointB Endpoint
}
