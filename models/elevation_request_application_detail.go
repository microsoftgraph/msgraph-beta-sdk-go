package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ElevationRequestApplicationDetail the details of the application which the user has requested to elevate
type ElevationRequestApplicationDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewElevationRequestApplicationDetail instantiates a new ElevationRequestApplicationDetail and sets the default values.
func NewElevationRequestApplicationDetail()(*ElevationRequestApplicationDetail) {
    m := &ElevationRequestApplicationDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateElevationRequestApplicationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateElevationRequestApplicationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewElevationRequestApplicationDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ElevationRequestApplicationDetail) GetAdditionalData()(map[string]any) {
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
func (m *ElevationRequestApplicationDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ElevationRequestApplicationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["fileHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHash(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
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
    res["productInternalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductInternalName(val)
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
    res["productVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersion(val)
        }
        return nil
    }
    res["publisherCert"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherCert(val)
        }
        return nil
    }
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    return res
}
// GetFileDescription gets the fileDescription property value. The path of the file in the request for elevation, for example, %programfiles%/git/cmd
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetFileDescription()(*string) {
    val, err := m.GetBackingStore().Get("fileDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileHash gets the fileHash property value. The SHA256 hash of the file in the request for elevation, for example, '18ee24150dcb1d96752a4d6dd0f20dfd8ba8c38527e40aa8509b7adecf78f9c6'
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetFileHash()(*string) {
    val, err := m.GetBackingStore().Get("fileHash")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileName gets the fileName property value. The name of the file in the request for elevation, for example, git.exe
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetFileName()(*string) {
    val, err := m.GetBackingStore().Get("fileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFilePath gets the filePath property value. The path of the file in the request for elevation, for example, %programfiles%/git/cmd
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetFilePath()(*string) {
    val, err := m.GetBackingStore().Get("filePath")
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
func (m *ElevationRequestApplicationDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductInternalName gets the productInternalName property value. The internal name of the application for which elevation request has been made. For example, 'git'
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetProductInternalName()(*string) {
    val, err := m.GetBackingStore().Get("productInternalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductName gets the productName property value. The product name of the application for which elevation request has been made. For example, 'Git'
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetProductName()(*string) {
    val, err := m.GetBackingStore().Get("productName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductVersion gets the productVersion property value. The product version of the application for which elevation request has been made. For example, '2.40.1.0'
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetProductVersion()(*string) {
    val, err := m.GetBackingStore().Get("productVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisherCert gets the publisherCert property value. The list of base64 encoded certificate for each signer, for example, string[encodedleafcert1, encodedleafcert2....]
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetPublisherCert()(*string) {
    val, err := m.GetBackingStore().Get("publisherCert")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisherName gets the publisherName property value. The certificate issuer name of the certificate used to sign the application, for example, 'Sectigo Public Code Signing CA R36'
// returns a *string when successful
func (m *ElevationRequestApplicationDetail) GetPublisherName()(*string) {
    val, err := m.GetBackingStore().Get("publisherName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ElevationRequestApplicationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileDescription", m.GetFileDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileHash", m.GetFileHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePath", m.GetFilePath())
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
        err := writer.WriteStringValue("productInternalName", m.GetProductInternalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productVersion", m.GetProductVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherCert", m.GetPublisherCert())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
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
func (m *ElevationRequestApplicationDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ElevationRequestApplicationDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFileDescription sets the fileDescription property value. The path of the file in the request for elevation, for example, %programfiles%/git/cmd
func (m *ElevationRequestApplicationDetail) SetFileDescription(value *string)() {
    err := m.GetBackingStore().Set("fileDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetFileHash sets the fileHash property value. The SHA256 hash of the file in the request for elevation, for example, '18ee24150dcb1d96752a4d6dd0f20dfd8ba8c38527e40aa8509b7adecf78f9c6'
func (m *ElevationRequestApplicationDetail) SetFileHash(value *string)() {
    err := m.GetBackingStore().Set("fileHash", value)
    if err != nil {
        panic(err)
    }
}
// SetFileName sets the fileName property value. The name of the file in the request for elevation, for example, git.exe
func (m *ElevationRequestApplicationDetail) SetFileName(value *string)() {
    err := m.GetBackingStore().Set("fileName", value)
    if err != nil {
        panic(err)
    }
}
// SetFilePath sets the filePath property value. The path of the file in the request for elevation, for example, %programfiles%/git/cmd
func (m *ElevationRequestApplicationDetail) SetFilePath(value *string)() {
    err := m.GetBackingStore().Set("filePath", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ElevationRequestApplicationDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProductInternalName sets the productInternalName property value. The internal name of the application for which elevation request has been made. For example, 'git'
func (m *ElevationRequestApplicationDetail) SetProductInternalName(value *string)() {
    err := m.GetBackingStore().Set("productInternalName", value)
    if err != nil {
        panic(err)
    }
}
// SetProductName sets the productName property value. The product name of the application for which elevation request has been made. For example, 'Git'
func (m *ElevationRequestApplicationDetail) SetProductName(value *string)() {
    err := m.GetBackingStore().Set("productName", value)
    if err != nil {
        panic(err)
    }
}
// SetProductVersion sets the productVersion property value. The product version of the application for which elevation request has been made. For example, '2.40.1.0'
func (m *ElevationRequestApplicationDetail) SetProductVersion(value *string)() {
    err := m.GetBackingStore().Set("productVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisherCert sets the publisherCert property value. The list of base64 encoded certificate for each signer, for example, string[encodedleafcert1, encodedleafcert2....]
func (m *ElevationRequestApplicationDetail) SetPublisherCert(value *string)() {
    err := m.GetBackingStore().Set("publisherCert", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisherName sets the publisherName property value. The certificate issuer name of the certificate used to sign the application, for example, 'Sectigo Public Code Signing CA R36'
func (m *ElevationRequestApplicationDetail) SetPublisherName(value *string)() {
    err := m.GetBackingStore().Set("publisherName", value)
    if err != nil {
        panic(err)
    }
}
type ElevationRequestApplicationDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFileDescription()(*string)
    GetFileHash()(*string)
    GetFileName()(*string)
    GetFilePath()(*string)
    GetOdataType()(*string)
    GetProductInternalName()(*string)
    GetProductName()(*string)
    GetProductVersion()(*string)
    GetPublisherCert()(*string)
    GetPublisherName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFileDescription(value *string)()
    SetFileHash(value *string)()
    SetFileName(value *string)()
    SetFilePath(value *string)()
    SetOdataType(value *string)()
    SetProductInternalName(value *string)()
    SetProductName(value *string)()
    SetProductVersion(value *string)()
    SetPublisherCert(value *string)()
    SetPublisherName(value *string)()
}
