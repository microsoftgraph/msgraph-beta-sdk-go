package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsRegressionSummaryable 
type UserExperienceAnalyticsRegressionSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetManufacturerRegression()([]UserExperienceAnalyticsMetricable)
    GetModelRegression()([]UserExperienceAnalyticsMetricable)
    GetOperatingSystemRegression()([]UserExperienceAnalyticsMetricable)
    SetManufacturerRegression(value []UserExperienceAnalyticsMetricable)()
    SetModelRegression(value []UserExperienceAnalyticsMetricable)()
    SetOperatingSystemRegression(value []UserExperienceAnalyticsMetricable)()
}
