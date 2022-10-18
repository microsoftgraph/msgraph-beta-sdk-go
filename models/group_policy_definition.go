package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinition the entity describes all of the information about a single group policy.
type GroupPolicyDefinition struct {
    Entity
    // The group policy category associated with the definition.
    category GroupPolicyCategoryable
    // The localized full category path for the policy.
    categoryPath *string
    // Group Policy Definition Class Type.
    classType *GroupPolicyDefinitionClassType
    // The group policy file associated with the definition.
    definitionFile GroupPolicyDefinitionFileable
    // The localized policy name.
    displayName *string
    // The localized explanation or help text associated with the policy. The default value is empty.
    explainText *string
    // The category id of the parent category
    groupPolicyCategoryId *string
    // Signifies whether or not there are related definitions to this definition
    hasRelatedDefinitions *bool
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Minimum required CSP version for device configuration in this definition
    minDeviceCspVersion *string
    // Minimum required CSP version for user configuration in this definition
    minUserCspVersion *string
    // Definition of the next version of this definition
    nextVersionDefinition GroupPolicyDefinitionable
    // Type of Group Policy File or Definition.
    policyType *GroupPolicyType
    // The group policy presentations associated with the definition.
    presentations []GroupPolicyPresentationable
    // Definition of the previous version of this definition
    previousVersionDefinition GroupPolicyDefinitionable
    // Localized string used to specify what operating system or application version is affected by the policy.
    supportedOn *string
    // Setting definition version
    version *string
}
// NewGroupPolicyDefinition instantiates a new groupPolicyDefinition and sets the default values.
func NewGroupPolicyDefinition()(*GroupPolicyDefinition) {
    m := &GroupPolicyDefinition{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyDefinition(), nil
}
// GetCategory gets the category property value. The group policy category associated with the definition.
func (m *GroupPolicyDefinition) GetCategory()(GroupPolicyCategoryable) {
    return m.category
}
// GetCategoryPath gets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) GetCategoryPath()(*string) {
    return m.categoryPath
}
// GetClassType gets the classType property value. Group Policy Definition Class Type.
func (m *GroupPolicyDefinition) GetClassType()(*GroupPolicyDefinitionClassType) {
    return m.classType
}
// GetDefinitionFile gets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) GetDefinitionFile()(GroupPolicyDefinitionFileable) {
    return m.definitionFile
}
// GetDisplayName gets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetExplainText gets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) GetExplainText()(*string) {
    return m.explainText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyCategoryFromDiscriminatorValue , m.SetCategory)
    res["categoryPath"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCategoryPath)
    res["classType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseGroupPolicyDefinitionClassType , m.SetClassType)
    res["definitionFile"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyDefinitionFileFromDiscriminatorValue , m.SetDefinitionFile)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["explainText"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExplainText)
    res["groupPolicyCategoryId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupPolicyCategoryId)
    res["hasRelatedDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHasRelatedDefinitions)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["minDeviceCspVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinDeviceCspVersion)
    res["minUserCspVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinUserCspVersion)
    res["nextVersionDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue , m.SetNextVersionDefinition)
    res["policyType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseGroupPolicyType , m.SetPolicyType)
    res["presentations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyPresentationFromDiscriminatorValue , m.SetPresentations)
    res["previousVersionDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue , m.SetPreviousVersionDefinition)
    res["supportedOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSupportedOn)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersion)
    return res
}
// GetGroupPolicyCategoryId gets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) GetGroupPolicyCategoryId()(*string) {
    return m.groupPolicyCategoryId
}
// GetHasRelatedDefinitions gets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) GetHasRelatedDefinitions()(*bool) {
    return m.hasRelatedDefinitions
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMinDeviceCspVersion gets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) GetMinDeviceCspVersion()(*string) {
    return m.minDeviceCspVersion
}
// GetMinUserCspVersion gets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) GetMinUserCspVersion()(*string) {
    return m.minUserCspVersion
}
// GetNextVersionDefinition gets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) GetNextVersionDefinition()(GroupPolicyDefinitionable) {
    return m.nextVersionDefinition
}
// GetPolicyType gets the policyType property value. Type of Group Policy File or Definition.
func (m *GroupPolicyDefinition) GetPolicyType()(*GroupPolicyType) {
    return m.policyType
}
// GetPresentations gets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) GetPresentations()([]GroupPolicyPresentationable) {
    return m.presentations
}
// GetPreviousVersionDefinition gets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) GetPreviousVersionDefinition()(GroupPolicyDefinitionable) {
    return m.previousVersionDefinition
}
// GetSupportedOn gets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) GetSupportedOn()(*string) {
    return m.supportedOn
}
// GetVersion gets the version property value. Setting definition version
func (m *GroupPolicyDefinition) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *GroupPolicyDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPresentations())
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
func (m *GroupPolicyDefinition) SetCategory(value GroupPolicyCategoryable)() {
    m.category = value
}
// SetCategoryPath sets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) SetCategoryPath(value *string)() {
    m.categoryPath = value
}
// SetClassType sets the classType property value. Group Policy Definition Class Type.
func (m *GroupPolicyDefinition) SetClassType(value *GroupPolicyDefinitionClassType)() {
    m.classType = value
}
// SetDefinitionFile sets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) SetDefinitionFile(value GroupPolicyDefinitionFileable)() {
    m.definitionFile = value
}
// SetDisplayName sets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExplainText sets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) SetExplainText(value *string)() {
    m.explainText = value
}
// SetGroupPolicyCategoryId sets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) SetGroupPolicyCategoryId(value *string)() {
    m.groupPolicyCategoryId = value
}
// SetHasRelatedDefinitions sets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) SetHasRelatedDefinitions(value *bool)() {
    m.hasRelatedDefinitions = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMinDeviceCspVersion sets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) SetMinDeviceCspVersion(value *string)() {
    m.minDeviceCspVersion = value
}
// SetMinUserCspVersion sets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) SetMinUserCspVersion(value *string)() {
    m.minUserCspVersion = value
}
// SetNextVersionDefinition sets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) SetNextVersionDefinition(value GroupPolicyDefinitionable)() {
    m.nextVersionDefinition = value
}
// SetPolicyType sets the policyType property value. Type of Group Policy File or Definition.
func (m *GroupPolicyDefinition) SetPolicyType(value *GroupPolicyType)() {
    m.policyType = value
}
// SetPresentations sets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) SetPresentations(value []GroupPolicyPresentationable)() {
    m.presentations = value
}
// SetPreviousVersionDefinition sets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) SetPreviousVersionDefinition(value GroupPolicyDefinitionable)() {
    m.previousVersionDefinition = value
}
// SetSupportedOn sets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) SetSupportedOn(value *string)() {
    m.supportedOn = value
}
// SetVersion sets the version property value. Setting definition version
func (m *GroupPolicyDefinition) SetVersion(value *string)() {
    m.version = value
}
