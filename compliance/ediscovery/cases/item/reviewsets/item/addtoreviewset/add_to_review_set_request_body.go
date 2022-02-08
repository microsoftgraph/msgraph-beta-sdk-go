package addtoreviewset

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// AddToReviewSetRequestBody 
type AddToReviewSetRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    additionalDataOptions *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.AdditionalDataOptions;
    // 
    sourceCollection *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection;
}
// NewAddToReviewSetRequestBody instantiates a new addToReviewSetRequestBody and sets the default values.
func NewAddToReviewSetRequestBody()(*AddToReviewSetRequestBody) {
    m := &AddToReviewSetRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddToReviewSetRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalDataOptions gets the additionalDataOptions property value. 
func (m *AddToReviewSetRequestBody) GetAdditionalDataOptions()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.AdditionalDataOptions) {
    if m == nil {
        return nil
    } else {
        return m.additionalDataOptions
    }
}
// GetSourceCollection gets the sourceCollection property value. 
func (m *AddToReviewSetRequestBody) GetSourceCollection()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollection
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddToReviewSetRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalDataOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseAdditionalDataOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDataOptions(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.AdditionalDataOptions))
        }
        return nil
    }
    res["sourceCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSourceCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection))
        }
        return nil
    }
    return res
}
func (m *AddToReviewSetRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AddToReviewSetRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAdditionalDataOptions() != nil {
        cast := (*m.GetAdditionalDataOptions()).String()
        err := writer.WriteStringValue("additionalDataOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
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
func (m *AddToReviewSetRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. 
func (m *AddToReviewSetRequestBody) SetAdditionalDataOptions(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.AdditionalDataOptions)() {
    if m != nil {
        m.additionalDataOptions = value
    }
}
// SetSourceCollection sets the sourceCollection property value. 
func (m *AddToReviewSetRequestBody) SetSourceCollection(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection)() {
    if m != nil {
        m.sourceCollection = value
    }
}
