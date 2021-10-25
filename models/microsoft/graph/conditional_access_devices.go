package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessDevices struct {
    additionalData map[string]interface{};
    deviceFilter *ConditionalAccessFilter;
    excludeDevices []string;
    excludeDeviceStates []string;
    includeDevices []string;
    includeDeviceStates []string;
}
func NewConditionalAccessDevices()(*ConditionalAccessDevices) {
    m := &ConditionalAccessDevices{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessDevices) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessDevices) GetDeviceFilter()(*ConditionalAccessFilter) {
    if m == nil {
        return nil
    } else {
        return m.deviceFilter
    }
}
func (m *ConditionalAccessDevices) GetExcludeDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeDevices
    }
}
func (m *ConditionalAccessDevices) GetExcludeDeviceStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeDeviceStates
    }
}
func (m *ConditionalAccessDevices) GetIncludeDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeDevices
    }
}
func (m *ConditionalAccessDevices) GetIncludeDeviceStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeDeviceStates
    }
}
func (m *ConditionalAccessDevices) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessFilter() })
        if err != nil {
            return err
        }
        m.SetDeviceFilter(val.(*ConditionalAccessFilter))
        return nil
    }
    res["excludeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeDevices(res)
        return nil
    }
    res["excludeDeviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeDeviceStates(res)
        return nil
    }
    res["includeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeDevices(res)
        return nil
    }
    res["includeDeviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeDeviceStates(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessDevices) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessDevices) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceFilter", m.GetDeviceFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("excludeDevices", m.GetExcludeDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("excludeDeviceStates", m.GetExcludeDeviceStates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeDevices", m.GetIncludeDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeDeviceStates", m.GetIncludeDeviceStates())
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
func (m *ConditionalAccessDevices) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessDevices) SetDeviceFilter(value *ConditionalAccessFilter)() {
    m.deviceFilter = value
}
func (m *ConditionalAccessDevices) SetExcludeDevices(value []string)() {
    m.excludeDevices = value
}
func (m *ConditionalAccessDevices) SetExcludeDeviceStates(value []string)() {
    m.excludeDeviceStates = value
}
func (m *ConditionalAccessDevices) SetIncludeDevices(value []string)() {
    m.includeDevices = value
}
func (m *ConditionalAccessDevices) SetIncludeDeviceStates(value []string)() {
    m.includeDeviceStates = value
}
