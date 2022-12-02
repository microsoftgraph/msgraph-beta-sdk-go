package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCompliancePoliciesItemAssignPostRequestBody provides operations to call the assign method.
type DeviceManagementCompliancePoliciesItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignments property
    assignments []DeviceManagementConfigurationPolicyAssignmentable
}
// NewDeviceManagementCompliancePoliciesItemAssignPostRequestBody instantiates a new DeviceManagementCompliancePoliciesItemAssignPostRequestBody and sets the default values.
func NewDeviceManagementCompliancePoliciesItemAssignPostRequestBody()(*DeviceManagementCompliancePoliciesItemAssignPostRequestBody) {
    m := &DeviceManagementCompliancePoliciesItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementCompliancePoliciesItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCompliancePoliciesItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCompliancePoliciesItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementCompliancePoliciesItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *DeviceManagementCompliancePoliciesItemAssignPostRequestBody) GetAssignments()([]DeviceManagementConfigurationPolicyAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCompliancePoliciesItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementCompliancePoliciesItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
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
func (m *DeviceManagementCompliancePoliciesItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *DeviceManagementCompliancePoliciesItemAssignPostRequestBody) SetAssignments(value []DeviceManagementConfigurationPolicyAssignmentable)() {
    m.assignments = value
}
