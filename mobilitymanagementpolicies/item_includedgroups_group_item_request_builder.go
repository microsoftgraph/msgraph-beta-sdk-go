package mobilitymanagementpolicies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemIncludedgroupsGroupItemRequestBuilder builds and executes requests for operations under \mobilityManagementPolicies\{mobilityManagementPolicy-id}\includedGroups\{group-id}
type ItemIncludedgroupsGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemIncludedgroupsGroupItemRequestBuilderInternal instantiates a new ItemIncludedgroupsGroupItemRequestBuilder and sets the default values.
func NewItemIncludedgroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemIncludedgroupsGroupItemRequestBuilder) {
    m := &ItemIncludedgroupsGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/mobilityManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}", pathParameters),
    }
    return m
}
// NewItemIncludedgroupsGroupItemRequestBuilder instantiates a new ItemIncludedgroupsGroupItemRequestBuilder and sets the default values.
func NewItemIncludedgroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemIncludedgroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemIncludedgroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of mobilityManagementPolicy entities.
// returns a *ItemIncludedgroupsItemRefRequestBuilder when successful
func (m *ItemIncludedgroupsGroupItemRequestBuilder) Ref()(*ItemIncludedgroupsItemRefRequestBuilder) {
    return NewItemIncludedgroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *ItemIncludedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *ItemIncludedgroupsGroupItemRequestBuilder) ServiceProvisioningErrors()(*ItemIncludedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewItemIncludedgroupsItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
