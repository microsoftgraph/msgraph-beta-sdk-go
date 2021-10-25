package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0d60b783cdde9e605ad590623f0f617487311ddada4e650851ae767fd4c60701 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/attachments"
    i60ffa45d68c3a1da2ee7cdbf8353296a72a2301992bdf6ae9c7fb4508f0ea355 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/multivalueextendedproperties"
    i90988a76de3564e1bc9e592c2f427a1a8825a2eafa08f4a66b7313a9d5dea2e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/singlevalueextendedproperties"
    iccd03e13b1d81f87d437c5c29b5a0db4cfbd6cf3821e75339ed520e0dbe9cfbf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/complete"
    i0df7fa7dae6d3fd2cbcf896d192507ad89ba48690b7c467cd8e9c6f4a12be5dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/multivalueextendedproperties/item"
    i8be48ea3d303b95811d463c98ddcfa42f165836ce37eaf0634c5eb246c021887 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/singlevalueextendedproperties/item"
    if75877972b617042c4c94629e1bf5e70f9da63d56a12fd0f3afec1f771266412 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item/attachments/item"
)

type OutlookTaskRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OutlookTaskRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *OutlookTaskRequestBuilder) Attachments()(*i0d60b783cdde9e605ad590623f0f617487311ddada4e650851ae767fd4c60701.AttachmentsRequestBuilder) {
    return i0d60b783cdde9e605ad590623f0f617487311ddada4e650851ae767fd4c60701.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) AttachmentsById(id string)(*if75877972b617042c4c94629e1bf5e70f9da63d56a12fd0f3afec1f771266412.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return if75877972b617042c4c94629e1bf5e70f9da63d56a12fd0f3afec1f771266412.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Complete()(*iccd03e13b1d81f87d437c5c29b5a0db4cfbd6cf3821e75339ed520e0dbe9cfbf.CompleteRequestBuilder) {
    return iccd03e13b1d81f87d437c5c29b5a0db4cfbd6cf3821e75339ed520e0dbe9cfbf.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewOutlookTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    m := &OutlookTaskRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/outlook/taskFolders/{outlookTaskFolder_id}/tasks/{outlookTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOutlookTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OutlookTaskRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OutlookTaskRequestBuilder) CreateGetRequestInformation(q func (value *OutlookTaskRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OutlookTaskRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OutlookTaskRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OutlookTaskRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OutlookTaskRequestBuilder) Get(q func (value *OutlookTaskRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTask() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask), nil
}
func (m *OutlookTaskRequestBuilder) MultiValueExtendedProperties()(*i60ffa45d68c3a1da2ee7cdbf8353296a72a2301992bdf6ae9c7fb4508f0ea355.MultiValueExtendedPropertiesRequestBuilder) {
    return i60ffa45d68c3a1da2ee7cdbf8353296a72a2301992bdf6ae9c7fb4508f0ea355.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0df7fa7dae6d3fd2cbcf896d192507ad89ba48690b7c467cd8e9c6f4a12be5dc.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0df7fa7dae6d3fd2cbcf896d192507ad89ba48690b7c467cd8e9c6f4a12be5dc.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OutlookTaskRequestBuilder) SingleValueExtendedProperties()(*i90988a76de3564e1bc9e592c2f427a1a8825a2eafa08f4a66b7313a9d5dea2e3.SingleValueExtendedPropertiesRequestBuilder) {
    return i90988a76de3564e1bc9e592c2f427a1a8825a2eafa08f4a66b7313a9d5dea2e3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8be48ea3d303b95811d463c98ddcfa42f165836ce37eaf0634c5eb246c021887.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8be48ea3d303b95811d463c98ddcfa42f165836ce37eaf0634c5eb246c021887.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
