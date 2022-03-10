package createdevicelogcollectionrequest

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreateDeviceLogCollectionRequestRequestBody provides operations to call the createDeviceLogCollectionRequest method.
type CreateDeviceLogCollectionRequestRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    templateType i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionRequestable;
}
// NewCreateDeviceLogCollectionRequestRequestBody instantiates a new createDeviceLogCollectionRequestRequestBody and sets the default values.
func NewCreateDeviceLogCollectionRequestRequestBody()(*CreateDeviceLogCollectionRequestRequestBody) {
    m := &CreateDeviceLogCollectionRequestRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCreateDeviceLogCollectionRequestRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreateDeviceLogCollectionRequestRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCreateDeviceLogCollectionRequestRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateDeviceLogCollectionRequestRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateDeviceLogCollectionRequestRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["templateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceLogCollectionRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionRequestable))
        }
        return nil
    }
    return res
}
// GetTemplateType gets the templateType property value. 
func (m *CreateDeviceLogCollectionRequestRequestBody) GetTemplateType()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionRequestable) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
}
func (m *CreateDeviceLogCollectionRequestRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateDeviceLogCollectionRequestRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("templateType", m.GetTemplateType())
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
func (m *CreateDeviceLogCollectionRequestRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTemplateType sets the templateType property value. 
func (m *CreateDeviceLogCollectionRequestRequestBody) SetTemplateType(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionRequestable)() {
    if m != nil {
        m.templateType = value
    }
}
