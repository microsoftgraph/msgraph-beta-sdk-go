package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SigningCertificateUpdateStatus 
type SigningCertificateUpdateStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    certificateUpdateResult *string;
    // 
    lastRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewSigningCertificateUpdateStatus instantiates a new signingCertificateUpdateStatus and sets the default values.
func NewSigningCertificateUpdateStatus()(*SigningCertificateUpdateStatus) {
    m := &SigningCertificateUpdateStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SigningCertificateUpdateStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificateUpdateResult gets the certificateUpdateResult property value. 
func (m *SigningCertificateUpdateStatus) GetCertificateUpdateResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateUpdateResult
    }
}
// GetLastRunDateTime gets the lastRunDateTime property value. 
func (m *SigningCertificateUpdateStatus) GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRunDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SigningCertificateUpdateStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificateUpdateResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateUpdateResult(val)
        }
        return nil
    }
    res["lastRunDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRunDateTime(val)
        }
        return nil
    }
    return res
}
func (m *SigningCertificateUpdateStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SigningCertificateUpdateStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("certificateUpdateResult", m.GetCertificateUpdateResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastRunDateTime", m.GetLastRunDateTime())
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
func (m *SigningCertificateUpdateStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificateUpdateResult sets the certificateUpdateResult property value. 
func (m *SigningCertificateUpdateStatus) SetCertificateUpdateResult(value *string)() {
    if m != nil {
        m.certificateUpdateResult = value
    }
}
// SetLastRunDateTime sets the lastRunDateTime property value. 
func (m *SigningCertificateUpdateStatus) SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRunDateTime = value
    }
}
