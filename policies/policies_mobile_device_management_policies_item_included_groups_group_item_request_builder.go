package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder builds and executes requests for operations under \policies\mobileDeviceManagementPolicies\{mobilityManagementPolicy-id}\includedGroups\{group-id}
type PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    m := &PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/mobileDeviceManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of policyRoot entities.
func (m *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) Ref()(*PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    return NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
