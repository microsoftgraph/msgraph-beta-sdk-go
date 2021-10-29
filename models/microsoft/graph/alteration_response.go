package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AlterationResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Defines the original user query string.
    originalQueryString *string;
    // Defines the details of alteration information for the spelling correction.
    queryAlteration *SearchAlteration;
    // Defines the type of the spelling correction. Possible values are suggestion, modification.
    queryAlterationType *SearchAlterationType;
}
// Instantiates a new alterationResponse and sets the default values.
func NewAlterationResponse()(*AlterationResponse) {
    m := &AlterationResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlterationResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the originalQueryString property value. Defines the original user query string.
func (m *AlterationResponse) GetOriginalQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalQueryString
    }
}
// Gets the queryAlteration property value. Defines the details of alteration information for the spelling correction.
func (m *AlterationResponse) GetQueryAlteration()(*SearchAlteration) {
    if m == nil {
        return nil
    } else {
        return m.queryAlteration
    }
}
// Gets the queryAlterationType property value. Defines the type of the spelling correction. Possible values are suggestion, modification.
func (m *AlterationResponse) GetQueryAlterationType()(*SearchAlterationType) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationType
    }
}
// The deserialization information for the current model
func (m *AlterationResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["originalQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginalQueryString(val)
        return nil
    }
    res["queryAlteration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchAlteration() })
        if err != nil {
            return err
        }
        m.SetQueryAlteration(val.(*SearchAlteration))
        return nil
    }
    res["queryAlterationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSearchAlterationType)
        if err != nil {
            return err
        }
        cast := val.(SearchAlterationType)
        m.SetQueryAlterationType(&cast)
        return nil
    }
    return res
}
func (m *AlterationResponse) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AlterationResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("originalQueryString", m.GetOriginalQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("queryAlteration", m.GetQueryAlteration())
        if err != nil {
            return err
        }
    }
    if m.GetQueryAlterationType() != nil {
        cast := m.GetQueryAlterationType().String()
        err := writer.WriteStringValue("queryAlterationType", &cast)
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
func (m *AlterationResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the originalQueryString property value. Defines the original user query string.
// Parameters:
//  - value : Value to set for the originalQueryString property.
func (m *AlterationResponse) SetOriginalQueryString(value *string)() {
    m.originalQueryString = value
}
// Sets the queryAlteration property value. Defines the details of alteration information for the spelling correction.
// Parameters:
//  - value : Value to set for the queryAlteration property.
func (m *AlterationResponse) SetQueryAlteration(value *SearchAlteration)() {
    m.queryAlteration = value
}
// Sets the queryAlterationType property value. Defines the type of the spelling correction. Possible values are suggestion, modification.
// Parameters:
//  - value : Value to set for the queryAlterationType property.
func (m *AlterationResponse) SetQueryAlterationType(value *SearchAlterationType)() {
    m.queryAlterationType = value
}
