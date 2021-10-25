package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessDeviceStates struct {
    additionalData map[string]interface{};
    excludeStates []string;
    includeStates []string;
}
func NewConditionalAccessDeviceStates()(*ConditionalAccessDeviceStates) {
    m := &ConditionalAccessDeviceStates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessDeviceStates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessDeviceStates) GetExcludeStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeStates
    }
}
func (m *ConditionalAccessDeviceStates) GetIncludeStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeStates
    }
}
func (m *ConditionalAccessDeviceStates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeStates(res)
        return nil
    }
    res["includeStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeStates(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessDeviceStates) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessDeviceStates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeStates", m.GetExcludeStates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeStates", m.GetIncludeStates())
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
func (m *ConditionalAccessDeviceStates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessDeviceStates) SetExcludeStates(value []string)() {
    m.excludeStates = value
}
func (m *ConditionalAccessDeviceStates) SetIncludeStates(value []string)() {
    m.includeStates = value
}
