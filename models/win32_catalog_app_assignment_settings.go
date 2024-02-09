package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32CatalogAppAssignmentSettings contains properties used to assign an Win32 catalog mobile app to a group.
type Win32CatalogAppAssignmentSettings struct {
    Win32LobAppAssignmentSettings
}
// NewWin32CatalogAppAssignmentSettings instantiates a new Win32CatalogAppAssignmentSettings and sets the default values.
func NewWin32CatalogAppAssignmentSettings()(*Win32CatalogAppAssignmentSettings) {
    m := &Win32CatalogAppAssignmentSettings{
        Win32LobAppAssignmentSettings: *NewWin32LobAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.win32CatalogAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32CatalogAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWin32CatalogAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32CatalogAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Win32CatalogAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppAssignmentSettings.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Win32CatalogAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type Win32CatalogAppAssignmentSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppAssignmentSettingsable
}
