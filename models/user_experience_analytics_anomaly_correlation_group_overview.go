package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsAnomalyCorrelationGroupOverview the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
type UserExperienceAnalyticsAnomalyCorrelationGroupOverview struct {
    Entity
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage instantiates a new UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage and sets the default values.
func NewUserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage()(*UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) {
    m := &UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewUserExperienceAnalyticsAnomalyCorrelationGroupOverview instantiates a new UserExperienceAnalyticsAnomalyCorrelationGroupOverview and sets the default values.
func NewUserExperienceAnalyticsAnomalyCorrelationGroupOverview()(*UserExperienceAnalyticsAnomalyCorrelationGroupOverview) {
    m := &UserExperienceAnalyticsAnomalyCorrelationGroupOverview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalyCorrelationGroupOverview(), nil
}
// GetAnomalyCorrelationGroupCount gets the anomalyCorrelationGroupCount property value. Indicates the number of correlation groups in the anomaly. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetAnomalyCorrelationGroupCount()(*int32) {
    val, err := m.GetBackingStore().Get("anomalyCorrelationGroupCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAnomalyId gets the anomalyId property value. The unique identifier of the anomaly. Anomaly details such as name and type can be found in the UserExperienceAnalyticsAnomalySeverityOverview entity.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetAnomalyId()(*string) {
    val, err := m.GetBackingStore().Get("anomalyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCorrelationGroupAnomalousDeviceCount gets the correlationGroupAnomalousDeviceCount property value. Indicates the total number of devices affected by the anomaly in the correlation group. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupAnomalousDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("correlationGroupAnomalousDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCorrelationGroupAtRiskDeviceCount gets the correlationGroupAtRiskDeviceCount property value. Indicates the total number of devices at risk in the correlation group. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupAtRiskDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("correlationGroupAtRiskDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCorrelationGroupDeviceCount gets the correlationGroupDeviceCount property value. Indicates the total number of devices in a correlation group. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("correlationGroupDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCorrelationGroupFeatures gets the correlationGroupFeatures property value. Describes the features of a device that are shared between all devices in a correlation group.
// returns a []UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupFeatures()([]UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable) {
    val, err := m.GetBackingStore().Get("correlationGroupFeatures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable)
    }
    return nil
}
// GetCorrelationGroupId gets the correlationGroupId property value. The unique identifier for the correlation group which will uniquely identify one of the correlation group within an anomaly. The correlation Id can be mapped to the correlation group name by concatinating the correlation group features. Example of correlation group name which is the indicative of concatenated features names are  for names, Contoso manufacture 4.4.1 and Windows 11.22621.1485.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupId()(*string) {
    val, err := m.GetBackingStore().Get("correlationGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCorrelationGroupPrevalence gets the correlationGroupPrevalence property value. Indicates the level of prevalence of the correlation group features in the anomaly. Possible values are: high, medium or low
// returns a *UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupPrevalence()(*UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence) {
    val, err := m.GetBackingStore().Get("correlationGroupPrevalence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence)
    }
    return nil
}
// GetCorrelationGroupPrevalencePercentage gets the correlationGroupPrevalencePercentage property value. The percentage of the devices in the correlation group that are anomalous. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetCorrelationGroupPrevalencePercentage()(UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable) {
    val, err := m.GetBackingStore().Get("correlationGroupPrevalencePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["anomalyCorrelationGroupCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyCorrelationGroupCount(val)
        }
        return nil
    }
    res["anomalyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyId(val)
        }
        return nil
    }
    res["correlationGroupAnomalousDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupAnomalousDeviceCount(val)
        }
        return nil
    }
    res["correlationGroupAtRiskDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupAtRiskDeviceCount(val)
        }
        return nil
    }
    res["correlationGroupDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupDeviceCount(val)
        }
        return nil
    }
    res["correlationGroupFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyCorrelationGroupFeatureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable)
                }
            }
            m.SetCorrelationGroupFeatures(res)
        }
        return nil
    }
    res["correlationGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupId(val)
        }
        return nil
    }
    res["correlationGroupPrevalence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalyCorrelationGroupPrevalence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupPrevalence(val.(*UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence))
        }
        return nil
    }
    res["correlationGroupPrevalencePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupPrevalencePercentage(val.(*UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentage))
        }
        return nil
    }
    res["totalDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetTotalDeviceCount gets the totalDeviceCount property value. Indicates the total number of devices in the tenant. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) GetTotalDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("anomalyCorrelationGroupCount", m.GetAnomalyCorrelationGroupCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("anomalyId", m.GetAnomalyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("correlationGroupAnomalousDeviceCount", m.GetCorrelationGroupAnomalousDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("correlationGroupAtRiskDeviceCount", m.GetCorrelationGroupAtRiskDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("correlationGroupDeviceCount", m.GetCorrelationGroupDeviceCount())
        if err != nil {
            return err
        }
    }
    if m.GetCorrelationGroupFeatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCorrelationGroupFeatures()))
        for i, v := range m.GetCorrelationGroupFeatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("correlationGroupFeatures", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationGroupId", m.GetCorrelationGroupId())
        if err != nil {
            return err
        }
    }
    if m.GetCorrelationGroupPrevalence() != nil {
        cast := (*m.GetCorrelationGroupPrevalence()).String()
        err = writer.WriteStringValue("correlationGroupPrevalence", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("correlationGroupPrevalencePercentage", m.GetCorrelationGroupPrevalencePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalDeviceCount", m.GetTotalDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnomalyCorrelationGroupCount sets the anomalyCorrelationGroupCount property value. Indicates the number of correlation groups in the anomaly. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetAnomalyCorrelationGroupCount(value *int32)() {
    err := m.GetBackingStore().Set("anomalyCorrelationGroupCount", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyId sets the anomalyId property value. The unique identifier of the anomaly. Anomaly details such as name and type can be found in the UserExperienceAnalyticsAnomalySeverityOverview entity.
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetAnomalyId(value *string)() {
    err := m.GetBackingStore().Set("anomalyId", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupAnomalousDeviceCount sets the correlationGroupAnomalousDeviceCount property value. Indicates the total number of devices affected by the anomaly in the correlation group. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupAnomalousDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("correlationGroupAnomalousDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupAtRiskDeviceCount sets the correlationGroupAtRiskDeviceCount property value. Indicates the total number of devices at risk in the correlation group. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupAtRiskDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("correlationGroupAtRiskDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupDeviceCount sets the correlationGroupDeviceCount property value. Indicates the total number of devices in a correlation group. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("correlationGroupDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupFeatures sets the correlationGroupFeatures property value. Describes the features of a device that are shared between all devices in a correlation group.
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupFeatures(value []UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable)() {
    err := m.GetBackingStore().Set("correlationGroupFeatures", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupId sets the correlationGroupId property value. The unique identifier for the correlation group which will uniquely identify one of the correlation group within an anomaly. The correlation Id can be mapped to the correlation group name by concatinating the correlation group features. Example of correlation group name which is the indicative of concatenated features names are  for names, Contoso manufacture 4.4.1 and Windows 11.22621.1485.
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupId(value *string)() {
    err := m.GetBackingStore().Set("correlationGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupPrevalence sets the correlationGroupPrevalence property value. Indicates the level of prevalence of the correlation group features in the anomaly. Possible values are: high, medium or low
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupPrevalence(value *UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence)() {
    err := m.GetBackingStore().Set("correlationGroupPrevalence", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupPrevalencePercentage sets the correlationGroupPrevalencePercentage property value. The percentage of the devices in the correlation group that are anomalous. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetCorrelationGroupPrevalencePercentage(value UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable)() {
    err := m.GetBackingStore().Set("correlationGroupPrevalencePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDeviceCount sets the totalDeviceCount property value. Indicates the total number of devices in the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverview) SetTotalDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("totalDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnomalyCorrelationGroupCount()(*int32)
    GetAnomalyId()(*string)
    GetCorrelationGroupAnomalousDeviceCount()(*int32)
    GetCorrelationGroupAtRiskDeviceCount()(*int32)
    GetCorrelationGroupDeviceCount()(*int32)
    GetCorrelationGroupFeatures()([]UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable)
    GetCorrelationGroupId()(*string)
    GetCorrelationGroupPrevalence()(*UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence)
    GetCorrelationGroupPrevalencePercentage()(UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable)
    GetTotalDeviceCount()(*int32)
    SetAnomalyCorrelationGroupCount(value *int32)()
    SetAnomalyId(value *string)()
    SetCorrelationGroupAnomalousDeviceCount(value *int32)()
    SetCorrelationGroupAtRiskDeviceCount(value *int32)()
    SetCorrelationGroupDeviceCount(value *int32)()
    SetCorrelationGroupFeatures(value []UserExperienceAnalyticsAnomalyCorrelationGroupFeatureable)()
    SetCorrelationGroupId(value *string)()
    SetCorrelationGroupPrevalence(value *UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence)()
    SetCorrelationGroupPrevalencePercentage(value UserExperienceAnalyticsAnomalyCorrelationGroupOverview_UserExperienceAnalyticsAnomalyCorrelationGroupOverview_correlationGroupPrevalencePercentageable)()
    SetTotalDeviceCount(value *int32)()
}
