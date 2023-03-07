package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ApplicableContentDeviceMatchCollectionResponse 
type ApplicableContentDeviceMatchCollectionResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewApplicableContentDeviceMatchCollectionResponse instantiates a new ApplicableContentDeviceMatchCollectionResponse and sets the default values.
func NewApplicableContentDeviceMatchCollectionResponse()(*ApplicableContentDeviceMatchCollectionResponse) {
    m := &ApplicableContentDeviceMatchCollectionResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateApplicableContentDeviceMatchCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicableContentDeviceMatchCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicableContentDeviceMatchCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicableContentDeviceMatchCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicableContentDeviceMatchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicableContentDeviceMatchable, len(val))
            for i, v := range val {
                res[i] = v.(ApplicableContentDeviceMatchable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ApplicableContentDeviceMatchCollectionResponse) GetValue()([]ApplicableContentDeviceMatchable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApplicableContentDeviceMatchable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApplicableContentDeviceMatchCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApplicableContentDeviceMatchCollectionResponse) SetValue(value []ApplicableContentDeviceMatchable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ApplicableContentDeviceMatchCollectionResponseable 
type ApplicableContentDeviceMatchCollectionResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ApplicableContentDeviceMatchable)
    SetValue(value []ApplicableContentDeviceMatchable)()
}
