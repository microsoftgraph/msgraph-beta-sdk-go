package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Approval 
type Approval struct {
    Entity
    // 
    steps []ApprovalStep;
}
// NewApproval instantiates a new approval and sets the default values.
func NewApproval()(*Approval) {
    m := &Approval{
        Entity: *NewEntity(),
    }
    return m
}
// GetSteps gets the steps property value. 
func (m *Approval) GetSteps()([]ApprovalStep) {
    if m == nil {
        return nil
    } else {
        return m.steps
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Approval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["steps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApprovalStep() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalStep, len(val))
            for i, v := range val {
                res[i] = *(v.(*ApprovalStep))
            }
            m.SetSteps(res)
        }
        return nil
    }
    return res
}
func (m *Approval) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Approval) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSteps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSteps()))
        for i, v := range m.GetSteps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("steps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSteps sets the steps property value. 
func (m *Approval) SetSteps(value []ApprovalStep)() {
    if m != nil {
        m.steps = value
    }
}
