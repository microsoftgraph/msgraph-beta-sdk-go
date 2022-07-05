package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppTokenRevokeLicensesActionResult 
type VppTokenRevokeLicensesActionResult struct {
    VppTokenActionResult
    // The reason for the revoke licenses action failure. Possible values are: none, appleFailure, internalError, expiredVppToken, expiredApplePushNotificationCertificate.
    actionFailureReason *VppTokenActionFailureReason
    // A count of the number of licenses that failed to revoke.
    failedLicensesCount *int32
    // A count of the number of licenses that were attempted to revoke.
    totalLicensesCount *int32
}
// NewVppTokenRevokeLicensesActionResult instantiates a new VppTokenRevokeLicensesActionResult and sets the default values.
func NewVppTokenRevokeLicensesActionResult()(*VppTokenRevokeLicensesActionResult) {
    m := &VppTokenRevokeLicensesActionResult{
        VppTokenActionResult: *NewVppTokenActionResult(),
    }
    return m
}
// CreateVppTokenRevokeLicensesActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVppTokenRevokeLicensesActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppTokenRevokeLicensesActionResult(), nil
}
// GetActionFailureReason gets the actionFailureReason property value. The reason for the revoke licenses action failure. Possible values are: none, appleFailure, internalError, expiredVppToken, expiredApplePushNotificationCertificate.
func (m *VppTokenRevokeLicensesActionResult) GetActionFailureReason()(*VppTokenActionFailureReason) {
    if m == nil {
        return nil
    } else {
        return m.actionFailureReason
    }
}
// GetFailedLicensesCount gets the failedLicensesCount property value. A count of the number of licenses that failed to revoke.
func (m *VppTokenRevokeLicensesActionResult) GetFailedLicensesCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedLicensesCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *VppTokenRevokeLicensesActionResult) GetTotalLicensesCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalLicensesCount
    }
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
// SetActionFailureReason sets the actionFailureReason property value. The reason for the revoke licenses action failure. Possible values are: none, appleFailure, internalError, expiredVppToken, expiredApplePushNotificationCertificate.
func (m *VppTokenRevokeLicensesActionResult) SetActionFailureReason(value *VppTokenActionFailureReason)() {
    if m != nil {
        m.actionFailureReason = value
    }
}
// SetFailedLicensesCount sets the failedLicensesCount property value. A count of the number of licenses that failed to revoke.
func (m *VppTokenRevokeLicensesActionResult) SetFailedLicensesCount(value *int32)() {
    if m != nil {
        m.failedLicensesCount = value
    }
}
// SetTotalLicensesCount sets the totalLicensesCount property value. A count of the number of licenses that were attempted to revoke.
func (m *VppTokenRevokeLicensesActionResult) SetTotalLicensesCount(value *int32)() {
    if m != nil {
        m.totalLicensesCount = value
    }
}
