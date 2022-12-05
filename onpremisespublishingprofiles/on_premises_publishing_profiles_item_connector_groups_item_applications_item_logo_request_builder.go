package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder provides operations to manage the media for the onPremisesPublishingProfile entity.
type OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderInternal instantiates a new LogoRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/applications/{application%2Did}/logo";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder instantiates a new LogoRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the main logo for the application. Not nullable.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePutRequestInformation the main logo for the application. Not nullable.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) CreatePutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the main logo for the application. Not nullable.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) Get(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the main logo for the application. Not nullable.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsItemLogoRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
