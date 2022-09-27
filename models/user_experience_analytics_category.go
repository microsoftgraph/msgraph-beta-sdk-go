package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsCategory 
type UserExperienceAnalyticsCategory struct {
    Entity
    // The insights for the user experience analytics category.
    insights []UserExperienceAnalyticsInsightable
    // The metric values for the user experience analytics category.
    metricValues []UserExperienceAnalyticsMetricable
}
// NewUserExperienceAnalyticsCategory instantiates a new userExperienceAnalyticsCategory and sets the default values.
func NewUserExperienceAnalyticsCategory()(*UserExperienceAnalyticsCategory) {
    m := &UserExperienceAnalyticsCategory{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsCategory";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightFromDiscriminatorValue , m.SetInsights)
    res["metricValues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricFromDiscriminatorValue , m.SetMetricValues)
    return res
}
// GetInsights gets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetInsights()([]UserExperienceAnalyticsInsightable) {
    return m.insights
}
// GetMetricValues gets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetMetricValues()([]UserExperienceAnalyticsMetricable) {
    return m.metricValues
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInsights() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInsights())
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetricValues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMetricValues())
        err = writer.WriteCollectionOfObjectValues("metricValues", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsights sets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) SetInsights(value []UserExperienceAnalyticsInsightable)() {
    m.insights = value
}
// SetMetricValues sets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) SetMetricValues(value []UserExperienceAnalyticsMetricable)() {
    m.metricValues = value
}
