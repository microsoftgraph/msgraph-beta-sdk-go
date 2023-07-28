package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifiableCredentialRetrieved 
type VerifiableCredentialRetrieved struct {
    VerifiableCredentialRequirementStatus
}
// NewVerifiableCredentialRetrieved instantiates a new verifiableCredentialRetrieved and sets the default values.
func NewVerifiableCredentialRetrieved()(*VerifiableCredentialRetrieved) {
    m := &VerifiableCredentialRetrieved{
        VerifiableCredentialRequirementStatus: *NewVerifiableCredentialRequirementStatus(),
    }
    odataTypeValue := "#microsoft.graph.verifiableCredentialRetrieved"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVerifiableCredentialRetrievedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifiableCredentialRetrievedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiableCredentialRetrieved(), nil
}
// GetExpiryDateTime gets the expiryDateTime property value. The specific date and time that the presentation request will expire and a new one will need to be generated.
func (m *VerifiableCredentialRetrieved) GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *VerifiableCredentialRetrieved) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *VerifiableCredentialRetrieved) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetExpiryDateTime sets the expiryDateTime property value. The specific date and time that the presentation request will expire and a new one will need to be generated.
func (m *VerifiableCredentialRetrieved) SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expiryDateTime", value)
    if err != nil {
        panic(err)
    }
}
// VerifiableCredentialRetrievedable 
type VerifiableCredentialRetrievedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VerifiableCredentialRequirementStatusable
    GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
