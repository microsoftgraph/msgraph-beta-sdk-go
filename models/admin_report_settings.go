package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminReportSettings 
type AdminReportSettings struct {
    Entity
    // If set to true, all reports will conceal user information such as usernames, groups, and sites. If false, all reports will show identifiable information. This property represents a setting in the Microsoft 365 admin center. Required.
    displayConcealedNames *bool
}
// NewAdminReportSettings instantiates a new AdminReportSettings and sets the default values.
func NewAdminReportSettings()(*AdminReportSettings) {
    m := &AdminReportSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminReportSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminReportSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminReportSettings(), nil
}
// GetDisplayConcealedNames gets the displayConcealedNames property value. If set to true, all reports will conceal user information such as usernames, groups, and sites. If false, all reports will show identifiable information. This property represents a setting in the Microsoft 365 admin center. Required.
func (m *AdminReportSettings) GetDisplayConcealedNames()(*bool) {
    return m.displayConcealedNames
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminReportSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayConcealedNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisplayConcealedNames)
    return res
}
// Serialize serializes information the current object
func (m *AdminReportSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("displayConcealedNames", m.GetDisplayConcealedNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayConcealedNames sets the displayConcealedNames property value. If set to true, all reports will conceal user information such as usernames, groups, and sites. If false, all reports will show identifiable information. This property represents a setting in the Microsoft 365 admin center. Required.
func (m *AdminReportSettings) SetDisplayConcealedNames(value *bool)() {
    m.displayConcealedNames = value
}
