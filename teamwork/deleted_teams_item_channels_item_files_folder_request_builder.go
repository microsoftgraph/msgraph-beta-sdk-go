package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemFilesFolderRequestBuilder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
type DeletedTeamsItemChannelsItemFilesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetQueryParameters metadata for the location where the channel's files are stored.
type DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetQueryParameters
}
// NewDeletedTeamsItemChannelsItemFilesFolderRequestBuilderInternal instantiates a new DeletedTeamsItemChannelsItemFilesFolderRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemFilesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemFilesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/filesFolder{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedTeamsItemChannelsItemFilesFolderRequestBuilder instantiates a new DeletedTeamsItemChannelsItemFilesFolderRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemFilesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemFilesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the teamwork entity.
// returns a *DeletedTeamsItemChannelsItemFilesFolderContentRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) Content()(*DeletedTeamsItemChannelsItemFilesFolderContentRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemFilesFolderContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ContentStream provides operations to manage the media for the teamwork entity.
// returns a *DeletedTeamsItemChannelsItemFilesFolderContentStreamRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) ContentStream()(*DeletedTeamsItemChannelsItemFilesFolderContentStreamRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemFilesFolderContentStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get metadata for the location where the channel's files are stored.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ToGetRequestInformation metadata for the location where the channel's files are stored.
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemFilesFolderRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeletedTeamsItemChannelsItemFilesFolderRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) WithUrl(rawUrl string)(*DeletedTeamsItemChannelsItemFilesFolderRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemFilesFolderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
