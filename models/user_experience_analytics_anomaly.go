package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomaly 
type UserExperienceAnalyticsAnomaly struct {
    Entity
}
// NewUserExperienceAnalyticsAnomaly instantiates a new UserExperienceAnalyticsAnomaly and sets the default values.
func NewUserExperienceAnalyticsAnomaly()(*UserExperienceAnalyticsAnomaly) {
    m := &UserExperienceAnalyticsAnomaly{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomaly(), nil
}
// GetAnomalyFirstOccurrenceDateTime gets the anomalyFirstOccurrenceDateTime property value. Indicates the first occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("anomalyFirstOccurrenceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAnomalyId gets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyId()(*string) {
    val, err := m.GetBackingStore().Get("anomalyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAnomalyLatestOccurrenceDateTime gets the anomalyLatestOccurrenceDateTime property value. Indicates the latest occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("anomalyLatestOccurrenceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAnomalyName gets the anomalyName property value. The name of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyName()(*string) {
    val, err := m.GetBackingStore().Get("anomalyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAnomalyType gets the anomalyType property value. Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyType()(*UserExperienceAnalyticsAnomalyType) {
    val, err := m.GetBackingStore().Get("anomalyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsAnomalyType)
    }
    return nil
}
// GetAssetName gets the assetName property value. The name of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAssetName()(*string) {
    val, err := m.GetBackingStore().Get("assetName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssetPublisher gets the assetPublisher property value. The publisher of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAssetPublisher()(*string) {
    val, err := m.GetBackingStore().Get("assetPublisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssetVersion gets the assetVersion property value. The version of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAssetVersion()(*string) {
    val, err := m.GetBackingStore().Get("assetVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetectionModelId gets the detectionModelId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) GetDetectionModelId()(*string) {
    val, err := m.GetBackingStore().Get("detectionModelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceImpactedCount gets the deviceImpactedCount property value. The number of devices impacted by the anomaly. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomaly) GetDeviceImpactedCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceImpactedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomaly) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["anomalyFirstOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyFirstOccurrenceDateTime(val)
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
    res["anomalyLatestOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyLatestOccurrenceDateTime(val)
        }
        return nil
    }
    res["anomalyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyName(val)
        }
        return nil
    }
    res["anomalyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyType(val.(*UserExperienceAnalyticsAnomalyType))
        }
        return nil
    }
    res["assetName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetName(val)
        }
        return nil
    }
    res["assetPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetPublisher(val)
        }
        return nil
    }
    res["assetVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetVersion(val)
        }
        return nil
    }
    res["detectionModelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionModelId(val)
        }
        return nil
    }
    res["deviceImpactedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceImpactedCount(val)
        }
        return nil
    }
    res["issueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueId(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalySeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*UserExperienceAnalyticsAnomalySeverity))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalyState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*UserExperienceAnalyticsAnomalyState))
        }
        return nil
    }
    return res
}
// GetIssueId gets the issueId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) GetIssueId()(*string) {
    val, err := m.GetBackingStore().Get("issueId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSeverity gets the severity property value. Indicates the severity of the anomaly. Eg: anomaly severity can be high, medium, low, informational or other.
func (m *UserExperienceAnalyticsAnomaly) GetSeverity()(*UserExperienceAnalyticsAnomalySeverity) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsAnomalySeverity)
    }
    return nil
}
// GetState gets the state property value. Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
func (m *UserExperienceAnalyticsAnomaly) GetState()(*UserExperienceAnalyticsAnomalyState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsAnomalyState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomaly) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("anomalyFirstOccurrenceDateTime", m.GetAnomalyFirstOccurrenceDateTime())
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
        err = writer.WriteTimeValue("anomalyLatestOccurrenceDateTime", m.GetAnomalyLatestOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("anomalyName", m.GetAnomalyName())
        if err != nil {
            return err
        }
    }
    if m.GetAnomalyType() != nil {
        cast := (*m.GetAnomalyType()).String()
        err = writer.WriteStringValue("anomalyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assetName", m.GetAssetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assetPublisher", m.GetAssetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assetVersion", m.GetAssetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("detectionModelId", m.GetDetectionModelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceImpactedCount", m.GetDeviceImpactedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issueId", m.GetIssueId())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnomalyFirstOccurrenceDateTime sets the anomalyFirstOccurrenceDateTime property value. Indicates the first occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("anomalyFirstOccurrenceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyId sets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyId(value *string)() {
    err := m.GetBackingStore().Set("anomalyId", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyLatestOccurrenceDateTime sets the anomalyLatestOccurrenceDateTime property value. Indicates the latest occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("anomalyLatestOccurrenceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyName sets the anomalyName property value. The name of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyName(value *string)() {
    err := m.GetBackingStore().Set("anomalyName", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyType sets the anomalyType property value. Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyType(value *UserExperienceAnalyticsAnomalyType)() {
    err := m.GetBackingStore().Set("anomalyType", value)
    if err != nil {
        panic(err)
    }
}
// SetAssetName sets the assetName property value. The name of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAssetName(value *string)() {
    err := m.GetBackingStore().Set("assetName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssetPublisher sets the assetPublisher property value. The publisher of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAssetPublisher(value *string)() {
    err := m.GetBackingStore().Set("assetPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetAssetVersion sets the assetVersion property value. The version of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAssetVersion(value *string)() {
    err := m.GetBackingStore().Set("assetVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionModelId sets the detectionModelId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) SetDetectionModelId(value *string)() {
    err := m.GetBackingStore().Set("detectionModelId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceImpactedCount sets the deviceImpactedCount property value. The number of devices impacted by the anomaly. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomaly) SetDeviceImpactedCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceImpactedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIssueId sets the issueId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) SetIssueId(value *string)() {
    err := m.GetBackingStore().Set("issueId", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. Indicates the severity of the anomaly. Eg: anomaly severity can be high, medium, low, informational or other.
func (m *UserExperienceAnalyticsAnomaly) SetSeverity(value *UserExperienceAnalyticsAnomalySeverity)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
func (m *UserExperienceAnalyticsAnomaly) SetState(value *UserExperienceAnalyticsAnomalyState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsAnomalyable 
type UserExperienceAnalyticsAnomalyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnomalyFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAnomalyId()(*string)
    GetAnomalyLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAnomalyName()(*string)
    GetAnomalyType()(*UserExperienceAnalyticsAnomalyType)
    GetAssetName()(*string)
    GetAssetPublisher()(*string)
    GetAssetVersion()(*string)
    GetDetectionModelId()(*string)
    GetDeviceImpactedCount()(*int32)
    GetIssueId()(*string)
    GetSeverity()(*UserExperienceAnalyticsAnomalySeverity)
    GetState()(*UserExperienceAnalyticsAnomalyState)
    SetAnomalyFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAnomalyId(value *string)()
    SetAnomalyLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAnomalyName(value *string)()
    SetAnomalyType(value *UserExperienceAnalyticsAnomalyType)()
    SetAssetName(value *string)()
    SetAssetPublisher(value *string)()
    SetAssetVersion(value *string)()
    SetDetectionModelId(value *string)()
    SetDeviceImpactedCount(value *int32)()
    SetIssueId(value *string)()
    SetSeverity(value *UserExperienceAnalyticsAnomalySeverity)()
    SetState(value *UserExperienceAnalyticsAnomalyState)()
}
