package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i228524311e9abe0e004f8f298c0b324e8b951cac6aaeb2e20f6fe676036aa4a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item/parentsectiongroup"
    i31079f17a18956b2d40155fdc92cde3b5198682e4177bb79076cd90c6a7c9398 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item/sectiongroups"
    i42f920a160b98ee19eae4f3dd441e1c060f4457631df6375863af840c9154edc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item/sections"
    i649f3d863995b36a98eaa6562795c7197440678bd000f9f35e6f3c62a78eeef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item/parentnotebook"
    i3c7783c31f5945958dc6bed6d3bf70e052aadd5d3b68fa7f54b99d5efa7506c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item/sections/item"
    i9a389425fbcd9ed2f96d181c6980e310a1e219ea0c1eaa9c6242f867a9d37bbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item/sectiongroups/item"
)

// SectionGroupRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\sectionGroups\{sectionGroup-id}
type SectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SectionGroupRequestBuilderDeleteOptions options for Delete
type SectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SectionGroupRequestBuilderGetOptions options for Get
type SectionGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SectionGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SectionGroupRequestBuilderGetQueryParameters the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type SectionGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SectionGroupRequestBuilderPatchOptions options for Patch
type SectionGroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSectionGroupRequestBuilderInternal instantiates a new SectionGroupRequestBuilder and sets the default values.
func NewSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    m := &SectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/sectionGroups/{sectionGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSectionGroupRequestBuilder instantiates a new SectionGroupRequestBuilder and sets the default values.
func NewSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupRequestBuilder) CreateDeleteRequestInformation(options *SectionGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupRequestBuilder) CreateGetRequestInformation(options *SectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupRequestBuilder) CreatePatchRequestInformation(options *SectionGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupRequestBuilder) Delete(options *SectionGroupRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupRequestBuilder) Get(options *SectionGroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSectionGroup() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup), nil
}
func (m *SectionGroupRequestBuilder) ParentNotebook()(*i649f3d863995b36a98eaa6562795c7197440678bd000f9f35e6f3c62a78eeef9.ParentNotebookRequestBuilder) {
    return i649f3d863995b36a98eaa6562795c7197440678bd000f9f35e6f3c62a78eeef9.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) ParentSectionGroup()(*i228524311e9abe0e004f8f298c0b324e8b951cac6aaeb2e20f6fe676036aa4a7.ParentSectionGroupRequestBuilder) {
    return i228524311e9abe0e004f8f298c0b324e8b951cac6aaeb2e20f6fe676036aa4a7.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *SectionGroupRequestBuilder) Patch(options *SectionGroupRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *SectionGroupRequestBuilder) SectionGroups()(*i31079f17a18956b2d40155fdc92cde3b5198682e4177bb79076cd90c6a7c9398.SectionGroupsRequestBuilder) {
    return i31079f17a18956b2d40155fdc92cde3b5198682e4177bb79076cd90c6a7c9398.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sectionGroups.item.sectionGroups.item collection
func (m *SectionGroupRequestBuilder) SectionGroupsById(id string)(*i9a389425fbcd9ed2f96d181c6980e310a1e219ea0c1eaa9c6242f867a9d37bbc.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i9a389425fbcd9ed2f96d181c6980e310a1e219ea0c1eaa9c6242f867a9d37bbc.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) Sections()(*i42f920a160b98ee19eae4f3dd441e1c060f4457631df6375863af840c9154edc.SectionsRequestBuilder) {
    return i42f920a160b98ee19eae4f3dd441e1c060f4457631df6375863af840c9154edc.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sectionGroups.item.sections.item collection
func (m *SectionGroupRequestBuilder) SectionsById(id string)(*i3c7783c31f5945958dc6bed6d3bf70e052aadd5d3b68fa7f54b99d5efa7506c8.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i3c7783c31f5945958dc6bed6d3bf70e052aadd5d3b68fa7f54b99d5efa7506c8.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
