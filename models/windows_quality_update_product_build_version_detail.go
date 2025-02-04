package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsQualityUpdateProductBuildVersionDetail represents the build version details of a product revision that is associated with a quality update.
type WindowsQualityUpdateProductBuildVersionDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsQualityUpdateProductBuildVersionDetail instantiates a new WindowsQualityUpdateProductBuildVersionDetail and sets the default values.
func NewWindowsQualityUpdateProductBuildVersionDetail()(*WindowsQualityUpdateProductBuildVersionDetail) {
    m := &WindowsQualityUpdateProductBuildVersionDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsQualityUpdateProductBuildVersionDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsQualityUpdateProductBuildVersionDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateProductBuildVersionDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBuildNumber gets the buildNumber property value. The build number of the product release, Allowed range is 0 - 2,147,483,647. For example: 19045. Read-only.
// returns a *int32 when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetBuildNumber()(*int32) {
    val, err := m.GetBackingStore().Get("buildNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildNumber(val)
        }
        return nil
    }
    res["majorVersionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorVersionNumber(val)
        }
        return nil
    }
    res["minorVersionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorVersionNumber(val)
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
    res["updateBuildRevision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateBuildRevision(val)
        }
        return nil
    }
    return res
}
// GetMajorVersionNumber gets the majorVersionNumber property value. The major version of the product release, Allowed range is 0 - 2,147,483,647. For example: 10. Read-only.
// returns a *int32 when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetMajorVersionNumber()(*int32) {
    val, err := m.GetBackingStore().Get("majorVersionNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinorVersionNumber gets the minorVersionNumber property value. The minor version of the product release, Allowed range is 0 - 2,147,483,647. For example: 0. Read-only.
// returns a *int32 when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetMinorVersionNumber()(*int32) {
    val, err := m.GetBackingStore().Get("minorVersionNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUpdateBuildRevision gets the updateBuildRevision property value. The update build revision number of the product revision for the corresponding patch, Allowed range is 0 - 2,147,483,647. For example: 4780. Read-only.
// returns a *int32 when successful
func (m *WindowsQualityUpdateProductBuildVersionDetail) GetUpdateBuildRevision()(*int32) {
    val, err := m.GetBackingStore().Get("updateBuildRevision")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateProductBuildVersionDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("buildNumber", m.GetBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("majorVersionNumber", m.GetMajorVersionNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minorVersionNumber", m.GetMinorVersionNumber())
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
        err := writer.WriteInt32Value("updateBuildRevision", m.GetUpdateBuildRevision())
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
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBuildNumber sets the buildNumber property value. The build number of the product release, Allowed range is 0 - 2,147,483,647. For example: 19045. Read-only.
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetBuildNumber(value *int32)() {
    err := m.GetBackingStore().Set("buildNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetMajorVersionNumber sets the majorVersionNumber property value. The major version of the product release, Allowed range is 0 - 2,147,483,647. For example: 10. Read-only.
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetMajorVersionNumber(value *int32)() {
    err := m.GetBackingStore().Set("majorVersionNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetMinorVersionNumber sets the minorVersionNumber property value. The minor version of the product release, Allowed range is 0 - 2,147,483,647. For example: 0. Read-only.
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetMinorVersionNumber(value *int32)() {
    err := m.GetBackingStore().Set("minorVersionNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateBuildRevision sets the updateBuildRevision property value. The update build revision number of the product revision for the corresponding patch, Allowed range is 0 - 2,147,483,647. For example: 4780. Read-only.
func (m *WindowsQualityUpdateProductBuildVersionDetail) SetUpdateBuildRevision(value *int32)() {
    err := m.GetBackingStore().Set("updateBuildRevision", value)
    if err != nil {
        panic(err)
    }
}
type WindowsQualityUpdateProductBuildVersionDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBuildNumber()(*int32)
    GetMajorVersionNumber()(*int32)
    GetMinorVersionNumber()(*int32)
    GetOdataType()(*string)
    GetUpdateBuildRevision()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBuildNumber(value *int32)()
    SetMajorVersionNumber(value *int32)()
    SetMinorVersionNumber(value *int32)()
    SetOdataType(value *string)()
    SetUpdateBuildRevision(value *int32)()
}
