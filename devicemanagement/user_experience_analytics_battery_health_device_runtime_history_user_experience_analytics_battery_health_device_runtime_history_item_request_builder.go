package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/{userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder instantiates a new UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation user Experience Analytics Battery Health Device Runtime History
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get user Experience Analytics Battery Health Device Runtime History
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable), nil
}
// Patch update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable), nil
}
