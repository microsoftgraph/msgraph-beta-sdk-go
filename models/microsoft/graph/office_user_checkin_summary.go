package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OfficeUserCheckinSummary struct {
    additionalData map[string]interface{};
    failedUserCount *int32;
    succeededUserCount *int32;
}
func NewOfficeUserCheckinSummary()(*OfficeUserCheckinSummary) {
    m := &OfficeUserCheckinSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OfficeUserCheckinSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OfficeUserCheckinSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
func (m *OfficeUserCheckinSummary) GetSucceededUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.succeededUserCount
    }
}
func (m *OfficeUserCheckinSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedUserCount(val)
        return nil
    }
    res["succeededUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSucceededUserCount(val)
        return nil
    }
    return res
}
func (m *OfficeUserCheckinSummary) IsNil()(bool) {
    return m == nil
}
func (m *OfficeUserCheckinSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("succeededUserCount", m.GetSucceededUserCount())
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
func (m *OfficeUserCheckinSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OfficeUserCheckinSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
func (m *OfficeUserCheckinSummary) SetSucceededUserCount(value *int32)() {
    m.succeededUserCount = value
}
