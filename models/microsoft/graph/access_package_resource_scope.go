package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageResourceScope struct {
    Entity
    // Read-only. Nullable.
    accessPackageResource *AccessPackageResource;
    // The description of the scope.
    description *string;
    // The display name of the scope.
    displayName *string;
    // True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
    isRootScope *bool;
    // The unique identifier for the scope in the resource as defined in the origin system.
    originId *string;
    // The origin system for the scope.
    originSystem *string;
    // The origin system for the role, if different.
    roleOriginId *string;
    // A resource locator for the scope.
    url *string;
}
// Instantiates a new accessPackageResourceScope and sets the default values.
func NewAccessPackageResourceScope()(*AccessPackageResourceScope) {
    m := &AccessPackageResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageResource property value. Read-only. Nullable.
func (m *AccessPackageResourceScope) GetAccessPackageResource()(*AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResource
    }
}
// Gets the description property value. The description of the scope.
func (m *AccessPackageResourceScope) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the scope.
func (m *AccessPackageResourceScope) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
func (m *AccessPackageResourceScope) GetIsRootScope()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRootScope
    }
}
// Gets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
func (m *AccessPackageResourceScope) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
// Gets the originSystem property value. The origin system for the scope.
func (m *AccessPackageResourceScope) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
// Gets the roleOriginId property value. The origin system for the role, if different.
func (m *AccessPackageResourceScope) GetRoleOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleOriginId
    }
}
// Gets the url property value. A resource locator for the scope.
func (m *AccessPackageResourceScope) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// The deserialization information for the current model
func (m *AccessPackageResourceScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        m.SetAccessPackageResource(val.(*AccessPackageResource))
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isRootScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRootScope(val)
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
    res["roleOriginId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleOriginId(val)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *AccessPackageResourceScope) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageResourceScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isRootScope", m.GetIsRootScope())
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
        err = writer.WriteStringValue("roleOriginId", m.GetRoleOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accessPackageResource property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageResource property.
func (m *AccessPackageResourceScope) SetAccessPackageResource(value *AccessPackageResource)() {
    m.accessPackageResource = value
}
// Sets the description property value. The description of the scope.
// Parameters:
//  - value : Value to set for the description property.
func (m *AccessPackageResourceScope) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the scope.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessPackageResourceScope) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
// Parameters:
//  - value : Value to set for the isRootScope property.
func (m *AccessPackageResourceScope) SetIsRootScope(value *bool)() {
    m.isRootScope = value
}
// Sets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
// Parameters:
//  - value : Value to set for the originId property.
func (m *AccessPackageResourceScope) SetOriginId(value *string)() {
    m.originId = value
}
// Sets the originSystem property value. The origin system for the scope.
// Parameters:
//  - value : Value to set for the originSystem property.
func (m *AccessPackageResourceScope) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// Sets the roleOriginId property value. The origin system for the role, if different.
// Parameters:
//  - value : Value to set for the roleOriginId property.
func (m *AccessPackageResourceScope) SetRoleOriginId(value *string)() {
    m.roleOriginId = value
}
// Sets the url property value. A resource locator for the scope.
// Parameters:
//  - value : Value to set for the url property.
func (m *AccessPackageResourceScope) SetUrl(value *string)() {
    m.url = value
}
