package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPFXCertificateCollectionResponse 
type UserPFXCertificateCollectionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []UserPFXCertificateable
}
// NewUserPFXCertificateCollectionResponse instantiates a new UserPFXCertificateCollectionResponse and sets the default values.
func NewUserPFXCertificateCollectionResponse()(*UserPFXCertificateCollectionResponse) {
    m := &UserPFXCertificateCollectionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserPFXCertificateCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserPFXCertificateCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPFXCertificateCollectionResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserPFXCertificateCollectionResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserPFXCertificateCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.nextLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataNextLink(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserPFXCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserPFXCertificateable, len(val))
            for i, v := range val {
                res[i] = v.(UserPFXCertificateable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *UserPFXCertificateCollectionResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *UserPFXCertificateCollectionResponse) GetValue()([]UserPFXCertificateable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UserPFXCertificateCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserPFXCertificateCollectionResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *UserPFXCertificateCollectionResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *UserPFXCertificateCollectionResponse) SetValue(value []UserPFXCertificateable)() {
    m.value = value
}
