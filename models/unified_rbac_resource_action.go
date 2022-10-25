package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRbacResourceAction provides operations to manage the collection of accessReview entities.
type UnifiedRbacResourceAction struct {
    Entity
    // HTTP method for the action, such as DELETE, GET, PATCH, POST, PUT, or null. Supports $filter (eq) but not for null values.
    actionVerb *string
    // Description for the action. Supports $filter (eq).
    description *string
    // Name for the action within the resource namespace, such as microsoft.insights/programs/update. Can include slash character (/). Case insensitive. Required. Supports $filter (eq).
    name *string
    // The resourceScope property
    resourceScope UnifiedRbacResourceScopeable
    // Not implemented.
    resourceScopeId *string
}
// NewUnifiedRbacResourceAction instantiates a new unifiedRbacResourceAction and sets the default values.
func NewUnifiedRbacResourceAction()(*UnifiedRbacResourceAction) {
    m := &UnifiedRbacResourceAction{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRbacResourceAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUnifiedRbacResourceActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRbacResourceActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacResourceAction(), nil
}
// GetActionVerb gets the actionVerb property value. HTTP method for the action, such as DELETE, GET, PATCH, POST, PUT, or null. Supports $filter (eq) but not for null values.
func (m *UnifiedRbacResourceAction) GetActionVerb()(*string) {
    return m.actionVerb
}
// GetDescription gets the description property value. Description for the action. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionVerb"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActionVerb)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["resourceScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUnifiedRbacResourceScopeFromDiscriminatorValue , m.SetResourceScope)
    res["resourceScopeId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourceScopeId)
    return res
}
// GetName gets the name property value. Name for the action within the resource namespace, such as microsoft.insights/programs/update. Can include slash character (/). Case insensitive. Required. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) GetName()(*string) {
    return m.name
}
// GetResourceScope gets the resourceScope property value. The resourceScope property
func (m *UnifiedRbacResourceAction) GetResourceScope()(UnifiedRbacResourceScopeable) {
    return m.resourceScope
}
// GetResourceScopeId gets the resourceScopeId property value. Not implemented.
func (m *UnifiedRbacResourceAction) GetResourceScopeId()(*string) {
    return m.resourceScopeId
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
        err = writer.WriteStringValue("description", m.GetDescription())
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
    m.actionVerb = value
}
// SetDescription sets the description property value. Description for the action. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. Name for the action within the resource namespace, such as microsoft.insights/programs/update. Can include slash character (/). Case insensitive. Required. Supports $filter (eq).
func (m *UnifiedRbacResourceAction) SetName(value *string)() {
    m.name = value
}
// SetResourceScope sets the resourceScope property value. The resourceScope property
func (m *UnifiedRbacResourceAction) SetResourceScope(value UnifiedRbacResourceScopeable)() {
    m.resourceScope = value
}
// SetResourceScopeId sets the resourceScopeId property value. Not implemented.
func (m *UnifiedRbacResourceAction) SetResourceScopeId(value *string)() {
    m.resourceScopeId = value
}
