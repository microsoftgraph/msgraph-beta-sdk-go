package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppSupersedence describes a supersedence relationship between two mobile apps.
type MobileAppSupersedence struct {
    MobileAppRelationship
}
// NewMobileAppSupersedence instantiates a new mobileAppSupersedence and sets the default values.
func NewMobileAppSupersedence()(*MobileAppSupersedence) {
    m := &MobileAppSupersedence{
        MobileAppRelationship: *NewMobileAppRelationship(),
    }
    odataTypeValue := "#microsoft.graph.mobileAppSupersedence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMobileAppSupersedenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppSupersedenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppSupersedence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppSupersedence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppRelationship.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MobileAppSupersedence) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupersededAppCount gets the supersededAppCount property value. The total number of apps directly or indirectly superseded by the child app.
func (m *MobileAppSupersedence) GetSupersededAppCount()(*int32) {
    val, err := m.GetBackingStore().Get("supersededAppCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSupersedenceType gets the supersedenceType property value. Indicates the supersedence type associated with a relationship between two mobile apps.
func (m *MobileAppSupersedence) GetSupersedenceType()(*MobileAppSupersedenceType) {
    val, err := m.GetBackingStore().Get("supersedenceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileAppSupersedenceType)
    }
    return nil
}
// GetSupersedingAppCount gets the supersedingAppCount property value. The total number of apps directly or indirectly superseding the parent app.
func (m *MobileAppSupersedence) GetSupersedingAppCount()(*int32) {
    val, err := m.GetBackingStore().Get("supersedingAppCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppSupersedence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppRelationship.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
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
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MobileAppSupersedence) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSupersededAppCount sets the supersededAppCount property value. The total number of apps directly or indirectly superseded by the child app.
func (m *MobileAppSupersedence) SetSupersededAppCount(value *int32)() {
    err := m.GetBackingStore().Set("supersededAppCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSupersedenceType sets the supersedenceType property value. Indicates the supersedence type associated with a relationship between two mobile apps.
func (m *MobileAppSupersedence) SetSupersedenceType(value *MobileAppSupersedenceType)() {
    err := m.GetBackingStore().Set("supersedenceType", value)
    if err != nil {
        panic(err)
    }
}
// SetSupersedingAppCount sets the supersedingAppCount property value. The total number of apps directly or indirectly superseding the parent app.
func (m *MobileAppSupersedence) SetSupersedingAppCount(value *int32)() {
    err := m.GetBackingStore().Set("supersedingAppCount", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppSupersedenceable 
type MobileAppSupersedenceable interface {
    MobileAppRelationshipable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetSupersededAppCount()(*int32)
    GetSupersedenceType()(*MobileAppSupersedenceType)
    GetSupersedingAppCount()(*int32)
    SetOdataType(value *string)()
    SetSupersededAppCount(value *int32)()
    SetSupersedenceType(value *MobileAppSupersedenceType)()
    SetSupersedingAppCount(value *int32)()
}
