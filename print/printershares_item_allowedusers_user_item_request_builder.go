package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrintersharesItemAllowedusersUserItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedUsers\{user-id}
type PrintersharesItemAllowedusersUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPrintersharesItemAllowedusersUserItemRequestBuilderInternal instantiates a new PrintersharesItemAllowedusersUserItemRequestBuilder and sets the default values.
func NewPrintersharesItemAllowedusersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemAllowedusersUserItemRequestBuilder) {
    m := &PrintersharesItemAllowedusersUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedUsers/{user%2Did}", pathParameters),
    }
    return m
}
// NewPrintersharesItemAllowedusersUserItemRequestBuilder instantiates a new PrintersharesItemAllowedusersUserItemRequestBuilder and sets the default values.
func NewPrintersharesItemAllowedusersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemAllowedusersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersharesItemAllowedusersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MailboxSettings the mailboxSettings property
// returns a *PrintersharesItemAllowedusersItemMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *PrintersharesItemAllowedusersUserItemRequestBuilder) MailboxSettings()(*PrintersharesItemAllowedusersItemMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewPrintersharesItemAllowedusersItemMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of print entities.
// returns a *PrintersharesItemAllowedusersItemRefRequestBuilder when successful
func (m *PrintersharesItemAllowedusersUserItemRequestBuilder) Ref()(*PrintersharesItemAllowedusersItemRefRequestBuilder) {
    return NewPrintersharesItemAllowedusersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *PrintersharesItemAllowedusersItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *PrintersharesItemAllowedusersUserItemRequestBuilder) ServiceProvisioningErrors()(*PrintersharesItemAllowedusersItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewPrintersharesItemAllowedusersItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
