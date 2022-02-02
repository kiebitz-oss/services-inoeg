// Kiebitz - Privacy-Friendly Appointment Scheduling
// Copyright (C) 2021-2021 The Kiebitz Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version. Additional terms
// as defined in section 7 of the license (e.g. regarding attribution)
// are specified at https://kiebitz.eu/en/docs/open-source/additional-terms.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package servers

import (
	//"encoding/json"
	"github.com/kiebitz-oss/services"
	//"github.com/kiebitz-oss/services/forms"
	//"time"
)

// The appointments backend acts as an interface between the API and the
// database. It is mostly concerned with ensuring data is propery serialized
// and deserialized when stored or fetched from the database.
type AppointmentsBackend struct {
	db services.Database
}

func (a *AppointmentsBackend) Neighbors(neighborType, zipCode string) *Neighbors {
	return nil
}

func (a *AppointmentsBackend) PriorityToken(name string) *PriorityToken {
	return nil
}

func (a *AppointmentsBackend) Keys(actor string) *Keys {
	return nil
}

func (a *AppointmentsBackend) Codes(actor string) *Codes {
	return nil
}

func (a *AppointmentsBackend) PublicProviderData() *PublicProviderData {
	return nil
}

func (a *AppointmentsBackend) ConfirmedProviderData() *ConfirmedProviderData {
	return nil
}

func (a *AppointmentsBackend) UnverifiedProviderData() *RawProviderData {
	return nil
}

func (a *AppointmentsBackend) VerifiedProviderData() *RawProviderData {
	return nil
}

func (a *AppointmentsBackend) AppointmentsByDate(
	providerID []byte,
	date string,
) *AppointmentsByDate {
	return nil
}

func (a *AppointmentsBackend) AppointmentDatesByID(providerID []byte) *AppointmentDatesByID {
	return nil
}

func (a *AppointmentsBackend) AppointmentDatesByProperty(
	providerID []byte,
	key string,
	value string,
) *AppointmentDatesByProperty {
	return nil
}

func (a *AppointmentsBackend) UsedTokens() *UsedTokens {
	return nil
}

type PriorityToken struct {
}

func (p *PriorityToken) IncrBy(value int64) (int64, error) {
	return 0, nil
}

func (p *PriorityToken) DecrBy(value int64) (int64, error) {
	return 0, nil
}

type Neighbors struct {
	Score int64
	Data  []byte
}

func (n *Neighbors) Add(to string, distance int64) error {
	return nil
}

func (n *Neighbors) Range(from, to int64) ([]Neighbors, error) {
	return nil, nil
}

type Keys struct {
}

func (k *Keys) Set(id []byte, key *services.ActorKey) error {
	return nil
}

func (k *Keys) Get(id []byte) (*services.ActorKey, error) {
	return nil, nil
}

func (k *Keys) GetAll() ([]*services.ActorKey, error) {
	return nil, nil
}

type Codes struct {
}

func (c *Codes) Has(code []byte) (bool, error) {
	return false, nil
}

func (c *Codes) Add(code []byte) error {
	return nil
}

func (c *Codes) Del(code []byte) error {
	return nil
}

func (c *Codes) Score(code []byte) (int64, error) {
	return 0, nil
}

func (c *Codes) AddToScore(code []byte, score int64) error {
	return nil
}

type ConfirmedProviderData struct {
}

func (c *ConfirmedProviderData) Set(providerID []byte, encryptedData *services.ConfirmedProviderData) error {
	return nil
}

func (c *ConfirmedProviderData) Get(providerID []byte) (*services.ConfirmedProviderData, error) {
	return nil, nil
}

type RawProviderData struct {
}

func (c *RawProviderData) Set(providerID []byte, rawData *services.RawProviderData) error {
	return nil
}

func (c *RawProviderData) Del(providerID []byte) error {
	return nil
}

func (c *RawProviderData) Get(providerID []byte) (*services.RawProviderData, error) {
	return nil, nil
}

func (c *RawProviderData) GetAll() (map[string]*services.RawProviderData, error) {
	return nil, nil
}

type UsedTokens struct {
}

func (t *UsedTokens) Del(token []byte) error {
	return nil
}

func (t *UsedTokens) Has(token []byte) (bool, error) {
	return false, nil
}

func (t *UsedTokens) Add(token []byte) error {
	return nil
}

type AppointmentDatesByID struct {
}

func (a *AppointmentDatesByID) GetAll() (map[string][]byte, error) {
	return nil, nil
}

func (a *AppointmentDatesByID) Get(id []byte) (string, error) {
	return "", nil
}

func (a *AppointmentDatesByID) Set(id []byte, date string) error {
	return nil
}

func (a *AppointmentDatesByID) Del(id []byte) error {
	return nil
}

type AppointmentDatesByProperty struct {
}

func (a *AppointmentDatesByProperty) GetAll() (map[string][]byte, error) {
	return nil, nil
}

func (a *AppointmentDatesByProperty) Get(id []byte) (string, error) {
	return "", nil
}

func (a *AppointmentDatesByProperty) Set(appId []byte, date string) error {
	return nil
}

func (a *AppointmentDatesByProperty) Del(id []byte) error {
	return nil
}

type PublicProviderData struct {
}

func (p *PublicProviderData) Get(id []byte) (*services.SignedProviderData, error) {
	return nil, nil
}

func (p *PublicProviderData) Set(id []byte, signedProviderData *services.SignedProviderData) error {
	return nil
}

type AppointmentsByDate struct {
}

func (a *AppointmentsByDate) Del(id []byte) error {
	return nil
}

func (a *AppointmentsByDate) Set(appointment *services.SignedAppointment) error {
	return nil
}

func (a *AppointmentsByDate) Get(validateSettings *services.ValidateSettings, id []byte) (*services.SignedAppointment, error) {
	return nil, nil
}

func (a *AppointmentsByDate) GetAll(validateSettings *services.ValidateSettings) (map[string]*services.SignedAppointment, error) {
	return nil, nil
}
