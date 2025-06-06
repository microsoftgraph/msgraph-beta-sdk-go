// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DailyUserInsightMetricsRoot struct {
    Entity
}
// NewDailyUserInsightMetricsRoot instantiates a new DailyUserInsightMetricsRoot and sets the default values.
func NewDailyUserInsightMetricsRoot()(*DailyUserInsightMetricsRoot) {
    m := &DailyUserInsightMetricsRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDailyUserInsightMetricsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDailyUserInsightMetricsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDailyUserInsightMetricsRoot(), nil
}
// GetActiveUsers gets the activeUsers property value. Insights for active users on apps registered in the tenant for a specified period.
// returns a []ActiveUsersMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetActiveUsers()([]ActiveUsersMetricable) {
    val, err := m.GetBackingStore().Get("activeUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActiveUsersMetricable)
    }
    return nil
}
// GetAuthentications gets the authentications property value. Insights for authentications on apps registered in the tenant for a specified period.
// returns a []AuthenticationsMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetAuthentications()([]AuthenticationsMetricable) {
    val, err := m.GetBackingStore().Get("authentications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationsMetricable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DailyUserInsightMetricsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActiveUsersMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActiveUsersMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActiveUsersMetricable)
                }
            }
            m.SetActiveUsers(res)
        }
        return nil
    }
    res["authentications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationsMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationsMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationsMetricable)
                }
            }
            m.SetAuthentications(res)
        }
        return nil
    }
    res["inactiveUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDailyInactiveUsersMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DailyInactiveUsersMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DailyInactiveUsersMetricable)
                }
            }
            m.SetInactiveUsers(res)
        }
        return nil
    }
    res["inactiveUsersByApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDailyInactiveUsersByApplicationMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DailyInactiveUsersByApplicationMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DailyInactiveUsersByApplicationMetricable)
                }
            }
            m.SetInactiveUsersByApplication(res)
        }
        return nil
    }
    res["mfaCompletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMfaCompletionMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MfaCompletionMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MfaCompletionMetricable)
                }
            }
            m.SetMfaCompletions(res)
        }
        return nil
    }
    res["mfaTelecomFraud"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMfaTelecomFraudMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MfaTelecomFraudMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MfaTelecomFraudMetricable)
                }
            }
            m.SetMfaTelecomFraud(res)
        }
        return nil
    }
    res["signUps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSignUpMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSignUpMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserSignUpMetricable)
                }
            }
            m.SetSignUps(res)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInsightSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InsightSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(InsightSummaryable)
                }
            }
            m.SetSummary(res)
        }
        return nil
    }
    res["userCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserCountMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserCountMetricable)
                }
            }
            m.SetUserCount(res)
        }
        return nil
    }
    return res
}
// GetInactiveUsers gets the inactiveUsers property value. The inactiveUsers property
// returns a []DailyInactiveUsersMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetInactiveUsers()([]DailyInactiveUsersMetricable) {
    val, err := m.GetBackingStore().Get("inactiveUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DailyInactiveUsersMetricable)
    }
    return nil
}
// GetInactiveUsersByApplication gets the inactiveUsersByApplication property value. The inactiveUsersByApplication property
// returns a []DailyInactiveUsersByApplicationMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetInactiveUsersByApplication()([]DailyInactiveUsersByApplicationMetricable) {
    val, err := m.GetBackingStore().Get("inactiveUsersByApplication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DailyInactiveUsersByApplicationMetricable)
    }
    return nil
}
// GetMfaCompletions gets the mfaCompletions property value. Insights for MFA usage on apps registered in the tenant for a specified period.
// returns a []MfaCompletionMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetMfaCompletions()([]MfaCompletionMetricable) {
    val, err := m.GetBackingStore().Get("mfaCompletions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MfaCompletionMetricable)
    }
    return nil
}
// GetMfaTelecomFraud gets the mfaTelecomFraud property value. The mfaTelecomFraud property
// returns a []MfaTelecomFraudMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetMfaTelecomFraud()([]MfaTelecomFraudMetricable) {
    val, err := m.GetBackingStore().Get("mfaTelecomFraud")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MfaTelecomFraudMetricable)
    }
    return nil
}
// GetSignUps gets the signUps property value. Total sign-ups on apps registered in the tenant for a specified period.
// returns a []UserSignUpMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetSignUps()([]UserSignUpMetricable) {
    val, err := m.GetBackingStore().Get("signUps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserSignUpMetricable)
    }
    return nil
}
// GetSummary gets the summary property value. Summary of all usage insights on apps registered in the tenant for a specified period.
// returns a []InsightSummaryable when successful
func (m *DailyUserInsightMetricsRoot) GetSummary()([]InsightSummaryable) {
    val, err := m.GetBackingStore().Get("summary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]InsightSummaryable)
    }
    return nil
}
// GetUserCount gets the userCount property value. Insights for total users on apps registered in the tenant for a specified period.
// returns a []UserCountMetricable when successful
func (m *DailyUserInsightMetricsRoot) GetUserCount()([]UserCountMetricable) {
    val, err := m.GetBackingStore().Get("userCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserCountMetricable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DailyUserInsightMetricsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActiveUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActiveUsers()))
        for i, v := range m.GetActiveUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activeUsers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthentications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthentications()))
        for i, v := range m.GetAuthentications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authentications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInactiveUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInactiveUsers()))
        for i, v := range m.GetInactiveUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("inactiveUsers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInactiveUsersByApplication() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInactiveUsersByApplication()))
        for i, v := range m.GetInactiveUsersByApplication() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("inactiveUsersByApplication", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMfaCompletions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMfaCompletions()))
        for i, v := range m.GetMfaCompletions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mfaCompletions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMfaTelecomFraud() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMfaTelecomFraud()))
        for i, v := range m.GetMfaTelecomFraud() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mfaTelecomFraud", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignUps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignUps()))
        for i, v := range m.GetSignUps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("signUps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSummary() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSummary()))
        for i, v := range m.GetSummary() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("summary", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserCount() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserCount()))
        for i, v := range m.GetUserCount() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userCount", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveUsers sets the activeUsers property value. Insights for active users on apps registered in the tenant for a specified period.
