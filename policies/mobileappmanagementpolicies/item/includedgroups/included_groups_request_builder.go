package includedgroups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i9430413a1b7030b897ed8a84222cdccd426ad2331aa0dba541d783c8e21e98c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobileappmanagementpolicies/item/includedgroups/ref"
    ic5b7d6629b4a5c986e3a9357b5228149ea2d4a33d343593c8d1d7ec0756e4fb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobileappmanagementpolicies/item/includedgroups/count"
)

// IncludedGroupsRequestBuilder provides operations to manage the includedGroups property of the microsoft.graph.mobilityManagementPolicy entity.
type IncludedGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IncludedGroupsRequestBuilderGetQueryParameters azure AD groups under the scope of the mobility management application if appliesTo is selected
type IncludedGroupsRequestBuilderGetQueryParameters struct {
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
// IncludedGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IncludedGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IncludedGroupsRequestBuilderGetQueryParameters
}
// NewIncludedGroupsRequestBuilderInternal instantiates a new IncludedGroupsRequestBuilder and sets the default values.
func NewIncludedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IncludedGroupsRequestBuilder) {
    m := &IncludedGroupsRequestBuilder{
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
// NewIncludedGroupsRequestBuilder instantiates a new IncludedGroupsRequestBuilder and sets the default values.
func NewIncludedGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IncludedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIncludedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *IncludedGroupsRequestBuilder) Count()(*ic5b7d6629b4a5c986e3a9357b5228149ea2d4a33d343593c8d1d7ec0756e4fb4.CountRequestBuilder) {
    return ic5b7d6629b4a5c986e3a9357b5228149ea2d4a33d343593c8d1d7ec0756e4fb4.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *IncludedGroupsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *IncludedGroupsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *IncludedGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *IncludedGroupsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *IncludedGroupsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *IncludedGroupsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// Ref the ref property
func (m *IncludedGroupsRequestBuilder) Ref()(*i9430413a1b7030b897ed8a84222cdccd426ad2331aa0dba541d783c8e21e98c7.RefRequestBuilder) {
    return i9430413a1b7030b897ed8a84222cdccd426ad2331aa0dba541d783c8e21e98c7.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
