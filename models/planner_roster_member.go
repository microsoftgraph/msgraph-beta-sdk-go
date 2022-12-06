package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerRosterMember provides operations to manage the collection of accessReviewDecision entities.
type PlannerRosterMember struct {
    Entity
    // Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents of the plannerRoster.
    roles []string
    // Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a plannerRoster.
    tenantId *string
    // Identifier of the user.
    userId *string
}
// NewPlannerRosterMember instantiates a new plannerRosterMember and sets the default values.
func NewPlannerRosterMember()(*PlannerRosterMember) {
    m := &PlannerRosterMember{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerRosterMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerRosterMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerRosterMember(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerRosterMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["roles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoles)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetRoles gets the roles property value. Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents of the plannerRoster.
func (m *PlannerRosterMember) GetRoles()([]string) {
    return m.roles
}
// GetTenantId gets the tenantId property value. Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a plannerRoster.
func (m *PlannerRosterMember) GetTenantId()(*string) {
    return m.tenantId
}
// GetUserId gets the userId property value. Identifier of the user.
func (m *PlannerRosterMember) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *PlannerRosterMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRoles() != nil {
        err = writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoles sets the roles property value. Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents of the plannerRoster.
func (m *PlannerRosterMember) SetRoles(value []string)() {
    m.roles = value
}
// SetTenantId sets the tenantId property value. Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a plannerRoster.
func (m *PlannerRosterMember) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUserId sets the userId property value. Identifier of the user.
func (m *PlannerRosterMember) SetUserId(value *string)() {
    m.userId = value
}
