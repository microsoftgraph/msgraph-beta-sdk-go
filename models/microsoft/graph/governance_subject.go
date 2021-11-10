package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new governanceSubject and sets the default values.
func NewGovernanceSubject()(*GovernanceSubject) {
    m := &GovernanceSubject{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name of the subject.
func (m *GovernanceSubject) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. The email address of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the principalName property value. The principal name of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
// Gets the type_escaped property value. The type of the subject. The value can be User, Group, and ServicePrincipal.
func (m *GovernanceSubject) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceSubject) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name of the subject.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GovernanceSubject) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. The email address of the user subject. If the subject is in other types, it is empty.
// Parameters:
//  - value : Value to set for the email property.
func (m *GovernanceSubject) SetEmail(value *string)() {
    m.email = value
}
// Sets the principalName property value. The principal name of the user subject. If the subject is in other types, it is empty.
// Parameters:
//  - value : Value to set for the principalName property.
func (m *GovernanceSubject) SetPrincipalName(value *string)() {
    m.principalName = value
}
// Sets the type_escaped property value. The type of the subject. The value can be User, Group, and ServicePrincipal.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *GovernanceSubject) SetType_escaped(value *string)() {
    m.type_escaped = value
}
