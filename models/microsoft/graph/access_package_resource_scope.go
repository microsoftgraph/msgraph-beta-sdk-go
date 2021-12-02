package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceScope 
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
// NewAccessPackageResourceScope instantiates a new accessPackageResourceScope and sets the default values.
func NewAccessPackageResourceScope()(*AccessPackageResourceScope) {
    m := &AccessPackageResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageResource gets the accessPackageResource property value. Read-only. Nullable.
func (m *AccessPackageResourceScope) GetAccessPackageResource()(*AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResource
    }
}
// GetDescription gets the description property value. The description of the scope.
func (m *AccessPackageResourceScope) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the scope.
func (m *AccessPackageResourceScope) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsRootScope gets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
func (m *AccessPackageResourceScope) GetIsRootScope()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRootScope
    }
}
// GetOriginId gets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
func (m *AccessPackageResourceScope) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
// GetOriginSystem gets the originSystem property value. The origin system for the scope.
func (m *AccessPackageResourceScope) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
// GetRoleOriginId gets the roleOriginId property value. The origin system for the role, if different.
func (m *AccessPackageResourceScope) GetRoleOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleOriginId
    }
}
// GetUrl gets the url property value. A resource locator for the scope.
func (m *AccessPackageResourceScope) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["isRootScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRootScope(val)
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
    res["roleOriginId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleOriginId(val)
        }
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageResourceScope) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAccessPackageResource sets the accessPackageResource property value. Read-only. Nullable.
func (m *AccessPackageResourceScope) SetAccessPackageResource(value *AccessPackageResource)() {
    if m != nil {
        m.accessPackageResource = value
    }
}
// SetDescription sets the description property value. The description of the scope.
func (m *AccessPackageResourceScope) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the scope.
func (m *AccessPackageResourceScope) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsRootScope sets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
func (m *AccessPackageResourceScope) SetIsRootScope(value *bool)() {
    if m != nil {
        m.isRootScope = value
    }
}
// SetOriginId sets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
func (m *AccessPackageResourceScope) SetOriginId(value *string)() {
    if m != nil {
        m.originId = value
    }
}
// SetOriginSystem sets the originSystem property value. The origin system for the scope.
func (m *AccessPackageResourceScope) SetOriginSystem(value *string)() {
    if m != nil {
        m.originSystem = value
    }
}
// SetRoleOriginId sets the roleOriginId property value. The origin system for the role, if different.
func (m *AccessPackageResourceScope) SetRoleOriginId(value *string)() {
    if m != nil {
        m.roleOriginId = value
    }
}
// SetUrl sets the url property value. A resource locator for the scope.
func (m *AccessPackageResourceScope) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
