package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestSignatureVerification 
type RequestSignatureVerification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies whether this application accepts weak algorithms.  The possible values are: rsaSha1, unknownFutureValue.
    allowedWeakAlgorithms *WeakAlgorithms
    // Specifies whether signed authentication requests for this application should be required.
    isSignedRequestRequired *bool
    // The OdataType property
    odataType *string
}
// NewRequestSignatureVerification instantiates a new requestSignatureVerification and sets the default values.
func NewRequestSignatureVerification()(*RequestSignatureVerification) {
    m := &RequestSignatureVerification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.requestSignatureVerification";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRequestSignatureVerificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestSignatureVerificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestSignatureVerification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestSignatureVerification) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedWeakAlgorithms gets the allowedWeakAlgorithms property value. Specifies whether this application accepts weak algorithms.  The possible values are: rsaSha1, unknownFutureValue.
func (m *RequestSignatureVerification) GetAllowedWeakAlgorithms()(*WeakAlgorithms) {
    return m.allowedWeakAlgorithms
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestSignatureVerification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedWeakAlgorithms"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWeakAlgorithms , m.SetAllowedWeakAlgorithms)
    res["isSignedRequestRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSignedRequestRequired)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsSignedRequestRequired gets the isSignedRequestRequired property value. Specifies whether signed authentication requests for this application should be required.
func (m *RequestSignatureVerification) GetIsSignedRequestRequired()(*bool) {
    return m.isSignedRequestRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RequestSignatureVerification) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *RequestSignatureVerification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedWeakAlgorithms() != nil {
        cast := (*m.GetAllowedWeakAlgorithms()).String()
        err := writer.WriteStringValue("allowedWeakAlgorithms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSignedRequestRequired", m.GetIsSignedRequestRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *RequestSignatureVerification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedWeakAlgorithms sets the allowedWeakAlgorithms property value. Specifies whether this application accepts weak algorithms.  The possible values are: rsaSha1, unknownFutureValue.
func (m *RequestSignatureVerification) SetAllowedWeakAlgorithms(value *WeakAlgorithms)() {
    m.allowedWeakAlgorithms = value
}
// SetIsSignedRequestRequired sets the isSignedRequestRequired property value. Specifies whether signed authentication requests for this application should be required.
func (m *RequestSignatureVerification) SetIsSignedRequestRequired(value *bool)() {
    m.isSignedRequestRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RequestSignatureVerification) SetOdataType(value *string)() {
    m.odataType = value
}
