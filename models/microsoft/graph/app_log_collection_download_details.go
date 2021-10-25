package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppLogCollectionDownloadDetails struct {
    additionalData map[string]interface{};
    appLogDecryptionAlgorithm *AppLogDecryptionAlgorithm;
    decryptionKey *string;
    downloadUrl *string;
}
func NewAppLogCollectionDownloadDetails()(*AppLogCollectionDownloadDetails) {
    m := &AppLogCollectionDownloadDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AppLogCollectionDownloadDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AppLogCollectionDownloadDetails) GetAppLogDecryptionAlgorithm()(*AppLogDecryptionAlgorithm) {
    if m == nil {
        return nil
    } else {
        return m.appLogDecryptionAlgorithm
    }
}
func (m *AppLogCollectionDownloadDetails) GetDecryptionKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decryptionKey
    }
}
func (m *AppLogCollectionDownloadDetails) GetDownloadUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.downloadUrl
    }
}
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
func (m *AppLogCollectionDownloadDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AppLogCollectionDownloadDetails) SetAppLogDecryptionAlgorithm(value *AppLogDecryptionAlgorithm)() {
    m.appLogDecryptionAlgorithm = value
}
func (m *AppLogCollectionDownloadDetails) SetDecryptionKey(value *string)() {
    m.decryptionKey = value
}
func (m *AppLogCollectionDownloadDetails) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
