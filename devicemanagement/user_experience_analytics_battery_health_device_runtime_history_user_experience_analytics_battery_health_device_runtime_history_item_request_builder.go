package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetQueryParameters user Experience Analytics Battery Health Device Runtime History
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/{userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder instantiates a new UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get user Experience Analytics Battery Health Device Runtime History
// returns a UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
// returns a UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user Experience Analytics Battery Health Device Runtime History
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
