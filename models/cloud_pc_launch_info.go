package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CloudPcLaunchInfo 
type CloudPcLaunchInfo struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcLaunchInfo instantiates a new cloudPcLaunchInfo and sets the default values.
func NewCloudPcLaunchInfo()(*CloudPcLaunchInfo) {
    m := &CloudPcLaunchInfo{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcLaunchInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcLaunchInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcLaunchInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcLaunchInfo) GetAdditionalData()(map[string]any) {
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
func (m *CloudPcLaunchInfo) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCloudPcId gets the cloudPcId property value. The unique identifier of the Cloud PC.
func (m *CloudPcLaunchInfo) GetCloudPcId()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCloudPcLaunchUrl gets the cloudPcLaunchUrl property value. The connect URL of the Cloud PC.
func (m *CloudPcLaunchInfo) GetCloudPcLaunchUrl()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcLaunchUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcLaunchInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudPcId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcId(val)
        }
        return nil
    }
    res["cloudPcLaunchUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcLaunchUrl(val)
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
    res["windows365SwitchCompatible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows365SwitchCompatible(val)
        }
        return nil
    }
    res["windows365SwitchNotCompatibleReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows365SwitchNotCompatibleReason(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcLaunchInfo) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWindows365SwitchCompatible gets the windows365SwitchCompatible property value. Indicates whether the Cloud PC supports switch functionality. If the value is true, it supports switch functionality; otherwise,  false.
func (m *CloudPcLaunchInfo) GetWindows365SwitchCompatible()(*bool) {
    val, err := m.GetBackingStore().Get("windows365SwitchCompatible")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWindows365SwitchNotCompatibleReason gets the windows365SwitchNotCompatibleReason property value. Indicates the reason the Cloud PC doesn't support switch. CPCOsVersionNotMeetRequirement indicates that the user needs to update their Cloud PC operation system version. CPCHardwareNotMeetRequirement indicates that the Cloud PC needs more CPU or RAM to support the functionality.
func (m *CloudPcLaunchInfo) GetWindows365SwitchNotCompatibleReason()(*string) {
    val, err := m.GetBackingStore().Get("windows365SwitchNotCompatibleReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcLaunchInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudPcLaunchUrl", m.GetCloudPcLaunchUrl())
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
        err := writer.WriteBoolValue("windows365SwitchCompatible", m.GetWindows365SwitchCompatible())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("windows365SwitchNotCompatibleReason", m.GetWindows365SwitchNotCompatibleReason())
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
func (m *CloudPcLaunchInfo) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudPcLaunchInfo) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCloudPcId sets the cloudPcId property value. The unique identifier of the Cloud PC.
func (m *CloudPcLaunchInfo) SetCloudPcId(value *string)() {
    err := m.GetBackingStore().Set("cloudPcId", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcLaunchUrl sets the cloudPcLaunchUrl property value. The connect URL of the Cloud PC.
func (m *CloudPcLaunchInfo) SetCloudPcLaunchUrl(value *string)() {
    err := m.GetBackingStore().Set("cloudPcLaunchUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcLaunchInfo) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetWindows365SwitchCompatible sets the windows365SwitchCompatible property value. Indicates whether the Cloud PC supports switch functionality. If the value is true, it supports switch functionality; otherwise,  false.
func (m *CloudPcLaunchInfo) SetWindows365SwitchCompatible(value *bool)() {
    err := m.GetBackingStore().Set("windows365SwitchCompatible", value)
    if err != nil {
        panic(err)
    }
}
// SetWindows365SwitchNotCompatibleReason sets the windows365SwitchNotCompatibleReason property value. Indicates the reason the Cloud PC doesn't support switch. CPCOsVersionNotMeetRequirement indicates that the user needs to update their Cloud PC operation system version. CPCHardwareNotMeetRequirement indicates that the Cloud PC needs more CPU or RAM to support the functionality.
func (m *CloudPcLaunchInfo) SetWindows365SwitchNotCompatibleReason(value *string)() {
    err := m.GetBackingStore().Set("windows365SwitchNotCompatibleReason", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcLaunchInfoable 
type CloudPcLaunchInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCloudPcId()(*string)
    GetCloudPcLaunchUrl()(*string)
    GetOdataType()(*string)
    GetWindows365SwitchCompatible()(*bool)
    GetWindows365SwitchNotCompatibleReason()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCloudPcId(value *string)()
    SetCloudPcLaunchUrl(value *string)()
    SetOdataType(value *string)()
    SetWindows365SwitchCompatible(value *bool)()
    SetWindows365SwitchNotCompatibleReason(value *string)()
}
