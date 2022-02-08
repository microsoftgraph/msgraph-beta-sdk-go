package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingDefinition 
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
    // Setting type, for example, configuration and compliance. Possible values are: none, configuration, compliance.
    settingUsage *DeviceManagementConfigurationSettingUsage;
    // Setting control type representation in the UX. Possible values are: default, dropdown, smallTextBox, largeTextBox, toggle, multiheaderGrid, contextPane.
    uxBehavior *DeviceManagementConfigurationControlType;
    // Item Version
    version *string;
    // Setting visibility scope to UX. Possible values are: none, settingsCatalog, template.
    visibility *DeviceManagementConfigurationSettingVisibility;
}
// NewDeviceManagementConfigurationSettingDefinition instantiates a new deviceManagementConfigurationSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationSettingDefinition()(*DeviceManagementConfigurationSettingDefinition) {
    m := &DeviceManagementConfigurationSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessTypes gets the accessTypes property value. Read/write access mode of the setting. Possible values are: none, add, copy, delete, get, replace, execute.
func (m *DeviceManagementConfigurationSettingDefinition) GetAccessTypes()(*DeviceManagementConfigurationSettingAccessTypes) {
    if m == nil {
        return nil
    } else {
        return m.accessTypes
    }
}
// GetApplicability gets the applicability property value. Details which device setting is applicable on
func (m *DeviceManagementConfigurationSettingDefinition) GetApplicability()(*DeviceManagementConfigurationSettingApplicability) {
    if m == nil {
        return nil
    } else {
        return m.applicability
    }
}
// GetBaseUri gets the baseUri property value. Base CSP Path
func (m *DeviceManagementConfigurationSettingDefinition) GetBaseUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseUri
    }
}
// GetCategoryId gets the categoryId property value. Specifies the area group under which the setting is configured in a specified configuration service provider (CSP)
func (m *DeviceManagementConfigurationSettingDefinition) GetCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.categoryId
    }
}
// GetDescription gets the description property value. Description of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHelpText gets the helpText property value. Help text of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetHelpText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpText
    }
}
// GetInfoUrls gets the infoUrls property value. List of links more info for the setting can be found at
func (m *DeviceManagementConfigurationSettingDefinition) GetInfoUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.infoUrls
    }
}
// GetKeywords gets the keywords property value. Tokens which to search settings on
func (m *DeviceManagementConfigurationSettingDefinition) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// GetName gets the name property value. Name of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOccurrence gets the occurrence property value. Indicates whether the setting is required or not
func (m *DeviceManagementConfigurationSettingDefinition) GetOccurrence()(*DeviceManagementConfigurationSettingOccurrence) {
    if m == nil {
        return nil
    } else {
        return m.occurrence
    }
}
// GetOffsetUri gets the offsetUri property value. Offset CSP Path from Base
func (m *DeviceManagementConfigurationSettingDefinition) GetOffsetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offsetUri
    }
}
// GetReferredSettingInformationList gets the referredSettingInformationList property value. List of referred setting information.
func (m *DeviceManagementConfigurationSettingDefinition) GetReferredSettingInformationList()([]DeviceManagementConfigurationReferredSettingInformation) {
    if m == nil {
        return nil
    } else {
        return m.referredSettingInformationList
    }
}
// GetRootDefinitionId gets the rootDefinitionId property value. Root setting definition if the setting is a child setting.
func (m *DeviceManagementConfigurationSettingDefinition) GetRootDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootDefinitionId
    }
}
// GetSettingUsage gets the settingUsage property value. Setting type, for example, configuration and compliance. Possible values are: none, configuration, compliance.
func (m *DeviceManagementConfigurationSettingDefinition) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    if m == nil {
        return nil
    } else {
        return m.settingUsage
    }
}
// GetUxBehavior gets the uxBehavior property value. Setting control type representation in the UX. Possible values are: default, dropdown, smallTextBox, largeTextBox, toggle, multiheaderGrid, contextPane.
func (m *DeviceManagementConfigurationSettingDefinition) GetUxBehavior()(*DeviceManagementConfigurationControlType) {
    if m == nil {
        return nil
    } else {
        return m.uxBehavior
    }
}
// GetVersion gets the version property value. Item Version
func (m *DeviceManagementConfigurationSettingDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetVisibility gets the visibility property value. Setting visibility scope to UX. Possible values are: none, settingsCatalog, template.
func (m *DeviceManagementConfigurationSettingDefinition) GetVisibility()(*DeviceManagementConfigurationSettingVisibility) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingAccessTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessTypes(val.(*DeviceManagementConfigurationSettingAccessTypes))
        }
        return nil
    }
    res["applicability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingApplicability() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicability(val.(*DeviceManagementConfigurationSettingApplicability))
        }
        return nil
    }
    res["baseUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUri(val)
        }
        return nil
    }
    res["categoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryId(val)
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
    res["helpText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpText(val)
        }
        return nil
    }
    res["infoUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInfoUrls(res)
        }
        return nil
    }
    res["keywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKeywords(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["occurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingOccurrence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrence(val.(*DeviceManagementConfigurationSettingOccurrence))
        }
        return nil
    }
    res["offsetUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffsetUri(val)
        }
        return nil
    }
    res["referredSettingInformationList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationReferredSettingInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationReferredSettingInformation, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationReferredSettingInformation))
            }
            m.SetReferredSettingInformationList(res)
        }
        return nil
    }
    res["rootDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootDefinitionId(val)
        }
        return nil
    }
    res["settingUsage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingUsage(val.(*DeviceManagementConfigurationSettingUsage))
        }
        return nil
    }
    res["uxBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationControlType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUxBehavior(val.(*DeviceManagementConfigurationControlType))
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
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingVisibility)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val.(*DeviceManagementConfigurationSettingVisibility))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessTypes() != nil {
        cast := (*m.GetAccessTypes()).String()
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
    if m.GetInfoUrls() != nil {
        err = writer.WriteCollectionOfStringValues("infoUrls", m.GetInfoUrls())
        if err != nil {
            return err
        }
    }
    if m.GetKeywords() != nil {
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
    if m.GetReferredSettingInformationList() != nil {
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
        cast := (*m.GetSettingUsage()).String()
        err = writer.WriteStringValue("settingUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUxBehavior() != nil {
        cast := (*m.GetUxBehavior()).String()
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
        cast := (*m.GetVisibility()).String()
        err = writer.WriteStringValue("visibility", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessTypes sets the accessTypes property value. Read/write access mode of the setting. Possible values are: none, add, copy, delete, get, replace, execute.
func (m *DeviceManagementConfigurationSettingDefinition) SetAccessTypes(value *DeviceManagementConfigurationSettingAccessTypes)() {
    if m != nil {
        m.accessTypes = value
    }
}
// SetApplicability sets the applicability property value. Details which device setting is applicable on
func (m *DeviceManagementConfigurationSettingDefinition) SetApplicability(value *DeviceManagementConfigurationSettingApplicability)() {
    if m != nil {
        m.applicability = value
    }
}
// SetBaseUri sets the baseUri property value. Base CSP Path
func (m *DeviceManagementConfigurationSettingDefinition) SetBaseUri(value *string)() {
    if m != nil {
        m.baseUri = value
    }
}
// SetCategoryId sets the categoryId property value. Specifies the area group under which the setting is configured in a specified configuration service provider (CSP)
func (m *DeviceManagementConfigurationSettingDefinition) SetCategoryId(value *string)() {
    if m != nil {
        m.categoryId = value
    }
}
// SetDescription sets the description property value. Description of the item
func (m *DeviceManagementConfigurationSettingDefinition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the item
func (m *DeviceManagementConfigurationSettingDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHelpText sets the helpText property value. Help text of the item
func (m *DeviceManagementConfigurationSettingDefinition) SetHelpText(value *string)() {
    if m != nil {
        m.helpText = value
    }
}
// SetInfoUrls sets the infoUrls property value. List of links more info for the setting can be found at
func (m *DeviceManagementConfigurationSettingDefinition) SetInfoUrls(value []string)() {
    if m != nil {
        m.infoUrls = value
    }
}
// SetKeywords sets the keywords property value. Tokens which to search settings on
func (m *DeviceManagementConfigurationSettingDefinition) SetKeywords(value []string)() {
    if m != nil {
        m.keywords = value
    }
}
// SetName sets the name property value. Name of the item
func (m *DeviceManagementConfigurationSettingDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOccurrence sets the occurrence property value. Indicates whether the setting is required or not
func (m *DeviceManagementConfigurationSettingDefinition) SetOccurrence(value *DeviceManagementConfigurationSettingOccurrence)() {
    if m != nil {
        m.occurrence = value
    }
}
// SetOffsetUri sets the offsetUri property value. Offset CSP Path from Base
func (m *DeviceManagementConfigurationSettingDefinition) SetOffsetUri(value *string)() {
    if m != nil {
        m.offsetUri = value
    }
}
// SetReferredSettingInformationList sets the referredSettingInformationList property value. List of referred setting information.
func (m *DeviceManagementConfigurationSettingDefinition) SetReferredSettingInformationList(value []DeviceManagementConfigurationReferredSettingInformation)() {
    if m != nil {
        m.referredSettingInformationList = value
    }
}
// SetRootDefinitionId sets the rootDefinitionId property value. Root setting definition if the setting is a child setting.
func (m *DeviceManagementConfigurationSettingDefinition) SetRootDefinitionId(value *string)() {
    if m != nil {
        m.rootDefinitionId = value
    }
}
// SetSettingUsage sets the settingUsage property value. Setting type, for example, configuration and compliance. Possible values are: none, configuration, compliance.
func (m *DeviceManagementConfigurationSettingDefinition) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    if m != nil {
        m.settingUsage = value
    }
}
// SetUxBehavior sets the uxBehavior property value. Setting control type representation in the UX. Possible values are: default, dropdown, smallTextBox, largeTextBox, toggle, multiheaderGrid, contextPane.
func (m *DeviceManagementConfigurationSettingDefinition) SetUxBehavior(value *DeviceManagementConfigurationControlType)() {
    if m != nil {
        m.uxBehavior = value
    }
}
// SetVersion sets the version property value. Item Version
func (m *DeviceManagementConfigurationSettingDefinition) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
// SetVisibility sets the visibility property value. Setting visibility scope to UX. Possible values are: none, settingsCatalog, template.
func (m *DeviceManagementConfigurationSettingDefinition) SetVisibility(value *DeviceManagementConfigurationSettingVisibility)() {
    if m != nil {
        m.visibility = value
    }
}
