package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccess provides operations to manage the collection of accessReview entities.
type PrivilegedAccess struct {
    Entity
    // The display name of the provider managed by PIM.
    displayName *string
    // A collection of resources for the provider.
    resources []GovernanceResourceable
    // A collection of role assignment requests for the provider.
    roleAssignmentRequests []GovernanceRoleAssignmentRequestable
    // A collection of role assignments for the provider.
    roleAssignments []GovernanceRoleAssignmentable
    // A collection of role defintions for the provider.
    roleDefinitions []GovernanceRoleDefinitionable
    // A collection of role settings for the provider.
    roleSettings []GovernanceRoleSettingable
}
// NewPrivilegedAccess instantiates a new privilegedAccess and sets the default values.
func NewPrivilegedAccess()(*PrivilegedAccess) {
    m := &PrivilegedAccess{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccess(), nil
}
// GetDisplayName gets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["resources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceResourceFromDiscriminatorValue , m.SetResources)
    res["roleAssignmentRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue , m.SetRoleAssignmentRequests)
    res["roleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceRoleAssignmentFromDiscriminatorValue , m.SetRoleAssignments)
    res["roleDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceRoleDefinitionFromDiscriminatorValue , m.SetRoleDefinitions)
    res["roleSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceRoleSettingFromDiscriminatorValue , m.SetRoleSettings)
    return res
}
// GetResources gets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) GetResources()([]GovernanceResourceable) {
    return m.resources
}
// GetRoleAssignmentRequests gets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable) {
    return m.roleAssignmentRequests
}
// GetRoleAssignments gets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) GetRoleAssignments()([]GovernanceRoleAssignmentable) {
    return m.roleAssignments
}
// GetRoleDefinitions gets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) GetRoleDefinitions()([]GovernanceRoleDefinitionable) {
    return m.roleDefinitions
}
// GetRoleSettings gets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) GetRoleSettings()([]GovernanceRoleSettingable) {
    return m.roleSettings
}
// Serialize serializes information the current object
func (m *PrivilegedAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResources())
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentRequests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignmentRequests())
        err = writer.WriteCollectionOfObjectValues("roleAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleDefinitions())
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleSettings())
        err = writer.WriteCollectionOfObjectValues("roleSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetResources sets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) SetResources(value []GovernanceResourceable)() {
    m.resources = value
}
// SetRoleAssignmentRequests sets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)() {
    m.roleAssignmentRequests = value
}
// SetRoleAssignments sets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) SetRoleAssignments(value []GovernanceRoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRoleDefinitions sets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) SetRoleDefinitions(value []GovernanceRoleDefinitionable)() {
    m.roleDefinitions = value
}
// SetRoleSettings sets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) SetRoleSettings(value []GovernanceRoleSettingable)() {
    m.roleSettings = value
}
