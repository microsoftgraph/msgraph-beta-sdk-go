package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditUserIdentity 
type AuditUserIdentity struct {
    UserIdentity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // For user sign ins, the identifier of the tenant that the user is a member of.
    homeTenantId *string
    // For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
    homeTenantName *string
}
// NewAuditUserIdentity instantiates a new AuditUserIdentity and sets the default values.
func NewAuditUserIdentity()(*AuditUserIdentity) {
    m := &AuditUserIdentity{
        UserIdentity: *NewUserIdentity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuditUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditUserIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditUserIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditUserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserIdentity.GetFieldDeserializers()
    res["homeTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
        }
        return nil
    }
    res["homeTenantName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantName(val)
        }
        return nil
    }
    return res
}
// GetHomeTenantId gets the homeTenantId property value. For user sign ins, the identifier of the tenant that the user is a member of.
func (m *AuditUserIdentity) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// GetHomeTenantName gets the homeTenantName property value. For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *AuditUserIdentity) GetHomeTenantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantName
    }
}
// Serialize serializes information the current object
func (m *AuditUserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantName", m.GetHomeTenantName())
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
func (m *AuditUserIdentity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHomeTenantId sets the homeTenantId property value. For user sign ins, the identifier of the tenant that the user is a member of.
func (m *AuditUserIdentity) SetHomeTenantId(value *string)() {
    if m != nil {
        m.homeTenantId = value
    }
}
// SetHomeTenantName sets the homeTenantName property value. For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *AuditUserIdentity) SetHomeTenantName(value *string)() {
    if m != nil {
        m.homeTenantName = value
    }
}
