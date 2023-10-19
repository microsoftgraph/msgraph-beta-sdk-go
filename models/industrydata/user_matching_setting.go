package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserMatchingSetting 
type UserMatchingSetting struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserMatchingSetting instantiates a new userMatchingSetting and sets the default values.
func NewUserMatchingSetting()(*UserMatchingSetting) {
    m := &UserMatchingSetting{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserMatchingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserMatchingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserMatchingSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserMatchingSetting) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *UserMatchingSetting) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserMatchingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["matchTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserMatchTargetReferenceValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchTarget(val.(UserMatchTargetReferenceValueable))
        }
        return nil
    }
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
    res["priorityOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriorityOrder(val)
        }
        return nil
    }
    res["roleGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleGroup(val.(RoleGroupable))
        }
        return nil
    }
    res["sourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentifierTypeReferenceValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdentifier(val.(IdentifierTypeReferenceValueable))
        }
        return nil
    }
    return res
}
// GetMatchTarget gets the matchTarget property value. The RefUserMatchTarget for matching a user from the source with a Microsoft Entra user object.
func (m *UserMatchingSetting) GetMatchTarget()(UserMatchTargetReferenceValueable) {
    val, err := m.GetBackingStore().Get("matchTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserMatchTargetReferenceValueable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserMatchingSetting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPriorityOrder gets the priorityOrder property value. The priority order to apply when a user has multiple RefRole codes assigned.
func (m *UserMatchingSetting) GetPriorityOrder()(*int32) {
    val, err := m.GetBackingStore().Get("priorityOrder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRoleGroup gets the roleGroup property value. The roleGroup property
func (m *UserMatchingSetting) GetRoleGroup()(RoleGroupable) {
    val, err := m.GetBackingStore().Get("roleGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RoleGroupable)
    }
    return nil
}
// GetSourceIdentifier gets the sourceIdentifier property value. The sourceIdentifier property
func (m *UserMatchingSetting) GetSourceIdentifier()(IdentifierTypeReferenceValueable) {
    val, err := m.GetBackingStore().Get("sourceIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentifierTypeReferenceValueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserMatchingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("matchTarget", m.GetMatchTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priorityOrder", m.GetPriorityOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roleGroup", m.GetRoleGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceIdentifier", m.GetSourceIdentifier())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserMatchingSetting) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *UserMatchingSetting) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetMatchTarget sets the matchTarget property value. The RefUserMatchTarget for matching a user from the source with a Microsoft Entra user object.
func (m *UserMatchingSetting) SetMatchTarget(value UserMatchTargetReferenceValueable)() {
    err := m.GetBackingStore().Set("matchTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserMatchingSetting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPriorityOrder sets the priorityOrder property value. The priority order to apply when a user has multiple RefRole codes assigned.
func (m *UserMatchingSetting) SetPriorityOrder(value *int32)() {
    err := m.GetBackingStore().Set("priorityOrder", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleGroup sets the roleGroup property value. The roleGroup property
func (m *UserMatchingSetting) SetRoleGroup(value RoleGroupable)() {
    err := m.GetBackingStore().Set("roleGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceIdentifier sets the sourceIdentifier property value. The sourceIdentifier property
func (m *UserMatchingSetting) SetSourceIdentifier(value IdentifierTypeReferenceValueable)() {
    err := m.GetBackingStore().Set("sourceIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// UserMatchingSettingable 
type UserMatchingSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetMatchTarget()(UserMatchTargetReferenceValueable)
    GetOdataType()(*string)
    GetPriorityOrder()(*int32)
    GetRoleGroup()(RoleGroupable)
    GetSourceIdentifier()(IdentifierTypeReferenceValueable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetMatchTarget(value UserMatchTargetReferenceValueable)()
    SetOdataType(value *string)()
    SetPriorityOrder(value *int32)()
    SetRoleGroup(value RoleGroupable)()
    SetSourceIdentifier(value IdentifierTypeReferenceValueable)()
}
