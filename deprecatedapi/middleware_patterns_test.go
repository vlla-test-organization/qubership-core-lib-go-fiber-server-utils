package deprecatedapi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPatternWithMethods(t *testing.T) {
	assert := require.New(t)
	patterns := []interface{}{"/deprecated-api/v1/** [POST GET DELETE]"}
	patternsMap := ConvertPatterns(patterns)
	methods := []string{"POST", "GET", "DELETE"}
	assert.Equal(methods, patternsMap["/deprecated-api/v1/**"])
}

func TestPatternWithSpacedMethods(t *testing.T) {
	assert := require.New(t)
	patterns := []interface{}{" /deprecated-api/v1/** [ POST  GET  DELETE OPTIONS ] "}
	patternsMap := ConvertPatterns(patterns)
	methods := []string{"POST", "GET", "DELETE", "OPTIONS"}
	assert.Equal(methods, patternsMap["/deprecated-api/v1/**"])
}

func TestPatternWithEmptyMethods(t *testing.T) {
	assert := require.New(t)
	patterns := []interface{}{" /deprecated-api/v1/** [  ] "}
	patternsMap := ConvertPatterns(patterns)
	methods := []string{"*"}
	assert.Equal(methods, patternsMap["/deprecated-api/v1/**"])
}

func TestPatternWithoutMethods(t *testing.T) {
	assert := require.New(t)
	patterns := []interface{}{" /deprecated-api/v1/** "}
	patternsMap := ConvertPatterns(patterns)
	methods := []string{"*"}
	assert.Equal(methods, patternsMap["/deprecated-api/v1/**"])
}

func TestComplexPatternWithoutMethods(t *testing.T) {
	assert := require.New(t)
	patterns := []interface{}{" /path?/**/path4/{param1}*{param2}*{param3} [GET]"}
	patternsMap := ConvertPatterns(patterns)
	methods := []string{"GET"}
	assert.Equal(methods, patternsMap["/path?/**/path4/{param1}*{param2}*{param3}"])
}
