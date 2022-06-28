package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADDevice 
type AzureADDevice struct {
    UpdatableAsset
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies areas of the service in which the device is enrolled. Read-only. Returned by default.
    enrollments []UpdatableAssetEnrollmentable
    // Specifies any errors that prevent the device from being enrolled in update management or receving deployed content. Read-only. Returned by default.
    errors []UpdatableAssetErrorable
}
// NewAzureADDevice instantiates a new AzureADDevice and sets the default values.
func NewAzureADDevice()(*AzureADDevice) {
    m := &AzureADDevice{
        UpdatableAsset: *NewUpdatableAsset(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAzureADDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADDevice(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADDevice) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnrollments gets the enrollments property value. Specifies areas of the service in which the device is enrolled. Read-only. Returned by default.
func (m *AzureADDevice) GetEnrollments()([]UpdatableAssetEnrollmentable) {
    if m == nil {
        return nil
    } else {
        return m.enrollments
    }
}
// GetErrors gets the errors property value. Specifies any errors that prevent the device from being enrolled in update management or receving deployed content. Read-only. Returned by default.
func (m *AzureADDevice) GetErrors()([]UpdatableAssetErrorable) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
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
                res[i] = v.(UpdatableAssetEnrollmentable)
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
                res[i] = v.(UpdatableAssetErrorable)
            }
            m.SetErrors(res)
        }
        return nil
    }
    return res
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enrollments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("errors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADDevice) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnrollments sets the enrollments property value. Specifies areas of the service in which the device is enrolled. Read-only. Returned by default.
func (m *AzureADDevice) SetEnrollments(value []UpdatableAssetEnrollmentable)() {
    if m != nil {
        m.enrollments = value
    }
}
// SetErrors sets the errors property value. Specifies any errors that prevent the device from being enrolled in update management or receving deployed content. Read-only. Returned by default.
func (m *AzureADDevice) SetErrors(value []UpdatableAssetErrorable)() {
    if m != nil {
        m.errors = value
    }
}
