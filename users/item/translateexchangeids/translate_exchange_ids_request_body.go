package translateexchangeids

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TranslateExchangeIdsRequestBody provides operations to call the translateExchangeIds method.
type TranslateExchangeIdsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The InputIds property
    inputIds []string
    // The SourceIdType property
    sourceIdType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat
    // The TargetIdType property
    targetIdType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat
}
// NewTranslateExchangeIdsRequestBody instantiates a new translateExchangeIdsRequestBody and sets the default values.
func NewTranslateExchangeIdsRequestBody()(*TranslateExchangeIdsRequestBody) {
    m := &TranslateExchangeIdsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTranslateExchangeIdsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTranslateExchangeIdsRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTranslateExchangeIdsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslateExchangeIdsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslateExchangeIdsRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["inputIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["sourceIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat))
        }
        return nil
    }
    res["targetIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetIdType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat))
        }
        return nil
    }
    return res
}
// GetInputIds gets the inputIds property value. The InputIds property
func (m *TranslateExchangeIdsRequestBody) GetInputIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inputIds
    }
}
// GetSourceIdType gets the sourceIdType property value. The SourceIdType property
func (m *TranslateExchangeIdsRequestBody) GetSourceIdType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat) {
    if m == nil {
        return nil
    } else {
        return m.sourceIdType
    }
}
// GetTargetIdType gets the targetIdType property value. The TargetIdType property
func (m *TranslateExchangeIdsRequestBody) GetTargetIdType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat) {
    if m == nil {
        return nil
    } else {
        return m.targetIdType
    }
}
// Serialize serializes information the current object
func (m *TranslateExchangeIdsRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetInputIds sets the inputIds property value. The InputIds property
func (m *TranslateExchangeIdsRequestBody) SetInputIds(value []string)() {
    if m != nil {
        m.inputIds = value
    }
}
// SetSourceIdType sets the sourceIdType property value. The SourceIdType property
func (m *TranslateExchangeIdsRequestBody) SetSourceIdType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat)() {
    if m != nil {
        m.sourceIdType = value
    }
}
// SetTargetIdType sets the targetIdType property value. The TargetIdType property
func (m *TranslateExchangeIdsRequestBody) SetTargetIdType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeIdFormat)() {
    if m != nil {
        m.targetIdType = value
    }
}
