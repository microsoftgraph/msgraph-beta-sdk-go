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
    // The recommendedConfidence property
    recommendedConfidence *int32
    // The type property
    type_escaped *string
    // The uniqueCount property
    uniqueCount *int32
}
// NewDetectedSensitiveContentBase instantiates a new detectedSensitiveContentBase and sets the default values.
func NewDetectedSensitiveContentBase()(*DetectedSensitiveContentBase) {
    m := &DetectedSensitiveContentBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    typeValue := "#microsoft.graph.detectedSensitiveContentBase";
    m.SetType(&typeValue);
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.detectedSensitiveContent":
                        return NewDetectedSensitiveContent(), nil
                    case "#microsoft.graph.exactMatchDetectedSensitiveContent":
                        return NewExactMatchDetectedSensitiveContent(), nil
                }
            }
        }
    }
    return NewDetectedSensitiveContentBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetectedSensitiveContentBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfidence gets the confidence property value. The confidence property
func (m *DetectedSensitiveContentBase) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *DetectedSensitiveContentBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContentBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confidence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["recommendedConfidence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedConfidence(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["uniqueCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueCount(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *DetectedSensitiveContentBase) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetRecommendedConfidence gets the recommendedConfidence property value. The recommendedConfidence property
func (m *DetectedSensitiveContentBase) GetRecommendedConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.recommendedConfidence
    }
}
// GetType gets the type property value. The type property
func (m *DetectedSensitiveContentBase) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUniqueCount gets the uniqueCount property value. The uniqueCount property
func (m *DetectedSensitiveContentBase) GetUniqueCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uniqueCount
    }
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
        err := writer.WriteInt32Value("recommendedConfidence", m.GetRecommendedConfidence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
    if m != nil {
        m.additionalData = value
    }
}
// SetConfidence sets the confidence property value. The confidence property
func (m *DetectedSensitiveContentBase) SetConfidence(value *int32)() {
    if m != nil {
        m.confidence = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *DetectedSensitiveContentBase) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. The id property
func (m *DetectedSensitiveContentBase) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetRecommendedConfidence sets the recommendedConfidence property value. The recommendedConfidence property
func (m *DetectedSensitiveContentBase) SetRecommendedConfidence(value *int32)() {
    if m != nil {
        m.recommendedConfidence = value
    }
}
// SetType sets the type property value. The type property
func (m *DetectedSensitiveContentBase) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUniqueCount sets the uniqueCount property value. The uniqueCount property
func (m *DetectedSensitiveContentBase) SetUniqueCount(value *int32)() {
    if m != nil {
        m.uniqueCount = value
    }
}
