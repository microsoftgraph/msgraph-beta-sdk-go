package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
type InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetQueryParameters read the properties of an inboundSharedUserProfile.
type InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetQueryParameters
}
// InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderInternal instantiates a new InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder and sets the default values.
func NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) {
    m := &InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/inboundSharedUserProfiles/{inboundSharedUserProfile%2DuserId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder instantiates a new InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder and sets the default values.
func NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property inboundSharedUserProfiles for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *InboundshareduserprofilesItemExportpersonaldataExportPersonalDataRequestBuilder when successful
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) ExportPersonalData()(*InboundshareduserprofilesItemExportpersonaldataExportPersonalDataRequestBuilder) {
    return NewInboundshareduserprofilesItemExportpersonaldataExportPersonalDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties of an inboundSharedUserProfile.
// returns a InboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/inboundshareduserprofile-get?view=graph-rest-beta
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, error) {
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
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, error) {
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
// returns a *InboundshareduserprofilesItemRemovepersonaldataRemovePersonalDataRequestBuilder when successful
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) RemovePersonalData()(*InboundshareduserprofilesItemRemovepersonaldataRemovePersonalDataRequestBuilder) {
    return NewInboundshareduserprofilesItemRemovepersonaldataRemovePersonalDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property inboundSharedUserProfiles for directory
// returns a *RequestInformation when successful
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties of an inboundSharedUserProfile.
// returns a *RequestInformation when successful
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder when successful
func (m *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) WithUrl(rawUrl string)(*InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) {
    return NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
