package classifyexactmatches

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ClassifyExactMatchesRequestBody provides operations to call the classifyExactMatches method.
type ClassifyExactMatchesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentClassifications property
    contentClassifications []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
    // The text property
    text *string
    // The timeoutInMs property
    timeoutInMs *string
}
// NewClassifyExactMatchesRequestBody instantiates a new classifyExactMatchesRequestBody and sets the default values.
func NewClassifyExactMatchesRequestBody()(*ClassifyExactMatchesRequestBody) {
    m := &ClassifyExactMatchesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassifyExactMatchesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassifyExactMatchesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassifyExactMatchesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyExactMatchesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentClassifications gets the contentClassifications property value. The contentClassifications property
func (m *ClassifyExactMatchesRequestBody) GetContentClassifications()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable) {
    if m == nil {
        return nil
    } else {
        return m.contentClassifications
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyExactMatchesRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentClassifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentClassificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable)
            }
            m.SetContentClassifications(res)
        }
        return nil
    }
    res["sensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["timeoutInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeoutInMs(val)
        }
        return nil
    }
    return res
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *ClassifyExactMatchesRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// GetText gets the text property value. The text property
func (m *ClassifyExactMatchesRequestBody) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetTimeoutInMs gets the timeoutInMs property value. The timeoutInMs property
func (m *ClassifyExactMatchesRequestBody) GetTimeoutInMs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeoutInMs
    }
}
// Serialize serializes information the current object
func (m *ClassifyExactMatchesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentClassifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentClassifications()))
        for i, v := range m.GetContentClassifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("contentClassifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err := writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeoutInMs", m.GetTimeoutInMs())
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
func (m *ClassifyExactMatchesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentClassifications sets the contentClassifications property value. The contentClassifications property
func (m *ClassifyExactMatchesRequestBody) SetContentClassifications(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable)() {
    if m != nil {
        m.contentClassifications = value
    }
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *ClassifyExactMatchesRequestBody) SetSensitiveTypeIds(value []string)() {
    if m != nil {
        m.sensitiveTypeIds = value
    }
}
// SetText sets the text property value. The text property
func (m *ClassifyExactMatchesRequestBody) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
// SetTimeoutInMs sets the timeoutInMs property value. The timeoutInMs property
func (m *ClassifyExactMatchesRequestBody) SetTimeoutInMs(value *string)() {
    if m != nil {
        m.timeoutInMs = value
    }
}
