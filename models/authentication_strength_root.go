package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationStrengthRoot 
type AuthenticationStrengthRoot struct {
    Entity
    // A collection of all valid authentication method combinations in the system.
    authenticationCombinations []AuthenticationMethodModes
    // Names and descriptions of all valid authentication method modes in the system.
    authenticationMethodModes []AuthenticationMethodModeDetailable
    // A collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
    policies []AuthenticationStrengthPolicyable
}
// NewAuthenticationStrengthRoot instantiates a new AuthenticationStrengthRoot and sets the default values.
func NewAuthenticationStrengthRoot()(*AuthenticationStrengthRoot) {
    m := &AuthenticationStrengthRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationStrengthRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationStrengthRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationStrengthRoot(), nil
}
// GetAuthenticationCombinations gets the authenticationCombinations property value. A collection of all valid authentication method combinations in the system.
func (m *AuthenticationStrengthRoot) GetAuthenticationCombinations()([]AuthenticationMethodModes) {
    return m.authenticationCombinations
}
// GetAuthenticationMethodModes gets the authenticationMethodModes property value. Names and descriptions of all valid authentication method modes in the system.
func (m *AuthenticationStrengthRoot) GetAuthenticationMethodModes()([]AuthenticationMethodModeDetailable) {
    return m.authenticationMethodModes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationStrengthRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationCombinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseAuthenticationMethodModes , m.SetAuthenticationCombinations)
    res["authenticationMethodModes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationMethodModeDetailFromDiscriminatorValue , m.SetAuthenticationMethodModes)
    res["policies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationStrengthPolicyFromDiscriminatorValue , m.SetPolicies)
    return res
}
// GetPolicies gets the policies property value. A collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *AuthenticationStrengthRoot) GetPolicies()([]AuthenticationStrengthPolicyable) {
    return m.policies
}
// Serialize serializes information the current object
func (m *AuthenticationStrengthRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationCombinations() != nil {
        err = writer.WriteCollectionOfStringValues("authenticationCombinations", SerializeAuthenticationMethodModes(m.GetAuthenticationCombinations()))
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationMethodModes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuthenticationMethodModes())
        err = writer.WriteCollectionOfObjectValues("authenticationMethodModes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPolicies())
        err = writer.WriteCollectionOfObjectValues("policies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationCombinations sets the authenticationCombinations property value. A collection of all valid authentication method combinations in the system.
func (m *AuthenticationStrengthRoot) SetAuthenticationCombinations(value []AuthenticationMethodModes)() {
    m.authenticationCombinations = value
}
// SetAuthenticationMethodModes sets the authenticationMethodModes property value. Names and descriptions of all valid authentication method modes in the system.
func (m *AuthenticationStrengthRoot) SetAuthenticationMethodModes(value []AuthenticationMethodModeDetailable)() {
    m.authenticationMethodModes = value
}
// SetPolicies sets the policies property value. A collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *AuthenticationStrengthRoot) SetPolicies(value []AuthenticationStrengthPolicyable)() {
    m.policies = value
}
