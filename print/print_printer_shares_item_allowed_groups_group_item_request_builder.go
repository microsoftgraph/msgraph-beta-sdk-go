package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedGroups\{group-id}
type PrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewPrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder) {
    m := &PrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *PrintPrinterSharesItemAllowedGroupsGroupItemRequestBuilder) Ref()(*PrintPrinterSharesItemAllowedGroupsItemRefRequestBuilder) {
    return NewPrintPrinterSharesItemAllowedGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
