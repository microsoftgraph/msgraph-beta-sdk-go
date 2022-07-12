package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminAccessContainer 
type DelegatedAdminAccessContainer struct {
    // The identifier of the access container (for example, a security group). For 'securityGroup' access containers, this must be a valid ID of an Azure AD security group in the Microsoft partner's tenant.
    accessContainerId *string
    // The accessContainerType property
    accessContainerType *DelegatedAdminAccessContainerType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewDelegatedAdminAccessContainer instantiates a new delegatedAdminAccessContainer and sets the default values.
func NewDelegatedAdminAccessContainer()(*DelegatedAdminAccessContainer) {
    m := &DelegatedAdminAccessContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDelegatedAdminAccessContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminAccessContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminAccessContainer(), nil
}
// GetAccessContainerId gets the accessContainerId property value. The identifier of the access container (for example, a security group). For 'securityGroup' access containers, this must be a valid ID of an Azure AD security group in the Microsoft partner's tenant.
func (m *DelegatedAdminAccessContainer) GetAccessContainerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessContainerId
    }
}
// GetAccessContainerType gets the accessContainerType property value. The accessContainerType property
func (m *DelegatedAdminAccessContainer) GetAccessContainerType()(*DelegatedAdminAccessContainerType) {
    if m == nil {
        return nil
    } else {
        return m.accessContainerType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DelegatedAdminAccessContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminAccessContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessContainerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessContainerId(val)
        }
        return nil
    }
    res["accessContainerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminAccessContainerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessContainerType(val.(*DelegatedAdminAccessContainerType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DelegatedAdminAccessContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accessContainerId", m.GetAccessContainerId())
        if err != nil {
            return err
        }
    }
    if m.GetAccessContainerType() != nil {
        cast := (*m.GetAccessContainerType()).String()
        err := writer.WriteStringValue("accessContainerType", &cast)
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
// SetAccessContainerId sets the accessContainerId property value. The identifier of the access container (for example, a security group). For 'securityGroup' access containers, this must be a valid ID of an Azure AD security group in the Microsoft partner's tenant.
func (m *DelegatedAdminAccessContainer) SetAccessContainerId(value *string)() {
    if m != nil {
        m.accessContainerId = value
    }
}
// SetAccessContainerType sets the accessContainerType property value. The accessContainerType property
func (m *DelegatedAdminAccessContainer) SetAccessContainerType(value *DelegatedAdminAccessContainerType)() {
    if m != nil {
        m.accessContainerType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DelegatedAdminAccessContainer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
