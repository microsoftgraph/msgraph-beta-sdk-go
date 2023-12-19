package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AuditLogQuery 
type AuditLogQuery struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAuditLogQuery instantiates a new auditLogQuery and sets the default values.
func NewAuditLogQuery()(*AuditLogQuery) {
    m := &AuditLogQuery{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAuditLogQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogQuery(), nil
}
// GetAdministrativeUnitIdFilters gets the administrativeUnitIdFilters property value. The administrativeUnitIdFilters property
func (m *AuditLogQuery) GetAdministrativeUnitIdFilters()([]string) {
    val, err := m.GetBackingStore().Get("administrativeUnitIdFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AuditLogQuery) GetDisplayName()(*string) {
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
func (m *AuditLogQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnitIdFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAdministrativeUnitIdFilters(res)
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
    res["filterEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterEndDateTime(val)
        }
        return nil
    }
    res["filterStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterStartDateTime(val)
        }
        return nil
    }
    res["ipAddressFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIpAddressFilters(res)
        }
        return nil
    }
    res["keywordFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeywordFilter(val)
        }
        return nil
    }
    res["objectIdFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetObjectIdFilters(res)
        }
        return nil
    }
    res["operationFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOperationFilters(res)
        }
        return nil
    }
    res["records"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditLogRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditLogRecordable)
                }
            }
            m.SetRecords(res)
        }
        return nil
    }
    res["recordTypeFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAuditLogQuery_recordTypeFilters)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogQuery_recordTypeFilters, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*AuditLogQuery_recordTypeFilters))
                }
            }
            m.SetRecordTypeFilters(res)
        }
        return nil
    }
    res["serviceFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetServiceFilters(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditLogQuery_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AuditLogQuery_status))
        }
        return nil
    }
    res["userPrincipalNameFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUserPrincipalNameFilters(res)
        }
        return nil
    }
    return res
}
// GetFilterEndDateTime gets the filterEndDateTime property value. The filterEndDateTime property
func (m *AuditLogQuery) GetFilterEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("filterEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFilterStartDateTime gets the filterStartDateTime property value. The filterStartDateTime property
func (m *AuditLogQuery) GetFilterStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("filterStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetIpAddressFilters gets the ipAddressFilters property value. The ipAddressFilters property
func (m *AuditLogQuery) GetIpAddressFilters()([]string) {
    val, err := m.GetBackingStore().Get("ipAddressFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKeywordFilter gets the keywordFilter property value. The keywordFilter property
func (m *AuditLogQuery) GetKeywordFilter()(*string) {
    val, err := m.GetBackingStore().Get("keywordFilter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetObjectIdFilters gets the objectIdFilters property value. The objectIdFilters property
func (m *AuditLogQuery) GetObjectIdFilters()([]string) {
    val, err := m.GetBackingStore().Get("objectIdFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOperationFilters gets the operationFilters property value. The operationFilters property
func (m *AuditLogQuery) GetOperationFilters()([]string) {
    val, err := m.GetBackingStore().Get("operationFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRecords gets the records property value. The records property
func (m *AuditLogQuery) GetRecords()([]AuditLogRecordable) {
    val, err := m.GetBackingStore().Get("records")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuditLogRecordable)
    }
    return nil
}
// GetRecordTypeFilters gets the recordTypeFilters property value. The recordTypeFilters property
func (m *AuditLogQuery) GetRecordTypeFilters()([]AuditLogQuery_recordTypeFilters) {
    val, err := m.GetBackingStore().Get("recordTypeFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuditLogQuery_recordTypeFilters)
    }
    return nil
}
// GetServiceFilters gets the serviceFilters property value. The serviceFilters property
func (m *AuditLogQuery) GetServiceFilters()([]string) {
    val, err := m.GetBackingStore().Get("serviceFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *AuditLogQuery) GetStatus()(*AuditLogQuery_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuditLogQuery_status)
    }
    return nil
}
// GetUserPrincipalNameFilters gets the userPrincipalNameFilters property value. The userPrincipalNameFilters property
func (m *AuditLogQuery) GetUserPrincipalNameFilters()([]string) {
    val, err := m.GetBackingStore().Get("userPrincipalNameFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuditLogQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnitIdFilters() != nil {
        err = writer.WriteCollectionOfStringValues("administrativeUnitIdFilters", m.GetAdministrativeUnitIdFilters())
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
        err = writer.WriteTimeValue("filterEndDateTime", m.GetFilterEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("filterStartDateTime", m.GetFilterStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetIpAddressFilters() != nil {
        err = writer.WriteCollectionOfStringValues("ipAddressFilters", m.GetIpAddressFilters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keywordFilter", m.GetKeywordFilter())
        if err != nil {
            return err
        }
    }
    if m.GetObjectIdFilters() != nil {
        err = writer.WriteCollectionOfStringValues("objectIdFilters", m.GetObjectIdFilters())
        if err != nil {
            return err
        }
    }
    if m.GetOperationFilters() != nil {
        err = writer.WriteCollectionOfStringValues("operationFilters", m.GetOperationFilters())
        if err != nil {
            return err
        }
    }
    if m.GetRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecords()))
        for i, v := range m.GetRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("records", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecordTypeFilters() != nil {
        err = writer.WriteCollectionOfStringValues("recordTypeFilters", SerializeAuditLogQuery_recordTypeFilters(m.GetRecordTypeFilters()))
        if err != nil {
            return err
        }
    }
    if m.GetServiceFilters() != nil {
        err = writer.WriteCollectionOfStringValues("serviceFilters", m.GetServiceFilters())
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
    if m.GetUserPrincipalNameFilters() != nil {
        err = writer.WriteCollectionOfStringValues("userPrincipalNameFilters", m.GetUserPrincipalNameFilters())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnitIdFilters sets the administrativeUnitIdFilters property value. The administrativeUnitIdFilters property
func (m *AuditLogQuery) SetAdministrativeUnitIdFilters(value []string)() {
    err := m.GetBackingStore().Set("administrativeUnitIdFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuditLogQuery) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFilterEndDateTime sets the filterEndDateTime property value. The filterEndDateTime property
func (m *AuditLogQuery) SetFilterEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("filterEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFilterStartDateTime sets the filterStartDateTime property value. The filterStartDateTime property
func (m *AuditLogQuery) SetFilterStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("filterStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddressFilters sets the ipAddressFilters property value. The ipAddressFilters property
func (m *AuditLogQuery) SetIpAddressFilters(value []string)() {
    err := m.GetBackingStore().Set("ipAddressFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetKeywordFilter sets the keywordFilter property value. The keywordFilter property
func (m *AuditLogQuery) SetKeywordFilter(value *string)() {
    err := m.GetBackingStore().Set("keywordFilter", value)
    if err != nil {
        panic(err)
    }
}
// SetObjectIdFilters sets the objectIdFilters property value. The objectIdFilters property
func (m *AuditLogQuery) SetObjectIdFilters(value []string)() {
    err := m.GetBackingStore().Set("objectIdFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetOperationFilters sets the operationFilters property value. The operationFilters property
func (m *AuditLogQuery) SetOperationFilters(value []string)() {
    err := m.GetBackingStore().Set("operationFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetRecords sets the records property value. The records property
func (m *AuditLogQuery) SetRecords(value []AuditLogRecordable)() {
    err := m.GetBackingStore().Set("records", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordTypeFilters sets the recordTypeFilters property value. The recordTypeFilters property
func (m *AuditLogQuery) SetRecordTypeFilters(value []AuditLogQuery_recordTypeFilters)() {
    err := m.GetBackingStore().Set("recordTypeFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceFilters sets the serviceFilters property value. The serviceFilters property
func (m *AuditLogQuery) SetServiceFilters(value []string)() {
    err := m.GetBackingStore().Set("serviceFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *AuditLogQuery) SetStatus(value *AuditLogQuery_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalNameFilters sets the userPrincipalNameFilters property value. The userPrincipalNameFilters property
func (m *AuditLogQuery) SetUserPrincipalNameFilters(value []string)() {
    err := m.GetBackingStore().Set("userPrincipalNameFilters", value)
    if err != nil {
        panic(err)
    }
}
// AuditLogQueryable 
type AuditLogQueryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnitIdFilters()([]string)
    GetDisplayName()(*string)
    GetFilterEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilterStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIpAddressFilters()([]string)
    GetKeywordFilter()(*string)
    GetObjectIdFilters()([]string)
    GetOperationFilters()([]string)
    GetRecords()([]AuditLogRecordable)
    GetRecordTypeFilters()([]AuditLogQuery_recordTypeFilters)
    GetServiceFilters()([]string)
    GetStatus()(*AuditLogQuery_status)
    GetUserPrincipalNameFilters()([]string)
    SetAdministrativeUnitIdFilters(value []string)()
    SetDisplayName(value *string)()
    SetFilterEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilterStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIpAddressFilters(value []string)()
    SetKeywordFilter(value *string)()
    SetObjectIdFilters(value []string)()
    SetOperationFilters(value []string)()
    SetRecords(value []AuditLogRecordable)()
    SetRecordTypeFilters(value []AuditLogQuery_recordTypeFilters)()
    SetServiceFilters(value []string)()
    SetStatus(value *AuditLogQuery_status)()
    SetUserPrincipalNameFilters(value []string)()
}
