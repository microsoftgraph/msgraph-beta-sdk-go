package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectSite 
type ProtectSite struct {
    LabelActionBase
    // The accessType property
    accessType *SiteAccessType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The conditionalAccessProtectionLevelId property
    conditionalAccessProtectionLevelId *string
}
// NewProtectSite instantiates a new ProtectSite and sets the default values.
func NewProtectSite()(*ProtectSite) {
    m := &ProtectSite{
        LabelActionBase: *NewLabelActionBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProtectSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectSite(), nil
}
// GetAccessType gets the accessType property value. The accessType property
func (m *ProtectSite) GetAccessType()(*SiteAccessType) {
    if m == nil {
        return nil
    } else {
        return m.accessType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectSite) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConditionalAccessProtectionLevelId gets the conditionalAccessProtectionLevelId property value. The conditionalAccessProtectionLevelId property
func (m *ProtectSite) GetConditionalAccessProtectionLevelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessProtectionLevelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectSite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["accessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSiteAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*SiteAccessType))
        }
        return nil
    }
    res["conditionalAccessProtectionLevelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessProtectionLevelId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProtectSite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err = writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conditionalAccessProtectionLevelId", m.GetConditionalAccessProtectionLevelId())
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
// SetAccessType sets the accessType property value. The accessType property
func (m *ProtectSite) SetAccessType(value *SiteAccessType)() {
    if m != nil {
        m.accessType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectSite) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConditionalAccessProtectionLevelId sets the conditionalAccessProtectionLevelId property value. The conditionalAccessProtectionLevelId property
func (m *ProtectSite) SetConditionalAccessProtectionLevelId(value *string)() {
    if m != nil {
        m.conditionalAccessProtectionLevelId = value
    }
}
