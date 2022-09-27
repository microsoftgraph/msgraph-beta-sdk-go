package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationEventsPolicy 
type AuthenticationEventsPolicy struct {
    Entity
    // A list of applicable actions to be taken on sign-up.
    onSignupStart []AuthenticationListenerable
}
// NewAuthenticationEventsPolicy instantiates a new AuthenticationEventsPolicy and sets the default values.
func NewAuthenticationEventsPolicy()(*AuthenticationEventsPolicy) {
    m := &AuthenticationEventsPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationEventsPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationEventsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationEventsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationEventsPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationEventsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["onSignupStart"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationListenerFromDiscriminatorValue , m.SetOnSignupStart)
    return res
}
// GetOnSignupStart gets the onSignupStart property value. A list of applicable actions to be taken on sign-up.
func (m *AuthenticationEventsPolicy) GetOnSignupStart()([]AuthenticationListenerable) {
    return m.onSignupStart
}
// Serialize serializes information the current object
func (m *AuthenticationEventsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOnSignupStart() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOnSignupStart())
        err = writer.WriteCollectionOfObjectValues("onSignupStart", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnSignupStart sets the onSignupStart property value. A list of applicable actions to be taken on sign-up.
func (m *AuthenticationEventsPolicy) SetOnSignupStart(value []AuthenticationListenerable)() {
    m.onSignupStart = value
}
