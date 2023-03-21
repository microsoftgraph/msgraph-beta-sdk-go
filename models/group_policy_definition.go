package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinition the entity describes all of the information about a single group policy.
type GroupPolicyDefinition struct {
    Entity
}
// NewGroupPolicyDefinition instantiates a new groupPolicyDefinition and sets the default values.
func NewGroupPolicyDefinition()(*GroupPolicyDefinition) {
    m := &GroupPolicyDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyDefinition(), nil
}
// GetCategory gets the category property value. The group policy category associated with the definition.
func (m *GroupPolicyDefinition) GetCategory()(GroupPolicyCategoryable) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyCategoryable)
    }
    return nil
}
// GetCategoryPath gets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) GetCategoryPath()(*string) {
    val, err := m.GetBackingStore().Get("categoryPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClassType gets the classType property value. Group Policy Definition Class Type.
func (m *GroupPolicyDefinition) GetClassType()(*GroupPolicyDefinitionClassType) {
    val, err := m.GetBackingStore().Get("classType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicyDefinitionClassType)
    }
    return nil
}
// GetDefinitionFile gets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) GetDefinitionFile()(GroupPolicyDefinitionFileable) {
    val, err := m.GetBackingStore().Get("definitionFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionFileable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExplainText gets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) GetExplainText()(*string) {
    val, err := m.GetBackingStore().Get("explainText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(GroupPolicyCategoryable))
        }
        return nil
    }
    res["categoryPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryPath(val)
        }
        return nil
    }
    res["classType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyDefinitionClassType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassType(val.(*GroupPolicyDefinitionClassType))
        }
        return nil
    }
    res["definitionFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionFile(val.(GroupPolicyDefinitionFileable))
        }
        return nil
    }
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
    res["explainText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExplainText(val)
        }
        return nil
    }
    res["groupPolicyCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyCategoryId(val)
        }
        return nil
    }
    res["hasRelatedDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasRelatedDefinitions(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["minDeviceCspVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinDeviceCspVersion(val)
        }
        return nil
    }
    res["minUserCspVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinUserCspVersion(val)
        }
        return nil
    }
    res["nextVersionDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextVersionDefinition(val.(GroupPolicyDefinitionable))
        }
        return nil
    }
    res["policyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyType(val.(*GroupPolicyType))
        }
        return nil
    }
    res["presentations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyPresentationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyPresentationable)
            }
            m.SetPresentations(res)
        }
        return nil
    }
    res["previousVersionDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousVersionDefinition(val.(GroupPolicyDefinitionable))
        }
        return nil
    }
    res["supportedOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedOn(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetGroupPolicyCategoryId gets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) GetGroupPolicyCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("groupPolicyCategoryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetHasRelatedDefinitions gets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) GetHasRelatedDefinitions()(*bool) {
    val, err := m.GetBackingStore().Get("hasRelatedDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMinDeviceCspVersion gets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) GetMinDeviceCspVersion()(*string) {
    val, err := m.GetBackingStore().Get("minDeviceCspVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinUserCspVersion gets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) GetMinUserCspVersion()(*string) {
    val, err := m.GetBackingStore().Get("minUserCspVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNextVersionDefinition gets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) GetNextVersionDefinition()(GroupPolicyDefinitionable) {
    val, err := m.GetBackingStore().Get("nextVersionDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionable)
    }
    return nil
}
// GetPolicyType gets the policyType property value. Type of Group Policy File or Definition.
func (m *GroupPolicyDefinition) GetPolicyType()(*GroupPolicyType) {
    val, err := m.GetBackingStore().Get("policyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicyType)
    }
    return nil
}
// GetPresentations gets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) GetPresentations()([]GroupPolicyPresentationable) {
    val, err := m.GetBackingStore().Get("presentations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyPresentationable)
    }
    return nil
}
// GetPreviousVersionDefinition gets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) GetPreviousVersionDefinition()(GroupPolicyDefinitionable) {
    val, err := m.GetBackingStore().Get("previousVersionDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionable)
    }
    return nil
}
// GetSupportedOn gets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) GetSupportedOn()(*string) {
    val, err := m.GetBackingStore().Get("supportedOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVersion gets the version property value. Setting definition version
func (m *GroupPolicyDefinition) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteUUIDValue("groupPolicyCategoryId", m.GetGroupPolicyCategoryId())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPresentations()))
        for i, v := range m.GetPresentations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *GroupPolicyDefinition) SetCategory(value GroupPolicyCategoryable)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetCategoryPath sets the categoryPath property value. The localized full category path for the policy.
