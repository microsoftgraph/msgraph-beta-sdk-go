package revokeuserlicense

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RevokeUserLicensePostRequestBody provides operations to call the revokeUserLicense method.
type RevokeUserLicensePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The notifyManagedDevices property
    notifyManagedDevices *bool
    // The userId property
    userId *string
}
// NewRevokeUserLicensePostRequestBody instantiates a new revokeUserLicensePostRequestBody and sets the default values.
func NewRevokeUserLicensePostRequestBody()(*RevokeUserLicensePostRequestBody) {
    m := &RevokeUserLicensePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRevokeUserLicensePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeUserLicensePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRevokeUserLicensePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeUserLicensePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeUserLicensePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeUserLicensePostRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
// GetUserId gets the userId property value. The userId property
func (m *RevokeUserLicensePostRequestBody) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Serialize serializes information the current object
func (m *RevokeUserLicensePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *RevokeUserLicensePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. The notifyManagedDevices property
func (m *RevokeUserLicensePostRequestBody) SetNotifyManagedDevices(value *bool)() {
    if m != nil {
        m.notifyManagedDevices = value
    }
}
// SetUserId sets the userId property value. The userId property
func (m *RevokeUserLicensePostRequestBody) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
