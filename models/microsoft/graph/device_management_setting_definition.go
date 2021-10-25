package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementSettingDefinition struct {
    Entity
    constraints []DeviceManagementConstraint;
    dependencies []DeviceManagementSettingDependency;
    description *string;
    displayName *string;
    documentationUrl *string;
    headerSubtitle *string;
    headerTitle *string;
    isTopLevel *bool;
    keywords []string;
    placeholderText *string;
    valueType *DeviceManangementIntentValueType;
}
func NewDeviceManagementSettingDefinition()(*DeviceManagementSettingDefinition) {
    m := &DeviceManagementSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementSettingDefinition) GetConstraints()([]DeviceManagementConstraint) {
    if m == nil {
        return nil
    } else {
        return m.constraints
    }
}
func (m *DeviceManagementSettingDefinition) GetDependencies()([]DeviceManagementSettingDependency) {
    if m == nil {
        return nil
    } else {
        return m.dependencies
    }
}
func (m *DeviceManagementSettingDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceManagementSettingDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceManagementSettingDefinition) GetDocumentationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentationUrl
    }
}
func (m *DeviceManagementSettingDefinition) GetHeaderSubtitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerSubtitle
    }
}
func (m *DeviceManagementSettingDefinition) GetHeaderTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerTitle
    }
}
func (m *DeviceManagementSettingDefinition) GetIsTopLevel()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTopLevel
    }
}
func (m *DeviceManagementSettingDefinition) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
func (m *DeviceManagementSettingDefinition) GetPlaceholderText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.placeholderText
    }
}
func (m *DeviceManagementSettingDefinition) GetValueType()(*DeviceManangementIntentValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
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
func (m *DeviceManagementSettingDefinition) SetConstraints(value []DeviceManagementConstraint)() {
    m.constraints = value
}
func (m *DeviceManagementSettingDefinition) SetDependencies(value []DeviceManagementSettingDependency)() {
    m.dependencies = value
}
func (m *DeviceManagementSettingDefinition) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceManagementSettingDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceManagementSettingDefinition) SetDocumentationUrl(value *string)() {
    m.documentationUrl = value
}
func (m *DeviceManagementSettingDefinition) SetHeaderSubtitle(value *string)() {
    m.headerSubtitle = value
}
func (m *DeviceManagementSettingDefinition) SetHeaderTitle(value *string)() {
    m.headerTitle = value
}
func (m *DeviceManagementSettingDefinition) SetIsTopLevel(value *bool)() {
    m.isTopLevel = value
}
func (m *DeviceManagementSettingDefinition) SetKeywords(value []string)() {
    m.keywords = value
}
func (m *DeviceManagementSettingDefinition) SetPlaceholderText(value *string)() {
    m.placeholderText = value
}
func (m *DeviceManagementSettingDefinition) SetValueType(value *DeviceManangementIntentValueType)() {
    m.valueType = value
}
