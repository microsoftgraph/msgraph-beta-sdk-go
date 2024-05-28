package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetQueryParameters user Experience Analytics Battery Health Device Performance
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    m := &UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/{userExperienceAnalyticsBatteryHealthDevicePerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder instantiates a new UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsBatteryHealthDevicePerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user Experience Analytics Battery Health Device Performance
// returns a UserExperienceAnalyticsBatteryHealthDevicePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsBatteryHealthDevicePerformance in deviceManagement
// returns a UserExperienceAnalyticsBatteryHealthDevicePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsBatteryHealthDevicePerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user Experience Analytics Battery Health Device Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsBatteryHealthDevicePerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
