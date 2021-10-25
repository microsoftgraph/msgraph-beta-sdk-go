package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementExchangeOnPremisesPolicy struct {
    Entity
    accessRules []DeviceManagementExchangeAccessRule;
    conditionalAccessSettings *OnPremisesConditionalAccessSettings;
    defaultAccessLevel *DeviceManagementExchangeAccessLevel;
    knownDeviceClasses []DeviceManagementExchangeDeviceClass;
    notificationContent []byte;
}
func NewDeviceManagementExchangeOnPremisesPolicy()(*DeviceManagementExchangeOnPremisesPolicy) {
    m := &DeviceManagementExchangeOnPremisesPolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementExchangeOnPremisesPolicy) GetAccessRules()([]DeviceManagementExchangeAccessRule) {
    if m == nil {
        return nil
    } else {
        return m.accessRules
    }
}
func (m *DeviceManagementExchangeOnPremisesPolicy) GetConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessSettings
    }
}
func (m *DeviceManagementExchangeOnPremisesPolicy) GetDefaultAccessLevel()(*DeviceManagementExchangeAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.defaultAccessLevel
    }
}
func (m *DeviceManagementExchangeOnPremisesPolicy) GetKnownDeviceClasses()([]DeviceManagementExchangeDeviceClass) {
    if m == nil {
        return nil
    } else {
        return m.knownDeviceClasses
    }
}
func (m *DeviceManagementExchangeOnPremisesPolicy) GetNotificationContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.notificationContent
    }
}
func (m *DeviceManagementExchangeOnPremisesPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeAccessRule() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementExchangeAccessRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementExchangeAccessRule))
        }
        m.SetAccessRules(res)
        return nil
    }
    res["conditionalAccessSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesConditionalAccessSettings() })
        if err != nil {
            return err
        }
        m.SetConditionalAccessSettings(val.(*OnPremisesConditionalAccessSettings))
        return nil
    }
    res["defaultAccessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessLevel)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeAccessLevel)
        m.SetDefaultAccessLevel(&cast)
        return nil
    }
    res["knownDeviceClasses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeDeviceClass() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementExchangeDeviceClass, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementExchangeDeviceClass))
        }
        m.SetKnownDeviceClasses(res)
        return nil
    }
    res["notificationContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetNotificationContent(val)
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeOnPremisesPolicy) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementExchangeOnPremisesPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessRules()))
        for i, v := range m.GetAccessRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccessSettings", m.GetConditionalAccessSettings())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultAccessLevel() != nil {
        cast := m.GetDefaultAccessLevel().String()
        err = writer.WriteStringValue("defaultAccessLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKnownDeviceClasses()))
        for i, v := range m.GetKnownDeviceClasses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("knownDeviceClasses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("notificationContent", m.GetNotificationContent())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementExchangeOnPremisesPolicy) SetAccessRules(value []DeviceManagementExchangeAccessRule)() {
    m.accessRules = value
}
func (m *DeviceManagementExchangeOnPremisesPolicy) SetConditionalAccessSettings(value *OnPremisesConditionalAccessSettings)() {
    m.conditionalAccessSettings = value
}
func (m *DeviceManagementExchangeOnPremisesPolicy) SetDefaultAccessLevel(value *DeviceManagementExchangeAccessLevel)() {
    m.defaultAccessLevel = value
}
func (m *DeviceManagementExchangeOnPremisesPolicy) SetKnownDeviceClasses(value []DeviceManagementExchangeDeviceClass)() {
    m.knownDeviceClasses = value
}
func (m *DeviceManagementExchangeOnPremisesPolicy) SetNotificationContent(value []byte)() {
    m.notificationContent = value
}
