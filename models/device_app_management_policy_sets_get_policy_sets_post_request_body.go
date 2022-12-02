package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody provides operations to call the getPolicySets method.
type DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The policySetIds property
    policySetIds []string
}
// NewDeviceAppManagementPolicySetsGetPolicySetsPostRequestBody instantiates a new DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody and sets the default values.
func NewDeviceAppManagementPolicySetsGetPolicySetsPostRequestBody()(*DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) {
    m := &DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementPolicySetsGetPolicySetsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementPolicySetsGetPolicySetsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementPolicySetsGetPolicySetsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["policySetIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPolicySetIds(res)
        }
        return nil
    }
    return res
}
// GetPolicySetIds gets the policySetIds property value. The policySetIds property
func (m *DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) GetPolicySetIds()([]string) {
    return m.policySetIds
}
// Serialize serializes information the current object
func (m *DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPolicySetIds() != nil {
        err := writer.WriteCollectionOfStringValues("policySetIds", m.GetPolicySetIds())
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
func (m *DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPolicySetIds sets the policySetIds property value. The policySetIds property
func (m *DeviceAppManagementPolicySetsGetPolicySetsPostRequestBody) SetPolicySetIds(value []string)() {
    m.policySetIds = value
}
