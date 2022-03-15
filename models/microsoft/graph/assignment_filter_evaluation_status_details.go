package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterEvaluationStatusDetails provides operations to manage the compliance singleton.
type AssignmentFilterEvaluationStatusDetails struct {
    Entity
    // PayloadId on which filter has been applied.
    payloadId *string;
}
// NewAssignmentFilterEvaluationStatusDetails instantiates a new assignmentFilterEvaluationStatusDetails and sets the default values.
func NewAssignmentFilterEvaluationStatusDetails()(*AssignmentFilterEvaluationStatusDetails) {
    m := &AssignmentFilterEvaluationStatusDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAssignmentFilterEvaluationStatusDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterEvaluationStatusDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    return res
}
// GetPayloadId gets the payloadId property value. PayloadId on which filter has been applied.
func (m *AssignmentFilterEvaluationStatusDetails) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
func (m *AssignmentFilterEvaluationStatusDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterEvaluationStatusDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPayloadId sets the payloadId property value. PayloadId on which filter has been applied.
func (m *AssignmentFilterEvaluationStatusDetails) SetPayloadId(value *string)() {
    if m != nil {
        m.payloadId = value
    }
}
