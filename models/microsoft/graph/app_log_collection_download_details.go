package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppLogCollectionDownloadDetails 
type AppLogCollectionDownloadDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // DecryptionAlgorithm for Content. Possible values are: aes256.
    appLogDecryptionAlgorithm *AppLogDecryptionAlgorithm;
    // DecryptionKey as string
    decryptionKey *string;
    // Download SAS Url for completed AppLogUploadRequest
    downloadUrl *string;
}
// NewAppLogCollectionDownloadDetails instantiates a new appLogCollectionDownloadDetails and sets the default values.
func NewAppLogCollectionDownloadDetails()(*AppLogCollectionDownloadDetails) {
    m := &AppLogCollectionDownloadDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppLogCollectionDownloadDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppLogDecryptionAlgorithm gets the appLogDecryptionAlgorithm property value. DecryptionAlgorithm for Content. Possible values are: aes256.
func (m *AppLogCollectionDownloadDetails) GetAppLogDecryptionAlgorithm()(*AppLogDecryptionAlgorithm) {
    if m == nil {
        return nil
    } else {
        return m.appLogDecryptionAlgorithm
    }
}
// GetDecryptionKey gets the decryptionKey property value. DecryptionKey as string
func (m *AppLogCollectionDownloadDetails) GetDecryptionKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decryptionKey
    }
}
// GetDownloadUrl gets the downloadUrl property value. Download SAS Url for completed AppLogUploadRequest
func (m *AppLogCollectionDownloadDetails) GetDownloadUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.downloadUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppLogCollectionDownloadDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appLogDecryptionAlgorithm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogDecryptionAlgorithm)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppLogDecryptionAlgorithm(val.(*AppLogDecryptionAlgorithm))
        }
        return nil
    }
    res["decryptionKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecryptionKey(val)
        }
        return nil
    }
    res["downloadUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUrl(val)
        }
        return nil
    }
    return res
}
func (m *AppLogCollectionDownloadDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppLogCollectionDownloadDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppLogCollectionDownloadDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppLogDecryptionAlgorithm sets the appLogDecryptionAlgorithm property value. DecryptionAlgorithm for Content. Possible values are: aes256.
func (m *AppLogCollectionDownloadDetails) SetAppLogDecryptionAlgorithm(value *AppLogDecryptionAlgorithm)() {
    if m != nil {
        m.appLogDecryptionAlgorithm = value
    }
}
// SetDecryptionKey sets the decryptionKey property value. DecryptionKey as string
func (m *AppLogCollectionDownloadDetails) SetDecryptionKey(value *string)() {
    if m != nil {
        m.decryptionKey = value
    }
}
// SetDownloadUrl sets the downloadUrl property value. Download SAS Url for completed AppLogUploadRequest
func (m *AppLogCollectionDownloadDetails) SetDownloadUrl(value *string)() {
    if m != nil {
        m.downloadUrl = value
    }
}
