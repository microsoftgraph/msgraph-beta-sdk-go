package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalySeverityOverview the user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
type UserExperienceAnalyticsAnomalySeverityOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of high severity anomalies which have been detected. Valid values -2147483648 to 2147483647
    highSeverityAnomalyCount *int32
    // The number of informational severity anomalies which have been detected. Valid values -2147483648 to 2147483647
    informationalSeverityAnomalyCount *int32
    // The number of low severity anomalies which have been detected. Valid values -2147483648 to 2147483647
    lowSeverityAnomalyCount *int32
    // The number of medium severity anomalies which have been detected. Valid values -2147483648 to 2147483647
    mediumSeverityAnomalyCount *int32
    // The OdataType property
    odataType *string
    // The number of anomalies which have been detected with undefined severity. Valid values -2147483648 to 2147483647
    otherSeverityAnomalyCount *int32
}
// NewUserExperienceAnalyticsAnomalySeverityOverview instantiates a new userExperienceAnalyticsAnomalySeverityOverview and sets the default values.
func NewUserExperienceAnalyticsAnomalySeverityOverview()(*UserExperienceAnalyticsAnomalySeverityOverview) {
    m := &UserExperienceAnalyticsAnomalySeverityOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsAnomalySeverityOverview";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalySeverityOverview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["highSeverityAnomalyCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetHighSeverityAnomalyCount)
    res["informationalSeverityAnomalyCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetInformationalSeverityAnomalyCount)
    res["lowSeverityAnomalyCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetLowSeverityAnomalyCount)
    res["mediumSeverityAnomalyCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMediumSeverityAnomalyCount)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["otherSeverityAnomalyCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetOtherSeverityAnomalyCount)
    return res
}
// GetHighSeverityAnomalyCount gets the highSeverityAnomalyCount property value. The number of high severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetHighSeverityAnomalyCount()(*int32) {
    return m.highSeverityAnomalyCount
}
// GetInformationalSeverityAnomalyCount gets the informationalSeverityAnomalyCount property value. The number of informational severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetInformationalSeverityAnomalyCount()(*int32) {
    return m.informationalSeverityAnomalyCount
}
// GetLowSeverityAnomalyCount gets the lowSeverityAnomalyCount property value. The number of low severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetLowSeverityAnomalyCount()(*int32) {
    return m.lowSeverityAnomalyCount
}
// GetMediumSeverityAnomalyCount gets the mediumSeverityAnomalyCount property value. The number of medium severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetMediumSeverityAnomalyCount()(*int32) {
    return m.mediumSeverityAnomalyCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetOdataType()(*string) {
    return m.odataType
}
// GetOtherSeverityAnomalyCount gets the otherSeverityAnomalyCount property value. The number of anomalies which have been detected with undefined severity. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetOtherSeverityAnomalyCount()(*int32) {
    return m.otherSeverityAnomalyCount
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalySeverityOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("highSeverityAnomalyCount", m.GetHighSeverityAnomalyCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("informationalSeverityAnomalyCount", m.GetInformationalSeverityAnomalyCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("lowSeverityAnomalyCount", m.GetLowSeverityAnomalyCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("mediumSeverityAnomalyCount", m.GetMediumSeverityAnomalyCount())
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
        err := writer.WriteInt32Value("otherSeverityAnomalyCount", m.GetOtherSeverityAnomalyCount())
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
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHighSeverityAnomalyCount sets the highSeverityAnomalyCount property value. The number of high severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetHighSeverityAnomalyCount(value *int32)() {
    m.highSeverityAnomalyCount = value
}
// SetInformationalSeverityAnomalyCount sets the informationalSeverityAnomalyCount property value. The number of informational severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetInformationalSeverityAnomalyCount(value *int32)() {
    m.informationalSeverityAnomalyCount = value
}
// SetLowSeverityAnomalyCount sets the lowSeverityAnomalyCount property value. The number of low severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetLowSeverityAnomalyCount(value *int32)() {
    m.lowSeverityAnomalyCount = value
}
// SetMediumSeverityAnomalyCount sets the mediumSeverityAnomalyCount property value. The number of medium severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetMediumSeverityAnomalyCount(value *int32)() {
    m.mediumSeverityAnomalyCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOtherSeverityAnomalyCount sets the otherSeverityAnomalyCount property value. The number of anomalies which have been detected with undefined severity. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetOtherSeverityAnomalyCount(value *int32)() {
    m.otherSeverityAnomalyCount = value
}
