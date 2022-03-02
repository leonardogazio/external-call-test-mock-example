package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetFoo(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	res, err := testInstance.GetFoo(context.Background(), &emptypb.Empty{})
	assert.Nil(err)
	require.NotNil(res)
	assert.Equal(res.GetFoo(), "mocked bar response")
}
