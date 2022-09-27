package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppLogCollectionDownloadDetails 
type AppLogCollectionDownloadDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appLogDecryptionAlgorithm property
    appLogDecryptionAlgorithm *AppLogDecryptionAlgorithm
    // DecryptionKey as string
    decryptionKey *string
    // Download SAS Url for completed AppLogUploadRequest
    downloadUrl *string
    // The OdataType property
    odataType *string
}
// NewAppLogCollectionDownloadDetails instantiates a new appLogCollectionDownloadDetails and sets the default values.
func NewAppLogCollectionDownloadDetails()(*AppLogCollectionDownloadDetails) {
    m := &AppLogCollectionDownloadDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.appLogCollectionDownloadDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppLogCollectionDownloadDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppLogCollectionDownloadDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppLogCollectionDownloadDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppLogCollectionDownloadDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppLogDecryptionAlgorithm gets the appLogDecryptionAlgorithm property value. The appLogDecryptionAlgorithm property
func (m *AppLogCollectionDownloadDetails) GetAppLogDecryptionAlgorithm()(*AppLogDecryptionAlgorithm) {
    return m.appLogDecryptionAlgorithm
}
// GetDecryptionKey gets the decryptionKey property value. DecryptionKey as string
func (m *AppLogCollectionDownloadDetails) GetDecryptionKey()(*string) {
    return m.decryptionKey
}
// GetDownloadUrl gets the downloadUrl property value. Download SAS Url for completed AppLogUploadRequest
func (m *AppLogCollectionDownloadDetails) GetDownloadUrl()(*string) {
    return m.downloadUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppLogCollectionDownloadDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appLogDecryptionAlgorithm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAppLogDecryptionAlgorithm , m.SetAppLogDecryptionAlgorithm)
    res["decryptionKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDecryptionKey)
    res["downloadUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDownloadUrl)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppLogCollectionDownloadDetails) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AppLogCollectionDownloadDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppLogDecryptionAlgorithm() != nil {
        cast := (*m.GetAppLogDecryptionAlgorithm()).String()
        err := writer.WriteStringValue("appLogDecryptionAlgorithm", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("decryptionKey", m.GetDecryptionKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("downloadUrl", m.GetDownloadUrl())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppLogCollectionDownloadDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppLogDecryptionAlgorithm sets the appLogDecryptionAlgorithm property value. The appLogDecryptionAlgorithm property
func (m *AppLogCollectionDownloadDetails) SetAppLogDecryptionAlgorithm(value *AppLogDecryptionAlgorithm)() {
    m.appLogDecryptionAlgorithm = value
}
// SetDecryptionKey sets the decryptionKey property value. DecryptionKey as string
func (m *AppLogCollectionDownloadDetails) SetDecryptionKey(value *string)() {
    m.decryptionKey = value
}
// SetDownloadUrl sets the downloadUrl property value. Download SAS Url for completed AppLogUploadRequest
func (m *AppLogCollectionDownloadDetails) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppLogCollectionDownloadDetails) SetOdataType(value *string)() {
    m.odataType = value
}
