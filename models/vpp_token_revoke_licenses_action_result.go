package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppTokenRevokeLicensesActionResult the status of the revoke licenses action performed on the Apple Volume Purchase Program token.
type VppTokenRevokeLicensesActionResult struct {
    VppTokenActionResult
}
// NewVppTokenRevokeLicensesActionResult instantiates a new VppTokenRevokeLicensesActionResult and sets the default values.
func NewVppTokenRevokeLicensesActionResult()(*VppTokenRevokeLicensesActionResult) {
    m := &VppTokenRevokeLicensesActionResult{
        VppTokenActionResult: *NewVppTokenActionResult(),
    }
    return m
}
// CreateVppTokenRevokeLicensesActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVppTokenRevokeLicensesActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppTokenRevokeLicensesActionResult(), nil
}
// GetActionFailureReason gets the actionFailureReason property value. Possible types of reasons for an Apple Volume Purchase Program token action failure.
// returns a *VppTokenActionFailureReason when successful
func (m *VppTokenRevokeLicensesActionResult) GetActionFailureReason()(*VppTokenActionFailureReason) {
    val, err := m.GetBackingStore().Get("actionFailureReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VppTokenActionFailureReason)
    }
    return nil
}
// GetFailedLicensesCount gets the failedLicensesCount property value. A count of the number of licenses that failed to revoke.
// returns a *int32 when successful
func (m *VppTokenRevokeLicensesActionResult) GetFailedLicensesCount()(*int32) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VppTokenRevokeLicensesActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VppTokenActionResult.GetFieldDeserializers()
    res["actionFailureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenActionFailureReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionFailureReason(val.(*VppTokenActionFailureReason))
        }
        return nil
    }
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
// GetTotalLicensesCount gets the totalLicensesCount property value. A count of the number of licenses that were attempted to revoke.
// returns a *int32 when successful
func (m *VppTokenRevokeLicensesActionResult) GetTotalLicensesCount()(*int32) {
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
func (m *VppTokenRevokeLicensesActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VppTokenActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionFailureReason() != nil {
        cast := (*m.GetActionFailureReason()).String()
        err = writer.WriteStringValue("actionFailureReason", &cast)
        if err != nil {
            return err
        }
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
// SetActionFailureReason sets the actionFailureReason property value. Possible types of reasons for an Apple Volume Purchase Program token action failure.
func (m *VppTokenRevokeLicensesActionResult) SetActionFailureReason(value *VppTokenActionFailureReason)() {
    err := m.GetBackingStore().Set("actionFailureReason", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedLicensesCount sets the failedLicensesCount property value. A count of the number of licenses that failed to revoke.
func (m *VppTokenRevokeLicensesActionResult) SetFailedLicensesCount(value *int32)() {
    err := m.GetBackingStore().Set("failedLicensesCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalLicensesCount sets the totalLicensesCount property value. A count of the number of licenses that were attempted to revoke.
func (m *VppTokenRevokeLicensesActionResult) SetTotalLicensesCount(value *int32)() {
    err := m.GetBackingStore().Set("totalLicensesCount", value)
    if err != nil {
        panic(err)
    }
}
type VppTokenRevokeLicensesActionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VppTokenActionResultable
    GetActionFailureReason()(*VppTokenActionFailureReason)
    GetFailedLicensesCount()(*int32)
    GetTotalLicensesCount()(*int32)
    SetActionFailureReason(value *VppTokenActionFailureReason)()
    SetFailedLicensesCount(value *int32)()
    SetTotalLicensesCount(value *int32)()
}
