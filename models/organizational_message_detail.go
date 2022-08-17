package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageDetail 
type OrganizationalMessageDetail struct {
    Entity
    // The content that will be displayed to clients for the message. This includes the text portion of the message and the displayed logo
    content OrganizationalMessageContentable
    // The date and time of when the message was created
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time of when the message will stop being displayed to clients
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The frequency at which a client will see the message. Possible values are: weeklyOnce, monthlyOnce, monthlyTwice.
    frequency *OrganizationalMessageFrequency
    // The date and time of when the message was last modified
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates the scenario for the message. Possible values are: onboarding, lifecycle.
    scenario *OrganizationalMessageScenario
    // The date and time of when the message will start being displayed to clients
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates the deployment status of the message. Possible values are: scheduled, active, completed, cancelled.
    status *OrganizationalMessageStatus
    // Indicates the area where content will be displayed to customers. Possible values are: actionCenter, getStarted, softLanding.
    surface *OrganizationalMessageSurface
    // The groups of devices that will receive the message. This also contains a list of excluded groups that will not receive the message regardless of the device being part of an included group
    targeting OrganizationalMessageTargetingable
    // Indicates the theme for the experience. Possible values are: update, training, welcomeToWindows.
    theme *OrganizationalMessageTheme
    // The statistics containing how the message was interacted with by clients. This includes the number of impressions, clicks, and dismisses from targeted clients.
    userEngagementStatistics OrganizationalMessageInsightsable
    // Indicates the corresponding variant for the experience
    variant *string
}
// NewOrganizationalMessageDetail instantiates a new OrganizationalMessageDetail and sets the default values.
func NewOrganizationalMessageDetail()(*OrganizationalMessageDetail) {
    m := &OrganizationalMessageDetail{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.organizationalMessageDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageDetail(), nil
}
// GetContent gets the content property value. The content that will be displayed to clients for the message. This includes the text portion of the message and the displayed logo
func (m *OrganizationalMessageDetail) GetContent()(OrganizationalMessageContentable) {
    return m.content
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time of when the message was created
func (m *OrganizationalMessageDetail) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetEndDateTime gets the endDateTime property value. The date and time of when the message will stop being displayed to clients
func (m *OrganizationalMessageDetail) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalMessageContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(OrganizationalMessageContentable))
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["frequency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageFrequency)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrequency(val.(*OrganizationalMessageFrequency))
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
    res["scenario"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageScenario)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScenario(val.(*OrganizationalMessageScenario))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*OrganizationalMessageStatus))
        }
        return nil
    }
    res["surface"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageSurface)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurface(val.(*OrganizationalMessageSurface))
        }
        return nil
    }
    res["targeting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalMessageTargetingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargeting(val.(OrganizationalMessageTargetingable))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageTheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val.(*OrganizationalMessageTheme))
        }
        return nil
    }
    res["userEngagementStatistics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalMessageInsightsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEngagementStatistics(val.(OrganizationalMessageInsightsable))
        }
        return nil
    }
    res["variant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariant(val)
        }
        return nil
    }
    return res
}
// GetFrequency gets the frequency property value. The frequency at which a client will see the message. Possible values are: weeklyOnce, monthlyOnce, monthlyTwice.
func (m *OrganizationalMessageDetail) GetFrequency()(*OrganizationalMessageFrequency) {
    return m.frequency
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time of when the message was last modified
func (m *OrganizationalMessageDetail) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetScenario gets the scenario property value. Indicates the scenario for the message. Possible values are: onboarding, lifecycle.
func (m *OrganizationalMessageDetail) GetScenario()(*OrganizationalMessageScenario) {
    return m.scenario
}
// GetStartDateTime gets the startDateTime property value. The date and time of when the message will start being displayed to clients
func (m *OrganizationalMessageDetail) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. Indicates the deployment status of the message. Possible values are: scheduled, active, completed, cancelled.
func (m *OrganizationalMessageDetail) GetStatus()(*OrganizationalMessageStatus) {
    return m.status
}
// GetSurface gets the surface property value. Indicates the area where content will be displayed to customers. Possible values are: actionCenter, getStarted, softLanding.
func (m *OrganizationalMessageDetail) GetSurface()(*OrganizationalMessageSurface) {
    return m.surface
}
// GetTargeting gets the targeting property value. The groups of devices that will receive the message. This also contains a list of excluded groups that will not receive the message regardless of the device being part of an included group
func (m *OrganizationalMessageDetail) GetTargeting()(OrganizationalMessageTargetingable) {
    return m.targeting
}
// GetTheme gets the theme property value. Indicates the theme for the experience. Possible values are: update, training, welcomeToWindows.
func (m *OrganizationalMessageDetail) GetTheme()(*OrganizationalMessageTheme) {
    return m.theme
}
// GetUserEngagementStatistics gets the userEngagementStatistics property value. The statistics containing how the message was interacted with by clients. This includes the number of impressions, clicks, and dismisses from targeted clients.
func (m *OrganizationalMessageDetail) GetUserEngagementStatistics()(OrganizationalMessageInsightsable) {
    return m.userEngagementStatistics
}
// GetVariant gets the variant property value. Indicates the corresponding variant for the experience
func (m *OrganizationalMessageDetail) GetVariant()(*string) {
    return m.variant
}
// Serialize serializes information the current object
func (m *OrganizationalMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
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
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetFrequency() != nil {
        cast := (*m.GetFrequency()).String()
        err = writer.WriteStringValue("frequency", &cast)
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
    if m.GetScenario() != nil {
        cast := (*m.GetScenario()).String()
        err = writer.WriteStringValue("scenario", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSurface() != nil {
        cast := (*m.GetSurface()).String()
        err = writer.WriteStringValue("surface", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targeting", m.GetTargeting())
        if err != nil {
            return err
        }
    }
    if m.GetTheme() != nil {
        cast := (*m.GetTheme()).String()
        err = writer.WriteStringValue("theme", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userEngagementStatistics", m.GetUserEngagementStatistics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("variant", m.GetVariant())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content that will be displayed to clients for the message. This includes the text portion of the message and the displayed logo
func (m *OrganizationalMessageDetail) SetContent(value OrganizationalMessageContentable)() {
    m.content = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time of when the message was created
func (m *OrganizationalMessageDetail) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetEndDateTime sets the endDateTime property value. The date and time of when the message will stop being displayed to clients
func (m *OrganizationalMessageDetail) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetFrequency sets the frequency property value. The frequency at which a client will see the message. Possible values are: weeklyOnce, monthlyOnce, monthlyTwice.
func (m *OrganizationalMessageDetail) SetFrequency(value *OrganizationalMessageFrequency)() {
    m.frequency = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time of when the message was last modified
func (m *OrganizationalMessageDetail) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetScenario sets the scenario property value. Indicates the scenario for the message. Possible values are: onboarding, lifecycle.
func (m *OrganizationalMessageDetail) SetScenario(value *OrganizationalMessageScenario)() {
    m.scenario = value
}
// SetStartDateTime sets the startDateTime property value. The date and time of when the message will start being displayed to clients
func (m *OrganizationalMessageDetail) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. Indicates the deployment status of the message. Possible values are: scheduled, active, completed, cancelled.
func (m *OrganizationalMessageDetail) SetStatus(value *OrganizationalMessageStatus)() {
    m.status = value
}
// SetSurface sets the surface property value. Indicates the area where content will be displayed to customers. Possible values are: actionCenter, getStarted, softLanding.
func (m *OrganizationalMessageDetail) SetSurface(value *OrganizationalMessageSurface)() {
    m.surface = value
}
// SetTargeting sets the targeting property value. The groups of devices that will receive the message. This also contains a list of excluded groups that will not receive the message regardless of the device being part of an included group
func (m *OrganizationalMessageDetail) SetTargeting(value OrganizationalMessageTargetingable)() {
    m.targeting = value
}
// SetTheme sets the theme property value. Indicates the theme for the experience. Possible values are: update, training, welcomeToWindows.
func (m *OrganizationalMessageDetail) SetTheme(value *OrganizationalMessageTheme)() {
    m.theme = value
}
// SetUserEngagementStatistics sets the userEngagementStatistics property value. The statistics containing how the message was interacted with by clients. This includes the number of impressions, clicks, and dismisses from targeted clients.
func (m *OrganizationalMessageDetail) SetUserEngagementStatistics(value OrganizationalMessageInsightsable)() {
    m.userEngagementStatistics = value
}
// SetVariant sets the variant property value. Indicates the corresponding variant for the experience
func (m *OrganizationalMessageDetail) SetVariant(value *string)() {
    m.variant = value
}
