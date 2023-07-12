package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationEventsPolicy 
type AuthenticationEventsPolicy struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewAuthenticationEventsPolicy instantiates a new authenticationEventsPolicy and sets the default values.
func NewAuthenticationEventsPolicy()(*AuthenticationEventsPolicy) {
    m := &AuthenticationEventsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationEventsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationEventsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationEventsPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationEventsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["onSignupStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationListenerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationListenerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationListenerable)
                }
            }
            m.SetOnSignupStart(res)
        }
        return nil
    }
    return res
}
// GetOnSignupStart gets the onSignupStart property value. A list of applicable actions to be taken on sign-up.
func (m *AuthenticationEventsPolicy) GetOnSignupStart()([]AuthenticationListenerable) {
    val, err := m.GetBackingStore().Get("onSignupStart")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationListenerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationEventsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOnSignupStart() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnSignupStart()))
        for i, v := range m.GetOnSignupStart() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("onSignupStart", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnSignupStart sets the onSignupStart property value. A list of applicable actions to be taken on sign-up.
func (m *AuthenticationEventsPolicy) SetOnSignupStart(value []AuthenticationListenerable)() {
    err := m.GetBackingStore().Set("onSignupStart", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationEventsPolicyable 
type AuthenticationEventsPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOnSignupStart()([]AuthenticationListenerable)
    SetOnSignupStart(value []AuthenticationListenerable)()
}
