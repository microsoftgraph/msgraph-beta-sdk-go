package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SuggestedEnrollmentLimit the suggestedEnrollmentLimit resource represents the suggested enrollment limit when given an enrollment type.
type SuggestedEnrollmentLimit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The suggested enrollment limit within a day
    suggestedDailyLimit *int32
}
// NewSuggestedEnrollmentLimit instantiates a new suggestedEnrollmentLimit and sets the default values.
func NewSuggestedEnrollmentLimit()(*SuggestedEnrollmentLimit) {
    m := &SuggestedEnrollmentLimit{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateSuggestedEnrollmentLimitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSuggestedEnrollmentLimitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSuggestedEnrollmentLimit(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SuggestedEnrollmentLimit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SuggestedEnrollmentLimit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["suggestedDailyLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestedDailyLimit(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SuggestedEnrollmentLimit) GetOdataType()(*string) {
    return m.odataType
}
// GetSuggestedDailyLimit gets the suggestedDailyLimit property value. The suggested enrollment limit within a day
func (m *SuggestedEnrollmentLimit) GetSuggestedDailyLimit()(*int32) {
    return m.suggestedDailyLimit
}
// Serialize serializes information the current object
func (m *SuggestedEnrollmentLimit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("suggestedDailyLimit", m.GetSuggestedDailyLimit())
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
func (m *SuggestedEnrollmentLimit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SuggestedEnrollmentLimit) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuggestedDailyLimit sets the suggestedDailyLimit property value. The suggested enrollment limit within a day
func (m *SuggestedEnrollmentLimit) SetSuggestedDailyLimit(value *int32)() {
    m.suggestedDailyLimit = value
}
