package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetQueryParameters user experience analytics devices without cloud identity.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal instantiates a new UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    m := &UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/{userExperienceAnalyticsDeviceWithoutCloudIdentity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder instantiates a new UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics devices without cloud identity.
// returns a UserExperienceAnalyticsDeviceWithoutCloudIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable), nil
}
// Patch update the navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity in deviceManagement
// returns a UserExperienceAnalyticsDeviceWithoutCloudIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics devices without cloud identity.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
