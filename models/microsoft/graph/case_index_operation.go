package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CaseIndexOperation 
type CaseIndexOperation struct {
    CaseOperation
}
// NewCaseIndexOperation instantiates a new caseIndexOperation and sets the default values.
func NewCaseIndexOperation()(*CaseIndexOperation) {
    m := &CaseIndexOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseIndexOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    return res
}
func (m *CaseIndexOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CaseIndexOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
