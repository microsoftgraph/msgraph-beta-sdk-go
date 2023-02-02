package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrinterSharesItemAllowedGroupsGroupItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedGroups\{group-id}
type PrinterSharesItemAllowedGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, groupId *string)(*PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) {
    m := &PrinterSharesItemAllowedGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if groupId != nil {
        urlTplParams["group%2Did"] = *groupId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Ref provides operations to manage the collection of print entities.
func (m *PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) Ref()(*PrinterSharesItemAllowedGroupsItemRefRequestBuilder) {
    return NewPrinterSharesItemAllowedGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
