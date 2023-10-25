package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder provides operations to call the getByUserIdAndRole method.
type VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters invoke function getByUserIdAndRole
type VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters
}
// NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal instantiates a new GetByUserIdAndRoleWithUserIdWithRoleRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, role *string, userId *string)(*VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    m := &VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/getByUserIdAndRole(userId='{userId}',role='{role}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if role != nil {
        m.BaseRequestBuilder.PathParameters["role"] = *role
    }
    if userId != nil {
        m.BaseRequestBuilder.PathParameters["userId"] = *userId
    }
    return m
}
// NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder instantiates a new GetByUserIdAndRoleWithUserIdWithRoleRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getByUserIdAndRole
// Deprecated: This method is obsolete. Use GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse instead.
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseable), nil
}
// GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse invoke function getByUserIdAndRole
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable), nil
}
// ToGetRequestInformation invoke function getByUserIdAndRole
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    return NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
