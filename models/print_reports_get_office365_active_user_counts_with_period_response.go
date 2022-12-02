package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse provides operations to call the getOffice365ActiveUserCounts method.
type PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Office365ActiveUserCountsable
}
// NewPrintReportsGetOffice365ActiveUserCountsWithPeriodResponse instantiates a new PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse and sets the default values.
func NewPrintReportsGetOffice365ActiveUserCountsWithPeriodResponse()(*PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse) {
    m := &PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrintReportsGetOffice365ActiveUserCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintReportsGetOffice365ActiveUserCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintReportsGetOffice365ActiveUserCountsWithPeriodResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOffice365ActiveUserCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Office365ActiveUserCountsable, len(val))
            for i, v := range val {
                res[i] = v.(Office365ActiveUserCountsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse) GetValue()([]Office365ActiveUserCountsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrintReportsGetOffice365ActiveUserCountsWithPeriodResponse) SetValue(value []Office365ActiveUserCountsable)() {
    m.value = value
}
