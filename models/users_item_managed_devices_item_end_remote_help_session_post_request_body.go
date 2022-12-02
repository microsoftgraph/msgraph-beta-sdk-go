package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody provides operations to call the endRemoteHelpSession method.
type UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The sessionKey property
    sessionKey *string
}
// NewUsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody instantiates a new UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody and sets the default values.
func NewUsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody()(*UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) {
    m := &UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sessionKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionKey(val)
        }
        return nil
    }
    return res
}
// GetSessionKey gets the sessionKey property value. The sessionKey property
func (m *UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) GetSessionKey()(*string) {
    return m.sessionKey
}
// Serialize serializes information the current object
func (m *UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sessionKey", m.GetSessionKey())
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
func (m *UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSessionKey sets the sessionKey property value. The sessionKey property
func (m *UsersItemManagedDevicesItemEndRemoteHelpSessionPostRequestBody) SetSessionKey(value *string)() {
    m.sessionKey = value
}
