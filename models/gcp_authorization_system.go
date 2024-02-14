package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GcpAuthorizationSystem struct {
    AuthorizationSystem
}
// NewGcpAuthorizationSystem instantiates a new GcpAuthorizationSystem and sets the default values.
func NewGcpAuthorizationSystem()(*GcpAuthorizationSystem) {
    m := &GcpAuthorizationSystem{
        AuthorizationSystem: *NewAuthorizationSystem(),
    }
    odataTypeValue := "#microsoft.graph.gcpAuthorizationSystem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpAuthorizationSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGcpAuthorizationSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpAuthorizationSystem(), nil
}
// GetActions gets the actions property value. List of actions for service in authorization system.
// returns a []GcpAuthorizationSystemTypeActionable when successful
func (m *GcpAuthorizationSystem) GetActions()([]GcpAuthorizationSystemTypeActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpAuthorizationSystemTypeActionable)
    }
    return nil
}
// GetAssociatedIdentities gets the associatedIdentities property value. Identities in the authorization system.
// returns a GcpAssociatedIdentitiesable when successful
func (m *GcpAuthorizationSystem) GetAssociatedIdentities()(GcpAssociatedIdentitiesable) {
    val, err := m.GetBackingStore().Get("associatedIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GcpAssociatedIdentitiesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GcpAuthorizationSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystem.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpAuthorizationSystemTypeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpAuthorizationSystemTypeActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpAuthorizationSystemTypeActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["associatedIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGcpAssociatedIdentitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociatedIdentities(val.(GcpAssociatedIdentitiesable))
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpAuthorizationSystemResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpAuthorizationSystemResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpRoleable)
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationSystemTypeServiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationSystemTypeServiceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationSystemTypeServiceable)
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. Resources associated with the authorization system type.
// returns a []GcpAuthorizationSystemResourceable when successful
func (m *GcpAuthorizationSystem) GetResources()([]GcpAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpAuthorizationSystemResourceable)
    }
    return nil
}
// GetRoles gets the roles property value. Roles associated with the authorization system type.
// returns a []GcpRoleable when successful
func (m *GcpAuthorizationSystem) GetRoles()([]GcpRoleable) {
    val, err := m.GetBackingStore().Get("roles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpRoleable)
    }
    return nil
}
// GetServices gets the services property value. Services associated with the authorization system type.
// returns a []AuthorizationSystemTypeServiceable when successful
func (m *GcpAuthorizationSystem) GetServices()([]AuthorizationSystemTypeServiceable) {
    val, err := m.GetBackingStore().Get("services")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationSystemTypeServiceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GcpAuthorizationSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("associatedIdentities", m.GetAssociatedIdentities())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. List of actions for service in authorization system.
func (m *GcpAuthorizationSystem) SetActions(value []GcpAuthorizationSystemTypeActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAssociatedIdentities sets the associatedIdentities property value. Identities in the authorization system.
func (m *GcpAuthorizationSystem) SetAssociatedIdentities(value GcpAssociatedIdentitiesable)() {
    err := m.GetBackingStore().Set("associatedIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. Resources associated with the authorization system type.
func (m *GcpAuthorizationSystem) SetResources(value []GcpAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// SetRoles sets the roles property value. Roles associated with the authorization system type.
func (m *GcpAuthorizationSystem) SetRoles(value []GcpRoleable)() {
    err := m.GetBackingStore().Set("roles", value)
    if err != nil {
        panic(err)
    }
}
// SetServices sets the services property value. Services associated with the authorization system type.
func (m *GcpAuthorizationSystem) SetServices(value []AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("services", value)
    if err != nil {
        panic(err)
    }
}
type GcpAuthorizationSystemable interface {
    AuthorizationSystemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]GcpAuthorizationSystemTypeActionable)
    GetAssociatedIdentities()(GcpAssociatedIdentitiesable)
    GetResources()([]GcpAuthorizationSystemResourceable)
    GetRoles()([]GcpRoleable)
    GetServices()([]AuthorizationSystemTypeServiceable)
    SetActions(value []GcpAuthorizationSystemTypeActionable)()
    SetAssociatedIdentities(value GcpAssociatedIdentitiesable)()
    SetResources(value []GcpAuthorizationSystemResourceable)()
    SetRoles(value []GcpRoleable)()
    SetServices(value []AuthorizationSystemTypeServiceable)()
}
