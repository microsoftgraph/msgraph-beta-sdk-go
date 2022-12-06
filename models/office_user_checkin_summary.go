package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeUserCheckinSummary 
type OfficeUserCheckinSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Total failed user check ins for the last 3 months.
    failedUserCount *int32
    // The OdataType property
    odataType *string
    // Total successful user check ins for the last 3 months.
    succeededUserCount *int32
}
// NewOfficeUserCheckinSummary instantiates a new officeUserCheckinSummary and sets the default values.
func NewOfficeUserCheckinSummary()(*OfficeUserCheckinSummary) {
    m := &OfficeUserCheckinSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOfficeUserCheckinSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeUserCheckinSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeUserCheckinSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeUserCheckinSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFailedUserCount gets the failedUserCount property value. Total failed user check ins for the last 3 months.
func (m *OfficeUserCheckinSummary) GetFailedUserCount()(*int32) {
    return m.failedUserCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeUserCheckinSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedUserCount)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["succeededUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSucceededUserCount)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OfficeUserCheckinSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetSucceededUserCount gets the succeededUserCount property value. Total successful user check ins for the last 3 months.
func (m *OfficeUserCheckinSummary) GetSucceededUserCount()(*int32) {
    return m.succeededUserCount
}
// Serialize serializes information the current object
func (m *OfficeUserCheckinSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
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
        err := writer.WriteInt32Value("succeededUserCount", m.GetSucceededUserCount())
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
func (m *OfficeUserCheckinSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFailedUserCount sets the failedUserCount property value. Total failed user check ins for the last 3 months.
func (m *OfficeUserCheckinSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OfficeUserCheckinSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSucceededUserCount sets the succeededUserCount property value. Total successful user check ins for the last 3 months.
func (m *OfficeUserCheckinSummary) SetSucceededUserCount(value *int32)() {
    m.succeededUserCount = value
}
