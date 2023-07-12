package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRbacResourceAction 
type UnifiedRbacResourceAction struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewUnifiedRbacResourceAction instantiates a new unifiedRbacResourceAction and sets the default values.
func NewUnifiedRbacResourceAction()(*UnifiedRbacResourceAction) {
    m := &UnifiedRbacResourceAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRbacResourceActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRbacResourceActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacResourceAction(), nil
}
// GetActionVerb gets the actionVerb property value. HTTP method for the action, such as DELETE, GET, PATCH, POST, PUT, or null. Supports $filter (eq) but not for null values.
func (m *UnifiedRbacResourceAction) GetActionVerb()(*string) {
    val, err := m.GetBackingStore().Get("actionVerb")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationContext gets the authenticationContext property value. The authenticationContext property
func (m *UnifiedRbacResourceAction) GetAuthenticationContext()(AuthenticationContextClassReferenceable) {
    val, err := m.GetBackingStore().Get("authenticationContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationContextClassReferenceable)
    }
    return nil
}
// GetAuthenticationContextId gets the authenticationContextId property value. The authenticationContextId property
func (m *UnifiedRbacResourceAction) GetAuthenticationContextId()(*string) {
    val, err := m.GetBackingStore().Get("authenticationContextId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description for the action. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionVerb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionVerb(val)
        }
        return nil
    }
    res["authenticationContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationContextClassReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationContext(val.(AuthenticationContextClassReferenceable))
        }
        return nil
    }
    res["authenticationContextId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationContextId(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["isAuthenticationContextSettable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAuthenticationContextSettable(val)
        }
        return nil
    }
    res["isPrivileged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivileged(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resourceScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRbacResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceScope(val.(UnifiedRbacResourceScopeable))
        }
        return nil
    }
    res["resourceScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceScopeId(val)
        }
        return nil
    }
    return res
}
// GetIsAuthenticationContextSettable gets the isAuthenticationContextSettable property value. The isAuthenticationContextSettable property
func (m *UnifiedRbacResourceAction) GetIsAuthenticationContextSettable()(*bool) {
    val, err := m.GetBackingStore().Get("isAuthenticationContextSettable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPrivileged gets the isPrivileged property value. The isPrivileged property
func (m *UnifiedRbacResourceAction) GetIsPrivileged()(*bool) {
    val, err := m.GetBackingStore().Get("isPrivileged")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. Name for the action within the resource namespace, such as microsoft.insights/programs/update. Can include slash character (/). Case insensitive. Required. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceScope gets the resourceScope property value. The resourceScope property
func (m *UnifiedRbacResourceAction) GetResourceScope()(UnifiedRbacResourceScopeable) {
    val, err := m.GetBackingStore().Get("resourceScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnifiedRbacResourceScopeable)
    }
    return nil
}
// GetResourceScopeId gets the resourceScopeId property value. Not implemented.
func (m *UnifiedRbacResourceAction) GetResourceScopeId()(*string) {
    val, err := m.GetBackingStore().Get("resourceScopeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRbacResourceAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionVerb", m.GetActionVerb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationContext", m.GetAuthenticationContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authenticationContextId", m.GetAuthenticationContextId())
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
        err = writer.WriteBoolValue("isAuthenticationContextSettable", m.GetIsAuthenticationContextSettable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPrivileged", m.GetIsPrivileged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourceScope", m.GetResourceScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceScopeId", m.GetResourceScopeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionVerb sets the actionVerb property value. HTTP method for the action, such as DELETE, GET, PATCH, POST, PUT, or null. Supports $filter (eq) but not for null values.
func (m *UnifiedRbacResourceAction) SetActionVerb(value *string)() {
    err := m.GetBackingStore().Set("actionVerb", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationContext sets the authenticationContext property value. The authenticationContext property
func (m *UnifiedRbacResourceAction) SetAuthenticationContext(value AuthenticationContextClassReferenceable)() {
    err := m.GetBackingStore().Set("authenticationContext", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationContextId sets the authenticationContextId property value. The authenticationContextId property
func (m *UnifiedRbacResourceAction) SetAuthenticationContextId(value *string)() {
    err := m.GetBackingStore().Set("authenticationContextId", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description for the action. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAuthenticationContextSettable sets the isAuthenticationContextSettable property value. The isAuthenticationContextSettable property
func (m *UnifiedRbacResourceAction) SetIsAuthenticationContextSettable(value *bool)() {
    err := m.GetBackingStore().Set("isAuthenticationContextSettable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPrivileged sets the isPrivileged property value. The isPrivileged property
func (m *UnifiedRbacResourceAction) SetIsPrivileged(value *bool)() {
    err := m.GetBackingStore().Set("isPrivileged", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name for the action within the resource namespace, such as microsoft.insights/programs/update. Can include slash character (/). Case insensitive. Required. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceScope sets the resourceScope property value. The resourceScope property
func (m *UnifiedRbacResourceAction) SetResourceScope(value UnifiedRbacResourceScopeable)() {
    err := m.GetBackingStore().Set("resourceScope", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceScopeId sets the resourceScopeId property value. Not implemented.
func (m *UnifiedRbacResourceAction) SetResourceScopeId(value *string)() {
    err := m.GetBackingStore().Set("resourceScopeId", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRbacResourceActionable 
type UnifiedRbacResourceActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionVerb()(*string)
    GetAuthenticationContext()(AuthenticationContextClassReferenceable)
    GetAuthenticationContextId()(*string)
    GetDescription()(*string)
    GetIsAuthenticationContextSettable()(*bool)
    GetIsPrivileged()(*bool)
    GetName()(*string)
    GetResourceScope()(UnifiedRbacResourceScopeable)
    GetResourceScopeId()(*string)
    SetActionVerb(value *string)()
    SetAuthenticationContext(value AuthenticationContextClassReferenceable)()
    SetAuthenticationContextId(value *string)()
    SetDescription(value *string)()
    SetIsAuthenticationContextSettable(value *bool)()
    SetIsPrivileged(value *bool)()
    SetName(value *string)()
    SetResourceScope(value UnifiedRbacResourceScopeable)()
    SetResourceScopeId(value *string)()
}
