package settings

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i604f8b04e00513fe80f3d85b5361d4a35dbdefc791356fd0e5d7cb3820a457bf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/profilecardproperties"
    ib5d9323a069790a40d8250ed97dce0595a704227bc01a0ff026952b59c07b6fa "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/peopleinsights"
    ic8dc6584c7fcf76e945a3a597a167477455672164e6ecef2a585705bb95f06f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/iteminsights"
    idad8116dcf0f508f2d5648089e624f18e9c892a32da2f4cd31cf8dfbb2acefad "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/microsoftapplicationdataaccess"
    i164665736586dec4189f70a600626997bf7f7b32e1d5debbea1adff37d4860df "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/profilecardproperties/item"
)

// SettingsRequestBuilder provides operations to manage the settings property of the microsoft.graph.organization entity.
type SettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SettingsRequestBuilderDeleteOptions options for Delete
type SettingsRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SettingsRequestBuilderGetOptions options for Get
type SettingsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SettingsRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SettingsRequestBuilderGetQueryParameters retrieve the properties and relationships of organizationSettings object. Nullable.
type SettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SettingsRequestBuilderPatchOptions options for Patch
type SettingsRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationSettingsable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewSettingsRequestBuilderInternal instantiates a new SettingsRequestBuilder and sets the default values.
func NewSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SettingsRequestBuilder) {
    m := &SettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/settings{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSettingsRequestBuilder instantiates a new SettingsRequestBuilder and sets the default values.
func NewSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property settings for organization
func (m *SettingsRequestBuilder) CreateDeleteRequestInformation(options *SettingsRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) CreateGetRequestInformation(options *SettingsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property settings in organization
func (m *SettingsRequestBuilder) CreatePatchRequestInformation(options *SettingsRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property settings for organization
func (m *SettingsRequestBuilder) Delete(options *SettingsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) Get(options *SettingsRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationSettingsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationSettingsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationSettingsable), nil
}
// ItemInsights the itemInsights property
func (m *SettingsRequestBuilder) ItemInsights()(*ic8dc6584c7fcf76e945a3a597a167477455672164e6ecef2a585705bb95f06f7.ItemInsightsRequestBuilder) {
    return ic8dc6584c7fcf76e945a3a597a167477455672164e6ecef2a585705bb95f06f7.NewItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftApplicationDataAccess the microsoftApplicationDataAccess property
func (m *SettingsRequestBuilder) MicrosoftApplicationDataAccess()(*idad8116dcf0f508f2d5648089e624f18e9c892a32da2f4cd31cf8dfbb2acefad.MicrosoftApplicationDataAccessRequestBuilder) {
    return idad8116dcf0f508f2d5648089e624f18e9c892a32da2f4cd31cf8dfbb2acefad.NewMicrosoftApplicationDataAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property settings in organization
func (m *SettingsRequestBuilder) Patch(options *SettingsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PeopleInsights the peopleInsights property
func (m *SettingsRequestBuilder) PeopleInsights()(*ib5d9323a069790a40d8250ed97dce0595a704227bc01a0ff026952b59c07b6fa.PeopleInsightsRequestBuilder) {
    return ib5d9323a069790a40d8250ed97dce0595a704227bc01a0ff026952b59c07b6fa.NewPeopleInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProfileCardProperties the profileCardProperties property
func (m *SettingsRequestBuilder) ProfileCardProperties()(*i604f8b04e00513fe80f3d85b5361d4a35dbdefc791356fd0e5d7cb3820a457bf.ProfileCardPropertiesRequestBuilder) {
    return i604f8b04e00513fe80f3d85b5361d4a35dbdefc791356fd0e5d7cb3820a457bf.NewProfileCardPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProfileCardPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.organization.item.settings.profileCardProperties.item collection
func (m *SettingsRequestBuilder) ProfileCardPropertiesById(id string)(*i164665736586dec4189f70a600626997bf7f7b32e1d5debbea1adff37d4860df.ProfileCardPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profileCardProperty%2Did"] = id
    }
    return i164665736586dec4189f70a600626997bf7f7b32e1d5debbea1adff37d4860df.NewProfileCardPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
