package copytodefaultcontentlocation

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CopyToDefaultContentLocationRequestBody provides operations to call the copyToDefaultContentLocation method.
type CopyToDefaultContentLocationRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    destinationFileName *string;
    // 
    sourceFile i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReferenceable;
}
// NewCopyToDefaultContentLocationRequestBody instantiates a new copyToDefaultContentLocationRequestBody and sets the default values.
func NewCopyToDefaultContentLocationRequestBody()(*CopyToDefaultContentLocationRequestBody) {
    m := &CopyToDefaultContentLocationRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCopyToDefaultContentLocationRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopyToDefaultContentLocationRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCopyToDefaultContentLocationRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyToDefaultContentLocationRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDestinationFileName gets the destinationFileName property value. 
func (m *CopyToDefaultContentLocationRequestBody) GetDestinationFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationFileName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopyToDefaultContentLocationRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["destinationFileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFileName(val)
        }
        return nil
    }
    res["sourceFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFile(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetSourceFile gets the sourceFile property value. 
func (m *CopyToDefaultContentLocationRequestBody) GetSourceFile()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReferenceable) {
    if m == nil {
        return nil
    } else {
        return m.sourceFile
    }
}
func (m *CopyToDefaultContentLocationRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CopyToDefaultContentLocationRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationFileName", m.GetDestinationFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceFile", m.GetSourceFile())
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
func (m *CopyToDefaultContentLocationRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDestinationFileName sets the destinationFileName property value. 
func (m *CopyToDefaultContentLocationRequestBody) SetDestinationFileName(value *string)() {
    if m != nil {
        m.destinationFileName = value
    }
}
// SetSourceFile sets the sourceFile property value. 
func (m *CopyToDefaultContentLocationRequestBody) SetSourceFile(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemReferenceable)() {
    if m != nil {
        m.sourceFile = value
    }
}
