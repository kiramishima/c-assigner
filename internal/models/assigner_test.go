package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssigner_Assign(t *testing.T) {
	assigner := Assigner{}

	t.Run("Eval 3000", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(3000)
		assert.NoError(t, err)
		assert.Equal(t, inv300, int32(10))
		assert.Equal(t, inv500, int32(0))
		assert.Equal(t, inv700, int32(0))
	})

	t.Run("Eval 6700", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(6700)
		assert.NoError(t, err)
		assert.Equal(t, inv300, int32(20))
		assert.Equal(t, inv500, int32(0))
		assert.Equal(t, inv700, int32(1))
	})

	t.Run("Eval 1100", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(1100)
		assert.NoError(t, err)
		assert.Equal(t, inv300, int32(2))
		assert.Equal(t, inv500, int32(1))
		assert.Equal(t, inv700, int32(0))
	})

	t.Run("Eval 5100", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(5100)
		assert.NoError(t, err)
		assert.Equal(t, inv300, int32(17))
		assert.Equal(t, inv500, int32(0))
		assert.Equal(t, inv700, int32(0))
	})

	t.Run("Eval 50", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(50)
		assert.Error(t, err)
		assert.Equal(t, inv300, int32(0))
		assert.Equal(t, inv500, int32(0))
		assert.Equal(t, inv700, int32(0))
	})

	t.Run("Eval 400", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(400)
		assert.Error(t, err)
		assert.Equal(t, inv300, int32(0))
		assert.Equal(t, inv500, int32(0))
		assert.Equal(t, inv700, int32(0))
	})

	t.Run("Eval -3300", func(t *testing.T) {
		inv300, inv500, inv700, err := assigner.Assign(-3300)
		assert.Error(t, err)
		assert.Equal(t, inv300, int32(0))
		assert.Equal(t, inv500, int32(0))
		assert.Equal(t, inv700, int32(0))
	})
}
