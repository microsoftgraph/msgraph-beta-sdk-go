package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupPolicyDefinition struct {
    Entity
    // The group policy category associated with the definition.
    category *GroupPolicyCategory;
    // The localized full category path for the policy.
    categoryPath *string;
    // Identifies the type of groups the policy can be applied to. Possible values are: user, machine.
    classType *GroupPolicyDefinitionClassType;
    // The group policy file associated with the definition.
    definitionFile *GroupPolicyDefinitionFile;
    // The localized policy name.
    displayName *string;
    // The localized explanation or help text associated with the policy. The default value is empty.
    explainText *string;
    // The category id of the parent category
    groupPolicyCategoryId *string;
    // Signifies whether or not there are related definitions to this definition
    hasRelatedDefinitions *bool;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Minimum required CSP version for device configuration in this definition
    minDeviceCspVersion *string;
    // Minimum required CSP version for user configuration in this definition
    minUserCspVersion *string;
    // Definition of the next version of this definition
    nextVersionDefinition *GroupPolicyDefinition;
    // Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
    policyType *GroupPolicyType;
    // The group policy presentations associated with the definition.
    presentations []GroupPolicyPresentation;
    // Definition of the previous version of this definition
    previousVersionDefinition *GroupPolicyDefinition;
    // Localized string used to specify what operating system or application version is affected by the policy.
    supportedOn *string;
    // Setting definition version
    version *string;
}
// Instantiates a new groupPolicyDefinition and sets the default values.
func NewGroupPolicyDefinition()(*GroupPolicyDefinition) {
    m := &GroupPolicyDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the category property value. The group policy category associated with the definition.
func (m *GroupPolicyDefinition) GetCategory()(*GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) GetCategoryPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.categoryPath
    }
}
// Gets the classType property value. Identifies the type of groups the policy can be applied to. Possible values are: user, machine.
func (m *GroupPolicyDefinition) GetClassType()(*GroupPolicyDefinitionClassType) {
    if m == nil {
        return nil
    } else {
        return m.classType
    }
}
// Gets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) GetDefinitionFile()(*GroupPolicyDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.definitionFile
    }
}
// Gets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) GetExplainText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.explainText
    }
}
// Gets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) GetGroupPolicyCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCategoryId
    }
}
// Gets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) GetHasRelatedDefinitions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasRelatedDefinitions
    }
}
// Gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) GetMinDeviceCspVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minDeviceCspVersion
    }
}
// Gets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) GetMinUserCspVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minUserCspVersion
    }
}
// Gets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) GetNextVersionDefinition()(*GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.nextVersionDefinition
    }
}
// Gets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
func (m *GroupPolicyDefinition) GetPolicyType()(*GroupPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
// Gets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) GetPresentations()([]GroupPolicyPresentation) {
    if m == nil {
        return nil
    } else {
        return m.presentations
    }
}
// Gets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) GetPreviousVersionDefinition()(*GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.previousVersionDefinition
    }
}
// Gets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) GetSupportedOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportedOn
    }
}
// Gets the version property value. Setting definition version
func (m *GroupPolicyDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *GroupPolicyDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*GroupPolicyCategory))
        }
        return nil
    }
    res["categoryPath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryPath(val)
        }
        return nil
    }
    res["classType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyDefinitionClassType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(GroupPolicyDefinitionClassType)
            m.SetClassType(&cast)
        }
        return nil
    }
    res["definitionFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinitionFile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionFile(val.(*GroupPolicyDefinitionFile))
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
    res["explainText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExplainText(val)
        }
        return nil
    }
    res["groupPolicyCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyCategoryId(val)
        }
        return nil
    }
    res["hasRelatedDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasRelatedDefinitions(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["minDeviceCspVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinDeviceCspVersion(val)
        }
        return nil
    }
    res["minUserCspVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinUserCspVersion(val)
        }
        return nil
    }
    res["nextVersionDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextVersionDefinition(val.(*GroupPolicyDefinition))
        }
        return nil
    }
    res["policyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(GroupPolicyType)
            m.SetPolicyType(&cast)
        }
        return nil
    }
    res["presentations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyPresentation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentation, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyPresentation))
            }
            m.SetPresentations(res)
        }
        return nil
    }
    res["previousVersionDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousVersionDefinition(val.(*GroupPolicyDefinition))
        }
        return nil
    }
    res["supportedOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedOn(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *GroupPolicyDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupPolicyDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("categoryPath", m.GetCategoryPath())
        if err != nil {
            return err
        }
    }
    if m.GetClassType() != nil {
        cast := m.GetClassType().String()
        err = writer.WriteStringValue("classType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definitionFile", m.GetDefinitionFile())
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
        err = writer.WriteStringValue("explainText", m.GetExplainText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupPolicyCategoryId", m.GetGroupPolicyCategoryId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasRelatedDefinitions", m.GetHasRelatedDefinitions())
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
        err = writer.WriteStringValue("minDeviceCspVersion", m.GetMinDeviceCspVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minUserCspVersion", m.GetMinUserCspVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("nextVersionDefinition", m.GetNextVersionDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := m.GetPolicyType().String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPresentations()))
        for i, v := range m.GetPresentations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("presentations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("previousVersionDefinition", m.GetPreviousVersionDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("supportedOn", m.GetSupportedOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the category property value. The group policy category associated with the definition.
// Parameters:
//  - value : Value to set for the category property.
func (m *GroupPolicyDefinition) SetCategory(value *GroupPolicyCategory)() {
    m.category = value
}
// Sets the categoryPath property value. The localized full category path for the policy.
// Parameters:
//  - value : Value to set for the categoryPath property.
func (m *GroupPolicyDefinition) SetCategoryPath(value *string)() {
    m.categoryPath = value
}
// Sets the classType property value. Identifies the type of groups the policy can be applied to. Possible values are: user, machine.
// Parameters:
//  - value : Value to set for the classType property.
func (m *GroupPolicyDefinition) SetClassType(value *GroupPolicyDefinitionClassType)() {
    m.classType = value
}
// Sets the definitionFile property value. The group policy file associated with the definition.
// Parameters:
//  - value : Value to set for the definitionFile property.
func (m *GroupPolicyDefinition) SetDefinitionFile(value *GroupPolicyDefinitionFile)() {
    m.definitionFile = value
}
// Sets the displayName property value. The localized policy name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GroupPolicyDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
// Parameters:
//  - value : Value to set for the explainText property.
func (m *GroupPolicyDefinition) SetExplainText(value *string)() {
    m.explainText = value
}
// Sets the groupPolicyCategoryId property value. The category id of the parent category
// Parameters:
//  - value : Value to set for the groupPolicyCategoryId property.
func (m *GroupPolicyDefinition) SetGroupPolicyCategoryId(value *string)() {
    m.groupPolicyCategoryId = value
}
// Sets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
// Parameters:
//  - value : Value to set for the hasRelatedDefinitions property.
func (m *GroupPolicyDefinition) SetHasRelatedDefinitions(value *bool)() {
    m.hasRelatedDefinitions = value
}
// Sets the lastModifiedDateTime property value. The date and time the entity was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *GroupPolicyDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
// Parameters:
//  - value : Value to set for the minDeviceCspVersion property.
func (m *GroupPolicyDefinition) SetMinDeviceCspVersion(value *string)() {
    m.minDeviceCspVersion = value
}
// Sets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
// Parameters:
//  - value : Value to set for the minUserCspVersion property.
func (m *GroupPolicyDefinition) SetMinUserCspVersion(value *string)() {
    m.minUserCspVersion = value
}
// Sets the nextVersionDefinition property value. Definition of the next version of this definition
// Parameters:
//  - value : Value to set for the nextVersionDefinition property.
func (m *GroupPolicyDefinition) SetNextVersionDefinition(value *GroupPolicyDefinition)() {
    m.nextVersionDefinition = value
}
// Sets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
// Parameters:
//  - value : Value to set for the policyType property.
func (m *GroupPolicyDefinition) SetPolicyType(value *GroupPolicyType)() {
    m.policyType = value
}
// Sets the presentations property value. The group policy presentations associated with the definition.
// Parameters:
//  - value : Value to set for the presentations property.
func (m *GroupPolicyDefinition) SetPresentations(value []GroupPolicyPresentation)() {
    m.presentations = value
}
// Sets the previousVersionDefinition property value. Definition of the previous version of this definition
// Parameters:
//  - value : Value to set for the previousVersionDefinition property.
func (m *GroupPolicyDefinition) SetPreviousVersionDefinition(value *GroupPolicyDefinition)() {
    m.previousVersionDefinition = value
}
// Sets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
// Parameters:
//  - value : Value to set for the supportedOn property.
func (m *GroupPolicyDefinition) SetSupportedOn(value *string)() {
    m.supportedOn = value
}
// Sets the version property value. Setting definition version
// Parameters:
//  - value : Value to set for the version property.
func (m *GroupPolicyDefinition) SetVersion(value *string)() {
    m.version = value
}
