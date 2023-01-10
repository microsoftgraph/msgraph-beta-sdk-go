package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody 
type WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The addressableUserName property
    addressableUserName *string
    // The deviceAccountPassword property
    deviceAccountPassword *string
    // The deviceAccountUpn property
    deviceAccountUpn *string
    // The deviceFriendlyName property
    deviceFriendlyName *string
    // The displayName property
    displayName *string
    // The groupTag property
    groupTag *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewWindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody instantiates a new WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody()(*WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) {
    m := &WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetDeviceAccountPassword gets the deviceAccountPassword property value. The deviceAccountPassword property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetDeviceAccountPassword()(*string) {
    return m.deviceAccountPassword
}
// GetDeviceAccountUpn gets the deviceAccountUpn property value. The deviceAccountUpn property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetDeviceAccountUpn()(*string) {
    return m.deviceAccountUpn
}
// GetDeviceFriendlyName gets the deviceFriendlyName property value. The deviceFriendlyName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetDeviceFriendlyName()(*string) {
    return m.deviceFriendlyName
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addressableUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["deviceAccountPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountPassword(val)
        }
        return nil
    }
    res["deviceAccountUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountUpn(val)
        }
        return nil
    }
    res["deviceFriendlyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFriendlyName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groupTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
        }
        return nil
    }
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
// GetGroupTag gets the groupTag property value. The groupTag property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetGroupTag()(*string) {
    return m.groupTag
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceAccountPassword", m.GetDeviceAccountPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceAccountUpn", m.GetDeviceAccountUpn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceFriendlyName", m.GetDeviceFriendlyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupTag", m.GetGroupTag())
        if err != nil {
            return err
        }
    }
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
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetDeviceAccountPassword sets the deviceAccountPassword property value. The deviceAccountPassword property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetDeviceAccountPassword(value *string)() {
    m.deviceAccountPassword = value
}
// SetDeviceAccountUpn sets the deviceAccountUpn property value. The deviceAccountUpn property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetDeviceAccountUpn(value *string)() {
    m.deviceAccountUpn = value
}
// SetDeviceFriendlyName sets the deviceFriendlyName property value. The deviceFriendlyName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetDeviceFriendlyName(value *string)() {
    m.deviceFriendlyName = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupTag sets the groupTag property value. The groupTag property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetGroupTag(value *string)() {
    m.groupTag = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeviceIdentitiesItemUpdateDevicePropertiesPostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
