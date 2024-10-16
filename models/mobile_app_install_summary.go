package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallSummary contains properties for the installation summary of a mobile app. This will be deprecated starting May, 2023 (Intune Release 2305).
type MobileAppInstallSummary struct {
    Entity
}
// NewMobileAppInstallSummary instantiates a new MobileAppInstallSummary and sets the default values.
func NewMobileAppInstallSummary()(*MobileAppInstallSummary) {
    m := &MobileAppInstallSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppInstallSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppInstallSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppInstallSummary(), nil
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Number of Devices that have failed to install this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetFailedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFailedUserCount gets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetFailedUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MobileAppInstallSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["failedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUserCount(val)
        }
        return nil
    }
    res["installedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledDeviceCount(val)
        }
        return nil
    }
    res["installedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledUserCount(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["notApplicableUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
        }
        return nil
    }
    res["notInstalledDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledDeviceCount(val)
        }
        return nil
    }
    res["notInstalledUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledUserCount(val)
        }
        return nil
    }
    res["pendingInstallDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInstallDeviceCount(val)
        }
        return nil
    }
    res["pendingInstallUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInstallUserCount(val)
        }
        return nil
    }
    return res
}
// GetInstalledDeviceCount gets the installedDeviceCount property value. Number of Devices that have successfully installed this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetInstalledDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("installedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetInstalledUserCount gets the installedUserCount property value. Number of Users whose devices have all succeeded to install this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetInstalledUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("installedUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of Devices that are not applicable for this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetNotApplicableDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of Users whose devices were all not applicable for this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetNotApplicableUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotInstalledDeviceCount gets the notInstalledDeviceCount property value. Number of Devices that does not have this app installed.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetNotInstalledDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("notInstalledDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotInstalledUserCount gets the notInstalledUserCount property value. Number of Users that have 1 or more devices that did not install this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetNotInstalledUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("notInstalledUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPendingInstallDeviceCount gets the pendingInstallDeviceCount property value. Number of Devices that have been notified to install this app.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetPendingInstallDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("pendingInstallDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPendingInstallUserCount gets the pendingInstallUserCount property value. Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
// returns a *int32 when successful
func (m *MobileAppInstallSummary) GetPendingInstallUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("pendingInstallUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppInstallSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Number of Devices that have failed to install this app.
func (m *MobileAppInstallSummary) SetFailedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("failedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedUserCount sets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this app.
func (m *MobileAppInstallSummary) SetFailedUserCount(value *int32)() {
    err := m.GetBackingStore().Set("failedUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInstalledDeviceCount sets the installedDeviceCount property value. Number of Devices that have successfully installed this app.
func (m *MobileAppInstallSummary) SetInstalledDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("installedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInstalledUserCount sets the installedUserCount property value. Number of Users whose devices have all succeeded to install this app.
func (m *MobileAppInstallSummary) SetInstalledUserCount(value *int32)() {
    err := m.GetBackingStore().Set("installedUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of Devices that are not applicable for this app.
func (m *MobileAppInstallSummary) SetNotApplicableDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of Users whose devices were all not applicable for this app.
func (m *MobileAppInstallSummary) SetNotApplicableUserCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotInstalledDeviceCount sets the notInstalledDeviceCount property value. Number of Devices that does not have this app installed.
func (m *MobileAppInstallSummary) SetNotInstalledDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("notInstalledDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotInstalledUserCount sets the notInstalledUserCount property value. Number of Users that have 1 or more devices that did not install this app.
func (m *MobileAppInstallSummary) SetNotInstalledUserCount(value *int32)() {
    err := m.GetBackingStore().Set("notInstalledUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingInstallDeviceCount sets the pendingInstallDeviceCount property value. Number of Devices that have been notified to install this app.
func (m *MobileAppInstallSummary) SetPendingInstallDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("pendingInstallDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingInstallUserCount sets the pendingInstallUserCount property value. Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
func (m *MobileAppInstallSummary) SetPendingInstallUserCount(value *int32)() {
    err := m.GetBackingStore().Set("pendingInstallUserCount", value)
    if err != nil {
        panic(err)
    }
}
type MobileAppInstallSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedDeviceCount()(*int32)
    GetFailedUserCount()(*int32)
    GetInstalledDeviceCount()(*int32)
    GetInstalledUserCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetNotInstalledDeviceCount()(*int32)
    GetNotInstalledUserCount()(*int32)
    GetPendingInstallDeviceCount()(*int32)
    GetPendingInstallUserCount()(*int32)
    SetFailedDeviceCount(value *int32)()
    SetFailedUserCount(value *int32)()
    SetInstalledDeviceCount(value *int32)()
    SetInstalledUserCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetNotInstalledDeviceCount(value *int32)()
    SetNotInstalledUserCount(value *int32)()
    SetPendingInstallDeviceCount(value *int32)()
    SetPendingInstallUserCount(value *int32)()
}
