package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentFilterEvaluationStatusDetails struct {
    Entity
    payloadId *string;
}
func NewAssignmentFilterEvaluationStatusDetails()(*AssignmentFilterEvaluationStatusDetails) {
    m := &AssignmentFilterEvaluationStatusDetails{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AssignmentFilterEvaluationStatusDetails) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
func (m *AssignmentFilterEvaluationStatusDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayloadId(val)
        return nil
    }
    return res
}
func (m *AssignmentFilterEvaluationStatusDetails) IsNil()(bool) {
    return m == nil
}
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
func (m *AssignmentFilterEvaluationStatusDetails) SetPayloadId(value *string)() {
    m.payloadId = value
}
