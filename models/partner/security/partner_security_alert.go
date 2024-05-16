package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type PartnerSecurityAlert struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewPartnerSecurityAlert instantiates a new PartnerSecurityAlert and sets the default values.
func NewPartnerSecurityAlert()(*PartnerSecurityAlert) {
    m := &PartnerSecurityAlert{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreatePartnerSecurityAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartnerSecurityAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPartnerSecurityAlert(), nil
}
// GetActivityLogs gets the activityLogs property value. Represents the activity by a partner and includes details of state transitions, who performed them, and when they occurred.
// returns a []ActivityLogable when successful
func (m *PartnerSecurityAlert) GetActivityLogs()([]ActivityLogable) {
    val, err := m.GetBackingStore().Get("activityLogs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActivityLogable)
    }
    return nil
}
// GetAdditionalDetails gets the additionalDetails property value. A bag of name-value pairs that contain additional details about an alert.
// returns a AdditionalDataDictionaryable when successful
func (m *PartnerSecurityAlert) GetAdditionalDetails()(AdditionalDataDictionaryable) {
    val, err := m.GetBackingStore().Get("additionalDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdditionalDataDictionaryable)
    }
    return nil
}
// GetAffectedResources gets the affectedResources property value. Contains details of the resources affected by the security alert.
// returns a []AffectedResourceable when successful
func (m *PartnerSecurityAlert) GetAffectedResources()([]AffectedResourceable) {
    val, err := m.GetBackingStore().Get("affectedResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AffectedResourceable)
    }
    return nil
}
// GetAlertType gets the alertType property value. The type of vulnerability that impacts the customer due to this alert.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetAlertType()(*string) {
    val, err := m.GetBackingStore().Get("alertType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCatalogOfferId gets the catalogOfferId property value. The modern offer category ID of the subscription.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetCatalogOfferId()(*string) {
    val, err := m.GetBackingStore().Get("catalogOfferId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfidenceLevel gets the confidenceLevel property value. The confidenceLevel property
// returns a *SecurityAlertConfidence when successful
func (m *PartnerSecurityAlert) GetConfidenceLevel()(*SecurityAlertConfidence) {
    val, err := m.GetBackingStore().Get("confidenceLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityAlertConfidence)
    }
    return nil
}
// GetCustomerTenantId gets the customerTenantId property value. The impacted customer tenant associated with the alert.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetCustomerTenantId()(*string) {
    val, err := m.GetBackingStore().Get("customerTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description for each alert.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetectedDateTime gets the detectedDateTime property value. Time when the alert was detected or created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *PartnerSecurityAlert) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("detectedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the alert.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartnerSecurityAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityLogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityLogable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActivityLogable)
                }
            }
            m.SetActivityLogs(res)
        }
        return nil
    }
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdditionalDataDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val.(AdditionalDataDictionaryable))
        }
        return nil
    }
    res["affectedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAffectedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AffectedResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AffectedResourceable)
                }
            }
            m.SetAffectedResources(res)
        }
        return nil
    }
    res["alertType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertType(val)
        }
        return nil
    }
    res["catalogOfferId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogOfferId(val)
        }
        return nil
    }
    res["confidenceLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAlertConfidence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidenceLevel(val.(*SecurityAlertConfidence))
        }
        return nil
    }
    res["customerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerTenantId(val)
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
    res["detectedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectedDateTime(val)
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
    res["firstObservedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstObservedDateTime(val)
        }
        return nil
    }
    res["isTest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTest(val)
        }
        return nil
    }
    res["lastObservedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastObservedDateTime(val)
        }
        return nil
    }
    res["resolvedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedBy(val)
        }
        return nil
    }
    res["resolvedOnDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedOnDateTime(val)
        }
        return nil
    }
    res["resolvedReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAlertResolvedReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedReason(val.(*SecurityAlertResolvedReason))
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*SecurityAlertSeverity))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SecurityAlertStatus))
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["valueAddedResellerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueAddedResellerTenantId(val)
        }
        return nil
    }
    return res
}
// GetFirstObservedDateTime gets the firstObservedDateTime property value. Time of the first activity associated with the alert. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  subscription.
// returns a *Time when successful
func (m *PartnerSecurityAlert) GetFirstObservedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstObservedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetIsTest gets the isTest property value. Indicates whehter an alert is a test alert.
// returns a *bool when successful
func (m *PartnerSecurityAlert) GetIsTest()(*bool) {
    val, err := m.GetBackingStore().Get("isTest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastObservedDateTime gets the lastObservedDateTime property value. Time of the latest activity associated with the alert. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *PartnerSecurityAlert) GetLastObservedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastObservedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetResolvedBy gets the resolvedBy property value. The UPN of the partner user who resolved the alert.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetResolvedBy()(*string) {
    val, err := m.GetBackingStore().Get("resolvedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResolvedOnDateTime gets the resolvedOnDateTime property value. Time when the alert was resolved. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *PartnerSecurityAlert) GetResolvedOnDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("resolvedOnDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetResolvedReason gets the resolvedReason property value. The reason provided by the partner for addressing the alert. The possible values are: legitimate, ignore, fraud, unknownFutureValue.
// returns a *SecurityAlertResolvedReason when successful
func (m *PartnerSecurityAlert) GetResolvedReason()(*SecurityAlertResolvedReason) {
    val, err := m.GetBackingStore().Get("resolvedReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityAlertResolvedReason)
    }
    return nil
}
// GetSeverity gets the severity property value. The severity property
// returns a *SecurityAlertSeverity when successful
func (m *PartnerSecurityAlert) GetSeverity()(*SecurityAlertSeverity) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityAlertSeverity)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *SecurityAlertStatus when successful
func (m *PartnerSecurityAlert) GetStatus()(*SecurityAlertStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityAlertStatus)
    }
    return nil
}
// GetSubscriptionId gets the subscriptionId property value. The subscription associated with the alert for the customer.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValueAddedResellerTenantId gets the valueAddedResellerTenantId property value. The value-added reseller tenant associated with the partner tenant and customer tenant.
// returns a *string when successful
func (m *PartnerSecurityAlert) GetValueAddedResellerTenantId()(*string) {
    val, err := m.GetBackingStore().Get("valueAddedResellerTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PartnerSecurityAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivityLogs()))
        for i, v := range m.GetActivityLogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activityLogs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    if m.GetAffectedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAffectedResources()))
        for i, v := range m.GetAffectedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("affectedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertType", m.GetAlertType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogOfferId", m.GetCatalogOfferId())
        if err != nil {
            return err
        }
    }
    if m.GetConfidenceLevel() != nil {
        cast := (*m.GetConfidenceLevel()).String()
        err = writer.WriteStringValue("confidenceLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerTenantId", m.GetCustomerTenantId())
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
        err = writer.WriteTimeValue("detectedDateTime", m.GetDetectedDateTime())
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
        err = writer.WriteTimeValue("firstObservedDateTime", m.GetFirstObservedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTest", m.GetIsTest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastObservedDateTime", m.GetLastObservedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resolvedBy", m.GetResolvedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("resolvedOnDateTime", m.GetResolvedOnDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetResolvedReason() != nil {
        cast := (*m.GetResolvedReason()).String()
        err = writer.WriteStringValue("resolvedReason", &cast)
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
        err = writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueAddedResellerTenantId", m.GetValueAddedResellerTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityLogs sets the activityLogs property value. Represents the activity by a partner and includes details of state transitions, who performed them, and when they occurred.
func (m *PartnerSecurityAlert) SetActivityLogs(value []ActivityLogable)() {
    err := m.GetBackingStore().Set("activityLogs", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalDetails sets the additionalDetails property value. A bag of name-value pairs that contain additional details about an alert.
func (m *PartnerSecurityAlert) SetAdditionalDetails(value AdditionalDataDictionaryable)() {
    err := m.GetBackingStore().Set("additionalDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetAffectedResources sets the affectedResources property value. Contains details of the resources affected by the security alert.
func (m *PartnerSecurityAlert) SetAffectedResources(value []AffectedResourceable)() {
    err := m.GetBackingStore().Set("affectedResources", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertType sets the alertType property value. The type of vulnerability that impacts the customer due to this alert.
func (m *PartnerSecurityAlert) SetAlertType(value *string)() {
    err := m.GetBackingStore().Set("alertType", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogOfferId sets the catalogOfferId property value. The modern offer category ID of the subscription.
func (m *PartnerSecurityAlert) SetCatalogOfferId(value *string)() {
    err := m.GetBackingStore().Set("catalogOfferId", value)
    if err != nil {
        panic(err)
    }
}
// SetConfidenceLevel sets the confidenceLevel property value. The confidenceLevel property
func (m *PartnerSecurityAlert) SetConfidenceLevel(value *SecurityAlertConfidence)() {
    err := m.GetBackingStore().Set("confidenceLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerTenantId sets the customerTenantId property value. The impacted customer tenant associated with the alert.
func (m *PartnerSecurityAlert) SetCustomerTenantId(value *string)() {
    err := m.GetBackingStore().Set("customerTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description for each alert.
func (m *PartnerSecurityAlert) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectedDateTime sets the detectedDateTime property value. Time when the alert was detected or created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PartnerSecurityAlert) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("detectedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the alert.
func (m *PartnerSecurityAlert) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstObservedDateTime sets the firstObservedDateTime property value. Time of the first activity associated with the alert. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  subscription.
func (m *PartnerSecurityAlert) SetFirstObservedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstObservedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTest sets the isTest property value. Indicates whehter an alert is a test alert.
func (m *PartnerSecurityAlert) SetIsTest(value *bool)() {
    err := m.GetBackingStore().Set("isTest", value)
    if err != nil {
        panic(err)
    }
}
// SetLastObservedDateTime sets the lastObservedDateTime property value. Time of the latest activity associated with the alert. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PartnerSecurityAlert) SetLastObservedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastObservedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResolvedBy sets the resolvedBy property value. The UPN of the partner user who resolved the alert.
func (m *PartnerSecurityAlert) SetResolvedBy(value *string)() {
    err := m.GetBackingStore().Set("resolvedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetResolvedOnDateTime sets the resolvedOnDateTime property value. Time when the alert was resolved. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PartnerSecurityAlert) SetResolvedOnDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("resolvedOnDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResolvedReason sets the resolvedReason property value. The reason provided by the partner for addressing the alert. The possible values are: legitimate, ignore, fraud, unknownFutureValue.
func (m *PartnerSecurityAlert) SetResolvedReason(value *SecurityAlertResolvedReason)() {
    err := m.GetBackingStore().Set("resolvedReason", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The severity property
func (m *PartnerSecurityAlert) SetSeverity(value *SecurityAlertSeverity)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *PartnerSecurityAlert) SetStatus(value *SecurityAlertStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionId sets the subscriptionId property value. The subscription associated with the alert for the customer.
func (m *PartnerSecurityAlert) SetSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("subscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetValueAddedResellerTenantId sets the valueAddedResellerTenantId property value. The value-added reseller tenant associated with the partner tenant and customer tenant.
func (m *PartnerSecurityAlert) SetValueAddedResellerTenantId(value *string)() {
    err := m.GetBackingStore().Set("valueAddedResellerTenantId", value)
    if err != nil {
        panic(err)
    }
}
type PartnerSecurityAlertable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityLogs()([]ActivityLogable)
    GetAdditionalDetails()(AdditionalDataDictionaryable)
    GetAffectedResources()([]AffectedResourceable)
    GetAlertType()(*string)
    GetCatalogOfferId()(*string)
    GetConfidenceLevel()(*SecurityAlertConfidence)
    GetCustomerTenantId()(*string)
    GetDescription()(*string)
    GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetFirstObservedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsTest()(*bool)
    GetLastObservedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResolvedBy()(*string)
    GetResolvedOnDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResolvedReason()(*SecurityAlertResolvedReason)
    GetSeverity()(*SecurityAlertSeverity)
    GetStatus()(*SecurityAlertStatus)
    GetSubscriptionId()(*string)
    GetValueAddedResellerTenantId()(*string)
    SetActivityLogs(value []ActivityLogable)()
    SetAdditionalDetails(value AdditionalDataDictionaryable)()
    SetAffectedResources(value []AffectedResourceable)()
    SetAlertType(value *string)()
    SetCatalogOfferId(value *string)()
    SetConfidenceLevel(value *SecurityAlertConfidence)()
    SetCustomerTenantId(value *string)()
    SetDescription(value *string)()
    SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetFirstObservedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsTest(value *bool)()
    SetLastObservedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResolvedBy(value *string)()
    SetResolvedOnDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResolvedReason(value *SecurityAlertResolvedReason)()
    SetSeverity(value *SecurityAlertSeverity)()
    SetStatus(value *SecurityAlertStatus)()
    SetSubscriptionId(value *string)()
    SetValueAddedResellerTenantId(value *string)()
}
