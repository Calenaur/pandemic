package response

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Default struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code        uint16 `json:"code"`
	Description string `json:"description"`
}

// Unknown Error
var UNKNOWN_ERROR uint16 = 2999

var DEBUG = false
var localError uint16

func MessageHandler(err error, message string, e echo.Context) error {
	errorDict := make(map[uint16]string)
	errorDict[1062] = "Duplicate entry"
	errorDict[1048] = "Not found"
	errorDict[1452] = "Method cannot be executed"
	if DEBUG {
		fmt.Println("In messageHandler")
	}

	if err != nil {
		localError, errorMessage := getLocalError(err)
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			// If known local error
			if localError != UNKNOWN_ERROR {
				if DEBUG {
					fmt.Println("In local error")
				}

				d := &Default{
					Message: message,
					Error: &Error{
						Code:        uint16(localError),
						Description: errorMessage,
					},
				}
				return e.JSON(getStatus(localError), d)
			}
			// Not local or SQL
			if DEBUG {
				fmt.Println("Not local not SQL")
				log.Fatal(err)
			}
			return err
		}
		// Else SQL error
		if DEBUG {
			fmt.Println("In SQL")
		}
		errorCode := me.Number
		d := &Default{
			Message: message,
			Error: &Error{
				Code:        errorCode,
				Description: errorDict[errorCode],
			},
		}
		if val, ok := errorDict[errorCode]; !ok {
			if DEBUG {
				fmt.Println(val)
				// fmt.Println(err.Error())
				log.Fatal(err.Error())
			}
		}
		return e.JSON(getStatus(uint16(errorCode)), d)
	}
	// No error
	if DEBUG {
		fmt.Println("No error")
	}
	d := &Default{
		Message: message,
	}
	return e.JSON(http.StatusOK, d)
}

func getStatus(code uint16) int {
	// TODO Implement all codes
	statusDict := make(map[uint16]int)
	statusDict[1062] = http.StatusForbidden
	statusDict[1048] = http.StatusNotFound
	statusDict[2000] = http.StatusForbidden
	statusDict[2001] = http.StatusForbidden
	statusDict[2010] = http.StatusForbidden
	statusDict[2011] = http.StatusForbidden
	statusDict[2012] = http.StatusForbidden
	statusDict[2100] = http.StatusUnauthorized
	statusDict[2200] = http.StatusNotFound
	statusDict[2201] = http.StatusBadRequest
	statusDict[2300] = http.StatusBadRequest
	statusDict[2400] = http.StatusForbidden
	if val, ok := statusDict[code]; ok {
		return val
	}
	return http.StatusInternalServerError
}

func getLocalError(err error) (uint16, string) {
	if DEBUG {
		fmt.Println(err.Error())
	}
	if err.Error() == "Password length must be between 8 and 64 characters" {
		return 2010, err.Error()
	} else if err.Error() == "Username length must be between 2 and 16 characters" {
		return 2000, err.Error()
	} else if err.Error() == "Username can not have special characters" {
		return 2001, err.Error()
	} else if strings.Contains(err.Error(), "upper case") {
		return 2011, err.Error()
	} else if strings.Contains(err.Error(), "numeric") {
		return 2012, err.Error()
	} else if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
		return 2100, "Password and hashed password mismatch"
	} else if err.Error() == "sql: no rows in result set" {
		return 2200, "User does not exist"
	} else if err.Error() == "Method failed" {
		return 2201, err.Error()
	} else if strings.Contains(err.Error(), "invalid syntax") {
		slicedError := (strings.SplitAfterN(strings.ReplaceAll(err.Error(), "\"", "'"), ": ", 2)[1])
		return 2300, formatError(slicedError)
	} else if err.Error() == "Restricted access" {
		return 2400, err.Error()
	}

	return 2999, "Unknown error"

}

func formatError(str string) string {
	return strings.SplitAfterN(strings.Title(str), " ", 2)[0] + strings.SplitAfterN(str, " ", 2)[1]
}
