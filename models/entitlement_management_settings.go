package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementSettings 
type EntitlementManagementSettings struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewEntitlementManagementSettings instantiates a new entitlementManagementSettings and sets the default values.
func NewEntitlementManagementSettings()(*EntitlementManagementSettings) {
    m := &EntitlementManagementSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEntitlementManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementSettings(), nil
}
// GetDaysUntilExternalUserDeletedAfterBlocked gets the daysUntilExternalUserDeletedAfterBlocked property value. If externalUserLifecycleAction is BlockSignInAndDelete, the number of days after an external user is blocked from sign in before their account is deleted.
func (m *EntitlementManagementSettings) GetDaysUntilExternalUserDeletedAfterBlocked()(*int32) {
    val, err := m.GetBackingStore().Get("daysUntilExternalUserDeletedAfterBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetExternalUserLifecycleAction gets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
func (m *EntitlementManagementSettings) GetExternalUserLifecycleAction()(*string) {
    val, err := m.GetBackingStore().Get("externalUserLifecycleAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["daysUntilExternalUserDeletedAfterBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaysUntilExternalUserDeletedAfterBlocked(val)
        }
        return nil
    }
    res["externalUserLifecycleAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserLifecycleAction(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EntitlementManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("daysUntilExternalUserDeletedAfterBlocked", m.GetDaysUntilExternalUserDeletedAfterBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUserLifecycleAction", m.GetExternalUserLifecycleAction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDaysUntilExternalUserDeletedAfterBlocked sets the daysUntilExternalUserDeletedAfterBlocked property value. If externalUserLifecycleAction is BlockSignInAndDelete, the number of days after an external user is blocked from sign in before their account is deleted.
func (m *EntitlementManagementSettings) SetDaysUntilExternalUserDeletedAfterBlocked(value *int32)() {
    err := m.GetBackingStore().Set("daysUntilExternalUserDeletedAfterBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalUserLifecycleAction sets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
func (m *EntitlementManagementSettings) SetExternalUserLifecycleAction(value *string)() {
    err := m.GetBackingStore().Set("externalUserLifecycleAction", value)
    if err != nil {
        panic(err)
    }
}
// EntitlementManagementSettingsable 
type EntitlementManagementSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDaysUntilExternalUserDeletedAfterBlocked()(*int32)
    GetExternalUserLifecycleAction()(*string)
    SetDaysUntilExternalUserDeletedAfterBlocked(value *int32)()
    SetExternalUserLifecycleAction(value *string)()
}
