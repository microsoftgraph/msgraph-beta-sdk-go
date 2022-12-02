package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody provides operations to call the triggerDeviceScopeAction method.
type DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody struct {
    // Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
    actionName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceScopeId property
    deviceScopeId *string
}
// NewDeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody instantiates a new DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody()(*DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) {
    m := &DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) GetActionName()(*string) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceScopeId gets the deviceScopeId property value. The deviceScopeId property
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) GetDeviceScopeId()(*string) {
    return m.deviceScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
    res["deviceScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceScopeId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceScopeId", m.GetDeviceScopeId())
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
// SetActionName sets the actionName property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) SetActionName(value *string)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceScopeId sets the deviceScopeId property value. The deviceScopeId property
func (m *DeviceManagementUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBody) SetDeviceScopeId(value *string)() {
    m.deviceScopeId = value
}
