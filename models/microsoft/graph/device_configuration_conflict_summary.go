package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceConfigurationConflictSummary struct {
    Entity
    conflictingDeviceConfigurations []SettingSource;
    contributingSettings []string;
    deviceCheckinsImpacted *int32;
}
func NewDeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummary) {
    m := &DeviceConfigurationConflictSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceConfigurationConflictSummary) GetConflictingDeviceConfigurations()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.conflictingDeviceConfigurations
    }
}
func (m *DeviceConfigurationConflictSummary) GetContributingSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contributingSettings
    }
}
func (m *DeviceConfigurationConflictSummary) GetDeviceCheckinsImpacted()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCheckinsImpacted
    }
}
func (m *DeviceConfigurationConflictSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictingDeviceConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingSource() })
        if err != nil {
            return err
        }
        res := make([]SettingSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingSource))
        }
        m.SetConflictingDeviceConfigurations(res)
        return nil
    }
    res["contributingSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetContributingSettings(res)
        return nil
    }
    res["deviceCheckinsImpacted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceCheckinsImpacted(val)
        return nil
    }
    return res
}
func (m *DeviceConfigurationConflictSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceConfigurationConflictSummary) SetConflictingDeviceConfigurations(value []SettingSource)() {
    m.conflictingDeviceConfigurations = value
}
func (m *DeviceConfigurationConflictSummary) SetContributingSettings(value []string)() {
    m.contributingSettings = value
}
func (m *DeviceConfigurationConflictSummary) SetDeviceCheckinsImpacted(value *int32)() {
    m.deviceCheckinsImpacted = value
}
