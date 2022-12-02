package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody provides operations to call the setScheduledActions method.
type DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The scheduledActions property
    scheduledActions []DeviceManagementComplianceScheduledActionForRuleable
}
// NewDeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody instantiates a new DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody and sets the default values.
func NewDeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody()(*DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) {
    m := &DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["scheduledActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementComplianceScheduledActionForRuleable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementComplianceScheduledActionForRuleable)
            }
            m.SetScheduledActions(res)
        }
        return nil
    }
    return res
}
// GetScheduledActions gets the scheduledActions property value. The scheduledActions property
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) GetScheduledActions()([]DeviceManagementComplianceScheduledActionForRuleable) {
    return m.scheduledActions
}
// Serialize serializes information the current object
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetScheduledActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScheduledActions()))
        for i, v := range m.GetScheduledActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("scheduledActions", cast)
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
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetScheduledActions sets the scheduledActions property value. The scheduledActions property
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsPostRequestBody) SetScheduledActions(value []DeviceManagementComplianceScheduledActionForRuleable)() {
    m.scheduledActions = value
}
