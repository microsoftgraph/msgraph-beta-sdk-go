package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ExactMatchDetectedSensitiveContent struct {
    DetectedSensitiveContentBase
    // 
    matches []SensitiveContentLocation;
}
// Instantiates a new exactMatchDetectedSensitiveContent and sets the default values.
func NewExactMatchDetectedSensitiveContent()(*ExactMatchDetectedSensitiveContent) {
    m := &ExactMatchDetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// Gets the matches property value. 
func (m *ExactMatchDetectedSensitiveContent) GetMatches()([]SensitiveContentLocation) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
}
// The deserialization information for the current model
func (m *ExactMatchDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["matches"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveContentLocation() })
        if err != nil {
            return err
        }
        res := make([]SensitiveContentLocation, len(val))
        for i, v := range val {
            res[i] = *(v.(*SensitiveContentLocation))
        }
        m.SetMatches(res)
        return nil
    }
    return res
}
func (m *ExactMatchDetectedSensitiveContent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExactMatchDetectedSensitiveContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the matches property value. 
// Parameters:
//  - value : Value to set for the matches property.
func (m *ExactMatchDetectedSensitiveContent) SetMatches(value []SensitiveContentLocation)() {
    m.matches = value
}
