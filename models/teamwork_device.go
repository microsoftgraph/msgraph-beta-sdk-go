package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDevice 
type TeamworkDevice struct {
    Entity
}
// NewTeamworkDevice instantiates a new teamworkDevice and sets the default values.
func NewTeamworkDevice()(*TeamworkDevice) {
    m := &TeamworkDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDevice(), nil
}
// GetActivity gets the activity property value. The activity properties that change based on the device usage.
func (m *TeamworkDevice) GetActivity()(TeamworkDeviceActivityable) {
    val, err := m.GetBackingStore().Get("activity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDeviceActivityable)
    }
    return nil
}
// GetActivityState gets the activityState property value. The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
func (m *TeamworkDevice) GetActivityState()(*TeamworkDeviceActivityState) {
    val, err := m.GetBackingStore().Get("activityState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamworkDeviceActivityState)
    }
    return nil
}
// GetCompanyAssetTag gets the companyAssetTag property value. The company asset tag assigned by the admin on the device.
func (m *TeamworkDevice) GetCompanyAssetTag()(*string) {
    val, err := m.GetBackingStore().Get("companyAssetTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfiguration gets the configuration property value. The configuration properties of the device.
func (m *TeamworkDevice) GetConfiguration()(TeamworkDeviceConfigurationable) {
    val, err := m.GetBackingStore().Get("configuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDeviceConfigurationable)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user who enrolled the device to the tenant.
func (m *TeamworkDevice) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device was enrolled to the tenant.
func (m *TeamworkDevice) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCurrentUser gets the currentUser property value. The signed-in user on the device.
func (m *TeamworkDevice) GetCurrentUser()(TeamworkUserIdentityable) {
    val, err := m.GetBackingStore().Get("currentUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkUserIdentityable)
    }
    return nil
}
// GetDeviceType gets the deviceType property value. The deviceType property
func (m *TeamworkDevice) GetDeviceType()(*TeamworkDeviceType) {
    val, err := m.GetBackingStore().Get("deviceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamworkDeviceType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDeviceActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(TeamworkDeviceActivityable))
        }
        return nil
    }
    res["activityState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceActivityState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityState(val.(*TeamworkDeviceActivityState))
        }
        return nil
    }
    res["companyAssetTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyAssetTag(val)
        }
        return nil
    }
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(TeamworkDeviceConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["currentUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUser(val.(TeamworkUserIdentityable))
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*TeamworkDeviceType))
        }
        return nil
    }
    res["hardwareDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHardwareDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareDetail(val.(TeamworkHardwareDetailable))
        }
        return nil
    }
    res["health"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDeviceHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealth(val.(TeamworkDeviceHealthable))
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*TeamworkDeviceHealthStatus))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
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
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkDeviceOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkDeviceOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamworkDeviceOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    return res
}
// GetHardwareDetail gets the hardwareDetail property value. The hardwareDetail property
func (m *TeamworkDevice) GetHardwareDetail()(TeamworkHardwareDetailable) {
    val, err := m.GetBackingStore().Get("hardwareDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkHardwareDetailable)
    }
    return nil
}
// GetHealth gets the health property value. The health properties of the device.
func (m *TeamworkDevice) GetHealth()(TeamworkDeviceHealthable) {
    val, err := m.GetBackingStore().Get("health")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDeviceHealthable)
    }
    return nil
}
// GetHealthStatus gets the healthStatus property value. The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
func (m *TeamworkDevice) GetHealthStatus()(*TeamworkDeviceHealthStatus) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamworkDeviceHealthStatus)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device details.
func (m *TeamworkDevice) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device detail was last modified.
func (m *TeamworkDevice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNotes gets the notes property value. The notes added by the admin to the device.
func (m *TeamworkDevice) GetNotes()(*string) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDevice) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperations gets the operations property value. The async operations on the device.
func (m *TeamworkDevice) GetOperations()([]TeamworkDeviceOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamworkDeviceOperationable)
    }
    return nil
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity properties that change based on the device usage.
func (m *TeamworkDevice) SetActivity(value TeamworkDeviceActivityable)() {
    err := m.GetBackingStore().Set("activity", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityState sets the activityState property value. The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
func (m *TeamworkDevice) SetActivityState(value *TeamworkDeviceActivityState)() {
    err := m.GetBackingStore().Set("activityState", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyAssetTag sets the companyAssetTag property value. The company asset tag assigned by the admin on the device.
func (m *TeamworkDevice) SetCompanyAssetTag(value *string)() {
    err := m.GetBackingStore().Set("companyAssetTag", value)
    if err != nil {
        panic(err)
    }
}
// SetConfiguration sets the configuration property value. The configuration properties of the device.
func (m *TeamworkDevice) SetConfiguration(value TeamworkDeviceConfigurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who enrolled the device to the tenant.
func (m *TeamworkDevice) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device was enrolled to the tenant.
func (m *TeamworkDevice) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentUser sets the currentUser property value. The signed-in user on the device.
func (m *TeamworkDevice) SetCurrentUser(value TeamworkUserIdentityable)() {
    err := m.GetBackingStore().Set("currentUser", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceType sets the deviceType property value. The deviceType property
func (m *TeamworkDevice) SetDeviceType(value *TeamworkDeviceType)() {
    err := m.GetBackingStore().Set("deviceType", value)
    if err != nil {
        panic(err)
    }
}
// SetHardwareDetail sets the hardwareDetail property value. The hardwareDetail property
func (m *TeamworkDevice) SetHardwareDetail(value TeamworkHardwareDetailable)() {
    err := m.GetBackingStore().Set("hardwareDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetHealth sets the health property value. The health properties of the device.
func (m *TeamworkDevice) SetHealth(value TeamworkDeviceHealthable)() {
    err := m.GetBackingStore().Set("health", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
func (m *TeamworkDevice) SetHealthStatus(value *TeamworkDeviceHealthStatus)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device details.
func (m *TeamworkDevice) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device detail was last modified.
func (m *TeamworkDevice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. The notes added by the admin to the device.
func (m *TeamworkDevice) SetNotes(value *string)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDevice) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. The async operations on the device.
func (m *TeamworkDevice) SetOperations(value []TeamworkDeviceOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkDeviceable 
type TeamworkDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(TeamworkDeviceActivityable)
    GetActivityState()(*TeamworkDeviceActivityState)
    GetCompanyAssetTag()(*string)
    GetConfiguration()(TeamworkDeviceConfigurationable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentUser()(TeamworkUserIdentityable)
    GetDeviceType()(*TeamworkDeviceType)
    GetHardwareDetail()(TeamworkHardwareDetailable)
    GetHealth()(TeamworkDeviceHealthable)
    GetHealthStatus()(*TeamworkDeviceHealthStatus)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()(*string)
    GetOdataType()(*string)
    GetOperations()([]TeamworkDeviceOperationable)
    SetActivity(value TeamworkDeviceActivityable)()
    SetActivityState(value *TeamworkDeviceActivityState)()
    SetCompanyAssetTag(value *string)()
    SetConfiguration(value TeamworkDeviceConfigurationable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentUser(value TeamworkUserIdentityable)()
    SetDeviceType(value *TeamworkDeviceType)()
    SetHardwareDetail(value TeamworkHardwareDetailable)()
    SetHealth(value TeamworkDeviceHealthable)()
    SetHealthStatus(value *TeamworkDeviceHealthStatus)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value *string)()
    SetOdataType(value *string)()
    SetOperations(value []TeamworkDeviceOperationable)()
}
