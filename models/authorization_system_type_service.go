package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthorizationSystemTypeService struct {
    Entity
}
// NewAuthorizationSystemTypeService instantiates a new AuthorizationSystemTypeService and sets the default values.
func NewAuthorizationSystemTypeService()(*AuthorizationSystemTypeService) {
    m := &AuthorizationSystemTypeService{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthorizationSystemTypeServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizationSystemTypeServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorizationSystemTypeService(), nil
}
// GetActions gets the actions property value. List of actions for the service in an authorization system that is onboarded to Permissions Management.
// returns a []AuthorizationSystemTypeActionable when successful
func (m *AuthorizationSystemTypeService) GetActions()([]AuthorizationSystemTypeActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationSystemTypeActionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthorizationSystemTypeService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationSystemTypeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationSystemTypeActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationSystemTypeActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthorizationSystemTypeService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
    return nil
}
// SetActions sets the actions property value. List of actions for the service in an authorization system that is onboarded to Permissions Management.
func (m *AuthorizationSystemTypeService) SetActions(value []AuthorizationSystemTypeActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
type AuthorizationSystemTypeServiceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]AuthorizationSystemTypeActionable)
    SetActions(value []AuthorizationSystemTypeActionable)()
}
