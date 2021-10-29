package outlook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0d34dad914e82d03ae5a89b298e4a19b8b8b6bf4eb26a5999888bd525b1fd9c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/supportedtimezoneswithtimezonestandard"
    i360efeb2914ada10b422929cbecf0f7f8fa695d7b1022878848afca7acfd0ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/mastercategories"
    i54cf645f94c231fabc20f9cde8f9bb429da86a32c740ad8c7fcbc3a398e8be24 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups"
    i6913acf90458035b6558bc17364451cf3efd2854e0918e1ebc3a438b11ba8b04 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/supportedtimezones"
    ia668b41f7ac335c1f54bbbbdde0f0ba9b8edb241acbb56fd81a1b8db3fb2ff10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/supportedlanguages"
    ib512bb7334ff30b38cfa6d102f049adfcb6a68103965b326b25efbd5874ed400 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks"
    idbf17811c520d1085f831150de278ee1bb194206c8b3ab498e0d0292a6511ace "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders"
    i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item"
    i74db178f0907a8ff4bbfcfc53a94d5520a643e2563712703082afd071884dd1e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item"
    iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/mastercategories/item"
    iaffe5863c3d1d6c8c2d7abc4fc3d594133e2a40279216bc540d6d0ad80f03059 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item"
)

type OutlookRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OutlookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/outlook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOutlookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OutlookRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutlookRequestBuilder) CreateGetRequestInformation(q func (value *OutlookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OutlookRequestBuilderGetQueryParameters)
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
func (m *OutlookRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutlookRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OutlookRequestBuilder) Get(q func (value *OutlookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser), nil
}
func (m *OutlookRequestBuilder) MasterCategories()(*i360efeb2914ada10b422929cbecf0f7f8fa695d7b1022878848afca7acfd0ddf.MasterCategoriesRequestBuilder) {
    return i360efeb2914ada10b422929cbecf0f7f8fa695d7b1022878848afca7acfd0ddf.NewMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399.OutlookCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory_id"] = id
    }
    return iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399.NewOutlookCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OutlookRequestBuilder) SupportedLanguages()(*ia668b41f7ac335c1f54bbbbdde0f0ba9b8edb241acbb56fd81a1b8db3fb2ff10.SupportedLanguagesRequestBuilder) {
    return ia668b41f7ac335c1f54bbbbdde0f0ba9b8edb241acbb56fd81a1b8db3fb2ff10.NewSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) SupportedTimeZones()(*i6913acf90458035b6558bc17364451cf3efd2854e0918e1ebc3a438b11ba8b04.SupportedTimeZonesRequestBuilder) {
    return i6913acf90458035b6558bc17364451cf3efd2854e0918e1ebc3a438b11ba8b04.NewSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*i0d34dad914e82d03ae5a89b298e4a19b8b8b6bf4eb26a5999888bd525b1fd9c5.SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return i0d34dad914e82d03ae5a89b298e4a19b8b8b6bf4eb26a5999888bd525b1fd9c5.NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
func (m *OutlookRequestBuilder) TaskFolders()(*idbf17811c520d1085f831150de278ee1bb194206c8b3ab498e0d0292a6511ace.TaskFoldersRequestBuilder) {
    return idbf17811c520d1085f831150de278ee1bb194206c8b3ab498e0d0292a6511ace.NewTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskFoldersById(id string)(*i74db178f0907a8ff4bbfcfc53a94d5520a643e2563712703082afd071884dd1e.OutlookTaskFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder_id"] = id
    }
    return i74db178f0907a8ff4bbfcfc53a94d5520a643e2563712703082afd071884dd1e.NewOutlookTaskFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskGroups()(*i54cf645f94c231fabc20f9cde8f9bb429da86a32c740ad8c7fcbc3a398e8be24.TaskGroupsRequestBuilder) {
    return i54cf645f94c231fabc20f9cde8f9bb429da86a32c740ad8c7fcbc3a398e8be24.NewTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskGroupsById(id string)(*iaffe5863c3d1d6c8c2d7abc4fc3d594133e2a40279216bc540d6d0ad80f03059.OutlookTaskGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup_id"] = id
    }
    return iaffe5863c3d1d6c8c2d7abc4fc3d594133e2a40279216bc540d6d0ad80f03059.NewOutlookTaskGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) Tasks()(*ib512bb7334ff30b38cfa6d102f049adfcb6a68103965b326b25efbd5874ed400.TasksRequestBuilder) {
    return ib512bb7334ff30b38cfa6d102f049adfcb6a68103965b326b25efbd5874ed400.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TasksById(id string)(*i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}