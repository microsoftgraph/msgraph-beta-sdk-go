package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingDefinition entity representing the defintion for a given setting
type DeviceManagementSettingDefinition struct {
    Entity
}
// NewDeviceManagementSettingDefinition instantiates a new deviceManagementSettingDefinition and sets the default values.
func NewDeviceManagementSettingDefinition()(*DeviceManagementSettingDefinition) {
    m := &DeviceManagementSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementAbstractComplexSettingDefinition":
                        return NewDeviceManagementAbstractComplexSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementCollectionSettingDefinition":
                        return NewDeviceManagementCollectionSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementComplexSettingDefinition":
                        return NewDeviceManagementComplexSettingDefinition(), nil
                }
            }
        }
    }
    return NewDeviceManagementSettingDefinition(), nil
}
// GetConstraints gets the constraints property value. Collection of constraints for the setting value
func (m *DeviceManagementSettingDefinition) GetConstraints()([]DeviceManagementConstraintable) {
    val, err := m.GetBackingStore().Get("constraints")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConstraintable)
    }
    return nil
}
// GetDependencies gets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) GetDependencies()([]DeviceManagementSettingDependencyable) {
    val, err := m.GetBackingStore().Get("dependencies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingDependencyable)
    }
    return nil
}
// GetDescription gets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDocumentationUrl gets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) GetDocumentationUrl()(*string) {
    val, err := m.GetBackingStore().Get("documentationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["constraints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConstraintable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConstraintable)
                }
            }
            m.SetConstraints(res)
        }
        return nil
    }
    res["dependencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingDependencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDependencyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingDependencyable)
                }
            }
            m.SetDependencies(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["documentationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
    res["headerSubtitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderSubtitle(val)
        }
        return nil
    }
    res["headerTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderTitle(val)
        }
        return nil
    }
    res["isTopLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTopLevel(val)
        }
        return nil
    }
    res["keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetKeywords(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["placeholderText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaceholderText(val)
        }
        return nil
    }
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetHeaderSubtitle gets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) GetHeaderSubtitle()(*string) {
    val, err := m.GetBackingStore().Get("headerSubtitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHeaderTitle gets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) GetHeaderTitle()(*string) {
    val, err := m.GetBackingStore().Get("headerTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsTopLevel gets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) GetIsTopLevel()(*bool) {
    val, err := m.GetBackingStore().Get("isTopLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeywords gets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) GetKeywords()([]string) {
    val, err := m.GetBackingStore().Get("keywords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingDefinition) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlaceholderText gets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) GetPlaceholderText()(*string) {
    val, err := m.GetBackingStore().Get("placeholderText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValueType gets the valueType property value. The valueType property
func (m *DeviceManagementSettingDefinition) GetValueType()(*DeviceManangementIntentValueType) {
    val, err := m.GetBackingStore().Get("valueType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManangementIntentValueType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConstraints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("constraints", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependencies()))
        for i, v := range m.GetDependencies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *DeviceManagementSettingDefinition) SetConstraints(value []DeviceManagementConstraintable)() {
    err := m.GetBackingStore().Set("constraints", value)
    if err != nil {
        panic(err)
    }
}
// SetDependencies sets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) SetDependencies(value []DeviceManagementSettingDependencyable)() {
    err := m.GetBackingStore().Set("dependencies", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentationUrl sets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) SetDocumentationUrl(value *string)() {
    err := m.GetBackingStore().Set("documentationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetHeaderSubtitle sets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) SetHeaderSubtitle(value *string)() {
    err := m.GetBackingStore().Set("headerSubtitle", value)
    if err != nil {
        panic(err)
    }
}
// SetHeaderTitle sets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) SetHeaderTitle(value *string)() {
    err := m.GetBackingStore().Set("headerTitle", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTopLevel sets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) SetIsTopLevel(value *bool)() {
    err := m.GetBackingStore().Set("isTopLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetKeywords sets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) SetKeywords(value []string)() {
    err := m.GetBackingStore().Set("keywords", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingDefinition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPlaceholderText sets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) SetPlaceholderText(value *string)() {
    err := m.GetBackingStore().Set("placeholderText", value)
    if err != nil {
        panic(err)
    }
}
// SetValueType sets the valueType property value. The valueType property
func (m *DeviceManagementSettingDefinition) SetValueType(value *DeviceManangementIntentValueType)() {
    err := m.GetBackingStore().Set("valueType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementSettingDefinitionable 
type DeviceManagementSettingDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConstraints()([]DeviceManagementConstraintable)
    GetDependencies()([]DeviceManagementSettingDependencyable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDocumentationUrl()(*string)
    GetHeaderSubtitle()(*string)
    GetHeaderTitle()(*string)
    GetIsTopLevel()(*bool)
    GetKeywords()([]string)
    GetOdataType()(*string)
    GetPlaceholderText()(*string)
    GetValueType()(*DeviceManangementIntentValueType)
    SetConstraints(value []DeviceManagementConstraintable)()
    SetDependencies(value []DeviceManagementSettingDependencyable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDocumentationUrl(value *string)()
    SetHeaderSubtitle(value *string)()
    SetHeaderTitle(value *string)()
    SetIsTopLevel(value *bool)()
    SetKeywords(value []string)()
    SetOdataType(value *string)()
    SetPlaceholderText(value *string)()
    SetValueType(value *DeviceManangementIntentValueType)()
}
