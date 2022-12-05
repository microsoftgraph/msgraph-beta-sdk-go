package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody provides operations to call the deleteUserFromSharedAppleDevice method.
type DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The userPrincipalName property
    userPrincipalName *string
}
// NewDeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody instantiates a new DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody and sets the default values.
func NewDeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody()(*DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) {
    m := &DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
