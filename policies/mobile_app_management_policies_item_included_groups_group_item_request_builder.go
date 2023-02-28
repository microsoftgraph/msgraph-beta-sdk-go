package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder builds and executes requests for operations under \policies\mobileAppManagementPolicies\{mobilityManagementPolicy-id}\includedGroups\{group-id}
type MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    m := &MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/mobileAppManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of policyRoot entities.
func (m *MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) Ref()(*MobileAppManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    return NewMobileAppManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
