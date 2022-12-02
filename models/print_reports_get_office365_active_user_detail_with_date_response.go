package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintReportsGetOffice365ActiveUserDetailWithDateResponse provides operations to call the getOffice365ActiveUserDetail method.
type PrintReportsGetOffice365ActiveUserDetailWithDateResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Office365ActiveUserDetailable
}
// NewPrintReportsGetOffice365ActiveUserDetailWithDateResponse instantiates a new PrintReportsGetOffice365ActiveUserDetailWithDateResponse and sets the default values.
func NewPrintReportsGetOffice365ActiveUserDetailWithDateResponse()(*PrintReportsGetOffice365ActiveUserDetailWithDateResponse) {
    m := &PrintReportsGetOffice365ActiveUserDetailWithDateResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrintReportsGetOffice365ActiveUserDetailWithDateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintReportsGetOffice365ActiveUserDetailWithDateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintReportsGetOffice365ActiveUserDetailWithDateResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintReportsGetOffice365ActiveUserDetailWithDateResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOffice365ActiveUserDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Office365ActiveUserDetailable, len(val))
            for i, v := range val {
                res[i] = v.(Office365ActiveUserDetailable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PrintReportsGetOffice365ActiveUserDetailWithDateResponse) GetValue()([]Office365ActiveUserDetailable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrintReportsGetOffice365ActiveUserDetailWithDateResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrintReportsGetOffice365ActiveUserDetailWithDateResponse) SetValue(value []Office365ActiveUserDetailable)() {
    m.value = value
}
