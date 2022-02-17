package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResource 
type AccessPackageResource struct {
    Entity
    // Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.
    accessPackageResourceEnvironment *AccessPackageResourceEnvironment;
    // Read-only. Nullable. Supports $expand.
    accessPackageResourceRoles []AccessPackageResourceRole;
    // Read-only. Nullable. Supports $expand.
    accessPackageResourceScopes []AccessPackageResourceScope;
    // The name of the user or application that first added this resource. Read-only.
    addedBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    addedOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Contains information about the attributes to be collected from the requestor and sent to the resource application.
    attributes []AccessPackageResourceAttribute;
    // A description for the resource.
    description *string;
    // The display name of the resource, such as the application name, group name or site name.
    displayName *string;
    // True if the resource is not yet available for assignment.
    isPendingOnboarding *bool;
    // The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
    originId *string;
    // The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
    originSystem *string;
    // The type of the resource, such as Application if it is an Azure AD connected application, or SharePoint Online Site for a SharePoint Online site.
    resourceType *string;
    // A unique resource locator for the resource, such as the URL for signing a user into an application.
    url *string;
}
// NewAccessPackageResource instantiates a new accessPackageResource and sets the default values.
func NewAccessPackageResource()(*AccessPackageResource) {
    m := &AccessPackageResource{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageResourceEnvironment gets the accessPackageResourceEnvironment property value. Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.
func (m *AccessPackageResource) GetAccessPackageResourceEnvironment()(*AccessPackageResourceEnvironment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceEnvironment
    }
}
// GetAccessPackageResourceRoles gets the accessPackageResourceRoles property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) GetAccessPackageResourceRoles()([]AccessPackageResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRoles
    }
}
// GetAccessPackageResourceScopes gets the accessPackageResourceScopes property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) GetAccessPackageResourceScopes()([]AccessPackageResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceScopes
    }
}
// GetAddedBy gets the addedBy property value. The name of the user or application that first added this resource. Read-only.
func (m *AccessPackageResource) GetAddedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addedBy
    }
}
// GetAddedOn gets the addedOn property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResource) GetAddedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.addedOn
    }
}
// GetAttributes gets the attributes property value. Contains information about the attributes to be collected from the requestor and sent to the resource application.
func (m *AccessPackageResource) GetAttributes()([]AccessPackageResourceAttribute) {
    if m == nil {
        return nil
    } else {
        return m.attributes
    }
}
// GetDescription gets the description property value. A description for the resource.
func (m *AccessPackageResource) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the resource, such as the application name, group name or site name.
func (m *AccessPackageResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsPendingOnboarding gets the isPendingOnboarding property value. True if the resource is not yet available for assignment.
func (m *AccessPackageResource) GetIsPendingOnboarding()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPendingOnboarding
    }
}
// GetOriginId gets the originId property value. The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
func (m *AccessPackageResource) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResource) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
// GetResourceType gets the resourceType property value. The type of the resource, such as Application if it is an Azure AD connected application, or SharePoint Online Site for a SharePoint Online site.
func (m *AccessPackageResource) GetResourceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceType
    }
}
// GetUrl gets the url property value. A unique resource locator for the resource, such as the URL for signing a user into an application.
func (m *AccessPackageResource) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResourceEnvironment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceEnvironment() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceEnvironment(val.(*AccessPackageResourceEnvironment))
        }
        return nil
    }
    res["accessPackageResourceRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRole() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRole, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResourceRole))
            }
            m.SetAccessPackageResourceRoles(res)
        }
        return nil
    }
    res["accessPackageResourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResourceScope))
            }
            m.SetAccessPackageResourceScopes(res)
        }
        return nil
    }
    res["addedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedBy(val)
        }
        return nil
    }
    res["addedOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedOn(val)
        }
        return nil
    }
    res["attributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttribute() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceAttribute, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResourceAttribute))
            }
            m.SetAttributes(res)
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
    res["isPendingOnboarding"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPendingOnboarding(val)
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
    res["resourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
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
func (m *AccessPackageResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceRoles()))
        for i, v := range m.GetAccessPackageResourceRoles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceScopes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceScopes()))
        for i, v := range m.GetAccessPackageResourceScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
// SetAccessPackageResourceEnvironment sets the accessPackageResourceEnvironment property value. Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.
func (m *AccessPackageResource) SetAccessPackageResourceEnvironment(value *AccessPackageResourceEnvironment)() {
    if m != nil {
        m.accessPackageResourceEnvironment = value
    }
}
// SetAccessPackageResourceRoles sets the accessPackageResourceRoles property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) SetAccessPackageResourceRoles(value []AccessPackageResourceRole)() {
    if m != nil {
        m.accessPackageResourceRoles = value
    }
}
// SetAccessPackageResourceScopes sets the accessPackageResourceScopes property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) SetAccessPackageResourceScopes(value []AccessPackageResourceScope)() {
    if m != nil {
        m.accessPackageResourceScopes = value
    }
}
// SetAddedBy sets the addedBy property value. The name of the user or application that first added this resource. Read-only.
func (m *AccessPackageResource) SetAddedBy(value *string)() {
    if m != nil {
        m.addedBy = value
    }
}
// SetAddedOn sets the addedOn property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResource) SetAddedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.addedOn = value
    }
}
// SetAttributes sets the attributes property value. Contains information about the attributes to be collected from the requestor and sent to the resource application.
func (m *AccessPackageResource) SetAttributes(value []AccessPackageResourceAttribute)() {
    if m != nil {
        m.attributes = value
    }
}
// SetDescription sets the description property value. A description for the resource.
func (m *AccessPackageResource) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the resource, such as the application name, group name or site name.
func (m *AccessPackageResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsPendingOnboarding sets the isPendingOnboarding property value. True if the resource is not yet available for assignment.
func (m *AccessPackageResource) SetIsPendingOnboarding(value *bool)() {
    if m != nil {
        m.isPendingOnboarding = value
    }
}
// SetOriginId sets the originId property value. The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group.
func (m *AccessPackageResource) SetOriginId(value *string)() {
    if m != nil {
        m.originId = value
    }
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResource) SetOriginSystem(value *string)() {
    if m != nil {
        m.originSystem = value
    }
}
// SetResourceType sets the resourceType property value. The type of the resource, such as Application if it is an Azure AD connected application, or SharePoint Online Site for a SharePoint Online site.
func (m *AccessPackageResource) SetResourceType(value *string)() {
    if m != nil {
        m.resourceType = value
    }
}
// SetUrl sets the url property value. A unique resource locator for the resource, such as the URL for signing a user into an application.
func (m *AccessPackageResource) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
