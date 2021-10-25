package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementScriptGroupAssignment struct {
    Entity
    targetGroupId *string;
}
func NewDeviceManagementScriptGroupAssignment()(*DeviceManagementScriptGroupAssignment) {
    m := &DeviceManagementScriptGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementScriptGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
func (m *DeviceManagementScriptGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *DeviceManagementScriptGroupAssignment) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceManagementScriptGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
