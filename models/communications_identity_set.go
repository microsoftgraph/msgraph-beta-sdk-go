package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsIdentitySet 
type CommunicationsIdentitySet struct {
    IdentitySet
    // The assertedIdentity property
    assertedIdentity Identityable
    // The azureCommunicationServicesUser property
    azureCommunicationServicesUser Identityable
    // The endpointType property
    endpointType *EndpointType
}
// NewCommunicationsIdentitySet instantiates a new CommunicationsIdentitySet and sets the default values.
func NewCommunicationsIdentitySet()(*CommunicationsIdentitySet) {
    m := &CommunicationsIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.communicationsIdentitySet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCommunicationsIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsIdentitySet(), nil
}
// GetAssertedIdentity gets the assertedIdentity property value. The assertedIdentity property
func (m *CommunicationsIdentitySet) GetAssertedIdentity()(Identityable) {
    if m == nil {
        return nil
    } else {
        return m.assertedIdentity
    }
}
// GetAzureCommunicationServicesUser gets the azureCommunicationServicesUser property value. The azureCommunicationServicesUser property
func (m *CommunicationsIdentitySet) GetAzureCommunicationServicesUser()(Identityable) {
    if m == nil {
        return nil
    } else {
        return m.azureCommunicationServicesUser
    }
}
// GetEndpointType gets the endpointType property value. The endpointType property
func (m *CommunicationsIdentitySet) GetEndpointType()(*EndpointType) {
    if m == nil {
        return nil
    } else {
        return m.endpointType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["assertedIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssertedIdentity(val.(Identityable))
        }
        return nil
    }
    res["azureCommunicationServicesUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureCommunicationServicesUser(val.(Identityable))
        }
        return nil
    }
    res["endpointType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointType(val.(*EndpointType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CommunicationsIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("assertedIdentity", m.GetAssertedIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("azureCommunicationServicesUser", m.GetAzureCommunicationServicesUser())
        if err != nil {
            return err
        }
    }
    if m.GetEndpointType() != nil {
        cast := (*m.GetEndpointType()).String()
        err = writer.WriteStringValue("endpointType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssertedIdentity sets the assertedIdentity property value. The assertedIdentity property
func (m *CommunicationsIdentitySet) SetAssertedIdentity(value Identityable)() {
    if m != nil {
        m.assertedIdentity = value
    }
}
// SetAzureCommunicationServicesUser sets the azureCommunicationServicesUser property value. The azureCommunicationServicesUser property
func (m *CommunicationsIdentitySet) SetAzureCommunicationServicesUser(value Identityable)() {
    if m != nil {
        m.azureCommunicationServicesUser = value
    }
}
// SetEndpointType sets the endpointType property value. The endpointType property
func (m *CommunicationsIdentitySet) SetEndpointType(value *EndpointType)() {
    if m != nil {
        m.endpointType = value
    }
}
