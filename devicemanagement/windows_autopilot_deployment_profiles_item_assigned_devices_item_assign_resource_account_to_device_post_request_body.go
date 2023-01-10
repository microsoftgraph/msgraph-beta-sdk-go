package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody 
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The addressableUserName property
    addressableUserName *string
    // The resourceAccountName property
    resourceAccountName *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody instantiates a new WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody and sets the default values.
func NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody()(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) {
    m := &WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resourceAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAccountName(val)
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
// GetResourceAccountName gets the resourceAccountName property value. The resourceAccountName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetResourceAccountName()(*string) {
    return m.resourceAccountName
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceAccountName", m.GetResourceAccountName())
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
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetResourceAccountName sets the resourceAccountName property value. The resourceAccountName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetResourceAccountName(value *string)() {
    m.resourceAccountName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
