package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageSubject 
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
    // A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
    onPremisesSecurityIdentifier *string;
    // The principal name, if known, of the subject.
    principalName *string;
    // The resource type of the subject.
    type_escaped *string;
}
// NewAccessPackageSubject instantiates a new accessPackageSubject and sets the default values.
func NewAccessPackageSubject()(*AccessPackageSubject) {
    m := &AccessPackageSubject{
        Entity: *NewEntity(),
    }
    return m
}
// GetAltSecId gets the altSecId property value. 
func (m *AccessPackageSubject) GetAltSecId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.altSecId
    }
}
// GetConnectedOrganization gets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) GetConnectedOrganization()(*ConnectedOrganization) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganization
    }
}
// GetConnectedOrganizationId gets the connectedOrganizationId property value. The identifier of the connected organization of the subject.
func (m *AccessPackageSubject) GetConnectedOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganizationId
    }
}
// GetDisplayName gets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. The email address of the subject.
func (m *AccessPackageSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetObjectId gets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
func (m *AccessPackageSubject) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// GetPrincipalName gets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
// GetType gets the type property value. The resource type of the subject.
func (m *AccessPackageSubject) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *AccessPackageSubject) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAltSecId sets the altSecId property value. 
func (m *AccessPackageSubject) SetAltSecId(value *string)() {
    if m != nil {
        m.altSecId = value
    }
}
// SetConnectedOrganization sets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) SetConnectedOrganization(value *ConnectedOrganization)() {
    if m != nil {
        m.connectedOrganization = value
    }
}
// SetConnectedOrganizationId sets the connectedOrganizationId property value. The identifier of the connected organization of the subject.
func (m *AccessPackageSubject) SetConnectedOrganizationId(value *string)() {
    if m != nil {
        m.connectedOrganizationId = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. The email address of the subject.
func (m *AccessPackageSubject) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetObjectId sets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) SetObjectId(value *string)() {
    if m != nil {
        m.objectId = value
    }
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
func (m *AccessPackageSubject) SetOnPremisesSecurityIdentifier(value *string)() {
    if m != nil {
        m.onPremisesSecurityIdentifier = value
    }
}
// SetPrincipalName sets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) SetPrincipalName(value *string)() {
    if m != nil {
        m.principalName = value
    }
}
// SetType sets the type property value. The resource type of the subject.
func (m *AccessPackageSubject) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
