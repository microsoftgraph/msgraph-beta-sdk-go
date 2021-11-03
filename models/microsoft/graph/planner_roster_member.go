package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerRosterMember struct {
    Entity
    // Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents of the plannerRoster.
    roles []string;
    // Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a plannerRoster.
    tenantId *string;
    // Identifier of the user.
    userId *string;
}
// Instantiates a new plannerRosterMember and sets the default values.
func NewPlannerRosterMember()(*PlannerRosterMember) {
    m := &PlannerRosterMember{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the roles property value. Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents of the plannerRoster.
func (m *PlannerRosterMember) GetRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roles
    }
}
// Gets the tenantId property value. Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a plannerRoster.
func (m *PlannerRosterMember) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the userId property value. Identifier of the user.
func (m *PlannerRosterMember) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *PlannerRosterMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["roles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoles(res)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *PlannerRosterMember) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PlannerRosterMember) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
// Sets the roles property value. Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents of the plannerRoster.
// Parameters:
//  - value : Value to set for the roles property.
func (m *PlannerRosterMember) SetRoles(value []string)() {
    m.roles = value
}
// Sets the tenantId property value. Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a plannerRoster.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *PlannerRosterMember) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the userId property value. Identifier of the user.
// Parameters:
//  - value : Value to set for the userId property.
func (m *PlannerRosterMember) SetUserId(value *string)() {
    m.userId = value
}
