// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package baggage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type BaggageRestrictionManager interface {
	// getBaggageRestrictions retrieves the baggage restrictions for a specific service.
	// Usually, baggageRestrictions apply to all services however there may be situations
	// where a baggageKey might only be allowed to be set by a specific service.
	//
	// Parameters:
	//  - ServiceName
	GetBaggageRestrictions(serviceName string) (r []*BaggageRestriction, err error)
}

type BaggageRestrictionManagerClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewBaggageRestrictionManagerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BaggageRestrictionManagerClient {
	return &BaggageRestrictionManagerClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewBaggageRestrictionManagerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BaggageRestrictionManagerClient {
	return &BaggageRestrictionManagerClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// getBaggageRestrictions retrieves the baggage restrictions for a specific service.
// Usually, baggageRestrictions apply to all services however there may be situations
// where a baggageKey might only be allowed to be set by a specific service.
//
// Parameters:
//  - ServiceName
func (p *BaggageRestrictionManagerClient) GetBaggageRestrictions(serviceName string) (r []*BaggageRestriction, err error) {
	if err = p.sendGetBaggageRestrictions(serviceName); err != nil {
		return
	}
	return p.recvGetBaggageRestrictions()
}

func (p *BaggageRestrictionManagerClient) sendGetBaggageRestrictions(serviceName string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getBaggageRestrictions", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := BaggageRestrictionManagerGetBaggageRestrictionsArgs{
		ServiceName: serviceName,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *BaggageRestrictionManagerClient) recvGetBaggageRestrictions() (value []*BaggageRestriction, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getBaggageRestrictions" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getBaggageRestrictions failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getBaggageRestrictions failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getBaggageRestrictions failed: invalid message type")
		return
	}
	result := BaggageRestrictionManagerGetBaggageRestrictionsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type BaggageRestrictionManagerProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      BaggageRestrictionManager
}

func (p *BaggageRestrictionManagerProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *BaggageRestrictionManagerProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *BaggageRestrictionManagerProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewBaggageRestrictionManagerProcessor(handler BaggageRestrictionManager) *BaggageRestrictionManagerProcessor {

	self2 := &BaggageRestrictionManagerProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["getBaggageRestrictions"] = &baggageRestrictionManagerProcessorGetBaggageRestrictions{handler: handler}
	return self2
}

func (p *BaggageRestrictionManagerProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type baggageRestrictionManagerProcessorGetBaggageRestrictions struct {
	handler BaggageRestrictionManager
}

func (p *baggageRestrictionManagerProcessorGetBaggageRestrictions) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := BaggageRestrictionManagerGetBaggageRestrictionsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getBaggageRestrictions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := BaggageRestrictionManagerGetBaggageRestrictionsResult{}
	var retval []*BaggageRestriction
	var err2 error
	if retval, err2 = p.handler.GetBaggageRestrictions(args.ServiceName); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getBaggageRestrictions: "+err2.Error())
		oprot.WriteMessageBegin("getBaggageRestrictions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getBaggageRestrictions", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - ServiceName
type BaggageRestrictionManagerGetBaggageRestrictionsArgs struct {
	ServiceName string `thrift:"serviceName,1" json:"serviceName"`
}

func NewBaggageRestrictionManagerGetBaggageRestrictionsArgs() *BaggageRestrictionManagerGetBaggageRestrictionsArgs {
	return &BaggageRestrictionManagerGetBaggageRestrictionsArgs{}
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) GetServiceName() string {
	return p.ServiceName
}
func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ServiceName = v
	}
	return nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getBaggageRestrictions_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serviceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.ServiceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serviceName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serviceName: ", p), err)
	}
	return err
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaggageRestrictionManagerGetBaggageRestrictionsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type BaggageRestrictionManagerGetBaggageRestrictionsResult struct {
	Success []*BaggageRestriction `thrift:"success,0" json:"success,omitempty"`
}

func NewBaggageRestrictionManagerGetBaggageRestrictionsResult() *BaggageRestrictionManagerGetBaggageRestrictionsResult {
	return &BaggageRestrictionManagerGetBaggageRestrictionsResult{}
}

var BaggageRestrictionManagerGetBaggageRestrictionsResult_Success_DEFAULT []*BaggageRestriction

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) GetSuccess() []*BaggageRestriction {
	return p.Success
}
func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*BaggageRestriction, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &BaggageRestriction{}
		if err := _elem4.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
		}
		p.Success = append(p.Success, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getBaggageRestrictions_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaggageRestrictionManagerGetBaggageRestrictionsResult(%+v)", *p)
}
