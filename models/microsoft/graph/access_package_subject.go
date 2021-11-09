package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageSubject struct {
    Entity
    // 
    altSecId *string;
    // The connected organization of the subject. Read-only. Nullable.
    connectedOrganization *ConnectedOrganization;
    // The identifier of the connected organization of the subject.
    connectedOrganizationId *string;
    // The display name of the subject.
    displayName *string;
    // The email address of the subject.
    email *string;
    // The object identifier of the subject. null if the subject is not yet a user in the tenant.
    objectId *string;
    // 
    onPremisesSecurityIdentifier *string;
    // The principal name, if known, of the subject.
    principalName *string;
    // The resource type of the subject.
    type_escaped *string;
}
// Instantiates a new accessPackageSubject and sets the default values.
func NewAccessPackageSubject()(*AccessPackageSubject) {
    m := &AccessPackageSubject{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the altSecId property value. 
func (m *AccessPackageSubject) GetAltSecId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.altSecId
    }
}
// Gets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) GetConnectedOrganization()(*ConnectedOrganization) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganization
    }
}
// Gets the connectedOrganizationId property value. The identifier of the connected organization of the subject.
func (m *AccessPackageSubject) GetConnectedOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganizationId
    }
}
// Gets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. The email address of the subject.
func (m *AccessPackageSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
// Gets the onPremisesSecurityIdentifier property value. 
func (m *AccessPackageSubject) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// Gets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
// Gets the type_escaped property value. The resource type of the subject.
func (m *AccessPackageSubject) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *AccessPackageSubject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["altSecId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltSecId(val)
        }
        return nil
    }
    res["connectedOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectedOrganization() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedOrganization(val.(*ConnectedOrganization))
        }
        return nil
    }
    res["connectedOrganizationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedOrganizationId(val)
        }
        return nil
    }
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
    res["objectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
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
func (m *AccessPackageSubject) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageSubject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("altSecId", m.GetAltSecId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectedOrganization", m.GetConnectedOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectedOrganizationId", m.GetConnectedOrganizationId())
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
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
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
// Sets the altSecId property value. 
// Parameters:
//  - value : Value to set for the altSecId property.
func (m *AccessPackageSubject) SetAltSecId(value *string)() {
    m.altSecId = value
}
// Sets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the connectedOrganization property.
func (m *AccessPackageSubject) SetConnectedOrganization(value *ConnectedOrganization)() {
    m.connectedOrganization = value
}
// Sets the connectedOrganizationId property value. The identifier of the connected organization of the subject.
// Parameters:
//  - value : Value to set for the connectedOrganizationId property.
func (m *AccessPackageSubject) SetConnectedOrganizationId(value *string)() {
    m.connectedOrganizationId = value
}
// Sets the displayName property value. The display name of the subject.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessPackageSubject) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. The email address of the subject.
// Parameters:
//  - value : Value to set for the email property.
func (m *AccessPackageSubject) SetEmail(value *string)() {
    m.email = value
}
// Sets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
// Parameters:
//  - value : Value to set for the objectId property.
func (m *AccessPackageSubject) SetObjectId(value *string)() {
    m.objectId = value
}
// Sets the onPremisesSecurityIdentifier property value. 
// Parameters:
//  - value : Value to set for the onPremisesSecurityIdentifier property.
func (m *AccessPackageSubject) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// Sets the principalName property value. The principal name, if known, of the subject.
// Parameters:
//  - value : Value to set for the principalName property.
func (m *AccessPackageSubject) SetPrincipalName(value *string)() {
    m.principalName = value
}
// Sets the type_escaped property value. The resource type of the subject.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *AccessPackageSubject) SetType_escaped(value *string)() {
    m.type_escaped = value
}
