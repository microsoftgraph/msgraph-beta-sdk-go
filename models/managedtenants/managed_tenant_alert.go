package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenantAlert 
type ManagedTenantAlert struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagedTenantAlert instantiates a new managedTenantAlert and sets the default values.
func NewManagedTenantAlert()(*ManagedTenantAlert) {
    m := &ManagedTenantAlert{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagedTenantAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantAlert(), nil
}
// GetAlertData gets the alertData property value. The alertData property
func (m *ManagedTenantAlert) GetAlertData()(AlertDataable) {
    val, err := m.GetBackingStore().Get("alertData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AlertDataable)
    }
    return nil
}
// GetAlertDataReferenceStrings gets the alertDataReferenceStrings property value. The alertDataReferenceStrings property
func (m *ManagedTenantAlert) GetAlertDataReferenceStrings()([]AlertDataReferenceStringable) {
    val, err := m.GetBackingStore().Get("alertDataReferenceStrings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AlertDataReferenceStringable)
    }
    return nil
}
// GetAlertLogs gets the alertLogs property value. The alertLogs property
func (m *ManagedTenantAlert) GetAlertLogs()([]ManagedTenantAlertLogable) {
    val, err := m.GetBackingStore().Get("alertLogs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantAlertLogable)
    }
    return nil
}
// GetAlertRule gets the alertRule property value. The alertRule property
func (m *ManagedTenantAlert) GetAlertRule()(ManagedTenantAlertRuleable) {
    val, err := m.GetBackingStore().Get("alertRule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedTenantAlertRuleable)
    }
    return nil
}
// GetAlertRuleDisplayName gets the alertRuleDisplayName property value. The alertRuleDisplayName property
func (m *ManagedTenantAlert) GetAlertRuleDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("alertRuleDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApiNotifications gets the apiNotifications property value. The apiNotifications property
func (m *ManagedTenantAlert) GetApiNotifications()([]ManagedTenantApiNotificationable) {
    val, err := m.GetBackingStore().Get("apiNotifications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantApiNotificationable)
    }
    return nil
}
// GetAssignedToUserId gets the assignedToUserId property value. The assignedToUserId property
func (m *ManagedTenantAlert) GetAssignedToUserId()(*string) {
    val, err := m.GetBackingStore().Get("assignedToUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCorrelationCount gets the correlationCount property value. The correlationCount property
func (m *ManagedTenantAlert) GetCorrelationCount()(*int32) {
    val, err := m.GetBackingStore().Get("correlationCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCorrelationId gets the correlationId property value. The correlationId property
func (m *ManagedTenantAlert) GetCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("correlationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlert) GetCreatedByUserId()(*string) {
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
func (m *ManagedTenantAlert) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEmailNotifications gets the emailNotifications property value. The emailNotifications property
func (m *ManagedTenantAlert) GetEmailNotifications()([]ManagedTenantEmailNotificationable) {
    val, err := m.GetBackingStore().Get("emailNotifications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantEmailNotificationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlertDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertData(val.(AlertDataable))
        }
        return nil
    }
    res["alertDataReferenceStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertDataReferenceStringFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertDataReferenceStringable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AlertDataReferenceStringable)
                }
            }
            m.SetAlertDataReferenceStrings(res)
        }
        return nil
    }
    res["alertLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertLogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertLogable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedTenantAlertLogable)
                }
            }
            m.SetAlertLogs(res)
        }
        return nil
    }
    res["alertRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedTenantAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRule(val.(ManagedTenantAlertRuleable))
        }
        return nil
    }
    res["alertRuleDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleDisplayName(val)
        }
        return nil
    }
    res["apiNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantApiNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantApiNotificationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedTenantApiNotificationable)
                }
            }
            m.SetApiNotifications(res)
        }
        return nil
    }
    res["assignedToUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedToUserId(val)
        }
        return nil
    }
    res["correlationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationCount(val)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
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
    res["emailNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantEmailNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantEmailNotificationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedTenantEmailNotificationable)
                }
            }
            m.SetEmailNotifications(res)
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
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AlertStatus))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlert) GetLastActionByUserId()(*string) {
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
func (m *ManagedTenantAlert) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMessage gets the message property value. The message property
func (m *ManagedTenantAlert) GetMessage()(*string) {
    val, err := m.GetBackingStore().Get("message")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedTenantAlert) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSeverity gets the severity property value. The severity property
func (m *ManagedTenantAlert) GetSeverity()(*AlertSeverity) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertSeverity)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *ManagedTenantAlert) GetStatus()(*AlertStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertStatus)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantAlert) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTitle gets the title property value. The title property
func (m *ManagedTenantAlert) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedTenantAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alertData", m.GetAlertData())
        if err != nil {
            return err
        }
    }
    if m.GetAlertDataReferenceStrings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertDataReferenceStrings()))
        for i, v := range m.GetAlertDataReferenceStrings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alertDataReferenceStrings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlertLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertLogs()))
        for i, v := range m.GetAlertLogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alertLogs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("alertRule", m.GetAlertRule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertRuleDisplayName", m.GetAlertRuleDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetApiNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApiNotifications()))
        for i, v := range m.GetApiNotifications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("apiNotifications", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignedToUserId", m.GetAssignedToUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("correlationCount", m.GetCorrelationCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
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
    if m.GetEmailNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailNotifications()))
        for i, v := range m.GetEmailNotifications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("emailNotifications", cast)
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
        err = writer.WriteStringValue("message", m.GetMessage())
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
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
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
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertData sets the alertData property value. The alertData property
func (m *ManagedTenantAlert) SetAlertData(value AlertDataable)() {
    err := m.GetBackingStore().Set("alertData", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertDataReferenceStrings sets the alertDataReferenceStrings property value. The alertDataReferenceStrings property
func (m *ManagedTenantAlert) SetAlertDataReferenceStrings(value []AlertDataReferenceStringable)() {
    err := m.GetBackingStore().Set("alertDataReferenceStrings", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertLogs sets the alertLogs property value. The alertLogs property
func (m *ManagedTenantAlert) SetAlertLogs(value []ManagedTenantAlertLogable)() {
    err := m.GetBackingStore().Set("alertLogs", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRule sets the alertRule property value. The alertRule property
func (m *ManagedTenantAlert) SetAlertRule(value ManagedTenantAlertRuleable)() {
    err := m.GetBackingStore().Set("alertRule", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRuleDisplayName sets the alertRuleDisplayName property value. The alertRuleDisplayName property
func (m *ManagedTenantAlert) SetAlertRuleDisplayName(value *string)() {
    err := m.GetBackingStore().Set("alertRuleDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetApiNotifications sets the apiNotifications property value. The apiNotifications property
func (m *ManagedTenantAlert) SetApiNotifications(value []ManagedTenantApiNotificationable)() {
    err := m.GetBackingStore().Set("apiNotifications", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedToUserId sets the assignedToUserId property value. The assignedToUserId property
func (m *ManagedTenantAlert) SetAssignedToUserId(value *string)() {
    err := m.GetBackingStore().Set("assignedToUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationCount sets the correlationCount property value. The correlationCount property
func (m *ManagedTenantAlert) SetCorrelationCount(value *int32)() {
    err := m.GetBackingStore().Set("correlationCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationId sets the correlationId property value. The correlationId property
func (m *ManagedTenantAlert) SetCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("correlationId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlert) SetCreatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("createdByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlert) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailNotifications sets the emailNotifications property value. The emailNotifications property
func (m *ManagedTenantAlert) SetEmailNotifications(value []ManagedTenantEmailNotificationable)() {
    err := m.GetBackingStore().Set("emailNotifications", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlert) SetLastActionByUserId(value *string)() {
    err := m.GetBackingStore().Set("lastActionByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlert) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMessage sets the message property value. The message property
func (m *ManagedTenantAlert) SetMessage(value *string)() {
    err := m.GetBackingStore().Set("message", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedTenantAlert) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The severity property
func (m *ManagedTenantAlert) SetSeverity(value *AlertSeverity)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *ManagedTenantAlert) SetStatus(value *AlertStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantAlert) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. The title property
func (m *ManagedTenantAlert) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// ManagedTenantAlertable 
type ManagedTenantAlertable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertData()(AlertDataable)
    GetAlertDataReferenceStrings()([]AlertDataReferenceStringable)
    GetAlertLogs()([]ManagedTenantAlertLogable)
    GetAlertRule()(ManagedTenantAlertRuleable)
    GetAlertRuleDisplayName()(*string)
    GetApiNotifications()([]ManagedTenantApiNotificationable)
    GetAssignedToUserId()(*string)
    GetCorrelationCount()(*int32)
    GetCorrelationId()(*string)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmailNotifications()([]ManagedTenantEmailNotificationable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMessage()(*string)
    GetOdataType()(*string)
    GetSeverity()(*AlertSeverity)
    GetStatus()(*AlertStatus)
    GetTenantId()(*string)
    GetTitle()(*string)
    SetAlertData(value AlertDataable)()
    SetAlertDataReferenceStrings(value []AlertDataReferenceStringable)()
    SetAlertLogs(value []ManagedTenantAlertLogable)()
    SetAlertRule(value ManagedTenantAlertRuleable)()
    SetAlertRuleDisplayName(value *string)()
    SetApiNotifications(value []ManagedTenantApiNotificationable)()
    SetAssignedToUserId(value *string)()
    SetCorrelationCount(value *int32)()
    SetCorrelationId(value *string)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmailNotifications(value []ManagedTenantEmailNotificationable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMessage(value *string)()
    SetOdataType(value *string)()
    SetSeverity(value *AlertSeverity)()
    SetStatus(value *AlertStatus)()
    SetTenantId(value *string)()
    SetTitle(value *string)()
}
