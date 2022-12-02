package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
type DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetQueryParameters user experience analytics devices without cloud identity.
type DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/{userExperienceAnalyticsDeviceWithoutCloudIdentity%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder instantiates a new UserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation user experience analytics devices without cloud identity.
func (m *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, requestConfiguration *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics devices without cloud identity.
func (m *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable), nil
}
// Patch update the navigation property userExperienceAnalyticsDevicesWithoutCloudIdentity in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, requestConfiguration *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable), nil
}
