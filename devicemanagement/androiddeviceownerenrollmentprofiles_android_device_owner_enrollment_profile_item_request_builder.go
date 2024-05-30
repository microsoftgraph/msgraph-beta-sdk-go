package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters android device owner enrollment profile entities.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal instantiates a new AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    m := &AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles/{androidDeviceOwnerEnrollmentProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder instantiates a new AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateToken provides operations to call the createToken method.
// returns a *AndroiddeviceownerenrollmentprofilesItemCreatetokenCreateTokenRequestBuilder when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateToken()(*AndroiddeviceownerenrollmentprofilesItemCreatetokenCreateTokenRequestBuilder) {
    return NewAndroiddeviceownerenrollmentprofilesItemCreatetokenCreateTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get android device owner enrollment profile entities.
// returns a AndroidDeviceOwnerEnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable), nil
}
// Patch update the navigation property androidDeviceOwnerEnrollmentProfiles in deviceManagement
// returns a AndroidDeviceOwnerEnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable), nil
}
// RevokeToken provides operations to call the revokeToken method.
// returns a *AndroiddeviceownerenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) RevokeToken()(*AndroiddeviceownerenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) {
    return NewAndroiddeviceownerenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation android device owner enrollment profile entities.
// returns a *RequestInformation when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property androidDeviceOwnerEnrollmentProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) WithUrl(rawUrl string)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    return NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
