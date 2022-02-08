package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessDevices 
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
// NewConditionalAccessDevices instantiates a new conditionalAccessDevices and sets the default values.
func NewConditionalAccessDevices()(*ConditionalAccessDevices) {
    m := &ConditionalAccessDevices{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceFilter gets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
func (m *ConditionalAccessDevices) GetDeviceFilter()(*ConditionalAccessFilter) {
    if m == nil {
        return nil
    } else {
        return m.deviceFilter
    }
}
// GetExcludeDevices gets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetExcludeDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeDevices
    }
}
// GetExcludeDeviceStates gets the excludeDeviceStates property value. 
func (m *ConditionalAccessDevices) GetExcludeDeviceStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeDeviceStates
    }
}
// GetIncludeDevices gets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetIncludeDevices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeDevices
    }
}
// GetIncludeDeviceStates gets the includeDeviceStates property value. 
func (m *ConditionalAccessDevices) GetIncludeDeviceStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeDeviceStates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessDevices) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessFilter() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFilter(val.(*ConditionalAccessFilter))
        }
        return nil
    }
    res["excludeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeDevices(res)
        }
        return nil
    }
    res["excludeDeviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeDeviceStates(res)
        }
        return nil
    }
    res["includeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeDevices(res)
        }
        return nil
    }
    res["includeDeviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeDeviceStates(res)
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessDevices) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessDevices) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceFilter", m.GetDeviceFilter())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeDevices() != nil {
        err := writer.WriteCollectionOfStringValues("excludeDevices", m.GetExcludeDevices())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeDeviceStates() != nil {
        err := writer.WriteCollectionOfStringValues("excludeDeviceStates", m.GetExcludeDeviceStates())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeDevices() != nil {
        err := writer.WriteCollectionOfStringValues("includeDevices", m.GetIncludeDevices())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeDeviceStates() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceFilter sets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
func (m *ConditionalAccessDevices) SetDeviceFilter(value *ConditionalAccessFilter)() {
    if m != nil {
        m.deviceFilter = value
    }
}
// SetExcludeDevices sets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) SetExcludeDevices(value []string)() {
    if m != nil {
        m.excludeDevices = value
    }
}
// SetExcludeDeviceStates sets the excludeDeviceStates property value. 
func (m *ConditionalAccessDevices) SetExcludeDeviceStates(value []string)() {
    if m != nil {
        m.excludeDeviceStates = value
    }
}
// SetIncludeDevices sets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) SetIncludeDevices(value []string)() {
    if m != nil {
        m.includeDevices = value
    }
}
// SetIncludeDeviceStates sets the includeDeviceStates property value. 
func (m *ConditionalAccessDevices) SetIncludeDeviceStates(value []string)() {
    if m != nil {
        m.includeDeviceStates = value
    }
}
