package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharedEmailDomain struct {
    Entity
    // 
    provisioningStatus *string;
}
// Instantiates a new sharedEmailDomain and sets the default values.
func NewSharedEmailDomain()(*SharedEmailDomain) {
    m := &SharedEmailDomain{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the provisioningStatus property value. 
func (m *SharedEmailDomain) GetProvisioningStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatus
    }
}
// The deserialization information for the current model
func (m *SharedEmailDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["provisioningStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProvisioningStatus(val)
        return nil
    }
    return res
}
func (m *SharedEmailDomain) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SharedEmailDomain) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("provisioningStatus", m.GetProvisioningStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the provisioningStatus property value. 
// Parameters:
//  - value : Value to set for the provisioningStatus property.
func (m *SharedEmailDomain) SetProvisioningStatus(value *string)() {
    m.provisioningStatus = value
}
