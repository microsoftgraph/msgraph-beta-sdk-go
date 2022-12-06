package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveContentLocation 
type SensitiveContentLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The confidence property
    confidence *int32
    // The evidences property
    evidences []SensitiveContentEvidenceable
    // The idMatch property
    idMatch *string
    // The length property
    length *int32
    // The OdataType property
    odataType *string
    // The offset property
    offset *int32
}
// NewSensitiveContentLocation instantiates a new sensitiveContentLocation and sets the default values.
func NewSensitiveContentLocation()(*SensitiveContentLocation) {
    m := &SensitiveContentLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSensitiveContentLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitiveContentLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitiveContentLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentLocation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfidence gets the confidence property value. The confidence property
func (m *SensitiveContentLocation) GetConfidence()(*int32) {
    return m.confidence
}
// GetEvidences gets the evidences property value. The evidences property
func (m *SensitiveContentLocation) GetEvidences()([]SensitiveContentEvidenceable) {
    return m.evidences
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveContentLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confidence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConfidence)
    res["evidences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitiveContentEvidenceFromDiscriminatorValue , m.SetEvidences)
    res["idMatch"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIdMatch)
    res["length"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetLength)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["offset"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetOffset)
    return res
}
// GetIdMatch gets the idMatch property value. The idMatch property
func (m *SensitiveContentLocation) GetIdMatch()(*string) {
    return m.idMatch
}
// GetLength gets the length property value. The length property
func (m *SensitiveContentLocation) GetLength()(*int32) {
    return m.length
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SensitiveContentLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetOffset gets the offset property value. The offset property
func (m *SensitiveContentLocation) GetOffset()(*int32) {
    return m.offset
}
// Serialize serializes information the current object
func (m *SensitiveContentLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    if m.GetEvidences() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEvidences())
        err := writer.WriteCollectionOfObjectValues("evidences", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("idMatch", m.GetIdMatch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("length", m.GetLength())
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
        err := writer.WriteInt32Value("offset", m.GetOffset())
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
func (m *SensitiveContentLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfidence sets the confidence property value. The confidence property
func (m *SensitiveContentLocation) SetConfidence(value *int32)() {
    m.confidence = value
}
// SetEvidences sets the evidences property value. The evidences property
func (m *SensitiveContentLocation) SetEvidences(value []SensitiveContentEvidenceable)() {
    m.evidences = value
}
// SetIdMatch sets the idMatch property value. The idMatch property
func (m *SensitiveContentLocation) SetIdMatch(value *string)() {
    m.idMatch = value
}
// SetLength sets the length property value. The length property
func (m *SensitiveContentLocation) SetLength(value *int32)() {
    m.length = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SensitiveContentLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffset sets the offset property value. The offset property
func (m *SensitiveContentLocation) SetOffset(value *int32)() {
    m.offset = value
}
