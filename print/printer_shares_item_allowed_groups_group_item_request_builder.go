package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrinterSharesItemAllowedGroupsGroupItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedGroups\{group-id}
type PrinterSharesItemAllowedGroupsGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) {
    m := &PrinterSharesItemAllowedGroupsGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedGroups/{group%2Did}", pathParameters),
    }
    return m
}
// NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemAllowedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) Ref()(*PrinterSharesItemAllowedGroupsItemRefRequestBuilder) {
    return NewPrinterSharesItemAllowedGroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
func (m *PrinterSharesItemAllowedGroupsGroupItemRequestBuilder) ServiceProvisioningErrors()(*PrinterSharesItemAllowedGroupsItemServiceProvisioningErrorsRequestBuilder) {
    return NewPrinterSharesItemAllowedGroupsItemServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
