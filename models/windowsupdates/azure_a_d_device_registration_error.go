package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADDeviceRegistrationError 
type AzureADDeviceRegistrationError struct {
    UpdatableAssetError
    // The OdataType property
    OdataType *string
}
// NewAzureADDeviceRegistrationError instantiates a new azureADDeviceRegistrationError and sets the default values.
func NewAzureADDeviceRegistrationError()(*AzureADDeviceRegistrationError) {
    m := &AzureADDeviceRegistrationError{
        UpdatableAssetError: *NewUpdatableAssetError(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.azureADDeviceRegistrationError"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureADDeviceRegistrationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADDeviceRegistrationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADDeviceRegistrationError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADDeviceRegistrationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAssetError.GetFieldDeserializers()
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADDeviceRegistrationErrorReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*AzureADDeviceRegistrationErrorReason))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *AzureADDeviceRegistrationError) GetReason()(*AzureADDeviceRegistrationErrorReason) {
    val, err := m.GetBackingStore().Get("reason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AzureADDeviceRegistrationErrorReason)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AzureADDeviceRegistrationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAssetError.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err = writer.WriteStringValue("reason", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetReason sets the reason property value. The reason property
func (m *AzureADDeviceRegistrationError) SetReason(value *AzureADDeviceRegistrationErrorReason)() {
    err := m.GetBackingStore().Set("reason", value)
    if err != nil {
        panic(err)
    }
}
// AzureADDeviceRegistrationErrorable 
type AzureADDeviceRegistrationErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UpdatableAssetErrorable
    GetReason()(*AzureADDeviceRegistrationErrorReason)
    SetReason(value *AzureADDeviceRegistrationErrorReason)()
}
