package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsGetOffice365ServicesUserCountsWithPeriodResponse provides operations to call the getOffice365ServicesUserCounts method.
type ReportsGetOffice365ServicesUserCountsWithPeriodResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Office365ServicesUserCountsable
}
// NewReportsGetOffice365ServicesUserCountsWithPeriodResponse instantiates a new ReportsGetOffice365ServicesUserCountsWithPeriodResponse and sets the default values.
func NewReportsGetOffice365ServicesUserCountsWithPeriodResponse()(*ReportsGetOffice365ServicesUserCountsWithPeriodResponse) {
    m := &ReportsGetOffice365ServicesUserCountsWithPeriodResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReportsGetOffice365ServicesUserCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsGetOffice365ServicesUserCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsGetOffice365ServicesUserCountsWithPeriodResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsGetOffice365ServicesUserCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOffice365ServicesUserCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Office365ServicesUserCountsable, len(val))
            for i, v := range val {
                res[i] = v.(Office365ServicesUserCountsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReportsGetOffice365ServicesUserCountsWithPeriodResponse) GetValue()([]Office365ServicesUserCountsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ReportsGetOffice365ServicesUserCountsWithPeriodResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReportsGetOffice365ServicesUserCountsWithPeriodResponse) SetValue(value []Office365ServicesUserCountsable)() {
    m.value = value
}
