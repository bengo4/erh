// Package erh provides simple errror handling functions.
//
// When an error returned from a nested function, it may be difficult for you to find the birth place or the pathway of the error.
//
// To solve this problem, the erh package provides functions below, similar to [pkg/errors].
//
//   - [Wrap]
//   - [Cause]
//   - [Errorf]
//
// Also erh provides function [As], a wrapper of [errors.As].
package erh
