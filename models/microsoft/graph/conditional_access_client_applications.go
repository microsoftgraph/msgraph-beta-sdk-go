package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessClientApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    excludeServicePrincipals []string;
    // 
    includeServicePrincipals []string;
}
// Instantiates a new conditionalAccessClientApplications and sets the default values.
func NewConditionalAccessClientApplications()(*ConditionalAccessClientApplications) {
    m := &ConditionalAccessClientApplications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessClientApplications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludeServicePrincipals property value. 
func (m *ConditionalAccessClientApplications) GetExcludeServicePrincipals()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeServicePrincipals
    }
}
// Gets the includeServicePrincipals property value. 
func (m *ConditionalAccessClientApplications) GetIncludeServicePrincipals()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeServicePrincipals
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessClientApplications) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeServicePrincipals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetExcludeServicePrincipals(res)
        return nil
    }
    res["includeServicePrincipals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetIncludeServicePrincipals(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessClientApplications) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessClientApplications) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeServicePrincipals", m.GetExcludeServicePrincipals())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeServicePrincipals", m.GetIncludeServicePrincipals())
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
func (m *ConditionalAccessClientApplications) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludeServicePrincipals property value. 
// Parameters:
//  - value : Value to set for the excludeServicePrincipals property.
func (m *ConditionalAccessClientApplications) SetExcludeServicePrincipals(value []string)() {
    m.excludeServicePrincipals = value
}
// Sets the includeServicePrincipals property value. 
// Parameters:
//  - value : Value to set for the includeServicePrincipals property.
func (m *ConditionalAccessClientApplications) SetIncludeServicePrincipals(value []string)() {
    m.includeServicePrincipals = value
}
