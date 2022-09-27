package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResource 
type AccessPackageResource struct {
    Entity
    // Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
    accessPackageResourceEnvironment AccessPackageResourceEnvironmentable
    // Read-only. Nullable. Supports $expand.
    accessPackageResourceRoles []AccessPackageResourceRoleable
    // Read-only. Nullable. Supports $expand.
    accessPackageResourceScopes []AccessPackageResourceScopeable
    // The name of the user or application that first added this resource. Read-only.
    addedBy *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    addedOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains information about the attributes to be collected from the requestor and sent to the resource application.
    attributes []AccessPackageResourceAttributeable
    // A description for the resource.
    description *string
    // The display name of the resource, such as the application name, group name or site name.
    displayName *string
    // True if the resource is not yet available for assignment.
    isPendingOnboarding *bool
    // The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
    originId *string
    // The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
    originSystem *string
    // The type of the resource, such as Application if it is an Azure AD connected application, or SharePoint Online Site for a SharePoint Online site.
    resourceType *string
    // A unique resource locator for the resource, such as the URL for signing a user into an application.
    url *string
}
// NewAccessPackageResource instantiates a new accessPackageResource and sets the default values.
func NewAccessPackageResource()(*AccessPackageResource) {
    m := &AccessPackageResource{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageResource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResource(), nil
}
// GetAccessPackageResourceEnvironment gets the accessPackageResourceEnvironment property value. Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *AccessPackageResource) GetAccessPackageResourceEnvironment()(AccessPackageResourceEnvironmentable) {
    return m.accessPackageResourceEnvironment
}
// GetAccessPackageResourceRoles gets the accessPackageResourceRoles property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable) {
    return m.accessPackageResourceRoles
}
// GetAccessPackageResourceScopes gets the accessPackageResourceScopes property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable) {
    return m.accessPackageResourceScopes
}
// GetAddedBy gets the addedBy property value. The name of the user or application that first added this resource. Read-only.
func (m *AccessPackageResource) GetAddedBy()(*string) {
    return m.addedBy
}
// GetAddedOn gets the addedOn property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResource) GetAddedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.addedOn
}
// GetAttributes gets the attributes property value. Contains information about the attributes to be collected from the requestor and sent to the resource application.
func (m *AccessPackageResource) GetAttributes()([]AccessPackageResourceAttributeable) {
    return m.attributes
}
// GetDescription gets the description property value. A description for the resource.
func (m *AccessPackageResource) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the resource, such as the application name, group name or site name.
func (m *AccessPackageResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResourceEnvironment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageResourceEnvironmentFromDiscriminatorValue , m.SetAccessPackageResourceEnvironment)
    res["accessPackageResourceRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceRoleFromDiscriminatorValue , m.SetAccessPackageResourceRoles)
    res["accessPackageResourceScopes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceScopeFromDiscriminatorValue , m.SetAccessPackageResourceScopes)
    res["addedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAddedBy)
    res["addedOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetAddedOn)
    res["attributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceAttributeFromDiscriminatorValue , m.SetAttributes)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isPendingOnboarding"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPendingOnboarding)
    res["originId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginId)
    res["originSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginSystem)
    res["resourceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourceType)
    res["url"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUrl)
    return res
}
// GetIsPendingOnboarding gets the isPendingOnboarding property value. True if the resource is not yet available for assignment.
func (m *AccessPackageResource) GetIsPendingOnboarding()(*bool) {
    return m.isPendingOnboarding
}
// GetOriginId gets the originId property value. The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
func (m *AccessPackageResource) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResource) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetResourceType gets the resourceType property value. The type of the resource, such as Application if it is an Azure AD connected application, or SharePoint Online Site for a SharePoint Online site.
func (m *AccessPackageResource) GetResourceType()(*string) {
    return m.resourceType
}
// GetUrl gets the url property value. A unique resource locator for the resource, such as the URL for signing a user into an application.
func (m *AccessPackageResource) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *AccessPackageResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceEnvironment", m.GetAccessPackageResourceEnvironment())
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceRoles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageResourceRoles())
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceScopes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageResourceScopes())
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("addedBy", m.GetAddedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("addedOn", m.GetAddedOn())
        if err != nil {
            return err
        }
    }
    if m.GetAttributes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttributes())
        err = writer.WriteCollectionOfObjectValues("attributes", cast)
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
        err = writer.WriteBoolValue("isPendingOnboarding", m.GetIsPendingOnboarding())
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
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
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
// SetAccessPackageResourceEnvironment sets the accessPackageResourceEnvironment property value. Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *AccessPackageResource) SetAccessPackageResourceEnvironment(value AccessPackageResourceEnvironmentable)() {
    m.accessPackageResourceEnvironment = value
}
// SetAccessPackageResourceRoles sets the accessPackageResourceRoles property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)() {
    m.accessPackageResourceRoles = value
}
// SetAccessPackageResourceScopes sets the accessPackageResourceScopes property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)() {
    m.accessPackageResourceScopes = value
}
// SetAddedBy sets the addedBy property value. The name of the user or application that first added this resource. Read-only.
func (m *AccessPackageResource) SetAddedBy(value *string)() {
    m.addedBy = value
}
// SetAddedOn sets the addedOn property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResource) SetAddedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.addedOn = value
}
// SetAttributes sets the attributes property value. Contains information about the attributes to be collected from the requestor and sent to the resource application.
func (m *AccessPackageResource) SetAttributes(value []AccessPackageResourceAttributeable)() {
    m.attributes = value
}
// SetDescription sets the description property value. A description for the resource.
func (m *AccessPackageResource) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the resource, such as the application name, group name or site name.
func (m *AccessPackageResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsPendingOnboarding sets the isPendingOnboarding property value. True if the resource is not yet available for assignment.
func (m *AccessPackageResource) SetIsPendingOnboarding(value *bool)() {
    m.isPendingOnboarding = value
}
// SetOriginId sets the originId property value. The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
func (m *AccessPackageResource) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResource) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetResourceType sets the resourceType property value. The type of the resource, such as Application if it is an Azure AD connected application, or SharePoint Online Site for a SharePoint Online site.
func (m *AccessPackageResource) SetResourceType(value *string)() {
    m.resourceType = value
}
// SetUrl sets the url property value. A unique resource locator for the resource, such as the URL for signing a user into an application.
func (m *AccessPackageResource) SetUrl(value *string)() {
    m.url = value
}
