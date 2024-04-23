package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
type InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetQueryParameters a collection of external users whose profile data is shared with the Microsoft Entra tenant. Nullable.
type InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetQueryParameters
}
// InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderInternal instantiates a new InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder and sets the default values.
func NewInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) {
    m := &InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/inboundSharedUserProfiles/{inboundSharedUserProfile%2DuserId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder instantiates a new InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder and sets the default values.
func NewInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property inboundSharedUserProfiles for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExportPersonalData provides operations to call the exportPersonalData method.
// returns a *InboundSharedUserProfilesItemExportPersonalDataRequestBuilder when successful
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) ExportPersonalData()(*InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) {
    return NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of external users whose profile data is shared with the Microsoft Entra tenant. Nullable.
// returns a InboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInboundSharedUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable), nil
}
// Patch update the navigation property inboundSharedUserProfiles in directory
// returns a InboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, requestConfiguration *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInboundSharedUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable), nil
}
// RemovePersonalData provides operations to call the removePersonalData method.
// returns a *InboundSharedUserProfilesItemRemovePersonalDataRequestBuilder when successful
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) RemovePersonalData()(*InboundSharedUserProfilesItemRemovePersonalDataRequestBuilder) {
    return NewInboundSharedUserProfilesItemRemovePersonalDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property inboundSharedUserProfiles for directory
// returns a *RequestInformation when successful
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of external users whose profile data is shared with the Microsoft Entra tenant. Nullable.
// returns a *RequestInformation when successful
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property inboundSharedUserProfiles in directory
// returns a *RequestInformation when successful
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, requestConfiguration *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder when successful
func (m *InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) WithUrl(rawUrl string)(*InboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) {
    return NewInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
