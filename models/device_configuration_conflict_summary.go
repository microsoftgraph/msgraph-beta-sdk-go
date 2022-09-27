package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationConflictSummary conflict summary for a set of device configuration policies.
type DeviceConfigurationConflictSummary struct {
    Entity
    // The set of policies in conflict with the given setting
    conflictingDeviceConfigurations []SettingSourceable
    // The set of settings in conflict with the given policies
    contributingSettings []string
    // The count of checkins impacted by the conflicting policies and settings
    deviceCheckinsImpacted *int32
}
// NewDeviceConfigurationConflictSummary instantiates a new deviceConfigurationConflictSummary and sets the default values.
func NewDeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummary) {
    m := &DeviceConfigurationConflictSummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceConfigurationConflictSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationConflictSummary(), nil
}
// GetConflictingDeviceConfigurations gets the conflictingDeviceConfigurations property value. The set of policies in conflict with the given setting
func (m *DeviceConfigurationConflictSummary) GetConflictingDeviceConfigurations()([]SettingSourceable) {
    return m.conflictingDeviceConfigurations
}
// GetContributingSettings gets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) GetContributingSettings()([]string) {
    return m.contributingSettings
}
// GetDeviceCheckinsImpacted gets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) GetDeviceCheckinsImpacted()(*int32) {
    return m.deviceCheckinsImpacted
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationConflictSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictingDeviceConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue , m.SetConflictingDeviceConfigurations)
    res["contributingSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetContributingSettings)
    res["deviceCheckinsImpacted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeviceCheckinsImpacted)
    return res
}
// Serialize serializes information the current object
func (m *DeviceConfigurationConflictSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConflictingDeviceConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConflictingDeviceConfigurations())
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
    m.conflictingDeviceConfigurations = value
}
// SetContributingSettings sets the contributingSettings property value. The set of settings in conflict with the given policies
func (m *DeviceConfigurationConflictSummary) SetContributingSettings(value []string)() {
    m.contributingSettings = value
}
// SetDeviceCheckinsImpacted sets the deviceCheckinsImpacted property value. The count of checkins impacted by the conflicting policies and settings
func (m *DeviceConfigurationConflictSummary) SetDeviceCheckinsImpacted(value *int32)() {
    m.deviceCheckinsImpacted = value
}
