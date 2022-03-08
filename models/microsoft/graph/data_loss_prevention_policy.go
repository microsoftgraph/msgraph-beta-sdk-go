package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DataLossPreventionPolicy provides operations to manage the compliance singleton.
type DataLossPreventionPolicy struct {
    Entity
    // 
    name *string;
}
// NewDataLossPreventionPolicy instantiates a new dataLossPreventionPolicy and sets the default values.
func NewDataLossPreventionPolicy()(*DataLossPreventionPolicy) {
    m := &DataLossPreventionPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataLossPreventionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataLossPreventionPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDataLossPreventionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataLossPreventionPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. 
func (m *DataLossPreventionPolicy) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *DataLossPreventionPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DataLossPreventionPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. 
func (m *DataLossPreventionPolicy) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
