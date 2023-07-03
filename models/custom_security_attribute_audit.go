package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomSecurityAttributeAudit 
type CustomSecurityAttributeAudit struct {
    Entity
}
// NewCustomSecurityAttributeAudit instantiates a new CustomSecurityAttributeAudit and sets the default values.
func NewCustomSecurityAttributeAudit()(*CustomSecurityAttributeAudit) {
    m := &CustomSecurityAttributeAudit{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomSecurityAttributeAuditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomSecurityAttributeAuditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomSecurityAttributeAudit(), nil
}
// GetActivityDateTime gets the activityDateTime property value. The activityDateTime property
func (m *CustomSecurityAttributeAudit) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("activityDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetActivityDisplayName gets the activityDisplayName property value. The activityDisplayName property
func (m *CustomSecurityAttributeAudit) GetActivityDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("activityDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdditionalDetails gets the additionalDetails property value. The additionalDetails property
func (m *CustomSecurityAttributeAudit) GetAdditionalDetails()([]KeyValueable) {
    val, err := m.GetBackingStore().Get("additionalDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValueable)
    }
    return nil
}
// GetCategory gets the category property value. The category property
func (m *CustomSecurityAttributeAudit) GetCategory()(*string) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCorrelationId gets the correlationId property value. The correlationId property
func (m *CustomSecurityAttributeAudit) GetCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("correlationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomSecurityAttributeAudit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["activityDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDisplayName(val)
        }
        return nil
    }
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValueable)
                }
            }
            m.SetAdditionalDetails(res)
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
    res["initiatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditActivityInitiatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val.(AuditActivityInitiatorable))
        }
        return nil
    }
    res["loggedByService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggedByService(val)
        }
        return nil
    }
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val)
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*OperationResult))
        }
        return nil
    }
    res["resultReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultReason(val)
        }
        return nil
    }
    res["targetResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TargetResourceable)
                }
            }
            m.SetTargetResources(res)
        }
        return nil
    }
    res["userAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val)
        }
        return nil
    }
    return res
}
// GetInitiatedBy gets the initiatedBy property value. The initiatedBy property
func (m *CustomSecurityAttributeAudit) GetInitiatedBy()(AuditActivityInitiatorable) {
    val, err := m.GetBackingStore().Get("initiatedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuditActivityInitiatorable)
    }
    return nil
}
// GetLoggedByService gets the loggedByService property value. The loggedByService property
func (m *CustomSecurityAttributeAudit) GetLoggedByService()(*string) {
    val, err := m.GetBackingStore().Get("loggedByService")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperationType gets the operationType property value. The operationType property
func (m *CustomSecurityAttributeAudit) GetOperationType()(*string) {
    val, err := m.GetBackingStore().Get("operationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResult gets the result property value. The result property
func (m *CustomSecurityAttributeAudit) GetResult()(*OperationResult) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperationResult)
    }
    return nil
}
// GetResultReason gets the resultReason property value. The resultReason property
func (m *CustomSecurityAttributeAudit) GetResultReason()(*string) {
    val, err := m.GetBackingStore().Get("resultReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetResources gets the targetResources property value. The targetResources property
func (m *CustomSecurityAttributeAudit) GetTargetResources()([]TargetResourceable) {
    val, err := m.GetBackingStore().Get("targetResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TargetResourceable)
    }
    return nil
}
// GetUserAgent gets the userAgent property value. The userAgent property
func (m *CustomSecurityAttributeAudit) GetUserAgent()(*string) {
    val, err := m.GetBackingStore().Get("userAgent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomSecurityAttributeAudit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityDisplayName", m.GetActivityDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetAdditionalDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalDetails()))
        for i, v := range m.GetAdditionalDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("additionalDetails", cast)
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
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loggedByService", m.GetLoggedByService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operationType", m.GetOperationType())
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err = writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultReason", m.GetResultReason())
        if err != nil {
            return err
        }
    }
    if m.GetTargetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetResources()))
        for i, v := range m.GetTargetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("targetResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userAgent", m.GetUserAgent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityDateTime sets the activityDateTime property value. The activityDateTime property
func (m *CustomSecurityAttributeAudit) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("activityDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityDisplayName sets the activityDisplayName property value. The activityDisplayName property
func (m *CustomSecurityAttributeAudit) SetActivityDisplayName(value *string)() {
    err := m.GetBackingStore().Set("activityDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalDetails sets the additionalDetails property value. The additionalDetails property
func (m *CustomSecurityAttributeAudit) SetAdditionalDetails(value []KeyValueable)() {
    err := m.GetBackingStore().Set("additionalDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *CustomSecurityAttributeAudit) SetCategory(value *string)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationId sets the correlationId property value. The correlationId property
func (m *CustomSecurityAttributeAudit) SetCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("correlationId", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatedBy sets the initiatedBy property value. The initiatedBy property
func (m *CustomSecurityAttributeAudit) SetInitiatedBy(value AuditActivityInitiatorable)() {
    err := m.GetBackingStore().Set("initiatedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLoggedByService sets the loggedByService property value. The loggedByService property
func (m *CustomSecurityAttributeAudit) SetLoggedByService(value *string)() {
    err := m.GetBackingStore().Set("loggedByService", value)
    if err != nil {
        panic(err)
    }
}
// SetOperationType sets the operationType property value. The operationType property
func (m *CustomSecurityAttributeAudit) SetOperationType(value *string)() {
    err := m.GetBackingStore().Set("operationType", value)
    if err != nil {
        panic(err)
    }
}
// SetResult sets the result property value. The result property
func (m *CustomSecurityAttributeAudit) SetResult(value *OperationResult)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
// SetResultReason sets the resultReason property value. The resultReason property
func (m *CustomSecurityAttributeAudit) SetResultReason(value *string)() {
    err := m.GetBackingStore().Set("resultReason", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetResources sets the targetResources property value. The targetResources property
func (m *CustomSecurityAttributeAudit) SetTargetResources(value []TargetResourceable)() {
    err := m.GetBackingStore().Set("targetResources", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAgent sets the userAgent property value. The userAgent property
func (m *CustomSecurityAttributeAudit) SetUserAgent(value *string)() {
    err := m.GetBackingStore().Set("userAgent", value)
    if err != nil {
        panic(err)
    }
}
// CustomSecurityAttributeAuditable 
type CustomSecurityAttributeAuditable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityDisplayName()(*string)
    GetAdditionalDetails()([]KeyValueable)
    GetCategory()(*string)
    GetCorrelationId()(*string)
    GetInitiatedBy()(AuditActivityInitiatorable)
    GetLoggedByService()(*string)
    GetOperationType()(*string)
    GetResult()(*OperationResult)
    GetResultReason()(*string)
    GetTargetResources()([]TargetResourceable)
    GetUserAgent()(*string)
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityDisplayName(value *string)()
    SetAdditionalDetails(value []KeyValueable)()
    SetCategory(value *string)()
    SetCorrelationId(value *string)()
    SetInitiatedBy(value AuditActivityInitiatorable)()
    SetLoggedByService(value *string)()
    SetOperationType(value *string)()
    SetResult(value *OperationResult)()
    SetResultReason(value *string)()
    SetTargetResources(value []TargetResourceable)()
    SetUserAgent(value *string)()
}
