package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationConflictSummary 
type DeviceConfigurationConflictSummary struct {
    Entity
    // The set of policies in conflict with the given setting
    conflictingDeviceConfigurations []SettingSourceable;
    // The set of settings in conflict with the given policies
    contributingSettings []string;
    // The count of checkins impacted by the conflicting policies and settings
    deviceCheckinsImpacted *int32;
}
// NewDeviceConfigurationConflictSummary instantiates a new deviceConfigurationConflictSummary and sets the default values.
func NewDeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummary) {
    m := &DeviceConfigurationConflictSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceConfigurationConflictSummary(), nil
}
// GetConflictingDeviceConfigurations gets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
func (m *DeviceConfigurationConflictSummary) GetConflictingDeviceConfigurations()([]SettingSourceable) {
    if m == nil {
        return nil
    } else {
        return m.conflictingDeviceConfigurations
    }
}
// GetContributingSettings gets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) GetContributingSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contributingSettings
    }
}
// GetDeviceCheckinsImpacted gets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) GetDeviceCheckinsImpacted()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCheckinsImpacted
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationConflictSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictingDeviceConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSourceable, len(val))
            for i, v := range val {
                res[i] = v.(SettingSourceable)
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
// Serialize serializes information the current object
func (m *DeviceConfigurationConflictSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConflictingDeviceConfigurations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConflictingDeviceConfigurations()))
        for i, v := range m.GetConflictingDeviceConfigurations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("conflictingDeviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContributingSettings() != nil {
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
// SetConflictingDeviceConfigurations sets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
func (m *DeviceConfigurationConflictSummary) SetConflictingDeviceConfigurations(value []SettingSourceable)() {
    if m != nil {
        m.conflictingDeviceConfigurations = value
    }
}
// SetContributingSettings sets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) SetContributingSettings(value []string)() {
    if m != nil {
        m.contributingSettings = value
    }
}
// SetDeviceCheckinsImpacted sets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) SetDeviceCheckinsImpacted(value *int32)() {
    if m != nil {
        m.deviceCheckinsImpacted = value
    }
}
