package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalyDeviceCollectionResponse 
type UserExperienceAnalyticsAnomalyDeviceCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewUserExperienceAnalyticsAnomalyDeviceCollectionResponse instantiates a new UserExperienceAnalyticsAnomalyDeviceCollectionResponse and sets the default values.
func NewUserExperienceAnalyticsAnomalyDeviceCollectionResponse()(*UserExperienceAnalyticsAnomalyDeviceCollectionResponse) {
    m := &UserExperienceAnalyticsAnomalyDeviceCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyDeviceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalyDeviceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalyDeviceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomalyDeviceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAnomalyDeviceable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UserExperienceAnalyticsAnomalyDeviceCollectionResponse) GetValue()([]UserExperienceAnalyticsAnomalyDeviceable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAnomalyDeviceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalyDeviceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UserExperienceAnalyticsAnomalyDeviceCollectionResponse) SetValue(value []UserExperienceAnalyticsAnomalyDeviceable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsAnomalyDeviceCollectionResponseable 
type UserExperienceAnalyticsAnomalyDeviceCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]UserExperienceAnalyticsAnomalyDeviceable)
    SetValue(value []UserExperienceAnalyticsAnomalyDeviceable)()
}
