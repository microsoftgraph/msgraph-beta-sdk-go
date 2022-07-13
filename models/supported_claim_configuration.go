package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SupportedClaimConfiguration 
type SupportedClaimConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The nameIdPolicyFormat property
    nameIdPolicyFormat *string
}
// NewSupportedClaimConfiguration instantiates a new supportedClaimConfiguration and sets the default values.
func NewSupportedClaimConfiguration()(*SupportedClaimConfiguration) {
    m := &SupportedClaimConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSupportedClaimConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSupportedClaimConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSupportedClaimConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SupportedClaimConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SupportedClaimConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nameIdPolicyFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameIdPolicyFormat(val)
        }
        return nil
    }
    return res
}
// GetNameIdPolicyFormat gets the nameIdPolicyFormat property value. The nameIdPolicyFormat property
func (m *SupportedClaimConfiguration) GetNameIdPolicyFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nameIdPolicyFormat
    }
}
// Serialize serializes information the current object
func (m *SupportedClaimConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("nameIdPolicyFormat", m.GetNameIdPolicyFormat())
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
func (m *SupportedClaimConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNameIdPolicyFormat sets the nameIdPolicyFormat property value. The nameIdPolicyFormat property
func (m *SupportedClaimConfiguration) SetNameIdPolicyFormat(value *string)() {
    if m != nil {
        m.nameIdPolicyFormat = value
    }
}
