package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodsRoot 
type AuthenticationMethodsRoot struct {
    Entity
    // Represents the state of a user's authentication methods, including which methods are registered and which features the user is registered and capable of (such as multi-factor authentication, self-service password reset, and passwordless authentication).
    userRegistrationDetails []UserRegistrationDetailsable
}
// NewAuthenticationMethodsRoot instantiates a new AuthenticationMethodsRoot and sets the default values.
func NewAuthenticationMethodsRoot()(*AuthenticationMethodsRoot) {
    m := &AuthenticationMethodsRoot{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationMethodsRoot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationMethodsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodsRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userRegistrationDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserRegistrationDetailsFromDiscriminatorValue , m.SetUserRegistrationDetails)
    return res
}
// GetUserRegistrationDetails gets the userRegistrationDetails property value. Represents the state of a user's authentication methods, including which methods are registered and which features the user is registered and capable of (such as multi-factor authentication, self-service password reset, and passwordless authentication).
func (m *AuthenticationMethodsRoot) GetUserRegistrationDetails()([]UserRegistrationDetailsable) {
    return m.userRegistrationDetails
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserRegistrationDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserRegistrationDetails())
        err = writer.WriteCollectionOfObjectValues("userRegistrationDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserRegistrationDetails sets the userRegistrationDetails property value. Represents the state of a user's authentication methods, including which methods are registered and which features the user is registered and capable of (such as multi-factor authentication, self-service password reset, and passwordless authentication).
func (m *AuthenticationMethodsRoot) SetUserRegistrationDetails(value []UserRegistrationDetailsable)() {
    m.userRegistrationDetails = value
}
