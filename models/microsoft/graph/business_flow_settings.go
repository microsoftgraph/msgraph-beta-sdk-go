package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BusinessFlowSettings struct {
    AccessReviewSettings
    durationInDays *int32;
}
func NewBusinessFlowSettings()(*BusinessFlowSettings) {
    m := &BusinessFlowSettings{
        AccessReviewSettings: *NewAccessReviewSettings(),
    }
    return m
}
func (m *BusinessFlowSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
func (m *BusinessFlowSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AccessReviewSettings.GetFieldDeserializers()
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInDays(val)
        return nil
    }
    return res
}
func (m *BusinessFlowSettings) IsNil()(bool) {
    return m == nil
}
func (m *BusinessFlowSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AccessReviewSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *BusinessFlowSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
