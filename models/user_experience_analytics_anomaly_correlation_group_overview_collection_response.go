package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse 
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse instantiates a new UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse and sets the default values.
func NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse()(*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse) {
    m := &UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse) GetValue()([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponse) SetValue(value []UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseable 
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)
    SetValue(value []UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)()
}
