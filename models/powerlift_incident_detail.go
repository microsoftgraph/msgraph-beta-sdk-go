package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PowerliftIncidentDetail this type contains specific information regarding a Powerlift incident, such as when it was uploaded, the platform the device was on, and a string array of files associated to the incident.
type PowerliftIncidentDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPowerliftIncidentDetail instantiates a new PowerliftIncidentDetail and sets the default values.
func NewPowerliftIncidentDetail()(*PowerliftIncidentDetail) {
    m := &PowerliftIncidentDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePowerliftIncidentDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePowerliftIncidentDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPowerliftIncidentDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PowerliftIncidentDetail) GetAdditionalData()(map[string]any) {
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
// GetApplicationName gets the applicationName property value. TThe name of the application for which the diagnostic is collected. Example: com.microsoft.CompanyPortal
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetApplicationName()(*string) {
    val, err := m.GetBackingStore().Get("applicationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *PowerliftIncidentDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClientApplicationVersion gets the clientApplicationVersion property value. The version of the application for which the diagnostic is collected. Example: 5.2203.1
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetClientApplicationVersion()(*string) {
    val, err := m.GetBackingStore().Get("clientApplicationVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The time the app diagnostic was created. The value cannot be modified and is automatically populated when the diagnostic is uploaded. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.Example: 2022-04-19T17:24:45.313Z
// returns a *Time when successful
func (m *PowerliftIncidentDetail) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEasyId gets the easyId property value. The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. This id is smaller compared to the powerliftId. Th Example: 8520467A
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetEasyId()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PowerliftIncidentDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
        }
        return nil
    }
    res["clientApplicationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientApplicationVersion(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["platformDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformDisplayName(val)
        }
        return nil
    }
    res["powerliftId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// returns a []string when successful
func (m *PowerliftIncidentDetail) GetFileNames()([]string) {
    val, err := m.GetBackingStore().Get("fileNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetLocale gets the locale property value. The locale information of the application for which the diagnostic is collected. Example: en-US
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetLocale()(*string) {
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
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatformDisplayName gets the platformDisplayName property value. The operating system of the device from which diagnostics are collected. Example: iOS
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetPlatformDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("platformDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPowerliftId gets the powerliftId property value. The unique identifier of the app diagnostic. This id is assigned to a diagnostic when it is uploaded to Powerlift. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
// returns a *string when successful
func (m *PowerliftIncidentDetail) GetPowerliftId()(*string) {
    val, err := m.GetBackingStore().Get("powerliftId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PowerliftIncidentDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationName", m.GetApplicationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientApplicationVersion", m.GetClientApplicationVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err := writer.WriteStringValue("platformDisplayName", m.GetPlatformDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("powerliftId", m.GetPowerliftId())
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
func (m *PowerliftIncidentDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationName sets the applicationName property value. TThe name of the application for which the diagnostic is collected. Example: com.microsoft.CompanyPortal
func (m *PowerliftIncidentDetail) SetApplicationName(value *string)() {
    err := m.GetBackingStore().Set("applicationName", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PowerliftIncidentDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClientApplicationVersion sets the clientApplicationVersion property value. The version of the application for which the diagnostic is collected. Example: 5.2203.1
func (m *PowerliftIncidentDetail) SetClientApplicationVersion(value *string)() {
    err := m.GetBackingStore().Set("clientApplicationVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time the app diagnostic was created. The value cannot be modified and is automatically populated when the diagnostic is uploaded. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.Example: 2022-04-19T17:24:45.313Z
func (m *PowerliftIncidentDetail) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEasyId sets the easyId property value. The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. This id is smaller compared to the powerliftId. Th Example: 8520467A
func (m *PowerliftIncidentDetail) SetEasyId(value *string)() {
    err := m.GetBackingStore().Set("easyId", value)
    if err != nil {
        panic(err)
    }
}
// SetFileNames sets the fileNames property value. A list of files that are associated with the diagnostic.
func (m *PowerliftIncidentDetail) SetFileNames(value []string)() {
    err := m.GetBackingStore().Set("fileNames", value)
    if err != nil {
        panic(err)
    }
}
// SetLocale sets the locale property value. The locale information of the application for which the diagnostic is collected. Example: en-US
func (m *PowerliftIncidentDetail) SetLocale(value *string)() {
    err := m.GetBackingStore().Set("locale", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PowerliftIncidentDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatformDisplayName sets the platformDisplayName property value. The operating system of the device from which diagnostics are collected. Example: iOS
func (m *PowerliftIncidentDetail) SetPlatformDisplayName(value *string)() {
    err := m.GetBackingStore().Set("platformDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPowerliftId sets the powerliftId property value. The unique identifier of the app diagnostic. This id is assigned to a diagnostic when it is uploaded to Powerlift. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
func (m *PowerliftIncidentDetail) SetPowerliftId(value *string)() {
    err := m.GetBackingStore().Set("powerliftId", value)
    if err != nil {
        panic(err)
    }
}
type PowerliftIncidentDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationName()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClientApplicationVersion()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasyId()(*string)
    GetFileNames()([]string)
    GetLocale()(*string)
    GetOdataType()(*string)
    GetPlatformDisplayName()(*string)
    GetPowerliftId()(*string)
    SetApplicationName(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClientApplicationVersion(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasyId(value *string)()
    SetFileNames(value []string)()
    SetLocale(value *string)()
    SetOdataType(value *string)()
    SetPlatformDisplayName(value *string)()
    SetPowerliftId(value *string)()
}
