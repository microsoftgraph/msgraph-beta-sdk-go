package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppLogCollectionDownloadDetails 
type AppLogCollectionDownloadDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The appLogDecryptionAlgorithm property
    appLogDecryptionAlgorithm *AppLogDecryptionAlgorithm
    // Decryption key that used to decrypt the log.
    decryptionKey *string
    // Download SAS (Shared Access Signature) Url for completed app log request.
    downloadUrl *string
    // The OdataType property
    odataType *string
}
// NewAppLogCollectionDownloadDetails instantiates a new appLogCollectionDownloadDetails and sets the default values.
func NewAppLogCollectionDownloadDetails()(*AppLogCollectionDownloadDetails) {
    m := &AppLogCollectionDownloadDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppLogCollectionDownloadDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppLogCollectionDownloadDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppLogCollectionDownloadDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppLogCollectionDownloadDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppLogDecryptionAlgorithm gets the appLogDecryptionAlgorithm property value. The appLogDecryptionAlgorithm property
func (m *AppLogCollectionDownloadDetails) GetAppLogDecryptionAlgorithm()(*AppLogDecryptionAlgorithm) {
    return m.appLogDecryptionAlgorithm
}
// GetDecryptionKey gets the decryptionKey property value. Decryption key that used to decrypt the log.
func (m *AppLogCollectionDownloadDetails) GetDecryptionKey()(*string) {
    return m.decryptionKey
}
// GetDownloadUrl gets the downloadUrl property value. Download SAS (Shared Access Signature) Url for completed app log request.
func (m *AppLogCollectionDownloadDetails) GetDownloadUrl()(*string) {
    return m.downloadUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppLogCollectionDownloadDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appLogDecryptionAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogDecryptionAlgorithm)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppLogDecryptionAlgorithm(val.(*AppLogDecryptionAlgorithm))
        }
        return nil
    }
    res["decryptionKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecryptionKey(val)
        }
        return nil
    }
    res["downloadUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUrl(val)
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
func (m *AppLogCollectionDownloadDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppLogDecryptionAlgorithm sets the appLogDecryptionAlgorithm property value. The appLogDecryptionAlgorithm property
func (m *AppLogCollectionDownloadDetails) SetAppLogDecryptionAlgorithm(value *AppLogDecryptionAlgorithm)() {
    m.appLogDecryptionAlgorithm = value
}
// SetDecryptionKey sets the decryptionKey property value. Decryption key that used to decrypt the log.
func (m *AppLogCollectionDownloadDetails) SetDecryptionKey(value *string)() {
    m.decryptionKey = value
}
// SetDownloadUrl sets the downloadUrl property value. Download SAS (Shared Access Signature) Url for completed app log request.
func (m *AppLogCollectionDownloadDetails) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppLogCollectionDownloadDetails) SetOdataType(value *string)() {
    m.odataType = value
}
