package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementScriptGroupAssignment struct {
    Entity
    // The Id of the Azure Active Directory group we are targeting the script to.
    targetGroupId *string;
}
// Instantiates a new deviceManagementScriptGroupAssignment and sets the default values.
func NewDeviceManagementScriptGroupAssignment()(*DeviceManagementScriptGroupAssignment) {
    m := &DeviceManagementScriptGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the targetGroupId property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementScriptGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementScriptGroupAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementScriptGroupAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the targetGroupId property value. The Id of the Azure Active Directory group we are targeting the script to.
// Parameters:
//  - value : Value to set for the targetGroupId property.
func (m *DeviceManagementScriptGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
