package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessDeviceStates 
type ConditionalAccessDeviceStates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
    excludeStates []string;
    // States in the scope of the policy. All is the only allowed value.
    includeStates []string;
}
// NewConditionalAccessDeviceStates instantiates a new conditionalAccessDeviceStates and sets the default values.
func NewConditionalAccessDeviceStates()(*ConditionalAccessDeviceStates) {
    m := &ConditionalAccessDeviceStates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDeviceStates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeStates gets the excludeStates property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
func (m *ConditionalAccessDeviceStates) GetExcludeStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeStates
    }
}
// GetIncludeStates gets the includeStates property value. States in the scope of the policy. All is the only allowed value.
func (m *ConditionalAccessDeviceStates) GetIncludeStates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeStates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessDeviceStates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeStates(res)
        }
        return nil
    }
    res["includeStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeStates(res)
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessDeviceStates) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessDeviceStates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetExcludeStates() != nil {
        err := writer.WriteCollectionOfStringValues("excludeStates", m.GetExcludeStates())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeStates() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDeviceStates) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeStates sets the excludeStates property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
func (m *ConditionalAccessDeviceStates) SetExcludeStates(value []string)() {
    if m != nil {
        m.excludeStates = value
    }
}
// SetIncludeStates sets the includeStates property value. States in the scope of the policy. All is the only allowed value.
func (m *ConditionalAccessDeviceStates) SetIncludeStates(value []string)() {
    if m != nil {
        m.includeStates = value
    }
}
