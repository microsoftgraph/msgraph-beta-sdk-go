package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppProvisioningConfigGroupAssignment struct {
    Entity
    // The ID of the AAD group in which the app provisioning configuration is being targeted.
    targetGroupId *string;
}
// Instantiates a new mobileAppProvisioningConfigGroupAssignment and sets the default values.
func NewMobileAppProvisioningConfigGroupAssignment()(*MobileAppProvisioningConfigGroupAssignment) {
    m := &MobileAppProvisioningConfigGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the targetGroupId property value. The ID of the AAD group in which the app provisioning configuration is being targeted.
func (m *MobileAppProvisioningConfigGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// The deserialization information for the current model
func (m *MobileAppProvisioningConfigGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetGroupId(val)
        return nil
    }
    return res
}
func (m *MobileAppProvisioningConfigGroupAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppProvisioningConfigGroupAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the targetGroupId property value. The ID of the AAD group in which the app provisioning configuration is being targeted.
// Parameters:
//  - value : Value to set for the targetGroupId property.
func (m *MobileAppProvisioningConfigGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
