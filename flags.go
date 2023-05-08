package args

import (
	"log"
	"strconv"
)

func (a *Args) FlagString(key string, required bool, defaultValue string, alternativeKeys ...string) string {
	value := a.getFlagByKeys(append([]string{key}, alternativeKeys...))
	if value == "" && required {
		log.Fatalf("flag: %s is required\n", key)
	} else if value == "" {
		return defaultValue
	}
	return value
}

func (a *Args) FlagInt(key string, required bool, defaultValue int, alternativeKeys ...string) int {
	value := a.getFlagByKeys(append([]string{key}, alternativeKeys...))
	if value == "" && required {
		log.Fatalf("flag: <%s> is required\n", key)
	} else if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Cannot convert value: [%s] to int (flag: <%s>)\n", value, key)
	}
	return intValue
}

func (a *Args) FlagUint(key string, required bool, defaultValue uint, alternativeKeys ...string) uint {
	value := a.getFlagByKeys(append([]string{key}, alternativeKeys...))
	if value == "" && required {
		log.Fatalf("flag: <%s> is required\n", key)
	} else if value == "" {
		return defaultValue
	}
	uintValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert value: [%s] to uint (flag: <%s>)\n", value, key)
	}
	return uint(uintValue)
}

func (a *Args) FlagFloat(key string, required bool, defaultValue float64, alternativeKeys ...string) float64 {
	value := a.getFlagByKeys(append([]string{key}, alternativeKeys...))
	if value == "" && required {
		log.Fatalf("flag: <%s> is required\n", key)
	} else if value == "" {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("Cannot convert value: [%s] to float64 (flay: <%s>)\n", value, key)
	}
	return floatValue
}

func (a *Args) FlagBool(key string, alternativeKeys ...string) bool {
	for _, k := range append([]string{key}, alternativeKeys...) {
		_, ok := a.Flags[k]
		if ok {
			return true
		}
	}
	return false
}

func (a *Args) getFlagByKeys(keys []string) string {
	for _, k := range keys {
		v, ok := a.Flags[k]
		if ok {
			if v == "" {
				log.Fatalf("flag: <%s> is provided but has no value\n", k)
			}
			return v
		}
	}
	return ""
}
