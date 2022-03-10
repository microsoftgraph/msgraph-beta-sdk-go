package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SymantecCodeSigningCertificateable 
type SymantecCodeSigningCertificateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContent()([]byte)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIssuer()(*string)
    GetIssuerName()(*string)
    GetPassword()(*string)
    GetStatus()(*CertificateStatus)
    GetSubject()(*string)
    GetSubjectName()(*string)
    GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContent(value []byte)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIssuer(value *string)()
    SetIssuerName(value *string)()
    SetPassword(value *string)()
    SetStatus(value *CertificateStatus)()
    SetSubject(value *string)()
    SetSubjectName(value *string)()
    SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
