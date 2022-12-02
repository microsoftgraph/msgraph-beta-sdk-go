package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder provides operations to manage the media for the organization entity.
type OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderInternal instantiates a new CustomCSSRequestBuilder and sets the default values.
func NewOrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) {
    m := &OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/branding/localizations/{organizationalBrandingLocalization%2Did}/customCSS";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder instantiates a new CustomCSSRequestBuilder and sets the default values.
func NewOrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get customCSS for the navigation property localizations from organization
func (m *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePutRequestInformation update customCSS for the navigation property localizations in organization
func (m *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) CreatePutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get customCSS for the navigation property localizations from organization
func (m *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put update customCSS for the navigation property localizations in organization
func (m *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderPutRequestConfiguration)(error) {
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
