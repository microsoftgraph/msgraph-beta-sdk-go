package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationCategory struct {
    Entity
    // List of child ids of the category.
    childCategoryIds []string;
    // Description of the item
    description *string;
    // Display name of the item
    displayName *string;
    // Help text of the item
    helpText *string;
    // Name of the item
    name *string;
    // Parent id of the category.
    parentCategoryId *string;
    // Platforms types, which settings in the category have. Possible values are: none, android, iOS, macOS, windows10X, windows10.
    platforms *DeviceManagementConfigurationPlatforms;
    // Root id of the category.
    rootCategoryId *string;
    // Indicates that the category contains settings that are used for Compliance or Configuration. Possible values are: none, configuration.
    settingUsage *DeviceManagementConfigurationSettingUsage;
    // Technologies types, which settings in the category have. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
    technologies *DeviceManagementConfigurationTechnologies;
}
// Instantiates a new deviceManagementConfigurationCategory and sets the default values.
func NewDeviceManagementConfigurationCategory()(*DeviceManagementConfigurationCategory) {
    m := &DeviceManagementConfigurationCategory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the childCategoryIds property value. List of child ids of the category.
func (m *DeviceManagementConfigurationCategory) GetChildCategoryIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.childCategoryIds
    }
}
// Gets the description property value. Description of the item
func (m *DeviceManagementConfigurationCategory) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name of the item
func (m *DeviceManagementConfigurationCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the helpText property value. Help text of the item
func (m *DeviceManagementConfigurationCategory) GetHelpText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpText
    }
}
// Gets the name property value. Name of the item
func (m *DeviceManagementConfigurationCategory) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the parentCategoryId property value. Parent id of the category.
func (m *DeviceManagementConfigurationCategory) GetParentCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentCategoryId
    }
}
// Gets the platforms property value. Platforms types, which settings in the category have. Possible values are: none, android, iOS, macOS, windows10X, windows10.
func (m *DeviceManagementConfigurationCategory) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// Gets the rootCategoryId property value. Root id of the category.
func (m *DeviceManagementConfigurationCategory) GetRootCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootCategoryId
    }
}
// Gets the settingUsage property value. Indicates that the category contains settings that are used for Compliance or Configuration. Possible values are: none, configuration.
func (m *DeviceManagementConfigurationCategory) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    if m == nil {
        return nil
    } else {
        return m.settingUsage
    }
}
// Gets the technologies property value. Technologies types, which settings in the category have. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
func (m *DeviceManagementConfigurationCategory) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    if m == nil {
        return nil
    } else {
        return m.technologies
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childCategoryIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetChildCategoryIds(res)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["parentCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentCategoryId(val)
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationPlatforms)
        m.SetPlatforms(&cast)
        return nil
    }
    res["rootCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRootCategoryId(val)
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
    res["technologies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationTechnologies)
        m.SetTechnologies(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationCategory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("childCategoryIds", m.GetChildCategoryIds())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentCategoryId", m.GetParentCategoryId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatforms() != nil {
        cast := m.GetPlatforms().String()
        err = writer.WriteStringValue("platforms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rootCategoryId", m.GetRootCategoryId())
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
    if m.GetTechnologies() != nil {
        cast := m.GetTechnologies().String()
        err = writer.WriteStringValue("technologies", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the childCategoryIds property value. List of child ids of the category.
// Parameters:
//  - value : Value to set for the childCategoryIds property.
func (m *DeviceManagementConfigurationCategory) SetChildCategoryIds(value []string)() {
    m.childCategoryIds = value
}
// Sets the description property value. Description of the item
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementConfigurationCategory) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name of the item
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementConfigurationCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the helpText property value. Help text of the item
// Parameters:
//  - value : Value to set for the helpText property.
func (m *DeviceManagementConfigurationCategory) SetHelpText(value *string)() {
    m.helpText = value
}
// Sets the name property value. Name of the item
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementConfigurationCategory) SetName(value *string)() {
    m.name = value
}
// Sets the parentCategoryId property value. Parent id of the category.
// Parameters:
//  - value : Value to set for the parentCategoryId property.
func (m *DeviceManagementConfigurationCategory) SetParentCategoryId(value *string)() {
    m.parentCategoryId = value
}
// Sets the platforms property value. Platforms types, which settings in the category have. Possible values are: none, android, iOS, macOS, windows10X, windows10.
// Parameters:
//  - value : Value to set for the platforms property.
func (m *DeviceManagementConfigurationCategory) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// Sets the rootCategoryId property value. Root id of the category.
// Parameters:
//  - value : Value to set for the rootCategoryId property.
func (m *DeviceManagementConfigurationCategory) SetRootCategoryId(value *string)() {
    m.rootCategoryId = value
}
// Sets the settingUsage property value. Indicates that the category contains settings that are used for Compliance or Configuration. Possible values are: none, configuration.
// Parameters:
//  - value : Value to set for the settingUsage property.
func (m *DeviceManagementConfigurationCategory) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    m.settingUsage = value
}
// Sets the technologies property value. Technologies types, which settings in the category have. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
// Parameters:
//  - value : Value to set for the technologies property.
func (m *DeviceManagementConfigurationCategory) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
