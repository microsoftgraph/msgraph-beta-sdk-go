package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse provides operations to call the setScheduledActions method.
type DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceManagementComplianceScheduledActionForRuleable
}
// NewDeviceManagementCompliancePoliciesItemSetScheduledActionsResponse instantiates a new DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse and sets the default values.
func NewDeviceManagementCompliancePoliciesItemSetScheduledActionsResponse()(*DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse) {
    m := &DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementCompliancePoliciesItemSetScheduledActionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCompliancePoliciesItemSetScheduledActionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCompliancePoliciesItemSetScheduledActionsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementComplianceScheduledActionForRuleable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementComplianceScheduledActionForRuleable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse) GetValue()([]DeviceManagementComplianceScheduledActionForRuleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceManagementCompliancePoliciesItemSetScheduledActionsResponse) SetValue(value []DeviceManagementComplianceScheduledActionForRuleable)() {
    m.value = value
}
