package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Custodian 
type Custodian struct {
    DataSourceContainer
    // Date and time the custodian acknowledged a hold notification.
    acknowledgedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identifies whether a custodian's sources were placed on hold during creation.
    applyHoldToSources *bool
    // Email address of the custodian.
    email *string
    // Data source entity for SharePoint sites associated with the custodian.
    siteSources []SiteSourceable
    // Data source entity for groups associated with the custodian.
    unifiedGroupSources []UnifiedGroupSourceable
    // Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
    userSources []UserSourceable
}
// NewCustodian instantiates a new Custodian and sets the default values.
func NewCustodian()(*Custodian) {
    m := &Custodian{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.custodian";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustodianFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustodianFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustodian(), nil
}
// GetAcknowledgedDateTime gets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *Custodian) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.acknowledgedDateTime
}
// GetApplyHoldToSources gets the applyHoldToSources property value. Identifies whether a custodian's sources were placed on hold during creation.
func (m *Custodian) GetApplyHoldToSources()(*bool) {
    return m.applyHoldToSources
}
// GetEmail gets the email property value. Email address of the custodian.
func (m *Custodian) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Custodian) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["acknowledgedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetAcknowledgedDateTime)
    res["applyHoldToSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetApplyHoldToSources)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["siteSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSiteSourceFromDiscriminatorValue , m.SetSiteSources)
    res["unifiedGroupSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedGroupSourceFromDiscriminatorValue , m.SetUnifiedGroupSources)
    res["userSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSourceFromDiscriminatorValue , m.SetUserSources)
    return res
}
// GetSiteSources gets the siteSources property value. Data source entity for SharePoint sites associated with the custodian.
func (m *Custodian) GetSiteSources()([]SiteSourceable) {
    return m.siteSources
}
// GetUnifiedGroupSources gets the unifiedGroupSources property value. Data source entity for groups associated with the custodian.
func (m *Custodian) GetUnifiedGroupSources()([]UnifiedGroupSourceable) {
    return m.unifiedGroupSources
}
// GetUserSources gets the userSources property value. Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
func (m *Custodian) GetUserSources()([]UserSourceable) {
    return m.userSources
}
// Serialize serializes information the current object
func (m *Custodian) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applyHoldToSources", m.GetApplyHoldToSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    if m.GetSiteSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSiteSources())
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUnifiedGroupSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUnifiedGroupSources())
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserSources())
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcknowledgedDateTime sets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *Custodian) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acknowledgedDateTime = value
}
// SetApplyHoldToSources sets the applyHoldToSources property value. Identifies whether a custodian's sources were placed on hold during creation.
func (m *Custodian) SetApplyHoldToSources(value *bool)() {
    m.applyHoldToSources = value
}
// SetEmail sets the email property value. Email address of the custodian.
func (m *Custodian) SetEmail(value *string)() {
    m.email = value
}
// SetSiteSources sets the siteSources property value. Data source entity for SharePoint sites associated with the custodian.
func (m *Custodian) SetSiteSources(value []SiteSourceable)() {
    m.siteSources = value
}
// SetUnifiedGroupSources sets the unifiedGroupSources property value. Data source entity for groups associated with the custodian.
func (m *Custodian) SetUnifiedGroupSources(value []UnifiedGroupSourceable)() {
    m.unifiedGroupSources = value
}
// SetUserSources sets the userSources property value. Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
func (m *Custodian) SetUserSources(value []UserSourceable)() {
    m.userSources = value
}
