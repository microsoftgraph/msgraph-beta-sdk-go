package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserInsightsSettings 
type UserInsightsSettings struct {
    Entity
    // true if user's itemInsights and meeting hours insights are enabled; false if user's itemInsights and meeting hours insights are disabled. Default is true. Optional.
    isEnabled *bool
}
// NewUserInsightsSettings instantiates a new userInsightsSettings and sets the default values.
func NewUserInsightsSettings()(*UserInsightsSettings) {
    m := &UserInsightsSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserInsightsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserInsightsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserInsightsSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserInsightsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    return res
}
// GetIsEnabled gets the isEnabled property value. true if user's itemInsights and meeting hours insights are enabled; false if user's itemInsights and meeting hours insights are disabled. Default is true. Optional.
func (m *UserInsightsSettings) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// Serialize serializes information the current object
func (m *UserInsightsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEnabled sets the isEnabled property value. true if user's itemInsights and meeting hours insights are enabled; false if user's itemInsights and meeting hours insights are disabled. Default is true. Optional.
func (m *UserInsightsSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
