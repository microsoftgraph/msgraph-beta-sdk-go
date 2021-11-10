package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceConfigurationConflictSummary struct {
    Entity
    // The set of policies in conflict with the given setting
    conflictingDeviceConfigurations []SettingSource;
    // The set of settings in conflict with the given policies
    contributingSettings []string;
    // The count of checkins impacted by the conflicting policies and settings
    deviceCheckinsImpacted *int32;
}
// Instantiates a new deviceConfigurationConflictSummary and sets the default values.
func NewDeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummary) {
    m := &DeviceConfigurationConflictSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
func (m *DeviceConfigurationConflictSummary) GetConflictingDeviceConfigurations()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.conflictingDeviceConfigurations
    }
}
// Gets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) GetContributingSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contributingSettings
    }
}
// Gets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) GetDeviceCheckinsImpacted()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCheckinsImpacted
    }
}
// The deserialization information for the current model
func (m *DeviceConfigurationConflictSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictingDeviceConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingSource))
            }
            m.SetConflictingDeviceConfigurations(res)
        }
        return nil
    }
    res["contributingSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContributingSettings(res)
        }
        return nil
    }
    res["deviceCheckinsImpacted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCheckinsImpacted(val)
        }
        return nil
    }
    return res
}
func (m *DeviceConfigurationConflictSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceConfigurationConflictSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConflictingDeviceConfigurations()))
        for i, v := range m.GetConflictingDeviceConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("conflictingDeviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("contributingSettings", m.GetContributingSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCheckinsImpacted", m.GetDeviceCheckinsImpacted())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
// Parameters:
//  - value : Value to set for the conflictingDeviceConfigurations property.
func (m *DeviceConfigurationConflictSummary) SetConflictingDeviceConfigurations(value []SettingSource)() {
    m.conflictingDeviceConfigurations = value
}
// Sets the contributingSettings property value. The set of settings in conflict with the given policies
// Parameters:
//  - value : Value to set for the contributingSettings property.
func (m *DeviceConfigurationConflictSummary) SetContributingSettings(value []string)() {
    m.contributingSettings = value
}
// Sets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
// Parameters:
//  - value : Value to set for the deviceCheckinsImpacted property.
func (m *DeviceConfigurationConflictSummary) SetDeviceCheckinsImpacted(value *int32)() {
    m.deviceCheckinsImpacted = value
}
