package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertEvidenceable 
type AlertEvidenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRemediationStatus()(*EvidenceRemediationStatus)
    GetRemediationStatusDetails()(*string)
    GetRoles()([]string)
    GetTags()([]string)
    GetType()(*string)
    GetVerdict()(*EvidenceVerdict)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRemediationStatus(value *EvidenceRemediationStatus)()
    SetRemediationStatusDetails(value *string)()
    SetRoles(value []string)()
    SetTags(value []string)()
    SetType(value *string)()
    SetVerdict(value *EvidenceVerdict)()
}
