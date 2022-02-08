package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingDefinition 
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
// NewDeviceManagementSettingDefinition instantiates a new deviceManagementSettingDefinition and sets the default values.
func NewDeviceManagementSettingDefinition()(*DeviceManagementSettingDefinition) {
    m := &DeviceManagementSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetConstraints gets the constraints property value. Collection of constraints for the setting value
func (m *DeviceManagementSettingDefinition) GetConstraints()([]DeviceManagementConstraint) {
    if m == nil {
        return nil
    } else {
        return m.constraints
    }
}
// GetDependencies gets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) GetDependencies()([]DeviceManagementSettingDependency) {
    if m == nil {
        return nil
    } else {
        return m.dependencies
    }
}
// GetDescription gets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDocumentationUrl gets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) GetDocumentationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentationUrl
    }
}
// GetHeaderSubtitle gets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) GetHeaderSubtitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerSubtitle
    }
}
// GetHeaderTitle gets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) GetHeaderTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerTitle
    }
}
// GetIsTopLevel gets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) GetIsTopLevel()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTopLevel
    }
}
// GetKeywords gets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// GetPlaceholderText gets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) GetPlaceholderText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.placeholderText
    }
}
// GetValueType gets the valueType property value. The data type of the value. Possible values are: integer, boolean, string, complex, collection, abstractComplex.
func (m *DeviceManagementSettingDefinition) GetValueType()(*DeviceManangementIntentValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["constraints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConstraint() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConstraint, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConstraint))
            }
            m.SetConstraints(res)
        }
        return nil
    }
    res["dependencies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingDependency() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDependency, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementSettingDependency))
            }
            m.SetDependencies(res)
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
    res["documentationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
    res["headerSubtitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderSubtitle(val)
        }
        return nil
    }
    res["headerTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderTitle(val)
        }
        return nil
    }
    res["isTopLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTopLevel(val)
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
    res["placeholderText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaceholderText(val)
        }
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManangementIntentValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*DeviceManangementIntentValueType))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementSettingDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConstraints() != nil {
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
    if m.GetDependencies() != nil {
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
    if m.GetKeywords() != nil {
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
        cast := (*m.GetValueType()).String()
        err = writer.WriteStringValue("valueType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConstraints sets the constraints property value. Collection of constraints for the setting value
func (m *DeviceManagementSettingDefinition) SetConstraints(value []DeviceManagementConstraint)() {
    if m != nil {
        m.constraints = value
    }
}
// SetDependencies sets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) SetDependencies(value []DeviceManagementSettingDependency)() {
    if m != nil {
        m.dependencies = value
    }
}
// SetDescription sets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDocumentationUrl sets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) SetDocumentationUrl(value *string)() {
    if m != nil {
        m.documentationUrl = value
    }
}
// SetHeaderSubtitle sets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) SetHeaderSubtitle(value *string)() {
    if m != nil {
        m.headerSubtitle = value
    }
}
// SetHeaderTitle sets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) SetHeaderTitle(value *string)() {
    if m != nil {
        m.headerTitle = value
    }
}
// SetIsTopLevel sets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) SetIsTopLevel(value *bool)() {
    if m != nil {
        m.isTopLevel = value
    }
}
// SetKeywords sets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) SetKeywords(value []string)() {
    if m != nil {
        m.keywords = value
    }
}
// SetPlaceholderText sets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) SetPlaceholderText(value *string)() {
    if m != nil {
        m.placeholderText = value
    }
}
// SetValueType sets the valueType property value. The data type of the value. Possible values are: integer, boolean, string, complex, collection, abstractComplex.
func (m *DeviceManagementSettingDefinition) SetValueType(value *DeviceManangementIntentValueType)() {
    if m != nil {
        m.valueType = value
    }
}
