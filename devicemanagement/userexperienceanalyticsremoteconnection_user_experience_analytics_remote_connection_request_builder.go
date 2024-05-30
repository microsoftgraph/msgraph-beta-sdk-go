package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetQueryParameters user experience analytics remote connection
type UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsRemoteConnectionId provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilder when successful
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) ByUserExperienceAnalyticsRemoteConnectionId(userExperienceAnalyticsRemoteConnectionId string)(*UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsRemoteConnectionId != "" {
        urlTplParams["userExperienceAnalyticsRemoteConnection%2Did"] = userExperienceAnalyticsRemoteConnectionId
    }
    return NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal instantiates a new UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder and sets the default values.
func NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    m := &UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsRemoteConnection{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder instantiates a new UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder and sets the default values.
func NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsremoteconnectionCountRequestBuilder when successful
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) Count()(*UserexperienceanalyticsremoteconnectionCountRequestBuilder) {
    return NewUserexperienceanalyticsremoteconnectionCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics remote connection
// returns a UserExperienceAnalyticsRemoteConnectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRemoteConnectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsRemoteConnectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRemoteConnectionCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsRemoteConnection for deviceManagement
// returns a UserExperienceAnalyticsRemoteConnectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRemoteConnectionable, requestConfiguration *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRemoteConnectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRemoteConnectionable), nil
}
// SummarizeDeviceRemoteConnectionWithSummarizeBy provides operations to call the summarizeDeviceRemoteConnection method.
// returns a *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder when successful
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) SummarizeDeviceRemoteConnectionWithSummarizeBy(summarizeBy *string)(*UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    return NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, summarizeBy)
}
// ToGetRequestInformation user experience analytics remote connection
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsRemoteConnection for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRemoteConnectionable, requestConfiguration *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder when successful
func (m *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    return NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
