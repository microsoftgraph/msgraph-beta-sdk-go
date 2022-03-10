package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Custodianable 
type Custodianable interface {
    DataSourceContainerable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetApplyHoldToSources()(*bool)
    GetEmail()(*string)
    GetSiteSources()([]SiteSourceable)
    GetUnifiedGroupSources()([]UnifiedGroupSourceable)
    GetUserSources()([]UserSourceable)
    SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetApplyHoldToSources(value *bool)()
    SetEmail(value *string)()
    SetSiteSources(value []SiteSourceable)()
    SetUnifiedGroupSources(value []UnifiedGroupSourceable)()
    SetUserSources(value []UserSourceable)()
}
