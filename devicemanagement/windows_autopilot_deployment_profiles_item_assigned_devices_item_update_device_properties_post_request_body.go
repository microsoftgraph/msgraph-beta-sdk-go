package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody 
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
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
// NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody instantiates a new WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody and sets the default values.
func NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody()(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) {
    m := &WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetDeviceAccountPassword gets the deviceAccountPassword property value. The deviceAccountPassword property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetDeviceAccountPassword()(*string) {
    return m.deviceAccountPassword
}
// GetDeviceAccountUpn gets the deviceAccountUpn property value. The deviceAccountUpn property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetDeviceAccountUpn()(*string) {
    return m.deviceAccountUpn
}
// GetDeviceFriendlyName gets the deviceFriendlyName property value. The deviceFriendlyName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetDeviceFriendlyName()(*string) {
    return m.deviceFriendlyName
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetGroupTag()(*string) {
    return m.groupTag
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetDeviceAccountPassword sets the deviceAccountPassword property value. The deviceAccountPassword property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetDeviceAccountPassword(value *string)() {
    m.deviceAccountPassword = value
}
// SetDeviceAccountUpn sets the deviceAccountUpn property value. The deviceAccountUpn property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetDeviceAccountUpn(value *string)() {
    m.deviceAccountUpn = value
}
// SetDeviceFriendlyName sets the deviceFriendlyName property value. The deviceFriendlyName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetDeviceFriendlyName(value *string)() {
    m.deviceFriendlyName = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupTag sets the groupTag property value. The groupTag property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetGroupTag(value *string)() {
    m.groupTag = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
