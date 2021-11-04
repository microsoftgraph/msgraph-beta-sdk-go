package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationSettingDefinition struct {
    Entity
    // Read/write access mode of the setting. Possible values are: none, add, copy, delete, get, replace, execute.
    accessTypes *DeviceManagementConfigurationSettingAccessTypes;
    // Details which device setting is applicable on
    applicability *DeviceManagementConfigurationSettingApplicability;
    // Base CSP Path
    baseUri *string;
    // Specifies the area group under which the setting is configured in a specified configuration service provider (CSP)
    categoryId *string;
    // Description of the item
    description *string;
    // Display name of the item
    displayName *string;
    // Help text of the item
    helpText *string;
    // List of links more info for the setting can be found at
    infoUrls []string;
    // Tokens which to search settings on
    keywords []string;
    // Name of the item
    name *string;
    // Indicates whether the setting is required or not
    occurrence *DeviceManagementConfigurationSettingOccurrence;
    // Offset CSP Path from Base
    offsetUri *string;
    // List of referred setting information.
    referredSettingInformationList []DeviceManagementConfigurationReferredSettingInformation;
    // Root setting definition if the setting is a child setting.
    rootDefinitionId *string;
    // Setting type, for example, configuration and compliance. Possible values are: none, configuration.
    settingUsage *DeviceManagementConfigurationSettingUsage;
    // Setting control type representation in the UX. Possible values are: default, dropdown, smallTextBox, largeTextBox, toggle, multiheaderGrid, contextPane.
    uxBehavior *DeviceManagementConfigurationControlType;
    // Item Version
    version *string;
    // Setting visibility scope to UX. Possible values are: none, settingsCatalog, template.
    visibility *DeviceManagementConfigurationSettingVisibility;
}
// Instantiates a new deviceManagementConfigurationSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationSettingDefinition()(*DeviceManagementConfigurationSettingDefinition) {
    m := &DeviceManagementConfigurationSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessTypes property value. Read/write access mode of the setting. Possible values are: none, add, copy, delete, get, replace, execute.
