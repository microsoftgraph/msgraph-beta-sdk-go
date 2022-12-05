package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody provides operations to call the assign method.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appProvisioningConfigurationGroupAssignments property
    appProvisioningConfigurationGroupAssignments []MobileAppProvisioningConfigGroupAssignmentable
    // The iOSLobAppProvisioningConfigAssignments property
    iOSLobAppProvisioningConfigAssignments []IosLobAppProvisioningConfigurationAssignmentable
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody instantiates a new DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody()(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) {
    m := &DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppProvisioningConfigurationGroupAssignments gets the appProvisioningConfigurationGroupAssignments property value. The appProvisioningConfigurationGroupAssignments property
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) GetAppProvisioningConfigurationGroupAssignments()([]MobileAppProvisioningConfigGroupAssignmentable) {
    return m.appProvisioningConfigurationGroupAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appProvisioningConfigurationGroupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppProvisioningConfigGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppProvisioningConfigGroupAssignmentable)
            }
            m.SetAppProvisioningConfigurationGroupAssignments(res)
        }
        return nil
    }
    res["iOSLobAppProvisioningConfigAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosLobAppProvisioningConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(IosLobAppProvisioningConfigurationAssignmentable)
            }
            m.SetIOSLobAppProvisioningConfigAssignments(res)
        }
        return nil
    }
    return res
}
// GetIOSLobAppProvisioningConfigAssignments gets the iOSLobAppProvisioningConfigAssignments property value. The iOSLobAppProvisioningConfigAssignments property
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) GetIOSLobAppProvisioningConfigAssignments()([]IosLobAppProvisioningConfigurationAssignmentable) {
    return m.iOSLobAppProvisioningConfigAssignments
}
// Serialize serializes information the current object
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppProvisioningConfigurationGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppProvisioningConfigurationGroupAssignments()))
        for i, v := range m.GetAppProvisioningConfigurationGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("appProvisioningConfigurationGroupAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIOSLobAppProvisioningConfigAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIOSLobAppProvisioningConfigAssignments()))
        for i, v := range m.GetIOSLobAppProvisioningConfigAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("iOSLobAppProvisioningConfigAssignments", cast)
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
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppProvisioningConfigurationGroupAssignments sets the appProvisioningConfigurationGroupAssignments property value. The appProvisioningConfigurationGroupAssignments property
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) SetAppProvisioningConfigurationGroupAssignments(value []MobileAppProvisioningConfigGroupAssignmentable)() {
    m.appProvisioningConfigurationGroupAssignments = value
}
// SetIOSLobAppProvisioningConfigAssignments sets the iOSLobAppProvisioningConfigAssignments property value. The iOSLobAppProvisioningConfigAssignments property
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignPostRequestBody) SetIOSLobAppProvisioningConfigAssignments(value []IosLobAppProvisioningConfigurationAssignmentable)() {
    m.iOSLobAppProvisioningConfigAssignments = value
}