func (m *DailyUserInsightMetricsRoot) SetActiveUsers(value []ActiveUsersMetricable)() {
    err := m.GetBackingStore().Set("activeUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthentications sets the authentications property value. Insights for authentications on apps registered in the tenant for a specified period.
func (m *DailyUserInsightMetricsRoot) SetAuthentications(value []AuthenticationsMetricable)() {
    err := m.GetBackingStore().Set("authentications", value)
    if err != nil {
        panic(err)
    }
}
// SetInactiveUsers sets the inactiveUsers property value. The inactiveUsers property
func (m *DailyUserInsightMetricsRoot) SetInactiveUsers(value []DailyInactiveUsersMetricable)() {
    err := m.GetBackingStore().Set("inactiveUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetInactiveUsersByApplication sets the inactiveUsersByApplication property value. The inactiveUsersByApplication property
func (m *DailyUserInsightMetricsRoot) SetInactiveUsersByApplication(value []DailyInactiveUsersByApplicationMetricable)() {
    err := m.GetBackingStore().Set("inactiveUsersByApplication", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaCompletions sets the mfaCompletions property value. Insights for MFA usage on apps registered in the tenant for a specified period.
func (m *DailyUserInsightMetricsRoot) SetMfaCompletions(value []MfaCompletionMetricable)() {
    err := m.GetBackingStore().Set("mfaCompletions", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaTelecomFraud sets the mfaTelecomFraud property value. The mfaTelecomFraud property
func (m *DailyUserInsightMetricsRoot) SetMfaTelecomFraud(value []MfaTelecomFraudMetricable)() {
    err := m.GetBackingStore().Set("mfaTelecomFraud", value)
    if err != nil {
        panic(err)
    }
}
// SetSignUps sets the signUps property value. Total sign-ups on apps registered in the tenant for a specified period.
func (m *DailyUserInsightMetricsRoot) SetSignUps(value []UserSignUpMetricable)() {
    err := m.GetBackingStore().Set("signUps", value)
    if err != nil {
        panic(err)
    }
}
// SetSummary sets the summary property value. Summary of all usage insights on apps registered in the tenant for a specified period.
func (m *DailyUserInsightMetricsRoot) SetSummary(value []InsightSummaryable)() {
    err := m.GetBackingStore().Set("summary", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCount sets the userCount property value. Insights for total users on apps registered in the tenant for a specified period.
func (m *DailyUserInsightMetricsRoot) SetUserCount(value []UserCountMetricable)() {
    err := m.GetBackingStore().Set("userCount", value)
    if err != nil {
        panic(err)
    }
}
type DailyUserInsightMetricsRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveUsers()([]ActiveUsersMetricable)
    GetAuthentications()([]AuthenticationsMetricable)
    GetInactiveUsers()([]DailyInactiveUsersMetricable)
    GetInactiveUsersByApplication()([]DailyInactiveUsersByApplicationMetricable)
    GetMfaCompletions()([]MfaCompletionMetricable)
    GetMfaTelecomFraud()([]MfaTelecomFraudMetricable)
    GetSignUps()([]UserSignUpMetricable)
    GetSummary()([]InsightSummaryable)
    GetUserCount()([]UserCountMetricable)
    SetActiveUsers(value []ActiveUsersMetricable)()
    SetAuthentications(value []AuthenticationsMetricable)()
    SetInactiveUsers(value []DailyInactiveUsersMetricable)()
    SetInactiveUsersByApplication(value []DailyInactiveUsersByApplicationMetricable)()
    SetMfaCompletions(value []MfaCompletionMetricable)()
    SetMfaTelecomFraud(value []MfaTelecomFraudMetricable)()
    SetSignUps(value []UserSignUpMetricable)()
    SetSummary(value []InsightSummaryable)()
    SetUserCount(value []UserCountMetricable)()
}
