package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkCustomConfiguration 
type AndroidForWorkCustomConfiguration struct {
    DeviceConfiguration
    // OMA settings. This collection can contain a maximum of 500 elements.
    omaSettings []OmaSettingable
}
// NewAndroidForWorkCustomConfiguration instantiates a new AndroidForWorkCustomConfiguration and sets the default values.
func NewAndroidForWorkCustomConfiguration()(*AndroidForWorkCustomConfiguration) {
    m := &AndroidForWorkCustomConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkCustomConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkCustomConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkCustomConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkCustomConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkCustomConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["omaSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOmaSettingFromDiscriminatorValue , m.SetOmaSettings)
    return res
}
// GetOmaSettings gets the omaSettings property value. OMA settings. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkCustomConfiguration) GetOmaSettings()([]OmaSettingable) {
    return m.omaSettings
}
// Serialize serializes information the current object
func (m *AndroidForWorkCustomConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOmaSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOmaSettings())
        err = writer.WriteCollectionOfObjectValues("omaSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOmaSettings sets the omaSettings property value. OMA settings. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkCustomConfiguration) SetOmaSettings(value []OmaSettingable)() {
    m.omaSettings = value
}
