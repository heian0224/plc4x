/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7Parameter struct {
	Child IS7ParameterChild
}

// The corresponding interface
type IS7Parameter interface {
	MessageType() uint8
	ParameterType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7ParameterParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7Parameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7ParameterChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7Parameter)
	GetTypeName() string
	IS7Parameter
}

func NewS7Parameter() *S7Parameter {
	return &S7Parameter{}
}

func CastS7Parameter(structType interface{}) *S7Parameter {
	castFunc := func(typ interface{}) *S7Parameter {
		if casted, ok := typ.(S7Parameter); ok {
			return &casted
		}
		if casted, ok := typ.(*S7Parameter); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7Parameter) GetTypeName() string {
	return "S7Parameter"
}

func (m *S7Parameter) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7Parameter) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *S7Parameter) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7Parameter) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterParse(readBuffer utils.ReadBuffer, messageType uint8) (*S7Parameter, error) {
	if pullErr := readBuffer.PullContext("S7Parameter"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType, _parameterTypeErr := readBuffer.ReadUint8("parameterType", 8)
	if _parameterTypeErr != nil {
		return nil, errors.Wrap(_parameterTypeErr, "Error parsing 'parameterType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7Parameter
	var typeSwitchError error
	switch {
	case parameterType == 0xF0: // S7ParameterSetupCommunication
		_parent, typeSwitchError = S7ParameterSetupCommunicationParse(readBuffer)
	case parameterType == 0x04 && messageType == 0x01: // S7ParameterReadVarRequest
		_parent, typeSwitchError = S7ParameterReadVarRequestParse(readBuffer)
	case parameterType == 0x04 && messageType == 0x03: // S7ParameterReadVarResponse
		_parent, typeSwitchError = S7ParameterReadVarResponseParse(readBuffer)
	case parameterType == 0x05 && messageType == 0x01: // S7ParameterWriteVarRequest
		_parent, typeSwitchError = S7ParameterWriteVarRequestParse(readBuffer)
	case parameterType == 0x05 && messageType == 0x03: // S7ParameterWriteVarResponse
		_parent, typeSwitchError = S7ParameterWriteVarResponseParse(readBuffer)
	case parameterType == 0x00 && messageType == 0x07: // S7ParameterUserData
		_parent, typeSwitchError = S7ParameterUserDataParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7Parameter"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *S7Parameter) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7Parameter) SerializeParent(writeBuffer utils.WriteBuffer, child IS7Parameter, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7Parameter"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType := uint8(child.ParameterType())
	_parameterTypeErr := writeBuffer.WriteUint8("parameterType", 8, (parameterType))

	if _parameterTypeErr != nil {
		return errors.Wrap(_parameterTypeErr, "Error serializing 'parameterType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Parameter"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7Parameter) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
