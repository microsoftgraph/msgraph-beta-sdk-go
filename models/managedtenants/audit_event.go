package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AuditEvent provides operations to manage the collection of activityStatistics entities.
type AuditEvent struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // A string which uniquely represents the operation that occurred. Required. Read-only.
    activity *string
    // The time when the activity ocurred. Required. Read-only.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identifier of the activity request that made the audit event. Required. Read-only.
    activityId *string
    // A category which represents a logical grouping of activities. Required. Read-only.
    category *string
    // The HTTP verb that was used when making the API request. Required. Read-only.
    httpVerb *string
    // The identifier of the app that was used to make the request. Required. Read-only.
    initiatedByAppId *string
    // The UPN of the user who initiated the activity. Required. Read-only.
    initiatedByUpn *string
    // The identifier of the user who initiated the activity. Required. Read-only.
    initiatedByUserId *string
    // The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
    ipAddress *string
    // The raw HTTP request body. Some sensitive information may be removed.
    requestBody *string
    // The raw HTTP request URL. Required. Read-only.
    requestUrl *string
    // The collection of Azure Active Directory tenant identifiers for the managed tenants that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
    tenantIds *string
    // The collection of tenant names that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
    tenantNames *string
}
// NewAuditEvent instantiates a new auditEvent and sets the default values.
func NewAuditEvent()(*AuditEvent) {
    m := &AuditEvent{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.auditEvent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuditEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditEvent(), nil
}
// GetActivity gets the activity property value. A string which uniquely represents the operation that occurred. Required. Read-only.
func (m *AuditEvent) GetActivity()(*string) {
    return m.activity
}
// GetActivityDateTime gets the activityDateTime property value. The time when the activity ocurred. Required. Read-only.
func (m *AuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activityDateTime
}
// GetActivityId gets the activityId property value. The identifier of the activity request that made the audit event. Required. Read-only.
func (m *AuditEvent) GetActivityId()(*string) {
    return m.activityId
}
// GetCategory gets the category property value. A category which represents a logical grouping of activities. Required. Read-only.
func (m *AuditEvent) GetCategory()(*string) {
    return m.category
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivity)
    res["activityDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetActivityDateTime)
    res["activityId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivityId)
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCategory)
    res["httpVerb"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHttpVerb)
    res["initiatedByAppId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInitiatedByAppId)
    res["initiatedByUpn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInitiatedByUpn)
    res["initiatedByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInitiatedByUserId)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["requestBody"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRequestBody)
    res["requestUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRequestUrl)
    res["tenantIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantIds)
    res["tenantNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantNames)
    return res
}
// GetHttpVerb gets the httpVerb property value. The HTTP verb that was used when making the API request. Required. Read-only.
func (m *AuditEvent) GetHttpVerb()(*string) {
    return m.httpVerb
}
// GetInitiatedByAppId gets the initiatedByAppId property value. The identifier of the app that was used to make the request. Required. Read-only.
func (m *AuditEvent) GetInitiatedByAppId()(*string) {
    return m.initiatedByAppId
}
// GetInitiatedByUpn gets the initiatedByUpn property value. The UPN of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) GetInitiatedByUpn()(*string) {
    return m.initiatedByUpn
}
// GetInitiatedByUserId gets the initiatedByUserId property value. The identifier of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) GetInitiatedByUserId()(*string) {
    return m.initiatedByUserId
}
// GetIpAddress gets the ipAddress property value. The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
func (m *AuditEvent) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetRequestBody gets the requestBody property value. The raw HTTP request body. Some sensitive information may be removed.
func (m *AuditEvent) GetRequestBody()(*string) {
    return m.requestBody
}
// GetRequestUrl gets the requestUrl property value. The raw HTTP request URL. Required. Read-only.
func (m *AuditEvent) GetRequestUrl()(*string) {
    return m.requestUrl
}
// GetTenantIds gets the tenantIds property value. The collection of Azure Active Directory tenant identifiers for the managed tenants that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) GetTenantIds()(*string) {
    return m.tenantIds
}
// GetTenantNames gets the tenantNames property value. The collection of tenant names that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) GetTenantNames()(*string) {
    return m.tenantNames
}
// Serialize serializes information the current object
func (m *AuditEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityId", m.GetActivityId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("httpVerb", m.GetHttpVerb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByAppId", m.GetInitiatedByAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByUpn", m.GetInitiatedByUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByUserId", m.GetInitiatedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestBody", m.GetRequestBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestUrl", m.GetRequestUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantIds", m.GetTenantIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantNames", m.GetTenantNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. A string which uniquely represents the operation that occurred. Required. Read-only.
func (m *AuditEvent) SetActivity(value *string)() {
    m.activity = value
}
// SetActivityDateTime sets the activityDateTime property value. The time when the activity ocurred. Required. Read-only.
func (m *AuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetActivityId sets the activityId property value. The identifier of the activity request that made the audit event. Required. Read-only.
func (m *AuditEvent) SetActivityId(value *string)() {
    m.activityId = value
}
// SetCategory sets the category property value. A category which represents a logical grouping of activities. Required. Read-only.
func (m *AuditEvent) SetCategory(value *string)() {
    m.category = value
}
// SetHttpVerb sets the httpVerb property value. The HTTP verb that was used when making the API request. Required. Read-only.
func (m *AuditEvent) SetHttpVerb(value *string)() {
    m.httpVerb = value
}
// SetInitiatedByAppId sets the initiatedByAppId property value. The identifier of the app that was used to make the request. Required. Read-only.
func (m *AuditEvent) SetInitiatedByAppId(value *string)() {
    m.initiatedByAppId = value
}
// SetInitiatedByUpn sets the initiatedByUpn property value. The UPN of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) SetInitiatedByUpn(value *string)() {
    m.initiatedByUpn = value
}
// SetInitiatedByUserId sets the initiatedByUserId property value. The identifier of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) SetInitiatedByUserId(value *string)() {
    m.initiatedByUserId = value
}
// SetIpAddress sets the ipAddress property value. The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
func (m *AuditEvent) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetRequestBody sets the requestBody property value. The raw HTTP request body. Some sensitive information may be removed.
func (m *AuditEvent) SetRequestBody(value *string)() {
    m.requestBody = value
}
// SetRequestUrl sets the requestUrl property value. The raw HTTP request URL. Required. Read-only.
func (m *AuditEvent) SetRequestUrl(value *string)() {
    m.requestUrl = value
}
// SetTenantIds sets the tenantIds property value. The collection of Azure Active Directory tenant identifiers for the managed tenants that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) SetTenantIds(value *string)() {
    m.tenantIds = value
}
// SetTenantNames sets the tenantNames property value. The collection of tenant names that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) SetTenantNames(value *string)() {
    m.tenantNames = value
}
