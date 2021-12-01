package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceRole 
type AccessPackageResourceRole struct {
    Entity
    // Read-only. Nullable.
    accessPackageResource *AccessPackageResource;
    // A description for the resource role.
    description *string;
    // The display name of the resource role such as the role defined by the application.
    displayName *string;
    // The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId will be the sequence number of the role in the site.
    originId *string;
    // The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
    originSystem *string;
}
// NewAccessPackageResourceRole instantiates a new accessPackageResourceRole and sets the default values.
func NewAccessPackageResourceRole()(*AccessPackageResourceRole) {
    m := &AccessPackageResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageResource gets the accessPackageResource property value. Read-only. Nullable.
func (m *AccessPackageResourceRole) GetAccessPackageResource()(*AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResource
    }
}
// GetDescription gets the description property value. A description for the resource role.
func (m *AccessPackageResourceRole) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the resource role such as the role defined by the application.
func (m *AccessPackageResourceRole) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetOriginId gets the originId property value. The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId will be the sequence number of the role in the site.
func (m *AccessPackageResourceRole) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResourceRole) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResource(val.(*AccessPackageResource))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    return res
}
func (m *AccessPackageResourceRole) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackageResource", m.GetAccessPackageResource())
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
    return nil
}
// SetAccessPackageResource sets the accessPackageResource property value. Read-only. Nullable.
func (m *AccessPackageResourceRole) SetAccessPackageResource(value *AccessPackageResource)() {
    if m != nil {
        m.accessPackageResource = value
    }
}
// SetDescription sets the description property value. A description for the resource role.
func (m *AccessPackageResourceRole) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the resource role such as the role defined by the application.
func (m *AccessPackageResourceRole) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetOriginId sets the originId property value. The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId will be the sequence number of the role in the site.
func (m *AccessPackageResourceRole) SetOriginId(value *string)() {
    if m != nil {
        m.originId = value
    }
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResourceRole) SetOriginSystem(value *string)() {
    if m != nil {
        m.originSystem = value
    }
}
