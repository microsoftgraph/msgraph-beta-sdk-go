package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MonthlyInactiveUsersMetricCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewMonthlyInactiveUsersMetricCollectionResponse instantiates a new MonthlyInactiveUsersMetricCollectionResponse and sets the default values.
func NewMonthlyInactiveUsersMetricCollectionResponse()(*MonthlyInactiveUsersMetricCollectionResponse) {
    m := &MonthlyInactiveUsersMetricCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMonthlyInactiveUsersMetricCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMonthlyInactiveUsersMetricCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonthlyInactiveUsersMetricCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MonthlyInactiveUsersMetricCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMonthlyInactiveUsersMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MonthlyInactiveUsersMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MonthlyInactiveUsersMetricable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []MonthlyInactiveUsersMetricable when successful
func (m *MonthlyInactiveUsersMetricCollectionResponse) GetValue()([]MonthlyInactiveUsersMetricable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MonthlyInactiveUsersMetricable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MonthlyInactiveUsersMetricCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MonthlyInactiveUsersMetricCollectionResponse) SetValue(value []MonthlyInactiveUsersMetricable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type MonthlyInactiveUsersMetricCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]MonthlyInactiveUsersMetricable)
    SetValue(value []MonthlyInactiveUsersMetricable)()
}
