package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchDetectedSensitiveContent 
type ExactMatchDetectedSensitiveContent struct {
    DetectedSensitiveContentBase
    // 
    matches []SensitiveContentLocation;
}
// NewExactMatchDetectedSensitiveContent instantiates a new exactMatchDetectedSensitiveContent and sets the default values.
func NewExactMatchDetectedSensitiveContent()(*ExactMatchDetectedSensitiveContent) {
    m := &ExactMatchDetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// GetMatches gets the matches property value. 
func (m *ExactMatchDetectedSensitiveContent) GetMatches()([]SensitiveContentLocation) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["matches"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveContentLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitiveContentLocation))
            }
            m.SetMatches(res)
        }
        return nil
    }
    return res
}
func (m *ExactMatchDetectedSensitiveContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExactMatchDetectedSensitiveContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatches() != nil {
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
// SetMatches sets the matches property value. 
func (m *ExactMatchDetectedSensitiveContent) SetMatches(value []SensitiveContentLocation)() {
    if m != nil {
        m.matches = value
    }
}
