package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifiableCredentialRequired 
type VerifiableCredentialRequired struct {
    VerifiableCredentialRequirementStatus
}
// NewVerifiableCredentialRequired instantiates a new VerifiableCredentialRequired and sets the default values.
func NewVerifiableCredentialRequired()(*VerifiableCredentialRequired) {
    m := &VerifiableCredentialRequired{
        VerifiableCredentialRequirementStatus: *NewVerifiableCredentialRequirementStatus(),
    }
    odataTypeValue := "#microsoft.graph.verifiableCredentialRequired"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVerifiableCredentialRequiredFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifiableCredentialRequiredFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiableCredentialRequired(), nil
}
// GetExpiryDateTime gets the expiryDateTime property value. When the presentation request will expire and a new one will need to be generated.
func (m *VerifiableCredentialRequired) GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expiryDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifiableCredentialRequired) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VerifiableCredentialRequirementStatus.GetFieldDeserializers()
    res["expiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDateTime(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. A URL that launches the digital wallet and starts the presentation process. You can present this URL to the user if they can't scan the QR code.
func (m *VerifiableCredentialRequired) GetUrl()(*string) {
    val, err := m.GetBackingStore().Get("url")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VerifiableCredentialRequired) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VerifiableCredentialRequirementStatus.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expiryDateTime", m.GetExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpiryDateTime sets the expiryDateTime property value. When the presentation request will expire and a new one will need to be generated.
func (m *VerifiableCredentialRequired) SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expiryDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUrl sets the url property value. A URL that launches the digital wallet and starts the presentation process. You can present this URL to the user if they can't scan the QR code.
func (m *VerifiableCredentialRequired) SetUrl(value *string)() {
    err := m.GetBackingStore().Set("url", value)
    if err != nil {
        panic(err)
    }
}
// VerifiableCredentialRequiredable 
type VerifiableCredentialRequiredable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VerifiableCredentialRequirementStatusable
    GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
