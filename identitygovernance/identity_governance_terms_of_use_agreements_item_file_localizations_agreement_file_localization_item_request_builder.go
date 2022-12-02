package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.agreementFile entity.
type IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetQueryParameters the localized version of the terms of use agreement files attached to the agreement.
type IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal instantiates a new AgreementFileLocalizationItemRequestBuilder and sets the default values.
func NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) {
    m := &IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/file/localizations/{agreementFileLocalization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder instantiates a new AgreementFileLocalizationItemRequestBuilder and sets the default values.
func NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property localizations for identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the localized version of the terms of use agreement files attached to the agreement.
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property localizations in identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileLocalizationable, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property localizations for identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the localized version of the terms of use agreement files attached to the agreement.
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileLocalizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgreementFileLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileLocalizationable), nil
}
// Patch update the navigation property localizations in identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileLocalizationable, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileLocalizationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgreementFileLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileLocalizationable), nil
}
// Versions provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Versions()(*IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) {
    return NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
func (m *IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) VersionsById(id string)(*IdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileVersion%2Did"] = id
    }
    return NewIdentityGovernanceTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
