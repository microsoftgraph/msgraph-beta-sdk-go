package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetQueryParameters get the list of groups that are included in a mobile app management policy.
type PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetQueryParameters
}
// PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) {
    m := &PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/mobileDeviceManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/$ref{?%24top,%24skip,%24search,%24filter,%24count,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the list of groups that are included in a mobile app management policy.
func (m *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property ref to includedGroups for policies
func (m *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceCreateable, requestConfiguration *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the list of groups that are included in a mobile app management policy.
func (m *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateStringCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringCollectionResponseable), nil
}
// Post create new navigation property ref to includedGroups for policies
func (m *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceCreateable, requestConfiguration *PoliciesMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
