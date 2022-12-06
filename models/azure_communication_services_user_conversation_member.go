package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureCommunicationServicesUserConversationMember 
type AzureCommunicationServicesUserConversationMember struct {
    ConversationMember
    // The azureCommunicationServicesId property
    azureCommunicationServicesId *string
}
// NewAzureCommunicationServicesUserConversationMember instantiates a new AzureCommunicationServicesUserConversationMember and sets the default values.
func NewAzureCommunicationServicesUserConversationMember()(*AzureCommunicationServicesUserConversationMember) {
    m := &AzureCommunicationServicesUserConversationMember{
        ConversationMember: *NewConversationMember(),
    }
    odataTypeValue := "#microsoft.graph.azureCommunicationServicesUserConversationMember";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAzureCommunicationServicesUserConversationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureCommunicationServicesUserConversationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureCommunicationServicesUserConversationMember(), nil
}
// GetAzureCommunicationServicesId gets the azureCommunicationServicesId property value. The azureCommunicationServicesId property
func (m *AzureCommunicationServicesUserConversationMember) GetAzureCommunicationServicesId()(*string) {
    return m.azureCommunicationServicesId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureCommunicationServicesUserConversationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConversationMember.GetFieldDeserializers()
    res["azureCommunicationServicesId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureCommunicationServicesId)
    return res
}
// Serialize serializes information the current object
func (m *AzureCommunicationServicesUserConversationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConversationMember.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureCommunicationServicesId", m.GetAzureCommunicationServicesId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureCommunicationServicesId sets the azureCommunicationServicesId property value. The azureCommunicationServicesId property
func (m *AzureCommunicationServicesUserConversationMember) SetAzureCommunicationServicesId(value *string)() {
    m.azureCommunicationServicesId = value
}
