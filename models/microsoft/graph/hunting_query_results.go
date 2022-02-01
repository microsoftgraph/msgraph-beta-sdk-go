package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HuntingQueryResults 
type HuntingQueryResults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    results []HuntingRowResult;
    // 
    schema []SinglePropertySchema;
}
// NewHuntingQueryResults instantiates a new huntingQueryResults and sets the default values.
func NewHuntingQueryResults()(*HuntingQueryResults) {
    m := &HuntingQueryResults{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HuntingQueryResults) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetResults gets the results property value. 
func (m *HuntingQueryResults) GetResults()([]HuntingRowResult) {
    if m == nil {
        return nil
    } else {
        return m.results
    }
}
// GetSchema gets the schema property value. 
func (m *HuntingQueryResults) GetSchema()([]SinglePropertySchema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *HuntingQueryResults) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetResults() != nil {
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
    if m.GetSchema() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HuntingQueryResults) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResults sets the results property value. 
func (m *HuntingQueryResults) SetResults(value []HuntingRowResult)() {
    if m != nil {
        m.results = value
    }
}
// SetSchema sets the schema property value. 
func (m *HuntingQueryResults) SetSchema(value []SinglePropertySchema)() {
    if m != nil {
        m.schema = value
    }
}
