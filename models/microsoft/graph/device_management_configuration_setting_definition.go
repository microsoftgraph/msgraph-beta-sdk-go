package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSettingDefinition struct {
    Entity
    accessTypes *DeviceManagementConfigurationSettingAccessTypes;
    applicability *DeviceManagementConfigurationSettingApplicability;
    baseUri *string;
    categoryId *string;
    description *string;
    displayName *string;
    helpText *string;
    infoUrls []string;
    keywords []string;
    name *string;
    occurrence *DeviceManagementConfigurationSettingOccurrence;
    offsetUri *string;
    referredSettingInformationList []DeviceManagementConfigurationReferredSettingInformation;
    rootDefinitionId *string;
    settingUsage *DeviceManagementConfigurationSettingUsage;
    uxBehavior *DeviceManagementConfigurationControlType;
    version *string;
    visibility *DeviceManagementConfigurationSettingVisibility;
}
func NewDeviceManagementConfigurationSettingDefinition()(*DeviceManagementConfigurationSettingDefinition) {
    m := &DeviceManagementConfigurationSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementConfigurationSettingDefinition) GetAccessTypes()(*DeviceManagementConfigurationSettingAccessTypes) {
    if m == nil {
        return nil
    } else {
        return m.accessTypes
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetApplicability()(*DeviceManagementConfigurationSettingApplicability) {
    if m == nil {
        return nil
    } else {
        return m.applicability
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetBaseUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseUri
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.categoryId
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetHelpText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpText
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetInfoUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.infoUrls
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetOccurrence()(*DeviceManagementConfigurationSettingOccurrence) {
    if m == nil {
        return nil
    } else {
        return m.occurrence
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetOffsetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offsetUri
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetReferredSettingInformationList()([]DeviceManagementConfigurationReferredSettingInformation) {
    if m == nil {
        return nil
    } else {
        return m.referredSettingInformationList
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetRootDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootDefinitionId
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    if m == nil {
        return nil
    } else {
        return m.settingUsage
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetUxBehavior()(*DeviceManagementConfigurationControlType) {
    if m == nil {
        return nil
    } else {
        return m.uxBehavior
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetVisibility()(*DeviceManagementConfigurationSettingVisibility) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
func (m *DeviceManagementConfigurationSettingDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingAccessTypes)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationSettingAccessTypes)
        m.SetAccessTypes(&cast)
        return nil
    }
    res["applicability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingApplicability() })
        if err != nil {
            return err
        }
        m.SetApplicability(val.(*DeviceManagementConfigurationSettingApplicability))
        return nil
    }
    res["baseUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBaseUri(val)
        return nil
    }
    res["categoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategoryId(val)
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
    res["helpText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHelpText(val)
        return nil
    }
    res["infoUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetInfoUrls(res)
        return nil
    }
    res["keywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetKeywords(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["occurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingOccurrence() })
        if err != nil {
            return err
        }
        m.SetOccurrence(val.(*DeviceManagementConfigurationSettingOccurrence))
        return nil
    }
    res["offsetUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOffsetUri(val)
        return nil
    }
    res["referredSettingInformationList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationReferredSettingInformation() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationReferredSettingInformation, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationReferredSettingInformation))
        }
        m.SetReferredSettingInformationList(res)
        return nil
    }
    res["rootDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRootDefinitionId(val)
        return nil
    }
    res["settingUsage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingUsage)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationSettingUsage)
        m.SetSettingUsage(&cast)
        return nil
    }
    res["uxBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationControlType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationControlType)
        m.SetUxBehavior(&cast)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingVisibility)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationSettingVisibility)
        m.SetVisibility(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingDefinition) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSettingDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessTypes() != nil {
        cast := m.GetAccessTypes().String()
        err = writer.WriteStringValue("accessTypes", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("applicability", m.GetApplicability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("baseUri", m.GetBaseUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("categoryId", m.GetCategoryId())
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
        err = writer.WriteStringValue("helpText", m.GetHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("infoUrls", m.GetInfoUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("keywords", m.GetKeywords())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("occurrence", m.GetOccurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("offsetUri", m.GetOffsetUri())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReferredSettingInformationList()))
        for i, v := range m.GetReferredSettingInformationList() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("referredSettingInformationList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rootDefinitionId", m.GetRootDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetSettingUsage() != nil {
        cast := m.GetSettingUsage().String()
        err = writer.WriteStringValue("settingUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUxBehavior() != nil {
        cast := m.GetUxBehavior().String()
        err = writer.WriteStringValue("uxBehavior", &cast)
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
    if m.GetVisibility() != nil {
        cast := m.GetVisibility().String()
        err = writer.WriteStringValue("visibility", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementConfigurationSettingDefinition) SetAccessTypes(value *DeviceManagementConfigurationSettingAccessTypes)() {
    m.accessTypes = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetApplicability(value *DeviceManagementConfigurationSettingApplicability)() {
    m.applicability = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetBaseUri(value *string)() {
    m.baseUri = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetCategoryId(value *string)() {
    m.categoryId = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetHelpText(value *string)() {
    m.helpText = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetInfoUrls(value []string)() {
    m.infoUrls = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetKeywords(value []string)() {
    m.keywords = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetName(value *string)() {
    m.name = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetOccurrence(value *DeviceManagementConfigurationSettingOccurrence)() {
    m.occurrence = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetOffsetUri(value *string)() {
    m.offsetUri = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetReferredSettingInformationList(value []DeviceManagementConfigurationReferredSettingInformation)() {
    m.referredSettingInformationList = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetRootDefinitionId(value *string)() {
    m.rootDefinitionId = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    m.settingUsage = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetUxBehavior(value *DeviceManagementConfigurationControlType)() {
    m.uxBehavior = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetVersion(value *string)() {
    m.version = value
}
func (m *DeviceManagementConfigurationSettingDefinition) SetVisibility(value *DeviceManagementConfigurationSettingVisibility)() {
    m.visibility = value
}
