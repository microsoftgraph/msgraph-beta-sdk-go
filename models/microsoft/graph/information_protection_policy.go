package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtectionPolicy provides operations to manage the compliance singleton.
type InformationProtectionPolicy struct {
    Entity
    // 
    labels []InformationProtectionLabelable;
}
// NewInformationProtectionPolicy instantiates a new informationProtectionPolicy and sets the default values.
func NewInformationProtectionPolicy()(*InformationProtectionPolicy) {
    m := &InformationProtectionPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInformationProtectionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInformationProtectionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionLabelable, len(val))
            for i, v := range val {
                res[i] = v.(InformationProtectionLabelable)
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. 
func (m *InformationProtectionPolicy) GetLabels()([]InformationProtectionLabelable) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabels sets the labels property value. 
func (m *InformationProtectionPolicy) SetLabels(value []InformationProtectionLabelable)() {
    if m != nil {
        m.labels = value
    }
}
