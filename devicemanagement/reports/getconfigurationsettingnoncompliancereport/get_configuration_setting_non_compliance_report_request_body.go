package getconfigurationsettingnoncompliancereport

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetConfigurationSettingNonComplianceReportRequestBody struct {
    additionalData map[string]interface{};
    filter *string;
    groupBy []string;
    name *string;
    orderBy []string;
    search *string;
    select_escaped []string;
    sessionId *string;
    skip *int32;
    top *int32;
}
func NewGetConfigurationSettingNonComplianceReportRequestBody()(*GetConfigurationSettingNonComplianceReportRequestBody) {
    m := &GetConfigurationSettingNonComplianceReportRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetGroupBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupBy
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetSearch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetSessionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionId
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFilter(val)
        return nil
    }
    res["groupBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetGroupBy(res)
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
    res["orderBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOrderBy(res)
        return nil
    }
    res["search"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSearch(val)
        return nil
    }
    res["select_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSelect_escaped(res)
        return nil
    }
    res["sessionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSessionId(val)
        return nil
    }
    res["skip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSkip(val)
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTop(val)
        return nil
    }
    return res
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("groupBy", m.GetGroupBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("select_escaped", m.GetSelect_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetFilter(value *string)() {
    m.filter = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetGroupBy(value []string)() {
    m.groupBy = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetName(value *string)() {
    m.name = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetOrderBy(value []string)() {
    m.orderBy = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetSearch(value *string)() {
    m.search = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
func (m *GetConfigurationSettingNonComplianceReportRequestBody) SetTop(value *int32)() {
    m.top = value
}
