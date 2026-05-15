// Package transformers is a stub for the upstream Enterprise transformer
// definitions. The OSS distribution ships an empty extras list so the
// transformers-service compiles without the proprietary set.
package transformers

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

// Transformers is the empty extras slice. The OSS transformers-service
// appends this to the public system transformer list.
var Transformers = []*mgmtv1alpha1.SystemTransformer{}
