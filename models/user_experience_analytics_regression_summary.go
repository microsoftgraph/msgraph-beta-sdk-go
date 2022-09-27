package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsRegressionSummary 
type UserExperienceAnalyticsRegressionSummary struct {
    Entity
    // The metric values for the user experience analytics Manufacturer regression.
    manufacturerRegression []UserExperienceAnalyticsMetricable
    // The metric values for the user experience analytics model regression.
    modelRegression []UserExperienceAnalyticsMetricable
    // The metric values for the user experience analytics operating system regression.
    operatingSystemRegression []UserExperienceAnalyticsMetricable
}
// NewUserExperienceAnalyticsRegressionSummary instantiates a new userExperienceAnalyticsRegressionSummary and sets the default values.
func NewUserExperienceAnalyticsRegressionSummary()(*UserExperienceAnalyticsRegressionSummary) {
    m := &UserExperienceAnalyticsRegressionSummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsRegressionSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsRegressionSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsRegressionSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsRegressionSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsRegressionSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["manufacturerRegression"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricFromDiscriminatorValue , m.SetManufacturerRegression)
    res["modelRegression"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricFromDiscriminatorValue , m.SetModelRegression)
    res["operatingSystemRegression"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricFromDiscriminatorValue , m.SetOperatingSystemRegression)
    return res
}
// GetManufacturerRegression gets the manufacturerRegression property value. The metric values for the user experience analytics Manufacturer regression.
func (m *UserExperienceAnalyticsRegressionSummary) GetManufacturerRegression()([]UserExperienceAnalyticsMetricable) {
    return m.manufacturerRegression
}
// GetModelRegression gets the modelRegression property value. The metric values for the user experience analytics model regression.
func (m *UserExperienceAnalyticsRegressionSummary) GetModelRegression()([]UserExperienceAnalyticsMetricable) {
    return m.modelRegression
}
// GetOperatingSystemRegression gets the operatingSystemRegression property value. The metric values for the user experience analytics operating system regression.
func (m *UserExperienceAnalyticsRegressionSummary) GetOperatingSystemRegression()([]UserExperienceAnalyticsMetricable) {
    return m.operatingSystemRegression
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsRegressionSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetManufacturerRegression() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManufacturerRegression())
        err = writer.WriteCollectionOfObjectValues("manufacturerRegression", cast)
        if err != nil {
            return err
        }
    }
    if m.GetModelRegression() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetModelRegression())
        err = writer.WriteCollectionOfObjectValues("modelRegression", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperatingSystemRegression() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOperatingSystemRegression())
        err = writer.WriteCollectionOfObjectValues("operatingSystemRegression", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManufacturerRegression sets the manufacturerRegression property value. The metric values for the user experience analytics Manufacturer regression.
func (m *UserExperienceAnalyticsRegressionSummary) SetManufacturerRegression(value []UserExperienceAnalyticsMetricable)() {
    m.manufacturerRegression = value
}
// SetModelRegression sets the modelRegression property value. The metric values for the user experience analytics model regression.
func (m *UserExperienceAnalyticsRegressionSummary) SetModelRegression(value []UserExperienceAnalyticsMetricable)() {
    m.modelRegression = value
}
// SetOperatingSystemRegression sets the operatingSystemRegression property value. The metric values for the user experience analytics operating system regression.
func (m *UserExperienceAnalyticsRegressionSummary) SetOperatingSystemRegression(value []UserExperienceAnalyticsMetricable)() {
    m.operatingSystemRegression = value
}
