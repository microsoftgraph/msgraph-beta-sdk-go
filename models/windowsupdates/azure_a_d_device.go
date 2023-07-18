package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADDevice 
type AzureADDevice struct {
    UpdatableAsset
}
// NewAzureADDevice instantiates a new azureADDevice and sets the default values.
func NewAzureADDevice()(*AzureADDevice) {
    m := &AzureADDevice{
        UpdatableAsset: *NewUpdatableAsset(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.azureADDevice"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureADDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADDevice(), nil
}
// GetEnrollments gets the enrollments property value. Specifies areas of the service in which the device is enrolled. Read-only. Returned by default.
func (m *AzureADDevice) GetEnrollments()([]UpdatableAssetEnrollmentable) {
    val, err := m.GetBackingStore().Get("enrollments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UpdatableAssetEnrollmentable)
    }
    return nil
}
// GetErrors gets the errors property value. Specifies any errors that prevent the device from being enrolled in update management or receving deployed content. Read-only. Returned by default.
func (m *AzureADDevice) GetErrors()([]UpdatableAssetErrorable) {
    val, err := m.GetBackingStore().Get("errors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UpdatableAssetErrorable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAsset.GetFieldDeserializers()
    res["enrollments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetEnrollmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetEnrollmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UpdatableAssetEnrollmentable)
                }
            }
            m.SetEnrollments(res)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetErrorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UpdatableAssetErrorable)
                }
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AzureADDevice) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AzureADDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAsset.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEnrollments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollments()))
        for i, v := range m.GetEnrollments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("enrollments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("errors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnrollments sets the enrollments property value. Specifies areas of the service in which the device is enrolled. Read-only. Returned by default.
func (m *AzureADDevice) SetEnrollments(value []UpdatableAssetEnrollmentable)() {
    err := m.GetBackingStore().Set("enrollments", value)
    if err != nil {
        panic(err)
    }
}
// SetErrors sets the errors property value. Specifies any errors that prevent the device from being enrolled in update management or receving deployed content. Read-only. Returned by default.
func (m *AzureADDevice) SetErrors(value []UpdatableAssetErrorable)() {
    err := m.GetBackingStore().Set("errors", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AzureADDevice) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AzureADDeviceable 
type AzureADDeviceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UpdatableAssetable
    GetEnrollments()([]UpdatableAssetEnrollmentable)
    GetErrors()([]UpdatableAssetErrorable)
    GetOdataType()(*string)
    SetEnrollments(value []UpdatableAssetEnrollmentable)()
    SetErrors(value []UpdatableAssetErrorable)()
    SetOdataType(value *string)()
}
