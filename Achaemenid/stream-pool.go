/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "sync"

// StreamPool set & get streams in a pool by ID!
type StreamPool struct {
	mutex               sync.Mutex         // TODO::: it is not efficient way and need more work
	p                   map[uint32]*Stream // key is Stream.ID
	freeIncomeStreamID  uint32
	freeOutcomeStreamID uint32
	totalOpenedStreams  uint32 // Manifest.TechnicalInfo.MaxStreamConnectionDaily
}

// Init initialize the pool
func (sp *StreamPool) Init() {
	sp.p = make(map[uint32]*Stream)
}

// RegisterStream save given Stream to pool
func (sp *StreamPool) RegisterStream(st *Stream) {
	sp.mutex.Lock()
	sp.p[st.ID] = st
	sp.mutex.Unlock()
}

// GetStreamByID returns Stream from pool if exists by given ID!
func (sp *StreamPool) GetStreamByID(id uint32) *Stream {
	// TODO::: Check stream isn't closed!!
	return sp.p[id]
}

// CloseStream delete given Stream from pool
func (sp *StreamPool) CloseStream(st *Stream) {
	sp.mutex.Lock()
	delete(sp.p, st.ID)
	sp.mutex.Unlock()
}
