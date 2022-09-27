package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleScopeTag role Scope Tag
type RoleScopeTag struct {
    Entity
    // The list of assignments for this Role Scope Tag.
    assignments []RoleScopeTagAutoAssignmentable
    // Description of the Role Scope Tag.
    description *string
    // The display or friendly name of the Role Scope Tag.
    displayName *string
    // Description of the Role Scope Tag. This property is read-only.
    isBuiltIn *bool
}
// NewRoleScopeTag instantiates a new roleScopeTag and sets the default values.
func NewRoleScopeTag()(*RoleScopeTag) {
    m := &RoleScopeTag{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.roleScopeTag";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRoleScopeTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleScopeTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTag(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) GetAssignments()([]RoleScopeTagAutoAssignmentable) {
    return m.assignments
}
// GetDescription gets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRoleScopeTagAutoAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isBuiltIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBuiltIn)
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// Serialize serializes information the current object
func (m *RoleScopeTag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) SetAssignments(value []RoleScopeTagAutoAssignmentable)() {
    m.assignments = value
}
// SetDescription sets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsBuiltIn sets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
