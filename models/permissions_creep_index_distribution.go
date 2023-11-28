package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionsCreepIndexDistribution 
type PermissionsCreepIndexDistribution struct {
    Entity
}
// NewPermissionsCreepIndexDistribution instantiates a new permissionsCreepIndexDistribution and sets the default values.
func NewPermissionsCreepIndexDistribution()(*PermissionsCreepIndexDistribution) {
    m := &PermissionsCreepIndexDistribution{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionsCreepIndexDistributionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionsCreepIndexDistributionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsCreepIndexDistribution(), nil
}
// GetAuthorizationSystem gets the authorizationSystem property value. The authorizationSystem property
func (m *PermissionsCreepIndexDistribution) GetAuthorizationSystem()(AuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("authorizationSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Defines when the PCI distribution was created.
func (m *PermissionsCreepIndexDistribution) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionsCreepIndexDistribution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authorizationSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystem(val.(AuthorizationSystemable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["highRiskProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRiskProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighRiskProfile(val.(RiskProfileable))
        }
        return nil
    }
    res["lowRiskProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRiskProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowRiskProfile(val.(RiskProfileable))
        }
        return nil
    }
    res["mediumRiskProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRiskProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediumRiskProfile(val.(RiskProfileable))
        }
        return nil
    }
    return res
}
// GetHighRiskProfile gets the highRiskProfile property value. The highRiskProfile property
func (m *PermissionsCreepIndexDistribution) GetHighRiskProfile()(RiskProfileable) {
    val, err := m.GetBackingStore().Get("highRiskProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RiskProfileable)
    }
    return nil
}
// GetLowRiskProfile gets the lowRiskProfile property value. The lowRiskProfile property
func (m *PermissionsCreepIndexDistribution) GetLowRiskProfile()(RiskProfileable) {
    val, err := m.GetBackingStore().Get("lowRiskProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RiskProfileable)
    }
    return nil
}
// GetMediumRiskProfile gets the mediumRiskProfile property value. The mediumRiskProfile property
func (m *PermissionsCreepIndexDistribution) GetMediumRiskProfile()(RiskProfileable) {
    val, err := m.GetBackingStore().Get("mediumRiskProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RiskProfileable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsCreepIndexDistribution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authorizationSystem", m.GetAuthorizationSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("highRiskProfile", m.GetHighRiskProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lowRiskProfile", m.GetLowRiskProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediumRiskProfile", m.GetMediumRiskProfile())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizationSystem sets the authorizationSystem property value. The authorizationSystem property
func (m *PermissionsCreepIndexDistribution) SetAuthorizationSystem(value AuthorizationSystemable)() {
    err := m.GetBackingStore().Set("authorizationSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Defines when the PCI distribution was created.
func (m *PermissionsCreepIndexDistribution) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetHighRiskProfile sets the highRiskProfile property value. The highRiskProfile property
func (m *PermissionsCreepIndexDistribution) SetHighRiskProfile(value RiskProfileable)() {
    err := m.GetBackingStore().Set("highRiskProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetLowRiskProfile sets the lowRiskProfile property value. The lowRiskProfile property
func (m *PermissionsCreepIndexDistribution) SetLowRiskProfile(value RiskProfileable)() {
    err := m.GetBackingStore().Set("lowRiskProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetMediumRiskProfile sets the mediumRiskProfile property value. The mediumRiskProfile property
func (m *PermissionsCreepIndexDistribution) SetMediumRiskProfile(value RiskProfileable)() {
    err := m.GetBackingStore().Set("mediumRiskProfile", value)
    if err != nil {
        panic(err)
    }
}
// PermissionsCreepIndexDistributionable 
type PermissionsCreepIndexDistributionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizationSystem()(AuthorizationSystemable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHighRiskProfile()(RiskProfileable)
    GetLowRiskProfile()(RiskProfileable)
    GetMediumRiskProfile()(RiskProfileable)
    SetAuthorizationSystem(value AuthorizationSystemable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHighRiskProfile(value RiskProfileable)()
    SetLowRiskProfile(value RiskProfileable)()
    SetMediumRiskProfile(value RiskProfileable)()
}
