package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SkypeForBusinessUserConversationMember 
type SkypeForBusinessUserConversationMember struct {
    ConversationMember
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // ID of the tenant that the user belongs to.
    tenantId *string
    // ID of the user.
    userId *string
}
// NewSkypeForBusinessUserConversationMember instantiates a new SkypeForBusinessUserConversationMember and sets the default values.
func NewSkypeForBusinessUserConversationMember()(*SkypeForBusinessUserConversationMember) {
    m := &SkypeForBusinessUserConversationMember{
        ConversationMember: *NewConversationMember(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSkypeForBusinessUserConversationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSkypeForBusinessUserConversationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSkypeForBusinessUserConversationMember(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SkypeForBusinessUserConversationMember) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SkypeForBusinessUserConversationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConversationMember.GetFieldDeserializers()
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. ID of the tenant that the user belongs to.
func (m *SkypeForBusinessUserConversationMember) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetUserId gets the userId property value. ID of the user.
func (m *SkypeForBusinessUserConversationMember) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Serialize serializes information the current object
func (m *SkypeForBusinessUserConversationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConversationMember.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
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
func (m *SkypeForBusinessUserConversationMember) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTenantId sets the tenantId property value. ID of the tenant that the user belongs to.
func (m *SkypeForBusinessUserConversationMember) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetUserId sets the userId property value. ID of the user.
func (m *SkypeForBusinessUserConversationMember) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
