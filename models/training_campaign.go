package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TrainingCampaign struct {
    Entity
}
// NewTrainingCampaign instantiates a new TrainingCampaign and sets the default values.
func NewTrainingCampaign()(*TrainingCampaign) {
    m := &TrainingCampaign{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrainingCampaignFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrainingCampaignFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrainingCampaign(), nil
}
// GetCampaignSchedule gets the campaignSchedule property value. The campaignSchedule property
// returns a CampaignScheduleable when successful
func (m *TrainingCampaign) GetCampaignSchedule()(CampaignScheduleable) {
    val, err := m.GetBackingStore().Get("campaignSchedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CampaignScheduleable)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
// returns a EmailIdentityable when successful
func (m *TrainingCampaign) GetCreatedBy()(EmailIdentityable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EmailIdentityable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
// returns a *Time when successful
func (m *TrainingCampaign) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TrainingCampaign) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *TrainingCampaign) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndUserNotificationSetting gets the endUserNotificationSetting property value. The endUserNotificationSetting property
// returns a EndUserNotificationSettingable when successful
func (m *TrainingCampaign) GetEndUserNotificationSetting()(EndUserNotificationSettingable) {
    val, err := m.GetBackingStore().Get("endUserNotificationSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EndUserNotificationSettingable)
    }
    return nil
}
// GetExcludedAccountTarget gets the excludedAccountTarget property value. The excludedAccountTarget property
// returns a AccountTargetContentable when successful
func (m *TrainingCampaign) GetExcludedAccountTarget()(AccountTargetContentable) {
    val, err := m.GetBackingStore().Get("excludedAccountTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccountTargetContentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrainingCampaign) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["campaignSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCampaignScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCampaignSchedule(val.(CampaignScheduleable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endUserNotificationSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEndUserNotificationSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndUserNotificationSetting(val.(EndUserNotificationSettingable))
        }
        return nil
    }
    res["excludedAccountTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountTargetContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedAccountTarget(val.(AccountTargetContentable))
        }
        return nil
    }
    res["includedAccountTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountTargetContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedAccountTarget(val.(AccountTargetContentable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
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
    res["report"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingCampaignReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReport(val.(TrainingCampaignReportable))
        }
        return nil
    }
    res["trainingSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingSetting(val.(TrainingSettingable))
        }
        return nil
    }
    return res
}
// GetIncludedAccountTarget gets the includedAccountTarget property value. The includedAccountTarget property
// returns a AccountTargetContentable when successful
func (m *TrainingCampaign) GetIncludedAccountTarget()(AccountTargetContentable) {
    val, err := m.GetBackingStore().Get("includedAccountTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccountTargetContentable)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
// returns a EmailIdentityable when successful
func (m *TrainingCampaign) GetLastModifiedBy()(EmailIdentityable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EmailIdentityable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *TrainingCampaign) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReport gets the report property value. The report property
// returns a TrainingCampaignReportable when successful
func (m *TrainingCampaign) GetReport()(TrainingCampaignReportable) {
    val, err := m.GetBackingStore().Get("report")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TrainingCampaignReportable)
    }
    return nil
}
// GetTrainingSetting gets the trainingSetting property value. The trainingSetting property
// returns a TrainingSettingable when successful
func (m *TrainingCampaign) GetTrainingSetting()(TrainingSettingable) {
    val, err := m.GetBackingStore().Get("trainingSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TrainingSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrainingCampaign) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("campaignSchedule", m.GetCampaignSchedule())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("endUserNotificationSetting", m.GetEndUserNotificationSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("excludedAccountTarget", m.GetExcludedAccountTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("includedAccountTarget", m.GetIncludedAccountTarget())
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
        err = writer.WriteObjectValue("report", m.GetReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trainingSetting", m.GetTrainingSetting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCampaignSchedule sets the campaignSchedule property value. The campaignSchedule property
func (m *TrainingCampaign) SetCampaignSchedule(value CampaignScheduleable)() {
    err := m.GetBackingStore().Set("campaignSchedule", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *TrainingCampaign) SetCreatedBy(value EmailIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *TrainingCampaign) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *TrainingCampaign) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *TrainingCampaign) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEndUserNotificationSetting sets the endUserNotificationSetting property value. The endUserNotificationSetting property
func (m *TrainingCampaign) SetEndUserNotificationSetting(value EndUserNotificationSettingable)() {
    err := m.GetBackingStore().Set("endUserNotificationSetting", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludedAccountTarget sets the excludedAccountTarget property value. The excludedAccountTarget property
func (m *TrainingCampaign) SetExcludedAccountTarget(value AccountTargetContentable)() {
    err := m.GetBackingStore().Set("excludedAccountTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludedAccountTarget sets the includedAccountTarget property value. The includedAccountTarget property
func (m *TrainingCampaign) SetIncludedAccountTarget(value AccountTargetContentable)() {
    err := m.GetBackingStore().Set("includedAccountTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *TrainingCampaign) SetLastModifiedBy(value EmailIdentityable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *TrainingCampaign) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReport sets the report property value. The report property
func (m *TrainingCampaign) SetReport(value TrainingCampaignReportable)() {
    err := m.GetBackingStore().Set("report", value)
    if err != nil {
        panic(err)
    }
}
// SetTrainingSetting sets the trainingSetting property value. The trainingSetting property
func (m *TrainingCampaign) SetTrainingSetting(value TrainingSettingable)() {
    err := m.GetBackingStore().Set("trainingSetting", value)
    if err != nil {
        panic(err)
    }
}
type TrainingCampaignable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCampaignSchedule()(CampaignScheduleable)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndUserNotificationSetting()(EndUserNotificationSettingable)
    GetExcludedAccountTarget()(AccountTargetContentable)
    GetIncludedAccountTarget()(AccountTargetContentable)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReport()(TrainingCampaignReportable)
    GetTrainingSetting()(TrainingSettingable)
    SetCampaignSchedule(value CampaignScheduleable)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndUserNotificationSetting(value EndUserNotificationSettingable)()
    SetExcludedAccountTarget(value AccountTargetContentable)()
    SetIncludedAccountTarget(value AccountTargetContentable)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReport(value TrainingCampaignReportable)()
    SetTrainingSetting(value TrainingSettingable)()
}
