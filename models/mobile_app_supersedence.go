package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppSupersedence 
type MobileAppSupersedence struct {
    MobileAppRelationship
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The total number of apps directly or indirectly superseded by the child app.
    supersededAppCount *int32
    // The supersedence relationship type between the parent and child apps. Possible values are: update, replace.
    supersedenceType *MobileAppSupersedenceType
    // The total number of apps directly or indirectly superseding the parent app.
    supersedingAppCount *int32
}
// NewMobileAppSupersedence instantiates a new MobileAppSupersedence and sets the default values.
func NewMobileAppSupersedence()(*MobileAppSupersedence) {
    m := &MobileAppSupersedence{
        MobileAppRelationship: *NewMobileAppRelationship(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppSupersedenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppSupersedenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppSupersedence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppSupersedence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppSupersedence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppRelationship.GetFieldDeserializers()
    res["supersededAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersededAppCount(val)
        }
        return nil
    }
    res["supersedenceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppSupersedenceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersedenceType(val.(*MobileAppSupersedenceType))
        }
        return nil
    }
    res["supersedingAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersedingAppCount(val)
        }
        return nil
    }
    return res
}
// GetSupersededAppCount gets the supersededAppCount property value. The total number of apps directly or indirectly superseded by the child app.
func (m *MobileAppSupersedence) GetSupersededAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supersededAppCount
    }
}
// GetSupersedenceType gets the supersedenceType property value. The supersedence relationship type between the parent and child apps. Possible values are: update, replace.
func (m *MobileAppSupersedence) GetSupersedenceType()(*MobileAppSupersedenceType) {
    if m == nil {
        return nil
    } else {
        return m.supersedenceType
    }
}
// GetSupersedingAppCount gets the supersedingAppCount property value. The total number of apps directly or indirectly superseding the parent app.
func (m *MobileAppSupersedence) GetSupersedingAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supersedingAppCount
    }
}
// Serialize serializes information the current object
func (m *MobileAppSupersedence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppRelationship.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("supersededAppCount", m.GetSupersededAppCount())
        if err != nil {
            return err
        }
    }
    if m.GetSupersedenceType() != nil {
        cast := (*m.GetSupersedenceType()).String()
        err = writer.WriteStringValue("supersedenceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supersedingAppCount", m.GetSupersedingAppCount())
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
func (m *MobileAppSupersedence) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSupersededAppCount sets the supersededAppCount property value. The total number of apps directly or indirectly superseded by the child app.
func (m *MobileAppSupersedence) SetSupersededAppCount(value *int32)() {
    if m != nil {
        m.supersededAppCount = value
    }
}
// SetSupersedenceType sets the supersedenceType property value. The supersedence relationship type between the parent and child apps. Possible values are: update, replace.
func (m *MobileAppSupersedence) SetSupersedenceType(value *MobileAppSupersedenceType)() {
    if m != nil {
        m.supersedenceType = value
    }
}
// SetSupersedingAppCount sets the supersedingAppCount property value. The total number of apps directly or indirectly superseding the parent app.
func (m *MobileAppSupersedence) SetSupersedingAppCount(value *int32)() {
    if m != nil {
        m.supersedingAppCount = value
    }
}
