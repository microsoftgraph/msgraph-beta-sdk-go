package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PowerliftIncidentMetadata collection of app diagnostics associated with a user.
type PowerliftIncidentMetadata struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPowerliftIncidentMetadata instantiates a new powerliftIncidentMetadata and sets the default values.
func NewPowerliftIncidentMetadata()(*PowerliftIncidentMetadata) {
    m := &PowerliftIncidentMetadata{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePowerliftIncidentMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePowerliftIncidentMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPowerliftIncidentMetadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PowerliftIncidentMetadata) GetAdditionalData()(map[string]any) {
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
// GetApplication gets the application property value. The name of the application the diagnostic is from. Example: com.microsoft.CompanyPortal
func (m *PowerliftIncidentMetadata) GetApplication()(*string) {
    val, err := m.GetBackingStore().Get("application")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *PowerliftIncidentMetadata) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClientVersion gets the clientVersion property value. The version of the application. Example: 5.2203.1
func (m *PowerliftIncidentMetadata) GetClientVersion()(*string) {
    val, err := m.GetBackingStore().Get("clientVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedAtDateTime gets the createdAtDateTime property value. The time the app diagnostic was created. Example: 2022-04-19T17:24:45.313Z
func (m *PowerliftIncidentMetadata) GetCreatedAtDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdAtDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEasyId gets the easyId property value. The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. Example: 8520467A
func (m *PowerliftIncidentMetadata) GetEasyId()(*string) {
    val, err := m.GetBackingStore().Get("easyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PowerliftIncidentMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val)
        }
        return nil
    }
    res["clientVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientVersion(val)
        }
        return nil
    }
    res["createdAtDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAtDateTime(val)
        }
        return nil
    }
    res["easyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasyId(val)
        }
        return nil
    }
    res["fileNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetFileNames(res)
        }
        return nil
    }
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val)
        }
        return nil
    }
    res["powerliftId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerliftId(val)
        }
        return nil
    }
    return res
}
// GetFileNames gets the fileNames property value. A list of files that are associated with the diagnostic.
func (m *PowerliftIncidentMetadata) GetFileNames()([]string) {
    val, err := m.GetBackingStore().Get("fileNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetLocale gets the locale property value. The locale information of the application. Example: en-US
func (m *PowerliftIncidentMetadata) GetLocale()(*string) {
    val, err := m.GetBackingStore().Get("locale")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PowerliftIncidentMetadata) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatform gets the platform property value. The device's OS the diagnostic is from. Example: iOS
func (m *PowerliftIncidentMetadata) GetPlatform()(*string) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPowerliftId gets the powerliftId property value. The unique identifier of the app diagnostic. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
func (m *PowerliftIncidentMetadata) GetPowerliftId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("powerliftId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PowerliftIncidentMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientVersion", m.GetClientVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdAtDateTime", m.GetCreatedAtDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("easyId", m.GetEasyId())
        if err != nil {
            return err
        }
    }
    if m.GetFileNames() != nil {
        err := writer.WriteCollectionOfStringValues("fileNames", m.GetFileNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
        err := writer.WriteStringValue("platform", m.GetPlatform())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("powerliftId", m.GetPowerliftId())
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
func (m *PowerliftIncidentMetadata) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApplication sets the application property value. The name of the application the diagnostic is from. Example: com.microsoft.CompanyPortal
func (m *PowerliftIncidentMetadata) SetApplication(value *string)() {
    err := m.GetBackingStore().Set("application", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PowerliftIncidentMetadata) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClientVersion sets the clientVersion property value. The version of the application. Example: 5.2203.1
func (m *PowerliftIncidentMetadata) SetClientVersion(value *string)() {
    err := m.GetBackingStore().Set("clientVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedAtDateTime sets the createdAtDateTime property value. The time the app diagnostic was created. Example: 2022-04-19T17:24:45.313Z
func (m *PowerliftIncidentMetadata) SetCreatedAtDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdAtDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEasyId sets the easyId property value. The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. Example: 8520467A
func (m *PowerliftIncidentMetadata) SetEasyId(value *string)() {
    err := m.GetBackingStore().Set("easyId", value)
    if err != nil {
        panic(err)
    }
}
// SetFileNames sets the fileNames property value. A list of files that are associated with the diagnostic.
func (m *PowerliftIncidentMetadata) SetFileNames(value []string)() {
    err := m.GetBackingStore().Set("fileNames", value)
    if err != nil {
        panic(err)
    }
}
// SetLocale sets the locale property value. The locale information of the application. Example: en-US
func (m *PowerliftIncidentMetadata) SetLocale(value *string)() {
    err := m.GetBackingStore().Set("locale", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PowerliftIncidentMetadata) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. The device's OS the diagnostic is from. Example: iOS
func (m *PowerliftIncidentMetadata) SetPlatform(value *string)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetPowerliftId sets the powerliftId property value. The unique identifier of the app diagnostic. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
func (m *PowerliftIncidentMetadata) SetPowerliftId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("powerliftId", value)
    if err != nil {
        panic(err)
    }
}
// PowerliftIncidentMetadataable 
type PowerliftIncidentMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClientVersion()(*string)
    GetCreatedAtDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasyId()(*string)
    GetFileNames()([]string)
    GetLocale()(*string)
    GetOdataType()(*string)
    GetPlatform()(*string)
    GetPowerliftId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetApplication(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClientVersion(value *string)()
    SetCreatedAtDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasyId(value *string)()
    SetFileNames(value []string)()
    SetLocale(value *string)()
    SetOdataType(value *string)()
    SetPlatform(value *string)()
    SetPowerliftId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
