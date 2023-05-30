package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceResource 
type GovernanceResource struct {
    Entity
}
// NewGovernanceResource instantiates a new governanceResource and sets the default values.
func NewGovernanceResource()(*GovernanceResource) {
    m := &GovernanceResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceResource(), nil
}
// GetDisplayName gets the displayName property value. The display name of the resource.
func (m *GovernanceResource) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
func (m *GovernanceResource) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(GovernanceResourceable))
        }
        return nil
    }
    res["registeredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredDateTime(val)
        }
        return nil
    }
    res["registeredRoot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredRoot(val)
        }
        return nil
    }
    res["roleAssignmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleAssignmentRequestable)
            }
            m.SetRoleAssignmentRequests(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleAssignmentable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleDefinitionable)
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleSettingable)
            }
            m.SetRoleSettings(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetParent gets the parent property value. Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
func (m *GovernanceResource) GetParent()(GovernanceResourceable) {
    val, err := m.GetBackingStore().Get("parent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceResourceable)
    }
    return nil
}
// GetRegisteredDateTime gets the registeredDateTime property value. Represents the date time when the resource is registered in PIM.
func (m *GovernanceResource) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("registeredDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRegisteredRoot gets the registeredRoot property value. The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
func (m *GovernanceResource) GetRegisteredRoot()(*string) {
    val, err := m.GetBackingStore().Get("registeredRoot")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleAssignmentRequests gets the roleAssignmentRequests property value. The collection of role assignment requests for the resource.
func (m *GovernanceResource) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable) {
    val, err := m.GetBackingStore().Get("roleAssignmentRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRoleAssignmentRequestable)
    }
    return nil
}
// GetRoleAssignments gets the roleAssignments property value. The collection of role assignments for the resource.
func (m *GovernanceResource) GetRoleAssignments()([]GovernanceRoleAssignmentable) {
    val, err := m.GetBackingStore().Get("roleAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRoleAssignmentable)
    }
    return nil
}
// GetRoleDefinitions gets the roleDefinitions property value. The collection of role defintions for the resource.
func (m *GovernanceResource) GetRoleDefinitions()([]GovernanceRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRoleDefinitionable)
    }
    return nil
}
// GetRoleSettings gets the roleSettings property value. The collection of role settings for the resource.
func (m *GovernanceResource) GetRoleSettings()([]GovernanceRoleSettingable) {
    val, err := m.GetBackingStore().Get("roleSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRoleSettingable)
    }
    return nil
}
// GetStatus gets the status property value. The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
func (m *GovernanceResource) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetType gets the type property value. Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
func (m *GovernanceResource) GetType()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GovernanceResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registeredDateTime", m.GetRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registeredRoot", m.GetRegisteredRoot())
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentRequests()))
        for i, v := range m.GetRoleAssignmentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleSettings()))
        for i, v := range m.GetRoleSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
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
// SetDisplayName sets the displayName property value. The display name of the resource.
func (m *GovernanceResource) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
func (m *GovernanceResource) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetParent sets the parent property value. Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
func (m *GovernanceResource) SetParent(value GovernanceResourceable)() {
    err := m.GetBackingStore().Set("parent", value)
    if err != nil {
        panic(err)
    }
}
// SetRegisteredDateTime sets the registeredDateTime property value. Represents the date time when the resource is registered in PIM.
func (m *GovernanceResource) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("registeredDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRegisteredRoot sets the registeredRoot property value. The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
func (m *GovernanceResource) SetRegisteredRoot(value *string)() {
    err := m.GetBackingStore().Set("registeredRoot", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignmentRequests sets the roleAssignmentRequests property value. The collection of role assignment requests for the resource.
func (m *GovernanceResource) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)() {
    err := m.GetBackingStore().Set("roleAssignmentRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignments sets the roleAssignments property value. The collection of role assignments for the resource.
func (m *GovernanceResource) SetRoleAssignments(value []GovernanceRoleAssignmentable)() {
    err := m.GetBackingStore().Set("roleAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. The collection of role defintions for the resource.
func (m *GovernanceResource) SetRoleDefinitions(value []GovernanceRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleSettings sets the roleSettings property value. The collection of role settings for the resource.
func (m *GovernanceResource) SetRoleSettings(value []GovernanceRoleSettingable)() {
    err := m.GetBackingStore().Set("roleSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
func (m *GovernanceResource) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetType sets the type property value. Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
func (m *GovernanceResource) SetType(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// GovernanceResourceable 
type GovernanceResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetParent()(GovernanceResourceable)
    GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegisteredRoot()(*string)
    GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable)
    GetRoleAssignments()([]GovernanceRoleAssignmentable)
    GetRoleDefinitions()([]GovernanceRoleDefinitionable)
    GetRoleSettings()([]GovernanceRoleSettingable)
    GetStatus()(*string)
    GetType()(*string)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetParent(value GovernanceResourceable)()
    SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegisteredRoot(value *string)()
    SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)()
    SetRoleAssignments(value []GovernanceRoleAssignmentable)()
    SetRoleDefinitions(value []GovernanceRoleDefinitionable)()
    SetRoleSettings(value []GovernanceRoleSettingable)()
    SetStatus(value *string)()
    SetType(value *string)()
}
