package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsApplicationIdentity 
type CommunicationsApplicationIdentity struct {
    Identity
    // The applicationType property
    applicationType *string
    // The hidden property
    hidden *bool
}
// NewCommunicationsApplicationIdentity instantiates a new CommunicationsApplicationIdentity and sets the default values.
func NewCommunicationsApplicationIdentity()(*CommunicationsApplicationIdentity) {
    m := &CommunicationsApplicationIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.communicationsApplicationIdentity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCommunicationsApplicationIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsApplicationIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsApplicationIdentity(), nil
}
// GetApplicationType gets the applicationType property value. The applicationType property
func (m *CommunicationsApplicationIdentity) GetApplicationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsApplicationIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["applicationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationType(val)
        }
        return nil
    }
    res["hidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidden(val)
        }
        return nil
    }
    return res
}
// GetHidden gets the hidden property value. The hidden property
func (m *CommunicationsApplicationIdentity) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// Serialize serializes information the current object
func (m *CommunicationsApplicationIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationType", m.GetApplicationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationType sets the applicationType property value. The applicationType property
func (m *CommunicationsApplicationIdentity) SetApplicationType(value *string)() {
    if m != nil {
        m.applicationType = value
    }
}
// SetHidden sets the hidden property value. The hidden property
func (m *CommunicationsApplicationIdentity) SetHidden(value *bool)() {
    if m != nil {
        m.hidden = value
    }
}
