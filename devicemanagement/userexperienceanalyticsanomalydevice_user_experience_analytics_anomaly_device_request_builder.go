package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetQueryParameters the user experience analytics anomaly entity contains device details.
type UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsAnomalyDeviceId provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilder when successful
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) ByUserExperienceAnalyticsAnomalyDeviceId(userExperienceAnalyticsAnomalyDeviceId string)(*UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsAnomalyDeviceId != "" {
        urlTplParams["userExperienceAnalyticsAnomalyDevice%2Did"] = userExperienceAnalyticsAnomalyDeviceId
    }
    return NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal instantiates a new UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    m := &UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAnomalyDevice{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder instantiates a new UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsanomalydeviceCountRequestBuilder when successful
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) Count()(*UserexperienceanalyticsanomalydeviceCountRequestBuilder) {
    return NewUserexperienceanalyticsanomalydeviceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the user experience analytics anomaly entity contains device details.
// returns a UserExperienceAnalyticsAnomalyDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyDeviceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAnomalyDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyDeviceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsAnomalyDevice for deviceManagement
// returns a UserExperienceAnalyticsAnomalyDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyDeviceable, requestConfiguration *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyDeviceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyDeviceable), nil
}
// ToGetRequestInformation the user experience analytics anomaly entity contains device details.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsAnomalyDevice for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyDeviceable, requestConfiguration *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder when successful
func (m *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    return NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
