package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
type ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetQueryParameters get sensitivityPolicySettings from users
type ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetQueryParameters
}
// ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal instantiates a new ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder and sets the default values.
func NewItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    m := &ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/informationProtection/sensitivityPolicySettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder instantiates a new ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder and sets the default values.
func NewItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sensitivityPolicySettings for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get sensitivityPolicySettings from users
// returns a SensitivityPolicySettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityPolicySettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable), nil
}
// Patch update the navigation property sensitivityPolicySettings in users
// returns a SensitivityPolicySettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, requestConfiguration *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityPolicySettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable), nil
}
// ToDeleteRequestInformation delete navigation property sensitivityPolicySettings for users
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get sensitivityPolicySettings from users
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sensitivityPolicySettings in users
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, requestConfiguration *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder when successful
func (m *ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    return NewItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
