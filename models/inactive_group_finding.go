package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InactiveGroupFinding 
type InactiveGroupFinding struct {
    Finding
}
// NewInactiveGroupFinding instantiates a new inactiveGroupFinding and sets the default values.
func NewInactiveGroupFinding()(*InactiveGroupFinding) {
    m := &InactiveGroupFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateInactiveGroupFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInactiveGroupFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInactiveGroupFinding(), nil
}
// GetActionSummary gets the actionSummary property value. The actionSummary property
func (m *InactiveGroupFinding) GetActionSummary()(ActionSummaryable) {
    val, err := m.GetBackingStore().Get("actionSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ActionSummaryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InactiveGroupFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["actionSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSummary(val.(ActionSummaryable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(AuthorizationSystemIdentityable))
        }
        return nil
    }
    res["permissionsCreepIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsCreepIndexFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsCreepIndex(val.(PermissionsCreepIndexable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *InactiveGroupFinding) GetGroup()(AuthorizationSystemIdentityable) {
    val, err := m.GetBackingStore().Get("group")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemIdentityable)
    }
    return nil
}
// GetPermissionsCreepIndex gets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *InactiveGroupFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InactiveGroupFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("actionSummary", m.GetActionSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("permissionsCreepIndex", m.GetPermissionsCreepIndex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionSummary sets the actionSummary property value. The actionSummary property
func (m *InactiveGroupFinding) SetActionSummary(value ActionSummaryable)() {
    err := m.GetBackingStore().Set("actionSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetGroup sets the group property value. The group property
func (m *InactiveGroupFinding) SetGroup(value AuthorizationSystemIdentityable)() {
    err := m.GetBackingStore().Set("group", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *InactiveGroupFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// InactiveGroupFindingable 
type InactiveGroupFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSummary()(ActionSummaryable)
    GetGroup()(AuthorizationSystemIdentityable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    SetActionSummary(value ActionSummaryable)()
    SetGroup(value AuthorizationSystemIdentityable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
}
