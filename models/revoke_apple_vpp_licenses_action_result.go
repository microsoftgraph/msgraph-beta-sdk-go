package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RevokeAppleVppLicensesActionResult revoke Apple Vpp licenses action result
type RevokeAppleVppLicensesActionResult struct {
    DeviceActionResult
}
// NewRevokeAppleVppLicensesActionResult instantiates a new revokeAppleVppLicensesActionResult and sets the default values.
func NewRevokeAppleVppLicensesActionResult()(*RevokeAppleVppLicensesActionResult) {
    m := &RevokeAppleVppLicensesActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateRevokeAppleVppLicensesActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeAppleVppLicensesActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRevokeAppleVppLicensesActionResult(), nil
}
// GetFailedLicensesCount gets the failedLicensesCount property value. Total number of Apple Vpp licenses that failed to revoke
func (m *RevokeAppleVppLicensesActionResult) GetFailedLicensesCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedLicensesCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeAppleVppLicensesActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["failedLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedLicensesCount(val)
        }
        return nil
    }
    res["totalLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicensesCount(val)
        }
        return nil
    }
    return res
}
// GetTotalLicensesCount gets the totalLicensesCount property value. Total number of Apple Vpp licenses associated
func (m *RevokeAppleVppLicensesActionResult) GetTotalLicensesCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalLicensesCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RevokeAppleVppLicensesActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("failedLicensesCount", m.GetFailedLicensesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicensesCount", m.GetTotalLicensesCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFailedLicensesCount sets the failedLicensesCount property value. Total number of Apple Vpp licenses that failed to revoke
func (m *RevokeAppleVppLicensesActionResult) SetFailedLicensesCount(value *int32)() {
    err := m.GetBackingStore().Set("failedLicensesCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalLicensesCount sets the totalLicensesCount property value. Total number of Apple Vpp licenses associated
func (m *RevokeAppleVppLicensesActionResult) SetTotalLicensesCount(value *int32)() {
    err := m.GetBackingStore().Set("totalLicensesCount", value)
    if err != nil {
        panic(err)
    }
}
// RevokeAppleVppLicensesActionResultable 
type RevokeAppleVppLicensesActionResultable interface {
    DeviceActionResultable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedLicensesCount()(*int32)
    GetTotalLicensesCount()(*int32)
    SetFailedLicensesCount(value *int32)()
    SetTotalLicensesCount(value *int32)()
}
