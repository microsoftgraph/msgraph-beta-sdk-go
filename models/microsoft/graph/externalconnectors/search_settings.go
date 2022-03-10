package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchSettings provides operations to manage the collection of externalConnection entities.
type SearchSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed.
    searchResultTemplates []DisplayTemplateable;
}
// NewSearchSettings instantiates a new searchSettings and sets the default values.
func NewSearchSettings()(*SearchSettings) {
    m := &SearchSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSearchSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["searchResultTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDisplayTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DisplayTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DisplayTemplateable)
            }
            m.SetSearchResultTemplates(res)
        }
        return nil
    }
    return res
}
// GetSearchResultTemplates gets the searchResultTemplates property value. Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed.
func (m *SearchSettings) GetSearchResultTemplates()([]DisplayTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.searchResultTemplates
    }
}
func (m *SearchSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetSearchResultTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSearchResultTemplates()))
        for i, v := range m.GetSearchResultTemplates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("searchResultTemplates", cast)
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
func (m *SearchSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSearchResultTemplates sets the searchResultTemplates property value. Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed.
func (m *SearchSettings) SetSearchResultTemplates(value []DisplayTemplateable)() {
    if m != nil {
        m.searchResultTemplates = value
    }
}
