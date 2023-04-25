package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingDefinition 
type DeviceManagementConfigurationSettingDefinition struct {
    Entity
}
// NewDeviceManagementConfigurationSettingDefinition instantiates a new deviceManagementConfigurationSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationSettingDefinition()(*DeviceManagementConfigurationSettingDefinition) {
    m := &DeviceManagementConfigurationSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionDefinition":
                        return NewDeviceManagementConfigurationChoiceSettingCollectionDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingDefinition":
                        return NewDeviceManagementConfigurationChoiceSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationRedirectSettingDefinition":
                        return NewDeviceManagementConfigurationRedirectSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionDefinition":
                        return NewDeviceManagementConfigurationSettingGroupCollectionDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupDefinition":
                        return NewDeviceManagementConfigurationSettingGroupDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionDefinition":
                        return NewDeviceManagementConfigurationSimpleSettingCollectionDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingDefinition":
                        return NewDeviceManagementConfigurationSimpleSettingDefinition(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingDefinition(), nil
}
// GetAccessTypes gets the accessTypes property value. The accessTypes property
func (m *DeviceManagementConfigurationSettingDefinition) GetAccessTypes()(*DeviceManagementConfigurationSettingAccessTypes) {
    val, err := m.GetBackingStore().Get("accessTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationSettingAccessTypes)
    }
    return nil
}
// GetApplicability gets the applicability property value. Details which device setting is applicable on. Supports: $filters.
func (m *DeviceManagementConfigurationSettingDefinition) GetApplicability()(DeviceManagementConfigurationSettingApplicabilityable) {
    val, err := m.GetBackingStore().Get("applicability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSettingApplicabilityable)
    }
    return nil
}
// GetBaseUri gets the baseUri property value. Base CSP Path
func (m *DeviceManagementConfigurationSettingDefinition) GetBaseUri()(*string) {
    val, err := m.GetBackingStore().Get("baseUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCategoryId gets the categoryId property value. Specify category in which the setting is under. Support $filters.
func (m *DeviceManagementConfigurationSettingDefinition) GetCategoryId()(*string) {
    val, err := m.GetBackingStore().Get("categoryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the setting.
func (m *DeviceManagementConfigurationSettingDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the setting. For example: Allow Toast.
func (m *DeviceManagementConfigurationSettingDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingAccessTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessTypes(val.(*DeviceManagementConfigurationSettingAccessTypes))
        }
        return nil
    }
    res["applicability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingApplicabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicability(val.(DeviceManagementConfigurationSettingApplicabilityable))
        }
        return nil
    }
    res["baseUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUri(val)
        }
        return nil
    }
    res["categoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryId(val)
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
    res["helpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpText(val)
        }
        return nil
    }
    res["infoUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["occurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingOccurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrence(val.(DeviceManagementConfigurationSettingOccurrenceable))
        }
        return nil
    }
    res["offsetUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffsetUri(val)
        }
        return nil
    }
    res["referredSettingInformationList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationReferredSettingInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationReferredSettingInformationable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationReferredSettingInformationable)
            }
            m.SetReferredSettingInformationList(res)
        }
        return nil
    }
    res["rootDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootDefinitionId(val)
        }
        return nil
    }
    res["settingUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingUsage(val.(*DeviceManagementConfigurationSettingUsage))
        }
        return nil
    }
    res["uxBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationControlType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUxBehavior(val.(*DeviceManagementConfigurationControlType))
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
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetHelpText gets the helpText property value. Help text of the setting. Give more details of the setting.
func (m *DeviceManagementConfigurationSettingDefinition) GetHelpText()(*string) {
    val, err := m.GetBackingStore().Get("helpText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInfoUrls gets the infoUrls property value. List of links more info for the setting can be found at.
func (m *DeviceManagementConfigurationSettingDefinition) GetInfoUrls()([]string) {
    val, err := m.GetBackingStore().Get("infoUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKeywords gets the keywords property value. Tokens which to search settings on
func (m *DeviceManagementConfigurationSettingDefinition) GetKeywords()([]string) {
    val, err := m.GetBackingStore().Get("keywords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetName gets the name property value. Name of the item
func (m *DeviceManagementConfigurationSettingDefinition) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOccurrence gets the occurrence property value. Indicates whether the setting is required or not
func (m *DeviceManagementConfigurationSettingDefinition) GetOccurrence()(DeviceManagementConfigurationSettingOccurrenceable) {
    val, err := m.GetBackingStore().Get("occurrence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSettingOccurrenceable)
    }
    return nil
}
// GetOffsetUri gets the offsetUri property value. Offset CSP Path from Base
func (m *DeviceManagementConfigurationSettingDefinition) GetOffsetUri()(*string) {
    val, err := m.GetBackingStore().Get("offsetUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReferredSettingInformationList gets the referredSettingInformationList property value. List of referred setting information.
func (m *DeviceManagementConfigurationSettingDefinition) GetReferredSettingInformationList()([]DeviceManagementConfigurationReferredSettingInformationable) {
    val, err := m.GetBackingStore().Get("referredSettingInformationList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationReferredSettingInformationable)
    }
    return nil
}
// GetRootDefinitionId gets the rootDefinitionId property value. Root setting definition id if the setting is a child setting.
func (m *DeviceManagementConfigurationSettingDefinition) GetRootDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("rootDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingUsage gets the settingUsage property value. Supported setting types
func (m *DeviceManagementConfigurationSettingDefinition) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    val, err := m.GetBackingStore().Get("settingUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationSettingUsage)
    }
    return nil
}
// GetUxBehavior gets the uxBehavior property value. Setting control type representation in the UX
func (m *DeviceManagementConfigurationSettingDefinition) GetUxBehavior()(*DeviceManagementConfigurationControlType) {
    val, err := m.GetBackingStore().Get("uxBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationControlType)
    }
    return nil
}
// GetVersion gets the version property value. Item Version
func (m *DeviceManagementConfigurationSettingDefinition) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVisibility gets the visibility property value. Supported setting types
func (m *DeviceManagementConfigurationSettingDefinition) GetVisibility()(*DeviceManagementConfigurationSettingVisibility) {
    val, err := m.GetBackingStore().Get("visibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationSettingVisibility)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferredSettingInformationList()))
        for i, v := range m.GetReferredSettingInformationList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetAccessTypes sets the accessTypes property value. The accessTypes property
func (m *DeviceManagementConfigurationSettingDefinition) SetAccessTypes(value *DeviceManagementConfigurationSettingAccessTypes)() {
    err := m.GetBackingStore().Set("accessTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicability sets the applicability property value. Details which device setting is applicable on. Supports: $filters.
func (m *DeviceManagementConfigurationSettingDefinition) SetApplicability(value DeviceManagementConfigurationSettingApplicabilityable)() {
    err := m.GetBackingStore().Set("applicability", value)
    if err != nil {
        panic(err)
    }
}
// SetBaseUri sets the baseUri property value. Base CSP Path
func (m *DeviceManagementConfigurationSettingDefinition) SetBaseUri(value *string)() {
    err := m.GetBackingStore().Set("baseUri", value)
    if err != nil {
        panic(err)
    }
}
// SetCategoryId sets the categoryId property value. Specify category in which the setting is under. Support $filters.
func (m *DeviceManagementConfigurationSettingDefinition) SetCategoryId(value *string)() {
    err := m.GetBackingStore().Set("categoryId", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the setting.
func (m *DeviceManagementConfigurationSettingDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the setting. For example: Allow Toast.
func (m *DeviceManagementConfigurationSettingDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHelpText sets the helpText property value. Help text of the setting. Give more details of the setting.
func (m *DeviceManagementConfigurationSettingDefinition) SetHelpText(value *string)() {
    err := m.GetBackingStore().Set("helpText", value)
    if err != nil {
        panic(err)
    }
}
// SetInfoUrls sets the infoUrls property value. List of links more info for the setting can be found at.
func (m *DeviceManagementConfigurationSettingDefinition) SetInfoUrls(value []string)() {
    err := m.GetBackingStore().Set("infoUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetKeywords sets the keywords property value. Tokens which to search settings on
func (m *DeviceManagementConfigurationSettingDefinition) SetKeywords(value []string)() {
    err := m.GetBackingStore().Set("keywords", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name of the item
func (m *DeviceManagementConfigurationSettingDefinition) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOccurrence sets the occurrence property value. Indicates whether the setting is required or not
func (m *DeviceManagementConfigurationSettingDefinition) SetOccurrence(value DeviceManagementConfigurationSettingOccurrenceable)() {
    err := m.GetBackingStore().Set("occurrence", value)
    if err != nil {
        panic(err)
    }
}
// SetOffsetUri sets the offsetUri property value. Offset CSP Path from Base
func (m *DeviceManagementConfigurationSettingDefinition) SetOffsetUri(value *string)() {
    err := m.GetBackingStore().Set("offsetUri", value)
    if err != nil {
        panic(err)
    }
}
// SetReferredSettingInformationList sets the referredSettingInformationList property value. List of referred setting information.
func (m *DeviceManagementConfigurationSettingDefinition) SetReferredSettingInformationList(value []DeviceManagementConfigurationReferredSettingInformationable)() {
    err := m.GetBackingStore().Set("referredSettingInformationList", value)
    if err != nil {
        panic(err)
    }
}
// SetRootDefinitionId sets the rootDefinitionId property value. Root setting definition id if the setting is a child setting.
func (m *DeviceManagementConfigurationSettingDefinition) SetRootDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("rootDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingUsage sets the settingUsage property value. Supported setting types
func (m *DeviceManagementConfigurationSettingDefinition) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    err := m.GetBackingStore().Set("settingUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetUxBehavior sets the uxBehavior property value. Setting control type representation in the UX
func (m *DeviceManagementConfigurationSettingDefinition) SetUxBehavior(value *DeviceManagementConfigurationControlType)() {
    err := m.GetBackingStore().Set("uxBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Item Version
func (m *DeviceManagementConfigurationSettingDefinition) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// SetVisibility sets the visibility property value. Supported setting types
func (m *DeviceManagementConfigurationSettingDefinition) SetVisibility(value *DeviceManagementConfigurationSettingVisibility)() {
    err := m.GetBackingStore().Set("visibility", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSettingDefinitionable 
type DeviceManagementConfigurationSettingDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessTypes()(*DeviceManagementConfigurationSettingAccessTypes)
    GetApplicability()(DeviceManagementConfigurationSettingApplicabilityable)
    GetBaseUri()(*string)
    GetCategoryId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHelpText()(*string)
    GetInfoUrls()([]string)
    GetKeywords()([]string)
    GetName()(*string)
    GetOccurrence()(DeviceManagementConfigurationSettingOccurrenceable)
    GetOffsetUri()(*string)
    GetReferredSettingInformationList()([]DeviceManagementConfigurationReferredSettingInformationable)
    GetRootDefinitionId()(*string)
    GetSettingUsage()(*DeviceManagementConfigurationSettingUsage)
    GetUxBehavior()(*DeviceManagementConfigurationControlType)
    GetVersion()(*string)
    GetVisibility()(*DeviceManagementConfigurationSettingVisibility)
    SetAccessTypes(value *DeviceManagementConfigurationSettingAccessTypes)()
    SetApplicability(value DeviceManagementConfigurationSettingApplicabilityable)()
    SetBaseUri(value *string)()
    SetCategoryId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHelpText(value *string)()
    SetInfoUrls(value []string)()
    SetKeywords(value []string)()
    SetName(value *string)()
    SetOccurrence(value DeviceManagementConfigurationSettingOccurrenceable)()
    SetOffsetUri(value *string)()
    SetReferredSettingInformationList(value []DeviceManagementConfigurationReferredSettingInformationable)()
    SetRootDefinitionId(value *string)()
    SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)()
    SetUxBehavior(value *DeviceManagementConfigurationControlType)()
    SetVersion(value *string)()
    SetVisibility(value *DeviceManagementConfigurationSettingVisibility)()
}
