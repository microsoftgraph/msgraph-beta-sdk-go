package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new governanceResource and sets the default values.
func NewGovernanceResource()(*GovernanceResource) {
    m := &GovernanceResource{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name of the resource.
func (m *GovernanceResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the externalId property value. The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
func (m *GovernanceResource) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the parent property value. Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
func (m *GovernanceResource) GetParent()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// Gets the registeredDateTime property value. Represents the date time when the resource is registered in PIM.
func (m *GovernanceResource) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registeredDateTime
    }
}
// Gets the registeredRoot property value. The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
func (m *GovernanceResource) GetRegisteredRoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.registeredRoot
    }
}
// Gets the roleAssignmentRequests property value. The collection of role assignment requests for the resource.
func (m *GovernanceResource) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentRequests
    }
}
// Gets the roleAssignments property value. The collection of role assignments for the resource.
func (m *GovernanceResource) GetRoleAssignments()([]GovernanceRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// Gets the roleDefinitions property value. The collection of role defintions for the resource.
func (m *GovernanceResource) GetRoleDefinitions()([]GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// Gets the roleSettings property value. The collection of role settings for the resource.
func (m *GovernanceResource) GetRoleSettings()([]GovernanceRoleSetting) {
    if m == nil {
        return nil
    } else {
        return m.roleSettings
    }
}
// Gets the status property value. The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
func (m *GovernanceResource) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the type_escaped property value. Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
func (m *GovernanceResource) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceResource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
    {
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
    {
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
    {
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
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name of the resource.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GovernanceResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the externalId property value. The external id of the resource, representing its original id in the external system. For example, a subscription resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *GovernanceResource) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the parent property value. Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs to.
// Parameters:
//  - value : Value to set for the parent property.
func (m *GovernanceResource) SetParent(value *GovernanceResource)() {
    m.parent = value
}
// Sets the registeredDateTime property value. Represents the date time when the resource is registered in PIM.
// Parameters:
//  - value : Value to set for the registeredDateTime property.
func (m *GovernanceResource) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registeredDateTime = value
}
// Sets the registeredRoot property value. The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent, or higher ancestor resources.
// Parameters:
//  - value : Value to set for the registeredRoot property.
func (m *GovernanceResource) SetRegisteredRoot(value *string)() {
    m.registeredRoot = value
}
// Sets the roleAssignmentRequests property value. The collection of role assignment requests for the resource.
// Parameters:
//  - value : Value to set for the roleAssignmentRequests property.
func (m *GovernanceResource) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequest)() {
    m.roleAssignmentRequests = value
}
// Sets the roleAssignments property value. The collection of role assignments for the resource.
// Parameters:
//  - value : Value to set for the roleAssignments property.
func (m *GovernanceResource) SetRoleAssignments(value []GovernanceRoleAssignment)() {
    m.roleAssignments = value
}
// Sets the roleDefinitions property value. The collection of role defintions for the resource.
// Parameters:
//  - value : Value to set for the roleDefinitions property.
func (m *GovernanceResource) SetRoleDefinitions(value []GovernanceRoleDefinition)() {
    m.roleDefinitions = value
}
// Sets the roleSettings property value. The collection of role settings for the resource.
// Parameters:
//  - value : Value to set for the roleSettings property.
func (m *GovernanceResource) SetRoleSettings(value []GovernanceRoleSetting)() {
    m.roleSettings = value
}
// Sets the status property value. The status of a given resource. For example, it could represent whether the resource is locked or not (values: Active/Locked). Note: This property may be extended in the future to support more scenarios.
// Parameters:
//  - value : Value to set for the status property.
func (m *GovernanceResource) SetStatus(value *string)() {
    m.status = value
}
// Sets the type_escaped property value. Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup', 'Microsoft.Sql/server', etc.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *GovernanceResource) SetType_escaped(value *string)() {
    m.type_escaped = value
}
