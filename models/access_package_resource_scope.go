package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceScope provides operations to manage the collection of activityStatistics entities.
type AccessPackageResourceScope struct {
    Entity
    // The accessPackageResource property
    accessPackageResource AccessPackageResourceable
    // The description of the scope.
    description *string
    // The display name of the scope.
    displayName *string
    // True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
    isRootScope *bool
    // The unique identifier for the scope in the resource as defined in the origin system.
    originId *string
    // The origin system for the scope.
    originSystem *string
    // The origin system for the role, if different.
    roleOriginId *string
    // A resource locator for the scope.
    url *string
}
// NewAccessPackageResourceScope instantiates a new accessPackageResourceScope and sets the default values.
func NewAccessPackageResourceScope()(*AccessPackageResourceScope) {
    m := &AccessPackageResourceScope{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageResourceScope";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageResourceScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceScope(), nil
}
// GetAccessPackageResource gets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceScope) GetAccessPackageResource()(AccessPackageResourceable) {
    return m.accessPackageResource
}
// GetDescription gets the description property value. The description of the scope.
func (m *AccessPackageResourceScope) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the scope.
func (m *AccessPackageResourceScope) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue , m.SetAccessPackageResource)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isRootScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRootScope)
    res["originId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginId)
    res["originSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginSystem)
    res["roleOriginId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRoleOriginId)
    res["url"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUrl)
    return res
}
// GetIsRootScope gets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
func (m *AccessPackageResourceScope) GetIsRootScope()(*bool) {
    return m.isRootScope
}
// GetOriginId gets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
func (m *AccessPackageResourceScope) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The origin system for the scope.
func (m *AccessPackageResourceScope) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetRoleOriginId gets the roleOriginId property value. The origin system for the role, if different.
func (m *AccessPackageResourceScope) GetRoleOriginId()(*string) {
    return m.roleOriginId
}
// GetUrl gets the url property value. A resource locator for the scope.
func (m *AccessPackageResourceScope) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *AccessPackageResourceScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAccessPackageResource sets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceScope) SetAccessPackageResource(value AccessPackageResourceable)() {
    m.accessPackageResource = value
}
// SetDescription sets the description property value. The description of the scope.
func (m *AccessPackageResourceScope) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the scope.
func (m *AccessPackageResourceScope) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsRootScope sets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
func (m *AccessPackageResourceScope) SetIsRootScope(value *bool)() {
    m.isRootScope = value
}
// SetOriginId sets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
func (m *AccessPackageResourceScope) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The origin system for the scope.
func (m *AccessPackageResourceScope) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetRoleOriginId sets the roleOriginId property value. The origin system for the role, if different.
func (m *AccessPackageResourceScope) SetRoleOriginId(value *string)() {
    m.roleOriginId = value
}
// SetUrl sets the url property value. A resource locator for the scope.
func (m *AccessPackageResourceScope) SetUrl(value *string)() {
    m.url = value
}
