package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
type DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetQueryParameters collection of Derived credential settings associated with account.
type DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetQueryParameters
}
// DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderInternal instantiates a new DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder and sets the default values.
func NewDerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) {
    m := &DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/derivedCredentials/{deviceManagementDerivedCredentialSettings%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder instantiates a new DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder and sets the default values.
func NewDerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property derivedCredentials for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of Derived credential settings associated with account.
// returns a DeviceManagementDerivedCredentialSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDerivedCredentialSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDerivedCredentialSettingsable), nil
}
// Patch update the navigation property derivedCredentials in deviceManagement
// returns a DeviceManagementDerivedCredentialSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDerivedCredentialSettingsable, requestConfiguration *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDerivedCredentialSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDerivedCredentialSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property derivedCredentials for deviceManagement
// returns a *RequestInformation when successful
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of Derived credential settings associated with account.
// returns a *RequestInformation when successful
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property derivedCredentials in deviceManagement
// returns a *RequestInformation when successful
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDerivedCredentialSettingsable, requestConfiguration *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder when successful
func (m *DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) WithUrl(rawUrl string)(*DerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) {
    return NewDerivedcredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
