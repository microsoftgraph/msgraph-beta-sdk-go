package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsAuthorizationSystem struct {
    AuthorizationSystem
}
// NewAwsAuthorizationSystem instantiates a new AwsAuthorizationSystem and sets the default values.
func NewAwsAuthorizationSystem()(*AwsAuthorizationSystem) {
    m := &AwsAuthorizationSystem{
        AuthorizationSystem: *NewAuthorizationSystem(),
    }
    odataTypeValue := "#microsoft.graph.awsAuthorizationSystem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsAuthorizationSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsAuthorizationSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsAuthorizationSystem(), nil
}
// GetActions gets the actions property value. List of actions for service in authorization system.
// returns a []AwsAuthorizationSystemTypeActionable when successful
func (m *AwsAuthorizationSystem) GetActions()([]AwsAuthorizationSystemTypeActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AwsAuthorizationSystemTypeActionable)
    }
    return nil
}
// GetAssociatedIdentities gets the associatedIdentities property value. Identities in the authorization system.
// returns a AwsAssociatedIdentitiesable when successful
func (m *AwsAuthorizationSystem) GetAssociatedIdentities()(AwsAssociatedIdentitiesable) {
    val, err := m.GetBackingStore().Get("associatedIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsAssociatedIdentitiesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsAuthorizationSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystem.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAwsAuthorizationSystemTypeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AwsAuthorizationSystemTypeActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AwsAuthorizationSystemTypeActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["associatedIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsAssociatedIdentitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociatedIdentities(val.(AwsAssociatedIdentitiesable))
        }
        return nil
    }
    res["policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAwsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AwsPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AwsPolicyable)
                }
            }
            m.SetPolicies(res)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAwsAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AwsAuthorizationSystemResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AwsAuthorizationSystemResourceable)
                }
            }
            m.SetResources(res)
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
// GetPolicies gets the policies property value. Policies associated with the AWS authorization system type.
// returns a []AwsPolicyable when successful
func (m *AwsAuthorizationSystem) GetPolicies()([]AwsPolicyable) {
    val, err := m.GetBackingStore().Get("policies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AwsPolicyable)
    }
    return nil
}
// GetResources gets the resources property value. Resources associated with the authorization system type.
// returns a []AwsAuthorizationSystemResourceable when successful
func (m *AwsAuthorizationSystem) GetResources()([]AwsAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AwsAuthorizationSystemResourceable)
    }
    return nil
}
// GetServices gets the services property value. Services associated with the authorization system type.
// returns a []AuthorizationSystemTypeServiceable when successful
func (m *AwsAuthorizationSystem) GetServices()([]AuthorizationSystemTypeServiceable) {
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
func (m *AwsAuthorizationSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("policies", cast)
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
func (m *AwsAuthorizationSystem) SetActions(value []AwsAuthorizationSystemTypeActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAssociatedIdentities sets the associatedIdentities property value. Identities in the authorization system.
func (m *AwsAuthorizationSystem) SetAssociatedIdentities(value AwsAssociatedIdentitiesable)() {
    err := m.GetBackingStore().Set("associatedIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicies sets the policies property value. Policies associated with the AWS authorization system type.
func (m *AwsAuthorizationSystem) SetPolicies(value []AwsPolicyable)() {
    err := m.GetBackingStore().Set("policies", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. Resources associated with the authorization system type.
func (m *AwsAuthorizationSystem) SetResources(value []AwsAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// SetServices sets the services property value. Services associated with the authorization system type.
func (m *AwsAuthorizationSystem) SetServices(value []AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("services", value)
    if err != nil {
        panic(err)
    }
}
type AwsAuthorizationSystemable interface {
    AuthorizationSystemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]AwsAuthorizationSystemTypeActionable)
    GetAssociatedIdentities()(AwsAssociatedIdentitiesable)
    GetPolicies()([]AwsPolicyable)
    GetResources()([]AwsAuthorizationSystemResourceable)
    GetServices()([]AuthorizationSystemTypeServiceable)
    SetActions(value []AwsAuthorizationSystemTypeActionable)()
    SetAssociatedIdentities(value AwsAssociatedIdentitiesable)()
    SetPolicies(value []AwsPolicyable)()
    SetResources(value []AwsAuthorizationSystemResourceable)()
    SetServices(value []AuthorizationSystemTypeServiceable)()
}
