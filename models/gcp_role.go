package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GcpRole 
type GcpRole struct {
    Entity
}
// NewGcpRole instantiates a new gcpRole and sets the default values.
func NewGcpRole()(*GcpRole) {
    m := &GcpRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGcpRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGcpRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpRole(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *GcpRole) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The externalId property
func (m *GcpRole) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GcpRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["gcpRoleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGcpRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGcpRoleType(val.(*GcpRoleType))
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpScopeable)
                }
            }
            m.SetScopes(res)
        }
        return nil
    }
    return res
}
// GetGcpRoleType gets the gcpRoleType property value. The gcpRoleType property
func (m *GcpRole) GetGcpRoleType()(*GcpRoleType) {
    val, err := m.GetBackingStore().Get("gcpRoleType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GcpRoleType)
    }
    return nil
}
// GetScopes gets the scopes property value. The scopes property
func (m *GcpRole) GetScopes()([]GcpScopeable) {
    val, err := m.GetBackingStore().Get("scopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpScopeable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GcpRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetGcpRoleType() != nil {
        cast := (*m.GetGcpRoleType()).String()
        err = writer.WriteStringValue("gcpRoleType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScopes()))
        for i, v := range m.GetScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("scopes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *GcpRole) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The externalId property
func (m *GcpRole) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetGcpRoleType sets the gcpRoleType property value. The gcpRoleType property
func (m *GcpRole) SetGcpRoleType(value *GcpRoleType)() {
    err := m.GetBackingStore().Set("gcpRoleType", value)
    if err != nil {
        panic(err)
    }
}
// SetScopes sets the scopes property value. The scopes property
func (m *GcpRole) SetScopes(value []GcpScopeable)() {
    err := m.GetBackingStore().Set("scopes", value)
    if err != nil {
        panic(err)
    }
}
// GcpRoleable 
type GcpRoleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetGcpRoleType()(*GcpRoleType)
    GetScopes()([]GcpScopeable)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetGcpRoleType(value *GcpRoleType)()
    SetScopes(value []GcpScopeable)()
}
