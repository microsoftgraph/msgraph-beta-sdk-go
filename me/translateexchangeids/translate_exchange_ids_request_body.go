package translateexchangeids

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// TranslateExchangeIdsRequestBody 
type TranslateExchangeIdsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    inputIds []string;
    // 
    sourceIdType *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat;
    // 
    targetIdType *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat;
}
// NewTranslateExchangeIdsRequestBody instantiates a new translateExchangeIdsRequestBody and sets the default values.
func NewTranslateExchangeIdsRequestBody()(*TranslateExchangeIdsRequestBody) {
    m := &TranslateExchangeIdsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslateExchangeIdsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetInputIds gets the inputIds property value. 
func (m *TranslateExchangeIdsRequestBody) GetInputIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inputIds
    }
}
// GetSourceIdType gets the sourceIdType property value. 
func (m *TranslateExchangeIdsRequestBody) GetSourceIdType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat) {
    if m == nil {
        return nil
    } else {
        return m.sourceIdType
    }
}
// GetTargetIdType gets the targetIdType property value. 
func (m *TranslateExchangeIdsRequestBody) GetTargetIdType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat) {
    if m == nil {
        return nil
    } else {
        return m.targetIdType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslateExchangeIdsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["inputIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInputIds(res)
        }
        return nil
    }
    res["sourceIdType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat))
        }
        return nil
    }
    res["targetIdType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetIdType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat))
        }
        return nil
    }
    return res
}
func (m *TranslateExchangeIdsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TranslateExchangeIdsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetInputIds() != nil {
        err := writer.WriteCollectionOfStringValues("inputIds", m.GetInputIds())
        if err != nil {
            return err
        }
    }
    if m.GetSourceIdType() != nil {
        cast := (*m.GetSourceIdType()).String()
        err := writer.WriteStringValue("sourceIdType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetIdType() != nil {
        cast := (*m.GetTargetIdType()).String()
        err := writer.WriteStringValue("targetIdType", &cast)
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
func (m *TranslateExchangeIdsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetInputIds sets the inputIds property value. 
func (m *TranslateExchangeIdsRequestBody) SetInputIds(value []string)() {
    if m != nil {
        m.inputIds = value
    }
}
// SetSourceIdType sets the sourceIdType property value. 
func (m *TranslateExchangeIdsRequestBody) SetSourceIdType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat)() {
    if m != nil {
        m.sourceIdType = value
    }
}
// SetTargetIdType sets the targetIdType property value. 
func (m *TranslateExchangeIdsRequestBody) SetTargetIdType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExchangeIdFormat)() {
    if m != nil {
        m.targetIdType = value
    }
}
