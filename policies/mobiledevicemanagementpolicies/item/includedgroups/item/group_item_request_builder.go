package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i61b420b9879bda500b60834850470fd5505e52498a22e08cf87bb1255987631b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobiledevicemanagementpolicies/item/includedgroups/item/ref"
)

// GroupItemRequestBuilder builds and executes requests for operations under \policies\mobileDeviceManagementPolicies\{mobilityManagementPolicy-id}\includedGroups\{group-id}
type GroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
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
// NewGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of policyRoot entities.
func (m *GroupItemRequestBuilder) Ref()(*i61b420b9879bda500b60834850470fd5505e52498a22e08cf87bb1255987631b.RefRequestBuilder) {
    return i61b420b9879bda500b60834850470fd5505e52498a22e08cf87bb1255987631b.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
