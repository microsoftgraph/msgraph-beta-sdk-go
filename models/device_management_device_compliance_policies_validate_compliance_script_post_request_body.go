package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody provides operations to call the validateComplianceScript method.
type DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceCompliancePolicyScript property
    deviceCompliancePolicyScript DeviceCompliancePolicyScriptable
}
// NewDeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody instantiates a new DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody()(*DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) {
    m := &DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceCompliancePolicyScript gets the deviceCompliancePolicyScript property value. The deviceCompliancePolicyScript property
func (m *DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) GetDeviceCompliancePolicyScript()(DeviceCompliancePolicyScriptable) {
    return m.deviceCompliancePolicyScript
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceCompliancePolicyScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyScript(val.(DeviceCompliancePolicyScriptable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceCompliancePolicyScript", m.GetDeviceCompliancePolicyScript())
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
func (m *DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceCompliancePolicyScript sets the deviceCompliancePolicyScript property value. The deviceCompliancePolicyScript property
func (m *DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) SetDeviceCompliancePolicyScript(value DeviceCompliancePolicyScriptable)() {
    m.deviceCompliancePolicyScript = value
}
