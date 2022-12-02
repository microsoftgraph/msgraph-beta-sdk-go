package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse provides operations to call the getOffice365GroupsActivityFileCounts method.
type ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Office365GroupsActivityFileCountsable
}
// NewReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse instantiates a new ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse and sets the default values.
func NewReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse()(*ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse) {
    m := &ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReportsGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOffice365GroupsActivityFileCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Office365GroupsActivityFileCountsable, len(val))
            for i, v := range val {
                res[i] = v.(Office365GroupsActivityFileCountsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse) GetValue()([]Office365GroupsActivityFileCountsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReportsGetOffice365GroupsActivityFileCountsWithPeriodResponse) SetValue(value []Office365GroupsActivityFileCountsable)() {
    m.value = value
}
