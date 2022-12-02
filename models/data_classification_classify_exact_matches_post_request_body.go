package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationClassifyExactMatchesPostRequestBody provides operations to call the classifyExactMatches method.
type DataClassificationClassifyExactMatchesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentClassifications property
    contentClassifications []ContentClassificationable
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
    // The text property
    text *string
    // The timeoutInMs property
    timeoutInMs *string
}
// NewDataClassificationClassifyExactMatchesPostRequestBody instantiates a new DataClassificationClassifyExactMatchesPostRequestBody and sets the default values.
func NewDataClassificationClassifyExactMatchesPostRequestBody()(*DataClassificationClassifyExactMatchesPostRequestBody) {
    m := &DataClassificationClassifyExactMatchesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDataClassificationClassifyExactMatchesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataClassificationClassifyExactMatchesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataClassificationClassifyExactMatchesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataClassificationClassifyExactMatchesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentClassifications gets the contentClassifications property value. The contentClassifications property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) GetContentClassifications()([]ContentClassificationable) {
    return m.contentClassifications
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationClassifyExactMatchesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentClassifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContentClassificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentClassificationable, len(val))
            for i, v := range val {
                res[i] = v.(ContentClassificationable)
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
func (m *DataClassificationClassifyExactMatchesPostRequestBody) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// GetText gets the text property value. The text property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) GetText()(*string) {
    return m.text
}
// GetTimeoutInMs gets the timeoutInMs property value. The timeoutInMs property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) GetTimeoutInMs()(*string) {
    return m.timeoutInMs
}
// Serialize serializes information the current object
func (m *DataClassificationClassifyExactMatchesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DataClassificationClassifyExactMatchesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentClassifications sets the contentClassifications property value. The contentClassifications property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) SetContentClassifications(value []ContentClassificationable)() {
    m.contentClassifications = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
// SetText sets the text property value. The text property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) SetText(value *string)() {
    m.text = value
}
// SetTimeoutInMs sets the timeoutInMs property value. The timeoutInMs property
func (m *DataClassificationClassifyExactMatchesPostRequestBody) SetTimeoutInMs(value *string)() {
    m.timeoutInMs = value
}
