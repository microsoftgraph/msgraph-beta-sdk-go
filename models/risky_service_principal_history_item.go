package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyServicePrincipalHistoryItem provides operations to manage the collection of activityStatistics entities.
type RiskyServicePrincipalHistoryItem struct {
    RiskyServicePrincipal
    // The activity related to service principal risk level change.
    activity RiskServicePrincipalActivityable
    // The identifier of the actor of the operation.
    initiatedBy *string
    // The identifier of the service principal.
    servicePrincipalId *string
}
// NewRiskyServicePrincipalHistoryItem instantiates a new riskyServicePrincipalHistoryItem and sets the default values.
func NewRiskyServicePrincipalHistoryItem()(*RiskyServicePrincipalHistoryItem) {
    m := &RiskyServicePrincipalHistoryItem{
        RiskyServicePrincipal: *NewRiskyServicePrincipal(),
    }
    odataTypeValue := "#microsoft.graph.riskyServicePrincipalHistoryItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskyServicePrincipalHistoryItem(), nil
}
// GetActivity gets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) GetActivity()(RiskServicePrincipalActivityable) {
    return m.activity
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyServicePrincipalHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RiskyServicePrincipal.GetFieldDeserializers()
    res["activity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRiskServicePrincipalActivityFromDiscriminatorValue , m.SetActivity)
    res["initiatedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInitiatedBy)
    res["servicePrincipalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServicePrincipalId)
    return res
}
// GetInitiatedBy gets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) GetInitiatedBy()(*string) {
    return m.initiatedBy
}
// GetServicePrincipalId gets the servicePrincipalId property value. The identifier of the service principal.
func (m *RiskyServicePrincipalHistoryItem) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// Serialize serializes information the current object
func (m *RiskyServicePrincipalHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RiskyServicePrincipal.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) SetActivity(value RiskServicePrincipalActivityable)() {
    m.activity = value
}
// SetInitiatedBy sets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) SetInitiatedBy(value *string)() {
    m.initiatedBy = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. The identifier of the service principal.
func (m *RiskyServicePrincipalHistoryItem) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
