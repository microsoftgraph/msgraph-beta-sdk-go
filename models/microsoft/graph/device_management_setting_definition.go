package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementSettingDefinition struct {
    Entity
    // Collection of constraints for the setting value
    constraints []DeviceManagementConstraint;
    // Collection of dependencies on other settings
    dependencies []DeviceManagementSettingDependency;
    // The setting's description
    description *string;
    // The setting's display name
    displayName *string;
    // Url to setting documentation
    documentationUrl *string;
    // subtitle of the setting header for more details about the category/section
    headerSubtitle *string;
    // title of the setting header represents a category/section of a setting/settings
    headerTitle *string;
    // If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
    isTopLevel *bool;
    // Keywords associated with the setting
    keywords []string;
    // Placeholder text as an example of valid input
    placeholderText *string;
    // The data type of the value. Possible values are: integer, boolean, string, complex, collection, abstractComplex.
    valueType *DeviceManangementIntentValueType;
}
// Instantiates a new deviceManagementSettingDefinition and sets the default values.
func NewDeviceManagementSettingDefinition()(*DeviceManagementSettingDefinition) {
    m := &DeviceManagementSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the constraints property value. Collection of constraints for the setting value
func (m *DeviceManagementSettingDefinition) GetConstraints()([]DeviceManagementConstraint) {
    if m == nil {
        return nil
    } else {
        return m.constraints
    }
}
// Gets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) GetDependencies()([]DeviceManagementSettingDependency) {
    if m == nil {
        return nil
    } else {
        return m.dependencies
    }
}
// Gets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) GetDocumentationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentationUrl
    }
}
// Gets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) GetHeaderSubtitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerSubtitle
    }
}
// Gets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) GetHeaderTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerTitle
    }
}
// Gets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) GetIsTopLevel()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTopLevel
    }
}
// Gets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// Gets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) GetPlaceholderText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.placeholderText
    }
}
// Gets the valueType property value. The data type of the value. Possible values are: integer, boolean, string, complex, collection, abstractComplex.
func (m *DeviceManagementSettingDefinition) GetValueType()(*DeviceManangementIntentValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// The deserialization information for the current model
func (m *DeviceManagementSettingDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["constraints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConstraint() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConstraint, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConstraint))
        }
        m.SetConstraints(res)
        return nil
    }
    res["dependencies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingDependency() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementSettingDependency, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementSettingDependency))
        }
        m.SetDependencies(res)
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
    res["documentationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDocumentationUrl(val)
        return nil
    }
    res["headerSubtitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHeaderSubtitle(val)
        return nil
    }
    res["headerTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHeaderTitle(val)
        return nil
    }
    res["isTopLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsTopLevel(val)
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
    res["placeholderText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPlaceholderText(val)
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManangementIntentValueType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManangementIntentValueType)
        m.SetValueType(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementSettingDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementSettingDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("constraints", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDependencies()))
        for i, v := range m.GetDependencies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dependencies", cast)
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
        err = writer.WriteStringValue("documentationUrl", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerSubtitle", m.GetHeaderSubtitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerTitle", m.GetHeaderTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTopLevel", m.GetIsTopLevel())
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
        err = writer.WriteStringValue("placeholderText", m.GetPlaceholderText())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := m.GetValueType().String()
        err = writer.WriteStringValue("valueType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the constraints property value. Collection of constraints for the setting value
// Parameters:
//  - value : Value to set for the constraints property.
func (m *DeviceManagementSettingDefinition) SetConstraints(value []DeviceManagementConstraint)() {
    m.constraints = value
}
// Sets the dependencies property value. Collection of dependencies on other settings
// Parameters:
//  - value : Value to set for the dependencies property.
func (m *DeviceManagementSettingDefinition) SetDependencies(value []DeviceManagementSettingDependency)() {
    m.dependencies = value
}
// Sets the description property value. The setting's description
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementSettingDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The setting's display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementSettingDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the documentationUrl property value. Url to setting documentation
// Parameters:
//  - value : Value to set for the documentationUrl property.
func (m *DeviceManagementSettingDefinition) SetDocumentationUrl(value *string)() {
    m.documentationUrl = value
}
// Sets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
// Parameters:
//  - value : Value to set for the headerSubtitle property.
func (m *DeviceManagementSettingDefinition) SetHeaderSubtitle(value *string)() {
    m.headerSubtitle = value
}
// Sets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
// Parameters:
//  - value : Value to set for the headerTitle property.
func (m *DeviceManagementSettingDefinition) SetHeaderTitle(value *string)() {
    m.headerTitle = value
}
// Sets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
// Parameters:
//  - value : Value to set for the isTopLevel property.
func (m *DeviceManagementSettingDefinition) SetIsTopLevel(value *bool)() {
    m.isTopLevel = value
}
// Sets the keywords property value. Keywords associated with the setting
// Parameters:
//  - value : Value to set for the keywords property.
func (m *DeviceManagementSettingDefinition) SetKeywords(value []string)() {
    m.keywords = value
}
// Sets the placeholderText property value. Placeholder text as an example of valid input
// Parameters:
//  - value : Value to set for the placeholderText property.
func (m *DeviceManagementSettingDefinition) SetPlaceholderText(value *string)() {
    m.placeholderText = value
}
// Sets the valueType property value. The data type of the value. Possible values are: integer, boolean, string, complex, collection, abstractComplex.
// Parameters:
//  - value : Value to set for the valueType property.
func (m *DeviceManagementSettingDefinition) SetValueType(value *DeviceManangementIntentValueType)() {
    m.valueType = value
}
