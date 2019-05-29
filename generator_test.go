package pgc

import (
	"testing"
)

func TestGenerateCode(t *testing.T) {
	GenerateGoCode(false)
	GenerateJavaCode(true)
}
