package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsGetOffice365GroupsActivityCountsWithPeriodResponse provides operations to call the getOffice365GroupsActivityCounts method.
type ReportsGetOffice365GroupsActivityCountsWithPeriodResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Office365GroupsActivityCountsable
}
// NewReportsGetOffice365GroupsActivityCountsWithPeriodResponse instantiates a new ReportsGetOffice365GroupsActivityCountsWithPeriodResponse and sets the default values.
func NewReportsGetOffice365GroupsActivityCountsWithPeriodResponse()(*ReportsGetOffice365GroupsActivityCountsWithPeriodResponse) {
    m := &ReportsGetOffice365GroupsActivityCountsWithPeriodResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReportsGetOffice365GroupsActivityCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsGetOffice365GroupsActivityCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsGetOffice365GroupsActivityCountsWithPeriodResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsGetOffice365GroupsActivityCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOffice365GroupsActivityCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Office365GroupsActivityCountsable, len(val))
            for i, v := range val {
                res[i] = v.(Office365GroupsActivityCountsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReportsGetOffice365GroupsActivityCountsWithPeriodResponse) GetValue()([]Office365GroupsActivityCountsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ReportsGetOffice365GroupsActivityCountsWithPeriodResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReportsGetOffice365GroupsActivityCountsWithPeriodResponse) SetValue(value []Office365GroupsActivityCountsable)() {
    m.value = value
}
