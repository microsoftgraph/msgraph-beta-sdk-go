package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type HuntingQueryResults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    results []HuntingRowResult;
    // 
    schema []SinglePropertySchema;
}
// Instantiates a new huntingQueryResults and sets the default values.
func NewHuntingQueryResults()(*HuntingQueryResults) {
    m := &HuntingQueryResults{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HuntingQueryResults) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the results property value. 
func (m *HuntingQueryResults) GetResults()([]HuntingRowResult) {
    if m == nil {
        return nil
    } else {
        return m.results
    }
}
// Gets the schema property value. 
func (m *HuntingQueryResults) GetSchema()([]SinglePropertySchema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// The deserialization information for the current model
func (m *HuntingQueryResults) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["results"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHuntingRowResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HuntingRowResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*HuntingRowResult))
            }
            m.SetResults(res)
        }
        return nil
    }
    res["schema"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSinglePropertySchema() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SinglePropertySchema, len(val))
            for i, v := range val {
                res[i] = *(v.(*SinglePropertySchema))
            }
            m.SetSchema(res)
        }
        return nil
    }
    return res
}
func (m *HuntingQueryResults) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *HuntingQueryResults) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("results", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchema()))
        for i, v := range m.GetSchema() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("schema", cast)
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
func (m *HuntingQueryResults) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the results property value. 
// Parameters:
//  - value : Value to set for the results property.
func (m *HuntingQueryResults) SetResults(value []HuntingRowResult)() {
    m.results = value
}
// Sets the schema property value. 
// Parameters:
//  - value : Value to set for the schema property.
func (m *HuntingQueryResults) SetSchema(value []SinglePropertySchema)() {
    m.schema = value
}
