package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AuditEvent 
type AuditEvent struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAuditEvent instantiates a new auditEvent and sets the default values.
func NewAuditEvent()(*AuditEvent) {
    m := &AuditEvent{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAuditEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditEvent(), nil
}
// GetActivity gets the activity property value. A string that uniquely represents the operation that occurred. Required. Read-only.
func (m *AuditEvent) GetActivity()(*string) {
    val, err := m.GetBackingStore().Get("activity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetActivityDateTime gets the activityDateTime property value. The time when the activity occurred. Required. Read-only.
func (m *AuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("activityDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetActivityId gets the activityId property value. The identifier of the activity request that made the audit event. Required. Read-only.
func (m *AuditEvent) GetActivityId()(*string) {
    val, err := m.GetBackingStore().Get("activityId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCategory gets the category property value. A category that represents a logical grouping of activities. Required. Read-only.
func (m *AuditEvent) GetCategory()(*string) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["activityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["activityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityId(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["httpVerb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpVerb(val)
        }
        return nil
    }
    res["initiatedByAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByAppId(val)
        }
        return nil
    }
    res["initiatedByUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUpn(val)
        }
        return nil
    }
    res["initiatedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUserId(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["requestBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestBody(val)
        }
        return nil
    }
    res["requestUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestUrl(val)
        }
        return nil
    }
    res["tenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantIds(val)
        }
        return nil
    }
    res["tenantNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantNames(val)
        }
        return nil
    }
    return res
}
// GetHttpVerb gets the httpVerb property value. The HTTP verb that was used when making the API request. Required. Read-only.
func (m *AuditEvent) GetHttpVerb()(*string) {
    val, err := m.GetBackingStore().Get("httpVerb")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInitiatedByAppId gets the initiatedByAppId property value. The identifier of the app that was used to make the request. Required. Read-only.
func (m *AuditEvent) GetInitiatedByAppId()(*string) {
    val, err := m.GetBackingStore().Get("initiatedByAppId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInitiatedByUpn gets the initiatedByUpn property value. The UPN of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) GetInitiatedByUpn()(*string) {
    val, err := m.GetBackingStore().Get("initiatedByUpn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInitiatedByUserId gets the initiatedByUserId property value. The identifier of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) GetInitiatedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("initiatedByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpAddress gets the ipAddress property value. The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
func (m *AuditEvent) GetIpAddress()(*string) {
    val, err := m.GetBackingStore().Get("ipAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestBody gets the requestBody property value. The raw HTTP request body. Some sensitive information may be removed.
func (m *AuditEvent) GetRequestBody()(*string) {
    val, err := m.GetBackingStore().Get("requestBody")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestUrl gets the requestUrl property value. The raw HTTP request URL. Required. Read-only.
func (m *AuditEvent) GetRequestUrl()(*string) {
    val, err := m.GetBackingStore().Get("requestUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantIds gets the tenantIds property value. The collection of Azure Active Directory tenant identifiers for the managed tenants that were affected by a change, and is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) GetTenantIds()(*string) {
    val, err := m.GetBackingStore().Get("tenantIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantNames gets the tenantNames property value. The collection of tenant names that were affected by a change, and is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) GetTenantNames()(*string) {
    val, err := m.GetBackingStore().Get("tenantNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
// SetActivity sets the activity property value. A string that uniquely represents the operation that occurred. Required. Read-only.
func (m *AuditEvent) SetActivity(value *string)() {
    err := m.GetBackingStore().Set("activity", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityDateTime sets the activityDateTime property value. The time when the activity occurred. Required. Read-only.
func (m *AuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("activityDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityId sets the activityId property value. The identifier of the activity request that made the audit event. Required. Read-only.
func (m *AuditEvent) SetActivityId(value *string)() {
    err := m.GetBackingStore().Set("activityId", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. A category that represents a logical grouping of activities. Required. Read-only.
func (m *AuditEvent) SetCategory(value *string)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetHttpVerb sets the httpVerb property value. The HTTP verb that was used when making the API request. Required. Read-only.
func (m *AuditEvent) SetHttpVerb(value *string)() {
    err := m.GetBackingStore().Set("httpVerb", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatedByAppId sets the initiatedByAppId property value. The identifier of the app that was used to make the request. Required. Read-only.
func (m *AuditEvent) SetInitiatedByAppId(value *string)() {
    err := m.GetBackingStore().Set("initiatedByAppId", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatedByUpn sets the initiatedByUpn property value. The UPN of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) SetInitiatedByUpn(value *string)() {
    err := m.GetBackingStore().Set("initiatedByUpn", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatedByUserId sets the initiatedByUserId property value. The identifier of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) SetInitiatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("initiatedByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddress sets the ipAddress property value. The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
func (m *AuditEvent) SetIpAddress(value *string)() {
    err := m.GetBackingStore().Set("ipAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestBody sets the requestBody property value. The raw HTTP request body. Some sensitive information may be removed.
func (m *AuditEvent) SetRequestBody(value *string)() {
    err := m.GetBackingStore().Set("requestBody", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestUrl sets the requestUrl property value. The raw HTTP request URL. Required. Read-only.
func (m *AuditEvent) SetRequestUrl(value *string)() {
    err := m.GetBackingStore().Set("requestUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantIds sets the tenantIds property value. The collection of Azure Active Directory tenant identifiers for the managed tenants that were affected by a change, and is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) SetTenantIds(value *string)() {
    err := m.GetBackingStore().Set("tenantIds", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantNames sets the tenantNames property value. The collection of tenant names that were affected by a change, and is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) SetTenantNames(value *string)() {
    err := m.GetBackingStore().Set("tenantNames", value)
    if err != nil {
        panic(err)
    }
}
// AuditEventable 
type AuditEventable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityId()(*string)
    GetCategory()(*string)
    GetHttpVerb()(*string)
    GetInitiatedByAppId()(*string)
    GetInitiatedByUpn()(*string)
    GetInitiatedByUserId()(*string)
    GetIpAddress()(*string)
    GetRequestBody()(*string)
    GetRequestUrl()(*string)
    GetTenantIds()(*string)
    GetTenantNames()(*string)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityId(value *string)()
    SetCategory(value *string)()
    SetHttpVerb(value *string)()
    SetInitiatedByAppId(value *string)()
    SetInitiatedByUpn(value *string)()
    SetInitiatedByUserId(value *string)()
    SetIpAddress(value *string)()
    SetRequestBody(value *string)()
    SetRequestUrl(value *string)()
    SetTenantIds(value *string)()
    SetTenantNames(value *string)()
}