func (m *GroupPolicyDefinition) SetCategoryPath(value *string)() {
    err := m.GetBackingStore().Set("categoryPath", value)
    if err != nil {
        panic(err)
    }
}
// SetClassType sets the classType property value. Group Policy Definition Class Type.
func (m *GroupPolicyDefinition) SetClassType(value *GroupPolicyDefinitionClassType)() {
    err := m.GetBackingStore().Set("classType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitionFile sets the definitionFile property value. The group policy file associated with the definition.
func (m *GroupPolicyDefinition) SetDefinitionFile(value GroupPolicyDefinitionFileable)() {
    err := m.GetBackingStore().Set("definitionFile", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The localized policy name.
func (m *GroupPolicyDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExplainText sets the explainText property value. The localized explanation or help text associated with the policy. The default value is empty.
func (m *GroupPolicyDefinition) SetExplainText(value *string)() {
    err := m.GetBackingStore().Set("explainText", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyCategoryId sets the groupPolicyCategoryId property value. The category id of the parent category
func (m *GroupPolicyDefinition) SetGroupPolicyCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("groupPolicyCategoryId", value)
    if err != nil {
        panic(err)
    }
}
// SetHasRelatedDefinitions sets the hasRelatedDefinitions property value. Signifies whether or not there are related definitions to this definition
func (m *GroupPolicyDefinition) SetHasRelatedDefinitions(value *bool)() {
    err := m.GetBackingStore().Set("hasRelatedDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMinDeviceCspVersion sets the minDeviceCspVersion property value. Minimum required CSP version for device configuration in this definition
func (m *GroupPolicyDefinition) SetMinDeviceCspVersion(value *string)() {
    err := m.GetBackingStore().Set("minDeviceCspVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinUserCspVersion sets the minUserCspVersion property value. Minimum required CSP version for user configuration in this definition
func (m *GroupPolicyDefinition) SetMinUserCspVersion(value *string)() {
    err := m.GetBackingStore().Set("minUserCspVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetNextVersionDefinition sets the nextVersionDefinition property value. Definition of the next version of this definition
func (m *GroupPolicyDefinition) SetNextVersionDefinition(value GroupPolicyDefinitionable)() {
    err := m.GetBackingStore().Set("nextVersionDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyType sets the policyType property value. Type of Group Policy File or Definition.
func (m *GroupPolicyDefinition) SetPolicyType(value *GroupPolicyType)() {
    err := m.GetBackingStore().Set("policyType", value)
    if err != nil {
        panic(err)
    }
}
// SetPresentations sets the presentations property value. The group policy presentations associated with the definition.
func (m *GroupPolicyDefinition) SetPresentations(value []GroupPolicyPresentationable)() {
    err := m.GetBackingStore().Set("presentations", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviousVersionDefinition sets the previousVersionDefinition property value. Definition of the previous version of this definition
func (m *GroupPolicyDefinition) SetPreviousVersionDefinition(value GroupPolicyDefinitionable)() {
    err := m.GetBackingStore().Set("previousVersionDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedOn sets the supportedOn property value. Localized string used to specify what operating system or application version is affected by the policy.
func (m *GroupPolicyDefinition) SetSupportedOn(value *string)() {
    err := m.GetBackingStore().Set("supportedOn", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Setting definition version
func (m *GroupPolicyDefinition) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyDefinitionable 
type GroupPolicyDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(GroupPolicyCategoryable)
    GetCategoryPath()(*string)
    GetClassType()(*GroupPolicyDefinitionClassType)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDisplayName()(*string)
    GetExplainText()(*string)
    GetGroupPolicyCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetHasRelatedDefinitions()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMinDeviceCspVersion()(*string)
    GetMinUserCspVersion()(*string)
    GetNextVersionDefinition()(GroupPolicyDefinitionable)
    GetPolicyType()(*GroupPolicyType)
    GetPresentations()([]GroupPolicyPresentationable)
    GetPreviousVersionDefinition()(GroupPolicyDefinitionable)
    GetSupportedOn()(*string)
    GetVersion()(*string)
    SetCategory(value GroupPolicyCategoryable)()
    SetCategoryPath(value *string)()
    SetClassType(value *GroupPolicyDefinitionClassType)()
    SetDefinitionFile(value GroupPolicyDefinitionFileable)()
    SetDisplayName(value *string)()
    SetExplainText(value *string)()
    SetGroupPolicyCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetHasRelatedDefinitions(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMinDeviceCspVersion(value *string)()
    SetMinUserCspVersion(value *string)()
    SetNextVersionDefinition(value GroupPolicyDefinitionable)()
    SetPolicyType(value *GroupPolicyType)()
    SetPresentations(value []GroupPolicyPresentationable)()
    SetPreviousVersionDefinition(value GroupPolicyDefinitionable)()
    SetSupportedOn(value *string)()
    SetVersion(value *string)()
}
