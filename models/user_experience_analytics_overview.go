package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsOverview 
type UserExperienceAnalyticsOverview struct {
    Entity
    // The user experience analytics insights.
    insights []UserExperienceAnalyticsInsightable
}
// NewUserExperienceAnalyticsOverview instantiates a new userExperienceAnalyticsOverview and sets the default values.
func NewUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    m := &UserExperienceAnalyticsOverview{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsOverview";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsOverview(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightFromDiscriminatorValue , m.SetInsights)
    return res
}
// GetInsights gets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) GetInsights()([]UserExperienceAnalyticsInsightable) {
    return m.insights
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetInsights sets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) SetInsights(value []UserExperienceAnalyticsInsightable)() {
    m.insights = value
}
