package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type DiscoveredCloudAppDetail struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDiscoveredCloudAppDetail instantiates a new DiscoveredCloudAppDetail and sets the default values.
func NewDiscoveredCloudAppDetail()(*DiscoveredCloudAppDetail) {
    m := &DiscoveredCloudAppDetail{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDiscoveredCloudAppDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDiscoveredCloudAppDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.endpointDiscoveredCloudAppDetail":
                        return NewEndpointDiscoveredCloudAppDetail(), nil
                }
            }
        }
    }
    return NewDiscoveredCloudAppDetail(), nil
}
// GetAppInfo gets the appInfo property value. The application information.
// returns a DiscoveredCloudAppInfoable when successful
func (m *DiscoveredCloudAppDetail) GetAppInfo()(DiscoveredCloudAppInfoable) {
    val, err := m.GetBackingStore().Get("appInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DiscoveredCloudAppInfoable)
    }
    return nil
}
// GetCategory gets the category property value. The category property
// returns a *AppCategory when successful
func (m *DiscoveredCloudAppDetail) GetCategory()(*AppCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppCategory)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *DiscoveredCloudAppDetail) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The app name.
// returns a *string when successful
func (m *DiscoveredCloudAppDetail) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomains gets the domains property value. The domain.
// returns a []string when successful
func (m *DiscoveredCloudAppDetail) GetDomains()([]string) {
    val, err := m.GetBackingStore().Get("domains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDownloadNetworkTrafficInBytes gets the downloadNetworkTrafficInBytes property value. The download traffic size.
// returns a *int64 when successful
func (m *DiscoveredCloudAppDetail) GetDownloadNetworkTrafficInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("downloadNetworkTrafficInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DiscoveredCloudAppDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDiscoveredCloudAppInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(DiscoveredCloudAppInfoable))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*AppCategory))
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
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["downloadNetworkTrafficInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadNetworkTrafficInBytes(val)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["ipAddressCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressCount(val)
        }
        return nil
    }
    res["ipAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredCloudAppIPAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredCloudAppIPAddressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DiscoveredCloudAppIPAddressable)
                }
            }
            m.SetIpAddresses(res)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    res["transactionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionCount(val)
        }
        return nil
    }
    res["uploadNetworkTrafficInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadNetworkTrafficInBytes(val)
        }
        return nil
    }
    res["userCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCount(val)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredCloudAppUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredCloudAppUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DiscoveredCloudAppUserable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
// returns a *Time when successful
func (m *DiscoveredCloudAppDetail) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetIpAddressCount gets the ipAddressCount property value. The IP address.
// returns a *int64 when successful
func (m *DiscoveredCloudAppDetail) GetIpAddressCount()(*int64) {
    val, err := m.GetBackingStore().Get("ipAddressCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetIpAddresses gets the ipAddresses property value. The list of IP addresses accessed by the app.
// returns a []DiscoveredCloudAppIPAddressable when successful
func (m *DiscoveredCloudAppDetail) GetIpAddresses()([]DiscoveredCloudAppIPAddressable) {
    val, err := m.GetBackingStore().Get("ipAddresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DiscoveredCloudAppIPAddressable)
    }
    return nil
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The last seen date of the discovered app. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *DiscoveredCloudAppDetail) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRiskScore gets the riskScore property value. The risk score of the app.
// returns a *int64 when successful
func (m *DiscoveredCloudAppDetail) GetRiskScore()(*int64) {
    val, err := m.GetBackingStore().Get("riskScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTags gets the tags property value. The tags applied to an app. Possible values include Unsanctioned, Sanctioned, Monitored, or a custom value.
// returns a []string when successful
func (m *DiscoveredCloudAppDetail) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTransactionCount gets the transactionCount property value. The app transaction count.
// returns a *int64 when successful
func (m *DiscoveredCloudAppDetail) GetTransactionCount()(*int64) {
    val, err := m.GetBackingStore().Get("transactionCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUploadNetworkTrafficInBytes gets the uploadNetworkTrafficInBytes property value. The app upload traffic size, in bytes.
// returns a *int64 when successful
func (m *DiscoveredCloudAppDetail) GetUploadNetworkTrafficInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("uploadNetworkTrafficInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUserCount gets the userCount property value. The count of users who use the app.
// returns a *int64 when successful
func (m *DiscoveredCloudAppDetail) GetUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("userCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUsers gets the users property value. The list of users who access the app.
// returns a []DiscoveredCloudAppUserable when successful
func (m *DiscoveredCloudAppDetail) GetUsers()([]DiscoveredCloudAppUserable) {
    val, err := m.GetBackingStore().Get("users")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DiscoveredCloudAppUserable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DiscoveredCloudAppDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
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
    if m.GetDomains() != nil {
        err = writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("downloadNetworkTrafficInBytes", m.GetDownloadNetworkTrafficInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("ipAddressCount", m.GetIpAddressCount())
        if err != nil {
            return err
        }
    }
    if m.GetIpAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpAddresses()))
        for i, v := range m.GetIpAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ipAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("transactionCount", m.GetTransactionCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("uploadNetworkTrafficInBytes", m.GetUploadNetworkTrafficInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("userCount", m.GetUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppInfo sets the appInfo property value. The application information.
func (m *DiscoveredCloudAppDetail) SetAppInfo(value DiscoveredCloudAppInfoable)() {
    err := m.GetBackingStore().Set("appInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *DiscoveredCloudAppDetail) SetCategory(value *AppCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *DiscoveredCloudAppDetail) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The app name.
func (m *DiscoveredCloudAppDetail) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomains sets the domains property value. The domain.
func (m *DiscoveredCloudAppDetail) SetDomains(value []string)() {
    err := m.GetBackingStore().Set("domains", value)
    if err != nil {
        panic(err)
    }
}
// SetDownloadNetworkTrafficInBytes sets the downloadNetworkTrafficInBytes property value. The download traffic size.
func (m *DiscoveredCloudAppDetail) SetDownloadNetworkTrafficInBytes(value *int64)() {
    err := m.GetBackingStore().Set("downloadNetworkTrafficInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DiscoveredCloudAppDetail) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddressCount sets the ipAddressCount property value. The IP address.
func (m *DiscoveredCloudAppDetail) SetIpAddressCount(value *int64)() {
    err := m.GetBackingStore().Set("ipAddressCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddresses sets the ipAddresses property value. The list of IP addresses accessed by the app.
func (m *DiscoveredCloudAppDetail) SetIpAddresses(value []DiscoveredCloudAppIPAddressable)() {
    err := m.GetBackingStore().Set("ipAddresses", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The last seen date of the discovered app. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DiscoveredCloudAppDetail) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskScore sets the riskScore property value. The risk score of the app.
func (m *DiscoveredCloudAppDetail) SetRiskScore(value *int64)() {
    err := m.GetBackingStore().Set("riskScore", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. The tags applied to an app. Possible values include Unsanctioned, Sanctioned, Monitored, or a custom value.
func (m *DiscoveredCloudAppDetail) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
// SetTransactionCount sets the transactionCount property value. The app transaction count.
func (m *DiscoveredCloudAppDetail) SetTransactionCount(value *int64)() {
    err := m.GetBackingStore().Set("transactionCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUploadNetworkTrafficInBytes sets the uploadNetworkTrafficInBytes property value. The app upload traffic size, in bytes.
func (m *DiscoveredCloudAppDetail) SetUploadNetworkTrafficInBytes(value *int64)() {
    err := m.GetBackingStore().Set("uploadNetworkTrafficInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCount sets the userCount property value. The count of users who use the app.
func (m *DiscoveredCloudAppDetail) SetUserCount(value *int64)() {
    err := m.GetBackingStore().Set("userCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUsers sets the users property value. The list of users who access the app.
func (m *DiscoveredCloudAppDetail) SetUsers(value []DiscoveredCloudAppUserable)() {
    err := m.GetBackingStore().Set("users", value)
    if err != nil {
        panic(err)
    }
}
type DiscoveredCloudAppDetailable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(DiscoveredCloudAppInfoable)
    GetCategory()(*AppCategory)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomains()([]string)
    GetDownloadNetworkTrafficInBytes()(*int64)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIpAddressCount()(*int64)
    GetIpAddresses()([]DiscoveredCloudAppIPAddressable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRiskScore()(*int64)
    GetTags()([]string)
    GetTransactionCount()(*int64)
    GetUploadNetworkTrafficInBytes()(*int64)
    GetUserCount()(*int64)
    GetUsers()([]DiscoveredCloudAppUserable)
    SetAppInfo(value DiscoveredCloudAppInfoable)()
    SetCategory(value *AppCategory)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomains(value []string)()
    SetDownloadNetworkTrafficInBytes(value *int64)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIpAddressCount(value *int64)()
    SetIpAddresses(value []DiscoveredCloudAppIPAddressable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRiskScore(value *int64)()
    SetTags(value []string)()
    SetTransactionCount(value *int64)()
    SetUploadNetworkTrafficInBytes(value *int64)()
    SetUserCount(value *int64)()
    SetUsers(value []DiscoveredCloudAppUserable)()
}
