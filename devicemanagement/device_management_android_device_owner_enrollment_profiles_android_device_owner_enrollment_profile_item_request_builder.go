package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters android device owner enrollment profile entities.
type DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal instantiates a new AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    m := &DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles/{androidDeviceOwnerEnrollmentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder instantiates a new AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation android device owner enrollment profile entities.
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property androidDeviceOwnerEnrollmentProfiles in deviceManagement
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateToken provides operations to call the createToken method.
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateToken()(*DeviceManagementAndroidDeviceOwnerEnrollmentProfilesItemCreateTokenRequestBuilder) {
    return NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesItemCreateTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get android device owner enrollment profile entities.
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable), nil
}
// Patch update the navigation property androidDeviceOwnerEnrollmentProfiles in deviceManagement
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable), nil
}
// RevokeToken provides operations to call the revokeToken method.
func (m *DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) RevokeToken()(*DeviceManagementAndroidDeviceOwnerEnrollmentProfilesItemRevokeTokenRequestBuilder) {
    return NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesItemRevokeTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
