package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder provides operations to call the importDeviceIdentityList method.
type ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderInternal instantiates a new ImportDeviceIdentityListRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) {
    m := &ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedDeviceIdentities/importDeviceIdentityList", pathParameters),
    }
    return m
}
// NewImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder instantiates a new ImportDeviceIdentityListRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importDeviceIdentityList
// Deprecated: This method is obsolete. Use PostAsImportDeviceIdentityListPostResponse instead.
func (m *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) Post(ctx context.Context, body ImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(ImportedDeviceIdentitiesImportDeviceIdentityListResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImportedDeviceIdentitiesImportDeviceIdentityListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedDeviceIdentitiesImportDeviceIdentityListResponseable), nil
}
// PostAsImportDeviceIdentityListPostResponse invoke action importDeviceIdentityList
func (m *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) PostAsImportDeviceIdentityListPostResponse(ctx context.Context, body ImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(ImportedDeviceIdentitiesImportDeviceIdentityListPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImportedDeviceIdentitiesImportDeviceIdentityListPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedDeviceIdentitiesImportDeviceIdentityListPostResponseable), nil
}
// ToPostRequestInformation invoke action importDeviceIdentityList
func (m *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) ToPostRequestInformation(ctx context.Context, body ImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) WithUrl(rawUrl string)(*ImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder) {
    return NewImportedDeviceIdentitiesImportDeviceIdentityListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
