package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserConsentRequest 
type UserConsentRequest struct {
    Request
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Approval decisions associated with a request.
    approval Approvalable
    // The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
    reason *string
}
// NewUserConsentRequest instantiates a new UserConsentRequest and sets the default values.
func NewUserConsentRequest()(*UserConsentRequest) {
    m := &UserConsentRequest{
        Request: *NewRequest(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserConsentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserConsentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserConsentRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserConsentRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApproval gets the approval property value. Approval decisions associated with a request.
func (m *UserConsentRequest) GetApproval()(Approvalable) {
    if m == nil {
        return nil
    } else {
        return m.approval
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserConsentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["approval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproval(val.(Approvalable))
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Serialize serializes information the current object
func (m *UserConsentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Request.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("approval", m.GetApproval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reason", m.GetReason())
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
func (m *UserConsentRequest) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApproval sets the approval property value. Approval decisions associated with a request.
func (m *UserConsentRequest) SetApproval(value Approvalable)() {
    if m != nil {
        m.approval = value
    }
}
// SetReason sets the reason property value. The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
func (m *UserConsentRequest) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
