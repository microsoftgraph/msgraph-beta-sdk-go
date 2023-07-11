package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleSummary 
type PrivilegedRoleSummary struct {
    Entity
}
// NewPrivilegedRoleSummary instantiates a new privilegedRoleSummary and sets the default values.
func NewPrivilegedRoleSummary()(*PrivilegedRoleSummary) {
    m := &PrivilegedRoleSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleSummary(), nil
}
// GetElevatedCount gets the elevatedCount property value. The elevatedCount property
func (m *PrivilegedRoleSummary) GetElevatedCount()(*int32) {
    val, err := m.GetBackingStore().Get("elevatedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["elevatedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElevatedCount(val)
        }
        return nil
    }
    res["managedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedCount(val)
        }
        return nil
    }
    res["mfaEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoleSummaryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RoleSummaryStatus))
        }
        return nil
    }
    res["usersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersCount(val)
        }
        return nil
    }
    return res
}
// GetManagedCount gets the managedCount property value. The managedCount property
func (m *PrivilegedRoleSummary) GetManagedCount()(*int32) {
    val, err := m.GetBackingStore().Get("managedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMfaEnabled gets the mfaEnabled property value. The mfaEnabled property
func (m *PrivilegedRoleSummary) GetMfaEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("mfaEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrivilegedRoleSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *PrivilegedRoleSummary) GetStatus()(*RoleSummaryStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RoleSummaryStatus)
    }
    return nil
}
// GetUsersCount gets the usersCount property value. The usersCount property
func (m *PrivilegedRoleSummary) GetUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("usersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegedRoleSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("elevatedCount", m.GetElevatedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("managedCount", m.GetManagedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mfaEnabled", m.GetMfaEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usersCount", m.GetUsersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetElevatedCount sets the elevatedCount property value. The elevatedCount property
func (m *PrivilegedRoleSummary) SetElevatedCount(value *int32)() {
    err := m.GetBackingStore().Set("elevatedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedCount sets the managedCount property value. The managedCount property
func (m *PrivilegedRoleSummary) SetManagedCount(value *int32)() {
    err := m.GetBackingStore().Set("managedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaEnabled sets the mfaEnabled property value. The mfaEnabled property
func (m *PrivilegedRoleSummary) SetMfaEnabled(value *bool)() {
    err := m.GetBackingStore().Set("mfaEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrivilegedRoleSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *PrivilegedRoleSummary) SetStatus(value *RoleSummaryStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUsersCount sets the usersCount property value. The usersCount property
func (m *PrivilegedRoleSummary) SetUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("usersCount", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegedRoleSummaryable 
type PrivilegedRoleSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetElevatedCount()(*int32)
    GetManagedCount()(*int32)
    GetMfaEnabled()(*bool)
    GetOdataType()(*string)
    GetStatus()(*RoleSummaryStatus)
    GetUsersCount()(*int32)
    SetElevatedCount(value *int32)()
    SetManagedCount(value *int32)()
    SetMfaEnabled(value *bool)()
    SetOdataType(value *string)()
    SetStatus(value *RoleSummaryStatus)()
    SetUsersCount(value *int32)()
}
