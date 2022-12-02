package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody provides operations to call the overrideComplianceState method.
type DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Administrator configured device compliance state Enum
    complianceState *AdministratorConfiguredDeviceComplianceState
    // The remediationUrl property
    remediationUrl *string
}
// NewDeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody instantiates a new DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody and sets the default values.
func NewDeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody()(*DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) {
    m := &DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComplianceState gets the complianceState property value. Administrator configured device compliance state Enum
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) GetComplianceState()(*AdministratorConfiguredDeviceComplianceState) {
    return m.complianceState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdministratorConfiguredDeviceComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*AdministratorConfiguredDeviceComplianceState))
        }
        return nil
    }
    res["remediationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationUrl(val)
        }
        return nil
    }
    return res
}
// GetRemediationUrl gets the remediationUrl property value. The remediationUrl property
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) GetRemediationUrl()(*string) {
    return m.remediationUrl
}
// Serialize serializes information the current object
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err := writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediationUrl", m.GetRemediationUrl())
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
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComplianceState sets the complianceState property value. Administrator configured device compliance state Enum
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) SetComplianceState(value *AdministratorConfiguredDeviceComplianceState)() {
    m.complianceState = value
}
// SetRemediationUrl sets the remediationUrl property value. The remediationUrl property
func (m *DeviceManagementManagedDevicesItemOverrideComplianceStatePostRequestBody) SetRemediationUrl(value *string)() {
    m.remediationUrl = value
}
