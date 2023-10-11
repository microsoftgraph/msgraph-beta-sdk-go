package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkDisplayConfiguration 
type TeamworkDisplayConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkDisplayConfiguration instantiates a new teamworkDisplayConfiguration and sets the default values.
func NewTeamworkDisplayConfiguration()(*TeamworkDisplayConfiguration) {
    m := &TeamworkDisplayConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkDisplayConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDisplayConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDisplayConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkDisplayConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConfiguredDisplays gets the configuredDisplays property value. The list of configured displays. Applicable only for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetConfiguredDisplays()([]TeamworkConfiguredPeripheralable) {
    val, err := m.GetBackingStore().Get("configuredDisplays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamworkConfiguredPeripheralable)
    }
    return nil
}
// GetDisplayCount gets the displayCount property value. Total number of connected displays, including the inbuilt display. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetDisplayCount()(*int32) {
    val, err := m.GetBackingStore().Get("displayCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDisplayConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuredDisplays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkConfiguredPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkConfiguredPeripheralable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamworkConfiguredPeripheralable)
                }
            }
            m.SetConfiguredDisplays(res)
        }
        return nil
    }
    res["displayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayCount(val)
        }
        return nil
    }
    res["inBuiltDisplayScreenConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDisplayScreenConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInBuiltDisplayScreenConfiguration(val.(TeamworkDisplayScreenConfigurationable))
        }
        return nil
    }
    res["isContentDuplicationAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentDuplicationAllowed(val)
        }
        return nil
    }
    res["isDualDisplayModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDualDisplayModeEnabled(val)
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
    return res
}
// GetInBuiltDisplayScreenConfiguration gets the inBuiltDisplayScreenConfiguration property value. Configuration for the inbuilt display. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetInBuiltDisplayScreenConfiguration()(TeamworkDisplayScreenConfigurationable) {
    val, err := m.GetBackingStore().Get("inBuiltDisplayScreenConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDisplayScreenConfigurationable)
    }
    return nil
}
// GetIsContentDuplicationAllowed gets the isContentDuplicationAllowed property value. True if content duplication is allowed. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetIsContentDuplicationAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("isContentDuplicationAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsDualDisplayModeEnabled gets the isDualDisplayModeEnabled property value. True if dual display mode is enabled. If isDualDisplayModeEnabled is true, then the content will be displayed on both front of room screens instead of just the one screen, when it is shared via the HDMI ingest module on the Microsoft Teams Rooms device. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetIsDualDisplayModeEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isDualDisplayModeEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkDisplayConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConfiguredDisplays() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfiguredDisplays()))
        for i, v := range m.GetConfiguredDisplays() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("configuredDisplays", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("displayCount", m.GetDisplayCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("inBuiltDisplayScreenConfiguration", m.GetInBuiltDisplayScreenConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentDuplicationAllowed", m.GetIsContentDuplicationAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDualDisplayModeEnabled", m.GetIsDualDisplayModeEnabled())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkDisplayConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConfiguredDisplays sets the configuredDisplays property value. The list of configured displays. Applicable only for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetConfiguredDisplays(value []TeamworkConfiguredPeripheralable)() {
    err := m.GetBackingStore().Set("configuredDisplays", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayCount sets the displayCount property value. Total number of connected displays, including the inbuilt display. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetDisplayCount(value *int32)() {
    err := m.GetBackingStore().Set("displayCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInBuiltDisplayScreenConfiguration sets the inBuiltDisplayScreenConfiguration property value. Configuration for the inbuilt display. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetInBuiltDisplayScreenConfiguration(value TeamworkDisplayScreenConfigurationable)() {
    err := m.GetBackingStore().Set("inBuiltDisplayScreenConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetIsContentDuplicationAllowed sets the isContentDuplicationAllowed property value. True if content duplication is allowed. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetIsContentDuplicationAllowed(value *bool)() {
    err := m.GetBackingStore().Set("isContentDuplicationAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDualDisplayModeEnabled sets the isDualDisplayModeEnabled property value. True if dual display mode is enabled. If isDualDisplayModeEnabled is true, then the content will be displayed on both front of room screens instead of just the one screen, when it is shared via the HDMI ingest module on the Microsoft Teams Rooms device. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetIsDualDisplayModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isDualDisplayModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkDisplayConfigurationable 
type TeamworkDisplayConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConfiguredDisplays()([]TeamworkConfiguredPeripheralable)
    GetDisplayCount()(*int32)
    GetInBuiltDisplayScreenConfiguration()(TeamworkDisplayScreenConfigurationable)
    GetIsContentDuplicationAllowed()(*bool)
    GetIsDualDisplayModeEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConfiguredDisplays(value []TeamworkConfiguredPeripheralable)()
    SetDisplayCount(value *int32)()
    SetInBuiltDisplayScreenConfiguration(value TeamworkDisplayScreenConfigurationable)()
    SetIsContentDuplicationAllowed(value *bool)()
    SetIsDualDisplayModeEnabled(value *bool)()
    SetOdataType(value *string)()
}
