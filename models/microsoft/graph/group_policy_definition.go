package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyDefinition 
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
// NewGroupPolicyDefinition instantiates a new groupPolicyDefinition and sets the default values.
func NewGroupPolicyDefinition()(*GroupPolicyDefinition) {
    m := &GroupPolicyDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategory gets the category property value. The group policy category associated with the definition.
func (m *GroupPolicyDefinition) GetCategory()(*GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetCategoryPath gets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) GetCategoryPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.categoryPath
    }
}
// GetClassType gets the classType property value. Identifies the type of groups the policy can be applied to. Possible values are: user, machine.
func (m *GroupPolicyDefinition) GetClassType()(*GroupPolicyDefinitionClassType) {
    if m == nil {
        return nil
    } else {
        return m.classType
    }
}
// GetDefinitionFile gets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) GetDefinitionFile()(*GroupPolicyDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.definitionFile
    }
}
// GetDisplayName gets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExplainText gets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) GetExplainText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.explainText
    }
}
// GetGroupPolicyCategoryId gets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) GetGroupPolicyCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCategoryId
    }
}
// GetHasRelatedDefinitions gets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) GetHasRelatedDefinitions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasRelatedDefinitions
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetMinDeviceCspVersion gets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) GetMinDeviceCspVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minDeviceCspVersion
    }
}
// GetMinUserCspVersion gets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) GetMinUserCspVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minUserCspVersion
    }
}
// GetNextVersionDefinition gets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) GetNextVersionDefinition()(*GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.nextVersionDefinition
    }
}
// GetPolicyType gets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
func (m *GroupPolicyDefinition) GetPolicyType()(*GroupPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
// GetPresentations gets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) GetPresentations()([]GroupPolicyPresentation) {
    if m == nil {
        return nil
    } else {
        return m.presentations
    }
}
// GetPreviousVersionDefinition gets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) GetPreviousVersionDefinition()(*GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.previousVersionDefinition
    }
}
// GetSupportedOn gets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) GetSupportedOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportedOn
    }
}
// GetVersion gets the version property value. Setting definition version
func (m *GroupPolicyDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetClassType(val.(*GroupPolicyDefinitionClassType))
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
            m.SetPolicyType(val.(*GroupPolicyType))
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
// Serialize serializes information the current object
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
        cast := (*m.GetClassType()).String()
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
        cast := (*m.GetPolicyType()).String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPresentations() != nil {
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
// SetCategory sets the category property value. The group policy category associated with the definition.
func (m *GroupPolicyDefinition) SetCategory(value *GroupPolicyCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetCategoryPath sets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) SetCategoryPath(value *string)() {
    if m != nil {
        m.categoryPath = value
    }
}
// SetClassType sets the classType property value. Identifies the type of groups the policy can be applied to. Possible values are: user, machine.
func (m *GroupPolicyDefinition) SetClassType(value *GroupPolicyDefinitionClassType)() {
    if m != nil {
        m.classType = value
    }
}
// SetDefinitionFile sets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) SetDefinitionFile(value *GroupPolicyDefinitionFile)() {
    if m != nil {
        m.definitionFile = value
    }
}
// SetDisplayName sets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExplainText sets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) SetExplainText(value *string)() {
    if m != nil {
        m.explainText = value
    }
}
// SetGroupPolicyCategoryId sets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) SetGroupPolicyCategoryId(value *string)() {
    if m != nil {
        m.groupPolicyCategoryId = value
    }
}
// SetHasRelatedDefinitions sets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) SetHasRelatedDefinitions(value *bool)() {
    if m != nil {
        m.hasRelatedDefinitions = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetMinDeviceCspVersion sets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) SetMinDeviceCspVersion(value *string)() {
    if m != nil {
        m.minDeviceCspVersion = value
    }
}
// SetMinUserCspVersion sets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) SetMinUserCspVersion(value *string)() {
    if m != nil {
        m.minUserCspVersion = value
    }
}
// SetNextVersionDefinition sets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) SetNextVersionDefinition(value *GroupPolicyDefinition)() {
    if m != nil {
        m.nextVersionDefinition = value
    }
}
// SetPolicyType sets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
func (m *GroupPolicyDefinition) SetPolicyType(value *GroupPolicyType)() {
    if m != nil {
        m.policyType = value
    }
}
// SetPresentations sets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) SetPresentations(value []GroupPolicyPresentation)() {
    if m != nil {
        m.presentations = value
    }
}
// SetPreviousVersionDefinition sets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) SetPreviousVersionDefinition(value *GroupPolicyDefinition)() {
    if m != nil {
        m.previousVersionDefinition = value
    }
}
// SetSupportedOn sets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) SetSupportedOn(value *string)() {
    if m != nil {
        m.supportedOn = value
    }
}
// SetVersion sets the version property value. Setting definition version
func (m *GroupPolicyDefinition) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
