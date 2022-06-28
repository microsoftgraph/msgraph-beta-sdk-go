package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcAuditEventable 
type CloudPcAuditEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityOperationType()(*CloudPcAuditActivityOperationType)
    GetActivityResult()(*CloudPcAuditActivityResult)
    GetActivityType()(*string)
    GetActor()(CloudPcAuditActorable)
    GetCategory()(*CloudPcAuditCategory)
    GetComponentName()(*string)
    GetCorrelationId()(*string)
    GetDisplayName()(*string)
    GetResources()([]CloudPcAuditResourceable)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityOperationType(value *CloudPcAuditActivityOperationType)()
    SetActivityResult(value *CloudPcAuditActivityResult)()
    SetActivityType(value *string)()
    SetActor(value CloudPcAuditActorable)()
    SetCategory(value *CloudPcAuditCategory)()
    SetComponentName(value *string)()
    SetCorrelationId(value *string)()
    SetDisplayName(value *string)()
    SetResources(value []CloudPcAuditResourceable)()
}
