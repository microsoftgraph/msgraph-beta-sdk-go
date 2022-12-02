package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse provides operations to call the getOffice365ServicesUserCounts method.
type PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Office365ServicesUserCountsable
}
// NewPrintReportsGetOffice365ServicesUserCountsWithPeriodResponse instantiates a new PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse and sets the default values.
func NewPrintReportsGetOffice365ServicesUserCountsWithPeriodResponse()(*PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse) {
    m := &PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrintReportsGetOffice365ServicesUserCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintReportsGetOffice365ServicesUserCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintReportsGetOffice365ServicesUserCountsWithPeriodResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse) GetValue()([]Office365ServicesUserCountsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrintReportsGetOffice365ServicesUserCountsWithPeriodResponse) SetValue(value []Office365ServicesUserCountsable)() {
    m.value = value
}
