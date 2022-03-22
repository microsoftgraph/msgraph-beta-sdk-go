package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// AuditEvent 
type AuditEvent struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // A string which uniquely represents the operation that occurred. Required. Read-only.
    activity *string;
    // The time when the activity ocurred. Required. Read-only.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier of the activity request that made the audit event. Required. Read-only.
    activityId *string;
    // A category which represents a logical grouping of activities. Required. Read-only.
    category *string;
    // The HTTP verb that was used when making the API request. Required. Read-only.
    httpVerb *string;
    // The identifier of the app that was used to make the request. Required. Read-only.
    initiatedByAppId *string;
    // The UPN of the user who initiated the activity. Required. Read-only.
    initiatedByUpn *string;
    // The identifier of the user who initiated the activity. Required. Read-only.
    initiatedByUserId *string;
    // The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
    ipAddress *string;
    // The raw HTTP request body. Some sensitive information may be removed.
    requestBody *string;
    // The raw HTTP request URL. Required. Read-only.
    requestUrl *string;
    // The collection of Azure Active Directory tenant identifiers for the managed tenants that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
    tenantIds *string;
    // The collection of tenant names that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
    tenantNames *string;
}
// NewAuditEvent instantiates a new auditEvent and sets the default values.
func NewAuditEvent()(*AuditEvent) {
    m := &AuditEvent{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateAuditEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditEventFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAuditEvent(), nil
}
// GetActivity gets the activity property value. A string which uniquely represents the operation that occurred. Required. Read-only.
func (m *AuditEvent) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetActivityDateTime gets the activityDateTime property value. The time when the activity ocurred. Required. Read-only.
func (m *AuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// GetActivityId gets the activityId property value. The identifier of the activity request that made the audit event. Required. Read-only.
func (m *AuditEvent) GetActivityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityId
    }
}
// GetCategory gets the category property value. A category which represents a logical grouping of activities. Required. Read-only.
func (m *AuditEvent) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["activityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityId(val)
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["httpVerb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpVerb(val)
        }
        return nil
    }
    res["initiatedByAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByAppId(val)
        }
        return nil
    }
    res["initiatedByUpn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUpn(val)
        }
        return nil
    }
    res["initiatedByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUserId(val)
        }
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["requestBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestBody(val)
        }
        return nil
    }
    res["requestUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestUrl(val)
        }
        return nil
    }
    res["tenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantIds(val)
        }
        return nil
    }
    res["tenantNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    if m == nil {
        return nil
    } else {
        return m.httpVerb
    }
}
// GetInitiatedByAppId gets the initiatedByAppId property value. The identifier of the app that was used to make the request. Required. Read-only.
func (m *AuditEvent) GetInitiatedByAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedByAppId
    }
}
// GetInitiatedByUpn gets the initiatedByUpn property value. The UPN of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) GetInitiatedByUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedByUpn
    }
}
// GetInitiatedByUserId gets the initiatedByUserId property value. The identifier of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) GetInitiatedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedByUserId
    }
}
// GetIpAddress gets the ipAddress property value. The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
func (m *AuditEvent) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetRequestBody gets the requestBody property value. The raw HTTP request body. Some sensitive information may be removed.
func (m *AuditEvent) GetRequestBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestBody
    }
}
// GetRequestUrl gets the requestUrl property value. The raw HTTP request URL. Required. Read-only.
func (m *AuditEvent) GetRequestUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestUrl
    }
}
// GetTenantIds gets the tenantIds property value. The collection of Azure Active Directory tenant identifiers for the managed tenants that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) GetTenantIds()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantIds
    }
}
// GetTenantNames gets the tenantNames property value. The collection of tenant names that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) GetTenantNames()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantNames
    }
}
// Serialize serializes information the current object
func (m *AuditEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.activity = value
    }
}
// SetActivityDateTime sets the activityDateTime property value. The time when the activity ocurred. Required. Read-only.
func (m *AuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.activityDateTime = value
    }
}
// SetActivityId sets the activityId property value. The identifier of the activity request that made the audit event. Required. Read-only.
func (m *AuditEvent) SetActivityId(value *string)() {
    if m != nil {
        m.activityId = value
    }
}
// SetCategory sets the category property value. A category which represents a logical grouping of activities. Required. Read-only.
func (m *AuditEvent) SetCategory(value *string)() {
    if m != nil {
        m.category = value
    }
}
// SetHttpVerb sets the httpVerb property value. The HTTP verb that was used when making the API request. Required. Read-only.
func (m *AuditEvent) SetHttpVerb(value *string)() {
    if m != nil {
        m.httpVerb = value
    }
}
// SetInitiatedByAppId sets the initiatedByAppId property value. The identifier of the app that was used to make the request. Required. Read-only.
func (m *AuditEvent) SetInitiatedByAppId(value *string)() {
    if m != nil {
        m.initiatedByAppId = value
    }
}
// SetInitiatedByUpn sets the initiatedByUpn property value. The UPN of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) SetInitiatedByUpn(value *string)() {
    if m != nil {
        m.initiatedByUpn = value
    }
}
// SetInitiatedByUserId sets the initiatedByUserId property value. The identifier of the user who initiated the activity. Required. Read-only.
func (m *AuditEvent) SetInitiatedByUserId(value *string)() {
    if m != nil {
        m.initiatedByUserId = value
    }
}
// SetIpAddress sets the ipAddress property value. The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
func (m *AuditEvent) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetRequestBody sets the requestBody property value. The raw HTTP request body. Some sensitive information may be removed.
func (m *AuditEvent) SetRequestBody(value *string)() {
    if m != nil {
        m.requestBody = value
    }
}
// SetRequestUrl sets the requestUrl property value. The raw HTTP request URL. Required. Read-only.
func (m *AuditEvent) SetRequestUrl(value *string)() {
    if m != nil {
        m.requestUrl = value
    }
}
// SetTenantIds sets the tenantIds property value. The collection of Azure Active Directory tenant identifiers for the managed tenants that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) SetTenantIds(value *string)() {
    if m != nil {
        m.tenantIds = value
    }
}
// SetTenantNames sets the tenantNames property value. The collection of tenant names that were impacted by this change. This is formatted as a list of comma-separated values. Required. Read-only.
func (m *AuditEvent) SetTenantNames(value *string)() {
    if m != nil {
        m.tenantNames = value
    }
}
