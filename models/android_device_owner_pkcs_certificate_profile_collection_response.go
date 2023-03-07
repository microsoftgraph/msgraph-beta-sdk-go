package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse 
type AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewAndroidDeviceOwnerPkcsCertificateProfileCollectionResponse instantiates a new AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse and sets the default values.
func NewAndroidDeviceOwnerPkcsCertificateProfileCollectionResponse()(*AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse) {
    m := &AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAndroidDeviceOwnerPkcsCertificateProfileCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerPkcsCertificateProfileCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerPkcsCertificateProfileCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerPkcsCertificateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerPkcsCertificateProfileable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerPkcsCertificateProfileable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse) GetValue()([]AndroidDeviceOwnerPkcsCertificateProfileable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerPkcsCertificateProfileable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AndroidDeviceOwnerPkcsCertificateProfileCollectionResponse) SetValue(value []AndroidDeviceOwnerPkcsCertificateProfileable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerPkcsCertificateProfileCollectionResponseable 
type AndroidDeviceOwnerPkcsCertificateProfileCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]AndroidDeviceOwnerPkcsCertificateProfileable)
    SetValue(value []AndroidDeviceOwnerPkcsCertificateProfileable)()
}
