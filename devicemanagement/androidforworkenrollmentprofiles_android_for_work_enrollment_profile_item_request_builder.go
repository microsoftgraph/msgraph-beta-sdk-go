package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetQueryParameters android for Work enrollment profile entities.
type AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderInternal instantiates a new AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) {
    m := &AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkEnrollmentProfiles/{androidForWorkEnrollmentProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder instantiates a new AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateToken provides operations to call the createToken method.
// returns a *AndroidforworkenrollmentprofilesItemCreatetokenCreateTokenRequestBuilder when successful
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) CreateToken()(*AndroidforworkenrollmentprofilesItemCreatetokenCreateTokenRequestBuilder) {
    return NewAndroidforworkenrollmentprofilesItemCreatetokenCreateTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property androidForWorkEnrollmentProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get android for Work enrollment profile entities.
// returns a AndroidForWorkEnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkEnrollmentProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkEnrollmentProfileable), nil
}
// Patch update the navigation property androidForWorkEnrollmentProfiles in deviceManagement
// returns a AndroidForWorkEnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkEnrollmentProfileable, requestConfiguration *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkEnrollmentProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkEnrollmentProfileable), nil
}
// RevokeToken provides operations to call the revokeToken method.
// returns a *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder when successful
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) RevokeToken()(*AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) {
    return NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property androidForWorkEnrollmentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation android for Work enrollment profile entities.
// returns a *RequestInformation when successful
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property androidForWorkEnrollmentProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkEnrollmentProfileable, requestConfiguration *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder when successful
func (m *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) WithUrl(rawUrl string)(*AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder) {
    return NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
