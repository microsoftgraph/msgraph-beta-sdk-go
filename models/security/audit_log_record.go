package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AuditLogRecord 
type AuditLogRecord struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAuditLogRecord instantiates a new auditLogRecord and sets the default values.
func NewAuditLogRecord()(*AuditLogRecord) {
    m := &AuditLogRecord{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAuditLogRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogRecord(), nil
}
// GetAdministrativeUnits gets the administrativeUnits property value. The administrativeUnits property
func (m *AuditLogRecord) GetAdministrativeUnits()([]string) {
    val, err := m.GetBackingStore().Get("administrativeUnits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAuditData gets the auditData property value. The auditData property
func (m *AuditLogRecord) GetAuditData()(AuditDataable) {
    val, err := m.GetBackingStore().Get("auditData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuditDataable)
    }
    return nil
}
// GetAuditLogRecordType gets the auditLogRecordType property value. The auditLogRecordType property
func (m *AuditLogRecord) GetAuditLogRecordType()(*AuditLogRecord_auditLogRecordType) {
    val, err := m.GetBackingStore().Get("auditLogRecordType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuditLogRecord_auditLogRecordType)
    }
    return nil
}
// GetClientIp gets the clientIp property value. The clientIp property
func (m *AuditLogRecord) GetClientIp()(*string) {
    val, err := m.GetBackingStore().Get("clientIp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *AuditLogRecord) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditLogRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAdministrativeUnits(res)
        }
        return nil
    }
    res["auditData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditData(val.(AuditDataable))
        }
        return nil
    }
    res["auditLogRecordType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditLogRecord_auditLogRecordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditLogRecordType(val.(*AuditLogRecord_auditLogRecordType))
        }
        return nil
    }
    res["clientIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientIp(val)
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
    res["objectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["operation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperation(val)
        }
        return nil
    }
    res["organizationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationId(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditLogRecord_userType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*AuditLogRecord_userType))
        }
        return nil
    }
    return res
}
// GetObjectId gets the objectId property value. The objectId property
func (m *AuditLogRecord) GetObjectId()(*string) {
    val, err := m.GetBackingStore().Get("objectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperation gets the operation property value. The operation property
func (m *AuditLogRecord) GetOperation()(*string) {
    val, err := m.GetBackingStore().Get("operation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationId gets the organizationId property value. The organizationId property
func (m *AuditLogRecord) GetOrganizationId()(*string) {
    val, err := m.GetBackingStore().Get("organizationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetService gets the service property value. The service property
func (m *AuditLogRecord) GetService()(*string) {
    val, err := m.GetBackingStore().Get("service")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The userId property
func (m *AuditLogRecord) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *AuditLogRecord) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserType gets the userType property value. The userType property
func (m *AuditLogRecord) GetUserType()(*AuditLogRecord_userType) {
    val, err := m.GetBackingStore().Get("userType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuditLogRecord_userType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuditLogRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnits() != nil {
        err = writer.WriteCollectionOfStringValues("administrativeUnits", m.GetAdministrativeUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("auditData", m.GetAuditData())
        if err != nil {
            return err
        }
    }
    if m.GetAuditLogRecordType() != nil {
        cast := (*m.GetAuditLogRecordType()).String()
        err = writer.WriteStringValue("auditLogRecordType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientIp", m.GetClientIp())
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
        err = writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operation", m.GetOperation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationId", m.GetOrganizationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnits sets the administrativeUnits property value. The administrativeUnits property
func (m *AuditLogRecord) SetAdministrativeUnits(value []string)() {
    err := m.GetBackingStore().Set("administrativeUnits", value)
    if err != nil {
        panic(err)
    }
}
// SetAuditData sets the auditData property value. The auditData property
func (m *AuditLogRecord) SetAuditData(value AuditDataable)() {
    err := m.GetBackingStore().Set("auditData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuditLogRecordType sets the auditLogRecordType property value. The auditLogRecordType property
func (m *AuditLogRecord) SetAuditLogRecordType(value *AuditLogRecord_auditLogRecordType)() {
    err := m.GetBackingStore().Set("auditLogRecordType", value)
    if err != nil {
        panic(err)
    }
}
// SetClientIp sets the clientIp property value. The clientIp property
func (m *AuditLogRecord) SetClientIp(value *string)() {
    err := m.GetBackingStore().Set("clientIp", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *AuditLogRecord) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetObjectId sets the objectId property value. The objectId property
func (m *AuditLogRecord) SetObjectId(value *string)() {
    err := m.GetBackingStore().Set("objectId", value)
    if err != nil {
        panic(err)
    }
}
// SetOperation sets the operation property value. The operation property
func (m *AuditLogRecord) SetOperation(value *string)() {
    err := m.GetBackingStore().Set("operation", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationId sets the organizationId property value. The organizationId property
func (m *AuditLogRecord) SetOrganizationId(value *string)() {
    err := m.GetBackingStore().Set("organizationId", value)
    if err != nil {
        panic(err)
    }
}
// SetService sets the service property value. The service property
func (m *AuditLogRecord) SetService(value *string)() {
    err := m.GetBackingStore().Set("service", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The userId property
func (m *AuditLogRecord) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *AuditLogRecord) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserType sets the userType property value. The userType property
func (m *AuditLogRecord) SetUserType(value *AuditLogRecord_userType)() {
    err := m.GetBackingStore().Set("userType", value)
    if err != nil {
        panic(err)
    }
}
// AuditLogRecordable 
type AuditLogRecordable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnits()([]string)
    GetAuditData()(AuditDataable)
    GetAuditLogRecordType()(*AuditLogRecord_auditLogRecordType)
    GetClientIp()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetObjectId()(*string)
    GetOperation()(*string)
    GetOrganizationId()(*string)
    GetService()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetUserType()(*AuditLogRecord_userType)
    SetAdministrativeUnits(value []string)()
    SetAuditData(value AuditDataable)()
    SetAuditLogRecordType(value *AuditLogRecord_auditLogRecordType)()
    SetClientIp(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetObjectId(value *string)()
    SetOperation(value *string)()
    SetOrganizationId(value *string)()
    SetService(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *AuditLogRecord_userType)()
}
