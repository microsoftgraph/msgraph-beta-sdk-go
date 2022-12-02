package mobilitymanagementpolicies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder builds and executes requests for operations under \mobilityManagementPolicies\{mobilityManagementPolicy-id}\includedGroups\{group-id}
type MobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewMobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    m := &MobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/mobilityManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewMobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of mobilityManagementPolicy entities.
func (m *MobilityManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) Ref()(*MobilityManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    return NewMobilityManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
