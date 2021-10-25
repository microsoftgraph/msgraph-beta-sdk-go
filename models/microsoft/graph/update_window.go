package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UpdateWindow struct {
    additionalData map[string]interface{};
    updateWindowEndTime *string;
    updateWindowStartTime *string;
}
func NewUpdateWindow()(*UpdateWindow) {
    m := &UpdateWindow{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UpdateWindow) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UpdateWindow) GetUpdateWindowEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.updateWindowEndTime
    }
}
func (m *UpdateWindow) GetUpdateWindowStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.updateWindowStartTime
    }
}
func (m *UpdateWindow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["updateWindowEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpdateWindowEndTime(val)
        return nil
    }
    res["updateWindowStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpdateWindowStartTime(val)
        return nil
    }
    return res
}
func (m *UpdateWindow) IsNil()(bool) {
    return m == nil
}
func (m *UpdateWindow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("updateWindowEndTime", m.GetUpdateWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updateWindowStartTime", m.GetUpdateWindowStartTime())
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
func (m *UpdateWindow) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UpdateWindow) SetUpdateWindowEndTime(value *string)() {
    m.updateWindowEndTime = value
}
func (m *UpdateWindow) SetUpdateWindowStartTime(value *string)() {
    m.updateWindowStartTime = value
}
