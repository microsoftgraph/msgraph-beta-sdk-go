package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceResource 
type GovernanceResource struct {
    Entity
    // The display name of the resource.
    displayName *string;
    // The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
    externalId *string;
    // Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
    parent *GovernanceResource;
    // Represents the date time when the resource is registered in PIM.
    registeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
    registeredRoot *string;
    // The collection of role assignment requests for the resource.
    roleAssignmentRequests []GovernanceRoleAssignmentRequest;
    // The collection of role assignments for the resource.
    roleAssignments []GovernanceRoleAssignment;
    // The collection of role defintions for the resource.
    roleDefinitions []GovernanceRoleDefinition;
    // The collection of role settings for the resource.
    roleSettings []GovernanceRoleSetting;
    // The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
    status *string;
    // Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
    type_escaped *string;
}
// NewGovernanceResource instantiates a new governanceResource and sets the default values.
func NewGovernanceResource()(*GovernanceResource) {
    m := &GovernanceResource{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The display name of the resource.
func (m *GovernanceResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalId gets the externalId property value. The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
func (m *GovernanceResource) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetParent gets the parent property value. Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
func (m *GovernanceResource) GetParent()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// GetRegisteredDateTime gets the registeredDateTime property value. Represents the date time when the resource is registered in PIM.
func (m *GovernanceResource) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registeredDateTime
    }
}
// GetRegisteredRoot gets the registeredRoot property value. The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
func (m *GovernanceResource) GetRegisteredRoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.registeredRoot
    }
}
// GetRoleAssignmentRequests gets the roleAssignmentRequests property value. The collection of role assignment requests for the resource.
func (m *GovernanceResource) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentRequests
    }
}
// GetRoleAssignments gets the roleAssignments property value. The collection of role assignments for the resource.
func (m *GovernanceResource) GetRoleAssignments()([]GovernanceRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. The collection of role defintions for the resource.
func (m *GovernanceResource) GetRoleDefinitions()([]GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// GetRoleSettings gets the roleSettings property value. The collection of role settings for the resource.
func (m *GovernanceResource) GetRoleSettings()([]GovernanceRoleSetting) {
    if m == nil {
        return nil
    } else {
        return m.roleSettings
    }
}
// GetStatus gets the status property value. The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
func (m *GovernanceResource) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetType gets the type property value. Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
func (m *GovernanceResource) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["parent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(*GovernanceResource))
        }
        return nil
    }
    res["registeredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredDateTime(val)
        }
        return nil
    }
    res["registeredRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredRoot(val)
        }
        return nil
    }
    res["roleAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleAssignmentRequest))
            }
            m.SetRoleAssignmentRequests(res)
        }
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleAssignment))
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleDefinition))
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceRoleSetting))
            }
            m.SetRoleSettings(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
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
func (m *GovernanceResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignmentRequests()))
        for i, v := range m.GetRoleAssignmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleSettings()))
        for i, v := range m.GetRoleSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    if m != nil {
        m.displayName = value
    }
}
// SetExternalId sets the externalId property value. The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
func (m *GovernanceResource) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetParent sets the parent property value. Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
func (m *GovernanceResource) SetParent(value *GovernanceResource)() {
    if m != nil {
        m.parent = value
    }
}
// SetRegisteredDateTime sets the registeredDateTime property value. Represents the date time when the resource is registered in PIM.
func (m *GovernanceResource) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.registeredDateTime = value
    }
}
// SetRegisteredRoot sets the registeredRoot property value. The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
func (m *GovernanceResource) SetRegisteredRoot(value *string)() {
    if m != nil {
        m.registeredRoot = value
    }
}
// SetRoleAssignmentRequests sets the roleAssignmentRequests property value. The collection of role assignment requests for the resource.
func (m *GovernanceResource) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequest)() {
    if m != nil {
        m.roleAssignmentRequests = value
    }
}
// SetRoleAssignments sets the roleAssignments property value. The collection of role assignments for the resource.
func (m *GovernanceResource) SetRoleAssignments(value []GovernanceRoleAssignment)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. The collection of role defintions for the resource.
func (m *GovernanceResource) SetRoleDefinitions(value []GovernanceRoleDefinition)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
// SetRoleSettings sets the roleSettings property value. The collection of role settings for the resource.
func (m *GovernanceResource) SetRoleSettings(value []GovernanceRoleSetting)() {
    if m != nil {
        m.roleSettings = value
    }
}
// SetStatus sets the status property value. The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
func (m *GovernanceResource) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetType sets the type property value. Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
func (m *GovernanceResource) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
