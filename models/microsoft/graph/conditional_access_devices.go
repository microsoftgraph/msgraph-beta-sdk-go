package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessDevices struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
    deviceFilter *ConditionalAccessFilter;
    // States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
    excludeDevices []string;
    // 
    excludeDeviceStates []string;
    // States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
    includeDevices []string;
    // 
    includeDeviceStates []string;
}
// Instantiates a new conditionalAccessDevices and sets the default values.
func NewConditionalAccessDevices()(*ConditionalAccessDevices) {
    m := &ConditionalAccessDevices{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
func (m *ConditionalAccessDevices) GetDeviceFilter()(*ConditionalAccessFilter) {
    if m == nil {
        return nil
    } else {
        return m.deviceFilter
    }
}
// Gets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetExcludeDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeDevices
    }
}
// Gets the excludeDeviceStates property value. 
func (m *ConditionalAccessDevices) GetExcludeDeviceStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeDeviceStates
    }
}
// Gets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetIncludeDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeDevices
    }
}
// Gets the includeDeviceStates property value. 
func (m *ConditionalAccessDevices) GetIncludeDeviceStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeDeviceStates
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ConditionalAccessDevices) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
// Parameters:
//  - value : Value to set for the deviceFilter property.
func (m *ConditionalAccessDevices) SetDeviceFilter(value *ConditionalAccessFilter)() {
    m.deviceFilter = value
}
// Sets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
// Parameters:
//  - value : Value to set for the excludeDevices property.
func (m *ConditionalAccessDevices) SetExcludeDevices(value []string)() {
    m.excludeDevices = value
}
// Sets the excludeDeviceStates property value. 
// Parameters:
//  - value : Value to set for the excludeDeviceStates property.
func (m *ConditionalAccessDevices) SetExcludeDeviceStates(value []string)() {
    m.excludeDeviceStates = value
}
// Sets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
// Parameters:
//  - value : Value to set for the includeDevices property.
func (m *ConditionalAccessDevices) SetIncludeDevices(value []string)() {
    m.includeDevices = value
}
// Sets the includeDeviceStates property value. 
// Parameters:
//  - value : Value to set for the includeDeviceStates property.
func (m *ConditionalAccessDevices) SetIncludeDeviceStates(value []string)() {
    m.includeDeviceStates = value
}
