package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrinterSharesItemAllowedUsersUserItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedUsers\{user-id}
type PrinterSharesItemAllowedUsersUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    m := &PrinterSharesItemAllowedUsersUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedUsers/{user%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPrinterSharesItemAllowedUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *PrinterSharesItemAllowedUsersUserItemRequestBuilder) Ref()(*PrinterSharesItemAllowedUsersItemRefRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
