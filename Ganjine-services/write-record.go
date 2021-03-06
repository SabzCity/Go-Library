/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// WriteRecordService store details about WriteRecord service
var WriteRecordService = achaemenid.Service{
	ID:                3836795965,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "WriteRecord",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `write some part of a record! Don't use this service until you force to use!
		Recalculate checksum do in database server that is not so efficient!`,
	},
	TAGS: []string{""},

	SRPCHandler: WriteRecordSRPC,
}

// WriteRecordSRPC is sRPC handler of WriteRecord service.
func WriteRecordSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &WriteRecordReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = WriteRecord(req)
}

// WriteRecordReq is request structure of WriteRecord()
type WriteRecordReq struct {
	Type     requestType
	RecordID [32]byte
	Offset   uint64 // start location of write data
	Data     []byte
}

// WriteRecord write some part of a record! Don't use this service until you force to use!
func WriteRecord(req *WriteRecordReq) (err error) {
	if req.Type == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		var i uint8
		for i = 1; i < cluster.Manifest.TotalZones; i++ {
			// Make new request-response streams
			var st *achaemenid.Stream
			st, err = cluster.Replications.Zones[i].Nodes[cluster.Node.ID].Conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			// Set WriteRecord ServiceID
			st.Service = &achaemenid.Service{ID: 3836795965}
			st.OutcomePayload = reqEncoded

			err = achaemenid.SrpcOutcomeRequestHandler(server, st)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = st.Err
		}
	}

	// Do for i=0 as local node
	err = persiaos.WriteStorageRecord(req.RecordID, req.Offset, req.Data)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *WriteRecordReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	copy(req.RecordID[:], buf[1:])
	req.Offset = syllab.GetUInt64(buf, 33)
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	req.Data = buf[req.syllabStackLen():]
	return
}

// SyllabEncoder encode req to buf
func (req *WriteRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.RecordID[:])
	syllab.SetUInt64(buf, 37, req.Offset)
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 45, res.syllabStackLen())
	// syllab.SetUInt32(buf, 49, uint32(len(res.Record)))
	copy(buf[45:], req.Data[:])
	return
}

func (req *WriteRecordReq) syllabStackLen() (ln uint32) {
	return 41
}

func (req *WriteRecordReq) syllabHeapLen() (ln uint32) {
	ln = uint32(len(req.Data))
	return
}

func (req *WriteRecordReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
