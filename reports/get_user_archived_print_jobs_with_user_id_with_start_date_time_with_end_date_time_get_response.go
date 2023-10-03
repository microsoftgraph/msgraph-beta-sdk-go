package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse 
type GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse instantiates a new GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse and sets the default values.
func NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse()(*GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse) {
    m := &GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateArchivedPrintJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseable 
type GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable)
    SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArchivedPrintJobable)()
}
