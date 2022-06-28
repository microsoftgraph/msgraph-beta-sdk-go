package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessFlowSettings 
type BusinessFlowSettings struct {
    AccessReviewSettings
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The durationInDays property
    durationInDays *int32
}
// NewBusinessFlowSettings instantiates a new BusinessFlowSettings and sets the default values.
func NewBusinessFlowSettings()(*BusinessFlowSettings) {
    m := &BusinessFlowSettings{
        AccessReviewSettings: *NewAccessReviewSettings(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBusinessFlowSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessFlowSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessFlowSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDurationInDays gets the durationInDays property value. The durationInDays property
func (m *BusinessFlowSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlowSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewSettings.GetFieldDeserializers()
    res["durationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BusinessFlowSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessFlowSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDurationInDays sets the durationInDays property value. The durationInDays property
func (m *BusinessFlowSettings) SetDurationInDays(value *int32)() {
    if m != nil {
        m.durationInDays = value
    }
}
