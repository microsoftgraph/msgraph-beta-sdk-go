package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtectionPolicy 
type InformationProtectionPolicy struct {
    Entity
    // 
    labels []InformationProtectionLabel;
}
// NewInformationProtectionPolicy instantiates a new informationProtectionPolicy and sets the default values.
func NewInformationProtectionPolicy()(*InformationProtectionPolicy) {
    m := &InformationProtectionPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetLabels gets the labels property value. 
func (m *InformationProtectionPolicy) GetLabels()([]InformationProtectionLabel) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationProtectionLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*InformationProtectionLabel))
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
func (m *InformationProtectionPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InformationProtectionPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLabels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabels sets the labels property value. 
func (m *InformationProtectionPolicy) SetLabels(value []InformationProtectionLabel)() {
    if m != nil {
        m.labels = value
    }
}
