package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceAccessProfilesItemAssignRequestBuilder provides operations to call the assign method.
type ResourceAccessProfilesItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResourceAccessProfilesItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceAccessProfilesItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewResourceAccessProfilesItemAssignRequestBuilderInternal instantiates a new AssignRequestBuilder and sets the default values.
func NewResourceAccessProfilesItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceAccessProfilesItemAssignRequestBuilder) {
    m := &ResourceAccessProfilesItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/resourceAccessProfiles/{deviceManagementResourceAccessProfileBase%2Did}/assign", pathParameters),
    }
    return m
}
// NewResourceAccessProfilesItemAssignRequestBuilder instantiates a new AssignRequestBuilder and sets the default values.
func NewResourceAccessProfilesItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceAccessProfilesItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceAccessProfilesItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// Deprecated: This method is obsolete. Use PostAsAssignPostResponse instead.
func (m *ResourceAccessProfilesItemAssignRequestBuilder) Post(ctx context.Context, body ResourceAccessProfilesItemAssignPostRequestBodyable, requestConfiguration *ResourceAccessProfilesItemAssignRequestBuilderPostRequestConfiguration)(ResourceAccessProfilesItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResourceAccessProfilesItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceAccessProfilesItemAssignResponseable), nil
}
// PostAsAssignPostResponse invoke action assign
func (m *ResourceAccessProfilesItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body ResourceAccessProfilesItemAssignPostRequestBodyable, requestConfiguration *ResourceAccessProfilesItemAssignRequestBuilderPostRequestConfiguration)(ResourceAccessProfilesItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResourceAccessProfilesItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceAccessProfilesItemAssignPostResponseable), nil
}
// ToPostRequestInformation invoke action assign
func (m *ResourceAccessProfilesItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body ResourceAccessProfilesItemAssignPostRequestBodyable, requestConfiguration *ResourceAccessProfilesItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ResourceAccessProfilesItemAssignRequestBuilder) WithUrl(rawUrl string)(*ResourceAccessProfilesItemAssignRequestBuilder) {
    return NewResourceAccessProfilesItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
