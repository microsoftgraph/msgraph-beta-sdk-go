package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionsManagement 
type PermissionsManagement struct {
    Entity
}
// NewPermissionsManagement instantiates a new permissionsManagement and sets the default values.
func NewPermissionsManagement()(*PermissionsManagement) {
    m := &PermissionsManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionsManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionsManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsManagement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionsManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["permissionsRequestChanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionsRequestChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionsRequestChangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionsRequestChangeable)
                }
            }
            m.SetPermissionsRequestChanges(res)
        }
        return nil
    }
    res["scheduledPermissionsRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScheduledPermissionsRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScheduledPermissionsRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScheduledPermissionsRequestable)
                }
            }
            m.SetScheduledPermissionsRequests(res)
        }
        return nil
    }
    return res
}
// GetPermissionsRequestChanges gets the permissionsRequestChanges property value. Represents a change event of the scheduledPermissionsRequest entity.
func (m *PermissionsManagement) GetPermissionsRequestChanges()([]PermissionsRequestChangeable) {
    val, err := m.GetBackingStore().Get("permissionsRequestChanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionsRequestChangeable)
    }
    return nil
}
// GetScheduledPermissionsRequests gets the scheduledPermissionsRequests property value. Represents a permissions request that Permissions Management uses to manage permissions for an identity on resources in the authorization system. This request can be granted, rejected or canceled by identities in Permissions Management.
func (m *PermissionsManagement) GetScheduledPermissionsRequests()([]ScheduledPermissionsRequestable) {
    val, err := m.GetBackingStore().Get("scheduledPermissionsRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ScheduledPermissionsRequestable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPermissionsRequestChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionsRequestChanges()))
        for i, v := range m.GetPermissionsRequestChanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("permissionsRequestChanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScheduledPermissionsRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScheduledPermissionsRequests()))
        for i, v := range m.GetScheduledPermissionsRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("scheduledPermissionsRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPermissionsRequestChanges sets the permissionsRequestChanges property value. Represents a change event of the scheduledPermissionsRequest entity.
func (m *PermissionsManagement) SetPermissionsRequestChanges(value []PermissionsRequestChangeable)() {
    err := m.GetBackingStore().Set("permissionsRequestChanges", value)
    if err != nil {
        panic(err)
    }
}
// SetScheduledPermissionsRequests sets the scheduledPermissionsRequests property value. Represents a permissions request that Permissions Management uses to manage permissions for an identity on resources in the authorization system. This request can be granted, rejected or canceled by identities in Permissions Management.
func (m *PermissionsManagement) SetScheduledPermissionsRequests(value []ScheduledPermissionsRequestable)() {
    err := m.GetBackingStore().Set("scheduledPermissionsRequests", value)
    if err != nil {
        panic(err)
    }
}
// PermissionsManagementable 
type PermissionsManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPermissionsRequestChanges()([]PermissionsRequestChangeable)
    GetScheduledPermissionsRequests()([]ScheduledPermissionsRequestable)
    SetPermissionsRequestChanges(value []PermissionsRequestChangeable)()
    SetScheduledPermissionsRequests(value []ScheduledPermissionsRequestable)()
}
