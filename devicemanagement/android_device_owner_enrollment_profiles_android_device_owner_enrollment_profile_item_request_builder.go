package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters android device owner enrollment profile entities.
type AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal instantiates a new AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    m := &AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles/{androidDeviceOwnerEnrollmentProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder instantiates a new AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateToken provides operations to call the createToken method.
// returns a *AndroidDeviceOwnerEnrollmentProfilesItemCreateTokenRequestBuilder when successful
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateToken()(*AndroidDeviceOwnerEnrollmentProfilesItemCreateTokenRequestBuilder) {
    return NewAndroidDeviceOwnerEnrollmentProfilesItemCreateTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
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
// returns a *AndroidDeviceOwnerEnrollmentProfilesItemRevokeTokenRequestBuilder when successful
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) RevokeToken()(*AndroidDeviceOwnerEnrollmentProfilesItemRevokeTokenRequestBuilder) {
    return NewAndroidDeviceOwnerEnrollmentProfilesItemRevokeTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder when successful
func (m *AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) WithUrl(rawUrl string)(*AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    return NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
