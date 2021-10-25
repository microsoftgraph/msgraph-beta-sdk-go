package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementTemplate struct {
    Entity
    categories []DeviceManagementTemplateSettingCategory;
    description *string;
    displayName *string;
    intentCount *int32;
    isDeprecated *bool;
    migratableTo []DeviceManagementTemplate;
    platformType *PolicyPlatformType;
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    settings []DeviceManagementSettingInstance;
    templateSubtype *DeviceManagementTemplateSubtype;
    templateType *DeviceManagementTemplateType;
    versionInfo *string;
}
func NewDeviceManagementTemplate()(*DeviceManagementTemplate) {
    m := &DeviceManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementTemplate) GetCategories()([]DeviceManagementTemplateSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
func (m *DeviceManagementTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceManagementTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceManagementTemplate) GetIntentCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.intentCount
    }
}
func (m *DeviceManagementTemplate) GetIsDeprecated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeprecated
    }
}
func (m *DeviceManagementTemplate) GetMigratableTo()([]DeviceManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.migratableTo
    }
}
func (m *DeviceManagementTemplate) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
func (m *DeviceManagementTemplate) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.publishedDateTime
    }
}
func (m *DeviceManagementTemplate) GetSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *DeviceManagementTemplate) GetTemplateSubtype()(*DeviceManagementTemplateSubtype) {
    if m == nil {
        return nil
    } else {
        return m.templateSubtype
    }
}
func (m *DeviceManagementTemplate) GetTemplateType()(*DeviceManagementTemplateType) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
}
func (m *DeviceManagementTemplate) GetVersionInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionInfo
    }
}
func (m *DeviceManagementTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTemplateSettingCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTemplateSettingCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTemplateSettingCategory))
        }
        m.SetCategories(res)
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
    res["intentCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIntentCount(val)
        return nil
    }
    res["isDeprecated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeprecated(val)
        return nil
    }
    res["migratableTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTemplate() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTemplate))
        }
        m.SetMigratableTo(res)
        return nil
    }
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        cast := val.(PolicyPlatformType)
        m.SetPlatformType(&cast)
        return nil
    }
    res["publishedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetPublishedDateTime(val)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingInstance() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementSettingInstance, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementSettingInstance))
        }
        m.SetSettings(res)
        return nil
    }
    res["templateSubtype"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateSubtype)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementTemplateSubtype)
        m.SetTemplateSubtype(&cast)
        return nil
    }
    res["templateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementTemplateType)
        m.SetTemplateType(&cast)
        return nil
    }
    res["versionInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersionInfo(val)
        return nil
    }
    return res
}
func (m *DeviceManagementTemplate) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
        err = writer.WriteInt32Value("intentCount", m.GetIntentCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeprecated", m.GetIsDeprecated())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMigratableTo()))
        for i, v := range m.GetMigratableTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("migratableTo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := m.GetPlatformType().String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateSubtype() != nil {
        cast := m.GetTemplateSubtype().String()
        err = writer.WriteStringValue("templateSubtype", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := m.GetTemplateType().String()
        err = writer.WriteStringValue("templateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionInfo", m.GetVersionInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementTemplate) SetCategories(value []DeviceManagementTemplateSettingCategory)() {
    m.categories = value
}
func (m *DeviceManagementTemplate) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceManagementTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceManagementTemplate) SetIntentCount(value *int32)() {
    m.intentCount = value
}
func (m *DeviceManagementTemplate) SetIsDeprecated(value *bool)() {
    m.isDeprecated = value
}
func (m *DeviceManagementTemplate) SetMigratableTo(value []DeviceManagementTemplate)() {
    m.migratableTo = value
}
func (m *DeviceManagementTemplate) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
func (m *DeviceManagementTemplate) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
func (m *DeviceManagementTemplate) SetSettings(value []DeviceManagementSettingInstance)() {
    m.settings = value
}
func (m *DeviceManagementTemplate) SetTemplateSubtype(value *DeviceManagementTemplateSubtype)() {
    m.templateSubtype = value
}
func (m *DeviceManagementTemplate) SetTemplateType(value *DeviceManagementTemplateType)() {
    m.templateType = value
}
func (m *DeviceManagementTemplate) SetVersionInfo(value *string)() {
    m.versionInfo = value
}
