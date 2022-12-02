package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody provides operations to call the assign method.
type DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The wdacPolicyAssignments property
    wdacPolicyAssignments []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable
}
// NewDeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody instantiates a new DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody and sets the default values.
func NewDeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody()(*DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) {
    m := &DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["wdacPolicyAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)
            }
            m.SetWdacPolicyAssignments(res)
        }
        return nil
    }
    return res
}
// GetWdacPolicyAssignments gets the wdacPolicyAssignments property value. The wdacPolicyAssignments property
func (m *DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) GetWdacPolicyAssignments()([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable) {
    return m.wdacPolicyAssignments
}
// Serialize serializes information the current object
func (m *DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWdacPolicyAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWdacPolicyAssignments()))
        for i, v := range m.GetWdacPolicyAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("wdacPolicyAssignments", cast)
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
func (m *DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetWdacPolicyAssignments sets the wdacPolicyAssignments property value. The wdacPolicyAssignments property
func (m *DeviceAppManagementWdacSupplementalPoliciesItemAssignPostRequestBody) SetWdacPolicyAssignments(value []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)() {
    m.wdacPolicyAssignments = value
}
