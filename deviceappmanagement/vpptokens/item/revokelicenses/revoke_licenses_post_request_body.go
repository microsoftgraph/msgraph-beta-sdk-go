package revokelicenses

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RevokeLicensesPostRequestBody provides operations to call the revokeLicenses method.
type RevokeLicensesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The notifyManagedDevices property
    notifyManagedDevices *bool
    // The revokeUntrackedLicenses property
    revokeUntrackedLicenses *bool
}
// NewRevokeLicensesPostRequestBody instantiates a new revokeLicensesPostRequestBody and sets the default values.
func NewRevokeLicensesPostRequestBody()(*RevokeLicensesPostRequestBody) {
    m := &RevokeLicensesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRevokeLicensesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeLicensesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRevokeLicensesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeLicensesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeLicensesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notifyManagedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetNotifyManagedDevices)
    res["revokeUntrackedLicenses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRevokeUntrackedLicenses)
    return res
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeLicensesPostRequestBody) GetNotifyManagedDevices()(*bool) {
    return m.notifyManagedDevices
}
// GetRevokeUntrackedLicenses gets the revokeUntrackedLicenses property value. The revokeUntrackedLicenses property
func (m *RevokeLicensesPostRequestBody) GetRevokeUntrackedLicenses()(*bool) {
    return m.revokeUntrackedLicenses
}
// Serialize serializes information the current object
func (m *RevokeLicensesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("revokeUntrackedLicenses", m.GetRevokeUntrackedLicenses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeLicensesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeLicensesPostRequestBody) SetNotifyManagedDevices(value *bool)() {
    m.notifyManagedDevices = value
}
// SetRevokeUntrackedLicenses sets the revokeUntrackedLicenses property value. The revokeUntrackedLicenses property
func (m *RevokeLicensesPostRequestBody) SetRevokeUntrackedLicenses(value *bool)() {
    m.revokeUntrackedLicenses = value
}
