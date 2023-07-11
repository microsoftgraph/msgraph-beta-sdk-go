package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationConflictSummary conflict summary for a set of device configuration policies.
type DeviceConfigurationConflictSummary struct {
    Entity
}
// NewDeviceConfigurationConflictSummary instantiates a new deviceConfigurationConflictSummary and sets the default values.
func NewDeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummary) {
    m := &DeviceConfigurationConflictSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationConflictSummary(), nil
}
// GetConflictingDeviceConfigurations gets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
func (m *DeviceConfigurationConflictSummary) GetConflictingDeviceConfigurations()([]SettingSourceable) {
    val, err := m.GetBackingStore().Get("conflictingDeviceConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SettingSourceable)
    }
    return nil
}
// GetContributingSettings gets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) GetContributingSettings()([]string) {
    val, err := m.GetBackingStore().Get("contributingSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDeviceCheckinsImpacted gets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) GetDeviceCheckinsImpacted()(*int32) {
    val, err := m.GetBackingStore().Get("deviceCheckinsImpacted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationConflictSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictingDeviceConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SettingSourceable)
                }
            }
            m.SetConflictingDeviceConfigurations(res)
        }
        return nil
    }
    res["contributingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetContributingSettings(res)
        }
        return nil
    }
    res["deviceCheckinsImpacted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCheckinsImpacted(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceConfigurationConflictSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceConfigurationConflictSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConflictingDeviceConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConflictingDeviceConfigurations()))
        for i, v := range m.GetConflictingDeviceConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConflictingDeviceConfigurations sets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
func (m *DeviceConfigurationConflictSummary) SetConflictingDeviceConfigurations(value []SettingSourceable)() {
    err := m.GetBackingStore().Set("conflictingDeviceConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetContributingSettings sets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) SetContributingSettings(value []string)() {
    err := m.GetBackingStore().Set("contributingSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCheckinsImpacted sets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) SetDeviceCheckinsImpacted(value *int32)() {
    err := m.GetBackingStore().Set("deviceCheckinsImpacted", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceConfigurationConflictSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceConfigurationConflictSummaryable 
type DeviceConfigurationConflictSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictingDeviceConfigurations()([]SettingSourceable)
    GetContributingSettings()([]string)
    GetDeviceCheckinsImpacted()(*int32)
    GetOdataType()(*string)
    SetConflictingDeviceConfigurations(value []SettingSourceable)()
    SetContributingSettings(value []string)()
    SetDeviceCheckinsImpacted(value *int32)()
    SetOdataType(value *string)()
}
