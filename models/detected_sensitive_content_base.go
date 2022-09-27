package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContentBase 
type DetectedSensitiveContentBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The confidence property
    confidence *int32
    // The displayName property
    displayName *string
    // The id property
    id *string
    // The OdataType property
    odataType *string
    // The recommendedConfidence property
    recommendedConfidence *int32
    // The uniqueCount property
    uniqueCount *int32
}
// NewDetectedSensitiveContentBase instantiates a new detectedSensitiveContentBase and sets the default values.
func NewDetectedSensitiveContentBase()(*DetectedSensitiveContentBase) {
    m := &DetectedSensitiveContentBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.detectedSensitiveContentBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDetectedSensitiveContentBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetectedSensitiveContentBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.detectedSensitiveContent":
                        return NewDetectedSensitiveContent(), nil
                    case "#microsoft.graph.exactMatchDetectedSensitiveContent":
                        return NewExactMatchDetectedSensitiveContent(), nil
                    case "#microsoft.graph.machineLearningDetectedSensitiveContent":
                        return NewMachineLearningDetectedSensitiveContent(), nil
                }
            }
        }
    }
    return NewDetectedSensitiveContentBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetectedSensitiveContentBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfidence gets the confidence property value. The confidence property
func (m *DetectedSensitiveContentBase) GetConfidence()(*int32) {
    return m.confidence
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *DetectedSensitiveContentBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContentBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confidence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConfidence)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["recommendedConfidence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRecommendedConfidence)
    res["uniqueCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUniqueCount)
    return res
}
// GetId gets the id property value. The id property
func (m *DetectedSensitiveContentBase) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DetectedSensitiveContentBase) GetOdataType()(*string) {
    return m.odataType
}
// GetRecommendedConfidence gets the recommendedConfidence property value. The recommendedConfidence property
func (m *DetectedSensitiveContentBase) GetRecommendedConfidence()(*int32) {
    return m.recommendedConfidence
}
// GetUniqueCount gets the uniqueCount property value. The uniqueCount property
func (m *DetectedSensitiveContentBase) GetUniqueCount()(*int32) {
    return m.uniqueCount
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContentBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("recommendedConfidence", m.GetRecommendedConfidence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("uniqueCount", m.GetUniqueCount())
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
func (m *DetectedSensitiveContentBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfidence sets the confidence property value. The confidence property
func (m *DetectedSensitiveContentBase) SetConfidence(value *int32)() {
    m.confidence = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *DetectedSensitiveContentBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The id property
func (m *DetectedSensitiveContentBase) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DetectedSensitiveContentBase) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecommendedConfidence sets the recommendedConfidence property value. The recommendedConfidence property
func (m *DetectedSensitiveContentBase) SetRecommendedConfidence(value *int32)() {
    m.recommendedConfidence = value
}
// SetUniqueCount sets the uniqueCount property value. The uniqueCount property
func (m *DetectedSensitiveContentBase) SetUniqueCount(value *int32)() {
    m.uniqueCount = value
}
