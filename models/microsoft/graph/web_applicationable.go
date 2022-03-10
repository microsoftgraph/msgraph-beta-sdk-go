package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WebApplicationable 
type WebApplicationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetHomePageUrl()(*string)
    GetImplicitGrantSettings()(ImplicitGrantSettingsable)
    GetLogoutUrl()(*string)
    GetOauth2AllowImplicitFlow()(*bool)
    GetRedirectUris()([]string)
    GetRedirectUriSettings()([]RedirectUriSettingsable)
    SetHomePageUrl(value *string)()
    SetImplicitGrantSettings(value ImplicitGrantSettingsable)()
    SetLogoutUrl(value *string)()
    SetOauth2AllowImplicitFlow(value *bool)()
    SetRedirectUris(value []string)()
    SetRedirectUriSettings(value []RedirectUriSettingsable)()
}
