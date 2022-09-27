package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpEvidence 
type IpEvidence struct {
    AlertEvidence
    // The value of the IP Address, can be either in V4 address or V6 address format.
    ipAddress *string
}
// NewIpEvidence instantiates a new IpEvidence and sets the default values.
func NewIpEvidence()(*IpEvidence) {
    m := &IpEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.ipEvidence";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIpEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    return res
}
// GetIpAddress gets the ipAddress property value. The value of the IP Address, can be either in V4 address or V6 address format.
func (m *IpEvidence) GetIpAddress()(*string) {
    return m.ipAddress
}
// Serialize serializes information the current object
func (m *IpEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIpAddress sets the ipAddress property value. The value of the IP Address, can be either in V4 address or V6 address format.
func (m *IpEvidence) SetIpAddress(value *string)() {
    m.ipAddress = value
}
