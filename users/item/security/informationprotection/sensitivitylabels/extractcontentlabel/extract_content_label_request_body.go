package extractcontentlabel

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExtractContentLabelRequestBody provides operations to call the extractContentLabel method.
type ExtractContentLabelRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    contentInfo i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentInfoable;
}
// NewExtractContentLabelRequestBody instantiates a new extractContentLabelRequestBody and sets the default values.
func NewExtractContentLabelRequestBody()(*ExtractContentLabelRequestBody) {
    m := &ExtractContentLabelRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExtractContentLabelRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExtractContentLabelRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExtractContentLabelRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExtractContentLabelRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentInfo gets the contentInfo property value. 
func (m *ExtractContentLabelRequestBody) GetContentInfo()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentInfoable) {
    if m == nil {
        return nil
    } else {
        return m.contentInfo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExtractContentLabelRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentInfoable))
        }
        return nil
    }
    return res
}
func (m *ExtractContentLabelRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExtractContentLabelRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
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
func (m *ExtractContentLabelRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentInfo sets the contentInfo property value. 
func (m *ExtractContentLabelRequestBody) SetContentInfo(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentInfoable)() {
    if m != nil {
        m.contentInfo = value
    }
}
