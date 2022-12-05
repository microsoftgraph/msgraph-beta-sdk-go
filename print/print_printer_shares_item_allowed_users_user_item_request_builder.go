package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrintPrinterSharesItemAllowedUsersUserItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedUsers\{user-id}
type PrintPrinterSharesItemAllowedUsersUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewPrintPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewPrintPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    m := &PrintPrinterSharesItemAllowedUsersUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedUsers/{user%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintPrinterSharesItemAllowedUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewPrintPrinterSharesItemAllowedUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *PrintPrinterSharesItemAllowedUsersUserItemRequestBuilder) Ref()(*PrintPrinterSharesItemAllowedUsersItemRefRequestBuilder) {
    return NewPrintPrinterSharesItemAllowedUsersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
