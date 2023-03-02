package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsAnomalySeverityOverview the user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
type UserExperienceAnalyticsAnomalySeverityOverview struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsAnomalySeverityOverview instantiates a new userExperienceAnalyticsAnomalySeverityOverview and sets the default values.
func NewUserExperienceAnalyticsAnomalySeverityOverview()(*UserExperienceAnalyticsAnomalySeverityOverview) {
    m := &UserExperienceAnalyticsAnomalySeverityOverview{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalySeverityOverview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["highSeverityAnomalyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighSeverityAnomalyCount(val)
        }
        return nil
    }
    res["informationalSeverityAnomalyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationalSeverityAnomalyCount(val)
        }
        return nil
    }
    res["lowSeverityAnomalyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowSeverityAnomalyCount(val)
        }
        return nil
    }
    res["mediumSeverityAnomalyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediumSeverityAnomalyCount(val)
        }
        return nil
    }
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
    return res
}
// GetHighSeverityAnomalyCount gets the highSeverityAnomalyCount property value. Indicates count of high severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetHighSeverityAnomalyCount()(*int32) {
    val, err := m.GetBackingStore().Get("highSeverityAnomalyCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetInformationalSeverityAnomalyCount gets the informationalSeverityAnomalyCount property value. Indicates count of informational severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetInformationalSeverityAnomalyCount()(*int32) {
    val, err := m.GetBackingStore().Get("informationalSeverityAnomalyCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLowSeverityAnomalyCount gets the lowSeverityAnomalyCount property value. Indicates count of low severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetLowSeverityAnomalyCount()(*int32) {
    val, err := m.GetBackingStore().Get("lowSeverityAnomalyCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMediumSeverityAnomalyCount gets the mediumSeverityAnomalyCount property value. Indicates count of medium severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetMediumSeverityAnomalyCount()(*int32) {
    val, err := m.GetBackingStore().Get("mediumSeverityAnomalyCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsAnomalySeverityOverview) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetHighSeverityAnomalyCount sets the highSeverityAnomalyCount property value. Indicates count of high severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetHighSeverityAnomalyCount(value *int32)() {
    err := m.GetBackingStore().Set("highSeverityAnomalyCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInformationalSeverityAnomalyCount sets the informationalSeverityAnomalyCount property value. Indicates count of informational severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetInformationalSeverityAnomalyCount(value *int32)() {
    err := m.GetBackingStore().Set("informationalSeverityAnomalyCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLowSeverityAnomalyCount sets the lowSeverityAnomalyCount property value. Indicates count of low severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetLowSeverityAnomalyCount(value *int32)() {
    err := m.GetBackingStore().Set("lowSeverityAnomalyCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMediumSeverityAnomalyCount sets the mediumSeverityAnomalyCount property value. Indicates count of medium severity anomalies which have been detected. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetMediumSeverityAnomalyCount(value *int32)() {
    err := m.GetBackingStore().Set("mediumSeverityAnomalyCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsAnomalySeverityOverview) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsAnomalySeverityOverviewable 
type UserExperienceAnalyticsAnomalySeverityOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetHighSeverityAnomalyCount()(*int32)
    GetInformationalSeverityAnomalyCount()(*int32)
    GetLowSeverityAnomalyCount()(*int32)
    GetMediumSeverityAnomalyCount()(*int32)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetHighSeverityAnomalyCount(value *int32)()
    SetInformationalSeverityAnomalyCount(value *int32)()
    SetLowSeverityAnomalyCount(value *int32)()
    SetMediumSeverityAnomalyCount(value *int32)()
    SetOdataType(value *string)()
}
