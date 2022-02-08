package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExchangeOnPremisesPolicy 
type DeviceManagementExchangeOnPremisesPolicy struct {
    Entity
    // The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
    accessRules []DeviceManagementExchangeAccessRule;
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings *OnPremisesConditionalAccessSettings;
    // Default access state in Exchange. This rule applies globally to the entire Exchange organization. Possible values are: none, allow, block, quarantine.
    defaultAccessLevel *DeviceManagementExchangeAccessLevel;
    // The list of device classes known to Exchange
    knownDeviceClasses []DeviceManagementExchangeDeviceClass;
    // Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
    notificationContent []byte;
}
// NewDeviceManagementExchangeOnPremisesPolicy instantiates a new deviceManagementExchangeOnPremisesPolicy and sets the default values.
func NewDeviceManagementExchangeOnPremisesPolicy()(*DeviceManagementExchangeOnPremisesPolicy) {
    m := &DeviceManagementExchangeOnPremisesPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessRules gets the accessRules property value. The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
func (m *DeviceManagementExchangeOnPremisesPolicy) GetAccessRules()([]DeviceManagementExchangeAccessRule) {
    if m == nil {
        return nil
    } else {
        return m.accessRules
    }
}
// GetConditionalAccessSettings gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagementExchangeOnPremisesPolicy) GetConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessSettings
    }
}
// GetDefaultAccessLevel gets the defaultAccessLevel property value. Default access state in Exchange. This rule applies globally to the entire Exchange organization. Possible values are: none, allow, block, quarantine.
func (m *DeviceManagementExchangeOnPremisesPolicy) GetDefaultAccessLevel()(*DeviceManagementExchangeAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.defaultAccessLevel
    }
}
// GetKnownDeviceClasses gets the knownDeviceClasses property value. The list of device classes known to Exchange
func (m *DeviceManagementExchangeOnPremisesPolicy) GetKnownDeviceClasses()([]DeviceManagementExchangeDeviceClass) {
    if m == nil {
        return nil
    } else {
        return m.knownDeviceClasses
    }
}
// GetNotificationContent gets the notificationContent property value. Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
func (m *DeviceManagementExchangeOnPremisesPolicy) GetNotificationContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.notificationContent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeOnPremisesPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeAccessRule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeAccessRule, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementExchangeAccessRule))
            }
            m.SetAccessRules(res)
        }
        return nil
    }
    res["conditionalAccessSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesConditionalAccessSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessSettings(val.(*OnPremisesConditionalAccessSettings))
        }
        return nil
    }
    res["defaultAccessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultAccessLevel(val.(*DeviceManagementExchangeAccessLevel))
        }
        return nil
    }
    res["knownDeviceClasses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeDeviceClass() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeDeviceClass, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementExchangeDeviceClass))
            }
            m.SetKnownDeviceClasses(res)
        }
        return nil
    }
    res["notificationContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationContent(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeOnPremisesPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeOnPremisesPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessRules() != nil {
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
        cast := (*m.GetDefaultAccessLevel()).String()
        err = writer.WriteStringValue("defaultAccessLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetKnownDeviceClasses() != nil {
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
// SetAccessRules sets the accessRules property value. The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
func (m *DeviceManagementExchangeOnPremisesPolicy) SetAccessRules(value []DeviceManagementExchangeAccessRule)() {
    if m != nil {
        m.accessRules = value
    }
}
// SetConditionalAccessSettings sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagementExchangeOnPremisesPolicy) SetConditionalAccessSettings(value *OnPremisesConditionalAccessSettings)() {
    if m != nil {
        m.conditionalAccessSettings = value
    }
}
// SetDefaultAccessLevel sets the defaultAccessLevel property value. Default access state in Exchange. This rule applies globally to the entire Exchange organization. Possible values are: none, allow, block, quarantine.
func (m *DeviceManagementExchangeOnPremisesPolicy) SetDefaultAccessLevel(value *DeviceManagementExchangeAccessLevel)() {
    if m != nil {
        m.defaultAccessLevel = value
    }
}
// SetKnownDeviceClasses sets the knownDeviceClasses property value. The list of device classes known to Exchange
func (m *DeviceManagementExchangeOnPremisesPolicy) SetKnownDeviceClasses(value []DeviceManagementExchangeDeviceClass)() {
    if m != nil {
        m.knownDeviceClasses = value
    }
}
// SetNotificationContent sets the notificationContent property value. Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
func (m *DeviceManagementExchangeOnPremisesPolicy) SetNotificationContent(value []byte)() {
    if m != nil {
        m.notificationContent = value
    }
}
