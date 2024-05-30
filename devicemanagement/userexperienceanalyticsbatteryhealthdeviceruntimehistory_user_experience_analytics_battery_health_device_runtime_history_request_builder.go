package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetQueryParameters user Experience Analytics Battery Health Device Runtime History
type UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) ByUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId(userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId string)(*UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory%2Did"] = userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId
    }
    return NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal instantiates a new UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    m := &UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder instantiates a new UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryCountRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) Count()(*UserexperienceanalyticsbatteryhealthdeviceruntimehistoryCountRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user Experience Analytics Battery Health Device Runtime History
// returns a UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
// returns a UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable), nil
}
// ToGetRequestInformation user Experience Analytics Battery Health Device Runtime History
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
