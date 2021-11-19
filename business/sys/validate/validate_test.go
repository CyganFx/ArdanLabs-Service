package validate_test

import (
	"github.com/CyganFx/ArdanLabs-Service/business/data/product"
	"github.com/CyganFx/ArdanLabs-Service/business/sys/validate"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestInterface interface {
	void()
}

type testStruct struct {
}

func (t *testStruct) void() {}

func TestCheck(t *testing.T) {
	np := product.NewProduct{
		Name:     "Comic Books",
		Cost:     10,
		Quantity: 55,
		UserID:   "5cf37266-3473-4006-984f-9325122678b7",
	}

	err := validate.Check(np)
	require.NoError(t, err)

	var testInterface TestInterface
	err = validate.Check(testInterface)
	require.Error(t, err)

	var testStruct testStruct
	err = validate.Check(testStruct)
	require.NoError(t, err)
}

func TestCheckID(t *testing.T) {
	id := validate.GenerateID() // smth like this: "123e4567-e89b-12d3-a456-426614174000"

	err := validate.CheckID(id)
	require.NoError(t, err)

	idWithoutDashes := "123e4567e89b12d3a456426614174000"
	err = validate.CheckID(idWithoutDashes)
	require.NoError(t, err)

	idWithInvalidLength := "123"
	err = validate.CheckID(idWithInvalidLength)
	require.Error(t, err)
}
