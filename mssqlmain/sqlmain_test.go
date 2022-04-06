package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		user       string
		password   string
		server     string
		database   string
		connString string
		port       int
	}{
		{"", "\n", " ", ";", "..", 1999},
		{".", "\t", ",", "?", "!!", 10001},
		{"\t", "NaN\null\n\n", "NaN\null\n", "\ntrue\n", "\t\n", 23409},
		{"Data for test", "Number 99999 to data test", ")([]$%", "Number 1234567890000", "\n&*@#", 9876},
		{"Yes, no", "No, or, yes, _, ops", "NULL, true", "false, or, yes", "-, ops", 65535},
	}

	var prevuser string
	for _, test := range tests {
		if test.user != prevuser {
			fmt.Printf("\n%s\n", test.user)
			prevuser = test.user
		}
	}

	var prevpassword string
	for _, test := range tests {
		if test.password != prevpassword {
			fmt.Printf("\n%s\n", test.password)
			prevpassword = test.password
		}
	}

	var prevserver string
	for _, test := range tests {
		if test.server != prevserver {
			fmt.Printf("\n%s\n", test.server)
			prevserver = test.server
		}
	}

	var prevdatabase string
	for _, test := range tests {
		if test.database != prevdatabase {
			fmt.Printf("\n%s\n", test.database)
			prevdatabase = test.database
		}
	}

	var prevconnString string
	for _, test := range tests {
		if test.connString != prevconnString {
			fmt.Printf("\n%s\n", test.connString)
			prevconnString = test.connString
		}
	}

	var prevport int
	for _, test := range tests {
		if test.port != prevport {
			fmt.Printf("\n%b\n", test.port)
			prevport = test.port
		}
	}
}
