package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewRecurrenceSettings struct {
    additionalData map[string]interface{};
    durationInDays *int32;
    recurrenceCount *int32;
    recurrenceEndType *string;
    recurrenceType *string;
}
func NewAccessReviewRecurrenceSettings()(*AccessReviewRecurrenceSettings) {
    m := &AccessReviewRecurrenceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewRecurrenceSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewRecurrenceSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
func (m *AccessReviewRecurrenceSettings) GetRecurrenceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceCount
    }
}
func (m *AccessReviewRecurrenceSettings) GetRecurrenceEndType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceEndType
    }
}
func (m *AccessReviewRecurrenceSettings) GetRecurrenceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceType
    }
}
func (m *AccessReviewRecurrenceSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInDays(val)
        return nil
    }
    res["recurrenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRecurrenceCount(val)
        return nil
    }
    res["recurrenceEndType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecurrenceEndType(val)
        return nil
    }
    res["recurrenceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecurrenceType(val)
        return nil
    }
    return res
}
func (m *AccessReviewRecurrenceSettings) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewRecurrenceSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("recurrenceCount", m.GetRecurrenceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceEndType", m.GetRecurrenceEndType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceType", m.GetRecurrenceType())
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
func (m *AccessReviewRecurrenceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessReviewRecurrenceSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
func (m *AccessReviewRecurrenceSettings) SetRecurrenceCount(value *int32)() {
    m.recurrenceCount = value
}
func (m *AccessReviewRecurrenceSettings) SetRecurrenceEndType(value *string)() {
    m.recurrenceEndType = value
}
func (m *AccessReviewRecurrenceSettings) SetRecurrenceType(value *string)() {
    m.recurrenceType = value
}
