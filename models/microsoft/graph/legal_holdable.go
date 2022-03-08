package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// LegalHoldable 
type LegalHoldable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContentQuery()(*string)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetErrors()([]string)
    GetIsEnabled()(*bool)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSiteSources()([]SiteSourceable)
    GetStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.LegalHoldStatus)
    GetUnifiedGroupSources()([]UnifiedGroupSourceable)
    GetUserSources()([]UserSourceable)
    SetContentQuery(value *string)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetErrors(value []string)()
    SetIsEnabled(value *bool)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSiteSources(value []SiteSourceable)()
    SetStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.LegalHoldStatus)()
    SetUnifiedGroupSources(value []UnifiedGroupSourceable)()
    SetUserSources(value []UserSourceable)()
}
