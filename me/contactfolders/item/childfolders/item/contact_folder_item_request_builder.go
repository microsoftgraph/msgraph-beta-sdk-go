package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07be5face02e2c1a4d150e0406cf2f38f7e750ce2acecd71ed70f2826713c15f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item/singlevalueextendedproperties"
    i147e6fb8b2bcf39ee199590249db2ecee28f7c975393ffcb830f025926dbd88c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item/contacts"
    i70e69e3d99aad2354270856e113c2473d4df686e8010f3a0d12dfc635a4890db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item/multivalueextendedproperties"
    i49acd1883e0dc334bb2f049b419b35c46baeadba192d92345f1d33e834bb7f93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item/contacts/item"
    i4fe953901040256e83f553396dd7370606f863f1d37918710b178c43cbbd073b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item/multivalueextendedproperties/item"
    ie305967796e678e5afc3f1a1a2600b8c5957abaecacbf74640d6682fea2f20e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item/singlevalueextendedproperties/item"
)

// ContactFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
type ContactFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContactFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContactFolderItemRequestBuilderGetQueryParameters the collection of child folders in the folder. Navigation property. Read-only. Nullable.
type ContactFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContactFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContactFolderItemRequestBuilderGetQueryParameters
}
// ContactFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContactFolderItemRequestBuilderInternal instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    m := &ContactFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder%2Did}/childFolders/{contactFolder%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactFolderItemRequestBuilder instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Contacts the contacts property
func (m *ContactFolderItemRequestBuilder) Contacts()(*i147e6fb8b2bcf39ee199590249db2ecee28f7c975393ffcb830f025926dbd88c.ContactsRequestBuilder) {
    return i147e6fb8b2bcf39ee199590249db2ecee28f7c975393ffcb830f025926dbd88c.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.childFolders.item.contacts.item collection
func (m *ContactFolderItemRequestBuilder) ContactsById(id string)(*i49acd1883e0dc334bb2f049b419b35c46baeadba192d92345f1d33e834bb7f93.ContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return i49acd1883e0dc334bb2f049b419b35c46baeadba192d92345f1d33e834bb7f93.NewContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property childFolders for me
func (m *ContactFolderItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property childFolders for me
func (m *ContactFolderItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ContactFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ContactFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property childFolders in me
func (m *ContactFolderItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property childFolders in me
func (m *ContactFolderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, requestConfiguration *ContactFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property childFolders for me
func (m *ContactFolderItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property childFolders for me
func (m *ContactFolderItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ContactFolderItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ContactFolderItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContactFolderFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedProperties()(*i70e69e3d99aad2354270856e113c2473d4df686e8010f3a0d12dfc635a4890db.MultiValueExtendedPropertiesRequestBuilder) {
    return i70e69e3d99aad2354270856e113c2473d4df686e8010f3a0d12dfc635a4890db.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4fe953901040256e83f553396dd7370606f863f1d37918710b178c43cbbd073b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i4fe953901040256e83f553396dd7370606f863f1d37918710b178c43cbbd073b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in me
func (m *ContactFolderItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property childFolders in me
func (m *ContactFolderItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, requestConfiguration *ContactFolderItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedProperties()(*i07be5face02e2c1a4d150e0406cf2f38f7e750ce2acecd71ed70f2826713c15f.SingleValueExtendedPropertiesRequestBuilder) {
    return i07be5face02e2c1a4d150e0406cf2f38f7e750ce2acecd71ed70f2826713c15f.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie305967796e678e5afc3f1a1a2600b8c5957abaecacbf74640d6682fea2f20e4.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ie305967796e678e5afc3f1a1a2600b8c5957abaecacbf74640d6682fea2f20e4.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
