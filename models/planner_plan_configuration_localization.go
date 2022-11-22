package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfigurationLocalization provides operations to manage the collection of activityStatistics entities.
type PlannerPlanConfigurationLocalization struct {
    Entity
    // The buckets property
    buckets []PlannerPlanConfigurationBucketLocalizationable
    // The languageTag property
    languageTag *string
    // The planTitle property
    planTitle *string
}
// NewPlannerPlanConfigurationLocalization instantiates a new plannerPlanConfigurationLocalization and sets the default values.
func NewPlannerPlanConfigurationLocalization()(*PlannerPlanConfigurationLocalization) {
    m := &PlannerPlanConfigurationLocalization{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanConfigurationLocalization(), nil
}
// GetBuckets gets the buckets property value. The buckets property
func (m *PlannerPlanConfigurationLocalization) GetBuckets()([]PlannerPlanConfigurationBucketLocalizationable) {
    return m.buckets
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanConfigurationLocalization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerPlanConfigurationBucketLocalizationFromDiscriminatorValue , m.SetBuckets)
    res["languageTag"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLanguageTag)
    res["planTitle"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPlanTitle)
    return res
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *PlannerPlanConfigurationLocalization) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetPlanTitle gets the planTitle property value. The planTitle property
func (m *PlannerPlanConfigurationLocalization) GetPlanTitle()(*string) {
    return m.planTitle
}
// Serialize serializes information the current object
func (m *PlannerPlanConfigurationLocalization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBuckets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetBuckets())
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planTitle", m.GetPlanTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. The buckets property
func (m *PlannerPlanConfigurationLocalization) SetBuckets(value []PlannerPlanConfigurationBucketLocalizationable)() {
    m.buckets = value
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *PlannerPlanConfigurationLocalization) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetPlanTitle sets the planTitle property value. The planTitle property
func (m *PlannerPlanConfigurationLocalization) SetPlanTitle(value *string)() {
    m.planTitle = value
}
