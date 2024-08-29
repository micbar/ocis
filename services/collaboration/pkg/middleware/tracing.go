package middleware

import (
	"net/http"

	"github.com/owncloud/ocis/v2/services/collaboration/pkg/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// CollaborationTracingMiddleware adds a new middleware in order to include
// more attributes in the traced span.
//
// In order not to mess with the expected responses, this middleware won't do
// anything if there is no available WOPI context set in the request (there is
// nothing to report). This means that the WopiContextAuthMiddleware should be
// set before this middleware.
func CollaborationTracingMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wopiContext, err := WopiContextFromCtx(r.Context())
		if err != nil {
			// if we can't get the context, skip this middleware
			next.ServeHTTP(w, r)
		}

		span := trace.SpanFromContext(r.Context())

		wopiMethod := r.Header.Get("X-WOPI-Override")

		wopiFile := wopiContext.FileReference
		//wopiUser := wopiContext.AccessToken

		attrs := []attribute.KeyValue{
			attribute.String("ocis.wopi.sessionid", r.Header.Get("X-WOPI-SessionId")),
			attribute.String("ocis.wopi.method", wopiMethod),
			attribute.String("ocis.wopi.resource.id.storage", wopiFile.GetResourceId().GetStorageId()),
			attribute.String("ocis.wopi.resource.id.opaque", wopiFile.GetResourceId().GetOpaqueId()),
			attribute.String("ocis.wopi.resource.id.space", wopiFile.GetResourceId().GetSpaceId()),
			attribute.String("ocis.wopi.resource.path", wopiFile.GetPath()),
			//attribute.String("ocis.wopi.user.idp", wopiUser.GetIdp()),
			//attribute.String("ocis.wopi.user.opaque", wopiUser.GetOpaqueId()),
			//attribute.String("ocis.wopi.user.type", wopiUser.GetType().String()),
		}
		span.SetAttributes(attrs...)

		next.ServeHTTP(w, r)
	})
}
