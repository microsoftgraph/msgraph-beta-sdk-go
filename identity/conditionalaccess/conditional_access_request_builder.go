package conditionalaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b765a6506e9c469fc394944929f3f3f68abc1cd67221843e72d0023cdc7b9a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationcontextclassreferences"
    i613a4d67bebbdad874d25a6af57693740b126f2be0ba8f0683a6c40352680356 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationstrengths"
    i74336fcd63e509c2dff28b731f926303088803045222a57ca21181ed72fcf778 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/policies"
    i8b450d93e3a0398b0e364118a8d1d65374bc93acb38b538683e67f6f68a0df16 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/namedlocations"
    ic97d3f37939bde9c02ae8d631da4f7cf27c863c3df73816d823df479eeae856e "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/templates"
    i1e1e27fc92484640c1240b9cb5fb9a5b81b2571a142348b22dc8e7cfa597bdea "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationcontextclassreferences/item"
    i236cf09f7afec6e3a311ec7d864db943dfb016d470e91251403c97b871f8c2be "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/templates/item"
    i4e66627208bf17ad44b017955b0a38ce27d6db98bb96808a5974cac73445707c "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/policies/item"
    ib6f953303fb5d962a08cf74168e4f0d71eb89d20c00a8164e1326cd252166379 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/namedlocations/item"
)

// ConditionalAccessRequestBuilder provides operations to manage the conditionalAccess property of the microsoft.graph.identityContainer entity.
type ConditionalAccessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConditionalAccessRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessRequestBuilderGetQueryParameters the entry point for the Conditional Access (CA) object model.
type ConditionalAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessRequestBuilderGetQueryParameters
}
// ConditionalAccessRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationContextClassReferences the authenticationContextClassReferences property
func (m *ConditionalAccessRequestBuilder) AuthenticationContextClassReferences()(*i0b765a6506e9c469fc394944929f3f3f68abc1cd67221843e72d0023cdc7b9a5.AuthenticationContextClassReferencesRequestBuilder) {
    return i0b765a6506e9c469fc394944929f3f3f68abc1cd67221843e72d0023cdc7b9a5.NewAuthenticationContextClassReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationContextClassReferencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.authenticationContextClassReferences.item collection
func (m *ConditionalAccessRequestBuilder) AuthenticationContextClassReferencesById(id string)(*i1e1e27fc92484640c1240b9cb5fb9a5b81b2571a142348b22dc8e7cfa597bdea.AuthenticationContextClassReferenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationContextClassReference%2Did"] = id
    }
    return i1e1e27fc92484640c1240b9cb5fb9a5b81b2571a142348b22dc8e7cfa597bdea.NewAuthenticationContextClassReferenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationStrengths the authenticationStrengths property
func (m *ConditionalAccessRequestBuilder) AuthenticationStrengths()(*i613a4d67bebbdad874d25a6af57693740b126f2be0ba8f0683a6c40352680356.AuthenticationStrengthsRequestBuilder) {
    return i613a4d67bebbdad874d25a6af57693740b126f2be0ba8f0683a6c40352680356.NewAuthenticationStrengthsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewConditionalAccessRequestBuilderInternal instantiates a new ConditionalAccessRequestBuilder and sets the default values.
func NewConditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    m := &ConditionalAccessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConditionalAccessRequestBuilder instantiates a new ConditionalAccessRequestBuilder and sets the default values.
func NewConditionalAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property conditionalAccess for identity
func (m *ConditionalAccessRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the entry point for the Conditional Access (CA) object model.
func (m *ConditionalAccessRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property conditionalAccess in identity
func (m *ConditionalAccessRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessRootable, requestConfiguration *ConditionalAccessRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property conditionalAccess for identity
func (m *ConditionalAccessRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the entry point for the Conditional Access (CA) object model.
func (m *ConditionalAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConditionalAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessRootable), nil
}
// NamedLocations the namedLocations property
func (m *ConditionalAccessRequestBuilder) NamedLocations()(*i8b450d93e3a0398b0e364118a8d1d65374bc93acb38b538683e67f6f68a0df16.NamedLocationsRequestBuilder) {
    return i8b450d93e3a0398b0e364118a8d1d65374bc93acb38b538683e67f6f68a0df16.NewNamedLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamedLocationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.namedLocations.item collection
func (m *ConditionalAccessRequestBuilder) NamedLocationsById(id string)(*ib6f953303fb5d962a08cf74168e4f0d71eb89d20c00a8164e1326cd252166379.NamedLocationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["namedLocation%2Did"] = id
    }
    return ib6f953303fb5d962a08cf74168e4f0d71eb89d20c00a8164e1326cd252166379.NewNamedLocationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property conditionalAccess in identity
func (m *ConditionalAccessRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessRootable, requestConfiguration *ConditionalAccessRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConditionalAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessRootable), nil
}
// Policies the policies property
func (m *ConditionalAccessRequestBuilder) Policies()(*i74336fcd63e509c2dff28b731f926303088803045222a57ca21181ed72fcf778.PoliciesRequestBuilder) {
    return i74336fcd63e509c2dff28b731f926303088803045222a57ca21181ed72fcf778.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.policies.item collection
func (m *ConditionalAccessRequestBuilder) PoliciesById(id string)(*i4e66627208bf17ad44b017955b0a38ce27d6db98bb96808a5974cac73445707c.ConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = id
    }
    return i4e66627208bf17ad44b017955b0a38ce27d6db98bb96808a5974cac73445707c.NewConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Templates the templates property
func (m *ConditionalAccessRequestBuilder) Templates()(*ic97d3f37939bde9c02ae8d631da4f7cf27c863c3df73816d823df479eeae856e.TemplatesRequestBuilder) {
    return ic97d3f37939bde9c02ae8d631da4f7cf27c863c3df73816d823df479eeae856e.NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.templates.item collection
func (m *ConditionalAccessRequestBuilder) TemplatesById(id string)(*i236cf09f7afec6e3a311ec7d864db943dfb016d470e91251403c97b871f8c2be.ConditionalAccessTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessTemplate%2Did"] = id
    }
    return i236cf09f7afec6e3a311ec7d864db943dfb016d470e91251403c97b871f8c2be.NewConditionalAccessTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
