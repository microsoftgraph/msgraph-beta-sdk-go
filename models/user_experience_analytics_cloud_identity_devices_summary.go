package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsCloudIdentityDevicesSummary the user experience analytics work from anywhere cloud identity devices summary.
type UserExperienceAnalyticsCloudIdentityDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The count of devices that are not cloud identity.
    deviceWithoutCloudIdentityCount *int32
    // The OdataType property
    odataType *string
}
// NewUserExperienceAnalyticsCloudIdentityDevicesSummary instantiates a new userExperienceAnalyticsCloudIdentityDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsCloudIdentityDevicesSummary()(*UserExperienceAnalyticsCloudIdentityDevicesSummary) {
    m := &UserExperienceAnalyticsCloudIdentityDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsCloudIdentityDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsCloudIdentityDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsCloudIdentityDevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceWithoutCloudIdentityCount gets the deviceWithoutCloudIdentityCount property value. The count of devices that are not cloud identity.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetDeviceWithoutCloudIdentityCount()(*int32) {
    return m.deviceWithoutCloudIdentityCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceWithoutCloudIdentityCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeviceWithoutCloudIdentityCount)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceWithoutCloudIdentityCount", m.GetDeviceWithoutCloudIdentityCount())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceWithoutCloudIdentityCount sets the deviceWithoutCloudIdentityCount property value. The count of devices that are not cloud identity.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) SetDeviceWithoutCloudIdentityCount(value *int32)() {
    m.deviceWithoutCloudIdentityCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