func (m *DeviceManagementConfigurationSettingDefinition) GetAccessTypes()(*DeviceManagementConfigurationSettingAccessTypes) {
    if m == nil {
        return nil
    } else {
        return m.accessTypes
    }
}
// Gets the applicability property value. Details which device setting is applicable on
func (m *DeviceManagementConfigurationSettingDefinition) GetApplicability()(*DeviceManagementConfigurationSettingApplicability) {
    if m == nil {
        return nil
    } else {
        return m.applicability
    }
}
// Gets the baseUri property value. Base CSP Path
func (m *DeviceManagementConfigurationSettingDefinition) GetBaseUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseUri
    }
}
// Gets the categoryId property value. Specifies the area group under which the setting is configured in a specified configuration service provider (CSP)
func (m *DeviceManagementConfigurationSettingDefinition) GetCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.categoryId
    }
}
// Gets the description property value. Description of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the helpText property value. Help text of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetHelpText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpText
    }
}
// Gets the infoUrls property value. List of links more info for the setting can be found at
func (m *DeviceManagementConfigurationSettingDefinition) GetInfoUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.infoUrls
    }
}
// Gets the keywords property value. Tokens which to search settings on
func (m *DeviceManagementConfigurationSettingDefinition) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// Gets the name property value. Name of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the occurrence property value. Indicates whether the setting is required or not
func (m *DeviceManagementConfigurationSettingDefinition) GetOccurrence()(*DeviceManagementConfigurationSettingOccurrence) {
    if m == nil {
        return nil
    } else {
        return m.occurrence
    }
}
// Gets the offsetUri property value. Offset CSP Path from Base
func (m *DeviceManagementConfigurationSettingDefinition) GetOffsetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offsetUri
    }
}
// Gets the referredSettingInformationList property value. List of referred setting information.
func (m *DeviceManagementConfigurationSettingDefinition) GetReferredSettingInformationList()([]DeviceManagementConfigurationReferredSettingInformation) {
    if m == nil {
        return nil
    } else {
        return m.referredSettingInformationList
    }
}
// Gets the rootDefinitionId property value. Root setting definition if the setting is a child setting.
func (m *DeviceManagementConfigurationSettingDefinition) GetRootDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootDefinitionId
    }
}
// Gets the settingUsage property value. Setting type, for example, configuration and compliance. Possible values are: none, configuration.
func (m *DeviceManagementConfigurationSettingDefinition) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    if m == nil {
        return nil
    } else {
        return m.settingUsage
    }
}
// Gets the uxBehavior property value. Setting control type representation in the UX. Possible values are: default, dropdown, smallTextBox, largeTextBox, toggle, multiheaderGrid, contextPane.
func (m *DeviceManagementConfigurationSettingDefinition) GetUxBehavior()(*DeviceManagementConfigurationControlType) {
    if m == nil {
        return nil
    } else {
        return m.uxBehavior
    }
}
// Gets the version property value. Item Version
func (m *DeviceManagementConfigurationSettingDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Gets the visibility property value. Setting visibility scope to UX. Possible values are: none, settingsCatalog, template.
func (m *DeviceManagementConfigurationSettingDefinition) GetVisibility()(*DeviceManagementConfigurationSettingVisibility) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// The deserialization information for the current model
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
            res[i] = *(v.(*string))
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
            res[i] = *(v.(*string))
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accessTypes property value. Read/write access mode of the setting. Possible values are: none, add, copy, delete, get, replace, execute.
// Parameters:
//  - value : Value to set for the accessTypes property.
func (m *DeviceManagementConfigurationSettingDefinition) SetAccessTypes(value *DeviceManagementConfigurationSettingAccessTypes)() {
    m.accessTypes = value
}
// Sets the applicability property value. Details which device setting is applicable on
// Parameters:
//  - value : Value to set for the applicability property.
func (m *DeviceManagementConfigurationSettingDefinition) SetApplicability(value *DeviceManagementConfigurationSettingApplicability)() {
    m.applicability = value
}
// Sets the baseUri property value. Base CSP Path
// Parameters:
//  - value : Value to set for the baseUri property.
func (m *DeviceManagementConfigurationSettingDefinition) SetBaseUri(value *string)() {
    m.baseUri = value
}
// Sets the categoryId property value. Specifies the area group under which the setting is configured in a specified configuration service provider (CSP)
// Parameters:
//  - value : Value to set for the categoryId property.
func (m *DeviceManagementConfigurationSettingDefinition) SetCategoryId(value *string)() {
    m.categoryId = value
}
// Sets the description property value. Description of the item
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementConfigurationSettingDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name of the item
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementConfigurationSettingDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the helpText property value. Help text of the item
// Parameters:
//  - value : Value to set for the helpText property.
func (m *DeviceManagementConfigurationSettingDefinition) SetHelpText(value *string)() {
    m.helpText = value
}
// Sets the infoUrls property value. List of links more info for the setting can be found at
// Parameters:
//  - value : Value to set for the infoUrls property.
func (m *DeviceManagementConfigurationSettingDefinition) SetInfoUrls(value []string)() {
    m.infoUrls = value
}
// Sets the keywords property value. Tokens which to search settings on
// Parameters:
//  - value : Value to set for the keywords property.
func (m *DeviceManagementConfigurationSettingDefinition) SetKeywords(value []string)() {
    m.keywords = value
}
// Sets the name property value. Name of the item
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementConfigurationSettingDefinition) SetName(value *string)() {
    m.name = value
}
// Sets the occurrence property value. Indicates whether the setting is required or not
// Parameters:
//  - value : Value to set for the occurrence property.
func (m *DeviceManagementConfigurationSettingDefinition) SetOccurrence(value *DeviceManagementConfigurationSettingOccurrence)() {
    m.occurrence = value
}
// Sets the offsetUri property value. Offset CSP Path from Base
// Parameters:
//  - value : Value to set for the offsetUri property.
func (m *DeviceManagementConfigurationSettingDefinition) SetOffsetUri(value *string)() {
    m.offsetUri = value
}
// Sets the referredSettingInformationList property value. List of referred setting information.
// Parameters:
//  - value : Value to set for the referredSettingInformationList property.
func (m *DeviceManagementConfigurationSettingDefinition) SetReferredSettingInformationList(value []DeviceManagementConfigurationReferredSettingInformation)() {
    m.referredSettingInformationList = value
}
// Sets the rootDefinitionId property value. Root setting definition if the setting is a child setting.
// Parameters:
//  - value : Value to set for the rootDefinitionId property.
func (m *DeviceManagementConfigurationSettingDefinition) SetRootDefinitionId(value *string)() {
    m.rootDefinitionId = value
}
// Sets the settingUsage property value. Setting type, for example, configuration and compliance. Possible values are: none, configuration.
// Parameters:
//  - value : Value to set for the settingUsage property.
func (m *DeviceManagementConfigurationSettingDefinition) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    m.settingUsage = value
}
// Sets the uxBehavior property value. Setting control type representation in the UX. Possible values are: default, dropdown, smallTextBox, largeTextBox, toggle, multiheaderGrid, contextPane.
// Parameters:
//  - value : Value to set for the uxBehavior property.
func (m *DeviceManagementConfigurationSettingDefinition) SetUxBehavior(value *DeviceManagementConfigurationControlType)() {
    m.uxBehavior = value
}
// Sets the version property value. Item Version
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceManagementConfigurationSettingDefinition) SetVersion(value *string)() {
    m.version = value
}
// Sets the visibility property value. Setting visibility scope to UX. Possible values are: none, settingsCatalog, template.
// Parameters:
//  - value : Value to set for the visibility property.
func (m *DeviceManagementConfigurationSettingDefinition) SetVisibility(value *DeviceManagementConfigurationSettingVisibility)() {
    m.visibility = value
}
