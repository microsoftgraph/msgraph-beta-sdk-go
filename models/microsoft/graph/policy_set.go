package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PolicySet struct {
    Entity
    // Assignments of the PolicySet.
    assignments []PolicySetAssignment;
    // Creation time of the PolicySet.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the PolicySet.
    description *string;
    // DisplayName of the PolicySet.
    displayName *string;
    // Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
    errorCode *ErrorCode;
    // Tags of the guided deployment
    guidedDeploymentTags []string;
    // Items of the PolicySet with maximum count 100.
    items []PolicySetItem;
    // Last modified time of the PolicySet.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // RoleScopeTags of the PolicySet
    roleScopeTags []string;
    // Validation/assignment status of the PolicySet. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
    status *PolicySetStatus;
}
// Instantiates a new policySet and sets the default values.
func NewPolicySet()(*PolicySet) {
    m := &PolicySet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. Assignments of the PolicySet.
func (m *PolicySet) GetAssignments()([]PolicySetAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. Creation time of the PolicySet.
func (m *PolicySet) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description of the PolicySet.
func (m *PolicySet) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. DisplayName of the PolicySet.
func (m *PolicySet) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the errorCode property value. Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
func (m *PolicySet) GetErrorCode()(*ErrorCode) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySet) GetGuidedDeploymentTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.guidedDeploymentTags
    }
}
// Gets the items property value. Items of the PolicySet with maximum count 100.
func (m *PolicySet) GetItems()([]PolicySetItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// Gets the lastModifiedDateTime property value. Last modified time of the PolicySet.
func (m *PolicySet) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the roleScopeTags property value. RoleScopeTags of the PolicySet
func (m *PolicySet) GetRoleScopeTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTags
    }
}
// Gets the status property value. Validation/assignment status of the PolicySet. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
func (m *PolicySet) GetStatus()(*PolicySetStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *PolicySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPolicySetAssignment() })
        if err != nil {
            return err
        }
        res := make([]PolicySetAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*PolicySetAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseErrorCode)
        if err != nil {
            return err
        }
        cast := val.(ErrorCode)
        m.SetErrorCode(&cast)
        return nil
    }
    res["guidedDeploymentTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetGuidedDeploymentTags(res)
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPolicySetItem() })
        if err != nil {
            return err
        }
        res := make([]PolicySetItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*PolicySetItem))
        }
        m.SetItems(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["roleScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoleScopeTags(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicySetStatus)
        if err != nil {
            return err
        }
        cast := val.(PolicySetStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *PolicySet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PolicySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetErrorCode() != nil {
        cast := m.GetErrorCode().String()
        err = writer.WriteStringValue("errorCode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("guidedDeploymentTags", m.GetGuidedDeploymentTags())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTags", m.GetRoleScopeTags())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. Assignments of the PolicySet.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *PolicySet) SetAssignments(value []PolicySetAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. Creation time of the PolicySet.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *PolicySet) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description of the PolicySet.
// Parameters:
//  - value : Value to set for the description property.
func (m *PolicySet) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. DisplayName of the PolicySet.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PolicySet) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the errorCode property value. Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *PolicySet) SetErrorCode(value *ErrorCode)() {
    m.errorCode = value
}
// Sets the guidedDeploymentTags property value. Tags of the guided deployment
// Parameters:
//  - value : Value to set for the guidedDeploymentTags property.
func (m *PolicySet) SetGuidedDeploymentTags(value []string)() {
    m.guidedDeploymentTags = value
}
// Sets the items property value. Items of the PolicySet with maximum count 100.
// Parameters:
//  - value : Value to set for the items property.
func (m *PolicySet) SetItems(value []PolicySetItem)() {
    m.items = value
}
// Sets the lastModifiedDateTime property value. Last modified time of the PolicySet.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *PolicySet) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the roleScopeTags property value. RoleScopeTags of the PolicySet
// Parameters:
//  - value : Value to set for the roleScopeTags property.
func (m *PolicySet) SetRoleScopeTags(value []string)() {
    m.roleScopeTags = value
}
// Sets the status property value. Validation/assignment status of the PolicySet. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
// Parameters:
//  - value : Value to set for the status property.
func (m *PolicySet) SetStatus(value *PolicySetStatus)() {
    m.status = value
}
