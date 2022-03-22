package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContactMergeSuggestions 
type ContactMergeSuggestions struct {
    Entity
    // true if the duplicate contact merge suggestions feature is enabled for the user; false if the feature is disabled. Default value is true.
    isEnabled *bool;
}
// NewContactMergeSuggestions instantiates a new contactMergeSuggestions and sets the default values.
func NewContactMergeSuggestions()(*ContactMergeSuggestions) {
    m := &ContactMergeSuggestions{
        Entity: *NewEntity(),
    }
    return m
}
// CreateContactMergeSuggestionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContactMergeSuggestionsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewContactMergeSuggestions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContactMergeSuggestions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. true if the duplicate contact merge suggestions feature is enabled for the user; false if the feature is disabled. Default value is true.
func (m *ContactMergeSuggestions) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Serialize serializes information the current object
func (m *ContactMergeSuggestions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEnabled sets the isEnabled property value. true if the duplicate contact merge suggestions feature is enabled for the user; false if the feature is disabled. Default value is true.
func (m *ContactMergeSuggestions) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
