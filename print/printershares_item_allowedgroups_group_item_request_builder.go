package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrintersharesItemAllowedgroupsGroupItemRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\allowedGroups\{group-id}
type PrintersharesItemAllowedgroupsGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPrintersharesItemAllowedgroupsGroupItemRequestBuilderInternal instantiates a new PrintersharesItemAllowedgroupsGroupItemRequestBuilder and sets the default values.
func NewPrintersharesItemAllowedgroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemAllowedgroupsGroupItemRequestBuilder) {
    m := &PrintersharesItemAllowedgroupsGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedGroups/{group%2Did}", pathParameters),
    }
    return m
}
// NewPrintersharesItemAllowedgroupsGroupItemRequestBuilder instantiates a new PrintersharesItemAllowedgroupsGroupItemRequestBuilder and sets the default values.
func NewPrintersharesItemAllowedgroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemAllowedgroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersharesItemAllowedgroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
// returns a *PrintersharesItemAllowedgroupsItemRefRequestBuilder when successful
func (m *PrintersharesItemAllowedgroupsGroupItemRequestBuilder) Ref()(*PrintersharesItemAllowedgroupsItemRefRequestBuilder) {
    return NewPrintersharesItemAllowedgroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *PrintersharesItemAllowedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *PrintersharesItemAllowedgroupsGroupItemRequestBuilder) ServiceProvisioningErrors()(*PrintersharesItemAllowedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewPrintersharesItemAllowedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
