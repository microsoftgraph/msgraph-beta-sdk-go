package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchDetectedSensitiveContent 
type ExactMatchDetectedSensitiveContent struct {
    DetectedSensitiveContentBase
    // 
    matches []SensitiveContentLocationable;
}
// NewExactMatchDetectedSensitiveContent instantiates a new exactMatchDetectedSensitiveContent and sets the default values.
func NewExactMatchDetectedSensitiveContent()(*ExactMatchDetectedSensitiveContent) {
    m := &ExactMatchDetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExactMatchDetectedSensitiveContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["matches"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitiveContentLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentLocationable, len(val))
            for i, v := range val {
                res[i] = v.(SensitiveContentLocationable)
            }
            m.SetMatches(res)
        }
        return nil
    }
    return res
}
// GetMatches gets the matches property value. 
func (m *ExactMatchDetectedSensitiveContent) GetMatches()([]SensitiveContentLocationable) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatches sets the matches property value. 
func (m *ExactMatchDetectedSensitiveContent) SetMatches(value []SensitiveContentLocationable)() {
    if m != nil {
        m.matches = value
    }
}
