package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder provides operations to manage the includedGroups property of the microsoft.graph.mobilityManagementPolicy entity.
type MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetQueryParameters get the list of groups that are included in a mobile app management policy.
type MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetQueryParameters
}
// NewMobileAppManagementPoliciesItemIncludedGroupsRequestBuilderInternal instantiates a new IncludedGroupsRequestBuilder and sets the default values.
func NewMobileAppManagementPoliciesItemIncludedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder) {
    m := &MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/mobileAppManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppManagementPoliciesItemIncludedGroupsRequestBuilder instantiates a new IncludedGroupsRequestBuilder and sets the default values.
func NewMobileAppManagementPoliciesItemIncludedGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppManagementPoliciesItemIncludedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder) Count()(*MobileAppManagementPoliciesItemIncludedGroupsCountRequestBuilder) {
    return NewMobileAppManagementPoliciesItemIncludedGroupsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the list of groups that are included in a mobile app management policy.
func (m *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the list of groups that are included in a mobile app management policy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/mobileappmanagementpolicies-list-includedgroups?view=graph-rest-1.0
func (m *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// Ref provides operations to manage the collection of policyRoot entities.
func (m *MobileAppManagementPoliciesItemIncludedGroupsRequestBuilder) Ref()(*MobileAppManagementPoliciesItemIncludedGroupsRefRequestBuilder) {
    return NewMobileAppManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
