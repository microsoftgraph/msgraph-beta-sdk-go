package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody provides operations to call the evaluate method.
type DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentLabel property
    currentLabel CurrentLabelable
    // The discoveredSensitiveTypes property
    discoveredSensitiveTypes []DiscoveredSensitiveTypeable
}
// NewDataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody instantiates a new DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody and sets the default values.
func NewDataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody()(*DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) {
    m := &DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetCurrentLabel()(CurrentLabelable) {
    return m.currentLabel
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveTypeable) {
    return m.discoveredSensitiveTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(CurrentLabelable))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredSensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(DiscoveredSensitiveTypeable)
            }
            m.SetDiscoveredSensitiveTypes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
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
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetCurrentLabel(value CurrentLabelable)() {
    m.currentLabel = value
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DataClassificationSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveTypeable)() {
    m.discoveredSensitiveTypes = value
}
