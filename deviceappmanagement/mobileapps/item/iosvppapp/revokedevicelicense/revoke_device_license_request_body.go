package revokedevicelicense

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RevokeDeviceLicenseRequestBody provides operations to call the revokeDeviceLicense method.
type RevokeDeviceLicenseRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedDeviceId property
    managedDeviceId *string
    // The notifyManagedDevices property
    notifyManagedDevices *bool
}
// NewRevokeDeviceLicenseRequestBody instantiates a new revokeDeviceLicenseRequestBody and sets the default values.
func NewRevokeDeviceLicenseRequestBody()(*RevokeDeviceLicenseRequestBody) {
    m := &RevokeDeviceLicenseRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRevokeDeviceLicenseRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeDeviceLicenseRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRevokeDeviceLicenseRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeDeviceLicenseRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeDeviceLicenseRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
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
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. The managedDeviceId property
func (m *RevokeDeviceLicenseRequestBody) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeDeviceLicenseRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
// Serialize serializes information the current object
func (m *RevokeDeviceLicenseRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
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
func (m *RevokeDeviceLicenseRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. The managedDeviceId property
func (m *RevokeDeviceLicenseRequestBody) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeDeviceLicenseRequestBody) SetNotifyManagedDevices(value *bool)() {
    if m != nil {
        m.notifyManagedDevices = value
    }
}
