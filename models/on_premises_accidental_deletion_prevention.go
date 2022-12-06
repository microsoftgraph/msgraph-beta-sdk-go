package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAccidentalDeletionPrevention 
type OnPremisesAccidentalDeletionPrevention struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The alertThreshold property
    alertThreshold *int32
    // The OdataType property
    odataType *string
    // The synchronizationPreventionType property
    synchronizationPreventionType *OnPremisesDirectorySynchronizationDeletionPreventionType
}
// NewOnPremisesAccidentalDeletionPrevention instantiates a new onPremisesAccidentalDeletionPrevention and sets the default values.
func NewOnPremisesAccidentalDeletionPrevention()(*OnPremisesAccidentalDeletionPrevention) {
    m := &OnPremisesAccidentalDeletionPrevention{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesAccidentalDeletionPreventionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesAccidentalDeletionPreventionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesAccidentalDeletionPrevention(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesAccidentalDeletionPrevention) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAlertThreshold gets the alertThreshold property value. The alertThreshold property
func (m *OnPremisesAccidentalDeletionPrevention) GetAlertThreshold()(*int32) {
    return m.alertThreshold
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesAccidentalDeletionPrevention) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertThreshold"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAlertThreshold)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["synchronizationPreventionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOnPremisesDirectorySynchronizationDeletionPreventionType , m.SetSynchronizationPreventionType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesAccidentalDeletionPrevention) GetOdataType()(*string) {
    return m.odataType
}
// GetSynchronizationPreventionType gets the synchronizationPreventionType property value. The synchronizationPreventionType property
func (m *OnPremisesAccidentalDeletionPrevention) GetSynchronizationPreventionType()(*OnPremisesDirectorySynchronizationDeletionPreventionType) {
    return m.synchronizationPreventionType
}
// Serialize serializes information the current object
func (m *OnPremisesAccidentalDeletionPrevention) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("alertThreshold", m.GetAlertThreshold())
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
    if m.GetSynchronizationPreventionType() != nil {
        cast := (*m.GetSynchronizationPreventionType()).String()
        err := writer.WriteStringValue("synchronizationPreventionType", &cast)
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
func (m *OnPremisesAccidentalDeletionPrevention) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAlertThreshold sets the alertThreshold property value. The alertThreshold property
func (m *OnPremisesAccidentalDeletionPrevention) SetAlertThreshold(value *int32)() {
    m.alertThreshold = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesAccidentalDeletionPrevention) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSynchronizationPreventionType sets the synchronizationPreventionType property value. The synchronizationPreventionType property
func (m *OnPremisesAccidentalDeletionPrevention) SetSynchronizationPreventionType(value *OnPremisesDirectorySynchronizationDeletionPreventionType)() {
    m.synchronizationPreventionType = value
}
