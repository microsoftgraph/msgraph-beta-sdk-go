package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody provides operations to call the revokeLicenses method.
type DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The notifyManagedDevices property
    notifyManagedDevices *bool
    // The revokeUntrackedLicenses property
    revokeUntrackedLicenses *bool
}
// NewDeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody instantiates a new DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody and sets the default values.
func NewDeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody()(*DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) {
    m := &DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementVppTokensItemRevokeLicensesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementVppTokensItemRevokeLicensesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notifyManagedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifyManagedDevices(val)
        }
        return nil
    }
    res["revokeUntrackedLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevokeUntrackedLicenses(val)
        }
        return nil
    }
    return res
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) GetNotifyManagedDevices()(*bool) {
    return m.notifyManagedDevices
}
// GetRevokeUntrackedLicenses gets the revokeUntrackedLicenses property value. The revokeUntrackedLicenses property
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) GetRevokeUntrackedLicenses()(*bool) {
    return m.revokeUntrackedLicenses
}
// Serialize serializes information the current object
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) SetNotifyManagedDevices(value *bool)() {
    m.notifyManagedDevices = value
}
// SetRevokeUntrackedLicenses sets the revokeUntrackedLicenses property value. The revokeUntrackedLicenses property
func (m *DeviceAppManagementVppTokensItemRevokeLicensesPostRequestBody) SetRevokeUntrackedLicenses(value *bool)() {
    m.revokeUntrackedLicenses = value
}
