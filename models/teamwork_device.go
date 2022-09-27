package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDevice provides operations to manage the collection of accessReviewDecision entities.
type TeamworkDevice struct {
    Entity
    // The activity properties that change based on the device usage.
    activity TeamworkDeviceActivityable
    // The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
    activityState *TeamworkDeviceActivityState
    // The company asset tag assigned by the admin on the device.
    companyAssetTag *string
    // The configuration properties of the device.
    configuration TeamworkDeviceConfigurationable
    // Identity of the user who enrolled the device to the tenant.
    createdBy IdentitySetable
    // The UTC date and time when the device was enrolled to the tenant.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The signed-in user on the device.
    currentUser TeamworkUserIdentityable
    // The deviceType property
    deviceType *TeamworkDeviceType
    // The hardwareDetail property
    hardwareDetail TeamworkHardwareDetailable
    // The health properties of the device.
    health TeamworkDeviceHealthable
    // The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
    healthStatus *TeamworkDeviceHealthStatus
    // Identity of the user who last modified the device details.
    lastModifiedBy IdentitySetable
    // The UTC date and time when the device detail was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The notes added by the admin to the device.
    notes *string
    // The async operations on the device.
    operations []TeamworkDeviceOperationable
}
// NewTeamworkDevice instantiates a new teamworkDevice and sets the default values.
func NewTeamworkDevice()(*TeamworkDevice) {
    m := &TeamworkDevice{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.teamworkDevice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamworkDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDevice(), nil
}
// GetActivity gets the activity property value. The activity properties that change based on the device usage.
func (m *TeamworkDevice) GetActivity()(TeamworkDeviceActivityable) {
    return m.activity
}
// GetActivityState gets the activityState property value. The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
func (m *TeamworkDevice) GetActivityState()(*TeamworkDeviceActivityState) {
    return m.activityState
}
// GetCompanyAssetTag gets the companyAssetTag property value. The company asset tag assigned by the admin on the device.
func (m *TeamworkDevice) GetCompanyAssetTag()(*string) {
    return m.companyAssetTag
}
// GetConfiguration gets the configuration property value. The configuration properties of the device.
func (m *TeamworkDevice) GetConfiguration()(TeamworkDeviceConfigurationable) {
    return m.configuration
}
// GetCreatedBy gets the createdBy property value. Identity of the user who enrolled the device to the tenant.
func (m *TeamworkDevice) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device was enrolled to the tenant.
func (m *TeamworkDevice) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCurrentUser gets the currentUser property value. The signed-in user on the device.
func (m *TeamworkDevice) GetCurrentUser()(TeamworkUserIdentityable) {
    return m.currentUser
}
// GetDeviceType gets the deviceType property value. The deviceType property
func (m *TeamworkDevice) GetDeviceType()(*TeamworkDeviceType) {
    return m.deviceType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkDeviceActivityFromDiscriminatorValue , m.SetActivity)
    res["activityState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTeamworkDeviceActivityState , m.SetActivityState)
    res["companyAssetTag"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompanyAssetTag)
    res["configuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkDeviceConfigurationFromDiscriminatorValue , m.SetConfiguration)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["currentUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkUserIdentityFromDiscriminatorValue , m.SetCurrentUser)
    res["deviceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTeamworkDeviceType , m.SetDeviceType)
    res["hardwareDetail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkHardwareDetailFromDiscriminatorValue , m.SetHardwareDetail)
    res["health"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkDeviceHealthFromDiscriminatorValue , m.SetHealth)
    res["healthStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTeamworkDeviceHealthStatus , m.SetHealthStatus)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["operations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamworkDeviceOperationFromDiscriminatorValue , m.SetOperations)
    return res
}
// GetHardwareDetail gets the hardwareDetail property value. The hardwareDetail property
func (m *TeamworkDevice) GetHardwareDetail()(TeamworkHardwareDetailable) {
    return m.hardwareDetail
}
// GetHealth gets the health property value. The health properties of the device.
func (m *TeamworkDevice) GetHealth()(TeamworkDeviceHealthable) {
    return m.health
}
// GetHealthStatus gets the healthStatus property value. The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
func (m *TeamworkDevice) GetHealthStatus()(*TeamworkDeviceHealthStatus) {
    return m.healthStatus
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device details.
func (m *TeamworkDevice) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device detail was last modified.
func (m *TeamworkDevice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNotes gets the notes property value. The notes added by the admin to the device.
func (m *TeamworkDevice) GetNotes()(*string) {
    return m.notes
}
// GetOperations gets the operations property value. The async operations on the device.
func (m *TeamworkDevice) GetOperations()([]TeamworkDeviceOperationable) {
    return m.operations
}
// Serialize serializes information the current object
func (m *TeamworkDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    if m.GetActivityState() != nil {
        cast := (*m.GetActivityState()).String()
        err = writer.WriteStringValue("activityState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyAssetTag", m.GetCompanyAssetTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currentUser", m.GetCurrentUser())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hardwareDetail", m.GetHardwareDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("health", m.GetHealth())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOperations())
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity properties that change based on the device usage.
func (m *TeamworkDevice) SetActivity(value TeamworkDeviceActivityable)() {
    m.activity = value
}
// SetActivityState sets the activityState property value. The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
func (m *TeamworkDevice) SetActivityState(value *TeamworkDeviceActivityState)() {
    m.activityState = value
}
// SetCompanyAssetTag sets the companyAssetTag property value. The company asset tag assigned by the admin on the device.
func (m *TeamworkDevice) SetCompanyAssetTag(value *string)() {
    m.companyAssetTag = value
}
// SetConfiguration sets the configuration property value. The configuration properties of the device.
func (m *TeamworkDevice) SetConfiguration(value TeamworkDeviceConfigurationable)() {
    m.configuration = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who enrolled the device to the tenant.
func (m *TeamworkDevice) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device was enrolled to the tenant.
func (m *TeamworkDevice) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCurrentUser sets the currentUser property value. The signed-in user on the device.
func (m *TeamworkDevice) SetCurrentUser(value TeamworkUserIdentityable)() {
    m.currentUser = value
}
// SetDeviceType sets the deviceType property value. The deviceType property
func (m *TeamworkDevice) SetDeviceType(value *TeamworkDeviceType)() {
    m.deviceType = value
}
// SetHardwareDetail sets the hardwareDetail property value. The hardwareDetail property
func (m *TeamworkDevice) SetHardwareDetail(value TeamworkHardwareDetailable)() {
    m.hardwareDetail = value
}
// SetHealth sets the health property value. The health properties of the device.
func (m *TeamworkDevice) SetHealth(value TeamworkDeviceHealthable)() {
    m.health = value
}
// SetHealthStatus sets the healthStatus property value. The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
func (m *TeamworkDevice) SetHealthStatus(value *TeamworkDeviceHealthStatus)() {
    m.healthStatus = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device details.
func (m *TeamworkDevice) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device detail was last modified.
func (m *TeamworkDevice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNotes sets the notes property value. The notes added by the admin to the device.
func (m *TeamworkDevice) SetNotes(value *string)() {
    m.notes = value
}
// SetOperations sets the operations property value. The async operations on the device.
func (m *TeamworkDevice) SetOperations(value []TeamworkDeviceOperationable)() {
    m.operations = value
}
