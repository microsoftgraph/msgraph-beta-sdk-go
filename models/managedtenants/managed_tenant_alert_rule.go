package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenantAlertRule provides operations to manage the collection of activityStatistics entities.
type ManagedTenantAlertRule struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The alertDisplayName property
    alertDisplayName *string
    // The alerts property
    alerts []ManagedTenantAlertable
    // The alertTTL property
    alertTTL *int32
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastRunDateTime property
    lastRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The notificationFinalDestinations property
    notificationFinalDestinations *NotificationDestination
    // The ruleDefinition property
    ruleDefinition ManagedTenantAlertRuleDefinitionable
    // The severity property
    severity *AlertSeverity
    // The targets property
    targets []NotificationTargetable
    // The tenantIds property
    tenantIds []TenantInfoable
}
// NewManagedTenantAlertRule instantiates a new managedTenantAlertRule and sets the default values.
func NewManagedTenantAlertRule()(*ManagedTenantAlertRule) {
    m := &ManagedTenantAlertRule{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managedTenantAlertRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedTenantAlertRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantAlertRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantAlertRule(), nil
}
// GetAlertDisplayName gets the alertDisplayName property value. The alertDisplayName property
func (m *ManagedTenantAlertRule) GetAlertDisplayName()(*string) {
    return m.alertDisplayName
}
// GetAlerts gets the alerts property value. The alerts property
func (m *ManagedTenantAlertRule) GetAlerts()([]ManagedTenantAlertable) {
    return m.alerts
}
// GetAlertTTL gets the alertTTL property value. The alertTTL property
func (m *ManagedTenantAlertRule) GetAlertTTL()(*int32) {
    return m.alertTTL
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlertRule) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlertRule) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description property
func (m *ManagedTenantAlertRule) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ManagedTenantAlertRule) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantAlertRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlertDisplayName)
    res["alerts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedTenantAlertFromDiscriminatorValue , m.SetAlerts)
    res["alertTTL"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAlertTTL)
    res["createdByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedByUserId)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastActionByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastActionByUserId)
    res["lastActionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastActionDateTime)
    res["lastRunDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastRunDateTime)
    res["notificationFinalDestinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseNotificationDestination , m.SetNotificationFinalDestinations)
    res["ruleDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue , m.SetRuleDefinition)
    res["severity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAlertSeverity , m.SetSeverity)
    res["targets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNotificationTargetFromDiscriminatorValue , m.SetTargets)
    res["tenantIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantInfoFromDiscriminatorValue , m.SetTenantIds)
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlertRule) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlertRule) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetLastRunDateTime gets the lastRunDateTime property value. The lastRunDateTime property
func (m *ManagedTenantAlertRule) GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRunDateTime
}
// GetNotificationFinalDestinations gets the notificationFinalDestinations property value. The notificationFinalDestinations property
func (m *ManagedTenantAlertRule) GetNotificationFinalDestinations()(*NotificationDestination) {
    return m.notificationFinalDestinations
}
// GetRuleDefinition gets the ruleDefinition property value. The ruleDefinition property
func (m *ManagedTenantAlertRule) GetRuleDefinition()(ManagedTenantAlertRuleDefinitionable) {
    return m.ruleDefinition
}
// GetSeverity gets the severity property value. The severity property
func (m *ManagedTenantAlertRule) GetSeverity()(*AlertSeverity) {
    return m.severity
}
// GetTargets gets the targets property value. The targets property
func (m *ManagedTenantAlertRule) GetTargets()([]NotificationTargetable) {
    return m.targets
}
// GetTenantIds gets the tenantIds property value. The tenantIds property
func (m *ManagedTenantAlertRule) GetTenantIds()([]TenantInfoable) {
    return m.tenantIds
}
// Serialize serializes information the current object
func (m *ManagedTenantAlertRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alertDisplayName", m.GetAlertDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetAlerts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAlerts())
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("alertTTL", m.GetAlertTTL())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
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
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRunDateTime", m.GetLastRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationFinalDestinations() != nil {
        cast := (*m.GetNotificationFinalDestinations()).String()
        err = writer.WriteStringValue("notificationFinalDestinations", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ruleDefinition", m.GetRuleDefinition())
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
    if m.GetTargets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTargets())
        err = writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantIds() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenantIds())
        err = writer.WriteCollectionOfObjectValues("tenantIds", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertDisplayName sets the alertDisplayName property value. The alertDisplayName property
func (m *ManagedTenantAlertRule) SetAlertDisplayName(value *string)() {
    m.alertDisplayName = value
}
// SetAlerts sets the alerts property value. The alerts property
func (m *ManagedTenantAlertRule) SetAlerts(value []ManagedTenantAlertable)() {
    m.alerts = value
}
// SetAlertTTL sets the alertTTL property value. The alertTTL property
func (m *ManagedTenantAlertRule) SetAlertTTL(value *int32)() {
    m.alertTTL = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlertRule) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlertRule) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *ManagedTenantAlertRule) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ManagedTenantAlertRule) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlertRule) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlertRule) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetLastRunDateTime sets the lastRunDateTime property value. The lastRunDateTime property
func (m *ManagedTenantAlertRule) SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRunDateTime = value
}
// SetNotificationFinalDestinations sets the notificationFinalDestinations property value. The notificationFinalDestinations property
func (m *ManagedTenantAlertRule) SetNotificationFinalDestinations(value *NotificationDestination)() {
    m.notificationFinalDestinations = value
}
// SetRuleDefinition sets the ruleDefinition property value. The ruleDefinition property
func (m *ManagedTenantAlertRule) SetRuleDefinition(value ManagedTenantAlertRuleDefinitionable)() {
    m.ruleDefinition = value
}
// SetSeverity sets the severity property value. The severity property
func (m *ManagedTenantAlertRule) SetSeverity(value *AlertSeverity)() {
    m.severity = value
}
// SetTargets sets the targets property value. The targets property
func (m *ManagedTenantAlertRule) SetTargets(value []NotificationTargetable)() {
    m.targets = value
}
// SetTenantIds sets the tenantIds property value. The tenantIds property
func (m *ManagedTenantAlertRule) SetTenantIds(value []TenantInfoable)() {
    m.tenantIds = value
}
