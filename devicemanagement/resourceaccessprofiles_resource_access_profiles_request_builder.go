package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceaccessprofilesResourceAccessProfilesRequestBuilder provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
type ResourceaccessprofilesResourceAccessProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetQueryParameters collection of resource access settings associated with account.
type ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetQueryParameters struct {
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
// ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetQueryParameters
}
// ResourceaccessprofilesResourceAccessProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesResourceAccessProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementResourceAccessProfileBaseId provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder when successful
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) ByDeviceManagementResourceAccessProfileBaseId(deviceManagementResourceAccessProfileBaseId string)(*ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementResourceAccessProfileBaseId != "" {
        urlTplParams["deviceManagementResourceAccessProfileBase%2Did"] = deviceManagementResourceAccessProfileBaseId
    }
    return NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewResourceaccessprofilesResourceAccessProfilesRequestBuilderInternal instantiates a new ResourceaccessprofilesResourceAccessProfilesRequestBuilder and sets the default values.
func NewResourceaccessprofilesResourceAccessProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesResourceAccessProfilesRequestBuilder) {
    m := &ResourceaccessprofilesResourceAccessProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/resourceAccessProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewResourceaccessprofilesResourceAccessProfilesRequestBuilder instantiates a new ResourceaccessprofilesResourceAccessProfilesRequestBuilder and sets the default values.
func NewResourceaccessprofilesResourceAccessProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesResourceAccessProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceaccessprofilesResourceAccessProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ResourceaccessprofilesCountRequestBuilder when successful
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) Count()(*ResourceaccessprofilesCountRequestBuilder) {
    return NewResourceaccessprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of resource access settings associated with account.
// returns a DeviceManagementResourceAccessProfileBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseCollectionResponseable), nil
}
// Post create new navigation property to resourceAccessProfiles for deviceManagement
// returns a DeviceManagementResourceAccessProfileBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *ResourceaccessprofilesResourceAccessProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
// QueryByPlatformType provides operations to call the queryByPlatformType method.
// returns a *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder when successful
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) QueryByPlatformType()(*ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) {
    return NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation collection of resource access settings associated with account.
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ResourceaccessprofilesResourceAccessProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to resourceAccessProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *ResourceaccessprofilesResourceAccessProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ResourceaccessprofilesResourceAccessProfilesRequestBuilder when successful
func (m *ResourceaccessprofilesResourceAccessProfilesRequestBuilder) WithUrl(rawUrl string)(*ResourceaccessprofilesResourceAccessProfilesRequestBuilder) {
    return NewResourceaccessprofilesResourceAccessProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
