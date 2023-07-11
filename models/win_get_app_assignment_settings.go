package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WinGetAppAssignmentSettings abstract class to contain properties used to assign a mobile app to a group.
type WinGetAppAssignmentSettings struct {
    MobileAppAssignmentSettings
}
// NewWinGetAppAssignmentSettings instantiates a new winGetAppAssignmentSettings and sets the default values.
func NewWinGetAppAssignmentSettings()(*WinGetAppAssignmentSettings) {
    m := &WinGetAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.winGetAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWinGetAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWinGetAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWinGetAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WinGetAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["installTimeSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWinGetAppInstallTimeSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallTimeSettings(val.(WinGetAppInstallTimeSettingsable))
        }
        return nil
    }
    res["notifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWinGetAppNotification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifications(val.(*WinGetAppNotification))
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
    res["restartSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWinGetAppRestartSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartSettings(val.(WinGetAppRestartSettingsable))
        }
        return nil
    }
    return res
}
// GetInstallTimeSettings gets the installTimeSettings property value. The install time settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) GetInstallTimeSettings()(WinGetAppInstallTimeSettingsable) {
    val, err := m.GetBackingStore().Get("installTimeSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WinGetAppInstallTimeSettingsable)
    }
    return nil
}
// GetNotifications gets the notifications property value. Contains value for notification status.
func (m *WinGetAppAssignmentSettings) GetNotifications()(*WinGetAppNotification) {
    val, err := m.GetBackingStore().Get("notifications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WinGetAppNotification)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WinGetAppAssignmentSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRestartSettings gets the restartSettings property value. The reboot settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) GetRestartSettings()(WinGetAppRestartSettingsable) {
    val, err := m.GetBackingStore().Get("restartSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WinGetAppRestartSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WinGetAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("installTimeSettings", m.GetInstallTimeSettings())
        if err != nil {
            return err
        }
    }
    if m.GetNotifications() != nil {
        cast := (*m.GetNotifications()).String()
        err = writer.WriteStringValue("notifications", &cast)
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
    {
        err = writer.WriteObjectValue("restartSettings", m.GetRestartSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstallTimeSettings sets the installTimeSettings property value. The install time settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) SetInstallTimeSettings(value WinGetAppInstallTimeSettingsable)() {
    err := m.GetBackingStore().Set("installTimeSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetNotifications sets the notifications property value. Contains value for notification status.
func (m *WinGetAppAssignmentSettings) SetNotifications(value *WinGetAppNotification)() {
    err := m.GetBackingStore().Set("notifications", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WinGetAppAssignmentSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartSettings sets the restartSettings property value. The reboot settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) SetRestartSettings(value WinGetAppRestartSettingsable)() {
    err := m.GetBackingStore().Set("restartSettings", value)
    if err != nil {
        panic(err)
    }
}
// WinGetAppAssignmentSettingsable 
type WinGetAppAssignmentSettingsable interface {
    MobileAppAssignmentSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallTimeSettings()(WinGetAppInstallTimeSettingsable)
    GetNotifications()(*WinGetAppNotification)
    GetOdataType()(*string)
    GetRestartSettings()(WinGetAppRestartSettingsable)
    SetInstallTimeSettings(value WinGetAppInstallTimeSettingsable)()
    SetNotifications(value *WinGetAppNotification)()
    SetOdataType(value *string)()
    SetRestartSettings(value WinGetAppRestartSettingsable)()
}
