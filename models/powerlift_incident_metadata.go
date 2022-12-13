package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PowerliftIncidentMetadata collection of app diagnostics associated with a user.
type PowerliftIncidentMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the application the diagnostic is from. Example: com.microsoft.CompanyPortal
    application *string
    // The version of the application. Example: 5.2203.1
    clientVersion *string
    // The time the app diagnostic was created. Example: 2022-04-19T17:24:45.313Z
    createdAtDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. Example: 8520467A
    easyId *string
    // A list of files that are associated with the diagnostic.
    fileNames []string
    // The locale information of the application. Example: en-US
    locale *string
    // The OdataType property
    odataType *string
    // The device's OS the diagnostic is from. Example: iOS
    platform *string
    // The unique identifier of the app diagnostic. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
    powerliftId *UUID
}
// NewPowerliftIncidentMetadata instantiates a new powerliftIncidentMetadata and sets the default values.
func NewPowerliftIncidentMetadata()(*PowerliftIncidentMetadata) {
    m := &PowerliftIncidentMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePowerliftIncidentMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePowerliftIncidentMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPowerliftIncidentMetadata(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PowerliftIncidentMetadata) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplication gets the application property value. The name of the application the diagnostic is from. Example: com.microsoft.CompanyPortal
func (m *PowerliftIncidentMetadata) GetApplication()(*string) {
    return m.application
}
// GetClientVersion gets the clientVersion property value. The version of the application. Example: 5.2203.1
func (m *PowerliftIncidentMetadata) GetClientVersion()(*string) {
    return m.clientVersion
}
// GetCreatedAtDateTime gets the createdAtDateTime property value. The time the app diagnostic was created. Example: 2022-04-19T17:24:45.313Z
func (m *PowerliftIncidentMetadata) GetCreatedAtDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAtDateTime
}
// GetEasyId gets the easyId property value. The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. Example: 8520467A
func (m *PowerliftIncidentMetadata) GetEasyId()(*string) {
    return m.easyId
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
                res[i] = *(v.(*string))
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
    return m.fileNames
}
// GetLocale gets the locale property value. The locale information of the application. Example: en-US
func (m *PowerliftIncidentMetadata) GetLocale()(*string) {
    return m.locale
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PowerliftIncidentMetadata) GetOdataType()(*string) {
    return m.odataType
}
// GetPlatform gets the platform property value. The device's OS the diagnostic is from. Example: iOS
func (m *PowerliftIncidentMetadata) GetPlatform()(*string) {
    return m.platform
}
// GetPowerliftId gets the powerliftId property value. The unique identifier of the app diagnostic. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
func (m *PowerliftIncidentMetadata) GetPowerliftId()(*UUID) {
    return m.powerliftId
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PowerliftIncidentMetadata) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplication sets the application property value. The name of the application the diagnostic is from. Example: com.microsoft.CompanyPortal
func (m *PowerliftIncidentMetadata) SetApplication(value *string)() {
    m.application = value
}
// SetClientVersion sets the clientVersion property value. The version of the application. Example: 5.2203.1
func (m *PowerliftIncidentMetadata) SetClientVersion(value *string)() {
    m.clientVersion = value
}
// SetCreatedAtDateTime sets the createdAtDateTime property value. The time the app diagnostic was created. Example: 2022-04-19T17:24:45.313Z
func (m *PowerliftIncidentMetadata) SetCreatedAtDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAtDateTime = value
}
// SetEasyId sets the easyId property value. The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. Example: 8520467A
func (m *PowerliftIncidentMetadata) SetEasyId(value *string)() {
    m.easyId = value
}
// SetFileNames sets the fileNames property value. A list of files that are associated with the diagnostic.
func (m *PowerliftIncidentMetadata) SetFileNames(value []string)() {
    m.fileNames = value
}
// SetLocale sets the locale property value. The locale information of the application. Example: en-US
func (m *PowerliftIncidentMetadata) SetLocale(value *string)() {
    m.locale = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PowerliftIncidentMetadata) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlatform sets the platform property value. The device's OS the diagnostic is from. Example: iOS
func (m *PowerliftIncidentMetadata) SetPlatform(value *string)() {
    m.platform = value
}
// SetPowerliftId sets the powerliftId property value. The unique identifier of the app diagnostic. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
func (m *PowerliftIncidentMetadata) SetPowerliftId(value *UUID)() {
    m.powerliftId = value
}
