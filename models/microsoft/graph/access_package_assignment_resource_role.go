package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentResourceRole 
type AccessPackageAssignmentResourceRole struct {
    Entity
    // The access package assignments resulting in this role assignment. Read-only. Nullable.
    accessPackageAssignments []AccessPackageAssignment;
    // Read-only. Nullable.
    accessPackageResourceRole *AccessPackageResourceRole;
    // Read-only. Nullable.
    accessPackageResourceScope *AccessPackageResourceScope;
    // Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
    accessPackageSubject *AccessPackageSubject;
    // A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
    originId *string;
    // The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
    originSystem *string;
    // The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
    status *string;
}
// NewAccessPackageAssignmentResourceRole instantiates a new accessPackageAssignmentResourceRole and sets the default values.
func NewAccessPackageAssignmentResourceRole()(*AccessPackageAssignmentResourceRole) {
    m := &AccessPackageAssignmentResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageAssignments gets the accessPackageAssignments property value. The access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageAssignments()([]AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignments
    }
}
// GetAccessPackageResourceRole gets the accessPackageResourceRole property value. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceRole()(*AccessPackageResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRole
    }
}
// GetAccessPackageResourceScope gets the accessPackageResourceScope property value. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceScope()(*AccessPackageResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceScope
    }
}
// GetAccessPackageSubject gets the accessPackageSubject property value. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageSubject()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageSubject
    }
}
// GetOriginId gets the originId property value. A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
// GetOriginSystem gets the originSystem property value. The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
// GetStatus gets the status property value. The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
func (m *AccessPackageAssignmentResourceRole) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentResourceRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAssignment))
            }
            m.SetAccessPackageAssignments(res)
        }
        return nil
    }
    res["accessPackageResourceRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRole() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceRole(val.(*AccessPackageResourceRole))
        }
        return nil
    }
    res["accessPackageResourceScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceScope() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceScope(val.(*AccessPackageResourceScope))
        }
        return nil
    }
    res["accessPackageSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageSubject(val.(*AccessPackageSubject))
        }
        return nil
    }
    res["originId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginId(val)
        }
        return nil
    }
    res["originSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginSystem(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentResourceRole) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentResourceRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignments()))
        for i, v := range m.GetAccessPackageAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceRole", m.GetAccessPackageResourceRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceScope", m.GetAccessPackageResourceScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageSubject", m.GetAccessPackageSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageAssignments sets the accessPackageAssignments property value. The access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageAssignments(value []AccessPackageAssignment)() {
    if m != nil {
        m.accessPackageAssignments = value
    }
}
// SetAccessPackageResourceRole sets the accessPackageResourceRole property value. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceRole(value *AccessPackageResourceRole)() {
    if m != nil {
        m.accessPackageResourceRole = value
    }
}
// SetAccessPackageResourceScope sets the accessPackageResourceScope property value. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceScope(value *AccessPackageResourceScope)() {
    if m != nil {
        m.accessPackageResourceScope = value
    }
}
// SetAccessPackageSubject sets the accessPackageSubject property value. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageSubject(value *AccessPackageSubject)() {
    if m != nil {
        m.accessPackageSubject = value
    }
}
// SetOriginId sets the originId property value. A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) SetOriginId(value *string)() {
    if m != nil {
        m.originId = value
    }
}
// SetOriginSystem sets the originSystem property value. The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) SetOriginSystem(value *string)() {
    if m != nil {
        m.originSystem = value
    }
}
// SetStatus sets the status property value. The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
func (m *AccessPackageAssignmentResourceRole) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
