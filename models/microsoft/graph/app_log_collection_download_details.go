package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new appLogCollectionDownloadDetails and sets the default values.
func NewAppLogCollectionDownloadDetails()(*AppLogCollectionDownloadDetails) {
    m := &AppLogCollectionDownloadDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppLogCollectionDownloadDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appLogDecryptionAlgorithm property value. DecryptionAlgorithm for Content. Possible values are: aes256.
func (m *AppLogCollectionDownloadDetails) GetAppLogDecryptionAlgorithm()(*AppLogDecryptionAlgorithm) {
    if m == nil {
        return nil
    } else {
        return m.appLogDecryptionAlgorithm
    }
}
// Gets the decryptionKey property value. DecryptionKey as string
func (m *AppLogCollectionDownloadDetails) GetDecryptionKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decryptionKey
    }
}
// Gets the downloadUrl property value. Download SAS Url for completed AppLogUploadRequest
func (m *AppLogCollectionDownloadDetails) GetDownloadUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.downloadUrl
    }
}
// The deserialization information for the current model
func (m *AppLogCollectionDownloadDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appLogDecryptionAlgorithm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogDecryptionAlgorithm)
        if err != nil {
            return err
        }
        cast := val.(AppLogDecryptionAlgorithm)
        m.SetAppLogDecryptionAlgorithm(&cast)
        return nil
    }
    res["decryptionKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDecryptionKey(val)
        return nil
    }
    res["downloadUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDownloadUrl(val)
        return nil
    }
    return res
}
func (m *AppLogCollectionDownloadDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppLogCollectionDownloadDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAppLogDecryptionAlgorithm() != nil {
        cast := m.GetAppLogDecryptionAlgorithm().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AppLogCollectionDownloadDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appLogDecryptionAlgorithm property value. DecryptionAlgorithm for Content. Possible values are: aes256.
// Parameters:
//  - value : Value to set for the appLogDecryptionAlgorithm property.
func (m *AppLogCollectionDownloadDetails) SetAppLogDecryptionAlgorithm(value *AppLogDecryptionAlgorithm)() {
    m.appLogDecryptionAlgorithm = value
}
// Sets the decryptionKey property value. DecryptionKey as string
// Parameters:
//  - value : Value to set for the decryptionKey property.
func (m *AppLogCollectionDownloadDetails) SetDecryptionKey(value *string)() {
    m.decryptionKey = value
}
// Sets the downloadUrl property value. Download SAS Url for completed AppLogUploadRequest
// Parameters:
//  - value : Value to set for the downloadUrl property.
func (m *AppLogCollectionDownloadDetails) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
