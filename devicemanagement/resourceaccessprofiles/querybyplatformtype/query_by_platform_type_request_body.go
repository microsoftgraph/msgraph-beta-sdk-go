package querybyplatformtype

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// QueryByPlatformTypeRequestBody provides operations to call the queryByPlatformType method.
type QueryByPlatformTypeRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    platformType *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicyPlatformType;
}
// NewQueryByPlatformTypeRequestBody instantiates a new queryByPlatformTypeRequestBody and sets the default values.
func NewQueryByPlatformTypeRequestBody()(*QueryByPlatformTypeRequestBody) {
    m := &QueryByPlatformTypeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateQueryByPlatformTypeRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQueryByPlatformTypeRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewQueryByPlatformTypeRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *QueryByPlatformTypeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *QueryByPlatformTypeRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicyPlatformType))
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. 
func (m *QueryByPlatformTypeRequestBody) GetPlatformType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
func (m *QueryByPlatformTypeRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *QueryByPlatformTypeRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err := writer.WriteStringValue("platformType", &cast)
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
func (m *QueryByPlatformTypeRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPlatformType sets the platformType property value. 
func (m *QueryByPlatformTypeRequestBody) SetPlatformType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicyPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
