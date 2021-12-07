package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceSubject 
type GovernanceSubject struct {
    Entity
    // The display name of the subject.
    displayName *string;
    // The email address of the user subject. If the subject is in other types, it is empty.
    email *string;
    // The principal name of the user subject. If the subject is in other types, it is empty.
    principalName *string;
    // The type of the subject. The value can be User, Group, and ServicePrincipal.
    type_escaped *string;
}
// NewGovernanceSubject instantiates a new governanceSubject and sets the default values.
func NewGovernanceSubject()(*GovernanceSubject) {
    m := &GovernanceSubject{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The display name of the subject.
func (m *GovernanceSubject) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. The email address of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetPrincipalName gets the principalName property value. The principal name of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
// GetType gets the type property value. The type of the subject. The value can be User, Group, and ServicePrincipal.
func (m *GovernanceSubject) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceSubject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["principalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalName(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceSubject) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceSubject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the subject.
func (m *GovernanceSubject) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. The email address of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetPrincipalName sets the principalName property value. The principal name of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) SetPrincipalName(value *string)() {
    if m != nil {
        m.principalName = value
    }
}
// SetType sets the type property value. The type of the subject. The value can be User, Group, and ServicePrincipal.
func (m *GovernanceSubject) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
