// Kiebitz - Privacy-Friendly Appointment Scheduling
// Copyright (C) 2021-2021 The Kiebitz Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package definitions

import (
	"github.com/kiebitz-oss/services"
	"github.com/kiebitz-oss/services/cmd"
	"github.com/kiebitz-oss/services/databases"
	"github.com/kiebitz-oss/services/meters"
)

var Default = services.Definitions{
	MeterDefinitions:    meters.Meters,
	DatabaseDefinitions: databases.Databases,
	CommandsDefinitions: cmd.Commands,
}
