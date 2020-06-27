/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../errors"

// Declare Errors Details
var (
	ErrGPPacketTooShort = errors.New("GPPacketTooShort", "GP packet is empty or too short than standard header. It must include 44Byte header plus 16Byte min Payload")

	ErrSRPCServiceNotFound = errors.New("SRPCServiceNotFound", "Requested sRPC Service is out range of services in this version of service")
	ErrSRPCPayloadEmpty    = errors.New("SRPCPayloadEmpty", "Stream data payload can't be empty")

	ErrHTTPServiceNotFound = errors.New("HTTPServiceNotFound", "Requested HTTP Service is not found in this instance of app")

	ErrStreamPayloadEmpty     = errors.New("StreamPayloadEmpty", "Stream data payload can't be empty")
	ErrPacketArrivedAnterior  = errors.New("PacketArrivedAnterior", "New packet arrive before some expected packet arrived. Usually cause of drop packet detection or high latency occur for some packet")
	ErrPacketArrivedPosterior = errors.New("PacketArrivedPosterior", "New packet arrive after some expected packet arrived. Usually cause of drop packet detection or high latency occur for some packet")
)
