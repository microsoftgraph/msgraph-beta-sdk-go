package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder provides operations to manage the microsoftApplicationDataAccess property of the microsoft.graph.organizationSettings entity.
type ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetQueryParameters get the settings in a microsoftApplicationDataAccessSettings object that specify access from Microsoft applications to Microsoft 365 user data in an organization.
type ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetQueryParameters
}
// ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderInternal instantiates a new ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder and sets the default values.
func NewItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) {
    m := &ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/settings/microsoftApplicationDataAccess{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder instantiates a new ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder and sets the default values.
func NewItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property microsoftApplicationDataAccess for organization
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the settings in a microsoftApplicationDataAccessSettings object that specify access from Microsoft applications to Microsoft 365 user data in an organization.
// returns a MicrosoftApplicationDataAccessSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/organizationsettings-list-microsoftapplicationdataaccess?view=graph-rest-beta
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftApplicationDataAccessSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftApplicationDataAccessSettingsable), nil
}
// Patch update the settings in a microsoftApplicationDataAccessSettings object that specify access from Microsoft applications to Microsoft 365 user data in an organization.
// returns a MicrosoftApplicationDataAccessSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/microsoftapplicationdataaccesssettings-update?view=graph-rest-beta
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftApplicationDataAccessSettingsable, requestConfiguration *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftApplicationDataAccessSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftApplicationDataAccessSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property microsoftApplicationDataAccess for organization
// returns a *RequestInformation when successful
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the settings in a microsoftApplicationDataAccessSettings object that specify access from Microsoft applications to Microsoft 365 user data in an organization.
// returns a *RequestInformation when successful
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the settings in a microsoftApplicationDataAccessSettings object that specify access from Microsoft applications to Microsoft 365 user data in an organization.
// returns a *RequestInformation when successful
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftApplicationDataAccessSettingsable, requestConfiguration *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder when successful
func (m *ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder) {
    return NewItemSettingsMicrosoftapplicationdataaccessMicrosoftApplicationDataAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
