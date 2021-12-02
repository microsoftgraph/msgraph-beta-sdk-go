package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedEmailDomain 
type SharedEmailDomain struct {
    Entity
    // 
    provisioningStatus *string;
}
// NewSharedEmailDomain instantiates a new sharedEmailDomain and sets the default values.
func NewSharedEmailDomain()(*SharedEmailDomain) {
    m := &SharedEmailDomain{
        Entity: *NewEntity(),
    }
    return m
}
// GetProvisioningStatus gets the provisioningStatus property value. 
func (m *SharedEmailDomain) GetProvisioningStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["provisioningStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatus(val)
        }
        return nil
    }
    return res
}
func (m *SharedEmailDomain) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetProvisioningStatus sets the provisioningStatus property value. 
func (m *SharedEmailDomain) SetProvisioningStatus(value *string)() {
    if m != nil {
        m.provisioningStatus = value
    }
}
