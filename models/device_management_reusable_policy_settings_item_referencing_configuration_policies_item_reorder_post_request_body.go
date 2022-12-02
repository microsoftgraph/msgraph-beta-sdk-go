package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody provides operations to call the reorder method.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The priority property
    priority *int32
}
// NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody instantiates a new DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody and sets the default values.
func NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) {
    m := &DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority property
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPriority sets the priority property value. The priority property
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderPostRequestBody) SetPriority(value *int32)() {
    m.priority = value
}
