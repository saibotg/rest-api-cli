package nagios

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func CheckValue(rule string, x int) (bool, error) {
	neg_inf := false
	pos_inf := false
	start := 0
	end := 0
	outside := true

	//create regex to check the rule
	rangeRegex, err := regexp.Compile(`^(@)?(-?\d*|~)(:?)(-?\d*)$`)
	if err != nil {
		return false, errors.New("invalid regex expression")
	}

	//check the rule with regex
	matches := rangeRegex.MatchString(rule)
	if !matches {
		return false, errors.New("invalid range definition")
	}

	if strings.HasPrefix(rule, "@") {
		outside = false
		rule = rule[1:]
	}

	//split the rule by ":"
	splitString := strings.Split(rule, ":")
	switch lenParts := len(splitString); lenParts {
	case 2:
		startSide := splitString[0]
		endSide := splitString[1]

		if startSide == "~" {
			neg_inf = true
		} else {
			neg_inf = false
			start, err = strconv.Atoi(startSide)
			if err != nil {
				return false, errors.New("invalid range definition")
			}
		}

		if endSide == "" {
			pos_inf = true
		} else {
			pos_inf = false
			end, err = strconv.Atoi(endSide)
			if err != nil {
				return false, errors.New("invalid range definition")
			}
		}
	case 1:
		endSide := splitString[0]
		end, err = strconv.Atoi(endSide)
		if err != nil {
			return false, errors.New("invalid range definition")
		}
		pos_inf = false
	default:
		return false, errors.New("invalid regex expression")
	}

	if outside {
		if !neg_inf && x < start {
			return true, nil
		}

		if !pos_inf && x > end {
			return true, nil
		}
	} else {
		if pos_inf {
			return (x >= start), nil
		}

		if neg_inf {
			return (x <= end), nil
		}

		if x >= start && x <= end {
			return true, nil
		}
	}

	return false, nil
}
