// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package filtergen

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type IDType int64

const (
	IDType_origin_id IDType = 1
	IDType_md5_id    IDType = 2
	IDType_sha1_id   IDType = 3
)

func (p IDType) String() string {
	switch p {
	case IDType_origin_id:
		return "origin_id"
	case IDType_md5_id:
		return "md5_id"
	case IDType_sha1_id:
		return "sha1_id"
	}
	return "<UNSET>"
}

func IDTypeFromString(s string) (IDType, error) {
	switch s {
	case "origin_id":
		return IDType_origin_id, nil
	case "md5_id":
		return IDType_md5_id, nil
	case "sha1_id":
		return IDType_sha1_id, nil
	}
	return IDType(0), fmt.Errorf("not a valid IDType string")
}

func IDTypePtr(v IDType) *IDType { return &v }

func (p IDType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *IDType) UnmarshalText(text []byte) error {
	q, err := IDTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - ID
//  - IdType
//  - IP
//  - SlotIds
//  - Ua
//  - SourceId
//  - Geo
//  - DevType
//  - IdDt
type AntiSpamRequest struct {
	ID       string   `thrift:"id,1" json:"id"`
	IdType   IDType   `thrift:"idType,2" json:"idType"`
	IP       string   `thrift:"ip,3" json:"ip"`
	SlotIds  []string `thrift:"slotIds,4" json:"slotIds"`
	Ua       string   `thrift:"ua,5" json:"ua"`
	SourceId string   `thrift:"sourceId,6" json:"sourceId"`
	Geo      string   `thrift:"geo,7" json:"geo"`
	DevType  int32    `thrift:"devType,8" json:"devType"`
	IdDt     int32    `thrift:"idDt,9" json:"idDt"`
}

func NewAntiSpamRequest() *AntiSpamRequest {
	return &AntiSpamRequest{}
}

func (p *AntiSpamRequest) GetID() string {
	return p.ID
}

func (p *AntiSpamRequest) GetIdType() IDType {
	return p.IdType
}

func (p *AntiSpamRequest) GetIP() string {
	return p.IP
}

func (p *AntiSpamRequest) GetSlotIds() []string {
	return p.SlotIds
}

func (p *AntiSpamRequest) GetUa() string {
	return p.Ua
}

func (p *AntiSpamRequest) GetSourceId() string {
	return p.SourceId
}

func (p *AntiSpamRequest) GetGeo() string {
	return p.Geo
}

func (p *AntiSpamRequest) GetDevType() int32 {
	return p.DevType
}

func (p *AntiSpamRequest) GetIdDt() int32 {
	return p.IdDt
}
func (p *AntiSpamRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.readField9(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AntiSpamRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *AntiSpamRequest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := IDType(v)
		p.IdType = temp
	}
	return nil
}

func (p *AntiSpamRequest) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *AntiSpamRequest) readField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.SlotIds = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.SlotIds = append(p.SlotIds, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AntiSpamRequest) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Ua = v
	}
	return nil
}

func (p *AntiSpamRequest) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.SourceId = v
	}
	return nil
}

func (p *AntiSpamRequest) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Geo = v
	}
	return nil
}

func (p *AntiSpamRequest) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.DevType = v
	}
	return nil
}

func (p *AntiSpamRequest) readField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.IdDt = v
	}
	return nil
}

func (p *AntiSpamRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AntiSpamRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AntiSpamRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteString(string(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("idType", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:idType: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.IdType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.idType (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:idType: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:ip: ", p), err)
	}
	if err := oprot.WriteString(string(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:ip: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("slotIds", thrift.LIST, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:slotIds: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.SlotIds)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.SlotIds {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:slotIds: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ua", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:ua: ", p), err)
	}
	if err := oprot.WriteString(string(p.Ua)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ua (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:ua: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sourceId", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:sourceId: ", p), err)
	}
	if err := oprot.WriteString(string(p.SourceId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sourceId (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:sourceId: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("geo", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:geo: ", p), err)
	}
	if err := oprot.WriteString(string(p.Geo)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.geo (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:geo: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("devType", thrift.I32, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:devType: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.DevType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.devType (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:devType: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("idDt", thrift.I32, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:idDt: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.IdDt)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.idDt (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:idDt: ", p), err)
	}
	return err
}

func (p *AntiSpamRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AntiSpamRequest(%+v)", *p)
}

// Attributes:
//  - Legals
//  - Reasons
type AntiSpamResponse struct {
	Legals  []bool   `thrift:"legals,1" json:"legals"`
	Reasons []string `thrift:"reasons,2" json:"reasons"`
}

func NewAntiSpamResponse() *AntiSpamResponse {
	return &AntiSpamResponse{}
}

func (p *AntiSpamResponse) GetLegals() []bool {
	return p.Legals
}

func (p *AntiSpamResponse) GetReasons() []string {
	return p.Reasons
}
func (p *AntiSpamResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AntiSpamResponse) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]bool, 0, size)
	p.Legals = tSlice
	for i := 0; i < size; i++ {
		var _elem1 bool
		if v, err := iprot.ReadBool(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.Legals = append(p.Legals, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AntiSpamResponse) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Reasons = tSlice
	for i := 0; i < size; i++ {
		var _elem2 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem2 = v
		}
		p.Reasons = append(p.Reasons, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AntiSpamResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AntiSpamResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AntiSpamResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("legals", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:legals: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.BOOL, len(p.Legals)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Legals {
		if err := oprot.WriteBool(bool(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:legals: ", p), err)
	}
	return err
}

func (p *AntiSpamResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("reasons", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:reasons: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Reasons)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Reasons {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:reasons: ", p), err)
	}
	return err
}

func (p *AntiSpamResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AntiSpamResponse(%+v)", *p)
}
