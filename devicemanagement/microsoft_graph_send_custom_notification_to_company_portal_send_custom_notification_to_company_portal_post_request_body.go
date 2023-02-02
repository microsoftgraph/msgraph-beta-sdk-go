package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody 
type MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The groupsToNotify property
    groupsToNotify []string
    // The notificationBody property
    notificationBody *string
    // The notificationTitle property
    notificationTitle *string
}
// NewMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody instantiates a new MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody and sets the default values.
func NewMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody()(*MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) {
    m := &MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupsToNotify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupsToNotify(res)
        }
        return nil
    }
    res["notificationBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationBody(val)
        }
        return nil
    }
    res["notificationTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTitle(val)
        }
        return nil
    }
    return res
}
// GetGroupsToNotify gets the groupsToNotify property value. The groupsToNotify property
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetGroupsToNotify()([]string) {
    return m.groupsToNotify
}
// GetNotificationBody gets the notificationBody property value. The notificationBody property
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationBody()(*string) {
    return m.notificationBody
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationTitle()(*string) {
    return m.notificationTitle
}
// Serialize serializes information the current object
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupsToNotify() != nil {
        err := writer.WriteCollectionOfStringValues("groupsToNotify", m.GetGroupsToNotify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationBody", m.GetNotificationBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTitle", m.GetNotificationTitle())
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
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupsToNotify sets the groupsToNotify property value. The groupsToNotify property
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetGroupsToNotify(value []string)() {
    m.groupsToNotify = value
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
