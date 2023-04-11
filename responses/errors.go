// SPDX-FileCopyrightText: 2023 Sidinggs Media
// SPDX-License-Identifier: MIT

package responses

type GeneralError struct {
  Code int `json:"code"`
  Message string `json:"message"`
}
