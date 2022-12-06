package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessRoot 
type ConditionalAccessRoot struct {
    Entity
    // Read-only. Nullable. Returns a collection of the specified authentication context class references.
    authenticationContextClassReferences []AuthenticationContextClassReferenceable
    // Defines the authentication strength policies, valid authentication method combinations, and authentication method mode details that can be required by a conditional access policy .
    authenticationStrengths AuthenticationStrengthRootable
    // Read-only. Nullable. Returns a collection of the specified named locations.
    namedLocations []NamedLocationable
    // Read-only. Nullable. Returns a collection of the specified Conditional Access policies.
    policies []ConditionalAccessPolicyable
    // Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
    templates []ConditionalAccessTemplateable
}
// NewConditionalAccessRoot instantiates a new ConditionalAccessRoot and sets the default values.
func NewConditionalAccessRoot()(*ConditionalAccessRoot) {
    m := &ConditionalAccessRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConditionalAccessRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessRoot(), nil
}
// GetAuthenticationContextClassReferences gets the authenticationContextClassReferences property value. Read-only. Nullable. Returns a collection of the specified authentication context class references.
func (m *ConditionalAccessRoot) GetAuthenticationContextClassReferences()([]AuthenticationContextClassReferenceable) {
    return m.authenticationContextClassReferences
}
// GetAuthenticationStrengths gets the authenticationStrengths property value. Defines the authentication strength policies, valid authentication method combinations, and authentication method mode details that can be required by a conditional access policy .
func (m *ConditionalAccessRoot) GetAuthenticationStrengths()(AuthenticationStrengthRootable) {
    return m.authenticationStrengths
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationContextClassReferences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationContextClassReferenceFromDiscriminatorValue , m.SetAuthenticationContextClassReferences)
    res["authenticationStrengths"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationStrengthRootFromDiscriminatorValue , m.SetAuthenticationStrengths)
    res["namedLocations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNamedLocationFromDiscriminatorValue , m.SetNamedLocations)
    res["policies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue , m.SetPolicies)
    res["templates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConditionalAccessTemplateFromDiscriminatorValue , m.SetTemplates)
    return res
}
// GetNamedLocations gets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
func (m *ConditionalAccessRoot) GetNamedLocations()([]NamedLocationable) {
    return m.namedLocations
}
// GetPolicies gets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access policies.
func (m *ConditionalAccessRoot) GetPolicies()([]ConditionalAccessPolicyable) {
    return m.policies
}
// GetTemplates gets the templates property value. Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
func (m *ConditionalAccessRoot) GetTemplates()([]ConditionalAccessTemplateable) {
    return m.templates
}
// Serialize serializes information the current object
func (m *ConditionalAccessRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationContextClassReferences() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuthenticationContextClassReferences())
        err = writer.WriteCollectionOfObjectValues("authenticationContextClassReferences", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationStrengths", m.GetAuthenticationStrengths())
        if err != nil {
            return err
        }
    }
    if m.GetNamedLocations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNamedLocations())
        err = writer.WriteCollectionOfObjectValues("namedLocations", cast)
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
    if m.GetTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTemplates())
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationContextClassReferences sets the authenticationContextClassReferences property value. Read-only. Nullable. Returns a collection of the specified authentication context class references.
func (m *ConditionalAccessRoot) SetAuthenticationContextClassReferences(value []AuthenticationContextClassReferenceable)() {
    m.authenticationContextClassReferences = value
}
// SetAuthenticationStrengths sets the authenticationStrengths property value. Defines the authentication strength policies, valid authentication method combinations, and authentication method mode details that can be required by a conditional access policy .
func (m *ConditionalAccessRoot) SetAuthenticationStrengths(value AuthenticationStrengthRootable)() {
    m.authenticationStrengths = value
}
// SetNamedLocations sets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
func (m *ConditionalAccessRoot) SetNamedLocations(value []NamedLocationable)() {
    m.namedLocations = value
}
// SetPolicies sets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access policies.
func (m *ConditionalAccessRoot) SetPolicies(value []ConditionalAccessPolicyable)() {
    m.policies = value
}
// SetTemplates sets the templates property value. Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
func (m *ConditionalAccessRoot) SetTemplates(value []ConditionalAccessTemplateable)() {
    m.templates = value
}
