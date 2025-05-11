package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
 * @Author zyz
 * @Date 2025/5/11 12:31
 **/

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	return number, nil
}
