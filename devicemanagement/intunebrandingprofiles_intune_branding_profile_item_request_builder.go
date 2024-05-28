package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
type IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetQueryParameters intune branding profiles targeted to AAD groups
type IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetQueryParameters
}
// IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *IntunebrandingprofilesItemAssignRequestBuilder when successful
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) Assign()(*IntunebrandingprofilesItemAssignRequestBuilder) {
    return NewIntunebrandingprofilesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.intuneBrandingProfile entity.
// returns a *IntunebrandingprofilesItemAssignmentsRequestBuilder when successful
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) Assignments()(*IntunebrandingprofilesItemAssignmentsRequestBuilder) {
    return NewIntunebrandingprofilesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderInternal instantiates a new IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder and sets the default values.
func NewIntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) {
    m := &IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intuneBrandingProfiles/{intuneBrandingProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder instantiates a new IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder and sets the default values.
func NewIntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property intuneBrandingProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get intune branding profiles targeted to AAD groups
// returns a IntuneBrandingProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IntuneBrandingProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIntuneBrandingProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IntuneBrandingProfileable), nil
}
// Patch update the navigation property intuneBrandingProfiles in deviceManagement
// returns a IntuneBrandingProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IntuneBrandingProfileable, requestConfiguration *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IntuneBrandingProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIntuneBrandingProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IntuneBrandingProfileable), nil
}
// ToDeleteRequestInformation delete navigation property intuneBrandingProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation intune branding profiles targeted to AAD groups
// returns a *RequestInformation when successful
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property intuneBrandingProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IntuneBrandingProfileable, requestConfiguration *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder when successful
func (m *IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) WithUrl(rawUrl string)(*IntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder) {
    return NewIntunebrandingprofilesIntuneBrandingProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
