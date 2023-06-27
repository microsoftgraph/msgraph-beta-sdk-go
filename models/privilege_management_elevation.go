package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegeManagementElevation 
type PrivilegeManagementElevation struct {
    Entity
}
// NewPrivilegeManagementElevation instantiates a new PrivilegeManagementElevation and sets the default values.
func NewPrivilegeManagementElevation()(*PrivilegeManagementElevation) {
    m := &PrivilegeManagementElevation{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegeManagementElevationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegeManagementElevationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeManagementElevation(), nil
}
// GetCertificatePayload gets the certificatePayload property value. The certificate payload of the application. This is computed by hashing the certificate information on the client. Example: `32c220482c68413fbf8290e3b1e49b0a85901cfcd62ab0738760568a2a6e8a50`
func (m *PrivilegeManagementElevation) GetCertificatePayload()(*string) {
    val, err := m.GetBackingStore().Get("certificatePayload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCompanyName gets the companyName property value. The company name of the application. This value is set by the creator of the application. Example: `Microsoft Corporation`
func (m *PrivilegeManagementElevation) GetCompanyName()(*string) {
    val, err := m.GetBackingStore().Get("companyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The Intune deviceId. Unique identifier for the managed device. Example: `92ce5047-9553-4731-817f-9b401a999a1b`
func (m *PrivilegeManagementElevation) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name associated with the device in the intune database. Example: `JOHNDOE-LAPTOP`.
func (m *PrivilegeManagementElevation) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetElevationType gets the elevationType property value. Indicates the type of elevation occured
func (m *PrivilegeManagementElevation) GetElevationType()(*PrivilegeManagementElevationType) {
    val, err := m.GetBackingStore().Get("elevationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrivilegeManagementElevationType)
    }
    return nil
}
// GetEventDateTime gets the eventDateTime property value. The date and time when the application was elevated. Example:`2014-01-01T00:00:00Z`
func (m *PrivilegeManagementElevation) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegeManagementElevation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificatePayload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificatePayload(val)
        }
        return nil
    }
    res["companyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["elevationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegeManagementElevationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElevationType(val.(*PrivilegeManagementElevationType))
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["fileDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileDescription(val)
        }
        return nil
    }
    res["filePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
        }
        return nil
    }
    res["fileVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVersion(val)
        }
        return nil
    }
    res["hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHash(val)
        }
        return nil
    }
    res["internalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalName(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val)
        }
        return nil
    }
    res["upn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegeManagementEndUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*PrivilegeManagementEndUserType))
        }
        return nil
    }
    return res
}
// GetFileDescription gets the fileDescription property value. The file description of the application. This value is set by the creator of the application. Example: `Editor of multiple coding languages.`
func (m *PrivilegeManagementElevation) GetFileDescription()(*string) {
    val, err := m.GetBackingStore().Get("fileDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFilePath gets the filePath property value. The full file path of the application including the filename and file extension. Example: `C:\Program Files\vscode.exe`
func (m *PrivilegeManagementElevation) GetFilePath()(*string) {
    val, err := m.GetBackingStore().Get("filePath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileVersion gets the fileVersion property value. The version of the application. This value is set by the creator of the application. Example: `6.2211.1035.1000`
func (m *PrivilegeManagementElevation) GetFileVersion()(*string) {
    val, err := m.GetBackingStore().Get("fileVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHash gets the hash property value. The sha256 hash of the application. Example: `32c220482c68413fbf8290e3b1e49b0a85901cfcd62ab0738760568a2a6e8a57`
func (m *PrivilegeManagementElevation) GetHash()(*string) {
    val, err := m.GetBackingStore().Get("hash")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInternalName gets the internalName property value. The internal name of the application. This value is set by the creator of the application. Example: `VS code`
func (m *PrivilegeManagementElevation) GetInternalName()(*string) {
    val, err := m.GetBackingStore().Get("internalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetJustification gets the justification property value. The justification to elevate the application. This is an input by the user when the privilegeManagementElevationType is of type userConfirmedElevation or support approved elevation. This will be null in all other scenarios. The length is capped at 256 char, enforced on the client side. Example: `To install debug tool.`.
func (m *PrivilegeManagementElevation) GetJustification()(*string) {
    val, err := m.GetBackingStore().Get("justification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductName gets the productName property value. The product name of the application. This value is set by the creator of the application. Example: `Visual Studio`
func (m *PrivilegeManagementElevation) GetProductName()(*string) {
    val, err := m.GetBackingStore().Get("productName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResult gets the result property value. The result of the elevation action with 0 being success, and everything else being exit code if the elevation was unsuccessful. The value will always be 0 on all unmanaged elevation. Example: `0`. Valid values 0 to 2147483647
func (m *PrivilegeManagementElevation) GetResult()(*int32) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUpn gets the upn property value. The User Principal Name of the user who performed the elevation. Example: `john@domain.com`
func (m *PrivilegeManagementElevation) GetUpn()(*string) {
    val, err := m.GetBackingStore().Get("upn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserType gets the userType property value. The type of user account on Windows that was used to performed the elevation.
func (m *PrivilegeManagementElevation) GetUserType()(*PrivilegeManagementEndUserType) {
    val, err := m.GetBackingStore().Get("userType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrivilegeManagementEndUserType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegeManagementElevation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificatePayload", m.GetCertificatePayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetElevationType() != nil {
        cast := (*m.GetElevationType()).String()
        err = writer.WriteStringValue("elevationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileDescription", m.GetFileDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filePath", m.GetFilePath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileVersion", m.GetFileVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hash", m.GetHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internalName", m.GetInternalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("result", m.GetResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upn", m.GetUpn())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificatePayload sets the certificatePayload property value. The certificate payload of the application. This is computed by hashing the certificate information on the client. Example: `32c220482c68413fbf8290e3b1e49b0a85901cfcd62ab0738760568a2a6e8a50`
func (m *PrivilegeManagementElevation) SetCertificatePayload(value *string)() {
    err := m.GetBackingStore().Set("certificatePayload", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyName sets the companyName property value. The company name of the application. This value is set by the creator of the application. Example: `Microsoft Corporation`
func (m *PrivilegeManagementElevation) SetCompanyName(value *string)() {
    err := m.GetBackingStore().Set("companyName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The Intune deviceId. Unique identifier for the managed device. Example: `92ce5047-9553-4731-817f-9b401a999a1b`
func (m *PrivilegeManagementElevation) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name associated with the device in the intune database. Example: `JOHNDOE-LAPTOP`.
func (m *PrivilegeManagementElevation) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetElevationType sets the elevationType property value. Indicates the type of elevation occured
func (m *PrivilegeManagementElevation) SetElevationType(value *PrivilegeManagementElevationType)() {
    err := m.GetBackingStore().Set("elevationType", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDateTime sets the eventDateTime property value. The date and time when the application was elevated. Example:`2014-01-01T00:00:00Z`
func (m *PrivilegeManagementElevation) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFileDescription sets the fileDescription property value. The file description of the application. This value is set by the creator of the application. Example: `Editor of multiple coding languages.`
func (m *PrivilegeManagementElevation) SetFileDescription(value *string)() {
    err := m.GetBackingStore().Set("fileDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetFilePath sets the filePath property value. The full file path of the application including the filename and file extension. Example: `C:\Program Files\vscode.exe`
func (m *PrivilegeManagementElevation) SetFilePath(value *string)() {
    err := m.GetBackingStore().Set("filePath", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVersion sets the fileVersion property value. The version of the application. This value is set by the creator of the application. Example: `6.2211.1035.1000`
func (m *PrivilegeManagementElevation) SetFileVersion(value *string)() {
    err := m.GetBackingStore().Set("fileVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetHash sets the hash property value. The sha256 hash of the application. Example: `32c220482c68413fbf8290e3b1e49b0a85901cfcd62ab0738760568a2a6e8a57`
func (m *PrivilegeManagementElevation) SetHash(value *string)() {
    err := m.GetBackingStore().Set("hash", value)
    if err != nil {
        panic(err)
    }
}
// SetInternalName sets the internalName property value. The internal name of the application. This value is set by the creator of the application. Example: `VS code`
func (m *PrivilegeManagementElevation) SetInternalName(value *string)() {
    err := m.GetBackingStore().Set("internalName", value)
    if err != nil {
        panic(err)
    }
}
// SetJustification sets the justification property value. The justification to elevate the application. This is an input by the user when the privilegeManagementElevationType is of type userConfirmedElevation or support approved elevation. This will be null in all other scenarios. The length is capped at 256 char, enforced on the client side. Example: `To install debug tool.`.
func (m *PrivilegeManagementElevation) SetJustification(value *string)() {
    err := m.GetBackingStore().Set("justification", value)
    if err != nil {
        panic(err)
    }
}
// SetProductName sets the productName property value. The product name of the application. This value is set by the creator of the application. Example: `Visual Studio`
func (m *PrivilegeManagementElevation) SetProductName(value *string)() {
    err := m.GetBackingStore().Set("productName", value)
    if err != nil {
        panic(err)
    }
}
// SetResult sets the result property value. The result of the elevation action with 0 being success, and everything else being exit code if the elevation was unsuccessful. The value will always be 0 on all unmanaged elevation. Example: `0`. Valid values 0 to 2147483647
func (m *PrivilegeManagementElevation) SetResult(value *int32)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
// SetUpn sets the upn property value. The User Principal Name of the user who performed the elevation. Example: `john@domain.com`
func (m *PrivilegeManagementElevation) SetUpn(value *string)() {
    err := m.GetBackingStore().Set("upn", value)
    if err != nil {
        panic(err)
    }
}
// SetUserType sets the userType property value. The type of user account on Windows that was used to performed the elevation.
func (m *PrivilegeManagementElevation) SetUserType(value *PrivilegeManagementEndUserType)() {
    err := m.GetBackingStore().Set("userType", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegeManagementElevationable 
type PrivilegeManagementElevationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificatePayload()(*string)
    GetCompanyName()(*string)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetElevationType()(*PrivilegeManagementElevationType)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFileDescription()(*string)
    GetFilePath()(*string)
    GetFileVersion()(*string)
    GetHash()(*string)
    GetInternalName()(*string)
    GetJustification()(*string)
    GetProductName()(*string)
    GetResult()(*int32)
    GetUpn()(*string)
    GetUserType()(*PrivilegeManagementEndUserType)
    SetCertificatePayload(value *string)()
    SetCompanyName(value *string)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetElevationType(value *PrivilegeManagementElevationType)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFileDescription(value *string)()
    SetFilePath(value *string)()
    SetFileVersion(value *string)()
    SetHash(value *string)()
    SetInternalName(value *string)()
    SetJustification(value *string)()
    SetProductName(value *string)()
    SetResult(value *int32)()
    SetUpn(value *string)()
    SetUserType(value *PrivilegeManagementEndUserType)()
}
