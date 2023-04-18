package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrinterSharesItemAllowedUsersUserItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedUsers\{user-id}
type PrinterSharesItemAllowedUsersUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    m := &PrinterSharesItemAllowedUsersUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedUsers/{user%2Did}", pathParameters),
    }
    return m
}
// NewPrinterSharesItemAllowedUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MailboxSettings the mailboxSettings property
func (m *PrinterSharesItemAllowedUsersUserItemRequestBuilder) MailboxSettings()(*PrinterSharesItemAllowedUsersItemMailboxSettingsRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersItemMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *PrinterSharesItemAllowedUsersUserItemRequestBuilder) Ref()(*PrinterSharesItemAllowedUsersItemRefRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
