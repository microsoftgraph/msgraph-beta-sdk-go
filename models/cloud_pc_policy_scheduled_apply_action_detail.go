package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcPolicyScheduledApplyActionDetail struct {
    Entity
}
// NewCloudPcPolicyScheduledApplyActionDetail instantiates a new CloudPcPolicyScheduledApplyActionDetail and sets the default values.
func NewCloudPcPolicyScheduledApplyActionDetail()(*CloudPcPolicyScheduledApplyActionDetail) {
    m := &CloudPcPolicyScheduledApplyActionDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcPolicyScheduledApplyActionDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcPolicyScheduledApplyActionDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcPolicyScheduledApplyActionDetail(), nil
}
// GetCronScheduleExpression gets the cronScheduleExpression property value. An expression that specifies the cron schedule. (For example, '0 0 0 20  ' means schedules a job to run at midnight on the 20th of every month) Administrators can set a cron expression to define the scheduling rules for automatic regular application. When auto-provision is disabled, cronScheduleExpression is set to null, stopping the automatic task scheduling. Read-Only.
// returns a *string when successful
func (m *CloudPcPolicyScheduledApplyActionDetail) GetCronScheduleExpression()(*string) {
    val, err := m.GetBackingStore().Get("cronScheduleExpression")
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
func (m *CloudPcPolicyScheduledApplyActionDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cronScheduleExpression"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCronScheduleExpression(val)
        }
        return nil
    }
    res["reservePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReservePercentage(val)
        }
        return nil
    }
    return res
}
// GetReservePercentage gets the reservePercentage property value. The percentage of Cloud PCs to keep available. Administrators can set this property to a value from 0 to 99. Cloud PCs are reprovisioned only when there are no active and connected Cloud PC users. Frontline shared only.
// returns a *int32 when successful
func (m *CloudPcPolicyScheduledApplyActionDetail) GetReservePercentage()(*int32) {
    val, err := m.GetBackingStore().Get("reservePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcPolicyScheduledApplyActionDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cronScheduleExpression", m.GetCronScheduleExpression())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("reservePercentage", m.GetReservePercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCronScheduleExpression sets the cronScheduleExpression property value. An expression that specifies the cron schedule. (For example, '0 0 0 20  ' means schedules a job to run at midnight on the 20th of every month) Administrators can set a cron expression to define the scheduling rules for automatic regular application. When auto-provision is disabled, cronScheduleExpression is set to null, stopping the automatic task scheduling. Read-Only.
func (m *CloudPcPolicyScheduledApplyActionDetail) SetCronScheduleExpression(value *string)() {
    err := m.GetBackingStore().Set("cronScheduleExpression", value)
    if err != nil {
        panic(err)
    }
}
// SetReservePercentage sets the reservePercentage property value. The percentage of Cloud PCs to keep available. Administrators can set this property to a value from 0 to 99. Cloud PCs are reprovisioned only when there are no active and connected Cloud PC users. Frontline shared only.
func (m *CloudPcPolicyScheduledApplyActionDetail) SetReservePercentage(value *int32)() {
    err := m.GetBackingStore().Set("reservePercentage", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcPolicyScheduledApplyActionDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCronScheduleExpression()(*string)
    GetReservePercentage()(*int32)
    SetCronScheduleExpression(value *string)()
    SetReservePercentage(value *int32)()
}
