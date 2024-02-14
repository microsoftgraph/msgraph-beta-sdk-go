package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserInsightsRoot struct {
    Entity
}
// NewUserInsightsRoot instantiates a new UserInsightsRoot and sets the default values.
func NewUserInsightsRoot()(*UserInsightsRoot) {
    m := &UserInsightsRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserInsightsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserInsightsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserInsightsRoot(), nil
}
// GetDaily gets the daily property value. Summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a DailyUserInsightMetricsRootable when successful
func (m *UserInsightsRoot) GetDaily()(DailyUserInsightMetricsRootable) {
    val, err := m.GetBackingStore().Get("daily")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DailyUserInsightMetricsRootable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserInsightsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["daily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDailyUserInsightMetricsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaily(val.(DailyUserInsightMetricsRootable))
        }
        return nil
    }
    res["monthly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMonthlyUserInsightMetricsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthly(val.(MonthlyUserInsightMetricsRootable))
        }
        return nil
    }
    return res
}
// GetMonthly gets the monthly property value. Summaries of monthly user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a MonthlyUserInsightMetricsRootable when successful
func (m *UserInsightsRoot) GetMonthly()(MonthlyUserInsightMetricsRootable) {
    val, err := m.GetBackingStore().Get("monthly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MonthlyUserInsightMetricsRootable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserInsightsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("daily", m.GetDaily())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("monthly", m.GetMonthly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDaily sets the daily property value. Summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
func (m *UserInsightsRoot) SetDaily(value DailyUserInsightMetricsRootable)() {
    err := m.GetBackingStore().Set("daily", value)
    if err != nil {
        panic(err)
    }
}
// SetMonthly sets the monthly property value. Summaries of monthly user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
func (m *UserInsightsRoot) SetMonthly(value MonthlyUserInsightMetricsRootable)() {
    err := m.GetBackingStore().Set("monthly", value)
    if err != nil {
        panic(err)
    }
}
type UserInsightsRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDaily()(DailyUserInsightMetricsRootable)
    GetMonthly()(MonthlyUserInsightMetricsRootable)
    SetDaily(value DailyUserInsightMetricsRootable)()
    SetMonthly(value MonthlyUserInsightMetricsRootable)()
}
