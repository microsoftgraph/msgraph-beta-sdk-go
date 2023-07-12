package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenantAlertRule 
type ManagedTenantAlertRule struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The OdataType property
    OdataType *string
}
// NewManagedTenantAlertRule instantiates a new managedTenantAlertRule and sets the default values.
func NewManagedTenantAlertRule()(*ManagedTenantAlertRule) {
    m := &ManagedTenantAlertRule{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagedTenantAlertRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantAlertRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantAlertRule(), nil
}
// GetAlertDisplayName gets the alertDisplayName property value. The alertDisplayName property
func (m *ManagedTenantAlertRule) GetAlertDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("alertDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlerts gets the alerts property value. The alerts property
func (m *ManagedTenantAlertRule) GetAlerts()([]ManagedTenantAlertable) {
    val, err := m.GetBackingStore().Get("alerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantAlertable)
    }
    return nil
}
// GetAlertTTL gets the alertTTL property value. The alertTTL property
func (m *ManagedTenantAlertRule) GetAlertTTL()(*int32) {
    val, err := m.GetBackingStore().Get("alertTTL")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlertRule) GetCreatedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("createdByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlertRule) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *ManagedTenantAlertRule) GetDescription()(*string) {
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
func (m *ManagedTenantAlertRule) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantAlertRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertDisplayName(val)
        }
        return nil
    }
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedTenantAlertable)
                }
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["alertTTL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertTTL(val)
        }
        return nil
    }
    res["createdByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
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
    res["lastActionByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["lastRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRunDateTime(val)
        }
        return nil
    }
    res["notificationFinalDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationDestination)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationFinalDestinations(val.(*NotificationDestination))
        }
        return nil
    }
    res["ruleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleDefinition(val.(ManagedTenantAlertRuleDefinitionable))
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*AlertSeverity))
        }
        return nil
    }
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NotificationTargetable)
                }
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["tenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TenantInfoable)
                }
            }
            m.SetTenantIds(res)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlertRule) GetLastActionByUserId()(*string) {
    val, err := m.GetBackingStore().Get("lastActionByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlertRule) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastRunDateTime gets the lastRunDateTime property value. The lastRunDateTime property
func (m *ManagedTenantAlertRule) GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastRunDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNotificationFinalDestinations gets the notificationFinalDestinations property value. The notificationFinalDestinations property
func (m *ManagedTenantAlertRule) GetNotificationFinalDestinations()(*NotificationDestination) {
    val, err := m.GetBackingStore().Get("notificationFinalDestinations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NotificationDestination)
    }
    return nil
}
// GetRuleDefinition gets the ruleDefinition property value. The ruleDefinition property
func (m *ManagedTenantAlertRule) GetRuleDefinition()(ManagedTenantAlertRuleDefinitionable) {
    val, err := m.GetBackingStore().Get("ruleDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedTenantAlertRuleDefinitionable)
    }
    return nil
}
// GetSeverity gets the severity property value. The severity property
func (m *ManagedTenantAlertRule) GetSeverity()(*AlertSeverity) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertSeverity)
    }
    return nil
}
// GetTargets gets the targets property value. The targets property
func (m *ManagedTenantAlertRule) GetTargets()([]NotificationTargetable) {
    val, err := m.GetBackingStore().Get("targets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NotificationTargetable)
    }
    return nil
}
// GetTenantIds gets the tenantIds property value. The tenantIds property
func (m *ManagedTenantAlertRule) GetTenantIds()([]TenantInfoable) {
    val, err := m.GetBackingStore().Get("tenantIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TenantInfoable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantIds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantIds()))
        for i, v := range m.GetTenantIds() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tenantIds", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertDisplayName sets the alertDisplayName property value. The alertDisplayName property
func (m *ManagedTenantAlertRule) SetAlertDisplayName(value *string)() {
    err := m.GetBackingStore().Set("alertDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAlerts sets the alerts property value. The alerts property
func (m *ManagedTenantAlertRule) SetAlerts(value []ManagedTenantAlertable)() {
    err := m.GetBackingStore().Set("alerts", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertTTL sets the alertTTL property value. The alertTTL property
func (m *ManagedTenantAlertRule) SetAlertTTL(value *int32)() {
    err := m.GetBackingStore().Set("alertTTL", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlertRule) SetCreatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("createdByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlertRule) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *ManagedTenantAlertRule) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ManagedTenantAlertRule) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlertRule) SetLastActionByUserId(value *string)() {
    err := m.GetBackingStore().Set("lastActionByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlertRule) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastRunDateTime sets the lastRunDateTime property value. The lastRunDateTime property
func (m *ManagedTenantAlertRule) SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastRunDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationFinalDestinations sets the notificationFinalDestinations property value. The notificationFinalDestinations property
func (m *ManagedTenantAlertRule) SetNotificationFinalDestinations(value *NotificationDestination)() {
    err := m.GetBackingStore().Set("notificationFinalDestinations", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleDefinition sets the ruleDefinition property value. The ruleDefinition property
func (m *ManagedTenantAlertRule) SetRuleDefinition(value ManagedTenantAlertRuleDefinitionable)() {
    err := m.GetBackingStore().Set("ruleDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The severity property
func (m *ManagedTenantAlertRule) SetSeverity(value *AlertSeverity)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// SetTargets sets the targets property value. The targets property
func (m *ManagedTenantAlertRule) SetTargets(value []NotificationTargetable)() {
    err := m.GetBackingStore().Set("targets", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantIds sets the tenantIds property value. The tenantIds property
func (m *ManagedTenantAlertRule) SetTenantIds(value []TenantInfoable)() {
    err := m.GetBackingStore().Set("tenantIds", value)
    if err != nil {
        panic(err)
    }
}
// ManagedTenantAlertRuleable 
type ManagedTenantAlertRuleable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertDisplayName()(*string)
    GetAlerts()([]ManagedTenantAlertable)
    GetAlertTTL()(*int32)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotificationFinalDestinations()(*NotificationDestination)
    GetRuleDefinition()(ManagedTenantAlertRuleDefinitionable)
    GetSeverity()(*AlertSeverity)
    GetTargets()([]NotificationTargetable)
    GetTenantIds()([]TenantInfoable)
    SetAlertDisplayName(value *string)()
    SetAlerts(value []ManagedTenantAlertable)()
    SetAlertTTL(value *int32)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotificationFinalDestinations(value *NotificationDestination)()
    SetRuleDefinition(value ManagedTenantAlertRuleDefinitionable)()
    SetSeverity(value *AlertSeverity)()
    SetTargets(value []NotificationTargetable)()
    SetTenantIds(value []TenantInfoable)()
}
