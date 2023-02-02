package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse 
type MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
    // The value property
    value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Office365GroupsActivityFileCountsable
}
// NewMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse instantiates a new MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse and sets the default values.
func NewMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse()(*MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse) {
    m := &MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOffice365GroupsActivityFileCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Office365GroupsActivityFileCountsable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Office365GroupsActivityFileCountsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Office365GroupsActivityFileCountsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Office365GroupsActivityFileCountsable)() {
    m.value = value
}
