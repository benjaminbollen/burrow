// Copyright 2015, 2016 Eris Industries (UK) Ltd.
// This file is part of Eris-RT

// Eris-RT is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Eris-RT is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Eris-RT.  If not, see <http://www.gnu.org/licenses/>.

package geth

import (
	"fmt"

	version "github.com/eris-ltd/eris-db/version"
)

const (
	// Client identifier to advertise over the network
	gethClientIdentifier = "geth"
	// Major version component of the current release
	gethVersionMajor     = 1
	// Minor version component of the current release
	gethVersionMinor     = 4
	// Patch version component of the current release
	gethVersionPatch     = 9
)


// Define the compatible consensus engines this application manager
// is compatible and has been tested with.
var compatibleConsensus = [...]string {
  "tendermint-0.6",
  // "tmsp-0.6",
}

func GetGethVersion() *version.VersionIdentifier {
  return version.New(gethClientIdentifier, gethVersionMajor,
    gethVersionMinor, gethVersionPatch)
}

func AssertCompatibleConsensus(consensusMinorVersion string) error {
  for _, supported := range compatibleConsensus {
    if consensusMinorVersion == supported {
      return nil
    }
  }
  return fmt.Errorf("Geth (%s) is not compatible with consensus engine %s",
    GetGethVersion(), consensusMinorVersion)
}
