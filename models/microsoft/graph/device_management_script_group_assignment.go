package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptGroupAssignment 
type DeviceManagementScriptGroupAssignment struct {
    Entity
    // The Id of the Azure Active Directory group we are targeting the script to.
    targetGroupId *string;
}
// NewDeviceManagementScriptGroupAssignment instantiates a new deviceManagementScriptGroupAssignment and sets the default values.
func NewDeviceManagementScriptGroupAssignment()(*DeviceManagementScriptGroupAssignment) {
    m := &DeviceManagementScriptGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetTargetGroupId gets the targetGroupId property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetTargetGroupId sets the targetGroupId property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptGroupAssignment) SetTargetGroupId(value *string)() {
    if m != nil {
        m.targetGroupId = value
    }
}
