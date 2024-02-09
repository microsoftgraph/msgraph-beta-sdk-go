package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureAuthorizationSystem struct {
    AuthorizationSystem
}
// NewAzureAuthorizationSystem instantiates a new AzureAuthorizationSystem and sets the default values.
func NewAzureAuthorizationSystem()(*AzureAuthorizationSystem) {
    m := &AzureAuthorizationSystem{
        AuthorizationSystem: *NewAuthorizationSystem(),
    }
    odataTypeValue := "#microsoft.graph.azureAuthorizationSystem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureAuthorizationSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureAuthorizationSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureAuthorizationSystem(), nil
}
// GetActions gets the actions property value. List of actions for service in authorization system.
// returns a []AzureAuthorizationSystemTypeActionable when successful
func (m *AzureAuthorizationSystem) GetActions()([]AzureAuthorizationSystemTypeActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureAuthorizationSystemTypeActionable)
    }
    return nil
}
// GetAssociatedIdentities gets the associatedIdentities property value. Identities in the authorization system.
// returns a AzureAssociatedIdentitiesable when successful
func (m *AzureAuthorizationSystem) GetAssociatedIdentities()(AzureAssociatedIdentitiesable) {
    val, err := m.GetBackingStore().Get("associatedIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AzureAssociatedIdentitiesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureAuthorizationSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystem.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureAuthorizationSystemTypeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureAuthorizationSystemTypeActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureAuthorizationSystemTypeActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["associatedIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureAssociatedIdentitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociatedIdentities(val.(AzureAssociatedIdentitiesable))
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureAuthorizationSystemResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureAuthorizationSystemResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureRoleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureRoleDefinitionable)
                }
            }
            m.SetRoleDefinitions(res)
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
// returns a []AzureAuthorizationSystemResourceable when successful
func (m *AzureAuthorizationSystem) GetResources()([]AzureAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureAuthorizationSystemResourceable)
    }
    return nil
}
// GetRoleDefinitions gets the roleDefinitions property value. Roles associated with the authorization system type.
// returns a []AzureRoleDefinitionable when successful
func (m *AzureAuthorizationSystem) GetRoleDefinitions()([]AzureRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureRoleDefinitionable)
    }
    return nil
}
// GetServices gets the services property value. Services associated with the authorization system type.
// returns a []AuthorizationSystemTypeServiceable when successful
func (m *AzureAuthorizationSystem) GetServices()([]AuthorizationSystemTypeServiceable) {
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
func (m *AzureAuthorizationSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
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
func (m *AzureAuthorizationSystem) SetActions(value []AzureAuthorizationSystemTypeActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAssociatedIdentities sets the associatedIdentities property value. Identities in the authorization system.
func (m *AzureAuthorizationSystem) SetAssociatedIdentities(value AzureAssociatedIdentitiesable)() {
    err := m.GetBackingStore().Set("associatedIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. Resources associated with the authorization system type.
func (m *AzureAuthorizationSystem) SetResources(value []AzureAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. Roles associated with the authorization system type.
func (m *AzureAuthorizationSystem) SetRoleDefinitions(value []AzureRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetServices sets the services property value. Services associated with the authorization system type.
func (m *AzureAuthorizationSystem) SetServices(value []AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("services", value)
    if err != nil {
        panic(err)
    }
}
type AzureAuthorizationSystemable interface {
    AuthorizationSystemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]AzureAuthorizationSystemTypeActionable)
    GetAssociatedIdentities()(AzureAssociatedIdentitiesable)
    GetResources()([]AzureAuthorizationSystemResourceable)
    GetRoleDefinitions()([]AzureRoleDefinitionable)
    GetServices()([]AuthorizationSystemTypeServiceable)
    SetActions(value []AzureAuthorizationSystemTypeActionable)()
    SetAssociatedIdentities(value AzureAssociatedIdentitiesable)()
    SetResources(value []AzureAuthorizationSystemResourceable)()
    SetRoleDefinitions(value []AzureRoleDefinitionable)()
    SetServices(value []AuthorizationSystemTypeServiceable)()
}
