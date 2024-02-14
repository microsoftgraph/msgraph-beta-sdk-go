package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type BuildVersionDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBuildVersionDetails instantiates a new BuildVersionDetails and sets the default values.
func NewBuildVersionDetails()(*BuildVersionDetails) {
    m := &BuildVersionDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBuildVersionDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBuildVersionDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBuildVersionDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BuildVersionDetails) GetAdditionalData()(map[string]any) {
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
func (m *BuildVersionDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBuildNumber gets the buildNumber property value. The build number of the product release. Read-only.
// returns a *int32 when successful
func (m *BuildVersionDetails) GetBuildNumber()(*int32) {
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
func (m *BuildVersionDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["majorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorVersion(val)
        }
        return nil
    }
    res["minorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorVersion(val)
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
// GetMajorVersion gets the majorVersion property value. The major version of the product release. Read-only.
// returns a *int32 when successful
func (m *BuildVersionDetails) GetMajorVersion()(*int32) {
    val, err := m.GetBackingStore().Get("majorVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinorVersion gets the minorVersion property value. The minor version of the product release. Read-only.
// returns a *int32 when successful
func (m *BuildVersionDetails) GetMinorVersion()(*int32) {
    val, err := m.GetBackingStore().Get("minorVersion")
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
func (m *BuildVersionDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUpdateBuildRevision gets the updateBuildRevision property value. The update build revision number of the product revision. Read-only.
// returns a *int32 when successful
func (m *BuildVersionDetails) GetUpdateBuildRevision()(*int32) {
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
func (m *BuildVersionDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("buildNumber", m.GetBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("majorVersion", m.GetMajorVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minorVersion", m.GetMinorVersion())
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
func (m *BuildVersionDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BuildVersionDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBuildNumber sets the buildNumber property value. The build number of the product release. Read-only.
func (m *BuildVersionDetails) SetBuildNumber(value *int32)() {
    err := m.GetBackingStore().Set("buildNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetMajorVersion sets the majorVersion property value. The major version of the product release. Read-only.
func (m *BuildVersionDetails) SetMajorVersion(value *int32)() {
    err := m.GetBackingStore().Set("majorVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinorVersion sets the minorVersion property value. The minor version of the product release. Read-only.
func (m *BuildVersionDetails) SetMinorVersion(value *int32)() {
    err := m.GetBackingStore().Set("minorVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BuildVersionDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateBuildRevision sets the updateBuildRevision property value. The update build revision number of the product revision. Read-only.
func (m *BuildVersionDetails) SetUpdateBuildRevision(value *int32)() {
    err := m.GetBackingStore().Set("updateBuildRevision", value)
    if err != nil {
        panic(err)
    }
}
type BuildVersionDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBuildNumber()(*int32)
    GetMajorVersion()(*int32)
    GetMinorVersion()(*int32)
    GetOdataType()(*string)
    GetUpdateBuildRevision()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBuildNumber(value *int32)()
    SetMajorVersion(value *int32)()
    SetMinorVersion(value *int32)()
    SetOdataType(value *string)()
    SetUpdateBuildRevision(value *int32)()
}
