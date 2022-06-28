package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsRegressionSummaryable 
type UserExperienceAnalyticsRegressionSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManufacturerRegression()([]UserExperienceAnalyticsMetricable)
    GetModelRegression()([]UserExperienceAnalyticsMetricable)
    GetOperatingSystemRegression()([]UserExperienceAnalyticsMetricable)
    SetManufacturerRegression(value []UserExperienceAnalyticsMetricable)()
    SetModelRegression(value []UserExperienceAnalyticsMetricable)()
    SetOperatingSystemRegression(value []UserExperienceAnalyticsMetricable)()
}
