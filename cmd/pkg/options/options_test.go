package options

import (
	"testing"

	testv1 "github.com/clly/proto-telemetry/test/open-telemetry/proto/gen/test/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protodesc"
)

func Test_GetTelemetryExcludeFile(t *testing.T) {
	// File_test_v1_exclude_proto

	require.True(t, GetTelemetryExcludeFile(protodesc.ToFileDescriptorProto(testv1.File_test_v1_exclude_proto)))
	require.False(t, GetTelemetryExcludeFile(protodesc.ToFileDescriptorProto(testv1.File_test_v1_test_proto)))
}
