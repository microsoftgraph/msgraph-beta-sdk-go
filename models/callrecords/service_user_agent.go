package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceUserAgent 
type ServiceUserAgent struct {
    UserAgent
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifies the role of the service used by this endpoint. Possible values are: unknown, customBot, skypeForBusinessMicrosoftTeamsGateway, skypeForBusinessAudioVideoMcu, skypeForBusinessApplicationSharingMcu, skypeForBusinessCallQueues, skypeForBusinessAutoAttendant, mediationServer, mediationServerCloudConnectorEdition, exchangeUnifiedMessagingService, mediaController, conferencingAnnouncementService, conferencingAttendant, audioTeleconferencerController, skypeForBusinessUnifiedCommunicationApplicationPlatform, responseGroupServiceAnnouncementService, gateway, skypeTranslator, skypeForBusinessAttendant, responseGroupService, voicemail, unknownFutureValue.
    role *ServiceRole
}
// NewServiceUserAgent instantiates a new ServiceUserAgent and sets the default values.
func NewServiceUserAgent()(*ServiceUserAgent) {
    m := &ServiceUserAgent{
        UserAgent: *NewUserAgent(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServiceUserAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceUserAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceUserAgent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceUserAgent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceUserAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserAgent.GetFieldDeserializers()
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*ServiceRole))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. Identifies the role of the service used by this endpoint. Possible values are: unknown, customBot, skypeForBusinessMicrosoftTeamsGateway, skypeForBusinessAudioVideoMcu, skypeForBusinessApplicationSharingMcu, skypeForBusinessCallQueues, skypeForBusinessAutoAttendant, mediationServer, mediationServerCloudConnectorEdition, exchangeUnifiedMessagingService, mediaController, conferencingAnnouncementService, conferencingAttendant, audioTeleconferencerController, skypeForBusinessUnifiedCommunicationApplicationPlatform, responseGroupServiceAnnouncementService, gateway, skypeTranslator, skypeForBusinessAttendant, responseGroupService, voicemail, unknownFutureValue.
func (m *ServiceUserAgent) GetRole()(*ServiceRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// Serialize serializes information the current object
func (m *ServiceUserAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserAgent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
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
func (m *ServiceUserAgent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRole sets the role property value. Identifies the role of the service used by this endpoint. Possible values are: unknown, customBot, skypeForBusinessMicrosoftTeamsGateway, skypeForBusinessAudioVideoMcu, skypeForBusinessApplicationSharingMcu, skypeForBusinessCallQueues, skypeForBusinessAutoAttendant, mediationServer, mediationServerCloudConnectorEdition, exchangeUnifiedMessagingService, mediaController, conferencingAnnouncementService, conferencingAttendant, audioTeleconferencerController, skypeForBusinessUnifiedCommunicationApplicationPlatform, responseGroupServiceAnnouncementService, gateway, skypeTranslator, skypeForBusinessAttendant, responseGroupService, voicemail, unknownFutureValue.
func (m *ServiceUserAgent) SetRole(value *ServiceRole)() {
    if m != nil {
        m.role = value
    }
}
