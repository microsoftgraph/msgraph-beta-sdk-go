package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessPackageAssignmentResourceRole struct {
    Entity
    accessPackageAssignments []AccessPackageAssignment;
    accessPackageResourceRole *AccessPackageResourceRole;
    accessPackageResourceScope *AccessPackageResourceScope;
    accessPackageSubject *AccessPackageSubject;
    originId *string;
    originSystem *string;
    status *string;
}
func NewAccessPackageAssignmentResourceRole()(*AccessPackageAssignmentResourceRole) {
    m := &AccessPackageAssignmentResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageAssignments()([]AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignments
    }
}
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceRole()(*AccessPackageResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRole
    }
}
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceScope()(*AccessPackageResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceScope
    }
}
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageSubject()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageSubject
    }
}
func (m *AccessPackageAssignmentResourceRole) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
func (m *AccessPackageAssignmentResourceRole) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
func (m *AccessPackageAssignmentResourceRole) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *AccessPackageAssignmentResourceRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignment))
        }
        m.SetAccessPackageAssignments(res)
        return nil
    }
    res["accessPackageResourceRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRole() })
        if err != nil {
            return err
        }
        m.SetAccessPackageResourceRole(val.(*AccessPackageResourceRole))
        return nil
    }
    res["accessPackageResourceScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceScope() })
        if err != nil {
            return err
        }
        m.SetAccessPackageResourceScope(val.(*AccessPackageResourceScope))
        return nil
    }
    res["accessPackageSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        m.SetAccessPackageSubject(val.(*AccessPackageSubject))
        return nil
    }
    res["originId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginId(val)
        return nil
    }
    res["originSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginSystem(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentResourceRole) IsNil()(bool) {
    return m == nil
}
func (m *AccessPackageAssignmentResourceRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageAssignments(value []AccessPackageAssignment)() {
    m.accessPackageAssignments = value
}
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceRole(value *AccessPackageResourceRole)() {
    m.accessPackageResourceRole = value
}
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceScope(value *AccessPackageResourceScope)() {
    m.accessPackageResourceScope = value
}
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageSubject(value *AccessPackageSubject)() {
    m.accessPackageSubject = value
}
func (m *AccessPackageAssignmentResourceRole) SetOriginId(value *string)() {
    m.originId = value
}
func (m *AccessPackageAssignmentResourceRole) SetOriginSystem(value *string)() {
    m.originSystem = value
}
func (m *AccessPackageAssignmentResourceRole) SetStatus(value *string)() {
    m.status = value
}
