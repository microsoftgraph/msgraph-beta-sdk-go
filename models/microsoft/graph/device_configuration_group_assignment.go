package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceConfigurationGroupAssignment struct {
    Entity
    deviceConfiguration *DeviceConfiguration;
    excludeGroup *bool;
    targetGroupId *string;
}
func NewDeviceConfigurationGroupAssignment()(*DeviceConfigurationGroupAssignment) {
    m := &DeviceConfigurationGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceConfigurationGroupAssignment) GetDeviceConfiguration()(*DeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
    }
}
func (m *DeviceConfigurationGroupAssignment) GetExcludeGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroup
    }
}
func (m *DeviceConfigurationGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
func (m *DeviceConfigurationGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfiguration() })
        if err != nil {
            return err
        }
        m.SetDeviceConfiguration(val.(*DeviceConfiguration))
        return nil
    }
    res["excludeGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExcludeGroup(val)
        return nil
    }
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
func (m *DeviceConfigurationGroupAssignment) IsNil()(bool) {
    return m == nil
}
func (m *DeviceConfigurationGroupAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceConfiguration", m.GetDeviceConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("excludeGroup", m.GetExcludeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceConfigurationGroupAssignment) SetDeviceConfiguration(value *DeviceConfiguration)() {
    m.deviceConfiguration = value
}
func (m *DeviceConfigurationGroupAssignment) SetExcludeGroup(value *bool)() {
    m.excludeGroup = value
}
func (m *DeviceConfigurationGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
