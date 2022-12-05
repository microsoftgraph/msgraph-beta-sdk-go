package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder provides operations to manage the media for the organization entity.
type OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOrganizationItemBrandingLocalizationsItemFaviconRequestBuilderInternal instantiates a new FaviconRequestBuilder and sets the default values.
func NewOrganizationItemBrandingLocalizationsItemFaviconRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) {
    m := &OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/branding/localizations/{organizationalBrandingLocalization%2Did}/favicon";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationItemBrandingLocalizationsItemFaviconRequestBuilder instantiates a new FaviconRequestBuilder and sets the default values.
func NewOrganizationItemBrandingLocalizationsItemFaviconRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemBrandingLocalizationsItemFaviconRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation a custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
func (m *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePutRequestInformation a custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
func (m *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) CreatePutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get a custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
func (m *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put a custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
func (m *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OrganizationItemBrandingLocalizationsItemFaviconRequestBuilderPutRequestConfiguration)(error) {
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
