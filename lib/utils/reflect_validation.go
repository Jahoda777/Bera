// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package utils

import (
	"errors"
	"fmt"
	"reflect"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
)

// validateArgumentAndReturnTypes checks if the precompile method argument and return types match
// the abi's argument return types. this is for overloaded Solidity functions and general
// validation.
func ValidateArgumentAndReturnTypes(implMethod reflect.Method, abiMethod abi.Method) error {
	// last parameter of go impl is an error (for reverts), so we don't check it.
	if implMethod.Type.NumOut()-1 != len(abiMethod.Outputs) {
		return errors.New("number of return types mismatch")
	}

	// receiver is 0th param, context is 1st param, so skip those.
	j := 0
	for i := 2; i < implMethod.Type.NumIn(); i++ {
		implMethodParamType := implMethod.Type.In(i)
		abiMethodParamType := abiMethod.Inputs[j].Type.GetType()
		if err := validate(implMethodParamType, abiMethodParamType); err != nil {
			return fmt.Errorf("argument type mismatch: %v != %v", implMethodParamType, abiMethodParamType)
		}
		j++
	}

	for i := 0; i < implMethod.Type.NumOut()-1; i++ {
		implMethodReturnType := implMethod.Type.Out(i)
		abiMethodReturnType := abiMethod.Outputs[i].Type.GetType()
		if err := validate(implMethodReturnType, abiMethodReturnType); err != nil {
			return fmt.Errorf("return type mismatch: %v != %v", implMethodReturnType, abiMethodReturnType)
		}
	}

	return nil
}

// helper function for validateArgumentAndReturnTypes. this function function uses reflection to see
// what types your implementation uses, and checks against the geth representation of the abi types.
func validate(implMethodVarType reflect.Type, abiMethodVarType reflect.Type) error {
	//nolint:exhaustive // nah, this is fine.
	switch implMethodVarType.Kind() {
	// if the type is any, we leave it to the user to make sure that it is used
	// correctly.
	case abiMethodVarType.Kind(), reflect.Interface:
		return nil
	case reflect.Struct:
		if err := validateStructFields(implMethodVarType, abiMethodVarType); err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		for j := 0; j < abiMethodVarType.Len(); j++ {
			// if it is a slice/array of structs, then we need to check if the struct fields match
			if abiMethodVarType.Elem().Kind() == reflect.Struct {
				if err := validateStructFields(
					implMethodVarType.Elem(),
					abiMethodVarType.Elem(),
				); err != nil {
					return err
				}
				// any other case, we just check the elements.
			} else if implMethodVarType.In(j) != abiMethodVarType.In(j) {
				return fmt.Errorf("return type mismatch: %v != %v",
					implMethodVarType.Elem(),
					abiMethodVarType.Elem(),
				)
			}
		}
	default:
		return fmt.Errorf("return type mismatch: %v != %v", implMethodVarType, abiMethodVarType)
	}
	return nil
}

// this function checks to make sure that the struct fields match. if there is a nested struct, then
// we use recursion until we reach the base case of primitive types.
func validateStructFields(implMethodVarType reflect.Type,
	abiMethodVarType reflect.Type,
) error {
	if implMethodVarType == nil && abiMethodVarType == nil {
		return nil
	}
	if implMethodVarType.NumField() != abiMethodVarType.NumField() {
		return errors.New("number of return types mismatch")
	}
	for j := 0; j < implMethodVarType.NumField(); j++ {
		// if the field is a nested struct, then we recurse
		//nolint:gocritic // nah, this is fine.
		if implMethodVarType.Field(j).Type.Kind() == reflect.Struct &&
			abiMethodVarType.Field(j).Type.Kind() == reflect.Struct {
			if err := validateStructFields(
				implMethodVarType.Field(j).Type,
				abiMethodVarType.Field(j).Type,
			); err != nil {
				return err
			}
		} else if implMethodVarType.Field(j).Type != abiMethodVarType.Field(j).Type {
			return fmt.Errorf("return type mismatch: %v != %v",
				implMethodVarType.Field(j).Type,
				abiMethodVarType.Field(j).Type,
			)
		} else {
			continue
		}
	}
	return nil
}
